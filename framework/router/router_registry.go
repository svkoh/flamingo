package router

type (
	// RouterRegistry holds a list of all routes and handlers to be registered
	// in modules.
	RouterRegistry struct {
		routes  map[string]string
		handler map[string]Controller
	}
)

// NewRouterRegistry builds a new RouterRegistry
func NewRouterRegistry() *RouterRegistry {
	return &RouterRegistry{
		routes:  make(map[string]string),
		handler: make(map[string]Controller),
	}
}

// Handle registers the controller for a named route
func (router *RouterRegistry) Handle(name string, controller Controller) {
	router.handler[name] = controller
}

// Router registers the path for a named route
func (router *RouterRegistry) Route(path, name string) {
	router.routes[name] = path
}
