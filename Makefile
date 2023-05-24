# note: call scripts from /scripts
.DEFAULT_GOAL := all

GOBUILD=GOOS=linux GOARCH=amd64 go build
GOSTATICBUILD=CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -extldflags "-static"'
# Clean all binaries
clean:
	rm -f build/package/*

go-update:
	go get -u ./...

# begin of build script
build: smtp imap

imap:
	$(GOBUILD) -o build/package/imap-proxy cmd/imap-proxy/*.go

smtp:
	$(GOBUILD) -o build/package/smtp-proxy cmd/smtp-proxy/*.go
