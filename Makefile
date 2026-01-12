generate:
	protoc --go_out=. --go_opt=module=sports-betting-helper \
		--go-grpc_out=. --go-grpc_opt=module=sports-betting-helper \
		proto/*.proto

up:
	docker-compose up --build -d

down:
	docker-compose down

down-all:
	docker-compose down -v

.PHONY: generate up down