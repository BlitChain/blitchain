FROM golang:bookworm

ENV DEBIAN_FRONTEND=noninteractive

WORKDIR /app
COPY go.mod go.sum .go-version ./
COPY cosmos-sdk ./cosmos-sdk
RUN go mod download -x

COPY . .

RUN GOWORK=off go build \
        -mod=readonly \
        -ldflags \
            "-X github.com/cosmos/cosmos-sdk/version.Name="blit" \
            -X github.com/cosmos/cosmos-sdk/version.AppName="blitd" \
            -X github.com/cosmos/cosmos-sdk/version.Version=${GIT_VERSION} \
            -X github.com/cosmos/cosmos-sdk/version.Commit=${GIT_COMMIT} \
            -X github.com/cosmos/cosmos-sdk/version.BuildTags=${BUILD_TAGS} \
            -w -s -linkmode=external -extldflags '-Wl,-z,muldefs -static'" \
        -trimpath \
        -o ./blitd \
        ./cmd/blitd

FROM debian:bookworm-slim

RUN apt-get update -y \
    && apt-get install -y \
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
        dnsutils \
    && rm -rf /var/lib/apt/lists/*

RUN groupadd -r user && useradd -r -g user -ms /bin/bash user

USER user
ENV HOME  /home/user
RUN chown -R user:user $HOME
WORKDIR $HOME/blit

ENV PYENV_ROOT "$HOME/.pyenv"
RUN git clone --depth=1 https://github.com/pyenv/pyenv.git $PYENV_ROOT
ENV PATH "$PYENV_ROOT/shims:$PYENV_ROOT/bin:$PATH"

COPY --chown=user:user .python-version .
COPY --chown=user:user patch .
COPY --chown=user:user blit-python .
COPY --chown=user:user .python-version .
RUN pyenv install --patch < patch
COPY --chown=user:user blitvm/requirements.txt ./blitvm/requirements.txt
RUN python -m pip install --user -r ./blitvm/requirements.txt
COPY --chown=user:user ./blitvm ./blitvm

COPY --from=0 --chown=user:user /app/blitd ./blitd

CMD ["./blitd", "start"]

EXPOSE 1317
EXPOSE 26657
EXPOSE 26656
