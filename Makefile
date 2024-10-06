# Container
fiap-run:
	# docker rmi fiap-challenge-ddd-app
	docker-compose up -d

fiap-stop:
	docker-compose down

fiap-logs:
	docker-compose logs -f

# Local
build:
	go build -o bin/kitchencontrol cmd/kitchencontrol/main.go
	
test:
	go test -v -cover ./...

tidy:
	go mod tidy

run:
	go run cmd/kitchencontrol/main.go

install-swag:
	go install github.com/swaggo/swag/cmd/swag@latest
	echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
	source ~/.bashrc
	swag --version

swaggo:
	go get -u github.com/swaggo/swag/cmd/swag
	go get -u github.com/swaggo/gin-swagger
	go get -u github.com/swaggo/files
	go get -u github.com/swaggo/gin-swagger/swaggerFiles
	swag init -g cmd/kitchencontrol/main.go

docs:
	#rm -rf docs
	swag init -g ./cmd/kitchencontrol/main.go -o ./docs

test-coverage:
	go test -coverprofile=coverage.out ./...

coverage: test-coverage
	go tool cover -func=coverage.out

coverage-html: test-coverage
	go tool cover -html=coverage.out


delete-identifiers:
	find . -name "*.Identifier" -type f -delete
