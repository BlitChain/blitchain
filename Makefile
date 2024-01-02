.PHONY: testnet build start installdeps

testnet:
	NODE_IP="149.248.76.117" python ./scripts/setconfig.py

mainnet:
	NODE_IP="172.104.250.57" python ./scripts/setconfig.py

installdeps:
	pyenv install -s --patch < patch
	python -m pip install -r ./blitvm/requirements.txt

build:
	@cd ./src && \
	export GIT_VERSION=`git symbolic-ref -q --short HEAD || git describe --tags --exact-match 2> /dev/null   || git rev-parse --short HEAD` && \
	export GIT_COMMIT=`git rev-parse HEAD` && \
	go build -mod=readonly \
		-ldflags "-X github.com/cosmos/cosmos-sdk/version.Name=blit \
		-X github.com/cosmos/cosmos-sdk/version.AppName=blitd \
		-X github.com/cosmos/cosmos-sdk/version.Version=$${GIT_VERSION} \
		-X github.com/cosmos/cosmos-sdk/version.Commit=$${GIT_COMMIT}" \
		-trimpath \
		-o ../bin/blitd \
		./cmd/blitd 

start:
	exec ./bin/blitd start
