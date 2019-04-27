package route

import (
	"github.com/YTPSourceCode/SoruApi/api/models"

	"github.com/gorilla/mux"
)

var routes = models.Routes{ //Route List
	models.Route{
		Name:        "UserGet",
		Path:        "/user/get",
		Method:      "Get",
		HandlerFunc: GetUser,
	},
	models.Route{
		Name:        "UserCreate",
		Path:        "/user/create",
		Method:      "Post",
		HandlerFunc: CreateUser,
	},
	models.Route{
		Name:        "UserGetOne",
		Path:        "/user/getone",
		Method:      "Post",
		HandlerFunc: GetOneUser,
	},
	models.Route{
		Name:        "UserDelete",
		Path:        "/user/getone",
		Method:      "Delete",
		HandlerFunc: DeleteUser,
	},
	models.Route{
		Name:        "UserUpdate",
		Path:        "/user/getone",
		Method:      "Put",
		HandlerFunc: UpdateUser,
	},
}

// SetRouter Function for create Api End Points.
func SetRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true) // Create router
	for _, route := range routes {              //Adding route list's property
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router // Return router
}
