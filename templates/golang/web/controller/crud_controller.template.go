package controller

// crud controller
const (
	CRUDController = `
	package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"gitlab.com/aifaniyi/go-libs/logger"
)

// %sController :
type %sController struct {
	config *settings.Settings
	deps   *settings.Dependencies
}

// New%sController : creates a controller instance
// add parameters to inject controller dependencies e.g
// casandra client, kafka client, etc
func New%sController(config *settings.Settings,
	deps *settings.Dependencies) *%sController {
	return &%sController{
		config,
		deps,
	}
}
`
	Create = `
// Create%sHandler :
func (ctrl *%sController) Create%sHandler(w http.ResponseWriter, r *http.Request) {
reqID := filter.GetRequestID(r.Context())

// validate request : 4XXXX error
data := &create%sRequest{}
if err := render.Bind(r, data); err != nil {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(NewAPIResponse(r.Context(),
		ErrInvalidRequest, err.Error()))
	return
}

// create in db
%sRepo := ctrl.deps.Db.Get%sRepo()
%s, err := %sRepo.Create(data.%s)
if err != nil {
	logger.Error.Printf("[%%s] internal error occured while processing request: %%v",
		reqID, err)
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(NewAPIResponse(r.Context(),
		ErrInternalServerError, err.Error()))
	return
}

json.NewEncoder(w).Encode(NewAPIResponse(r.Context(), StatusOK, create%sResponse{
	%s: %s,
}))
return
}`

	CreateDto = `
// create%sRequest :
type create%sRequest struct {
	%s *models.%s ` + "`json" + `:"%s,omitempty"` + "`" + `
}

func (t *create%sRequest) Bind(r *http.Request) error {
	if t.%s == nil {
		return errors.New("%s object is not provided")
	}

	return nil
}

// create%sResponse :
type create%sResponse struct {
	%s *models.%s ` + "`json" + `:"%s,omitempty"` + "`" + `
}`

	Read = `
	// Read%sHandler :
func (ctrl *%sController) Read%sHandler(w http.ResponseWriter, r *http.Request) {
	reqID := filter.GetRequestID(r.Context())

	// validate request : 4XXXX error
	data := &read%sRequest{}
	if err := render.Bind(r, data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(NewAPIResponse(r.Context(),
			ErrInvalidRequest, err.Error()))
		return
	}

	// create in db
	%sRepo := ctrl.deps.Db.Get%sRepo()
	%ss, err := %sRepo.Read(data.Offset, data.Size)
	if err != nil {
		logger.Error.Printf("[%%s] internal error occured while processing request: %%v",
			reqID, err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(NewAPIResponse(r.Context(),
			ErrInternalServerError, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(NewAPIResponse(r.Context(), StatusOK, read%sResponse{
		%ss: %ss,
	}))
	return
}`

	ReadDto = `
	// read%sRequest :
type read%sRequest struct {
	Offset int ` + "`json" + `:"offset,omitempty"` + "`" + `
	Size   int ` + "`json" + `:"size,omitempty"` + "`" + `
}

func (t *read%sRequest) Bind(r *http.Request) error {
	if t.Size > 100 {
		return fmt.Errorf("requested size %%d is too large", t.Size)
	}

	return nil
}

type read%sResponse struct {
	%ss []models.%s  ` + "`json" + `:"%ss,omitempty"` + "`" + `
}`

	Update = `
	// Update%sHandler :
func (ctrl *%sController) Update%sHandler(w http.ResponseWriter, r *http.Request) {
	reqID := filter.GetRequestID(r.Context())

	// validate request : 4XXXX error
	data := &update%sRequest{}
	if err := render.Bind(r, data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(NewAPIResponse(r.Context(),
			ErrInvalidRequest, err.Error()))
		return
	}

	// update in db
	%sRepo := ctrl.deps.Db.Get%sRepo()
	err := %sRepo.Update(data.%s)
	if err != nil {
		logger.Error.Printf("[%%s] internal error occured while processing request: %%v",
			reqID, err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(NewAPIResponse(r.Context(),
			ErrInternalServerError, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(NewAPIResponse(r.Context(), StatusOK, update%sResponse{
		%s: data.%s,
	}))
	return
}`

	UpdateDto = `
	// update%sRequest :
type update%sRequest struct {
	%s *models.%s ` + "`json" + `:"%s,omitempty"` + "`" + `
}

func (t *update%sRequest) Bind(r *http.Request) error {
	if t.%s == nil {
		return errors.New("%s object is not provided")
	}

	return nil
}

type update%sResponse struct {
	%s *models.%s ` + "`json" + `:"%s,omitempty"` + "`" + `
}`

	Delete = `
	// Delete%sHandler :
func (ctrl *%sController) Delete%sHandler(w http.ResponseWriter, r *http.Request) {
	reqID := filter.GetRequestID(r.Context())

	// validate request : 4XXXX error
	data := &delete%sRequest{}
	if err := render.Bind(r, data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(NewAPIResponse(r.Context(),
			ErrInvalidRequest, err.Error()))
		return
	}

	// delete from db
	%sRepo := ctrl.deps.Db.Get%sRepo()
	err := %sRepo.Delete(data.%s)
	if err != nil {
		logger.Error.Printf("[%%s] internal error occured while processing request: %%v",
			reqID, err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(NewAPIResponse(r.Context(),
			ErrInternalServerError, err.Error()))
		return
	}

	json.NewEncoder(w).Encode(NewAPIResponse(r.Context(), StatusOK, delete%sResponse{
		%s: data.%s,
	}))
	return
}`

	DeleteDto = `
// delete%sRequest :
type delete%sRequest struct {
	%s *models.%s ` + "`json" + `:"%s,omitempty"` + "`" + `
}

func (t *delete%sRequest) Bind(r *http.Request) error {
	if t.%s == nil {
		return errors.New("%s object is not provided")
	}

	return nil
}

type delete%sResponse struct {
	%s *models.%s ` + "`json" + `:"%s,omitempty"` + "`" + `
}`
)
