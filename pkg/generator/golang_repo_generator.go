package generator

import (
	"fmt"
	"strings"

	gotemplate "github.com/aifaniyi/bootstrap/templates/golang"
	"github.com/iancoleman/strcase"
)

// content of db.go
func getGolangDbContent(schema *schema) string {
	lines := make([]string, 0)

	for _, entity := range schema.Entities {
		lines = append(lines, fmt.Sprintf("Get%sRepo() %srepo.Service",
			strcase.ToCamel(entity.Name), strings.ToLower(entity.Name)))
	}

	return fmt.Sprintf(gotemplate.DbService, strings.Join(lines, "\n"))
}

// content of db_impl.go
func getGolangDbImplContent(schema *schema) string {
	lines := make([]string, 0)
	migrateModelLines := make([]string, 0)
	functionLines := make([]string, 0)

	for _, entity := range schema.Entities {
		lower := strings.ToLower(entity.Name)
		upperCamel := strcase.ToCamel(entity.Name)
		lowerCamel := strcase.ToLowerCamel(entity.Name)

		lines = append(lines, fmt.Sprintf("%sRepo %srepo.Service",
			lowerCamel, lower))

		migrateModelLines = append(migrateModelLines, fmt.Sprintf("&models.%s{},",
			upperCamel))

		functionLines = append(functionLines, fmt.Sprintf(gotemplate.DbServiceGetRepoFunction,
			upperCamel, lower, upperCamel, lower,
			lowerCamel, lowerCamel, lower, upperCamel, lowerCamel))
	}

	return fmt.Sprintf(gotemplate.DbServiceImpl,
		strings.Join(lines, "\n"),
		strings.Join(migrateModelLines, "\n"),
		strings.Join(functionLines, "\n"))
}

// content of db_impl.go
func getGolangEntityRepoContent(entity entity) string {
	lower := strings.ToLower(entity.Name)
	lowerCamel := strcase.ToLowerCamel(entity.Name)
	upperCamel := strcase.ToCamel(entity.Name)

	// TODO : implement read by specific fields and paginate on read all
	return fmt.Sprintf(gotemplate.DbRepoService,
		lower,
		lowerCamel, upperCamel, upperCamel,
		upperCamel,
		lowerCamel, upperCamel,
		lowerCamel, upperCamel)
}

// content of db_impl.go
func getGolangEntityRepoImplContent(entity entity) string {
	lower := strings.ToLower(entity.Name)
	lowerCamel := strcase.ToLowerCamel(entity.Name)
	upperCamel := strcase.ToCamel(entity.Name)

	// TODO : implement read by specific fields and paginate on read all
	return fmt.Sprintf(gotemplate.DbRepoServiceImpl,
		lower,
		upperCamel, lower,
		upperCamel,
		lower, lowerCamel, upperCamel,
		upperCamel,
		lowerCamel,
		lower,
		lowerCamel,

		lower,
		upperCamel,
		lowerCamel, upperCamel,
		lowerCamel,
		lower,
		lower,
		lowerCamel,

		lower,
		lowerCamel, upperCamel,
		lowerCamel, upperCamel,
		lowerCamel, lowerCamel,
		lowerCamel, lowerCamel,

		lower,
		lowerCamel, upperCamel,
		lowerCamel,
	)
}
