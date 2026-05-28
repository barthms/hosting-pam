package database

import (
	"net/url"
	"time"
)

const (
	postgresDriverName string = "pgx"
	postgresSchema     string = "postgres"

	mysqlDriverName string = "mysql"
	mysqlSchema     string = "mysql"
)

const (
	// Postgres ...
	Postgres DBType = "postgres"
	// Mysql ...
	Mysql DBType = "mysql"

	// WriteConn write connection type
	WriteConn ConnType = "write"
	// ReadConn read connection type
	ReadConn ConnType = "read"
)

type (
	DBType   string
	ConnType string
	Args     struct {
		DSN             string
		Username        string
		Password        string
		Host            string
		Port            int
		Database        string
		Schema          string
		Flavor          string
		MaxIdleConns    int
		MaxOpenConns    int
		ConnMaxLifetime time.Duration
		Location        string
		Timeout         time.Duration

		// Other params
		ConnType ConnType
		DBType   DBType
		Values   url.Values
	}
)

func (a *Args) IsValid() bool {
	if a.DSN != "" {
		return a.DBType != ""
	}

	return a.Username != "" && a.Password != "" && a.Host != "" && a.Port != 0 && a.Database != "" && a.DBType != ""
}
