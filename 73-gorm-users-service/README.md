## go get 

go get -u gorm.io/gorm
go get -u github.com/gin-gonic/gin

or 

go mod tidy


## Docker network

```docker network create demo-network
```

## docker postgres

```
docker run -d --name pg -p 5432:5432 --network demo-network -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=usersdb postgres:16

docker run -d --name dbui -p 28080:8080 --network demo-network adminer

docker ps
```