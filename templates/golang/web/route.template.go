package web

// route file content
const (
	Route = `
	package web

// SetupRouteControllers : create controllers
// and sets up route handlers
func SetupRouteControllers(router *chi.Mux, config *settings.Settings,
	deps *settings.Dependencies) {

	// filters
	router.Use(filter.AddHeaders)
	router.Use(filter.AddContext)
	router.Use(filter.LoggingMiddleware)
	router.Use(corsMW().Handler)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)

	// create oauth handlers
	%s

	// create controller instances
	healthCheckController := controllers.NewHealthCheckController()
	%s

	// map route to handler
	router.Get("/health", healthCheckController.HealthCheckerHandler)

	%s
}

func corsMW() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})
}
`
)
