test:
	go test -v ./... -cover

build: test
	go build -v -o ./bin/ ./cmd/consulta-cep.go 