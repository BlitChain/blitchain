.PHONY: watch build

build: 
	ignite chain build -v --release -t linux:amd64 -t darwin:amd64 -t darwin:arm64 --check-dependencies


watch:
	@echo "Reading file list from stdin..."
	@while read -r filepath; do \
		echo $$filepath; \
	done | entr -s ' \
		TXHASH=`blitd tx script run --caller-address $$ADDRESS  --script-address $$ADDRESS  --extra-code "`cat /tmp/extracode.py`" --from $$ADDRESS -o json -y | jq | tee /dev/tty | jq -r .txhash` ; \
		sleep 2 ; \
		blitd q tx $$TXHASH -o json | jq -r .raw_log | tee /dev/tty | jq | tee /dev/tty | jq -r .[0].events[1].attributes[3].value | jq \
	'

