package bdb

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

const baseURL = "http://geekpi.top:15522" // http://localhost:9984

var (
	httpClient = http.DefaultClient
	httpHeader = http.Header{}
)

func TestNew(t *testing.T) {
	type args struct {
		bdbURL string
	}
	tests := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
	}{
		{name: "case1", args: args{baseURL}, want: nil, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.args.bdbURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_GetNodeInfo(t *testing.T) {
	type fields struct {
		baseURL    string
		httpClient *http.Client
		baseHeader http.Header
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "case1", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL:    tt.fields.baseURL,
				httpClient: tt.fields.httpClient,
				baseHeader: tt.fields.baseHeader,
			}
			got, err := c.GetNodeInfo(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNodeInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			val, _ := json.Marshal(got)
			t.Log(string(val))
		})
	}
}

func TestClient_GetEndpoint(t *testing.T) {
	type fields struct {
		baseURL    string
		httpClient *http.Client
		baseHeader http.Header
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "case1", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background()}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL:    tt.fields.baseURL,
				httpClient: tt.fields.httpClient,
				baseHeader: tt.fields.baseHeader,
			}
			got, err := c.GetEndpoint(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEndpoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			val, _ := json.Marshal(got)
			t.Log(string(val))
		})
	}
}

func TestClient_GetTransaction(t *testing.T) {
	type fields struct {
		baseURL    string
		httpClient *http.Client
		baseHeader http.Header
	}
	type args struct {
		ctx context.Context
		tid string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		// wantTransaction types.Transaction
		wantErr bool
	}{
		{name: "case1", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background(), "d36547a9a5f3b5a88a5bb97943f4074e23d2bebb494e603268209b8af121e127"}, wantErr: false},
		{name: "case2", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background(), "18da0b91c0d638c522b2cf58e2bc9a2195069694b27b3ded40f2e916225c1b21"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL:    tt.fields.baseURL,
				httpClient: tt.fields.httpClient,
				baseHeader: tt.fields.baseHeader,
			}
			got, err := c.GetTransaction(tt.args.ctx, tt.args.tid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			val, _ := json.Marshal(got)
			t.Log(string(val))
		})
	}
}

func TestClient_GetTransactionList(t *testing.T) {
	type fields struct {
		baseURL    string
		httpClient *http.Client
		baseHeader http.Header
	}
	type args struct {
		ctx       context.Context
		assetID   string
		operation string
		lastTx    bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		// wantTransactions []types.Transaction
		wantErr bool
	}{
		{name: "case1", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background(), "18da0b91c0d638c522b2cf58e2bc9a2195069694b27b3ded40f2e916225c1b21", "TRANSFER", false}, wantErr: false},
		{name: "case2", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background(), "18da0b91c0d638c522b2cf58e2bc9a2195069694b27b3ded40f2e916225c1b21", "TRANSFER", true}, wantErr: false},
		{name: "case3", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background(), "18da0b91c0d638c522b2cf58e2bc9a2195069694b27b3ded40f2e916225c1b21", "CREATE", false}, wantErr: false},
		{name: "case4", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background(), "18da0b91c0d638c522b2cf58e2bc9a2195069694b27b3ded40f2e916225c1b21", "", false}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL:    tt.fields.baseURL,
				httpClient: tt.fields.httpClient,
				baseHeader: tt.fields.baseHeader,
			}
			gotTransactions, err := c.GetTransactionList(tt.args.ctx, tt.args.assetID, tt.args.operation, tt.args.lastTx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTransactionList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(gotTransactions, tt.wantTransactions) {
			//	t.Errorf("GetTransactionList() gotTransactions = %v, want %v", gotTransactions, tt.wantTransactions)
			//}

			val, _ := json.Marshal(gotTransactions)
			t.Log(string(val))
		})
	}
}

func TestClient_GetOutputs(t *testing.T) {
	type fields struct {
		baseURL    string
		httpClient *http.Client
		baseHeader http.Header
	}
	type args struct {
		ctx       context.Context
		publicKey string
		status    string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		// wantOutputs []types.OutputLocation
		wantErr bool
	}{
		{name: "case1", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background(), "2LeKKDqhnYvHx8oXDE4PxXpg31NM7d2NgwXXCAxasuuv", spent}, wantErr: false},
		{name: "case2", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background(), "2LeKKDqhnYvHx8oXDE4PxXpg31NM7d2NgwXXCAxasuuv", unspent}, wantErr: false},
		{name: "case3", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background(), "2LeKKDqhnYvHx8oXDE4PxXpg31NM7d2NgwXXCAxasuuv", ""}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL:    tt.fields.baseURL,
				httpClient: tt.fields.httpClient,
				baseHeader: tt.fields.baseHeader,
			}
			gotOutputs, err := c.GetOutputs(tt.args.ctx, tt.args.publicKey, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOutputs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(gotOutputs, tt.wantOutputs) {
			//	t.Errorf("GetOutputs() gotOutputs = %v, want %v", gotOutputs, tt.wantOutputs)
			//}
			val, _ := json.Marshal(gotOutputs)
			t.Log(string(val))
		})
	}
}

