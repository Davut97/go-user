.PHONY:  test run

run:
	docker-compose build && docker-compose up

test:
	docker run -d --rm --name mongodb-test -p 27017:27017 mongo
	go test -v ./...
	docker stop mongodb-test