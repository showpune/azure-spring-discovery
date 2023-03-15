rm discovery-l
GOOS=linux 
GOARCH=amd64
echo "Building discovery with GOOS=$GOOS GOARCH=$GOARCH"
go build -mod=mod -o discovery-l cli/main.go

