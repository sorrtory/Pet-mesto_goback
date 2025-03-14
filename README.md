# Mesto project backend

Mesto API implemented with Golang

## Check on the result

> [mesto.sorrtory.ru](https://mesto.sorrtory.ru)

## Deploy
>
> Tested on Go 1.24.1 and Ubuntu 24.04
>
> With `docker compose`

1. Create .env

    - To start in docker

        ```bash
        POSTGRES_HOST=db
        POSTGRES_USER=postgres
        POSTGRES_PASSWORD=123

        BACKEND_HOST=0.0.0.0
        BACKEND_PORT=80
        BACKEND_PUBLIC="./public"

        ALLOW_MIGRATION=true
        MIGRATIONS_PATH=file://migrations
        ```

    - To start in system

        ```bash
        POSTGRES_HOST=localhost
        POSTGRES_USER=postgres
        POSTGRES_PASSWORD=123

        BACKEND_HOST=localhost
        BACKEND_PORT=8080
        BACKEND_PUBLIC="web/public"

        ALLOW_MIGRATION=true
        MIGRATIONS_PATH=file://cmd/migrate/migrations
        ```

2. Start
    - In docker\
      `make prod`

    - In system\
      `make up-db` and `make migrate-up` and `make`
3. Test with
    - Postman

      *OR*
    - [Visit](http://localhost:8080/static)

      ```bash
      http://${BACKEND_HOST}:${BACKEND_PORT}/static
      ```

## Conclusion

Anyway there's a container with Postgres and backend which is connected to it.

Frontend fetches to the backend and so a client could use the api to interact with db.

*SCREENSHOTS FROM FRONT...*

## Sources

- [Complete Backend API in Golang (JWT, MySQL & Tests)](https://youtu.be/7VLmLOiQ3ck?si=xTfy9YvVPjcD8sLc)
  - [Github](https://github.com/sikozonpc/ecom.git)
- [Go API Tutorial - Make An API With Go](https://youtu.be/bj77B59nkTQ?si=XQ9tkz9qa21LHFOd)
  - [Github](https://github.com/techwithtim/Go-API-Tutorial.git)
  - [Stolen from doc](https://go.dev/doc/tutorial/web-service-gin)
- [How To Build A Complete JSON API In Golang](https://youtu.be/pwZuNmAzaH8?si=fBmc8d3Bffjgt7UT)
- [Dockerize Go&Gin+Postgres](https://ramadhansalmanalfarisi8.medium.com/how-to-dockerize-your-api-with-go-postgresql-gin-docker-9a2b16548520)
  - [Github](https://github.com/ramadhanalfarisi/go-simple-dockerizing.git)
- [Docker Go&Mux+Postgres](https://dev.to/francescoxx/build-a-crud-rest-api-in-go-using-mux-postgres-docker-and-docker-compose-2a75)
- [Go API](https://habr.com/ru/companies/otus/articles/667308/)
- [Fiber+postgres](https://blog.logrocket.com/building-simple-app-go-postgresql/)
