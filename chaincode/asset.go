package chaincode

type Asset struct {
    DocType     string  `json:"docType"`
    DealerID    string  `json:"dealerId"`
    MSISDN      string  `json:"msisdn"`
    MPIN        string  `json:"mpin"`
    Balance     float64 `json:"balance"`
    Status      string  `json:"status"`
    TransAmount float64 `json:"transAmount"`
    TransType   string  `json:"transType"`
    Remarks     string  `json:"remarks"`
}