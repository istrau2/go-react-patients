# FullStack Task Go Back End

## Getting Started
In order to operate on the server application, you must navigate to the `server` directory before running any of the commands below.

Be sure you have a modern version of [Go](https://golang.org/) intalled.

## Available Scripts

Before you can run the server, you need to seed the db. By default it will run with a file-backed sqlite db. To create the db file and seed, run:

### `go run seed.go`

Which will run and seed the db with 200 patient records. You can keep running it, and it will seed the db with an additional 200 records each time.

### `go run main.go`

Runs the server in development mode.<br />
Open [http://localhost:8080/v1/patient](http://localhost:8080v1/patient) to see the only implemented endpoint.