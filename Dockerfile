FROM golang:1.3-onbuild
ENV DATABASE_URL_DOCKER "dbname=postgres user=postgres password=password host=198.199.124.226 port=5432 sslmode=disable"
EXPOSE 8080

