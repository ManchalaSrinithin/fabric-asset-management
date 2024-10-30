## Financial Institution Asset Management System

This is a blockchain-based system using Hyperledger Fabric for managing and tracking financial assets.

### Prerequisites

1. Install Go 1.20 or later
2. Install Docker and Docker Compose
3. Install Hyperledger Fabric prerequisites
4. Set up Hyperledger Fabric test network

### Setup Instructions

1. Start the Fabric test network:
```bash
cd fabric-samples/test-network
./network.sh up createChannel -c mychannel -ca
```

2. Deploy the chaincode:
```bash
./network.sh deployCC -ccn basic -ccp ../asset-management/chaincode -ccl go
```

3. Build and run the API:
```bash
cd ../asset-management
go mod tidy
go run ./api
```

### API Endpoints

- POST /assets - Create a new asset
- GET /assets/{id} - Read an asset
- PUT /assets/{id} - Update an asset
- GET /assets/{id}/history - Get asset history

### Docker Build

```bash
docker build -t asset-management .
docker run -p 8080:8080 asset-management
```