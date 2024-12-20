package controllers

import (
    "github.com/beego/beego/v2/server/web"
    "rbac-beego-api/services"
    "rbac-beego-api/models"
    "encoding/json"
    "strconv"
)

type AuthRolesUserController struct {
    web.Controller
    roleUserService *services.AuthRolesUserService
}

func (c *AuthRolesUserController) Prepare() {
    c.roleUserService = services.NewAuthRolesUserService()
}

func (c *AuthRolesUserController) Create() {
    var payload struct {
        UserId int    `json:"UserId"`
        Code   string `json:"Code"`
    }
    
    if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&payload); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid JSON data",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    roleUser := &models.AuthRolesUser{
        UserId: &models.User{Id: payload.UserId},
        RoleId: &models.AuthRoles{Code: payload.Code},
    }

    if err := c.roleUserService.Create(roleUser); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to assign role to user",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Role assigned to user successfully",
            "data":    roleUser,
        }
    }
    c.ServeJSON()
}


func (c *AuthRolesUserController) GetUserRoles() {
    UserId, _ := c.GetInt(":userId")
    roles, err := c.roleUserService.GetRolesByUserId(UserId)
    
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to get user roles",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "data":    roles,
        }
    }
    c.ServeJSON()
}

func (c *AuthRolesUserController) GetRoleUsers() {
    roleId := c.Ctx.Input.Param(":roleId")
    users, err := c.roleUserService.GetUsersByRoleId(roleId)
    
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to get role users",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "data":    users,
        }
    }
    c.ServeJSON()
}



func (c *AuthRolesUserController) Delete() {
    userId := c.Ctx.Input.Param(":userId")
    userIdInt, _ := strconv.Atoi(userId)

    if err := c.roleUserService.Delete(userIdInt); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to remove role from user",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Role removed from user successfully",
        }
    }
    c.ServeJSON()
}

