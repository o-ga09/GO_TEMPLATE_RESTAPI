run:
	ENV="dev" DATABASE_URL="api:P@ssw0rd@tcp(localhost:3306)/api?charset=utf8&parseTime=True&loc=Local" PROJECTID="0000000" go run ./api/cmd/main.go
test:
	go test --race ./...