package evm

import (
	_ "embed"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
)

var (
	//go:embed addismya_ext_data_hashes.json
	rawAddismyaExtDataHashes []byte
	addismyaExtDataHashes    map[common.Hash]common.Hash

	//go:embed mainnet_ext_data_hashes.json
	rawMainnetExtDataHashes []byte
	mainnetExtDataHashes    map[common.Hash]common.Hash
)

func init() {
	if err := json.Unmarshal(rawAddismyaExtDataHashes, &addismyaExtDataHashes); err != nil {
		panic(err)
	}
	rawAddismyaExtDataHashes = nil
	if err := json.Unmarshal(rawMainnetExtDataHashes, &mainnetExtDataHashes); err != nil {
		panic(err)
	}
	rawMainnetExtDataHashes = nil
}
