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

# endpoints
post -> http://localhost:8080/api/v1/books
        
        body: {
            "title": "a book2",
            "author": "c",
            "publicationYear": 2060
        }

get -> http://localhost:8080/api/v1/books?page=1&limit=20

get -> http://localhost:8080/api/v1/books/b7164383-5b73-4ae8-82ee-0e3ce8df5c8b

put -> http://localhost:8080/api/v1/books/b7164383-5b73-4ae8-82ee-0e3ce8df5c8b

        body: {
            "title": "a book2",
            "author": "an author3",
            "publicationYear": 3030
        }

delete -> http://localhost:8080/api/v1/books/b7164383-5b73-4ae8-82ee-0e3ce8df5c8b