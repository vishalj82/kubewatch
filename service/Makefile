DIR="$( cd "$( dirname "$0" )" && pwd )"
GOBIN="$DIR/dependencies/bin"
GOPATH="$DIR/dependencies/pkg" 
GO111MODULE=auto

docker: compile
	docker build -t vishalj82/kubewatch-service  .
	docker push vishalj82/kubewatch-service

compile:
	GOOS=linux GOARCH=386 go build -o bin/service-linux-386 webhookService.go reviewResponse.go scheme.go
