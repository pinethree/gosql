postgresinit:
	docker run --name postgresalpine -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d postgres:alpine

postgres:
	docker exec -it postgresalpine psql -U root

createdb:
	docker exec -it postgresalpine createdb -U root -O root gosql

dropdb:
	docker exec -it postgresalpine dropdb -U root gosql

.PHONY: postgresinit postgres createdb dropdb
