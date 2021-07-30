package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/labstack/echo/v4"
	"server/app"
	"server/config"
	"server/infrastructure/dal"
	"server/infrastructure/dal/migrations"
	"server/infrastructure/httpcontrollers"
)

func main() {
	sConfig := config.NewServerConfig()

	linkValidator := app.NewLinkValidator()
	idValidator := app.NewIDValidator()
	hasher := app.NewLinkHasher(sConfig)

	storage := dal.NewDatabaseConnector(sConfig)
	defer storage.CloseDatabaseConnection()

	service := app.NewService(linkValidator, idValidator, hasher, storage, sConfig)

	getController := httpcontrollers.NewGetController(service)
	postController := httpcontrollers.NewPostController(service)

	s:=bindata.Resource(migrations.AssetNames(), migrations.Asset)
	runDBMigrate("postgres://postgres:Cc030789@localhost:5432/postgres?sslmode=disable", s)

	e := echo.New()
	e.GET("/:id", getController.Get)
	e.POST("/", postController.Post)
	e.Logger.Fatal(e.Start(sConfig.Host.ServerStartPort))
}

func runDBMigrate(dsn string, source *bindata.AssetSource)  {
	d, err := bindata.WithInstance(source)
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithSourceInstance("go-bindata", d, dsn)
	if err != nil {
		panic(err)
	}

	if err = m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println(err)
		} else {
			panic(err)
		}
	}
}
