# HackZurich 2018 - Drugify Server
 
## Build with Docker

The Docker build is using a multi-stage approach which means that the final container contains the bare minimum to run application
````
docker build -t drugify-server .
````

## Run with Docker

Make sure to start a Mongo DB container like this (the name is important!)
````
docker run -d --name mongo -p 27017:27017 -v mongo_data:/etc/mongo mongo:4
````

After the Mongo DB is up and running start the `drugify-server` application

````
docker run -d --name drugify-server --link mongo -p 3000:3000 drugify-server:latest 
````

## Run with Docker Compose

Make sure you are within the projects root directory and you have the [docker-compose](https://docs.docker.com/compose/install/#prerequisites) tool installed

````
docker-compose up
````

## Run with Docker Swarm

Make sure you initialized your Servers via `docker swarm init` or joined an existing swarm

````
docker stack deploy -c docker-compose.yml drugify-server
````

**Important**: In Docker Swarm mode the `depends_on` feature will be ignored which means that the drugify-server is probably failing a few times until the `mongo` service is up and running.


## Development

#### Dependencies

You need to install the following dependencies:

````
go get -u github.com/BurntSushi/toml 
go get -u gopkg.in/mgo.v2
go get -u github.com/sirupsen/logrus
go get -u github.com/gin-gonic/gin
````