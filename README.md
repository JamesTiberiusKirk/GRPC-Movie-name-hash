# Movie name hasher

- To run with docker compose:
```sh
docker-compose up
```

- To test the api:
```sh
curl -X 'POST' localhost:3001/hash-movie-name -H 'Content-Type: application/json' -d '{"name":"Test movie"}' 
```

- To run unit tests:
```sh
make test
```


- To run locally without docker:
```sh
export $(grep -v '^#' .env | xargs) # To load .env
make generate
go run ./cryptoservice
go run ./apiservice
```
