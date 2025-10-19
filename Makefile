.PHONY: run test start-db start-test-db stop-db stop-test-db

swagger:
	swag init -g cmd/app/main.go -o cmd/docs
## 🧪 Run the development environment
run: swagger
	air

## 🧪 Run tests with test database
test:
	go test ./...

## Start Dev Services
start-db:
	docker-compose -p dev -f docker-compose.local-db.yml up -d 

## Start Test Services
start-test-db:
	docker-compose -p test -f docker-compose.test-db.yml up -d 

## 🛑 Stop dev services
stop-db:
	docker-compose -p dev -f docker-compose.local-db.yml down
	
## 🛑 Stop test services
stop-test-db:
	docker-compose -p test -f docker-compose.test-db.yml down
