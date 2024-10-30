package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func (g *Gateway) readAsset(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    result, err := g.contract.EvaluateTransaction("ReadAsset", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(result)
}

func (g *Gateway) updateAsset(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var update struct {
        Balance     float64 `json:"balance"`
        Status      string  `json:"status"`
        TransAmount float64 `json:"transAmount"`
        TransType   string  `json:"transType"`
        Remarks     string  `json:"remarks"`
    }

    if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err := g.contract.SubmitTransaction(
        "UpdateAsset",
        id,
        fmt.Sprintf("%f", update.Balance),
        update.Status,
        fmt.Sprintf("%f", update.TransAmount),
        update.TransType,
        update.Remarks,
    )

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func (g *Gateway) getAssetHistory(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    result, err := g.contract.EvaluateTransaction("GetAssetHistory", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(result)
}