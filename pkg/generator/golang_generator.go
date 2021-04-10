package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/aifaniyi/bootstrap/templates/golang/web/controller"
	"github.com/aifaniyi/bootstrap/templates/golang/web/filter"
)

const (
	webfiles    = "packages/web"
	oauthDir    = "packages/web/auth/oauth"
	controllers = "packages/web/controllers"
	filters     = "packages/web/filters"
	models      = "packages/models"
	dbRepo      = "packages/services/db"
	settings    = "packages/settings"
)

func GenerateGolang(inputFile, outputDir, project string) error {
	// parse
	schema, err := getSchema(inputFile)
	if err != nil {
		return err
	}

	// generate entities
	entities, err := generateGolangModels(schema)
	if err != nil {
		return err
	}

	// prepare output directories
	dirs := []string{
		"cmd/" + project,
		oauthDir,
		models,
		dbRepo,
		filters,
		controllers}

	for _, directory := range dirs {
		dir := filepath.Join(outputDir, directory)
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating directory %s: %v", dir, err)
		}
	}

	/*******************************
	 write main file and entities **
	********************************/
	filename := filepath.Join(outputDir, "cmd/"+project+"/main.go")
	err = writeFile(getGolangMainContent(), filename)
	if err != nil {
		return fmt.Errorf("error writing main file: %v", err)
	}

	filename = filepath.Join(outputDir, "cmd/"+project+"/server.go")
	err = writeFile(getGolangServerContent(), filename)
	if err != nil {
		return fmt.Errorf("error writing main file: %v", err)
	}

	// write base entity
	filename = filepath.Join(outputDir, models, "base.go")
	err = writeFile(getGolangBaseModelContent(), filename)
	if err != nil {
		return fmt.Errorf("error writing base model file: %v", err)
	}

	// write entities
	for key, entity := range entities {
		filename := filepath.Join(outputDir, models, strings.ToLower(key)+".go")
		err = writeFile(entity, filename)
		if err != nil {
			return fmt.Errorf("error writing entity [%s] to output file [%s]: %v", key, filename, err)
		}
	}

	/********************************
	 write db repos *****************
	*********************************/
	// generate main db and dbimpl file
	filename = filepath.Join(outputDir, dbRepo, "db.go")
	err = writeFile(getGolangDbContent(schema), filename)
	if err != nil {
		return fmt.Errorf("error writing db service file: %v", err)
	}

	filename = filepath.Join(outputDir, dbRepo, "db_impl.go")
	err = writeFile(getGolangDbImplContent(schema), filename)
	if err != nil {
		return fmt.Errorf("error writing db service file: %v", err)
	}

	// generate each entity repo package
	for _, entity := range schema.Entities {
		lower := strings.ToLower(entity.Name)

		dir := filepath.Join(outputDir, dbRepo, lower+"repo")
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating directory %s: %v", dir, err)
		}

		filename := filepath.Join(outputDir, dbRepo, lower+"repo", lower+"_repo.go")
		err = writeFile(getGolangEntityRepoContent(entity), filename)
		if err != nil {
			return fmt.Errorf("error writing entity [%s] to output file [%s]: %v", lower, filename, err)
		}

		filename = filepath.Join(outputDir, dbRepo, lower+"repo", lower+"_repo_impl.go")
		err = writeFile(getGolangEntityRepoImplContent(entity), filename)
		if err != nil {
			return fmt.Errorf("error writing entity impl file [%s] to output file [%s]: %v", lower, filename, err)
		}
	}

	/***************************************
	********** write settings file**********
	***************************************/
	dir := filepath.Join(outputDir, settings)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating directory %s: %v", dir, err)
	}

	filename = filepath.Join(outputDir, settings, "settings.go")
	err = writeFile(getGolangSettingsContent(), filename)
	if err != nil {
		return fmt.Errorf("error writing settings file: %v", err)
	}

	/***************************************
	********** write oauth file*************
	***************************************/
	dir = filepath.Join(outputDir, oauthDir)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating directory %s: %v", dir, err)
	}

	filename = filepath.Join(outputDir, oauthDir, "randomstate.go")
	err = writeFile(generateRandomStateString(), filename)
	if err != nil {
		return fmt.Errorf("error writing oauth randomstate file: %v", err)
	}

	for _, authEntry := range schema.Auth {
		filename = filepath.Join(outputDir, oauthDir, authEntry+"_config.go")
		err = writeFile(generateOauthConfig(authEntry), filename)
		if err != nil {
			return fmt.Errorf("error writing oauth config file: %v", err)
		}

		filename = filepath.Join(outputDir, oauthDir, authEntry+"_handler.go")
		err = writeFile(generateOauthHandler(authEntry), filename)
		if err != nil {
			return fmt.Errorf("error writing oauth handler file: %v", err)
		}
	}

	/*****************************************
	*** write response files and filters *****
	*******************************************/
	filename = filepath.Join(outputDir, controllers, "response.go")
	err = writeFile(controller.Response, filename)
	if err != nil {
		return fmt.Errorf("error writing response file: %v", err)
	}

	// api response writer filter
	filename = filepath.Join(outputDir, filters, "response_writer.go")
	err = writeFile(filter.ResponseWriter, filename)
	if err != nil {
		return fmt.Errorf("error writing response writer filter file: %v", err)
	}

	// context filter
	filename = filepath.Join(outputDir, filters, "context.go")
	err = writeFile(filter.Context, filename)
	if err != nil {
		return fmt.Errorf("error writing context filter file: %v", err)
	}

	// headers filter
	filename = filepath.Join(outputDir, filters, "header.go")
	err = writeFile(filter.Headers, filename)
	if err != nil {
		return fmt.Errorf("error writing headers filter file: %v", err)
	}

	// logger filter
	filename = filepath.Join(outputDir, filters, "logger.go")
	err = writeFile(filter.Logger, filename)
	if err != nil {
		return fmt.Errorf("error writing logger filter file: %v", err)
	}

	/*****************************************
	*********** write routes *****************
	*******************************************/

	/*****************************************
	*********** write controllers ************
	*******************************************/
	// health controller
	filename = filepath.Join(outputDir, controllers, "healthcheck_controller.go")
	err = writeFile(controller.HealthCheckController, filename)
	if err != nil {
		return fmt.Errorf("error writing healthcheck controller file: %v", err)
	}

	// routes
	filename = filepath.Join(outputDir, webfiles, "routes.go")
	err = writeFile(getRoutes(schema), filename)
	if err != nil {
		return fmt.Errorf("error writing routes file: %v", err)
	}

	// generate each entity repo package
	for _, entity := range schema.Entities {
		lower := strings.ToLower(entity.Name)

		filename := filepath.Join(outputDir, controllers, lower+"_controller.go")
		err = writeFile(getCRUDController(entity), filename)
		if err != nil {
			return fmt.Errorf("error writing entity [%s] to output file [%s]: %v", lower, filename, err)
		}
	}
	return nil
}
