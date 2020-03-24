package db

import (
	"fmt"
	"log"

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
	cs := connectionString(c)
	log.Println(cs)
	DB, err := gorm.Open(c.Driver, cs)

	return &Database{DB}, err
}

func (d *Database) Info() {
	log.Println("Successfully established connection with database")
}

func connectionString(c *Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.Name)
}
