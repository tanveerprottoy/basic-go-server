# run in container with everything included
docker-compose up

# run service locally postgres in docker
postgresql and book_service_db in postgres must be present for the service to run
``` start postgres if not available in the system through docker
docker compose -f postgres.yml up

```cli
go run cmd/bookservice/main.go
```
```makefile
make run
```
# build executable for linux 
GOOS=linux GOARCH=amd64 go build -o bin/app cmd/basicserver/main.go