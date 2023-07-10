build:
	npm run --prefix ./web/namegpt-ui build
	go build -o ./bin/namegpt ./cmd/namegpt/main.go
deploy:
	./bin/namegpt
