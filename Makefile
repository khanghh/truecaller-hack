server:
	go build -ldflags "-w -s \
		-X 'main.CommitHash=$(shell git rev-list -1 HEAD)' \
		-X 'main.Version=$(shell git describe --tags)' \
		-X 'main.BuiltTime=$(shell date)'" \
		-o bin/$@ cmd/server/main.go

all: server
