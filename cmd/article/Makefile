# DATABASE
DB_USER=postgres
DB_PASSWORD=bismillah
DB_HOST=127.0.0.1
DB_PORT=5432
DB_NAME=sagara_crud_article
DB_SSL=disable


gen-proto:
	mkdir -p cmd/auth/proto
	protoc --proto_path=cmd/auth/proto --go_out=cmd/auth/proto \
		--go_opt=paths=source_relative cmd/auth/proto/*.proto;
	protoc --proto_path=cmd/auth/proto --go-grpc_out=cmd/auth/proto \
		--go-grpc_opt=paths=source_relative cmd/auth/proto/*.proto;

migrate-up:
	migrate -source file:./migrations/ -verbose -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL} up 1

migrate-down:
	migrate -source file:./migrations/ -verbose -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL} down 1

migrate-lastest:
	migrate -source file:./migrations/ -verbose -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL} up

migrate-drop:
	migrate -source file:./migrations/ -verbose -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL} drop

migration-new-%:
	migrate create -ext sql -dir ./cmd/auth/migrations -seq $*