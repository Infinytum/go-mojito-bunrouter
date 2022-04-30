package bunrouter

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"sync"

	"github.com/infinytum/go-mojito"
	"github.com/infinytum/go-mojito/log"
	"github.com/uptrace/bunrouter"
	brouter "github.com/uptrace/bunrouter"
)

func init() {
	_ = mojito.Register(func() mojito.Router {
		return NewBunRouter()
	}, true)

	_ = mojito.RegisterNamed("bun", func() mojito.Router {
		return NewBunRouter()
	}, true)
}

type bunRouterRouter struct {
	middlewares []interface{}
	router      *brouter.CompatRouter
	routeMap    map[string]*mojito.Handler
	sync.Mutex
	http.Server
}

//// Convenience functions for registering routes

func (r *bunRouterRouter) DELETE(path string, handler interface{}) error {
	return r.WithRoute(mojito.MethodDelete, path, handler)
}

func (r *bunRouterRouter) GET(path string, handler interface{}) error {
	return r.WithRoute(mojito.MethodGet, path, handler)
}

func (r *bunRouterRouter) HEAD(path string, handler interface{}) error {
	return r.WithRoute(mojito.MethodHead, path, handler)
}

func (r *bunRouterRouter) POST(path string, handler interface{}) error {
	return r.WithRoute(mojito.MethodPost, path, handler)
}

func (r *bunRouterRouter) PUT(path string, handler interface{}) error {
	return r.WithRoute(mojito.MethodPut, path, handler)
}

//// Generic functions for adding routes and middleware

// Group is the deprecated version of WithGroup, they do the same thing.
//
// Deprecated: For consistency in naming and due to the Routeable interface
// change, the Group method has been deprecated in favour of WithGroup and
// its new interface.
func (r *bunRouterRouter) Group(path string, callback func(group mojito.Routeable)) error {
	return r.WithGroup(path, func(group mojito.RouteGroup) {
		callback(group)
	})
}

// WithGroup will create a new route group for the given prefix
func (r *bunRouterRouter) WithGroup(path string, callback func(group mojito.RouteGroup)) error {
	rg := mojito.NewRouteGroup()
	callback(rg)
	return rg.ApplyToRouter(r, path)
}

// WithMiddleware will add a middleware to the router
func (r *bunRouterRouter) WithMiddleware(handler interface{}) error {
	r.Lock()
	defer r.Unlock()
	for _, h := range r.routeMap {
		if err := h.AddMiddleware(handler); err != nil {
			log.Error(err)
			return err
		}
	}
	r.middlewares = append(r.middlewares, handler)
	return nil
}

// WithRoute will add a new route with the given RouteMethod to the router
func (r *bunRouterRouter) WithRoute(method RouteMethod, path string, handler interface{}) error {
	r.Lock()
	defer r.Unlock()

	// If the handler is already a mojito handler, skip the handler creation
	h, err := mojito.GetHandler(handler)
	if err != nil {
		// Check if the handler is of kind func, else this is not a valid handler.
		if reflect.TypeOf(handler).Kind() != reflect.Func {
			return errors.New("The given route handler is neither a func nor a mojito.Handler and is therefore not valid")
		}

		// The handler is of kind func, attempt to create a new mojito.Handler for it
		h, err = mojito.NewHandler(handler)
		if err != nil {
			log.Field("method", method).Field("path", path).Errorf("Failed to create a new mojito.Handler for given route handler: %s", err)
			return err
		}
	}

	// Safety check, this should never happen
	if h == nil {
		return errors.New("mojito.Handler was unexpectedly nil")
	}

	// Chain router-wide middleware to the (new) handler
	for _, middleware := range r.middlewares {
		if err := h.AddMiddleware(middleware); err != nil {
			log.Field("method", method).Field("path", path).Errorf("Failed to chain middleware to route: %s", err)
			return err
		}
	}

	switch method {
	case mojito.MethodDelete:
		r.router.Group.Compat().DELETE(path, r.withMojitoHandler(h))
	case mojito.MethodGet:
		r.router.Group.Compat().GET(path, r.withMojitoHandler(h))
	case mojito.MethodHead:
		r.router.Group.Compat().HEAD(path, r.withMojitoHandler(h))
	case mojito.MethodPost:
		r.router.Group.Compat().POST(path, r.withMojitoHandler(h))
	case mojito.MethodPut:
		r.router.Group.Compat().PUT(path, r.withMojitoHandler(h))
	default:
		log.Field("method", method).Field("path", path).Error("The bunrouter router implementation unfortunately does not support this HTTP method")
		return errors.New("The given HTTP method is not available on this router")
	}
	r.routeMap[path] = h
	return nil
}

// ListenAndServe will start an HTTP webserver on the given address
func (r *bunRouterRouter) ListenAndServe(address string) error {
	r.Server = http.Server{
		Addr:    address,
		Handler: r.router,
	}
	return r.Server.ListenAndServe()
}

// Shutdown will stop the HTTP webserver
func (r *bunRouterRouter) Shutdown() error {
	return r.Server.Shutdown(context.TODO())
}

// NewBunRouter will create new instance of the mojito bun router implementation
func NewBunRouterRouter() mojito.Router {
	router := brouter.New().Compat()
	return &bunRouterRouter{
		router:   router,
		routeMap: make(map[string]*mojito.Handler),
		Mutex:    sync.Mutex{},
	}
}

//// Internal functions

func (r *bunRouterRouter) withMojitoHandler(handler *mojito.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := mojito.NewRequest(r)
		res := mojito.NewResponse(w)
		req.Params = bunrouter.ParamsFromContext(r.Context()).Map()
		if err := handler.Serve(req, res); err != nil {
			log.Field("router", "bun").Error(err)
		}
	}
}
