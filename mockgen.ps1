go install github.com/golang/mock/mockgen@v1.6.0
mockgen -source="repository/repository.go" -destination="mocks/repository/mock.go"
mockgen -source="service/service.go" -destination="mocks/service/mock.go"
mockgen -source="handler/handler.go" -destination="mocks/handler/mock.go"
go get github.com/golang/mock/mockgen

go install golang.org/x/tools/...@latest
go test -coverprofile="coverage.out" ./...
go tool cover -html="coverage.out"

go mod tidy
exit 0