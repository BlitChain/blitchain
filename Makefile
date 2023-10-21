.PHONY: watch testnet build start

watch:
	@echo "Reading file list from stdin..."
	@while read -r filepath; do \
		echo $$filepath; \
	done | entr -s ' \
		TXHASH=`blitd tx script run --caller-address $$ADDRESS  --script-address $$ADDRESS  --extra-code "`cat /tmp/extracode.py`" --from $$ADDRESS -o json -y | jq | tee /dev/tty | jq -r .txhash` ; \
		sleep 2 ; \
		./blitd q tx $$TXHASH -o json | jq -r .raw_log | tee /dev/tty | jq | tee /dev/tty | jq -r .[0].events[1].attributes[3].value | jq '


testnet:
	RPC_HOST="192.46.233.149" bash ./scripts/setenv.bash



build:
	@cd ./src && \
	export GIT_VERSION=`git rev-parse --abbrev-ref HEAD` && \
	export GIT_COMMIT=`git rev-parse HEAD` && \
	go build -mod=readonly \
		-ldflags "-X github.com/cosmos/cosmos-sdk/version.Name=blit \
		-X github.com/cosmos/cosmos-sdk/version.AppName=blitd \
		-X github.com/cosmos/cosmos-sdk/version.Version=$${GIT_VERSION} \
		-X github.com/cosmos/cosmos-sdk/version.Commit=$${GIT_COMMIT}" \
		-trimpath \
		-o ../blitd \
		./cmd/blitd 


start:
	./blitd start
