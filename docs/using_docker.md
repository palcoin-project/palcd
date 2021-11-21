# Using Docker


## Introduction

With Docker you can easily set up *palcd* to run your Palcoin full node. You can find the official *palcd* Docker images on Docker Hub [palcoin-project/palcd](https://hub.docker.com/r/palcoin-project/palcd). The Docker source file of this image is located at [Dockerfile](https://github.com/palcoin-project/palcd/blob/master/Dockerfile).

This documentation focuses on running Docker container with *docker-compose.yml* files. These files are better to read and you can use them as a template for your own use. For more information about Docker and Docker compose visit the official [Docker documentation](https://docs.docker.com/).

## Docker volumes

**Special diskspace hint**: The following examples are using a Docker managed volume. The volume is named *palcd-data* This will use a lot of disk space, because it contains the full Bitcoin blockchain. Please make yourself familiar with [Docker volumes](https://docs.docker.com/storage/volumes/).

The *palcd-data* volume will be reused, if you upgrade your *docker-compose.yml* file. Keep in mind, that it is not automatically removed by Docker, if you delete the palcd container. If you don't need the volume anymore, please delete it manually with the command:

```bash
docker volume ls
docker volume rm palcd-data
```

For binding a local folder to your *palcd* container please read the [Docker documentation](https://docs.docker.com/). The preferred way is to use a Docker managed volume.

## Known error messages when starting the palcd container

We pass all needed arguments to *palcd* as command line parameters in our *docker-compose.yml* file. It doesn't make sense to create a *palcd.conf* file. This would make things too complicated. Anyhow *palcd* will complain with following log messages when starting. These messages can be ignored:

```bash
Error creating a default config file: open /sample-palcd.conf: no such file or directory
...
[WRN] palcd: open /root/.palcd/palcd.conf: no such file or directory
```

## Examples

### Preamble

All following examples uses some defaults:

- container_name: palcd
  Name of the docker container that is be shown by e.g. ```docker ps -a```

- hostname: palcd **(very important to set a fixed name before first start)**
  The internal hostname in the docker container. By default, docker is recreating the hostname every time you change the *docker-compose.yml* file. The default hostnames look like *ef00548d4fa5*. This is a problem when using the *palcd* RPC port. The RPC port is using a certificate to validate the hostname. If the hostname changes you need to recreate the certificate. To avoid this, you should set a fixed hostname before the first start. This ensures, that the docker volume is created with a certificate with this hostname.

- restart: unless-stopped
  Starts the *palcd* container when Docker starts, except that when the container is stopped (manually or otherwise), it is not restarted even after Docker restarts.

To use the following examples create an empty directory. In this directory create a file named *docker-compose.yml*, copy and paste the example into the *docker-compose.yml* file and run it.

```bash
mkdir ~/palcd-docker
cd ~/palcd-docker
touch docker-compose.yaml
nano docker-compose.yaml (use your favourite editor to edit the compose file)
docker-compose up (creates and starts a new palcd container)
```

With the following commands you can control *docker-compose*:

```docker-compose up -d``` (creates and starts the container in background)

```docker-compose down``` (stops and delete the container. **The docker volume palcd-data will not be deleted**)

```docker-compose stop``` (stops the container)

```docker-compose start``` (starts the container)

```docker ps -a``` (list all running and stopped container)

```docker volume ls``` (lists all docker volumes)

```docker logs palcd``` (shows the log )

```docker-compose help``` (brings up some helpful information)

### Full node without RPC port

Let's start with an easy example. If you just want to create a full node without the need of using the RPC port, you can use the following example. This example will launch *palcd* and exposes only the default p2p port 8333 to the outside world:

```yaml
version: "2"

services:
  palcd:
    container_name: palcd
    hostname: palcd
    image: palcoin-project/palcd:latest
    restart: unless-stopped
    volumes:
      - palcd-data:/root/.palcd
    ports:
      - 8333:8333

volumes:
  palcd-data:
```

### Full node with RPC port

To use the RPC port of *palcd* you need to specify a *username* and a very strong *password*. If you want to connect to the RPC port from the internet, you need to expose port 8334(RPC) as well.

```yaml
version: "2"

services:
  palcd:
    container_name: palcd
    hostname: palcd
    image: palcoin-project/palcd:latest
    restart: unless-stopped
    volumes:
      - palcd-data:/root/.palcd
    ports:
      - 8333:8333
      - 8334:8334
    command: [
        "--rpcuser=[CHOOSE_A_USERNAME]",
        "--rpcpass=[CREATE_A_VERY_HARD_PASSWORD]"
    ]

volumes:
  palcd-data:
```

### Full node with RPC port running on TESTNET

To run a node on testnet, you need to provide the *--testnet* argument. The ports for testnet are 18333 (p2p) and 18334 (RPC):

```yaml
version: "2"

services:
  palcd:
    container_name: palcd
    hostname: palcd
    image: palcoin-project/palcd:latest
    restart: unless-stopped
    volumes:
      - palcd-data:/root/.palcd
    ports:
      - 18333:18333
      - 18334:18334
    command: [
        "--testnet",
        "--rpcuser=[CHOOSE_A_USERNAME]",
        "--rpcpass=[CREATE_A_VERY_HARD_PASSWORD]"
    ]

volumes:
  palcd-data:
```
