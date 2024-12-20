package routers

import (
    "github.com/beego/beego/v2/server/web"
    "rbac-beego-api/controllers"
    "rbac-beego-api/middleware"
    "rbac-beego-api/database"
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
    web.Router("/api/roles/:id", &controllers.AuthRolesController{}, "get:Get")
    web.Router("/api/roles", &controllers.AuthRolesController{}, "get:List")
    web.Router("/api/roles/:id", &controllers.AuthRolesController{}, "put:Update")
    web.Router("/api/roles/:id", &controllers.AuthRolesController{}, "delete:Delete")
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
    web.Router("/api/auth-items/bulk", &controllers.AuthItemController{}, "post:CreateBulk")

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
    web.Router("/user/:id", &controllers.UserController{}, "put:UpdateUser")
    web.Router("/user/:id", &controllers.UserController{}, "delete:DeleteUser")
    web.Router("/users", &controllers.UserController{}, "get:ListUsers")
    web.Router("/user/:id/visits", &controllers.UserVisitLogController{}, "get:GetUserVisits")


}