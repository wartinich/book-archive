run:
	go run cmd/app/main.go

build:
	docker-compose build book-archive && docker-compose up book-archive

down:
	docker-compose down book-archive