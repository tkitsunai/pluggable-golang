# meta data
NAME := pacman
ENTRYPOINT := main.go

dev: dep test osx-build-development

dep:
	dep ensure

test:
	go test -v ./...

# Build for Darwin/Amd64
osx-build-development:
	GOOS=darwin GOARCH=amd64 go build -o bin/$(NAME)-darwin-amd64 ./$(ENTRYPOINT)