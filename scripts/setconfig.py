import requests
import json
import os
import toml
import logging

# Setup logging
logging.basicConfig(level=logging.INFO)


# Define Constants
NODE_IP = os.environ.get("NODE_IP", "127.0.0.1")
HOME_DIR = os.path.expanduser("~/.blit")

def fetch_data_from_rpc(endpoint):
    try:
        response = requests.get(f"http://{NODE_IP}:26657/{endpoint}")
        response.raise_for_status()
        return response.json()
    except requests.RequestException as e:
        logging.error(f"Failed to fetch data from {endpoint}. Error: {e}")
        exit(1)

def update_config_file(config_path, updates):
    try:
        with open(config_path, "r") as f:
            config = toml.load(f)

        # Update values in the TOML config
        for key, value in updates.items():
            print(f"Updating {config_path}: {key}: {value}")
            keys = key.split('.')
            d = config
            for k in keys[:-1]:
                d = d[k]
            d[keys[-1]] = value

        with open(config_path, "w") as f:
            toml.dump(config, f)

    except Exception as e:
        logging.error(f"Failed to update config file. Error: {e}")
        exit(1)

def main():
    # Fetch node ID
    data = fetch_data_from_rpc("status")
    NODE_ID = data['result']['node_info']['id']

    # Fetch block data
    data = fetch_data_from_rpc("commit")
    BLOCK_HEIGHT = data['result']['signed_header']['header']['height']
    TRUST_HASH = data['result']['signed_header']['commit']['block_id']['hash']

    logging.info(f"TRUST HEIGHT: {BLOCK_HEIGHT}")
    logging.info(f"TRUST HASH: {TRUST_HASH}")

    # Fetch and save genesis file
    data = fetch_data_from_rpc("genesis")
    genesis_data = data['result']['genesis']
    try:
        with open(os.path.join(HOME_DIR, "config", "genesis.json"), "w") as f:
            json.dump(genesis_data, f, indent=4)
    except Exception as e:
        logging.error(f"Failed to write to genesis.json. Error: {e}")
        exit(1)

    # Define config updates
    BLIT_P2P_MAX_NUM_OUTBOUND_PEERS = 100
    BLIT_P2P_MAX_NUM_INBOUND_PEERS = 100
    BLIT_STATESYNC_RPC_SERVERS = f"http://{NODE_IP}:26657,http://{NODE_IP}:26657"
    BLIT_P2P_SEEDS = f"{NODE_ID}@{NODE_IP}:26656"
    BLIT_P2P_PERSISTENT_PEERS = f"{NODE_ID}@{NODE_IP}:26656"

    updates = {
        'p2p.persistent_peers': BLIT_P2P_PERSISTENT_PEERS,
        'p2p.seeds': BLIT_P2P_SEEDS,
        'p2p.max_num_outbound_peers': BLIT_P2P_MAX_NUM_OUTBOUND_PEERS,
        'p2p.max_num_inbound_peers': BLIT_P2P_MAX_NUM_INBOUND_PEERS,
        'statesync.enable': True,
        'statesync.rpc_servers': BLIT_STATESYNC_RPC_SERVERS,
        'statesync.trust_height': BLOCK_HEIGHT,
        'statesync.trust_hash': TRUST_HASH,
        'statesync.double_sign_check_height': 1,
        'statesync.discovery_time': '5s'
    }

    # Update the config file
    config_path = os.path.join(HOME_DIR, "config", "config.toml")
    update_config_file(config_path, updates)

    app_config_path = os.path.join(HOME_DIR, "config", "app.toml")
    update_config_file(app_config_path, {
        'state-sync.enable': True,
        'minimum-gas-prices': '0.0ublit',
        'api.enable': True,
        'api.swagger': True,
        'consensus.timeout_propose': '3s',
        'consensus.timeout_propose_delta': '500ms',
        'consensus.timeout_prevote': '1s',
        'consensus.timeout_prevote_delta': '500ms',
        'consensus.timeout_precommit': '1s',
        'consensus.timeout_precommit_delta': '500ms',
        'consensus.timeout_commit': '5s',
    })


if __name__ == "__main__":
    main()
