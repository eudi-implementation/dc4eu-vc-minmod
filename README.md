# DC4EU VC

This repo bundles the original DC4EU Verifiable Credential (VC) service at a fixed, known‑good version and spins it up via Docker containers. The gateway is exposed on port **9080**, so external issuers (e.g. the [EUDI Reference Issuer](https://github.com/eudi-implementation/eudi-srv-web-issuing-eudiw-py-modified)) can simply call the `collect_id` endpoint to pull stored documents.

Since the API‑Gateway handles all calls to the issuer, datastore, etc., we deploy the full DC4EU VC repo but only use the `document/collect_id` endpoint. A `collect_id` is a short‑lived reference to a document in the datastore; when queried with the matching identity, the corresponding document is returned.

We use the `collect_id` in our modified [EUDI Reference Issuer](https://github.com/eudi-implementation/eudi-srv-web-issuing-eudiw-py-modified), by including it in the credential offer and together with PID authentication, we can successfully query the DC4EU datastore.

## Quickstart

1. **Clone this repo**  
  ```bash
  git clone git@github.com:eudi-implementation/dc4eu-vc-minmod.git
  
  ```
2. **Start the service**
  ```bash
  cd dc4eu-vc-minmod
  docker-compose up -d
  ```

3. **Retrieve Document**

In your issuer, call `http://your.local.ip.address:9080/api/v1/document/collect_id` with a `collect_id` and the corresponding identity information. For more details, refer to the [API Specification](https://github.com/eudi-implementation/dc4eu-vc-minmod/blob/main/standards/api_specification.md) or the [Swagger File](https://github.com/eudi-implementation/dc4eu-vc-minmod/blob/main/docs/apigw/swagger.yaml).


----------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Find the origianl README below

----------------------------------------------------------------------------------------------------------------------------------------------------------------------------

This repository consists of the source code for the VC EU project, but also tools and make targets that's making sense for developers, please do not use for anything else.

Are you looking for running this, and need some sort of starting point, please take a look at <https://github.com/dc4eu/vc_up_and_running>

## docker release version

`latest` tracks the latest tag available and is build from branch `main`.

## branches

`main` is the stable development branch.

## How to build

### Docker

For convenience all services are build within a docker container.

Each service has its own make target, `make docker-build-<service>` or build all of them at once `make docker-build`

### Static binary without Docker

run `make build-<service>` to build a specific service or `make build` to build all services.

linux/amd64 is consider supported, other build options may work.

## Start, Stop & Restart

`make start` or `docker-compose -f docker-compose.yaml up -d --remove-orphans`

`make stop` or `docker-compose -f docker-compose.yaml rm -s -f`

`make restart`

## Swagger

### Endpoint

`GET http://<apigw-url>/swagger/doc.json`

or with web browser: `http://<apigw-url>/swagger/index.html`
