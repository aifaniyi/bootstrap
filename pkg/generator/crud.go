package generator

import (
	"fmt"

	"github.com/aifaniyi/bootstrap/templates/golang/web/controller"
	"github.com/iancoleman/strcase"
)

func getCRUDController(entity entity) string {
	return fmt.Sprintf(`
	%s
	
	%s
	
	%s
	
	%s
	
	%s
	
	%s
	
	%s
	
	%s
	
	%s`,
		getController(entity),
		getCreate(entity),
		getRead(entity),
		getUpdate(entity),
		getDelete(entity),
		getCreateDto(entity),
		getReadDto(entity),
		getUpdateDto(entity),
		getDeleteDto(entity))
}

func getController(entity entity) string {
	upperCamel := strcase.ToCamel(entity.Name)

	return fmt.Sprintf(controller.CRUDController,
		upperCamel,
		upperCamel,
		upperCamel,
		upperCamel,
		upperCamel,
		upperCamel,
	)
}

func getCreate(entity entity) string {
	// lower := strings.ToLower(entity.Name)
	lowerCamel := strcase.ToLowerCamel(entity.Name)
	upperCamel := strcase.ToCamel(entity.Name)

	return fmt.Sprintf(controller.Create,
		upperCamel,
		upperCamel, upperCamel,
		upperCamel,
		lowerCamel, upperCamel,
		lowerCamel, lowerCamel, upperCamel,
		upperCamel,
		upperCamel, lowerCamel,
	)
}

func getCreateDto(entity entity) string {
	lowerCamel := strcase.ToLowerCamel(entity.Name)
	upperCamel := strcase.ToCamel(entity.Name)

	return fmt.Sprintf(controller.CreateDto,
		upperCamel,
		upperCamel,
		upperCamel, upperCamel, lowerCamel,
		upperCamel,
		upperCamel,
		lowerCamel,
		upperCamel,
		upperCamel,
		upperCamel, upperCamel, lowerCamel,
	)
}

func getRead(entity entity) string {
	lowerCamel := strcase.ToLowerCamel(entity.Name)
	upperCamel := strcase.ToCamel(entity.Name)

	return fmt.Sprintf(controller.Read,
		upperCamel,
		upperCamel, upperCamel,
		upperCamel,
		lowerCamel, upperCamel,
		lowerCamel, lowerCamel,
		upperCamel,
		upperCamel, lowerCamel,
	)
}

func getReadDto(entity entity) string {
	lowerCamel := strcase.ToLowerCamel(entity.Name)
	upperCamel := strcase.ToCamel(entity.Name)

	return fmt.Sprintf(controller.ReadDto,
		upperCamel,
		upperCamel,
		upperCamel,
		upperCamel,
		upperCamel, upperCamel, lowerCamel,
	)
}

func getUpdate(entity entity) string {
	lowerCamel := strcase.ToLowerCamel(entity.Name)
	upperCamel := strcase.ToCamel(entity.Name)

	return fmt.Sprintf(controller.Update,
		upperCamel,
		upperCamel, upperCamel,
		upperCamel,
		lowerCamel, upperCamel,
		lowerCamel, upperCamel,
		upperCamel,
		upperCamel, upperCamel,
	)
}

func getUpdateDto(entity entity) string {
	lowerCamel := strcase.ToLowerCamel(entity.Name)
	upperCamel := strcase.ToCamel(entity.Name)

	return fmt.Sprintf(controller.UpdateDto,
		upperCamel,
		upperCamel,
		upperCamel, upperCamel, lowerCamel,
		upperCamel,
		upperCamel,
		lowerCamel,
		upperCamel,
		upperCamel, upperCamel, lowerCamel,
	)
}

func getDelete(entity entity) string {
	lowerCamel := strcase.ToLowerCamel(entity.Name)
	upperCamel := strcase.ToCamel(entity.Name)

	return fmt.Sprintf(controller.Delete,
		upperCamel,
		upperCamel, upperCamel,
		upperCamel,
		lowerCamel, upperCamel,
		lowerCamel, upperCamel,
		upperCamel,
		upperCamel, upperCamel,
	)
}

func getDeleteDto(entity entity) string {
	lowerCamel := strcase.ToLowerCamel(entity.Name)
	upperCamel := strcase.ToCamel(entity.Name)

	return fmt.Sprintf(controller.DeleteDto,
		upperCamel,
		upperCamel,
		upperCamel, upperCamel, lowerCamel,
		upperCamel,
		upperCamel,
		lowerCamel,
		upperCamel,
		upperCamel, upperCamel, lowerCamel,
	)
}
