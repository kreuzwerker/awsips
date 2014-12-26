TOKEN = `cat .token`
REPO := awsips
USER := kreuzwerker
VERSION := "v0.0.2"

build:
	mkdir -p out/darwin out/linux
	GOOS=darwin go build -o out/darwin/awsips -ldflags "-X main.build `git rev-parse --short HEAD`" bin/awsips.go
	GOOS=linux go build -o out/linux/awsips -ldflags "-X main.build `git rev-parse --short HEAD`" bin/awsips.go

clean:
	rm -rf out

release: clean build
	github-release release --user $(USER) --repo $(REPO) --tag $(VERSION) -s $(TOKEN)
	github-release upload --user $(USER) --repo $(REPO) --tag $(VERSION) -s $(TOKEN) --name awsips-osx --file out/darwin/awsips
	github-release upload --user $(USER) --repo $(REPO) --tag $(VERSION) -s $(TOKEN) --name awsips-linux --file out/linux/awsips

test:
	go test -cover
