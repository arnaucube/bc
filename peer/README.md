# Peer

To run as a normal peer:
```
./peer
```

To run as a p2p server:
```
./peer server
```

Needs the config.json file:
```json
{
    "ip": "127.0.0.1",
    "port": "3001",
    "serverip": "127.0.0.1",
    "serverport": "3000"
}
```


## Peer REST API
- GET /
    - Returns the peer.ID (where peer.ID = hash(peer.IP + ":" + peer.Port))


- GET /peers
    - Returns the peer outcomingPeersList (the peers which the peer have connection)

```json
{
    "PeerID": "VOnL-15rFsUiCnRoyGFksKvWKcwNBRz5iarRem0Ilvo=",
    "peerslist": [
        {
            "id": "VOnL-15rFsUiCnRoyGFksKvWKcwNBRz5iarRem0Ilvo=",
            "ip": "127.0.0.1",
            "port": "3000",
            "role": "server",
            "conn": null
        },
        {
            "id": "Lk9jEP1YcOAzl51yY61GdWADNe35_g5Teh12JeguHhA=",
            "ip": "127.0.0.1",
            "port": "3003",
            "role": "client",
            "conn": {}
        },
        {
            "id": "xj78wuyN2_thFBsXOUXnwij4L8vualxQ9GnVRK6RS4c=",
            "ip": "127.0.0.1",
            "port": "3005",
            "role": "client",
            "conn": {}
        }
    ],
    "date": "0001-01-01T00:00:00Z"
}
```

- POST /register
    - Adds the address (pubK of the user) to the blockchain


## TODO
- When a peer connects to the network, sends his last Block, and receives the new Blocks from this last Block --> DONE with REST petitions, maybe is better with tcp conn
- Delete the peer from the peers list when the connection is closed --> DONE
- REST:
    - endpoint to get if the address is in the blockchain (to verify users)
- parameters Date or LastUpdate on the structs needs to be updated values
- implement rsa encryption between peers
- store blockchain in a .data file
