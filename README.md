## Run Postgres Example

Steps to run the example are :

1) Run the docker image

   > `docker run --name gofr-pgsql -e POSTGRES_DB=cats -e POSTGRES_PASSWORD=password -p 2006:5432 -d postgres:latest
   `
2) Now on the project path `zopsmart/gofr` run the following command to load the schema

   > `docker exec -i gofr-pgsql psql -U postgres cats < .github/setups/setup.sql`

3) Run server

   > `go run main.go`
