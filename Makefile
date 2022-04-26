gmsDir := ./cmd/gitme-shelter

.PHONY: build


test:
	@go test -v --cover ./...

build:
	@go build -o bin/ ${gmsDir}

run:
	@go run ${gmsDir}

demo:
	@LOG_LEVEL=trace go run ${gmsDir} -cfg ./assets/default-test.yaml

demo2:
	@LOG_LEVEL=debug go run ${gmsDir} -cfg ./assets/default-test.yaml

docker-build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/gitme-shelter ${gmsDir}
	docker build -f ./build/Dockerfile -t gitmeshelter .
