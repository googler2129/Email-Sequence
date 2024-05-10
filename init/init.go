package appinit

import (
	"context"
	"fmt"
	"github.com/depender/email-sequence-service/pkg/config"
	"github.com/depender/email-sequence-service/pkg/db"
	log "github.com/depender/email-sequence-service/pkg/logger"
	newrelic2 "github.com/depender/email-sequence-service/pkg/newrelic"
	_ "github.com/newrelic/go-agent/v3/integrations/nrpq"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func init() {
	err := log.NewLogger()
	if err != nil {
		panic(err.Error())
	}
}

func Initialise(ctx context.Context) {
	initNewrelic(ctx)
	initDbConnection(ctx)
}

func initNewrelic(ctx context.Context) {
	application, err := newrelic.NewApplication(
		newrelic.ConfigAppName(config.GetString(ctx, "newrelic.appName")),
		newrelic.ConfigLicense(config.GetString(ctx, "newrelic.licence")),
		newrelic.ConfigDistributedTracerEnabled(config.GetBool(ctx, "newrelic.distributedTracer")),
		newrelic.ConfigEnabled(config.GetBool(ctx, "newrelic.enabled")),
		func(c *newrelic.Config) {
			c.CrossApplicationTracer.Enabled = config.GetBool(ctx, "newrelic.crossApplicationTracer")
		},
		func(c *newrelic.Config) {
			c.ErrorCollector.RecordPanics = true
		},
	)

	if err != nil {
		panic("Could not initialize newrelic: " + err.Error())
	}

	newrelic2.SetNewrelicApplication(application)
}

func initDbConnection(ctx context.Context) {
	maxOpenConnections := config.GetInt(ctx, "postgresql.maxOpenConns")
	maxIdleConnections := config.GetInt(ctx, "postgresql.maxIdleConns")

	database := config.GetString(ctx, "postgresql.database")
	connIdleTimeout := 10 * time.Minute

	// Read Write endpoint config
	mysqlWriteServer := config.GetString(ctx, "postgresql.master.host")
	mysqlWritePort := config.GetString(ctx, "postgresql.master.port")
	mysqlWritePassword := config.GetString(ctx, "postgresql.master.password")
	mysqlWriterUsername := config.GetString(ctx, "postgresql.master.username")

	debugMode := config.GetBool(ctx, "postgresql.debugMode")

	gormLogger := logger.Default
	if debugMode {
		gormLogger = gormLogger.LogMode(logger.Info)
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", mysqlWriteServer, mysqlWritePort, mysqlWriterUsername, mysqlWritePassword, database,
	)

	gormDB, err := gorm.Open(postgres.Dialector{
		Config: &postgres.Config{
			DriverName: "nrpostgres", // For newrelic tracking
			DSN:        dsn,
		},
	}, &gorm.Config{
		Logger:                 gormLogger,
		SkipDefaultTransaction: false,
		PrepareStmt:            false,
	})
	if err != nil {
		log.Errorf("Unable to make gorm connection | Error: %v", err)
		panic("Unable to make gorm connection")
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Errorf("Unable to get sqlDB from gormDB | Error: %v", err)
		panic("Unable to get sqlDB from gormDB")
	}

	sqlDB.SetMaxOpenConns(maxOpenConnections)
	sqlDB.SetMaxIdleConns(maxIdleConnections)
	sqlDB.SetConnMaxLifetime(connIdleTimeout)

	retries := 3
	for retries > 0 {
		err = sqlDB.Ping()
		if err != nil {
			log.Errorf("Unable to ping database server: %s, waiting 2 seconds before trying %d more times", err.Error(), retries)
			time.Sleep(time.Second * 2)
			retries--
		} else {
			err = nil
			break
		}
	}
	if err != nil {
		panic("Db not initialised")
	}

	db.SetCluster(gormDB)
}
