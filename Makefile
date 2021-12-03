build: 
	@go build

run:
	@go run .

demo:
	@LOG_LEVEL=trace go run . -cfg default-test.yaml

demo2:
	@LOG_LEVEL=debug go run . -cfg default-test.yaml

docker-build:
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/gitme-shelter .
	@docker build . -t gitmeshelter
