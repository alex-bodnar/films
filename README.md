# Films api

Films api is a rest-api server for searching films.
App using [Redis](https://redis.io/) for caching request.

## Usage
To run the application on your computer, you will need:
- installed go 1.18
- installed redis store
- change in config ```./volume/config.yaml``` redis parameters
- run ```make mod``` for getting dependencies
- run ```make run``` for start app
- test app in postman with collection ```./tests/postman/Films-api.postman_collection.json```

To run the application on [docker](https://www.docker.com/), you will need:
- installed docker
- run ```make start-docker-compose```
- test app using postman with collection ```./tests/postman/Films-api.postman_collection.json```

If your wont to generate binary file and run in, yur will need:
- installed go 1.18
- installed redis store
- change in config ```./volume/config.yaml``` redis parameters
- run ```make run-build``` for build and start app

## Swagger
If yor wont to see swagger documentation, yor will need:
- installed [redoc-cli](https://redocly.com/docs/redoc/deployment/cli/)
- run ```make swagger-serve```
- open http://127.0.0.1:8080 on your browser

## Test
For running service test, yor will need:
- run ```make run``` for start app
- in other terminal run ```make test-service``` for start test
