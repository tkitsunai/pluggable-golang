# meta data
NAME := pluggable
ENTRYPOINT := main.go

dev: dep test osx-build-development

dep:
	dep ensure

update:
	dep ensure -update

test:
	go test -v ./...

# Build for Darwin/Amd64
osx-build-development:
	GOOS=darwin GOARCH=amd64 go build -o bin/$(NAME)-darwin-amd64 ./$(ENTRYPOINT)