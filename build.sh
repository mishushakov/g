GOOS=darwin go build -o ./bin/mac/g -ldflags="-s -w" ./main.go 
GOARCH=amd64 GOOS=linux go build -o ./bin/linux/g -ldflags="-s -w" ./main.go
GOOS=windows go build -o ./bin/windows/g -ldflags="-s -w" ./main.go