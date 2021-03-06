example:
	go build -o ./dist/bin/cosmosapi-examples ./examples/cosmosapi/main.go
	go build -o ./dist/bin/cosmos-examples ./examples/cosmos/main.go

test:
	go build cmd/cosmosdb-apply/main.go
	go test -v `go list ./cosmosapi`
	go test -tags=offline -v `go list ./cosmos`
	go test -v `go list ./cosmostest`

vet: exttools/bin/shadow
	go vet ./...
	go vet -vettool=exttools/bin/shadow ./...

exttools/bin/shadow: exttools

exttools:
	cd exttools && ./build.sh

.PHONY: example test vet exttools
