# Predix Enterprise Connect Configurator

A simple app to guide you through the Enterprise Connect configuration on Predix.io

## What is Enterprise Connect (EC)?

Provide manageable and scalable connectivity between the Predix cloud and your enterprise. Enterprise Connect uses Websockets to securely support all TCP protocols for scalable and flexible tunneling deployment.

The service also provides integrated security with UAA and flexible deployment options to connect to multiple data sources. Because there is no limit to the number of concurrent connections, the service is more versatile and scalable than most other VPN offerings.

--

<cite>
  <a href="https://www.predix.io/services/service.html?id=2184">Predix.io Documentation</a>
</cite>

--------------------------------------------------------------------------------

## Why I need Enterprise Connect and how it works?

Enterprise Connect addresses two scenarios:

1. Connect on-premise data source to Predix cloud;

  - Make accessible in a secure way an on-premise resource directly from the cloud

2. Connect Predix data sources to on-premise systems.

  - Make accessible in a secure way a Predix data source like PostgreSQL or RabibtMQ (running on an internal Predix network) out of Predix cloud, so that you can push data to the cloud without any additional code to be written (e.g. ETL, EAI or custom scripts).

Enterprise Connect (EC) consists of two main components:

- **EC Service**: automatically instanciated on Predix.io cloud subscribing to the service through the Predix.io catalog or CF CLI.
- **EC Agent**: securely engages heavy data flow at the TCP packet level, targeting lightweight adoption, delivers uncompromised data quality, and makes itself more application friendly for developers.

EC Agent is equipped with three explicit functional modes:

- **Client mode**: The agent provisions a resource access and is consumed by host applications.
- **Server mode**: The server has sole access to a target resource, is tasked to transmitting the data flow between the resource and the Gateway.
- **Gateway mode**: It handles security handshakes, IP filtering, and seeks for the permission from EC service instance by passing on the Client/Server credentials to authorizing requests. Upon authentication, the Gateway performs two-way binding (Client/Server), induces a session, and signifies the requesters for readiness.

A single binary like EC Agent is used to address both scenarios and needs several parameters to work.

--------------------------------------------------------------------------------

## What does this app and how it can help?

This app speed up the configuration of ECAgent script based on your selected scenario using an automated approach to avoid errors due manual activities. The app asks you a very short list of information to make the rest in background for you. Once you have finished, the app:

- download latest version for EC Agent SDK
- generate all the scripts you need to setup the scenario
- structure the folders in a reasonable way to be easily used
- create a step-by-step tutorial to guide you to the goal.

--------------------------------------------------------------------------------

## Setup - What I need to use this app?

Two options are available:

### 1\. Vagrant box

If you are looking for a quick way to use this app instead to install and configure your computer, refer to the Vagrant box setup available [here](https://github.com/indaco/predix-ec-configurator-vagrant).

### 2\. Docker container

If you are looking for a quick way to use this app as docker container, refer to the docker version [here](https://github.com/indaco/predix-ec-configurator-docker).

### 3\. Build the app for yourself

Below what is needed on your local machine to use this application and get started with Enterprise Connect.

Be sure to have the following software installed and configured on your machine:

- A working internet connection:

  - it is used to download the latest release for ECAgent-SDK from the GitHub repository
  - if you are behind a proxy be sure to setup `HTTP_PROXY` and `HTTPS_PROXY`environmental variables on your machine

- [Git](https://git-scm.com/): Distributed version control system
- Go (v.1.9.x)

  - Download and install [Go](https://golang.org/doc/install) on your operating system and configure your `GOPATH` environment

- GoVendor: install it to manage dependencies

  ```shell
  $ go get -u github.com/kardianos/govendor
  ```

- [Cloud Foundry CLI](https://github.com/cloudfoundry/cli/releases): Official command line client for Cloud Foundry

- [Predix CLI](https://github.com/PredixDev/predix-cli) (Optional): Command line utility meant to simplify interaction with the Predix Cloud

#### Get, Build and Run!

Get the code:

```shell
$ go get github.com/indaco/predix-ec-configurator
```

Move to the app folder:

```shell
$ cd $GOPATH/src/github.com/indaco/predix-ec-configurator
```

And get all dependencies:

```shell
$ $GOPATH/bin/govendor sync
$ go get
```

Edit `config.json` file with your Predix.io credentials. If you are not on Predix Basic make sure to update the domain and the api endpoint too:

```json
"predix": {
    "domain": "run.aws-usw02-pr.ice.predix.io",
    "api": "https://api.system.aws-usw02-pr.ice.predix.io",
    "username": "<your-predix-username>",
    "password": "<your-predix-password>"
}
...
```

Run the app:

```shell
$ go run main.go
```

or:

```shell
$ go build
$ ./predix-ec-configurator
```

Open a browser window at `http://localhost:9000`

--------------------------------------------------------------------------------

### Resources

- [Predix.io](https://predix.io)
- [Enterprise Connect Service Description](https://www.predix.io/services/service.html?id=2184)
- [Enterprise Connect SDK](https://github.com/Enterprise-connect/ec-sdk)
- [Enterprise Connect Wiki](https://github.com/Enterprise-connect/ec-sdk/wiki)
- [Predix UAA]((https://www.predix.io/services/service.html?id=1172))

--------------------------------------------------------------------------------

### Contribution

- [Alberto Gorni](https://github.com/gorniAbertoGeDigital)

--------------------------------------------------------------------------------

### DISCLAIMER

This is **not** an official development neither from the [GE Digital's Predix Team](https://github.com/predixdev) and [GE Digital's Enterprise Connect Team](https://github.com/Enterprise-Connect)
