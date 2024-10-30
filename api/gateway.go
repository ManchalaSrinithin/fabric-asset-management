package main

import (
    "crypto/x509"
    "fmt"
    "os"
    "path"

    "github.com/hyperledger/fabric-gateway/pkg/identity"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
)

func newGrpcConnection() *grpc.ClientConn {
    certificate, err := loadCertificate(tlsCertPath)
    if err != nil {
        panic(fmt.Errorf("failed to load certificate: %v", err))
    }

    certPool := x509.NewCertPool()
    certPool.AddCert(certificate)
    transportCredentials := credentials.NewClientTLSFromCert(certPool, "peer0.org1.example.com")

    connection, err := grpc.Dial("localhost:7051", grpc.WithTransportCredentials(transportCredentials))
    if err != nil {
        panic(fmt.Errorf("failed to create gRPC connection: %v", err))
    }

    return connection
}

func newIdentity() *identity.X509Identity {
    certificate, err := loadCertificate(certPath)
    if err != nil {
        panic(fmt.Errorf("failed to load certificate: %v", err))
    }

    id, err := identity.NewX509Identity("Org1MSP", certificate)
    if err != nil {
        panic(fmt.Errorf("failed to create identity: %v", err))
    }

    return id
}

func newSign() identity.Sign {
    files, err := os.ReadDir(keyPath)
    if err != nil {
        panic(fmt.Errorf("failed to read private key directory: %v", err))
    }

    privateKeyPEM, err := os.ReadFile(path.Join(keyPath, files[0].Name()))
    if err != nil {
        panic(fmt.Errorf("failed to read private key file: %v", err))
    }

    privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
    if err != nil {
        panic(fmt.Errorf("failed to load private key: %v", err))
    }

    sign, err := identity.NewPrivateKeySign(privateKey)
    if err != nil {
        panic(fmt.Errorf("failed to create signer: %v", err))
    }

    return sign
}