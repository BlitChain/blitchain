FROM debian:bookworm-slim

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update -y \
    && apt-get install -y \
	make \
    wget \
	build-essential \
	git \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app/

ENV HOME  /root
ENV GOENV_ROOT "$HOME/.goenv"

RUN git clone https://github.com/go-nv/goenv.git $GOENV_ROOT
ENV PATH "$GOENV_ROOT/shims:$GOENV_ROOT/bin:$PATH"

COPY .go-version .

RUN goenv install 

COPY src/go.mod src/go.mod
COPY src/go.sum src/go.sum

RUN cd src && go mod download -x

COPY src src
COPY Makefile Makefile

COPY .git .git

RUN make build

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
COPY --chown=user:user blitvm/requirements.txt ./blitvm/requirements.txt
COPY --chown=user:user ./blitvm ./blitvm
COPY --chown=user:user ./scripts/ ./scripts
COPY --chown=user:user ./Makefile .
RUN make installdeps

COPY --from=0 --chown=user:user /app/bin/blitd ./bin/blitd

CMD ["./bin/blitd", "start"]

EXPOSE 1317
EXPOSE 26657
EXPOSE 26656
