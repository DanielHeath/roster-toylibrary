package web

import (
	"encoding/csv"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gocraft/web"
	"log"
	"net/http"
	"public"
)

type Context struct {
	MultipartContext
	CsvFileContext
}

func Router() http.Handler {
	router := web.New(Context{}).
		Middleware(web.LoggerMiddleware).
		Middleware(web.ShowErrorsMiddleware).
		Middleware(web.StaticMiddlewareFromDir(
		&assetfs.AssetFS{public.Asset, public.AssetDir, "assets"},
	)).
		Post("/upload", (*Context).ReceiveUpload)

	return router
}

// TODO: Middleware to identify valid kinds of file
// to upload, parse them && add to context.
func (c *Context) ReceiveUpload(rw web.ResponseWriter, r *web.Request) {
	f, err := c.GetSingleFile(r, "members.csv") // see also sessions.csv
	panicIf(err)
	defer f.Close()
	records, err := csv.NewReader(f).ReadAll()
	panicIf(err)
	for _, record := range records {
		log.Println(record)
	}
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
