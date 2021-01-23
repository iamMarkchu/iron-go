package model

import "github.com/tal-tech/go-zero/core/stores/sqlx"

const (
	StatusOk      = 1
	StatusDeleted = 2
)

var ErrNotFound = sqlx.ErrNotFound
