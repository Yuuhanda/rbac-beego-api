package services

import (
	"rbac-beego-api/models"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type APIRouteService struct {
	ormer orm.Ormer
}

func NewAPIRouteService() *APIRouteService {
	return &APIRouteService{
		ormer: orm.NewOrm(),
	}
}

func (s *APIRouteService) Create(route *models.ApiRoute) error {
	route.CreatedAt = time.Now()
	route.UpdatedAt = time.Now()
	_, err := s.ormer.Insert(route)
	return err
}

func (s *APIRouteService) List() ([]*models.ApiRoute, error) {
	var routes []*models.ApiRoute
	_, err := s.ormer.QueryTable(new(models.ApiRoute)).All(&routes)
	return routes, err
}

func (s *APIRouteService) ScanAndSaveRoutes(routes []map[string]string) error {
	// Get all existing routes from DB
	var existingRoutes []struct {
		ID          int
		Path        string
		Method      string
		Controller  string
		Action      string
		Description string
	}
	_, err := s.ormer.Raw("SELECT id, path, method, controller, action, description FROM api_route").QueryRows(&existingRoutes)
	if err != nil {
		return err
	}

	// Track which routes still exist in router.go
	currentRouteKeys := make(map[string]bool)

	// Add new routes or update existing ones
	for _, route := range routes {
		path := route["path"]
		method := route["method"]
		controller := route["controller"]
		action := route["action"]
		description := route["description"]
		routeKey := path + "-" + method
		currentRouteKeys[routeKey] = true

		exists, err := s.RouteExists(path, method)
		if err != nil {
			logs.Error("Error checking route existence:", err)
			continue
		}

		if exists {
			_, err := s.ormer.Raw("UPDATE api_route SET controller = ?, action = ?, description = ?, updated_at = NOW() WHERE path = ? AND method = ?",
				controller, action, description, path, method).Exec()
			if err != nil {
				logs.Error("Error updating route:", err)
				continue
			}
			logs.Info("Updated route:", method, path)
			continue
		}

		_, err = s.ormer.Raw("INSERT INTO api_route (path, method, controller, action, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?, NOW(), NOW())",
			path, method, controller, action, description).Exec()
		if err != nil {
			logs.Error("Error saving route:", err)
			continue
		}
		logs.Info("Added new route:", method, path)
	}

	// Delete routes that no longer exist in router.go
	for _, existingRoute := range existingRoutes {
		routeKey := existingRoute.Path + "-" + existingRoute.Method
		if currentRouteKeys[routeKey] {
			continue
		}

		_, err := s.ormer.Raw("DELETE FROM api_route WHERE id = ?", existingRoute.ID).Exec()
		if err != nil {
			logs.Error("Error deleting route:", err)
			continue
		}
		logs.Info("Deleted removed route:", existingRoute.Method, existingRoute.Path)
	}

	return nil
}

func routeMatches(existingRoute map[string]string, currentRoute map[string]string) bool {
	if existingRoute == nil || currentRoute == nil {
		return false
	}

	return existingRoute["path"] == currentRoute["path"] && existingRoute["method"] == currentRoute["method"]
}

func (s *APIRouteService) GetRoute(id int) (*models.ApiRoute, error) {
	route := &models.ApiRoute{Id: id}
	err := s.ormer.Read(route)
	return route, err
}

func (s *APIRouteService) RouteExists(path, method string) (bool, error) {
	var count int
	err := s.ormer.Raw("SELECT COUNT(*) FROM api_route WHERE path = ? AND method = ?", path, method).QueryRow(&count)
	return count > 0, err
}

// RouteExist checks whether any route with the given path exists (ignoring method)
func (s *APIRouteService) RouteExist(path string) (bool, error) {
	var count int
	err := s.ormer.Raw("SELECT COUNT(*) FROM api_route WHERE path = ?", path).QueryRow(&count)
	return count > 0, err
}

func (s *APIRouteService) UpdateRoute(id int, routeData map[string]string) error {
	route := &models.ApiRoute{Id: id}
	if err := s.ormer.Read(route); err != nil {
		logs.Error("Error reading route before update:", err)
		return err
	}

	if path, ok := routeData["path"]; ok && path != "" {
		route.Path = path
	}
	if method, ok := routeData["method"]; ok && method != "" {
		route.Method = strings.ToUpper(method)
	}
	if controller, ok := routeData["controller"]; ok {
		route.Controller = controller
	}
	if action, ok := routeData["action"]; ok {
		route.Action = action
	}
	if description, ok := routeData["description"]; ok {
		route.Description = description
	}

	route.UpdatedAt = time.Now()
	_, err := s.ormer.Update(route)
	if err != nil {
		logs.Error("Error updating route:", err)
		return err
	}

	logs.Info("Successfully updated route with ID:", id)
	return nil
}

func (s *APIRouteService) DeleteRoute(id int) error {
	_, err := s.ormer.Raw("DELETE FROM api_route WHERE id = ?", id).Exec()
	if err != nil {
		logs.Error("Error deleting route:", err)
		return err
	}
	logs.Info("Successfully deleted route with ID:", id)
	return nil
}
