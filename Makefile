VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")
BUILDDATE := $(shell date +%FT%T%Z)

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.GitCommit=$(BUILD) -X=main.BuildDate=$(BUILDDATE)"

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

build:
	@go build $(LDFLAGS) -o bin/$(PROJECTNAME)
