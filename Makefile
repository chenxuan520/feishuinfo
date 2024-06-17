all: feishuinfo

feishuinfo: ./cmd/main.go
	go build -o feishuinfo ./cmd/main.go
