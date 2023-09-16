default:
	docker build -t scim-client .
	docker run -v `pwd`/config:/config scim-client

run:
	go run ./cmd/scim-client-go

test:
	go test ./...


docker-build:
	docker build -t scim-client .

docker-run:
	docker run -v `pwd`/config:/config scim-client
