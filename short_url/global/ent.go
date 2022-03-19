package global

import (
	"context"
	"log"
	"time"

	"shor_url/ent"
	"shor_url/ent/migrate"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DBClient *ent.Client

func initMySQL() {
	drv, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/shor_url?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatalln("Open err: ", err)
	}

	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	DBClient = ent.NewClient(ent.Driver(drv))
	createSchema(DBClient)
}

func createSchema(client *ent.Client) {
	err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(false),
	)

	if err != nil {
		log.Fatalln("Schema Create err", err)
	}
}
