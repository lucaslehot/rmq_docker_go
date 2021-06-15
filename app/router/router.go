package router

import (
	"net/http"
	"app/controllers"
	"app/middleware"

	"github.com/gorilla/mux"
)

//This is a router, powered by mux!
//The route type contains a name for the route, a method (PUT, GET, POST...) a pattern (the url, basically) and a handler function that
//makes the connection between a controller and the route.

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// NewRouter is the register for every public route
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	//Check every route created below register it
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

//This is gonna be parsed, so this must contain every public route.
var routes = Routes{

	// User management routes

	Route{
		Name:        "Create user",
		Method:      "POST",
		Pattern:     "/user/create",
		HandlerFunc: middleware.VerifyJwt(controllers.CreateUser),
	},

	Route{
		Name:        "Read user",
		Method:      "GET",
		Pattern:     `/user/{ID}`,
		HandlerFunc: controllers.ReadUser,
	},

	Route{
		Name:        "Update user",
		Method:      "POST",
		Pattern:     "/user/update",
		HandlerFunc: middleware.VerifyJwt(controllers.UpdateUser),
	},

	Route{
		Name:        "Delete user",
		Method:      "POST",
		Pattern:     `/delete/delete`,
		HandlerFunc: middleware.VerifyJwt(controllers.DeleteUser),
	},

	// Simple 

	Route{
		Name:        "Welcome user",
		Method:      "GET",
		Pattern:     "/welcome",
		HandlerFunc: middleware.VerifyJwt(controllers.Welcome),
	},
}
