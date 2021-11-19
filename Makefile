build: 
	@go build

run:
	@go run .

demo:
	@LOG_LEVEL=trace go run . -cfg default-test.yaml

demo2:
	@LOG_LEVEL=debug go run . -cfg default-test.yaml

docker-build:
	@docker build . -t gitmeshelter
