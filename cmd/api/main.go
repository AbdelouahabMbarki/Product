package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/AbdelouahabMbarki/Product/config"
	"github.com/AbdelouahabMbarki/Product/product"
	_ "github.com/lib/pq"

	"github.com/go-kit/kit/log"

	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log/level"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	var db *mongo.Client

	{
		var err error
		db, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Database.Url))
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(1)
		}
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = db.Connect(ctx)
		if err != nil {
			level.Info(logger).Log("msg", err)
		}
		defer db.Disconnect(ctx)
	}

	flag.Parse()
	ctx := context.Background()
	var srv product.Service
	{
		repository := product.NewReponoSql(db, logger)

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
