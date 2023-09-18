package infrastructure

import (
	"github.com/Renos-id/go-starter-template/database"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

func InitQueryBuilder() goqu.Database {
	db := database.Open()
	// goqu.SetDefaultPrepared(true)
	dialect := goqu.Dialect("postgres")
	goquDB := dialect.DB(db)
	return *goquDB
}
