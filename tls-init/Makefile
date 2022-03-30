DIR="$( cd "$( dirname "$0" )" && pwd )"
GOBIN="$DIR/dependencies/bin"
GOPATH="$DIR/dependencies/pkg" 
GO111MODULE=auto

docker: compile
	docker build -t vishalj82/kubewatch-tls-init  .
	docker push vishalj82/kubewatch-tls-init

compile:
	GOOS=linux GOARCH=386 go build -o bin/tlssetup-linux-386 tlssetup.go webhookConfig.go
