package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
    "path/filepath"

    "github.com/gorilla/mux"
    "github.com/hyperledger/fabric-gateway/pkg/client"
    "github.com/hyperledger/fabric-gateway/pkg/identity"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
)

type Gateway struct {
    contract *client.Contract
}

func main() {
    // Initialize the gateway connection
    gateway, err := initGateway()
    if err != nil {
        log.Fatalf("Failed to initialize gateway: %v", err)
    }

    router := mux.NewRouter()
    
    // Register routes
    router.HandleFunc("/assets", gateway.createAsset).Methods("POST")
    router.HandleFunc("/assets/{id}", gateway.readAsset).Methods("GET")
    router.HandleFunc("/assets/{id}", gateway.updateAsset).Methods("PUT")
    router.HandleFunc("/assets/{id}/history", gateway.getAssetHistory).Methods("GET")

    log.Printf("Server starting on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func initGateway() (*Gateway, error) {
    clientConnection := newGrpcConnection()
    id := newIdentity()
    sign := newSign()

    gw, err := client.Connect(
        id,
        client.WithSign(sign),
        client.WithClientConnection(clientConnection),
        client.WithEvaluateTimeout(5),
        client.WithEndorseTimeout(15),
        client.WithSubmitTimeout(5),
        client.WithCommitStatusTimeout(1),
    )
    if err != nil {
        return nil, err
    }

    network := gw.GetNetwork("mychannel")
    contract := network.GetContract("basic")
    
    return &Gateway{contract: contract}, nil
}

func (g *Gateway) createAsset(w http.ResponseWriter, r *http.Request) {
    var asset struct {
        ID        string  `json:"id"`
        DealerID  string  `json:"dealerId"`
        MSISDN    string  `json:"msisdn"`
        MPIN      string  `json:"mpin"`
        Balance   float64 `json:"balance"`
        Status    string  `json:"status"`
    }

    if err := json.NewDecoder(r.Body).Decode(&asset); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err := g.contract.SubmitTransaction(
        "CreateAsset",
        asset.ID,
        asset.DealerID,
        asset.MSISDN,
        asset.MPIN,
        fmt.Sprintf("%f", asset.Balance),
        asset.Status,
    )

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}