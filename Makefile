.PHONY: build test cover
all: build test 
	@./pbkdf2-go  
	
build:
	@go build .

test:
	@go test --run 100 --bench=. -coverprofile=coverage.out 

cover: test
	@go tool cover -html=coverage.out