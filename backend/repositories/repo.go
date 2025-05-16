package repositories

import (
	"database/sql"
	"strings"
)

type AppRepository struct {
	db *sql.DB
}

// NewPostRepository creates a new repository
func NewAppRepository(db *sql.DB) *AppRepository {
	return &AppRepository{db: db}
}

// this struct will be used for queries
// specially for the categories filtering ...

type Query struct {
	b      strings.Builder
	params []any
}

func (q *Query) Query(s string) {
	q.b.WriteString(s)
}

func (q *Query) Param(val any) {
	q.b.WriteString("?")
	q.params = append(q.params, val)
}

func (q *Query) Get() (string, []any) {
	return q.b.String(), q.params
}
