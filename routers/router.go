package routers

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"rbac-beego-api/controllers"
	"rbac-beego-api/database"
	"rbac-beego-api/middleware"
	"rbac-beego-api/services"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	database.GetInstance()
	InitRoutes()
}

func InitRoutes() {
	// Get database instance first
	database.GetInstance()

	// Public routes
	web.Router("/auth/login", &controllers.AuthController{}, "post:Login")

	//RBAC
	// Routes Scanner
	web.Router("/api/routes/scan", &controllers.APIRouteController{}, "post:ScanRoutes")
	web.Router("/api/routes/list", &controllers.APIRouteController{}, "get:ListRoutes")
	web.Router("/api/routes/:id", &controllers.APIRouteController{}, "get:Get;delete:DeleteRoute")

	// Auth Roles Routes
	web.Router("/api/roles", &controllers.AuthRolesController{}, "post:Create")
	web.Router("/api/roles/:code", &controllers.AuthRolesController{}, "get:Get;put:Update;delete:Delete")
	web.Router("/api/roles", &controllers.AuthRolesController{}, "get:List")

	// Auth Roles User Routes
	web.Router("/api/user-roles", &controllers.AuthRolesUserController{}, "post:Create")
	web.Router("/api/user-roles/user/:userId", &controllers.AuthRolesUserController{}, "get:GetUserRoles")
	web.Router("/api/user-roles/role/:roleId", &controllers.AuthRolesUserController{}, "get:GetRoleUsers")
	web.Router("/api/user-roles/:userId/:roleId", &controllers.AuthRolesUserController{}, "delete:Delete")

	// Auth Item Routes
	web.Router("/api/auth-items", &controllers.AuthItemController{}, "post:Create")
	web.Router("/api/auth-items/:id", &controllers.AuthItemController{}, "get:Get")
	web.Router("/api/auth-items", &controllers.AuthItemController{}, "get:List")
	web.Router("/api/auth-items/:id", &controllers.AuthItemController{}, "put:Update")
	web.Router("/api/auth-items/:id", &controllers.AuthItemController{}, "delete:Delete")
	// bulk endpoint removed: use POST /api/auth-items for single create

	// Admin-only route with multiple middleware
	web.InsertFilter("/user", web.BeforeRouter, middleware.AuthMiddleware())
	web.InsertFilter("/user", web.BeforeRouter, middleware.AdminMiddleware())
	web.InsertFilter("/api/routes/*", web.BeforeRouter, middleware.AuthMiddleware())
	web.InsertFilter("/api/routes/*", web.BeforeRouter, middleware.AdminMiddleware())
	web.InsertFilter("/api/roles/*", web.BeforeRouter, middleware.AuthMiddleware())
	web.InsertFilter("/api/roles/*", web.BeforeRouter, middleware.AdminMiddleware())
	web.Router("/user", &controllers.UserController{}, "post:CreateUser")

	// Other protected routes
	web.InsertFilter("/user/*", web.BeforeRouter, middleware.AuthMiddleware())
	web.InsertFilter("/users", web.BeforeRouter, middleware.AuthMiddleware())
	web.InsertFilter("/api/*", web.BeforeRouter, middleware.AuthMiddleware())

	// User Routes
	web.Router("/auth/logout", &controllers.AuthController{}, "post:Logout")
	web.Router("/user/:id", &controllers.UserController{}, "get:GetUser")
	web.Router("/user-update/:id", &controllers.UserController{}, "put:UpdateUser")
	web.Router("/user-delete/:id", &controllers.UserController{}, "delete:DeleteUser")
	web.Router("/users", &controllers.UserController{}, "get:ListUsers")
	web.Router("/user/:id/visits", &controllers.UserVisitLogController{}, "get:GetUserVisits")

}

// scanAndSyncRoutes parses this router file and syncs routes to DB
func scanAndSyncRoutes() error {
	routeData := make([]map[string]string, 0)

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "routers/router.go", nil, parser.ParseComments)
	if err != nil {
		return err
	}

	ast.Inspect(node, func(n ast.Node) bool {
		if funcDecl, ok := n.(*ast.FuncDecl); ok && funcDecl.Name.Name == "InitRoutes" {
			ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
				if callExpr, ok := n.(*ast.CallExpr); ok {
					if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
						if ident, ok := selExpr.X.(*ast.Ident); ok && ident.Name == "web" && selExpr.Sel.Name == "Router" {
							if len(callExpr.Args) >= 3 {
								// extract path
								path := ""
								if lit, ok := callExpr.Args[0].(*ast.BasicLit); ok {
									path = strings.Trim(lit.Value, "\"")
								}
								// extract method: third arg
								methodAction := ""
								if lit, ok := callExpr.Args[2].(*ast.BasicLit); ok {
									methodAction = strings.Trim(lit.Value, "\"")
								}
								parts := strings.Split(methodAction, ":")
								if len(parts) == 2 {
									method := strings.ToUpper(parts[0])
									routeData = append(routeData, map[string]string{
										"path":   path,
										"method": method,
										// controller/action/description can be set by ScanAndSaveRoutes if needed
										"controller":  "",
										"action":      parts[1],
										"description": fmt.Sprintf("API endpoint for %s %s", method, path),
									})
								}
							}
						}
					}
				}
				return true
			})
		}
		return true
	})

	svc := services.NewAPIRouteService()
	return svc.ScanAndSaveRoutes(routeData)
}
