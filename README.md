# Whalebone Assignment

A simple microservice which provides two endpoints.

# Requirements
- Go v1.23
- Docker

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

