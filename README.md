# NameGPT

NameGPT is a simple microservice application you can use for domain name recommendation.

This application has been built using following technologies:
- Golang (gin-gonic)
- GORM, MySQL (database)
- ChatGPT API

## How to use

```shell
# Creates Swagger document
swag init --parseDependency --parseInternal -g ./cmd/namegpt/main.go -o ./api/swagger
# Build service 
go build github.com/copolio/namegpt/cmd/namegpt
```
