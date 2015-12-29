package mssql

import (
	"fmt"
	"strings"
)

// Options custom options for mssql connection
type Options struct {
	Server   string
	Port     int
	UserID   string
	Password string
	Database string
}

// MsSQLConnString creates a connection string for mssql
func (opts *Options) MsSQLConnString() string {
	args := []string{}

	server := opts.Server
	port := opts.Port
	userid := opts.UserID
	password := opts.Password

	if server == "" {
		server = `127.0.0.1\SQLExpress`
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
