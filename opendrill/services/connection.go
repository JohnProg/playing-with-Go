package services

import (
	"../app/models"
	"gopkg.in/mgo.v2"
	"log"
)

// Connection represents the database session
type Connection struct {
	Session *mgo.Session
	Db      *mgo.Database
}

// NewDatabaseConn initializes the database connection
func NewDatabaseConn(config *Config) (*Connection, error) {
	conn := new(Connection)
	var err error

	if conn.Session, err = mgo.Dial(config.DatabaseUrl); err != nil {
		log.Fatalf("Error connecting to MongoDB '%s'", config.DatabaseUrl)
		return nil, err
	}
	conn.Session.SetMode(mgo.Strong, true)
	conn.Db = conn.Session.DB(config.DatabaseName)
	models.SetDB(conn.Db)

	return conn, nil
}

// Close closes mongodb open connection
func (c *Connection) Close() {
	c.Session.Close()
}
