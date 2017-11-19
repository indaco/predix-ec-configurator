## Predix-EC-Configurator - Scenario 2: Step-by-Step Doc

**_This is an auto-generated file based on your inputs_. You can find it and all other auto-generated scripts saved into the `output` folder**

Here is a step-to-step document to setup EC for the selected scenario.

### 0. Login to Predix.io

```shell
$ cf login // or predix login if you use Predix CLI
```

### 1. Install Diego-Enabler plugin

Cloud Foundry uses the Diego architecture to manage application containers. Diego components assume application scheduling and management responsibility from the Cloud Controller.

Enable Diego support for an app running on Cloud Foundry.

```shell
$ cf add-plugin-repo CF-Community https://plugins.cloudfoundry.org/
$ cf install-plugin Diego-Enabler -r CF-Community
```

### 2. EC Agent Gateway

Here is the content for `ec-gateway.sh` file

```shell
<gateway_script_content_here>
```

Here is the content for `manifest.yml` file

```sh
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

Now it is time to map CF Route to the Gateway app with

```shell
$ cf map-route <ecagent_gateway_name> <predix_domain> -n <ecagent_gateway_name>
```

and start the EC Agent Gateway

```shell
$ cf start <ecagent_gateway_name>
```

Check if it works opening a browser windows at `https://<ecagent_gateway_name>.<predix_domain>/health`

### 3. EC Agent Server

Here is the content for `ec-server.sh` file

```shell
<server_script_content_here>
```

Here is the content for `manifest.yml` file

```shell
<server_manifest_content_here>
```

#### Deploy the EC Agent Server to the Predix cloud

It is time now to push the EC Agent Server app to Predix.io

**NOTE:** The `output/server` folder contains a `push-server` script to do all the following in a single step.

```shell
$ cd output/server
$ cf push
```

Enable Diego support:

```shell
$ cf enable-diego <ecagent_server_name>
```

Now, it is time to map CF Route to the Gateway app

```shell
 $ cf map-route <ecagent_server_name> <predix_domain> -n <ecagent_server_name>
```

and start the EC Agent Server

```shell
$ cf start <ecagent_server_name>
```

Check if it works opening a browser windows at `https://<ecagent_server_name>.<predix_domain>/health`

**NOTE:** Verify the server appears as "SupperConns" belongs to the gateway: `https://<ecagent_gateway_name>.<predix_domain>/health` (it may take a minute)

### 4. EC Agent Client

Here is the content for `ec-client` file

```shell
<client_script_content_here>
```

### 5. Connect to the Predix data source from you local machine

You should now be able to use a local client for Predix resource and connect to it.

E.g. If you want to connect to PostgreSQL on Predix, you could download and install on your local machine [PGAdmin](https://www.pgadmin.org/) and create a new server configuration as below:

- Hostname: localhost
- Port: <local_port>
- Database: **postgresql-database-name-from-cf-vcap**
- Username: **postgresql-user-from-cf-vcap**
- Password: **postgresql-password-from-cf-vcap**
