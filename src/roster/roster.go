package main

import (
	"database/sql"
	"fmt"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gocraft/web"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"public"
	"strings"
)

type Context struct {
	HelloCount int
}

func (c *Context) SetHelloCount(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HelloCount = 3
	next(rw, req)
}

func (c *Context) SayHello(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, strings.Repeat("Hello ", c.HelloCount), "World!")
}

var port string

func init() {
	_, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
	}

	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
}

func main() {
	router := web.New(Context{}).
		Middleware(web.LoggerMiddleware).
		Middleware(web.ShowErrorsMiddleware).
		Middleware(web.StaticMiddlewareFromDir(
		&assetfs.AssetFS{public.Asset, public.AssetDir, "assets"},
	)).
		Middleware((*Context).SetHelloCount).
		Get("/foo", (*Context).SayHello)

	err := http.ListenAndServe(":"+port, router)
	panic(err)
}
