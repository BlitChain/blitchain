# Blit

Manual Installation
-------------------

Follow the steps below to manually install and build the project:

### 1\. Install Dependencies

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

Also install Go from the [Offical Docs](https://go.dev/dl/) or using goenv



```
curl -OL https://go.dev/dl/go1.20.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xvf ./go1.20.3.linux-amd64.tar.gz

echo '
export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:/usr/local/go/bin:$GOBIN' >> ~/.bashrc
```

### 3\. Install Pyenv

Similarly, for managing Python versions, you can use `pyenv`. Here's how to install it:

```bash
git clone --depth=1 https://github.com/pyenv/pyenv.git ~/.pyenv
echo 'export PYENV_ROOT="$HOME/.pyenv"' >> ~/.bashrc
echo 'export PATH="$PYENV_ROOT/bin:$PATH"' >> ~/.bashrc
echo -e 'if command -v pyenv 1>/dev/null 2>&1; then\n  eval "$(pyenv init -)"\nfi' >> ~/.bashrc
source ~/.bashrc
```

### 4\. Clone the Repository

Now, clone the Blit repository to your local machine:

```bash
git clone --depth=1  --recursive  https://github.com/BlitChain/blitchain
cd blitchain
```

### 5\. Pyenv build with Python Patch

To build the patched version of Python in `.python-version` run:

```bash
pyenv install --patch < patch
```

### 6\. Install Python Requirements

Now install the required Python packages using pip:


```bash
python -m pip install -r ./blitvm/requirements.txt
```

Now, you should have all the necessary dependencies installed and have built the project. You can start `blitd` with the following command:


### 7\. Build blitd

Now you need to build `blitd` from the source. First, ensure you have the right version of Go installed by checking the `.go-version` file in the project. Then run:

```bash
make build
```


### Init the node


```
./bin/blitd init $MY_MONIKER 
```


## Connect to the Testnet and start the node
```
make testnet start
```



