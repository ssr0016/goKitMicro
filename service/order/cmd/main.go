package main

import (
	"database/sql"
	"flag"
	"net/http"
	"os"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "order",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var db *sql.DB{
		var err error
		// Connect to the "ordersdb" database
		db, err = sql.Open("postgres",
	"host=localhost port=5432 user=postgres password=postgres dbname=ordersdb sslmode=disable")
	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}
 }

 var h http.Handler
 {
	endpoints := transport.MakeEndpoints(svc)
	h = httptransport.NewService(endpoints, logger)
 }

 errs := make(chan error)
 go func(){
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	err <- fmt.Errorf("%s", <-c)
 }

 go func(){
	level.Info(logger).Log("tranport", "HTTP", "addr", *httpAddr)
	server := &http.Server{
		Addr: *httpAddr,
		Handler: h,
	}
	errs <- server.ListenAndServe()
 }()

 level.Error(logger).Log("exit", <-errs)



}
