package database

import (
	"context"
	"fmt"
	"max_inventory/settings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {

	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", s.DB.Username, s.DB.Password, s.DB.Host, s.DB.Port, s.DB.Database)

	return sqlx.ConnectContext(ctx, "mysql", conn)
}
