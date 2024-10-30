package chaincode

import (
    "encoding/json"
    "fmt"
    "log"

    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
    contractapi.Contract
}

func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, id string, dealerId string, msisdn string, mpin string, balance float64, status string) error {
    exists, err := s.AssetExists(ctx, id)
    if err != nil {
        return fmt.Errorf("failed to check asset existence: %v", err)
    }
    if exists {
        return fmt.Errorf("asset already exists: %s", id)
    }

    asset := Asset{
        DocType:     "asset",
        DealerID:    dealerId,
        MSISDN:      msisdn,
        MPIN:        mpin,
        Balance:     balance,
        Status:      status,
        TransAmount: 0,
        TransType:   "CREATE",
        Remarks:     "Asset creation",
    }

    assetJSON, err := json.Marshal(asset)
    if err != nil {
        return err
    }

    return ctx.GetStub().PutState(id, assetJSON)
}

func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
    assetJSON, err := ctx.GetStub().GetState(id)
    if err != nil {
        return nil, fmt.Errorf("failed to read asset: %v", err)
    }
    if assetJSON == nil {
        return nil, fmt.Errorf("asset does not exist: %s", id)
    }

    var asset Asset
    err = json.Unmarshal(assetJSON, &asset)
    if err != nil {
        return nil, err
    }

    return &asset, nil
}

func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, id string, balance float64, status string, transAmount float64, transType string, remarks string) error {
    asset, err := s.ReadAsset(ctx, id)
    if err != nil {
        return err
    }

    asset.Balance = balance
    asset.Status = status
    asset.TransAmount = transAmount
    asset.TransType = transType
    asset.Remarks = remarks

    assetJSON, err := json.Marshal(asset)
    if err != nil {
        return err
    }

    return ctx.GetStub().PutState(id, assetJSON)
}

func (s *SmartContract) GetAssetHistory(ctx contractapi.TransactionContextInterface, id string) ([]Asset, error) {
    resultsIterator, err := ctx.GetStub().GetHistoryForKey(id)
    if err != nil {
        return nil, err
    }
    defer resultsIterator.Close()

    var assets []Asset
    for resultsIterator.HasNext() {
        response, err := resultsIterator.Next()
        if err != nil {
            return nil, err
        }

        var asset Asset
        if err := json.Unmarshal(response.Value, &asset); err != nil {
            return nil, err
        }
        assets = append(assets, asset)
    }

    return assets, nil
}

func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
    assetJSON, err := ctx.GetStub().GetState(id)
    if err != nil {
        return false, fmt.Errorf("failed to read asset: %v", err)
    }
    return assetJSON != nil, nil
}