package controllers

import (
	"encoding/json"
	"rbac-beego-api/models"
	"rbac-beego-api/services"

	"github.com/beego/beego/v2/server/web"
)

type AuthItemController struct {
	web.Controller
	authItemService *services.AuthItemService
}

func (c *AuthItemController) Prepare() {
	c.authItemService = services.NewAuthItemService()
}

func (c *AuthItemController) Create() {
	var payload struct {
		Role   string `json:"Role"`
		Path   string `json:"Path"`
		Method string `json:"Method"`
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

	// validate route exists for path+method
	apiSvc := services.NewAPIRouteService()
	exists, err := apiSvc.RouteExists(payload.Path, payload.Method)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"success": false, "message": "Failed to validate route", "error": err.Error()}
		c.ServeJSON()
		return
	}
	if !exists {
		c.Data["json"] = map[string]interface{}{"success": false, "message": "Route not found"}
		c.ServeJSON()
		return
	}

	authItem := models.AuthItem{Role: payload.Role, Path: payload.Path}

	if err := c.authItemService.Create(&authItem); err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Failed to create auth item",
			"error":   err.Error(),
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"message": "Auth item created successfully",
			"data":    authItem,
		}
	}
	c.ServeJSON()
}

func (c *AuthItemController) Get() {
	id, _ := c.GetInt(":id")
	authItem, err := c.authItemService.GetByID(id)

	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Auth item not found",
			"error":   err.Error(),
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"data":    authItem,
		}
	}
	c.ServeJSON()
}

func (c *AuthItemController) List() {
	page, _ := c.GetInt("page", 1)
	pageSize, _ := c.GetInt("page_size", 10)

	authItems, total, err := c.authItemService.List(page, pageSize)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Failed to retrieve auth items",
			"error":   err.Error(),
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"data": map[string]interface{}{
				"items": authItems,
				"total": total,
				"page":  page,
			},
		}
	}
	c.ServeJSON()
}

func (c *AuthItemController) Update() {
	id, _ := c.GetInt(":id")
	var payload struct {
		Role   string `json:"Role"`
		Path   string `json:"Path"`
		Method string `json:"Method"`
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
	// validate route exists for path+method
	apiSvc := services.NewAPIRouteService()
	exists, err := apiSvc.RouteExists(payload.Path, payload.Method)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"success": false, "message": "Failed to validate route", "error": err.Error()}
		c.ServeJSON()
		return
	}
	if !exists {
		c.Data["json"] = map[string]interface{}{"success": false, "message": "Route not found"}
		c.ServeJSON()
		return
	}

	var authItem models.AuthItem
	authItem.Id = id
	authItem.Role = payload.Role
	authItem.Path = payload.Path

	if err := c.authItemService.Update(&authItem); err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Failed to update auth item",
			"error":   err.Error(),
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"message": "Auth item updated successfully",
			"data":    authItem,
		}
	}
	c.ServeJSON()
}

func (c *AuthItemController) Delete() {
	id, _ := c.GetInt(":id")

	err := c.authItemService.Delete(id)
	if err != nil {
		message := err.Error()
		if message == "no auth item found with this id" {
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"message": message,
			}
		} else {
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"message": "Failed to delete auth item",
				"error":   err.Error(),
			}
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"message": "Auth item deleted successfully",
		}
	}
	c.ServeJSON()
}
