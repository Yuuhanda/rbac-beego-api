package controllers

import (
    "github.com/beego/beego/v2/server/web"
    "rbac-beego-api/models"
    "rbac-beego-api/services"
	"fmt"
    "golang.org/x/crypto/bcrypt"
    "encoding/json"
    "io"
    "time"
    "strings"
)

type UserController struct {
    web.Controller
    userService *services.UserService
}

func (c *UserController) Prepare() {
    c.userService = services.NewUserService()
}

// CreateUser handles user creation
// @router /user [post]
func (c *UserController) CreateUser() {
    var userForm struct {
        Username string `json:"username"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    // Read the raw request body
    body, _ := io.ReadAll(c.Ctx.Request.Body)
    
    if err := json.Unmarshal(body, &userForm); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid JSON data",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userForm.Password), bcrypt.DefaultCost)

    user := &models.User{
        Username:       userForm.Username,
        Email:         userForm.Email,
        PasswordHash:  string(hashedPassword),
        Status:        1,
        CreatedAt:     time.Now(),
        UpdatedAt:     time.Now(),
        RegistrationIP: c.Ctx.Input.IP(),
    }

    if err := c.userService.Create(user); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to create user",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "User created successfully",
            "data":    user,
        }
    }
    c.ServeJSON()
}



// GetUser retrieves user by ID
// @router /user/:id [get]
func (c *UserController) GetUser() {
    idStr := c.Ctx.Input.Param(":id")
    var id int
    if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid user ID",
        }
        c.ServeJSON()
        return
    }

    user, err := c.userService.GetByID(id)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "User not found",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "data":    user,
        }
    }
    c.ServeJSON()
}

// ListUsers retrieves paginated users
// @router /users [get]
func (c *UserController) ListUsers() {
    page, _ := c.GetInt("page", 1)
    pageSize, _ := c.GetInt("pageSize", 10)

    users, total, err := c.userService.List(page, pageSize)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to retrieve users",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "data": map[string]interface{}{
                "users": users,
                "total": total,
                "page":  page,
            },
        }
    }
    c.ServeJSON()
}

// UpdateUser updates user information
// @router /user/:id [put]
func (c *UserController) UpdateUser() {
    idStr := c.Ctx.Input.Param(":id")
    var id int
    if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid user ID",
        }
        c.ServeJSON()
        return
    }

    var updateForm struct {
        Username string `json:"username"`
        Email    string `json:"email"`
        Password string `json:"password"`
        Status   int    `json:"status"`
    }

    body, _ := io.ReadAll(c.Ctx.Request.Body)
    if err := json.Unmarshal(body, &updateForm); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid form data",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    user, err := c.userService.GetByID(id)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "User not found",
        }
        c.ServeJSON()
        return
    }

    // Update fields if provided
    if updateForm.Username != "" {
        user.Username = updateForm.Username
    }
    if updateForm.Email != "" {
        user.Email = updateForm.Email
    }
    if updateForm.Status != 0 {
        user.Status = updateForm.Status
    }
    
    // Update password if provided
    if updateForm.Password != "" {
        hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(updateForm.Password), bcrypt.DefaultCost)
        user.PasswordHash = string(hashedPassword)
    }

    user.UpdatedAt = time.Now()

    if err := c.userService.Update(user); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to update user",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "User updated successfully",
            "data":    user,
        }
    }
    c.ServeJSON()
}

// DeleteUser deletes a user
// @router /user/:id [delete]
func (c *UserController) DeleteUser() {
    idStr := c.Ctx.Input.Param(":id")
    var id int
    if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid user ID",
        }
        c.ServeJSON()
        return
    }

    if err := c.userService.Delete(id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to delete user",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "User deleted successfully",
        }
    }
    c.ServeJSON()
}

func GetUserSystemLanguage(c *web.Controller) string {
    // Get language from browser's Accept-Language header
    acceptLanguage := c.Ctx.Input.Header("Accept-Language")
    
    // Parse the Accept-Language header which looks like: "en-US,en;q=0.9,es;q=0.8"
    if len(acceptLanguage) > 0 {
        // Split by comma to get primary language
        languages := strings.Split(acceptLanguage, ",")
        if len(languages) > 0 {
            // Get first preferred language and trim any quality values
            primaryLang := strings.Split(languages[0], ";")[0]
            return primaryLang
        }
    }
    
    return "en" // Default fallback language
}
