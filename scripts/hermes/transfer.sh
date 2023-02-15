#!/bin/sh

CHAIN_A_ARGS="--from juno1 --keyring-backend test --chain-id local-1 --home ~/.baobab2/ --node http://localhost:26657 --yes"

# Send from local-1 to local-2 via the relayer
baobabd tx ibc-transfer transfer transfer channel-0 juno1hj5fveer5cjtn4wd6wstzugjfdxzl0xps73ftl 9ubaobab "$CHAIN_A_ARGS" --packet-timeout-height 0-0

sleep 6

# check the query on the other chain to ensure it went through
baobabd q bank balances juno1hj5fveer5cjtn4wd6wstzugjfdxzl0xps73ftl --chain-id local-2 --node http://localhost:36657
