package dal

import (
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // for migrations?
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // for database connections?
	"log"
	"server/config"
	"server/domain"
)

type DatabaseConnector struct {
	Database *sqlx.DB
}

func NewDatabaseConnector(sConfig *config.ServerConfig) *DatabaseConnector {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		sConfig.Database.DBHost, sConfig.Database.DBPort, sConfig.Database.DBUser,
		sConfig.Database.DBPassword, sConfig.Database.DBName)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//s:=bindata.Resource(migrations.AssetNames(), migrations.Asset)
	//runDBMigrate("postgres://postgres:Cc030789@localhost:5432/testdb?sslmode=disable", s)

	return &DatabaseConnector{
		Database: db,
	}
}

func (d *DatabaseConnector) GetInitialLinkFromStorage(id string) (string, error) {
	var urlForGetting domain.URLDTO
	var err  = d.Database.Get(&urlForGetting, "select * from url where hash = $1", id)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return urlForGetting.InitialURL, nil
}

func (d *DatabaseConnector) SaveInitialLinkToStorage(url, id string) error {
	urlData := domain.URLDTO{
		ID:         id,
		InitialURL: url,
	}

	_, err := d.Database.NamedExec("insert into url (hash, initial_url) values (:hash, :initial_url) on conflict (hash) do nothing",
		urlData)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (d *DatabaseConnector) CloseDatabaseConnection() {
	err := d.Database.Close()
	if err != nil {
		log.Println(err)
	}
}

//func runDBMigrate(dsn string, source *bindata.AssetSource)  {
//	d, err := bindata.WithInstance(source)
//	if err != nil {
//		panic(err)
//	}
//
//	m, err := migrate.NewWithSourceInstance("go-bindata", d, dsn)
//	if err != nil {
//		panic(err)
//	}
//
//	if err = m.Up(); err != nil {
//		if err == migrate.ErrNoChange {
//			fmt.Println(err)
//		} else {
//			panic(err)
//		}
//	}
//}
