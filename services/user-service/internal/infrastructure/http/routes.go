package http

import "github.com/gin-gonic/gin"

// Route defines a single route with its handler and method
type Route struct {
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
}

// RegisterRoutes is a helper to map all routes to the Gin router
func RegisterRoutes(router *gin.Engine, routes []Route) {
	for _, route := range routes {
		switch route.Method {
		case "GET":
			router.GET(route.Path, route.HandlerFunc)
		case "POST":
			router.POST(route.Path, route.HandlerFunc)
		case "PUT":
			router.PUT(route.Path, route.HandlerFunc)
		case "DELETE":
			router.DELETE(route.Path, route.HandlerFunc)
		}
	}
}
func CollectRoutes(router *gin.Engine, handlerMap map[string]interface{}) []Route {
	allRoutes := []Route{}

	// Example for adding routes of User module
	if handler, ok := handlerMap["UserHandler"]; ok {
		allRoutes = append(allRoutes, UserRoutes(handler.(*UserHandler))...)
	}

	// Register all routes to the router
	RegisterRoutes(router, allRoutes)

	return allRoutes
}
