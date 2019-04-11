# Simple GraphQL Server with Gin
## Development
Make sure [Go](https://ahmadawais.com/install-go-lang-on-macos-with-homebrew/) is installed and setup correctly

In your Go workspace (`$GOPATH/github.com/itstang/`):
```bash
$ git clone git@github.com:itstang/graphql.git
```

```bash
$ make setup  # downloads dependencies
$ make schema # generates schema code
$ make dev    # start local server
```

Navigate to `http://localhost:8080/` to play around in the GraphQL Explorer
