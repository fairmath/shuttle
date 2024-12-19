# SHUTTLE
SHUTTLE is proxy server to translate Ethereum JSON-RPC requests to CosmosSDK based chains. Was developed to use [Blockscout](https://github.com/blockscout/blockscout) block explorer for CosmosSDK-based chains.

## Description
Shuttle works as a proxy on websocket connection and HTTP requests. Incoming JSON RPC requests is translated to corresponding Cosmos HTTP requests and then a result is transmitted in Ethereum format. Websocket connection transmits a new block info from cosmos chain to the client.
The are several settings for the proxy that should be installed thorugh environments variables:
* `HTTP_TENDERMINT_URL` - address of cosmos node 
* `WS_TENDERMINT_URL` - address of cosmos node supported a websocket connection
* `HTTP_LISTEN_ADDR` - listening address
* `LOG_LEVEL` - log level 
* `BECH_PREFIX` - prefix of addresses in the cosmos-based chain
* `COIN_DENOM` - coin denomination in the cosmos-based chain

> [!IMPORTANT] 
> This project is under development right now. Some requests could be unavailable and some fields in responses could be inaccurate. 

## Tools
* Go 1.23.3
* golangci-lint 1.55.2
* [Swagger](https://api.github.com/repos/go-swagger/go-swagger/releases/latest)
* GNU Make

## How To Build 
* Install necessary tools, for build will be enought the following:
  * Go 1.23.3
  * GNU Make
* Run the following command
```shell
$ make build
```
Resulted binary file will be located in `bin` directory
Instead of installation all nececcary tools locally it is possible to use prepared development docker container:
* Install docker and docker-compose
* Run the following command
```shell
$ docker-compose run dev
```
* Once container is run print the following command
```shell
$ make build
```

## How to Run
The simplest way to run shuttle - use docker compose:
* Install docker and docker-compose
* Run:
```shell
$ docker-compose up proxy
```
