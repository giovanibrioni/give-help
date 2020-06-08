package runtime

import (
	"os"

	firebase "firebase.google.com/go"
	"github.com/alexwbaule/give-help/v2/internal/common"
	"github.com/alexwbaule/give-help/v2/internal/fireadmin"
	"github.com/alexwbaule/give-help/v2/internal/storage/connection"
)

// NewRuntime creates a new application level runtime that encapsulates the shared services for this application
func NewRuntime() (*Runtime, error) {

	c := connection.New(&common.DbConfig{
		DBName: os.Getenv("DATABASE_DBNAME"),
		Host:   os.Getenv("DATABASE_HOST"),
		Pass:   os.Getenv("DATABASE_PASS"),
		User:   os.Getenv("DATABASE_USER"),
	})

	rt := &Runtime{
		fbase:    fireadmin.InitializeAppWithServiceAccount(),
		database: c,
	}

	return rt, nil
}

// Runtime encapsulates the shared services for this application
type Runtime struct {
	fbase    *firebase.App
	database *connection.Connection
}

func (rt *Runtime) GetFirebase() *firebase.App {
	return rt.fbase
}

func (rt *Runtime) GetDatabase() *connection.Connection {
	return rt.database
}

func (rt *Runtime) CloseDatabase() {
	rt.database.Close()
}
