package resource

import (
	"github.com/jmoiron/sqlx"
	"github.com/tanveerprottoy/basic-go-server/pkg/data/sqlxpkg"
)

const TableName = "resources"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	r := new(Repository)
	r.db = db
	return r
}

func (r *Repository) Create(e *Resource) error {
	var lastId string
	err := r.db.QueryRow("INSERT INTO "+TableName+" (title, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id", e.Name, e.CreatedAt, e.UpdatedAt).Scan(&lastId)
	if err != nil {
		return err
	}
	e.Id = lastId
	return nil
}

func (r *Repository) ReadMany(limit, offset int) ([]Resource, error) {
	d := []Resource{}
	err := r.db.Select(&d, "SELECT * FROM "+TableName+" LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (r *Repository) ReadOne(id string) (Resource, error) {
	b := Resource{}
	err := r.db.Get(&b, "SELECT * FROM "+TableName+" WHERE id = $1 LIMIT 1", id)
	if err != nil {
		return b, err
	}
	return b, nil
}

func (r *Repository) Update(id string, e *Resource) (int64, error) {
	q := "UPDATE " + TableName + " SET name = $2, updated_at = $3 WHERE id = $1"
	res, err := r.db.Exec(q, id, e.Name, e.UpdatedAt)
	if err != nil {
		return -1, err
	}
	return sqlxpkg.GetRowsAffected(res), nil
}

func (r *Repository) Delete(id string) (int64, error) {
	q := "DELETE FROM " + TableName + " WHERE id = $1"
	res, err := r.db.Exec(q, id)
	if err != nil {
		return -1, err
	}
	return sqlxpkg.GetRowsAffected(res), nil
}
