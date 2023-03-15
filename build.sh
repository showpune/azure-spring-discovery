GOOS=linux 
echo $GOOS
GOARCH=amd64
echo $GOARCH
go build -o discovery cli/main.go