func TestClient_GetAssets(t *testing.T) {
	type fields struct {
		baseURL    string
		httpClient *http.Client
		baseHeader http.Header
	}
	type args struct {
		ctx    context.Context
		search string
		limit  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		// wantAssets []types.Asset
		wantErr bool
	}{
		{name: "case1", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background(), "BigchainDB", 0}, wantErr: false},
		{name: "case2", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background(), "BigchainDB", 2}, wantErr: false},
		{name: "case3", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background(), "BigchainDB", 10}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL:    tt.fields.baseURL,
				httpClient: tt.fields.httpClient,
				baseHeader: tt.fields.baseHeader,
			}
			gotAssets, err := c.GetAssets(tt.args.ctx, tt.args.search, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAssets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(gotAssets, tt.wantAssets) {
			//	t.Errorf("GetAssets() gotAssets = %v, want %v", gotAssets, tt.wantAssets)
			//}

			val, _ := json.Marshal(gotAssets)
			t.Log(string(val))
		})
	}
}

func TestClient_GetMetadatas(t *testing.T) {
	type fields struct {
		baseURL    string
		httpClient *http.Client
		baseHeader http.Header
	}
	type args struct {
		ctx    context.Context
		search string
		limit  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		// wantMetadatas []types.Metadata
		wantErr bool
	}{
		{name: "case1", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background(), "1201", 0}, wantErr: false},
		{name: "case2", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background(), "bigchaindb", 2}, wantErr: false},
		{name: "case3", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background(), "仁爱路", 10}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL:    tt.fields.baseURL,
				httpClient: tt.fields.httpClient,
				baseHeader: tt.fields.baseHeader,
			}
			gotMetadatas, err := c.GetMetadatas(tt.args.ctx, tt.args.search, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMetadatas() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(gotMetadatas, tt.wantMetadatas) {
			//	t.Errorf("GetMetadatas() gotMetadatas = %v, want %v", gotMetadatas, tt.wantMetadatas)
			//}
			val, _ := json.Marshal(gotMetadatas)
			t.Log(string(val))
		})
	}
}

func TestClient_GetValidators(t *testing.T) {
	type fields struct {
		baseURL    string
		httpClient *http.Client
		baseHeader http.Header
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		// wantValidators []types.Validator
		wantErr bool
	}{
		{name: "case1", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background()}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL:    tt.fields.baseURL,
				httpClient: tt.fields.httpClient,
				baseHeader: tt.fields.baseHeader,
			}
			gotValidators, err := c.GetValidators(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetValidators() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(gotValidators, tt.wantValidators) {
			//	t.Errorf("GetValidators() gotValidators = %v, want %v", gotValidators, tt.wantValidators)
			//}
			val, _ := json.Marshal(gotValidators)
			t.Log(string(val))
		})
	}
}

func TestClient_GetBlock(t *testing.T) {
	type fields struct {
		baseURL    string
		httpClient *http.Client
		baseHeader http.Header
	}
	type args struct {
		ctx    context.Context
		height int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		// wantBlock types.Block
		wantErr bool
	}{
		{name: "case1", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background(), 1}, wantErr: false},
		{name: "case2", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background(), 2}, wantErr: false},
		{name: "case3", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background(), 4}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL:    tt.fields.baseURL,
				httpClient: tt.fields.httpClient,
				baseHeader: tt.fields.baseHeader,
			}
			gotBlock, err := c.GetBlock(tt.args.ctx, tt.args.height)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(gotBlock, tt.wantBlock) {
			//	t.Errorf("GetBlock() gotBlock = %v, want %v", gotBlock, tt.wantBlock)
			//}
			val, _ := json.Marshal(gotBlock)
			t.Log(string(val))
		})
	}
}

func TestClient_GetBlockHeight(t *testing.T) {
	type fields struct {
		baseURL    string
		httpClient *http.Client
		baseHeader http.Header
	}
	type args struct {
		ctx           context.Context
		transactionID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		// wantHeights []int
		wantErr bool
	}{
		{name: "case1", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background(), "d36547a9a5f3b5a88a5bb97943f4074e23d2bebb494e603268209b8af121e127"}, wantErr: false},
		{name: "case2", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background(), "4e25fcade9172b3dc2310d7565150d4d600247cd24797c2de8e48c2cc7fb8f19"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL:    tt.fields.baseURL,
				httpClient: tt.fields.httpClient,
				baseHeader: tt.fields.baseHeader,
			}
			gotHeight, err := c.GetBlockHeight(tt.args.ctx, tt.args.transactionID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBlockHeights() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(gotHeights, tt.wantHeights) {
			//	t.Errorf("GetBlockHeights() gotHeights = %v, want %v", gotHeights, tt.wantHeights)
			//}
			t.Log(gotHeight)
		})
	}
}

func TestClient_NewKeyPair(t *testing.T) {
	type fields struct {
		baseURL    string
		httpClient *http.Client
		baseHeader http.Header
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		// wantKeys *types.KeyPair
		wantErr bool
	}{
		{name: "case1", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background()}, wantErr: false},
		{name: "case2", fields: fields{baseURL: baseURL, httpClient: httpClient, baseHeader: httpHeader}, args: args{context.Background()}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				baseURL:    tt.fields.baseURL,
				httpClient: tt.fields.httpClient,
				baseHeader: tt.fields.baseHeader,
			}
			gotKeys, err := c.NewKeyPair(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewKeyPair() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(gotKeys, tt.wantKeys) {
			//	t.Errorf("NewKeyPair() gotKeys = %v, want %v", gotKeys, tt.wantKeys)
			//}
			val, _ := json.Marshal(gotKeys)
			t.Log(string(val))
		})
	}
}
