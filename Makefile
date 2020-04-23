IMAGE ?= thief

all: build

build:
	GOOS=linux GOARCH=amd64 go build -o bin/thief ./cmd/

docker: build
	docker build -t ${IMAGE} .

minikube: docker
	mkdir -p out/
	docker save -o out/${IMAGE}.tar ${IMAGE}
	scp -i $$(minikube ssh-key) out/${IMAGE}.tar docker@$$(minikube ip):/tmp
	minikube ssh -- sudo ctr cri load /tmp/${IMAGE}.tar
