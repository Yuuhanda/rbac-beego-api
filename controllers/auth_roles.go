package controllers

import (
	"encoding/json"
	"rbac-beego-api/models"
	"rbac-beego-api/services"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type AuthRolesController struct {
	web.Controller
	roleService *services.AuthRolesService
}

func (c *AuthRolesController) Prepare() {
	c.roleService = services.NewAuthRolesService()
}

func (c *AuthRolesController) Create() {
	var role models.AuthRoles
	if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&role); err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Invalid JSON data",
			"error":   err.Error(),
		}
		c.ServeJSON()
		return
	}

	if err := c.roleService.Create(&role); err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Failed to create role",
			"error":   err.Error(),
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"message": "Role created successfully",
			"data":    role,
		}
	}
	c.ServeJSON()
}

func (c *AuthRolesController) Get() {
	code := strings.TrimSpace(c.Ctx.Input.Param(":code"))
	if code == "" {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Role code is required",
		}
		c.ServeJSON()
		return
	}

	role, err := c.roleService.GetByID(code)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Role not found",
			"error":   err.Error(),
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"data":    role,
		}
	}
	c.ServeJSON()
}

func (c *AuthRolesController) List() {
	page, _ := c.GetInt("page", 1)
	pageSize, _ := c.GetInt("page_size", 10)

	roles, total, err := c.roleService.List(page, pageSize)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Failed to retrieve roles",
			"error":   err.Error(),
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"data": map[string]interface{}{
				"roles": roles,
				"total": total,
				"page":  page,
			},
		}
	}
	c.ServeJSON()
}

func (c *AuthRolesController) Update() {
	code := strings.TrimSpace(c.Ctx.Input.Param(":code"))
	if code == "" {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Role code is required",
		}
		c.ServeJSON()
		return
	}

	var role models.AuthRoles
	if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&role); err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Invalid JSON data",
			"error":   err.Error(),
		}
		c.ServeJSON()
		return
	}

	role.Code = code
	if err := c.roleService.Update(&role); err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Failed to create role",
			"error":   err.Error(),
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"message": "Role updated successfully",
			"data":    role,
		}
	}
	c.ServeJSON()
}

func (c *AuthRolesController) Delete() {
	code := strings.TrimSpace(c.Ctx.Input.Param(":code"))
	if code == "" {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Role code is required",
		}
		c.ServeJSON()
		return
	}

	// Check if role exists first
	role, err := c.roleService.GetByID(code)
	if err != nil || role == nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "No roles associated with role code found",
		}
		c.ServeJSON()
		return
	}

	if err := c.roleService.Delete(code); err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Failed to delete role",
			"error":   err.Error(),
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"message": "Role deleted successfully",
		}
	}
	c.ServeJSON()
}
