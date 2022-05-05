# golang-with-mongo
## Prerequisites

* docker

## Installation

### a. docker
pull mongo image https://hub.docker.com/_/mongo
```sh
docker pull mongo
```
check if pull success
```sh
docker images  # should see mongo
```
config the ``docker-compose.yaml``
```yaml
version: '3.9'

services:
  mongodb:
    image: mongo  # specify image name
    ports:
      - 27017:27017  # open mongo default port 27017
    volumes:
      - ./data:/data/db  # bind the folder from container /data/db -> current ./data
    environment:
      - MONGO_INITDB_ROOT_USERNAME=root  # sample name to login
      - MONGO_INITDB_ROOT_PASSWORD=123   # sample pwd to login
```
run the container with docker compose
```sh
docker-compose up
```
>ctrl+c to close

check if container exists
```sh
docker ps
```
### b. clone the repo
```sh
git clone https://github.com/vincent0426/golang-with-mongo.git
```
### c. golang
```sh
cd golang-with-mongo
go mod init golang-with-mongo
go run main.go
```
