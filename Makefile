APP?=projectone
PORT?=8000
test:
	go test ./... -v
build:
	go build ./...

CONTAINER_IMAGE?=docker.io/msarajam/${APP}


container: build
	docker build -t $(CONTAINER_IMAGE) .
push: container
	docker push $(CONTAINER_IMAGE)
