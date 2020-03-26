package bdb

import (
	"context"
	"encoding/json"
	"testing"
)

const (
	BdbURL = "http://192.168.10.252:9984"

	// Transaction ID
	TransactionID = "02179f9cd757c06203b1da0abb012e537d034999b1baffdf02f63033c1f386d5"
	AssetID       = "02179f9cd757c06203b1da0abb012e537d034999b1baffdf02f63033c1f386d5"

	PubKey = "8Ztf3wLbeqWR4mP35mowMxGhhUMx3QrPv1DRjTUczgHo"
)

func TestClient(t *testing.T) {
	// test client create
	client, err := New(BdbURL)
	if err != nil {
		t.Log(err)
		return
	}

	// get node info
	node, err := client.GetNodeInfo(context.Background())
	if err != nil {
		t.Log(err)
		return
	}
	val, _ := json.Marshal(node)
	t.Logf("node info:%s", string(val))

	// get node info
	end, err := client.GetEndpoint(context.Background())
	if err != nil {
		t.Log(err)
		return
	}
	val, _ = json.Marshal(end)
	t.Logf("end point:%s", string(val))

	// get transaction
	tx, err := client.GetTransaction(context.Background(), TransactionID)
	if err != nil {
		t.Log(err)
		return
	}
	val, _ = json.Marshal(tx)
	t.Logf("tx:%s", string(val))

	// get transaction list
	txs, err := client.GetTransactionList(context.Background(), AssetID, create, false)
	if err != nil {
		t.Log(err)
		return
	}
	val, _ = json.Marshal(txs)
	t.Logf("tx list:%s", string(val))

	// get output
	outputs, err := client.GetOutputs(context.Background(), PubKey, unspent)
	if err != nil {
		t.Log(err)
		return
	}
	val, _ = json.Marshal(outputs)
	t.Logf("get outputs:%s", string(val))

	// get assets
	assets, err := client.GetAssets(context.Background(), "BigchainDB", 10)
	if err != nil {
		t.Log(err)
		return
	}
	val, _ = json.Marshal(assets)
	t.Logf("get assets:%s", string(val))

	// get metadatas
	metas, err := client.GetMetadatas(context.Background(), "2020-03-25", 10)
	if err != nil {
		t.Log(err)
		return
	}
	val, _ = json.Marshal(metas)
	t.Logf("get metadatas:%s", string(val))

	// get validators
	vs, err := client.GetValidators(context.Background())
	if err != nil {
		t.Log(err)
		return
	}
	val, _ = json.Marshal(vs)
	t.Logf("get validators:%s", string(val))

	// get block
	block, err := client.GetBlock(context.Background(), 1)
	if err != nil {
		t.Log(err)
		return
	}
	val, _ = json.Marshal(block)
	t.Logf("get block:%s", string(val))

	// get block height
	height, err := client.GetBlockHeight(context.Background(), TransactionID)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("get block height:%d", height)
}

func TestPostTransaction(t *testing.T) {
	txStr := `{"inputs":[{"owners_before":["N6Bgg1eF2wZoqwoVz8dNmuk5ZxC68QfNiuNH9NurbUa"],"fulfills":null,"fulfillment":"pGSAIAVnDw0BhLSI4KR-Ao3L2ET6f00tN8iQ98jdT25IEJcPgUAJ2bvw3jl7uS1iYMLfjyvJ5QrNC5DrxErpaw4ProFcprRLGzDEgVYLA1xE1k61_AZ-Am4UwtVbtY4pOcpyuz8P"}],"outputs":[{"public_keys":["N6Bgg1eF2wZoqwoVz8dNmuk5ZxC68QfNiuNH9NurbUa"],"condition":{"details":{"type":"ed25519-sha-256","public_key":"N6Bgg1eF2wZoqwoVz8dNmuk5ZxC68QfNiuNH9NurbUa"},"uri":"ni:///sha-256;mmKx6Gzcrshxg-zqi2Q-3P704j6R1akp4NYy_4JRUPA?fpt=ed25519-sha-256&cost=131072"},"amount":"1"}],"operation":"CREATE","metadata":{"date":"2020-03-25"},"asset":{"data":{"name":"hello Bigchaindb"}},"version":"2.0","id":"6c5ad1be3f9ba6d24a7ddeff55ec1dd61cecdca5e80d7ed77823a1fdbb47ddae"}`
	// test client create
	client, err := New(BdbURL)
	if err != nil {
		t.Log(err)
		return
	}
	var tx Transaction
	err = json.Unmarshal([]byte(txStr), &tx)
	if err != nil {
		t.Log(err)
		return
	}

	result, err := client.PostTransaction(context.Background(), "", tx)
	if err != nil {
		t.Log(err)
		return
	}

	val, _ := json.Marshal(result)
	t.Logf("get block:%s", string(val))
}
