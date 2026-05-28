package config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	DBPlatformPostgres = `Postgres`
	DBPlatformMySQL    = `MySQL`
)

type DatabasePlatform struct {
	MySQL    DatabaseInstance `mapstructure:"mysql" json:"mysql"`
	Postgres DatabaseInstance `mapstructure:"postgres" json:"postgres"`
}

type DatabaseInstance struct {
	Read            Database `mapstructure:"read" json:"read"`
	Write           Database `mapstructure:"write" json:"write"`
	Schema          string   `mapstructure:"schema" json:"schema"`
	SchemaChameleon string   `mapstructure:"schema_chameleon" json:"schema_chameleon"`
	LogLevel        string   `mapstructure:"log_level" json:"log_level"`
}

type Database struct {
	DSN          string        `mapstructure:"dsn" json:"dsn"`
	Username     string        `mapstructure:"username" json:"username"`
	Password     string        `mapstructure:"password" json:"password"`
	URL          string        `mapstructure:"url" json:"url"`
	Port         int           `mapstructure:"port" json:"port"`
	Name         string        `mapstructure:"name" json:"name"`
	Schema       string        `mapstructure:"schema" json:"schema"`
	Flavor       string        `mapstructure:"flavor" json:"flavor"`
	MaxIdleConns int           `mapstructure:"max_idle_conns" json:"max_idle_conns"`
	MaxOpenConns int           `mapstructure:"max_open_conns" json:"max_open_conns"`
	MaxLifetime  time.Duration `mapstructure:"max_lifetime" json:"max_lifetime"`
	Location     string        `mapstructure:"location" json:"location"`
	Timeout      time.Duration `mapstructure:"timeout" json:"timeout"`
}

func (m *Config) Postgres() *DatabaseInstance { return &m.Database.Postgres }
func (m *Config) MySQL() *DatabaseInstance    { return &m.Database.MySQL }

func LoadDatabaseConfig() DatabasePlatform {
	return DatabasePlatform{
		MySQL: DatabaseInstance{
			Read: Database{
				Username:     viper.GetString("DB_MYSQL_USER"),
				Password:     viper.GetString("DB_MYSQL_PASSWORD"),
				URL:          viper.GetString("DB_MYSQL_HOST"),
				Port:         viper.GetInt("DB_MYSQL_PORT"),
				Name:         viper.GetString("DB_MYSQL_NAME"),
				MaxIdleConns: viper.GetInt("DB_MYSQL_MAX_IDLE_CONNS"),
				MaxOpenConns: viper.GetInt("DB_MYSQL_MAX_OPEN_CONNS"),
				MaxLifetime:  viper.GetDuration("DB_MYSQL_MAX_LIFE_TIME"),
				Flavor:       viper.GetString("DB_MYSQL_FLAVOR"),
				Location:     viper.GetString("DB_MYSQL_LOCATION"),
				Timeout:      viper.GetDuration("DB_MYSQL_TIMEOUT"),
			},
			Write: Database{
				Username:     viper.GetString("DB_MYSQL_USER"),
				Password:     viper.GetString("DB_MYSQL_PASSWORD"),
				URL:          viper.GetString("DB_MYSQL_HOST"),
				Port:         viper.GetInt("DB_MYSQL_PORT"),
				Name:         viper.GetString("DB_MYSQL_NAME"),
				MaxIdleConns: viper.GetInt("DB_MYSQL_MAX_IDLE_CONNS"),
				MaxOpenConns: viper.GetInt("DB_MYSQL_MAX_OPEN_CONNS"),
				MaxLifetime:  viper.GetDuration("DB_MYSQL_MAX_LIFE_TIME"),
				Flavor:       viper.GetString("DB_MYSQL_FLAVOR"),
				Location:     viper.GetString("DB_MYSQL_LOCATION"),
				Timeout:      viper.GetDuration("DB_MYSQL_TIMEOUT"),
			},
		},
		Postgres: DatabaseInstance{
			Read: Database{
				DSN:          viper.GetString("DB_POSTGRES_DSN"),
				Username:     viper.GetString("DB_POSTGRES_USER"),
				Password:     viper.GetString("DB_POSTGRES_PASSWORD"),
				URL:          viper.GetString("DB_POSTGRES_HOST"),
				Port:         viper.GetInt("DB_POSTGRES_PORT"),
				Name:         viper.GetString("DB_POSTGRES_NAME"),
				Schema:       viper.GetString("DB_POSTGRES_SCHEMA"),
				MaxIdleConns: viper.GetInt("DB_POSTGRES_MAX_IDLE_CONNS"),
				MaxOpenConns: viper.GetInt("DB_POSTGRES_MAX_OPEN_CONNS"),
				MaxLifetime:  viper.GetDuration("DB_POSTGRES_MAX_LIFE_TIME"),
				Flavor:       viper.GetString("DB_POSTGRES_FLAVOR"),
				Location:     viper.GetString("DB_POSTGRES_LOCATION"),
				Timeout:      viper.GetDuration("DB_POSTGRES_TIMEOUT"),
			},
			Write: Database{
				DSN:          viper.GetString("DB_POSTGRES_DSN"),
				Username:     viper.GetString("DB_POSTGRES_USER"),
				Password:     viper.GetString("DB_POSTGRES_PASSWORD"),
				URL:          viper.GetString("DB_POSTGRES_HOST"),
				Port:         viper.GetInt("DB_POSTGRES_PORT"),
				Name:         viper.GetString("DB_POSTGRES_NAME"),
				Schema:       viper.GetString("DB_POSTGRES_SCHEMA"),
				MaxIdleConns: viper.GetInt("DB_POSTGRES_MAX_IDLE_CONNS"),
				MaxOpenConns: viper.GetInt("DB_POSTGRES_MAX_OPEN_CONNS"),
				MaxLifetime:  viper.GetDuration("DB_POSTGRES_MAX_LIFE_TIME"),
				Flavor:       viper.GetString("DB_POSTGRES_FLAVOR"),
				Location:     viper.GetString("DB_POSTGRES_LOCATION"),
				Timeout:      viper.GetDuration("DB_POSTGRES_TIMEOUT"),
			},
		},
	}

}
