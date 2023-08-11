# Auction House Service - WIP

Auction House Service is an example web API I would envision what a basic auction house service would need for a game.

## Install

    go mod vendor
    go build

## Run the service

    go run .

## Run the tests

    go test

## REST API
The REST API for the Auction House service is described below.

## Get a Bid
### Request
```GET \bid\ ```

    curl -i -H 'Accept: application/json' http://localhost:8080/bid/

## Post a Bid
### Request
```POST \bid\ ```

## Get a listing
### Request
```GET \listing\ ```

## Post a listing
### Request
```POST \listing\ ```

## Patch a listing
### Request
```PATCH \listing\ ```

## Delete a listing
### Request
```DELETE \listing\ ```


