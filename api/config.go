package main

const (
    mspID        = "Org1MSP"
    cryptoPath   = "../../test-network/organizations/peerOrganizations/org1.example.com"
    certPath     = cryptoPath + "/users/User1@org1.example.com/msp/signcerts/cert.pem"
    keyPath      = cryptoPath + "/users/User1@org1.example.com/msp/keystore/"
    tlsCertPath  = cryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt"
)