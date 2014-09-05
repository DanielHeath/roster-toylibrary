
.PHONY: bin/roster bin/go-bindata src/public

bin/roster: src src/public
	go build -o bin/roster roster

bin/go-bindata: src/github.com/jteeuwen/go-bindata
	go build -o bin/go-bindata github.com/jteeuwen/go-bindata/go-bindata

src/public: assets bin/go-bindata
	./bin/go-bindata -pkg public -o src/public/public.go assets/...
