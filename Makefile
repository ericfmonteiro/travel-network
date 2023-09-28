start-db:
	docker run --name postgres-db -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres

start-app:
	go run main.go

start: start-db start-app