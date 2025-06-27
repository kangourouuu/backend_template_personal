package sqlclient

import (
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	MYSQL = "mysql"
	POSTGRESQL = "postgresql"
)

type IGormSqlClientConn interface {
	GetDB() *gorm.DB
	GetDriver() string
}

type GormSqlConfig struct {
	Driver   string
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

type GormSqlClientConn struct {
	GormSqlConfig
	DB *gorm.DB
}

func NewGormSqlClient(config GormSqlConfig) IGormSqlClientConn {
	client := &GormSqlClientConn{}
	client.GormSqlConfig = config
	if err := client.Connect(); err != nil {
		log.Fatal(err)
		return nil
	}
	// if err := client.DB.Ping(); err != nil {
	// 	log.Fatal(err)
	// 	return nil
	// }

	return client
}

func (c *GormSqlClientConn) Connect() (err error) {
	switch c.Driver {
	case MYSQL:
		return nil
	case POSTGRESQL:
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
			c.Host,
			c.Username,
			c.Password,
			c.Database,
			c.Port,
		)
		
		maxRetries := 5
		for i := 0; i < maxRetries; i++ {
			db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err == nil {
				c.DB = db
				log.Printf("Successfully connected to database on attempt %d", i+1)
				return nil
			}

			log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
			if i < maxRetries-1 {
				waitTime := time.Duration(i+1) * 2 * time.Second
				log.Printf("Retrying in %v...", waitTime)
				time.Sleep(waitTime)
			}
		}

		log.Fatal("Failed to connect to DB after all retries")
		return err
	default:
		log.Fatal("driver is missing")
		return errors.New("driver is missing")
	}
}

func (c *GormSqlClientConn) GetDB() *gorm.DB {
	return c.DB
}

func (c *GormSqlClientConn) GetDriver() string {
	return c.Driver
}
