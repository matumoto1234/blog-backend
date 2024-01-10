package main

import (
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"

	"github.com/matumoto1234/blog-backend/app/infra"
	"github.com/matumoto1234/blog-backend/app/infra/api"
	"github.com/matumoto1234/blog-backend/app/middleware/applog"
	"github.com/matumoto1234/blog-backend/app/ui"
	"github.com/matumoto1234/blog-backend/app/usecase"
	"github.com/matumoto1234/blog-backend/db"
	"github.com/matumoto1234/blog-backend/gen/article"
	"github.com/matumoto1234/blog-backend/gen/http/article/server"
)

// https://www.prisma.io/docs/orm/prisma-client/setup-and-configuration/databases-connections#prismaclient-in-serverless-environments
var client = db.NewClient()

func main() {
	logger := applog.New()
	logger.Info("server started")

	if err := client.Connect(); err != nil {
		logger.Error("connect failed: " + err.Error())
		os.Exit(1)
	}
	defer client.Disconnect()

	gateway := api.NewGateway(http.DefaultClient)
	repo := infra.NewArticleRepository(client, gateway)
	u := usecase.NewArticleUseCase(repo)
	h := ui.NewArticleHandler(u)

	endpoints := article.NewEndpoints(h)
	mux := goahttp.NewMuxer()
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder
	svr := server.New(endpoints, mux, dec, enc, nil, nil)
	server.Mount(mux, svr)
	httpsvr := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	logger.Info(fmt.Sprintf("listening on %s...", httpsvr.Addr))

	if err := httpsvr.ListenAndServe(); err != nil {
		logger.Error(err.Error())
	}
}
