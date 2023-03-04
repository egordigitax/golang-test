build:
	go build ./cmd/app
run:
	go run ./cmd/app
docker:
	sudo docker-compose up --build -d
