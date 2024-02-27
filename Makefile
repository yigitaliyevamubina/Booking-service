CURRENT_DIR := $(shell pwd)
DB_URL := postgres://bobo:1234@localhost:5432/bookingdb?sslmode=disable

proto-gen:
	chmod +x ./scripts/genproto.sh
	./scripts/genproto.sh

migrate-up:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" -verbose down

migrate-force:
	migrate -path migrations -database "$(DB_URL)" -verbose force 1

migrate-file:
	migrate create -ext sql -dir migrations/ -seq create_comments_table
