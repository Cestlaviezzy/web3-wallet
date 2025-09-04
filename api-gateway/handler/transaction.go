package handler

import (
	"net/http"

	"web3-wallet/api-gateway/client"
	"web3-wallet/api-gateway/types"
)

// TransactionHandler 交易处理器
type TransactionHandler struct {
	clients *client.GRPCClients
}

// NewTransactionHandler 创建交易处理器
func NewTransactionHandler(clients *client.GRPCClients) *TransactionHandler {
	return &TransactionHandler{
		clients: clients,
	}
}

// GetTxReceipt 查询交易回执
func (h *TransactionHandler) GetTxReceipt(w http.ResponseWriter, r *http.Request) {
	txHash := extractPathParam(r.URL.Path, "/tx_receipt/")
	if txHash == "" {
		writeJSONResponse(w, http.StatusOK, types.Error(types.StatusBadRequest, "Missing tx_hash parameter"))
		return
	}

	// TODO: 调用chain-service的gRPC接口
	response := map[string]interface{}{
		"txHash":  txHash,
		"status":  1,
		"receipt": nil,
		"txInfo":  map[string]interface{}{},
	}

	writeJSONResponse(w, http.StatusOK, types.Success(response))
}

// GetSubmitters 获取提交者状态 (兼容MetaSignX-Wallet响应格式)
func (h *TransactionHandler) GetSubmitters(w http.ResponseWriter, r *http.Request) {
	// TODO: 调用transaction-service获取提交者状态
	response := types.SubmittersResponse{
		Status: []types.SubmitterStatus{
			{
				Address:              "0x6813eb9362372eef6200f3b1dbc3f819671cba69",
				Nonce:                72,
				Balance:              "0",
				IsSendingTransaction: false,
				IsBlocking:           false,
			},
		},
		FeeReceiver: "0x6506b82cd478c27434bcd7152c0a58bef5ed23e9",
	}

	writeJSONResponse(w, http.StatusOK, types.Success(response))
}
