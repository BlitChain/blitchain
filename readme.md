# Blit

## Set the Current Version
Set the Blit version (testnet) as an environment variable. It will be used in the rest of the commands.

```bash
$ export BLIT_VERSION=$(curl http://testnet.blitchain.net/cosmos/base/tendermint/v1beta1/node_info | jq -r .application_version.version)
```

|:exclamation:  Note for Docker Machine  |
|:-----------------------------------------|
| You may need to change the host from `localhost` to `0.0.0.0` in `~/.blit/config/app.yaml` in order to acceess the node services. Docker machine runs the containers in a VM and binding to `localhost` is overly restrictive in this case. |

# Quick Start

To hit the ground running paste this in your terminal. Requires docker to be installed. Use the `$BLIT_VERSION` environment variable set above.

```bash
$ docker run --init -it --rm \
    --pull always \
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

|:exclamation:  Note for running a node  |
|:-----------------------------------------|
| When **starting** the node you MUST use `$ $DAEMON_HOME/cosmovisor/current/bin/blitd start` from within the project directory and NOT using the globally linked binary. Otherwise you will get consensus errors. When using the binary as a client you can use the globally linked `$ blitd` |

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

# Systemd + Cosmovisor

## 1\. Update Packages and Install Dependencies

Update your package lists and install the necessary tools.

```bash
sudo apt update && sudo apt install git jq make wget
```

## 2\. Install GoEnv

Set up GoEnv to manage Go versions.

```bash
git clone https://github.com/go-nv/goenv.git ~/.goenv

echo 'export GOENV_ROOT="$HOME/.goenv"' >> ~/.bashrc
echo 'export PATH="$GOENV_ROOT/bin:$PATH"' >> ~/.bashrc
echo 'eval "$(goenv init -)"' >> ~/.bashrc
echo 'export PATH="$GOROOT/bin:$PATH"' >> ~/.bashrc
echo 'export PATH="$PATH:$GOPATH/bin"' >> ~/.bashrc

# Reload bash configuration
source ~/.bashrc
```
## 3\. Install Go

Install and set the global Go version (e.g., 1.21.5).

```bash
goenv install 1.21.5
goenv global 1.21.5
```

## 4\. Install Cosmovisor

Install Cosmovisor for managing blockchain daemon upgrades.

```
go install cosmossdk.io/tools/cosmovisor/cmd/cosmovisor@latest
```

## 5\. Install PyEnv and Dependencies

Set up PyEnv for Python version management.

```bash
sudo apt update && sudo apt install build-essential libssl-dev zlib1g-dev \
libbz2-dev libreadline-dev libsqlite3-dev curl \
libncursesw5-dev xz-utils tk-dev libxml2-dev libxmlsec1-dev libffi-dev liblzma-dev libre2-dev

git clone https://github.com/pyenv/pyenv.git ~/.pyenv

echo 'export PYENV_ROOT="$HOME/.pyenv"' >> ~/.bashrc
echo 'command -v pyenv >/dev/null || export PATH="$PYENV_ROOT/bin:$PATH"' >> ~/.bashrc
echo 'eval "$(pyenv init -)"' >> ~/.bashrc

# Reload bash configuration
source ~/.bashrc
```

## 6\. Set Up Cosmovisor and Systemd for BlitChain

Prepare Cosmovisor and Systemd to run the BlitChain daemon.


```bash

