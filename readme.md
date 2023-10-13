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


### 3\. Install Pyenv

Similarly, for managing Python versions, you can use `pyenv`. Here's how to install it:

```bash
git clone --depth=1 https://github.com/pyenv/pyenv.git ~/.pyenv
echo 'export PYENV_ROOT="$HOME/.pyenv"' >> ~/.bash_profile
echo 'export PATH="$PYENV_ROOT/bin:$PATH"' >> ~/.bash_profile
echo -e 'if command -v pyenv 1>/dev/null 2>&1; then\n  eval "$(pyenv init -)"\nfi' >> ~/.bash_profile
source ~/.bash_profile
```

### 4\. Clone the Repository

Now, clone the BlitChain repository to your local machine:

```bash
git clone https://github.com/BlitChain/blitchain
cd blitchain
```

### 5\. Build blitd

Now you need to build `blitd` from the source. First, ensure you have the right version of Go installed by checking the `.go-version` file in the project. Then run:

```bash
export GIT_VERSION=`git rev-parse --abbrev-ref HEAD`
export GIT_COMMIT=`git rev-parse HEAD`
go mod download -x
go build\
    -mod=readonly\
    -ldflags\
        "-X github.com/cosmos/cosmos-sdk/version.Name=blit\
        -X github.com/cosmos/cosmos-sdk/version.AppName=blitd\
        -X github.com/cosmos/cosmos-sdk/version.Version=${GIT_VERSION}\
        -X github.com/cosmos/cosmos-sdk/version.Commit=${GIT_COMMIT}\
        -w -s -linkmode=external -extldflags '-Wl,-z,muldefs -static'"\
    -trimpath\
    -o ./blitd\
    ./cmd/blitd
```

### 6\. Pyenv build with Python Patch

To build the patched version of Python in `.python-version` run:

```bash
pyenv install --patch < patch
```

### 7\. Install Python Requirements

Now install the required Python packages using pip:


```bash
python -m pip install -r ./blitvm/requirements.txt
```

Now, you should have all the necessary dependencies installed and have built the project. You can start `blitd` with the following command:


```bash
./blitd start
```

