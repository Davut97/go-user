# go-user

go-user is the stare of the art when it comes to creating new user and getting jokes from the internet, however the app is still under active development :)

## How to Run the app

First make sure you have docker and docker-compose installed on your machine
To run the tests simply run:

    make test

To run the server run:

    make run

## Endpoints

    Post /user { "email":"fo@fgo.com", "password":"214112412523", "firstName":"Lucky","lastName":"McLucky"}
    Post /login {"email":"fo@fgo.com", "password":"214112412523" }
    Get  /jokes
