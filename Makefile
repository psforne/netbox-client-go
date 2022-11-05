deps:
	go mod download
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest

generate: deps
	~/go/bin/swagger generate client --target=./netbox

clean:
	rm -rf netbox/client netbox/models

integration:
	go test ./... -tags=integration
