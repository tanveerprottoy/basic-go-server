package dto

// CreateUpdateResourceDto dto for create update
type CreateUpdateResourceDto struct {
	Name string `json:"name" validate:"required"`
}
