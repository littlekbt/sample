GO=go
GOBUILD=go build
BINARY=bin/server
DOCKER_BUILD=docker build
DOCKER_RMI=docker rmi
IMAGE=nginx-server-push-sample

build:
	docker build . -t $(IMAGE) && GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY) .

clean:
		rm -rf bin && $(DOCKER_RMI) $(IMAGE)
