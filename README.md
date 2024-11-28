# Whalebone Assignment

A simple microservice which provides two endpoints.

# Requirements
- Go v1.23
- Docker
- Make (optional)

# Setup
from project root copy `.env-example` into `.env`
```bash
copy .env-example .env
```

run one of the `make` commands (see the file for all option)
```bash
make run
```

```bash
make itest
```

or you can run it manually
```bash
go run cmd/api/main.go
go build -o main cmd/api/main.go
```

# Run dockerized application
There is a docker file prepared in the project root.

```bash


```

# Examples

post request to save user:
```bash
curl -X POST http://localhost:8080/save -d'{
"id": "ae593b85-b9a2-4386-ad71-7b62287d7c24",
"name": "example name",
"email": "example@gmail.com",
"date_of_birth": "2020-01-01T12:12:34+00:00" }'
```

get request for user created above:
```bash
curl -X GET http://localhost:8080/ae593b85-b9a2-4386-ad71-7b62287d7c24
```

