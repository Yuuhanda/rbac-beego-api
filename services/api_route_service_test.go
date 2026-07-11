package services

import "testing"

func TestRouteMatchesByPathAndMethod(t *testing.T) {
	existingRoute := map[string]string{"path": "/users", "method": "GET"}
	currentRoute := map[string]string{"path": "/users", "method": "GET", "controller": "UserController", "action": "List"}

	if !routeMatches(existingRoute, currentRoute) {
		t.Fatalf("expected routes with the same path and method to match")
	}
}

func TestRouteMatchesFailsForDifferentPathOrMethod(t *testing.T) {
	existingRoute := map[string]string{"path": "/users", "method": "GET"}
	currentRoute := map[string]string{"path": "/users", "method": "POST"}

	if routeMatches(existingRoute, currentRoute) {
		t.Fatalf("expected routes with different path or method to not match")
	}
}
