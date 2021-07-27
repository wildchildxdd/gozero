#!/bin/zsh
rm -rf greet
goctl api go -api greet.api -dir greet
cd greet
go get gorm.io/gorm 
go get gorm.io/driver/mysql
go mod init && go mod tidy && go run greet.go -f etc/greet-api.yaml
