package main

import (
	"fmt"
	"strings"

	"github.com/aifaniyi/bootstrap/templates/golang/web"
	"github.com/iancoleman/strcase"
)

func getRoutes(schema *schema) string {
	ctrls, routes := getEntityRoutes(schema)

	return fmt.Sprintf(web.Route,
		getOauthRoutes(schema),
		ctrls,
		routes,
	)
}

func getOauthRoutes(schema *schema) string {
	lines := make([]string, 0)
	for _, authEntry := range schema.Auth {
		lowerCamel := strcase.ToLowerCamel(authEntry)
		upperCamel := strcase.ToCamel(authEntry)

		lines = append(lines, fmt.Sprintf(`http.HandleFunc("/api/v1/%s/login", oauth.Handle%sLogin)
			http.HandleFunc("/api/v1/%s/callback", oauth.Handle%sCallback)`,
			lowerCamel, upperCamel,
			lowerCamel, upperCamel))
	}

	return strings.Join(lines, "\n")
}

func getEntityRoutes(schema *schema) (string, string) {
	ctrls := make([]string, 0)
	routes := make([]string, 0)

	for _, entity := range schema.Entities {
		lower := strings.ToLower(entity.Name)
		lowerCamel := strcase.ToLowerCamel(entity.Name)
		upperCamel := strcase.ToCamel(entity.Name)

		ctrls = append(ctrls, fmt.Sprintf(`%sController := controllers.New%sController(config, deps)`,
			lowerCamel,
			upperCamel,
		))

		routes = append(routes, fmt.Sprintf(`
		// %s routes
		router.Route("/api/v1/%s/", func(router chi.Router) {
			router.Post("/create", %sController.Create%sHandler)
			router.Post("/read", %sController.Read%sHandler)
			router.Post("/update", %sController.Update%sHandler)
		})
		`,
			upperCamel,
			lower,
			lowerCamel, upperCamel,
			lowerCamel, upperCamel,
			lowerCamel, upperCamel,
		))
	}

	return strings.Join(ctrls, "\n"), strings.Join(routes, "\n\n")
}
