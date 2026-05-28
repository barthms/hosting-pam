package database

import (
	"fmt"
	"net/url"
	"strconv"
)

var typeDBSchemaMapping = map[DBType]string{
	Mysql:    mysqlSchema,
	Postgres: postgresSchema,
}

func GetURL(arg *Args) (res *url.URL) {
	if arg == nil || !arg.IsValid() {
		return nil
	}

	scheme := typeDBSchemaMapping[arg.DBType]
	if scheme == "" {
		scheme = mysqlSchema
	}

	res = &url.URL{
		Scheme: scheme,
		User:   url.UserPassword(arg.Username, arg.Password),
		Path:   arg.Database,
		Host:   fmt.Sprintf("%s:%d", arg.Host, arg.Port),
	}

	isUrlValNil := arg.Values == nil
	if arg.Values == nil {
		arg.Values = make(url.Values)
	}

	switch scheme {
	case postgresSchema:
		if isUrlValNil {
			arg.Values.Add("sslmode", "disable")
			arg.Values.Add("extra_float_digits", "-1")
		}

		if arg.Schema != "" {
			arg.Values.Add("search_path", arg.Schema)
		}

		timeout := int32(arg.Timeout.Seconds())
		arg.Values.Add("connect_timeout", strconv.FormatInt(int64(timeout), 10))
	case mysqlSchema:
		if isUrlValNil {
			arg.Values.Add("autocommit", "true")
			arg.Values.Add("parseTime", "true")
		}

		if arg.Location != "" {
			arg.Values.Add("loc", arg.Location)
		}
	}

	res.RawQuery = arg.Values.Encode()
	return
}

func GetURLString(arg *Args) (res string) {
	if arg != nil && arg.DSN != "" {
		return arg.DSN
	}

	url := GetURL(arg)
	if url == nil {
		return
	}

	res = url.String()
	return
}
