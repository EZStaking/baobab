# https://github.com/osmosis-labs/osmosis/blob/main/tests/ibc-hooks/test_hooks.sh

CHAIN_A_ARGS="--keyring-backend test --chain-id local-1 --home $HOME/.baobab1/ --node http://localhost:26657 --gas 5000000 -b block --yes"
CHAIN_B_ARGS="--keyring-backend test --chain-id local-2 --home $HOME/.baobab2/ --node http://localhost:36657 --gas 5000000 -b block --yes"

# upload contract on chain B (receiving chain)
baobabd tx wasm store ./scripts/ibcmemo/counter.wasm --from juno1 $CHAIN_B_ARGS
CONTRACT_ID=$(baobabd query wasm list-code --node http://localhost:36657 -o json | jq -r '.code_infos[-1].code_id')
echo "contract id: $CONTRACT_ID"

baobabd tx wasm instantiate "$CONTRACT_ID" '{"count":0}' --from juno1 --no-admin --label=counter $CHAIN_B_ARGS
export CONTRACT_ADDRESS=$(baobabd query wasm list-contract-by-code 1 --node http://localhost:36657 -o json | jq -r '.contracts | [last][0]')
echo "contract address: $CONTRACT_ADDRESS" # no balance, new contract

# Send a memo with the wasm message to execute on Chain B from chain A
MEMO=$(jenv -c '{"wasm":{"contract":$CONTRACT_ADDRESS,"msg": {"increment": {}} }}' )
baobabd tx ibc-transfer transfer transfer channel-0 $CONTRACT_ADDRESS 1ubaobab --from juno1 $CHAIN_A_ARGS --memo "$MEMO" --packet-timeout-height 0-0

# Wait for relay
sleep 6

# Ensure the balance has gone up and both are NOT null
denom=$(baobabd query bank balances "$CONTRACT_ADDRESS" --node http://localhost:36657 -o json | jq -r '.balances[0].denom')
balance=$(baobabd query bank balances "$CONTRACT_ADDRESS" --node http://localhost:36657 -o json | jq -r '.balances[0].amount')
echo "denom: $denom"
echo "balance: $balance"

export ADDR_IN_CHAIN_A=$(baobabd q ibchooks wasm-sender channel-0 "juno1hj5fveer5cjtn4wd6wstzugjfdxzl0xps73ftl" --node http://localhost:26657)

# Total Funds
QUERY=$(jenv -c -r '{"get_total_funds": {"addr": $ADDR_IN_CHAIN_A}}')
baobabd query wasm contract-state smart "$CONTRACT_ADDRESS" "$QUERY" --node http://localhost:36657 -o json

# Count
QUERY=$(jenv -c -r '{"get_count": {"addr": $ADDR_IN_CHAIN_A}}')
baobabd query wasm contract-state smart "$CONTRACT_ADDRESS" "$QUERY" --node http://localhost:36657 -o json
