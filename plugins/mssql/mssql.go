package mssql

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/mitchellh/mapstructure"
	"github.com/webcanvas/pinch/plugins"
	"github.com/webcanvas/pinch/shared/commanders"
	"github.com/webcanvas/pinch/shared/models"
)

var versionex = regexp.MustCompile("[0-9.]+")

type mssql struct {
	commander *commanders.Commander
	Version   string
}

type serviceopts struct {
	Server   string
	Port     int
	UserID   string
	Password string
	Database string
}

func (opts *serviceopts) MsSQLConnString() string {
	args := []string{}

	server := opts.Server
	port := opts.Port
	userid := opts.UserID
	password := opts.Password

	if server == "" {
		server = "127.0.0.1\\SQLExpress"
	}
	if port == 0 {
		port = 1433
	}

	args = append(args, fmt.Sprintf("server=%s", server))
	args = append(args, fmt.Sprintf("port=%d", port))
	args = append(args, fmt.Sprintf("user id=%s", userid))
	args = append(args, fmt.Sprintf("password=%s", password))

	return strings.Join(args, ";")
}

// Setup runs all the pre plugin stuff. IE finding versions
func (g *mssql) Setup() error {
	return nil
}

// Ensure setups the service
func (g *mssql) Ensure(data map[string]string) (result models.Result, err error) {
	opts := new(serviceopts)
	mapstructure.Decode(data, &opts)

	connstr := opts.MsSQLConnString()
	logrus.WithFields(logrus.Fields{"connstr": connstr}).Debug("What's the connection?")

	db, err := sql.Open("mssql", connstr)
	if err != nil {
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return
	}

	var version string
	err = db.QueryRow("SELECT @@VERSION as version").Scan(&version)
	if err != nil {
		return
	}

	logrus.WithFields(logrus.Fields{"version": version, "database": opts.Database}).Debug("We have a version of mssql")

	// TODO drop database
	// TODO create database
	// TODO create local user accounts

	// TODO return service information

	return
}

func init() {
	g := &mssql{}
	plugins.RegisterServicePlugin("mssql", g)
}
