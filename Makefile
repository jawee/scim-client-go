default:
	go run ./cmd/scim-client-go

run:
	go run ./cmd/scim-client-go

test:
	go test ./...


docker-build:
	docker build -t scim-client .

docker-run:
	docker run scim-client
