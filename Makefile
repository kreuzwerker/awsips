build:
	mkdir -p out
	go build -o out/awsips -ldflags "-X main.build `git rev-parse --short HEAD`" bin/awsips.go

test:
	go test -cover
