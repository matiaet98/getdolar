build.linux:
	GOOS=linux GOARCH=amd64 go build -o ./getdolar main.go
build.docker:
	docker build -t get-dolar:latest .