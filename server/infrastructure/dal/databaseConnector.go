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
