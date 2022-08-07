test:
	go test -v ./... -cover

build: test
	go build -v -o ./bin/ ./cmd/consulta-cep/consulta-cep.go 
	go build -v -o ./bin/ ./cmd/consulta-correios/consulta-correios.go 
