package dal

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"server/domain"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

const host = "localhost"
const port = 5432
const user = "postgres"
const password = "Cc030789"
const dbname = "postgres"

type DatabaseConnector struct {
	Database *sqlx.DB
}

func NewDatabaseConnector() *DatabaseConnector {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//s:=bindata.Resource(migrations.AssetNames(), migrations.Asset)
	//runDBMigrate("postgres://postgres:Cc030789@localhost:5432/testdb?sslmode=disable", s)

	return &DatabaseConnector{
		db,
	}
}

func (d *DatabaseConnector) GetInitialLinkFromStorage(id string) (string, error) {
	var urlForGetting domain.UrlDTO
	var err error
	err = d.Database.Get(&urlForGetting, "select * from url where hash = $1", id)
	if err != nil {
		return "", err
	}
	return urlForGetting.InitialURL, nil
}

func (d *DatabaseConnector) SaveInitialLinkToStorage(url string, id string) error {
	urlData := domain.UrlDTO{
		id,
		url,
	}

	_, err := d.Database.NamedExec("insert into url (hash, initial_url) values (:hash, :initial_url) on conflict (hash) do nothing",
		urlData)
	if err != nil {
		return err
	}

	return nil
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
