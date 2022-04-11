check_install:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	swagger generate spec -o ./swagger.yaml --scan-models

serve_swagger:
	swagger serve -F=swagger swagger.yaml

run:
	go run ./main.go

build:
	go build -o bin/main main.go