# The current Blitchain version
export BLIT_VERSION=$(curl http://testnet.blitchain.net/cosmos/base/tendermint/v1beta1/node_info | jq -r .application_version.version)
echo $BLIT_VERSION

# This is the normal location 
export DAEMON_HOME=$HOME/.blit

# Prepare Cosmovisor
mkdir -p $DAEMON_HOME/cosmovisor/
```

### Get up the pre-upgrade helper
Write your choice of upgrade helper to `$DAEMON_HOME/cosmovisor/cosmovisor-preupgrade.sh`

#### (option 1) Get the pre-upgrade helper that pulls the Docker containter
This will pull the pre build Docker container. It is still possible to use `$ blitd` directly with the shell script that is a wrapper around the container.
See: https://github.com/BlitChain/blitchain/blob/develop/scripts/cosmovisor-preupgrade-pull-docker.sh#L10-L50

```bash
curl https://raw.githubusercontent.com/BlitChain/blitchain/develop/scripts/cosmovisor-preupgrade-docker.sh > $DAEMON_HOME/cosmovisor/cosmovisor-preupgrade.sh
```
#### (option 2) Get the pre-upgrade helper that builds from source
This will build every upgrade locally and link it globally.

|:exclamation:  Note for running a node  |
|:-----------------------------------------|
| When **starting** the node you MUST use `$ $DAEMON_HOME/cosmovisor/current/bin/blitd start` from within the project directory and NOT using the globally linked binary. Otherwise you will get consensus errors. When using the binary as a client you can use the globally linked `$ blitd` |

```bash
curl https://raw.githubusercontent.com/BlitChain/blitchain/develop/scripts/cosmovisor-preupgrade-build-from-source.sh > $DAEMON_HOME/cosmovisor/cosmovisor-preupgrade.sh
```

### Set up the current Blitchain version
```bash
# Run cosmovisor-preupgrade.sh
bash $DAEMON_HOME/cosmovisor/cosmovisor-preupgrade.sh $BLIT_VERSION
ln -s $DAEMON_HOME/cosmovisor/upgrades/$BLIT_VERSION $DAEMON_HOME/cosmovisor/current

# Link the binary for global access
sudo ln -s $DAEMON_HOME/cosmovisor/current/bin/blitd /usr/local/bin/blitd

# Clear the binary cache
hash -r
```

### Initialize your node if you haven't already. 
Replace 'my_node_name' with your desired node name. 

If you chose **option 1** and are running blitd with Docker run the commands int he container to initialize the node and set the config for the testnet.
```
docker run -it --rm \
    -u $(id -u):$(id -g) \
    -v $DAEMON_HOME:/home/user/.blit \
    blitchain/blitchain:$BLIT_VERSION \
    ./bin/blitd init my_node_name

docker run -it --rm \
    -u $(id -u):$(id -g) \
    -v $DAEMON_HOME:/home/user/.blit \
    blitchain/blitchain:$BLIT_VERSION \
    make testnet
```

If you chose **option 2** and are running blitd from source init the node and set up the config for the testnet.
```bash
cd $DAEMON_HOME/cosmovisor/current/
./bin/blitd init my_node_name

make testnet
```


### Create the Systemd service for Blitchain

```bash
sudo tee /etc/systemd/system/blit-cosmovisor.service > /dev/null <<EOF
[Unit]
Description=Blitchain Daemon
After=network-online.target

[Service]
User=$USER
ExecStart=$(goenv which cosmovisor) run start
Restart=always
RestartSec=3
Environment="DAEMON_HOME=$DAEMON_HOME"
Environment="DAEMON_NAME=blitd"
Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=false"
Environment="DAEMON_RESTART_AFTER_UPGRADE=true"
Environment=DAEMON_POLL_INTERVAL=1s
Environment=DAEMON_LOG_BUFFER_SIZE=512
Environment=DAEMON_PREUPGRADE_MAX_RETRIES=10
Environment=COSMOVISOR_CUSTOM_PREUPGRADE=cosmovisor-preupgrade.sh
WorkingDirectory=$DAEMON_HOME/cosmovisor/current/
Environment=PATH=$HOME/.pyenv/shims:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
LimitNOFILE=infinity
LimitNPROC=infinity

[Install]
WantedBy=multi-user.target
EOF
```

## 7\. Start and Monitor the Service

Enable and start the BlitChain service, and check its status.

```bash
sudo -S systemctl daemon-reload
sudo -S systemctl enable blit-cosmovisor
sudo -S systemctl start blit-cosmovisor
sudo service blit-cosmovisor status
 
# Monitor the logs
journalctl -f
```
