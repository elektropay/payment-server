# Payment Server
[![Go Report Card](https://goreportcard.com/badge/github.com/reactivex/rxgo)](https://goreportcard.com/report/github.com/teivah/payment-server)

## Getting Started

`payment-server` is a Go example application exposing a RESTful API and managing the persistence using Mongo 3.

The API is based on [Swagger](/api/swagger.yml). It is a subpart of the [Form3 Swagger](https://github.com/form3tech-oss/go-form3/blob/master/swagger.yaml) with minor variations.

Here is an example of a payment request payload which can be handled by `payment-server`: [payment.json](/test/data/payment.json). 

## Project Structure

The project structure is based on https://github.com/golang-standards/project-layout.

* [/api](api): Swagger file.
* [/githooks](githooks): Git hooks to be used by project maintainers.
* [/main](main): Main package.
* [/payment](payment): Payment package.
* [/scripts](scripts): Utility bash scripts.
* [/swagger](swagger): Generated go files from Swagger.
* [/test](test): Test resources.
* [/utils](utils): Utilities package (logger, config).

## Design

Please refer to [DESIGN.md](design.md) or [DESIGN.pdf](DESIGN.pdf).

## API Testing

Before testing, please run [start_local_env.sh](scripts/start_local_env.sh) that will execute a Docker Compose with:
* `mongo:3.4.20-jessie`
    * Ports: 27017 the Mongo HTTP port.
* `mongo-express:0.49`
    * Ports: 8081 to navigate to the Mongo database.

In order to test the payment API, please refer to [payment_test.go](payment_test.go) file.
It contains BDD/Contract-based tests using `ginkgo` and `gomega`. 

## Application Properties

| Argument                      | Default                   | Description 
|---                            |---                        |---
| `logging.level`               | info                      | Logging level (`debug`, `info`, `warn`, `error` or `panic`).
| `server.port`                 | 8080                      | Server port.
| `server.external.hostname`    | 8080                      | External hostname used in HATEOAS responses.
| `server.external.port`        | localhost                 | External port used in HATEOAS responses.
| `mongo.uri`                   | mongodb://localhost:27017 | Mongo URI.
| `mongo.payment.db`            | payment                   | Mongo payment database name.
| `mongo.payment.collection`    | payment                   | Mongo payment collection name.
| `mongo.connection.timeout.ms` | 5000                      | Mongo connection timeout in milliseconds.
| `mongo.request.timeout.ms`    | 500                       | Mongo request timeout in milliseconds.