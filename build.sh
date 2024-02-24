#!/bin/bash
go build -o ./bin/fcompress ./cmd/fcompress/main.go && cp ./bin/fcompress ~/.local/bin/
go build -o ./bin/jsoncheck ./cmd/jsonchk/main.go
go build -o ./bin/ytcut ./cmd/ytcut/main.go && cp ./bin/ytcut ~/.local/bin
