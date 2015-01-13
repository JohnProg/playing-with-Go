package app

import (
	"../router"
	"../services"
)

// App represents the application, it contains the instances of the Config and Connection services
// to be closed after running the application
type App struct {
	Config     *services.Config
	Connection *services.Connection
}

// NewApp creates a new application given a path to a config file
func NewApp(configPath string) (*App, error) {
	var err error

	// Create config service
	config, err := services.NewConfig(configPath)
	if err != nil {
		return nil, err
	}

	// Create database service
	conn, err := services.NewDatabaseConn(config)
	if err != nil {
		return nil, err
	}

	// Add routes
	router.Init()

	return &App{config, conn}, nil
}
