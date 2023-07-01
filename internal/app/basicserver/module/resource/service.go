package resource

import (
	"net/http"

	"github.com/tanveerprottoy/basic-go-server/internal/app/basicserver/module/resource/dto"
	"github.com/tanveerprottoy/basic-go-server/internal/app/basicserver/module/resource/entity"
	"github.com/tanveerprottoy/basic-go-server/pkg/config"
	"github.com/tanveerprottoy/basic-go-server/pkg/data/sqlxpkg"
	"github.com/tanveerprottoy/basic-go-server/pkg/errorpkg"
	"github.com/tanveerprottoy/basic-go-server/pkg/response"
	"github.com/tanveerprottoy/basic-go-server/pkg/timepkg"
)

type Service struct {
	repository sqlxpkg.Repository[entity.Resource]
}

func NewService(r sqlxpkg.Repository[entity.Resource]) *Service {
	s := new(Service)
	s.repository = r
	return s
}

func (s *Service) readOneInternal(id string, w http.ResponseWriter) (entity.Resource, error) {
	return s.repository.ReadOne(id)
}

func (s *Service) GetBasicData(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]any)
	m["message"] = config.GetEnvValue("MESSAGE")
	response.Respond(http.StatusOK, response.BuildData(m), w)
}

func (s *Service) Create(d *dto.CreateUpdateResourceDto, w http.ResponseWriter, r *http.Request) {
	// convert dto to entity
	b := entity.Resource{}
	b.Name = d.Name
	n := timepkg.NowUnixMilli()
	b.CreatedAt = n
	b.UpdatedAt = n
	err := s.repository.Create(&b)
	if err != nil {
		errorpkg.HandleDBError(err, w)
		return
	}
	response.Respond(http.StatusCreated, response.BuildData(d), w)
}

func (s *Service) ReadMany(limit, page int, w http.ResponseWriter, r *http.Request) {
	offset := limit * (page - 1)
	d, err := s.repository.ReadMany(limit, offset)
	if err != nil {
		errorpkg.HandleDBError(err, w)
		return
	}
	m := make(map[string]any)
	m["items"] = d
	m["limit"] = limit
	m["page"] = page
	response.Respond(http.StatusOK, response.BuildData(m), w)
}

func (s *Service) ReadOne(id string, w http.ResponseWriter, r *http.Request) {
	b, err := s.readOneInternal(id, w)
	if err != nil {
		errorpkg.HandleDBError(err, w)
		return
	}
	response.Respond(http.StatusOK, response.BuildData(b), w)
}

func (s *Service) Update(id string, d *dto.CreateUpdateResourceDto, w http.ResponseWriter, r *http.Request) {
	b, err := s.readOneInternal(id, w)
	if err != nil {
		errorpkg.HandleDBError(err, w)
		return
	}
	b.Name = d.Name
	b.UpdatedAt = timepkg.NowUnixMilli()
	rows, err := s.repository.Update(id, &b)
	if err != nil {
		errorpkg.HandleDBError(err, w)
		return
	}
	if rows > 0 {
		response.Respond(http.StatusOK, response.BuildData(b), w)
		return
	}
	response.RespondError(http.StatusBadRequest, errorpkg.NewError("operation was not successful"), w)
}

func (s *Service) Delete(id string, w http.ResponseWriter, r *http.Request) {
	b, err := s.readOneInternal(id, w)
	if err != nil {
		errorpkg.HandleDBError(err, w)
		return
	}
	rows, err := s.repository.Delete(id)
	if err != nil {
		errorpkg.HandleDBError(err, w)
		return
	}
	if rows > 0 {
		response.Respond(http.StatusOK, response.BuildData(b), w)
		return
	}
	response.RespondError(http.StatusBadRequest, errorpkg.NewError("operation was not successful"), w)
}
