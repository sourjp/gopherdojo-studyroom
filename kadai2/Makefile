build:
	@go build -o converter cmd/conv/main.go

test:
	@go test -v ./ -cover
	@go test -v ./cmd/conv/ -cover

test-cover-out:
	@go test -v -coverprofile=c.out
	@go tool cover -html=c.out

