build: 
	@go build

run:
	@go run .

dry:
	@go run . -cfg "x" --dryrun

docker-build:
	@docker build . -t gitmeshelter
