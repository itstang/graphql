clean:
	$(RM) schema/_bindata.go

dep:
	dep ensure

dev: schema
	go run cmd/graphql/main.go

schema:
	go generate ./schema

setup:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/jteeuwen/go-bindata/...
