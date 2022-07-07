package config

import (
	"flag"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type AppConfig struct {
	DbConfig   DbConfig
	HTTPServer HTTPServer
}

type DbConfig struct {
	Engine       *string
	ConnString   *string
	MaxOpenConns *int
	MaxIdleConns *int
}

type HTTPServer struct {
	ListenAddr          *string
	WriteTimeout        *time.Duration
	ReadTimeout         *time.Duration
	IdleTimeout         *time.Duration
	GracefulWaitTimeout *time.Duration
	MaxRequestSize      *int64
}

func BuildConfig() *AppConfig {
	fs := flag.NewFlagSet("app-config", flag.ExitOnError)
	config := AppConfig{}
	config.DbConfig.Engine = fs.String("dbEngine", "postgres", "DB Engine Name")
	config.DbConfig.ConnString = fs.String("dbConnString", "postgresql://postgres:postgres@0.0.0.0:5432/postgres?sslmode=disable", "Postgres DSN")
	config.DbConfig.MaxOpenConns = fs.Int("dbMaxOpenConns", 5, "Max Open Connections for DB")
	config.DbConfig.MaxIdleConns = fs.Int("dbMaxIdleConns", 5, "Max Idle Connections for DB")

	config.HTTPServer.ListenAddr = fs.String("listenAddr", "0.0.0.0:8080", "Address at which application runs")
	config.HTTPServer.WriteTimeout = fs.Duration("writeTimeout", 60*time.Second, "Write Timeout")
	config.HTTPServer.ReadTimeout = fs.Duration("readTimeout", 60*time.Second, "Read Timeout")
	config.HTTPServer.IdleTimeout = fs.Duration("idleTimeout", 60*time.Second, "Idle Timeout")
	config.HTTPServer.GracefulWaitTimeout = fs.Duration("gracefulWaitTimeout", 15*time.Second, "Graceful Wait Timeout")
	config.HTTPServer.MaxRequestSize = fs.Int64("maxRequestSize", 10485760, "Max Request Size")
	return &config
}
