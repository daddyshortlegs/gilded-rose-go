# Gilded Rose in Go

I wasn't going to learn Go, but an interesting distraction
nonetheless. Different to work with test library with no 
assertions. 


- Run :

```shell
go run texttest_fixture.go gilded-rose.go
```

- Run tests :

```shell
go test
```

- Run tests and coverage :

```shell
go test -coverprofile=coverage.out

go tool cover -html=coverage.out
```