package bunrouter

import "github.com/infinytum/go-mojito"

// AsDefault registers this router as the default router
func AsDefault() {
	err := mojito.Register(func() mojito.Router {
		return NewBunRouterRouter()
	}, true)
	if err != nil {
		panic(err)
	}
}

// As registers this router under a given name
func As(name string) {
	err := mojito.RegisterNamed(name, func() mojito.Router {
		return NewBunRouterRouter()
	}, true)
	if err != nil {
		panic(err)
	}
}
