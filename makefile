
GOPATH:=$(shell go env GOPATH)

.PHONY: run
run:build
	go run migration/init.go
	./dst/bin/app
.PHONY: build
build:
	rm -rf ./dst
	CGO_ENABLED=1 go build -o ./dst/bin/app ./*.go
.PHONY: migrateup
migrateup:
	go run migration/init.go migration/up.go
.PHONY: migratedown
migratedown:
	go run migration/init.go migration/down.go
