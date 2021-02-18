package handler

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//Config struct
type Config struct {
	Server   string
	User     string
	Password string
	Port     string
	Database string
}

//Initialize func
func (c *Config) Initialize() {
	c.Server = os.Getenv("SERVER_SIDOC")
	c.User = os.Getenv("USER_SIDOC")
	c.Password = os.Getenv("PASSWORD_SIDOC")
	c.Port = os.Getenv("PORT_SIDOC")
	c.Database = os.Getenv("DATABASE_SIDOC")
}

//Connect to DB
func (c *Config) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable",
		c.User,
		c.Password,
		c.Server,
		c.Port,
		c.Database,
	)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		return nil, err
	}
	fmt.Println("conectado a la base de datos")
	return db, nil
}
