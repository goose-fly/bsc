package downstream

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type Block struct {
	Number           *big.Int `json:"_id"`
	Hash             string   `json:"blockHash"`
	ParentHash       string   `json:"parentHash"`
	Nonce            string   `json:"nonce"`
	Sha3Uncles       string   `json:"sha3Uncles"`
	LogsBloom        string   `json:"logsBloom"`
	TransactionsRoot string   `json:"transactionsRoot"`
	StateRoot        string   `json:"stateRoot"`
	ReceiptRoot      string   `json:"receiptsRoot"`
	Miner            string   `json:"miner"`
	MixHash          string   `json:"mixHash"`
	Difficulty       string   `json:"difficulty"`
	TotalDifficulty  string   `json:"totalDifficulty"`
	ExtraData        string   `json:"extraData"`
	BlockSize        uint64   `json:"blockSize"`
	GasLimit         uint64   `json:"gasLimit"`
	GasUsed          uint64   `json:"gasUsed"`
	BlockTimestamp   uint64   `json:"blockTimestamp"`
	TransactionCount int      `json:"transactionCount"`
	Transactions     []*Tx    `json:"transactions"`
	Class            string   `json:"_class"`
}

func FromOrgBlock(oBlock *types.Block) *Block {
	oHeader := oBlock.Header()

	return &Block{
		Number:           oHeader.Number,
		Hash:             oHeader.Hash().Hex(),
		ParentHash:       oHeader.ParentHash.Hex(),
		Nonce:            fmt.Sprintf("0x%x", oHeader.Nonce),
		Sha3Uncles:       oHeader.UncleHash.Hex(),
		LogsBloom:        hexutil.Encode(hexutil.Bytes(oHeader.Bloom[:])),
		TransactionsRoot: oHeader.TxHash.Hex(),
		StateRoot:        oHeader.Root.Hex(),
		ReceiptRoot:      oHeader.ReceiptHash.Hex(),
		Miner:            strings.ToLower(oHeader.Coinbase.Hex()),
		MixHash:          oHeader.MixDigest.Hex(),
		Difficulty:       fmt.Sprintf("0x%x", oHeader.Difficulty),
		TotalDifficulty:  "0x0",
		ExtraData:        hexutil.Encode(oHeader.Extra),
		BlockSize:        uint64(oBlock.Size()),
		GasLimit:         oHeader.GasLimit,
		GasUsed:          oHeader.GasUsed,
		BlockTimestamp:   oHeader.Time,
		TransactionCount: 0,
		Transactions:     make([]*Tx, 0),
		Class:            "com.mingsi.data.connector.entity.Block",
	}
}
