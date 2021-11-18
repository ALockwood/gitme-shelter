build: 
	@go build

run:
	@go run .

demo:
	@GITHUB_USERNAME=x GITHUB_ACCESS_TOKEN=a AWS_ACCESS_KEY_ID=b AWS_SECRET_ACCESS_KEY=c LOG_LEVEL=trace go run . -cfg default-test.yaml

docker-build:
	@docker build . -t gitmeshelter
