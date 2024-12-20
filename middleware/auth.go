package middleware

import (
    "github.com/beego/beego/v2/server/web/context"
    "github.com/beego/beego/v2/server/web"
    "rbac-beego-api/services"
)

func AuthMiddleware() web.FilterFunc {
    return func(ctx *context.Context) {
        token := ctx.Input.Header("Authorization")
        if token == "" {
            ctx.Output.JSON(map[string]interface{}{
                "success": false,
                "message": "Unauthorized access",
            }, true, false)
            return
        }

        authService := services.NewAuthService()
        user, err := authService.GetUserFromToken(token)
        if err != nil {
            ctx.Output.JSON(map[string]interface{}{
                "success": false,
                "message": "Invalid token",
            }, true, false)
            return
        }

        // Grant immediate access if user is superadmin
        if user.Superadmin == 1 {
            ctx.Input.SetData("user", user)
            return
        }

        // Continue with regular role checking for non-superadmin users
        roleUserService := services.NewAuthRolesUserService()
        roles, err := roleUserService.GetRolesByUserId(user.Id)
        if err != nil {
            ctx.Output.JSON(map[string]interface{}{
                "success": false,
                "message": "Failed to get user roles",
            }, true, false)
            return
        }

        currentPath := ctx.Input.URL()
        currentMethod := ctx.Input.Method()
        hasPermission := false

        authItemService := services.NewAuthItemService()
        for _, role := range roles {
            permitted, _ := authItemService.CheckPermission(role.Code, currentPath, currentMethod)
            if permitted {
                hasPermission = true
                break
            }
        }

        if !hasPermission {
            ctx.Output.JSON(map[string]interface{}{
                "success": false,
                "message": "Access denied",
            }, true, false)
            return
        }

        ctx.Input.SetData("user", user)
    }
}

