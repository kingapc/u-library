package model

import (
	"database/sql"
	"time"
)

type Entity struct {
	ID        string       `json:"id"`
	CreatedBy string       `json:"created_by"`
	UpdatedBy string       `json:"updated_by"`
	DeletedBy string       `json:"deleted_by"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
	IsDeleted string       `json:"is_deleted"`
}
