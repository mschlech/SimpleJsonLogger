# Go parameters
GOCMD=go
GOBUILD=CGO_ENABLED=0 GOOS=linux $(GOCMD) build -installsuffix cgo
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

HAVE_DOCKER=$(shell docker version >/dev/null 2>&1 && echo have_docker)

APPS := $(shell ls -1 cmd)

PRIVATE_LIB=$(shell test -d private/lib && echo private/lib)

all:: init test

init:
	mkdir -p bin container gen/lib/grpc

define make-cmd-target =
bin/$1: $(shell find cmd private/app/$1 $(PRIVATE_LIB) gen -name '*.go')
	$(GOBUILD) -o bin/$1 -v ./cmd/$1
all:: bin/$1

ifdef HAVE_DOCKER
container/Dockerfile-$1: Dockerfile.template
	sed <Dockerfile.template >container/Dockerfile-$1 "s,APP,$1,g"

container/$1: bin/$1 container/Dockerfile-$1
	docker build -f container/Dockerfile-$1 -t lcm/$1 .
	touch container/$1

container:: container/$1

endif
endef

$(foreach _,$(APPS),$(eval $(call make-cmd-target,$_)))

build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN) ./...
	rm -rf bin container gen

deps:
	$(GOGET) github.com/markbates/goth
	$(GOGET) github.com/markbates/pop
