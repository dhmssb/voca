postgres:
	docker run --rm --name dockPost -e POSTGRES_PASSWORD=123456 -e POSTGRES_USER=postgres -p 5432:5432 -d postgres:latest  

createdb:
	docker exec -it dockPost createdb --username=postgres voca

migrateup:
	migrate -path database/migrations -database "postgresql://postgres:123456@localhost:5432/voca?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migrations -database "postgresql://postgres:123456@localhost:5432/voca?sslmode=disable" -verbose down



.PHONY: postgres createdb migrateup migratedown