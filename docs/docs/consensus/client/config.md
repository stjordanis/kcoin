# Config

| Field                   | Description                |
| ----------------------- | -------------------------- |
| NetworkId               | The network identification |
| SyncMode                |                            |
| LightPeers              | Ignore for now             |
| DatabaseCache           |                            |
| GasPrice                |                            |
| EnablePreimageRecording |                            |

## Transaction Pool

| Field        | Description                |
| ------------ | -------------------------- |
| NoLocals     | The network identification |
| Journal      |                            |
| Rejournal    | Ignore for now             |
| PriceLimit   |                            |
| PriceBump    |                            |
| AccountSlots |                            |
| GlobalSlots  |                            |
| AccountQueue |                            |
| GlobalQueue  |                            |
| Lifetime     |                            |

## GPO

| Field      | Description |
| ---------- | ----------- |
| Blocks     |             |
| Percentile |             |

## Node

| Field   | Description |
| ------- | ----------- |
| DataDir |             |

### P2P

| Field           | Description |
| --------------- | ----------- |
| MaxPeers        |             |
| NoDiscovery     |             |
| DiscoveryV5Addr |             |
| BootstrapNodes  |             |
| StaticNodes     |             |
| TrustedNodes    |             |
| ListenAddr      |             |

# Context

## NetworkID vs ChainID

Chain ID was introduced to prevent replay attacks between the main ETH and ETC chains - both have a network id of 1. Chain ID is an additional way to differentiate chains and it's used during the transaction signing process which will result in different transactions for both networks.

## Sync Modes

# Config Sample

```
[Kowala]
NetworkId = 1
SyncMode = "fast"
LightPeers = 20
DatabaseCache = 128
GasPrice = 25
EnablePreimageRecording = false

    [Kowala.TxPool]
    NoLocals = false
    Journal = "transactions.rlp"
    Rejournal = 3600000000000
    PriceLimit = 1
    PriceBump = 10
    AccountSlots = 16
    GlobalSlots = 4096
    AccountQueue = 64
    GlobalQueue = 1024
    Lifetime = 10800000000000

    [Kowala.GPO]
    Blocks = 10
    Percentile = 50

[Node]
DataDir = ".kowala"

    [Node.P2P]
    MaxPeers = 25
    NoDiscovery = false
    DiscoveryV5Addr = ":30304"
    BootstrapNodes = ["enode://111929beaa10c4eef786996c79267ad041b919830d0c66587f213418d759e83416047df57c9dca9cb86af9b4f53b4491112a905ec7bc2540ec73f1fa066bfb8f@54.162.29.69:33445"]
    # BootstrapNodesV5 = ["enode://bae38251dece370a6c99482364d433244f183ce74169ed148a57960e46b2ab36995c0ee15d54cdff2edfa92a423ab62faecdb8185a5c536c9090507284c87427@192.168.12.171:33445"]
    StaticNodes = []
    TrustedNodes = []
    ListenAddr = ":30303"

[Stats]
URL = "{{.Hostname}}:DVagynuHLdn9sK6c@testnet.kowala.io:80"
```