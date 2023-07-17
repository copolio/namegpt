ifndef ($(GOPATH))
	GOPATH = $(HOME)/go
endif

generate-swagger:
	$(GOPATH)/bin/swag init --parseDependency --parseInternal -g ./cmd/namegpt/main.go -o ./api/swagger
build-frontend:
	npm install --prefix ./web/namegpt-ui
	npm run --prefix ./web/namegpt-ui build
build-backend:
	go build -o ./bin/namegpt ./cmd/namegpt/main.go
deploy:
	./bin/namegpt
