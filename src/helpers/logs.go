package helpers

import (
	"fmt"

	"github.com/gorilla/mux"
)

func LogRoutes(router *mux.Router) {
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, _ := route.GetPathTemplate()
		methods, _ := route.GetMethods()
		fmt.Printf("Route: %s Methods: %v\n", path, methods)
		return nil
	})
}
