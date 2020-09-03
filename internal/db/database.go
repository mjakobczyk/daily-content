package db

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"

	// PostgreSQL driver
	_ "github.com/lib/pq"
)

// Database of application.
type Database struct {
	*gorm.DB
}

// NewDatabase is a default constuctor of Database.
func NewDatabase(c *Config) (*Database, error) {
	connection, err := connect(c.Driver, connectionString(c))

	return &Database{connection}, err
}

// Info that is printed when successful connection is established with database.
func (d *Database) Info() {
	log.Println("Successfully established connection with database")
}

func connect(driver string, connection string) (*gorm.DB, error) {
	attempts := 10
	var DB *gorm.DB
	var err error
	log.Println(connection)

	for attempts > 0 {
		DB, err = gorm.Open(driver, connection)
		if err != nil {
			log.Printf("Could not connect to database, %d more attempts...", attempts)
			attempts--
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}

	return DB, err
}

func connectionString(c *Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.Name)
}
