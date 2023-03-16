package model

import (
	"strings"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound

func getValuesMark(n int) string {
	return strings.Repeat("?,", n-1) + "?"
}
