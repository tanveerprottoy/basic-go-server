package resource

import (
	"github.com/jmoiron/sqlx"
	"github.com/tanveerprottoy/basic-go-server/internal/app/basicserver/module/resource/entity"
	"github.com/tanveerprottoy/basic-go-server/pkg/data/sqlxpkg"
)

const TableName = "resources"

type Repository[T entity.Resource] struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository[entity.Resource] {
	r := new(Repository[entity.Resource])
	r.db = db
	return r
}

func (r *Repository[T]) Create(e *entity.Resource) error {
	var lastId string
	err := r.db.QueryRow("INSERT INTO "+TableName+" (title, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id", e.Name, e.CreatedAt, e.UpdatedAt).Scan(&lastId)
	if err != nil {
		return err
	}
	e.Id = lastId
	return nil
}

func (r *Repository[T]) ReadMany(limit, offset int) ([]entity.Resource, error) {
	d := []entity.Resource{}
	err := r.db.Select(&d, "SELECT * FROM "+TableName+" LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (r *Repository[T]) ReadOne(id string) (entity.Resource, error) {
	b := entity.Resource{}
	err := r.db.Get(&b, "SELECT * FROM "+TableName+" WHERE id = $1 LIMIT 1", id)
	if err != nil {
		return b, err
	}
	return b, nil
}

func (r *Repository[T]) Update(id string, e *entity.Resource) (int64, error) {
	q := "UPDATE " + TableName + " SET name = $2, updated_at = $3 WHERE id = $1"
	res, err := r.db.Exec(q, id, e.Name, e.UpdatedAt)
	if err != nil {
		return -1, err
	}
	return sqlxpkg.GetRowsAffected(res), nil
}

func (r *Repository[T]) Delete(id string) (int64, error) {
	q := "DELETE FROM " + TableName + " WHERE id = $1"
	res, err := r.db.Exec(q, id)
	if err != nil {
		return -1, err
	}
	return sqlxpkg.GetRowsAffected(res), nil
}
