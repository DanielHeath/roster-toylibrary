package main

import (
	"database/sql"
	"encoding/csv"
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

const (
	Kilobyte       = 1024
	Megabyte       = 1024 * Kilobyte
	TimeFormat     = "3:04PM"
	DateFormat     = "Jan 02 2006"
	TimeDateFormat = DateFormat + " " + TimeFormat
)

type Context struct {
}

// TODO: Middleware to identify valid kinds of file
// to upload, parse them && add to context.
func (c *Context) ReceiveUpload(rw web.ResponseWriter, r *web.Request) {
	log.Println(string(r.URL.String()))
	err := r.ParseMultipartForm(10 * Megabyte)
	panicIf(err)
	m := r.MultipartForm
	header := m.File["members.csv"]
	if len(header) != 1 {
		panic("You must upload one file")
	}
	f, err := header[0].Open()
	panicIf(err)
	defer f.Close()
	records, err := csv.NewReader(f).ReadAll()
	panicIf(err)
	for _, record := range records {
		log.Println(record)
	}
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
		Post("/upload", (*Context).ReceiveUpload)

	err := http.ListenAndServe(":"+port, router)
	panic(err)
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
