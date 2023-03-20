# Articles

## Prerequisites
* Docker should be up and running

## Database
* Mongodb

## Assumptions
* Mongodb should not be running in the system as the default port has been exposed from docker. If the mongodb is running in the system on default port, docker compose will not bring mongo container up
* This was developed and test on MacOS(Apple silicon)

## Configuration
You can configure application port & mongo credentials using env variables
* PORT: Use this env varibale to run the application on desired port. Default : `8080`
* MONGOURI: This holds the mongodb URI which is used to connect to db. Sample: `mongodb://[username]:[password]@[host]:[port]/myarticles?retryWrites=true&w=majority`

## Running the application
* To run the application, go to the root folder and run `start.sh`
* You can access the endpoint by going to below url:

    `http://localhost:8080/[endpoint]`



## Unit tests
 You run the unit test using the following cmd

 `go test ./...`

## Integration tests
 You run the unit test using the following cmd

 `go test ./... --tags=integration`