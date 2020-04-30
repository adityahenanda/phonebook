# phonebook

clean architecture:

1. Models : layer for register model struct / request / response / entities
2. Repository : layer for database handler / query
3. Usecase : layer for bussiness logic
4. Delivery : output hanlder layer, http rquest / grpc (not implement yet)


quick set up (clone to your gopath )
1.edit config.json (set your mysql db connection), automatic migrate table
2.go get
3.go run main.go

