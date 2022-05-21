package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

	"github.com/AbdelouahabMbarki/Product/config"
	"github.com/AbdelouahabMbarki/Product/product"
	_ "github.com/lib/pq"

	"github.com/go-kit/kit/log"

	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log/level"
)

func main() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "product	",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	config, err := config.LoadConfig(".")
	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(1)
	}
	var httpAddr = flag.String("http", config.Server.Port, "http listen address")
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var db *sql.DB
	{
		var err error

		db, err = sql.Open("postgres", config.Database.Url)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(1)
		}

	}

	flag.Parse()
	ctx := context.Background()
	var srv product.Service
	{
		repository := product.NewRepo(db, logger)

		srv = product.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := product.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := product.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
