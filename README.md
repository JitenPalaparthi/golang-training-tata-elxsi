# Golang 

- Create a directory 
- cd to the directory 

- To create a go project (binary project/library project)

```bash
go mod init demo
```

- Creates a go.mod file. go mod is a package manager, kind of a build tool

- To run Go application

```bash
go run main.go
go run .
```

- To run multiple Go files

```bash
go run main.go greet.go
go run .
```

- To build

```bash
go build main.go
# build with name 
go build -o demo main.go
# release build
go build -ldflags="-w -s" -o release_demo main.go
```

