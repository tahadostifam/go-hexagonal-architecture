package restful_adapter

type AppOption func(a *App)

func WithPort(port int) AppOption {
	return func(a *App) {
		a.port = port
	}
}
