# Payment Server
[![Go Report Card](https://goreportcard.com/badge/github.com/reactivex/rxgo)](https://goreportcard.com/report/github.com/teivah/payment-server)

## Getting Started

`payment-server` is a Go example application exposing a RESTful API to manage payments using Mongo 3 for the persistence.

The API is based on [Swagger](/api/swagger.yml). It is a subpart of the [Form3 Swagger](https://github.com/form3tech-oss/go-form3/blob/master/swagger.yaml) with minor variations.

Here is an example of a payment request payload which can be handled by `payment-server`: [payment.json](/test/data/payment.json). 

## Project Structure

The project structure is based on https://github.com/golang-standards/project-layout.

* [/api](api): Swagger file.
* [/docs](docs): Additional documentation (design).
* [/githooks](githooks): Git hooks to be used by project maintainers.
* [/main](main): Main package.
* [/payment](payment): Payment package.
* [/scripts](scripts): Utility bash scripts.
* [/swagger](swagger): Generated go files from Swagger.
* [/test](test): Test resources.
* [/utils](utils): Utilities package (logger, config).

## Design

Please refer to the design document ([markdown](/docs/design/README.md) or [pdf](/docs/design/design.pdf) version).

## Running the Application

### Prerequisites

Before testing, please run [start_local_env.sh](scripts/start_local_env.sh) that will execute a Docker Compose with:
* `mongo:3.4.20-jessie`
    * Ports: 27017 the Mongo HTTP port.
* `mongo-express:0.49`
    * Ports: 8081 to navigate to the Mongo database.

As a first test, you can run the following `POST` command:

```
curl --header "Content-Type: application/json" --request POST --data '{"data":{"type":"payments","version":0,"organisation_id":"ee2fb143-6dfe-4787-b183-ca8ddd4164d2","attributes":{"amount":"26.00","batch_booking_indicator":"string","batch_id":"string","batch_type":"string","beneficiary_party":{"account_name":"James Bond","account_number":"71268996","account_number_code":"IBAN","account_type":0,"account_with":{"bank_address":["Liverpool Customer Service Centre","Stevenson Way","Wavertree","L13 1NW"],"bank_id":"333333","bank_id_code":"SWBIC","bank_name":"NATIONAL WESTMINSTER BANK PLC","bank_party_id":"//AT12345"},"address":["1 Clarence Mew","Horsforth","Leeds Ls18 4EP"],"birth_date":"1977-02-28","birth_city":"PARIS","birth_country":"FR","birth_province":"NORTHSIDE","country":"DE","name":"James Bond","organisation_identification":"ID1234656","organisation_identification_code":"BIC","organisation_identification_issuer":"BANK","telephone_number":"+447921123987"},"category_purpose_coded":null,"category_purpose":null,"charges_information":{"bearer_code":"string","receiver_charges_amount":"string","receiver_charges_currency":"string","sender_charges":[{"amount":"string","currency":"string"}]},"clearing_id":null,"currency":"EUR","debtor_party":{"account_name":"Jane Bond","account_number":"12345678","account_number_code":"IBAN","account_with":{"bank_address":["Liverpool Customer Service Centre","Stevenson Way","Wavertree","L13 1NW"],"bank_id":"333333","bank_id_code":"SWBIC","bank_name":"NATIONAL WESTMINSTER BANK PLC","bank_party_id":"//AT12345"},"address":["63 St Mary Axe","London","EC3A 8AA"],"birth_date":"1973-01-31","birth_city":"PARIS","birth_country":"FR","birth_province":"SOUTH SIDE","country":"GB","customer_id":"BARCGB22","customer_id_code":"SWBIC","name":"Norman Smith","organisation_identification":"ID1234656","organisation_identification_code":"BIC","organisation_identification_issuer":"BANK"},"end_to_end_reference":"PAYMENT REF: 20094","file_number":"string","fx":{"contract_reference":"FXCONTRACT/REF/123567","exchange_rate":"0.13343","original_amount":"100.00","original_currency":"EUR"},"instruction_id":"ID1245799","intermediary_bank":{"bank_address":["Liverpool Customer Service Centre","Stevenson Way","Wavertree","L13 1NW"],"bank_id":"333333","bank_id_code":"SWBIC","bank_name":"XYZ BANK PLC","bank_party_id":"//AT12345"},"numeric_reference":null,"payment_acceptance_datetime":"2017-09-30T12:36:02.123+01:00","scheme_transaction_id":"123456789012345678","unique_scheme_id":"L5W48NDWYW7JV9MRO71020180301826040011","payment_purpose":"X","payment_purpose_coded":"string","payment_scheme":"FPS","payment_type":"string","processing_date":"2015-02-12","scheme_processing_date":"2015-02-12","receivers_correspondent":{"bank_address":["Liverpool Customer Service Centre","Stevenson Way","Wavertree","L13 1NW"],"bank_id":"333333","bank_id_code":"SWBIC","bank_name":"XYZ BANK PLC","bank_party_id":"//AT12345"},"reference":"rent for oct","regulatory_reporting":"May be required for some foreign originated payments","reimbursement":{"bank_address":["Liverpool Customer Service Centre","Stevenson Way","Wavertree","L13 1NW"],"bank_id":"333333","bank_id_code":"SWBIC","bank_name":"NATIONAL WESTMINSTER BANK PLC","bank_party_id":"//AT12345"},"remittance_information":"Additional remittance information over and above reference information","scheme_payment_sub_type":"TelephoneBanking","scheme_payment_type":"ImmediatePayment","senders_correspondent":{"bank_address":["Liverpool Customer Service Centre","Stevenson Way","Wavertree","L13 1NW"],"bank_id":"333333","bank_id_code":"SWBIC","bank_name":"XYZ BANK PLC","bank_party_id":"//AT12345"},"structured_reference":{"issuer":null,"reference":null},"swift":{"bank_operation_code":"CRED","header":{"destination":"MIDLGB22XABC","message_type":"MT103","priority":"Priority","recipient":null,"source":null,"user_reference":null},"instruction_code":"INTC","sender_receiver_information":"/INS/ABNANL2A","time_indication":"/CLSTIME/0915+0200"},"ultimate_beneficiary":{"name":"string","country":"string","address":"string","organisation_identification":"string","organisation_identification_code":"string","organisation_identification_issuer":"string","birth_date":"2019-03-21","birth_city":"string","birth_country":"string","birth_province":"string"},"ultimate_debtor":{"name":"string","country":"string","address":"string","organisation_identification":"string","organisation_identification_code":"string","organisation_identification_issuer":"string","birth_date":"2019-03-21","birth_city":"string","birth_country":"string","birth_province":"string"}},"relationships":{"payment_submission":{"data":[{"type":"string","id":"3fa85f64-5717-4562-b3fc-2c963f66afa6","version":0,"organisation_id":"3fa85f64-5717-4562-b3fc-2c963f66afa6","attributes":{"status":"accepted","scheme_status_code":"string","status_reason":"string","settlement_date":"2019-03-21","settlement_cycle":0,"redirected_bank_id":"string","redirected_account_number":"string"},"relationships":{"payment":{"data":[{"type":"string","id":"3fa85f64-5717-4562-b3fc-2c963f66afa6"}]},"validations":{"data":[{"type":"string","id":"3fa85f64-5717-4562-b3fc-2c963f66afa6"}]}}}]},"payment_return":{"data":[{"type":"string","id":"3fa85f64-5717-4562-b3fc-2c963f66afa6","version":0,"organisation_id":"3fa85f64-5717-4562-b3fc-2c963f66afa6","attributes":{"amount":"string","currency":"string","return_code":"string","scheme_transaction_id":"string"},"relationships":{"payment":{"data":[{"type":"string","id":"3fa85f64-5717-4562-b3fc-2c963f66afa6"}]},"return_admission":{"data":[{"type":"string","id":"3fa85f64-5717-4562-b3fc-2c963f66afa6"}]},"return_submission":{"data":[{"type":"string","id":"3fa85f64-5717-4562-b3fc-2c963f66afa6"}]}}}]},"payment_admission":{"data":[{"type":"string","id":"3fa85f64-5717-4562-b3fc-2c963f66afa6","version":0,"organisation_id":"3fa85f64-5717-4562-b3fc-2c963f66afa6","attributes":{"status":"confirmed","scheme_status_code":"string","status_reason":"accepted","settlement_date":"2019-03-21","settlement_cycle":0},"relationships":{"payment":{"data":[{"type":"string","id":"3fa85f64-5717-4562-b3fc-2c963f66afa6"}]}}}]},"payment_reversal":{"data":[{"type":"string","id":"3fa85f64-5717-4562-b3fc-2c963f66afa6","version":0,"organisation_id":"3fa85f64-5717-4562-b3fc-2c963f66afa6","attributes":{},"relationships":{"payment":{"data":[{"type":"string","id":"3fa85f64-5717-4562-b3fc-2c963f66afa6"}]},"reversal_admission":{"data":[{"type":"string","id":"3fa85f64-5717-4562-b3fc-2c963f66afa6"}]}}}]}}}}' http://localhost:8080/v1/payment
```

To make sure a payment has been created, you navigate in Mongo Express: [http://dockerhost:8081/db/payment/payment/](http://dockerhost:8081/db/payment/payment/) (please make sure to change `dockerhost`).

### API Testing

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