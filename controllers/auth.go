package controllers

import (
    "github.com/beego/beego/v2/server/web"
    "rbac-beego-api/services"
	"golang.org/x/crypto/bcrypt"
	"encoding/json"
	"io"
)

type AuthController struct {
    web.Controller
    userService *services.UserService
    visitLogService *services.UserVisitLogService
}

func (c *AuthController) Prepare() {
    c.userService = services.NewUserService()
    c.visitLogService = services.NewUserVisitLogService()
}

// Login handles user authentication
// @router /auth/login [post]
func (c *AuthController) Login() {
    var loginForm struct {
        Email    string `json:"Email"`
        Password string `json:"Password"`
    }

    body, _ := io.ReadAll(c.Ctx.Request.Body)
    if err := json.Unmarshal(body, &loginForm); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid form data",
        }
        c.ServeJSON()
        return
    }

    user, err := c.userService.GetByEmail(loginForm.Email)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid credentials",
        }
        c.ServeJSON()
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginForm.Password)); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid credentials",
        }
        c.ServeJSON()
        return
    }

    // Generate auth token
    if err := c.userService.GenerateAuthToken(user); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to generate auth token",
        }
        c.ServeJSON()
        return
    }

    language := GetUserSystemLanguage(&c.Controller)
    // Log the visit
    c.visitLogService.LogVisit(
        user.Id,
        c.Ctx.Input.IP(),
        c.Ctx.Input.UserAgent(),
        //c.Ctx.Input.Header("Accept-Language"),
        language,
    )

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "message": "Login successful",
        "data": map[string]interface{}{
            "user": user,
            "token": user.AuthKey,
        },
    }

    // Generate auth token
    token := user.AuthKey // Use the generated AuthKey

    // Get user roles
    roleUserService := services.NewAuthRolesUserService()
    roles, _ := roleUserService.GetRolesByUserId(user.Id)

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "message": "Login successful",
        "data": map[string]interface{}{
            "token": token,
            "user":  user,
            "roles": roles,
        },
    }
    c.ServeJSON()
}

// Logout handles user logout
// @router /auth/logout [post]
func (c *AuthController) Logout() {
    // Clear auth token from current user
    authHeader := c.Ctx.Input.Header("Authorization")
    if user, err := c.userService.GetByAuthKey(authHeader); err == nil {
        user.AuthKey = ""
        c.userService.Update(user)
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "message": "Logout successful",
    }
    c.ServeJSON()
}
