package mssql

import (
	"database/sql"

	"github.com/Sirupsen/logrus"
	"github.com/webcanvas/pinch/shared/models"
)

// Runner holds information required for a run
type Runner struct {
	opts Options
	conn sql.DB

	Version string
}

// Run setup the service
func (r *Runner) Run() (result models.Result, err error) {

	logrus.WithFields(logrus.Fields{"version": r.Version, "database": r.opts.Database}).Debug("We have a version of mssql")

	// TODO drop database
	// TODO create database
	// TODO create local user accounts

	// TODO return service information

	return
}

// CreateRunner creates a new runner
func CreateRunner(opts Options) (*Runner, error) {

	connstr := opts.MsSQLConnString()
	logrus.WithFields(logrus.Fields{"connstr": connstr}).Debug("What's the connection?")

	db, err := sql.Open("mssql", connstr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	var version string
	err = db.QueryRow("SELECT @@VERSION as version").Scan(&version)
	if err != nil {
		return nil, err
	}

	return &Runner{
		opts:    opts,
		conn:    *db,
		Version: version,
	}, nil
}
