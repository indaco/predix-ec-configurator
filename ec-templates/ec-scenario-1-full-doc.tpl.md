## Predix-EC-Configurator - Scenario 1: Step-by-Step Doc

**_This is an auto-generated file based on your inputs_. You can find it and all other auto-generated scripts saved into the `output` folder**

Here is a step-to-step document to setup EC for the selected scenario.

### 1. Login to Predix.io

```shell
$ cf login // or predix login if you use Predix CLI
```

### 2. Install Diego-Enabler plugin

Cloud Foundry uses the Diego architecture to manage application containers. Diego components assume application scheduling and management responsibility from the Cloud Controller.

Enable Diego support for an app running on Cloud Foundry.

```shell
$ cf add-plugin-repo CF-Community https://plugins.cloudfoundry.org/
$ cf install-plugin Diego-Enabler -r CF-Community
```

### 3. EC Agent Gateway

Here is the content for `ec-gateway.sh` file

```shell
<gateway_script_content_here>
```

Here is the content for `manifest.yml` file

```shell
<gateway_manifest_content_here>
```

#### Deploy the Agent Gateway to the Predix cloud

It is time now to push the EC Agent Gateway app to Predix.io

**NOTE:** The `output/gateway` folder contains a `push-gateway` script to do all the following in a single step.

```shell
$ cd output/gateway
$ cf push
```

Enable Diego support:

```shell
$ cf enable-diego <ecagent_gateway_name>
```

Now it is time to map CF Route to the Gateway app with:

```shell
$ cf map-route <ecagent_gateway_name> <predix_domain> -n <ecagent_gateway_name>
```

and start the EC Agent Gateway:

```shell
$ cf start <ecagent_gateway_name>
```

Check if it works opening a browser windows at `https://<ecagent_gateway_name>.<predix_domain>/health`

### 4. EC Agent Server

Here is the content for `ec-server` file

```shell
<server_script_content_here>
```

This script has to be executed locally on your machine.

### 5. EC Agent Client

Here is the content for `ec-client.sh` file

```shell
<client_script_content_here>
```

#### Connect to the local data source from your Predix App

Now, th ecagent-client has to be embedded into your application code to be executed on Predix cloud. E.g. in case of access to an on-premise PostgreSQL instance, it means that your PostgreSQL client app has to execute the ECAgent Client script before create a connection to the PostgreSQL server.

You can have a look to the `main.go` file for a sample app available on [GitHub](https://github.com/indaco/ec-go-sample-app) to see what does it mean.

Sample applications in other programming languages are available on the [official GitHub repository](https://github.com/Enterprise-connect/ec-sdk) for Enterprise Connect
