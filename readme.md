# Blit

## Set the Current Version
Set the Blit version (testnet) as an environment variable. It will be used in the rest of the commands.

```bash
$ export BLIT_VERSION=$(curl http://testnet.blitchain.net/cosmos/base/tendermint/v1beta1/node_info | jq -r .application_version.version)
```

|:exclamation:  Note for Docker Machine  |
|:-----------------------------------------|
| You may need to set the host from `localhost` to `0.0.0.0` in `~/.blit/config/app.yaml` in order to acceess the node services. Docker machine runs the containers in a VM and binding to `localhost` is overly restrictive in this case. |

# Quick Start

To hit the ground running paste this in your terminal. Requires docker to be installed.

```bash
$ docker run --init -it --rm \
    -v ~/.blit:/home/user/.blit \
    -p 127.0.0.1:26656:26656 \
    -p 127.0.0.1:26657:26657 \
    -p 127.0.0.1:1317:1317 \
    -p 127.0.0.1:9090:9090 \
    blitchain/blitchain:$BLIT_VERSION \
    /bin/bash -c './bin/blitd init my_node_name ; make testnet start'
```

# The long way

To get the code and build from source.

## Get the Code
Clone the repo for the first time.
```
$ git clone -b $BLIT_VERSION  https://github.com/BlitChain/blitchain
$ cd blitchain
```

_Or_ fetch and checkout the version of the existing repo.


```bash
$ cd blitchain
$ git fetch origin $BLIT_VERSION:$BLIT_VERSION
$ git checkout $BLIT_VERSION
```

## Run with Docker (recommended)

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

Make sure `BLIT_VERSION` is set from above.

### Build the Blit container from source

```bash
$ docker compose build
```

### Initialize the node
Initialize your node with and replace `$MY_MONIKOR` with your node name

```bash
$ docker compose run --rm blit ./bin/blitd init $MY_MONIKOR
```

Update config to connect to the testnet

```bash
$ docker compose run --rm blit make testnet
```

### Start syncing the testnet

```bash
$ docker compose up blit
```

-------

## Manual Installation

Follow the steps below to manually install and build the project:

### Prerequisites

- [Pyenv](https://github.com/pyenv/pyenv)
- [Goenv](https://github.com/go-nv/goenv)

### Install Dependencies

First, you'll need to install some dependencies. On a Debian-based system, you can do this with the following command:

```bash
sudo apt-get update -y && sudo apt-get install -y \
    make \
    build-essential \
    libssl-dev \
    zlib1g-dev \
    libbz2-dev \
    libreadline-dev \
    libsqlite3-dev \
    wget \
    curl \
    llvm \
    libncurses5-dev \
    libncursesw5-dev \
    xz-utils \
    tk-dev \
    libffi-dev \
    liblzma-dev \
    python3-openssl \
    git \
    libre2-dev \
    dnsutils
```


### Pyenv build with Python Patch

To build the patched version of Python in `.python-version` run:

```bash
pyenv install --patch < patch
```

###  Install Python Requirements

Now install the required Python packages using pip:


```bash
python -m pip install -r ./blitvm/requirements.txt
```

Now, you should have all the necessary dependencies installed and have built the project. You can start `blitd` with the following command:

### Install the proper Goversion

```bash
$ goenv install
```

###  Build blitd

Now you need to build `blitd` from the source. First, ensure you have the right version of Go installed by checking the `.go-version` file in the project. Then run:

```bash
make build
```


### Init the node


```
./bin/blitd init $MY_MONIKER 
```


### Configure to connect to the Testnet
```
make testnet 
```

### Start the node

Start syncing the testnet
```bash
$ make start
```

## Configure

Depending on your needs the following configs may need to be changed:

### ~/.blitd/config/app.toml
```toml
###############################################################################
###                           API Configuration                             ###
###############################################################################

[api]
# Enable defines if the API server should be enabled.
enable = true

# Swagger defines if swagger documentation should automatically be registered.
swagger = true

# Allow other domains and ports to access the api. This is useful when developing locally.
enabled-unsafe-cors = true
```

 ### ~/.blit/config/config.toml
```toml
[rpc]
laddr = "tcp://127.0.0.1:26657"
cors_allowed_origins = [ "*" ]
```
