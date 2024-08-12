CURRENT_DIR=$(shell pwd)
DBURL := postgres://macbookpro:1111@localhost:5432/auth?sslmode=disable

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}


swag:
	~/go/bin/swag init -g ./api/router.go -o api/docs
run-service:
	go run cmd/main.go
