package resource

import (
	"fmt"
	"net/http"

	"github.com/tanveerprottoy/basic-go-server/internal/pkg/constant"
	"github.com/tanveerprottoy/basic-go-server/pkg/adapter"
	"github.com/tanveerprottoy/basic-go-server/pkg/httppkg"
	"github.com/tanveerprottoy/basic-go-server/pkg/jsonpkg"
	"github.com/tanveerprottoy/basic-go-server/pkg/response"

	"github.com/go-playground/validator/v10"
)

// Hanlder is responsible for extracting data
// from request body and building and seding response
type Handler struct {
	service  *Service
	validate *validator.Validate
}

func NewHandler(s *Service, v *validator.Validate) *Handler {
	h := new(Handler)
	h.service = s
	h.validate = v
	return h
}

func (h *Handler) parseValidateRequestBody(r *http.Request) (CreateUpdateResourceDto, error) {
	var d CreateUpdateResourceDto
	err := jsonpkg.Decode(r.Body, &d)
	if err != nil {
		return d, err
	}
	// validate request body
	err = h.validate.Struct(d)
	if err != nil {
		// Struct is invalid
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag())
		}
	}
	return d, err
}

func (h *Handler) GetBasicData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	m := h.service.GetBasicData(ctx)
	response.Respond(http.StatusOK, response.BuildData(m), w)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	d, err := h.parseValidateRequestBody(r)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	ctx := r.Context()
	e, httpErr := h.service.Create(&d, ctx)
	if httpErr != nil {
		response.RespondError(httpErr.Code, httpErr.Err, w)
		return
	}
	response.Respond(http.StatusCreated, e, w)
}

func (h *Handler) ReadMany(w http.ResponseWriter, r *http.Request) {
	limit := 10
	page := 1
	var err error
	limitStr := httppkg.GetQueryParam(r, constant.KeyLimit)
	if limitStr != "" {
		limit, err = adapter.StringToInt(limitStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err, w)
			return
		}
	}
	pageStr := httppkg.GetQueryParam(r, constant.KeyPage)
	if pageStr != "" {
		page, err = adapter.StringToInt(pageStr)
		if err != nil {
			response.RespondError(http.StatusBadRequest, err, w)
			return
		}
	}
	e, httpErr := h.service.ReadMany(limit, page, nil)
	if httpErr != nil {
		response.RespondError(httpErr.Code, httpErr.Err, w)
	}
	response.Respond(http.StatusOK, e, w)
}

func (h *Handler) ReadOne(w http.ResponseWriter, r *http.Request) {
	id := httppkg.GetURLParam(r, constant.KeyId)
	e, httpErr := h.service.ReadOne(id, nil)
	if httpErr != nil {
		response.RespondError(httpErr.Code, httpErr.Err, w)
	}
	response.Respond(http.StatusOK, e, w)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := httppkg.GetURLParam(r, constant.KeyId)
	d, err := h.parseValidateRequestBody(r)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	e, httpErr := h.service.Update(id, &d, nil)
	if httpErr != nil {
		response.RespondError(httpErr.Code, httpErr.Err, w)
	}
	response.Respond(http.StatusOK, e, w)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := httppkg.GetURLParam(r, constant.KeyId)
	e, httpErr := h.service.Delete(id, nil)
	if httpErr != nil {
		response.RespondError(httpErr.Code, httpErr.Err, w)
	}
	response.Respond(http.StatusOK, e, w)
}
