# Running Validator Node using Ansible
<!-- markdownlint-disable MD033 -->

## Prerequisites

Make sure you have all [prerequisites](./prerequisites.md) set up

## Configure DCL network parameters (local machine)

### 1. Set network chain ID

[`deployment/ansible/inventory/hosts.yml`]

```yaml
all:
  vars:
    chain_id: <chain-id>
  ...
```

Every network must have a unique chain ID (e.g. `test-net`, `main-net` etc.)

<details>
<summary>Example for Testnet 2.0 (clickable) </summary>

```yaml
all:
  vars:
    chain_id: testnet-2.0
  ...
```

</details>

### 2. Put `genesis.json` file under specific directory

- Get or download `genesis.json` file of a network your node will be joining and put it under the following path:

  ```text
  deployment/persistent_chains/<chain-id>/genesis.json
  ```

  where `<chain-id>` is the chain ID of a network spefied in the previous step.

  <details>
  <summary>Example for Testnet 2.0 (clickable) </summary>

  For `testnet-2.0` the genesis file is already in place. So you don't need to do anything!

  ```text
  deployment/persistent_chains/testnet-2.0/genesis.json
  ```

  </details>

## Configure node type specific parameters (local machine)

### 1. Specify target instance address in the inventory file

[`deployment/ansible/inventory/hosts.yml`]

```yaml
all:
  ...
  children:
    ...
    validators:
      hosts:
        <validator node IP address or hostname>
    ...
```

### 2. Set persistent peers string in validator configuration

[`deployment/ansible/roles/configure/vars/validator.yml`]

```yaml
config:
  p2p:
    persistent_peers: "<node1-ID>@<node1-IP>:26656,..."
  ...
```

- For `testnet-2.0` get the latest `persistent_peers` string from the CSA slack channel

### 3. (Optional) If you are joining a long-running network, enable `statesync` or use one of the options in [running-node-in-existing-network.md](../advanced/running-node-in-existing-network.md)

[`deployment/ansible/roles/configure/vars/validator.yml`]

```yaml
config:
...
  statesync:
    enable: true
    rpc_servers: "http(s):<node1-IP>:26657, ..."
    trust_height: <trust-height>
    trust_hash: "<trust-hash>"
...
```

</details>

> **_NOTE:_**  You should provide at least 2 addresses for `rpc_servers`. It can be 2 identical addresses

You can use the following command to obtain `<trust-height>` and `<trust-hash>` of your network

```bash
curl -s http(s)://<host>:<port>/commit | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
```

<details>
<summary>Example for Testnet 2.0 (clickable) </summary>

```bash
curl -s https://on.test-net.dcl.csa-iot.org:26657/commit | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
```

</details>

- `<host>` - RPC endpoint host of the network being joined
- `<port>` - RPC endpoint port of the network being joined

> **_NOTE:_** State sync is not attempted if the node has any local state (LastBlockHeight > 0)

## Run ansible (local machine)

### 1. Verify that all the configuration parameters from the previous section are correct

### 2. Run ansible

```bash
ansible-playbook -i ./deployment/ansible/inventory  -u <target-host-ssh-user> ./deployment/ansible/deploy.yml
```

- `<target-host-ssh-username>` - target host ssh user
- Ansible provisioning can take several minutes depending on number of nodes being provisioned

## Deployment Verification (target machine)

### 1. Switch to cosmovisor user

```bash
sudo su -s /bin/bash cosmovisor
```

### 2. Query status

```bash
dcld status
```

## Make your node a Validator (target machine)

### 1. Switch to cosmovisor user in the Validator

```bash
sudo su -s /bin/bash cosmovisor
```

### 2. Generate NodeAdmin keys

```bash
dcld keys add "<admin-account-name>" 2>&1 | tee "<admin-account-name>.dclkey.data"
```

IMPORTANT keep generated data (especially the mnemonic) securely.

### 3. Share generated address and pubkey with the community

address and pubkey can be found using

```bash
dcld keys show --output text "<admin-account-name>"
```

### 4. Wait until your NodeAdmin key is proposed and approved by the quorum of trustees

- Make sure the Node Admin account is proposed by a Trustee (usually CSA). The account will appear in the "Accounts" / "All Proposed Accounts" tab in <https://testnet.iotledger.io/accounts> (for `testnet-2.0` only).

- Make sure that the proposed account is approved by at least 2/3 of Trustees. The account must disappear from the "Accounts" / "All Proposed Accounts" tab in <https://testnet.iotledger.io/accounts>, and appear in  "Accounts" / "All Active Accounts" tab.

### 5. Check the account presence on the ledger

`dcld query auth account --address="<address>"`

### 6. Make the node a validator

```bash
dcld tx validator add-node --pubkey="<protobuf JSON encoded validator-pubkey>" --moniker="<node-name>" --from="<admin-account-name>"
```

> **_Note:_** Get `<protobuf JSON encoded validator-pubkey>` using `dcld tendermint show-validator` command

(once transaction is successfully written you should see "code": 0 in the JSON output.)

### 7. Make sure the VN participates in consensus

`dcld query tendermint-validator-set` must contain the VN's address

>**_Note:_** Get your VN's address using `dcld tendermint show-address` command.
