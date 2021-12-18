package repository

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
	"log"
	"testproyect/entity"
	"time"
)

const (
	mySQLURL    = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
	displaySQL  = "DBName %s, %s@%s:/%s:%s"
	psqldialect = "postgres"
)

const (
	UniqueViolation = "unique_violation"
)

type Base struct {
	SQLDB  *gorm.DB
	DBName string
}

func NewPostgresRepository() Repository {
	return &Repo{}
}

func (r *Repo) Init(dbIP, dbPort, dbUser, dbPass, dbName string, retryCount int) error {
	err := r.Base.ConnectInit(dbIP, dbPort, dbUser, dbPass, dbName, retryCount)
	if err != nil {
		return err
	}

	err = r.Base.Migrate()
	if err != nil {
		return err
	}
	return nil
}

func (b *Base) ConnectInit(dbIP, dbPort, dbUser, dbPass, dbName string, retryCount int) error {
	if len(dbIP) == 0 {
		return errors.New("missing database ip address configuration")
	}

	url, displayURL := buildURLs(dbIP, dbPort, dbUser, dbPass, dbName)
	dbc, err := connect(url, dbName, retryCount)
	if err != nil {
		log.Fatalf("db connect failure to: %s  Reason: %s\n", displayURL, err.Error())
	}

	b.SQLDB = dbc
	b.DBName = dbName
	return nil
}

func buildURLs(mySQLIP, port, user, pass, dbName string) (url string, displayURL string) {
	url = fmt.Sprintf(mySQLURL, mySQLIP, port, user, pass, dbName)
	displayURL = fmt.Sprintf(displaySQL, dbName, "****", "****", mySQLIP, port)
	return
}

func connect(url string, dbName string, retryCount int) (*gorm.DB, error) {
	var delay = time.Second * 2

	var dbc *gorm.DB
	var err error
	limit := retryCount >= 0
	for ; !limit || retryCount >= 0; retryCount-- {
		dbc, err = gorm.Open(psqldialect, url)
		if err != nil {
			fmt.Printf("Database connect failed, error: %v\n", err.Error())
			time.Sleep(delay)
			// backoff delay
			if delay < 30 {
				delay *= 2
			}
			continue
		}
		if dbc != nil {
			break
		}
	}
	if err != nil {
		return nil, err
	}

	// Catch the case where we failed to connect before we ran out of retries.
	if dbc == nil {
		log.Fatalln("error connecting to db [" + dbName + "]")
	}

	fmt.Println("connected to '" + dbName + "' database")

	return dbc, nil
}

func (b *Base) Migrate() error {
	if err := b.SQLDB.AutoMigrate(&entity.Beer{}).Error; err != nil {
		return err
	}
	return nil
}

func ValidatePSQLError(err error, keyError string) bool {
	if err, ok := err.(*pq.Error); ok {
		return err.Code.Name() == keyError
	}

	return false
}
