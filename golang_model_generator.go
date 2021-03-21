package main

import (
	"fmt"
	"strings"

	gotemplate "github.com/aifaniyi/bootstrap/templates/golang"
	"github.com/iancoleman/strcase"
)

var (
	types        = make(map[string]bool)
	crossUpdates = make(map[string]string) // fields to be added to an entity because of associations to other entities
)

func generateGolangModels(schema *schema) (map[string]string, error) {
	entities := make(map[string]string)

	for _, entity := range schema.Entities {
		types[entity.Name] = true
	}

	for _, entity := range schema.Entities {
		entityString, err := generateGolangEntity(entity, schema)
		if err != nil {
			return nil, fmt.Errorf("error processing input file: %v", err)
		}
		entities[entity.Name] = entityString
	}

	return entities, nil
}

func generateGolangEntity(entity entity, schema *schema) (string, error) {
	// add ID, CreatedAt and UpdatedAt columns by default
	fields, err := generateFields(entity, schema)
	if err != nil {
		return "", err
	}

	content := fmt.Sprintf(`
	package models

	import (
		"time"
	)

	// %s : %s
	type %s struct {
		ID int `+"`json:\"-\" gorm:\"primary_key\"`"+`
		%s
		CreatedAt time.Time `+"`json:\"createdAt\"` // creation time"+`
		UpdatedAt time.Time `+"`json:\"updatedAt\" gorm:\"default:CURRENT_TIMESTAMP\"`  // update time"+`
	}
	`, entity.Name, entity.Description,
		entity.Name, fields)

	return content, nil
}

func generateFields(entity entity, schema *schema) (string, error) {
	upperCamel := strcase.ToCamel(entity.Name)
	content := ""
	for index, property := range entity.Properties {
		propertyString, err := generateProperty(property)
		if err != nil {
			return "", err
		}

		if index == 0 {
			content += propertyString
		} else {
			content += "\n" + propertyString
		}
	}

	// for entity relations
	for _, relation := range entity.Relations {
		relationString, err := generateRelation(relation, entity)
		if err != nil {
			return "", err
		}

		content += "\n" + relationString
	}

	// add cross update fields
	if val, ok := crossUpdates[upperCamel]; ok {
		content += "\n" + val
	}

	return content, nil
}

func generateProperty(property property) (string, error) {
	switch property.Type.Name {
	case "string", "integer":
	default:
		if _, ok := types[property.Type.Name]; ok {
			goto process
		}
		return "", fmt.Errorf("unrecognized type %s", property.Type.Name)
	}

process:
	name := strcase.ToCamel(property.Name)
	if strings.ToLower(name) == "uuid" {
		name = "UUID"
	}

	json := "-"
	if property.Dto {
		json = fmt.Sprintf("%s", property.Name)
	}

	content := fmt.Sprintf("%s %s `json:\"%s\" %s` // %s",
		name, getGolangType(property.Type.Name), json,
		getGormString(property),
		property.Description)

	return content, nil
}

func generateRelation(relation relation, entity entity) (string, error) {
	upperCamel := strcase.ToCamel(relation.Entity)

	switch strcase.ToLowerCamel(relation.Type) {
	case "belongsTo":
		return fmt.Sprintf("%sID int\n%s %s", upperCamel, upperCamel, upperCamel), nil

	case "hasOne":
		crossUpdates[upperCamel] = fmt.Sprintf("%sID int", upperCamel)
		return fmt.Sprintf("%s %s", upperCamel, upperCamel), nil

	case "hasMany":
		crossUpdates[upperCamel] = fmt.Sprintf("%sID int", upperCamel)
		return fmt.Sprintf("%s []%s", upperCamel, upperCamel), nil

	case "manyToMany":
		return fmt.Sprintf("%s []%s", upperCamel, upperCamel), nil
	}

	return "", fmt.Errorf("unknown relation type %s. Accepts only [belongsTo, hasOne, hasMany, manyToMany]", relation.Type)
}

func getGormString(property property) string {
	str := "gorm:\""
	parts := make([]string, 0)

	if property.Indexable {
		parts = append(parts, "index")
	}

	if !property.Nullable {
		parts = append(parts, "not null")
	}

	if property.Unique {
		parts = append(parts, "unique")
	}

	if property.Width != nil {
		parts = append(parts, fmt.Sprintf("size:%d", *property.Width))
	}

	return str + strings.Join(parts, ";") + "\""
}

func getGolangType(str string) string {
	switch str {
	case "integer":
		return "int"

	case "string":
		return "string"

	default:
		return "string"
	}
}

// content of main.go
func getGolangMainContent() string {
	return gotemplate.Main
}

// content of server.go
func getGolangServerContent() string {
	return gotemplate.Server
}

// content of settings.go
func getGolangSettingsContent() string {
	return gotemplate.Settings
}
