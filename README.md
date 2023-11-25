# go-user

go-user is the stare of the art when it comes to creating new user and getting jokes from the internet, however the app is still under active development :)

#### Create an application that connects to a MongoDB database and performs CRUD (Create, Read, Update, and Delete) operations on a collection.

Done

#### Implement a concurrency mechanism that allows the application to handle a large number of concurrent requests.

Used concurrency in the implementation of get Jokes endpoint (Jokes.go)

#### Implement a custom business logic that involves data encryption and decryption for sensitive fields in the collection.

Done (Hashing the password)

#### Use an Object-Document Mapping (ODM) framework (e.g. Mongoose) to handle the database operations.

Golang does not support ODM

#### Create a REST API for the application.

Done

#### Add unit tests to test the main parts of the application.

Done

#### Use Git for version control and GitHub for code sharing.

Done

#### Create a detailed document that explains the architecture of the application, the technologies used, and the reasoning behind the design decisions made, especially the encryption and decryption mechanism.

Used technologies: golang, docker.

#### Create a Dockerfile to containerize the application, and a docker-compose.yml file that sets up the application and the MongoDB database.

Done

#### Share your code repository with commitsmart-jobs github account.

Done

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
