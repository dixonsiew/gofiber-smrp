go mod tidy
swag init
go build -ldflags "-s -w" -buildvcs=false