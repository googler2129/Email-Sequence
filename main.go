package main

import (
	"context"
	"errors"
	"flag"
	"github.com/depender/email-sequence-service/init"
	"github.com/depender/email-sequence-service/pkg/config"
	log "github.com/depender/email-sequence-service/pkg/logger"
	"github.com/depender/email-sequence-service/router"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"net/http"
	"strings"
	"time"
)

const (
	modeHttp      = "http"
	modeMigration = "migration"
)

func main() {
	err := config.Init(10 * time.Second)
	if err != nil {
		panic(err)
	}

	ctx, err := config.TODOContext()
	if err != nil {
		panic(err)
	}

	appinit.Initialise(ctx)

	var mode string
	flag.StringVar(
		&mode,
		"mode",
		modeHttp,
		"Pass the flag to run in different modes (worker or default)",
	)
	flag.Parse()

	switch strings.ToLower(mode) {
	case modeHttp:
		runHttpServer(ctx)
	case modeMigration:
		runMigration(ctx)
	default:
		runHttpServer(ctx)
	}
}

func runHttpServer(ctx context.Context) {
	// Setting gin to releaseMode
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	s := &http.Server{
		Addr:         ":8085",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := router.Initialize(ctx, r)
	if err != nil {
		panic(err)
	}

	err = s.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error(err)
	}
}

func runMigration(ctx context.Context) {
	database := config.GetString(ctx, "postgresql.database")
	mysqlWriteHost := config.GetString(ctx, "postgresql.master.host")
	mysqlWritePort := config.GetString(ctx, "postgresql.master.port")
	mysqlWritePassword := config.GetString(ctx, "postgresql.master.password")
	mysqlWriterUsername := config.GetString(ctx, "postgresql.master.username")

	m, err := migrate.New("file://deployment/migration", "postgres://"+mysqlWriteHost+":"+mysqlWritePort+"/"+database+"?user="+mysqlWriterUsername+"&password="+mysqlWritePassword+"&sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic(err)
	}
}
