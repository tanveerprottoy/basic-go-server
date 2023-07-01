package resource

import (
	"fmt"
	"net/http"

	"github.com/tanveerprottoy/basic-go-server/internal/app/basicserver/module/resource/dto"
	"github.com/tanveerprottoy/basic-go-server/internal/pkg/constant"
	"github.com/tanveerprottoy/basic-go-server/pkg/adapter"
	"github.com/tanveerprottoy/basic-go-server/pkg/httppkg"
	"github.com/tanveerprottoy/basic-go-server/pkg/jsonpkg"
	"github.com/tanveerprottoy/basic-go-server/pkg/response"

	"github.com/go-playground/validator/v10"
)

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

func (h *Handler) parseValidateRequestBody(r *http.Request) (dto.CreateUpdateResourceDto, error) {
	var d dto.CreateUpdateResourceDto
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
	h.service.GetBasicData(w, r)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	d, err := h.parseValidateRequestBody(r)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	h.service.Create(&d, w, r)
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
	h.service.ReadMany(limit, page, w, r)
}

func (h *Handler) ReadOne(w http.ResponseWriter, r *http.Request) {
	id := httppkg.GetURLParam(r, constant.KeyId)
	h.service.ReadOne(id, w, r)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := httppkg.GetURLParam(r, constant.KeyId)
	d, err := h.parseValidateRequestBody(r)
	if err != nil {
		response.RespondError(http.StatusBadRequest, err, w)
		return
	}
	h.service.Update(id, &d, w, r)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := httppkg.GetURLParam(r, constant.KeyId)
	h.service.Delete(id, w, r)
}
