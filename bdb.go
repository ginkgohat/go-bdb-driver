package bdb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	baseHeader http.Header
}

// NewClient Client Constructor
func NewClient(bdbURL string, client *http.Client, header http.Header) (*Client, error) {
	_, err := url.Parse(bdbURL)
	if err != nil {
		return nil, errors.Wrap(err, "Could not parse bigchaindb url")
	}

	return &Client{
		baseURL:    bdbURL,
		httpClient: client,
		baseHeader: header,
	}, nil
}

// New default bigchaindb client
// bdbURL host+port https://example.com:9984、
func New(bdbURL string) (*Client, error) {
	return NewClient(bdbURL, http.DefaultClient, http.Header{})
}

func (c *Client) get(ctx context.Context, path string, out interface{}) error {
	targetURL := fmt.Sprintf("%s/%s", c.baseURL, path)
	resp, err := c.httpClient.Get(strings.TrimSpace(targetURL))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(&out)
}

func (c *Client) post(ctx context.Context, path string, in, out interface{}) error {
	targetURL := fmt.Sprintf("%s/%s", c.baseURL, path)

	var buf io.ReadWriter
	if in != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(in)
		if err != nil {
			return err
		}
	}

	resp, err := c.httpClient.Post(targetURL, "json", buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	return json.NewDecoder(resp.Body).Decode(&out)
}

// GetNodeInfo get bigchaindb node info
func (c *Client) GetNodeInfo(ctx context.Context) (node Node, err error) {
	err = c.get(ctx, "", &node)
	return
}

// GetNodeInfo get bigchaindb root endpoint info
func (c *Client) GetEndpoint(ctx context.Context) (endpoint Endpoint, err error) {
	err = c.get(ctx, apiPath, &endpoint)
	return
}

// GetTransaction get the transaction with the ID
func (c *Client) GetTransaction(ctx context.Context, transactionID string) (transaction Transaction, err error) {
	path := fmt.Sprintf("%s/%s/%s", apiPath, transactionPath, transactionID)
	err = c.get(ctx, path, &transaction)
	return
}

// GetTransactionList Get a list of transactions that use an asset with the ID
// asset_id the asset id
// operation is CREATE or TRANSFER or ""(all)
func (c *Client) GetTransactionList(ctx context.Context, assetID, operation string, lastTx bool) (transactions []Transaction, err error) {
	path := fmt.Sprintf("%s/%s/?asset_id=%s&last_tx=%v", apiPath, transactionPath, assetID, lastTx)
	if operation == "CREATE" || operation == "TRANSFER" {
		path = fmt.Sprintf("%s&operation=%s", path, operation)
	}
	err = c.get(ctx, path, &transactions)
	return
}

// GetOutputs Get transaction outputs by public key
// status spent | unspent | ""(all)
func (c *Client) GetOutputs(ctx context.Context, publicKey string, status string) (outputs []OutputLocation, err error) {
	path := fmt.Sprintf("%s/%s/?public_key=%s", apiPath, outputPath, publicKey)
	if status == unspent {
		path = fmt.Sprintf("%s&spent=%v", path, false)
	}
	if status == spent {
		path = fmt.Sprintf("%s&spent=%v", path, true)
	}
	err = c.get(ctx, path, &outputs)
	return
}

// GetAssets get all the assets that match a given text search.
func (c *Client) GetAssets(ctx context.Context, search string, limit int) (assets []Asset, err error) {
	path := fmt.Sprintf("%s/%s/?search=%s", apiPath, assetPath, search)
	if limit > 0 {
		path = fmt.Sprintf("%s&limit=%v", path, limit)
	}
	err = c.get(ctx, path, &assets)
	return
}

// GetMetadatas get all the metadatas that match a given text search.
func (c *Client) GetMetadatas(ctx context.Context, search string, limit int) (metadatas []Metadata, err error) {
	path := fmt.Sprintf("%s/%s/?search=%s", apiPath, metadataPath, search)
	if limit > 0 {
		path = fmt.Sprintf("%s&limit=%v", path, limit)
	}
	err = c.get(ctx, path, &metadatas)
	return
}

// GetValidators the local validators set of a given node.
func (c *Client) GetValidators(ctx context.Context) (validators []Validator, err error) {
	path := fmt.Sprintf("%s/%s", apiPath, validatorPath)
	err = c.get(ctx, path, &validators)
	return
}

// GetBlock get the block with the height
func (c *Client) GetBlock(ctx context.Context, height int) (block Block, err error) {
	path := fmt.Sprintf("%s/%s/%d", apiPath, blockPath, height)
	err = c.get(ctx, path, &block)
	return
}

// GetBlockHeights  Retrieve a list of block IDs (block heights),
// such that the blocks with those IDs contain a transaction with the ID transaction_id.
// A correct response may consist of an empty list or a list with one block ID.
func (c *Client) GetBlockHeight(ctx context.Context, transactionID string) (heights []int, err error) {
	path := fmt.Sprintf("%s/%s?transaction_id=%s", apiPath, blockPath, transactionID)
	err = c.get(ctx, path, &heights)
	return
}

func (c *Client) PostTransaction(ctx context.Context, mode string, tx interface{}) (*Transaction, error) {
	var out *Transaction
	if !(mode == sync || mode == commit) {
		mode = async
	}
	path := fmt.Sprintf("%s/%s?mode=%s", apiPath, transactionPath, mode)
	err := c.post(ctx, path, tx, out)
	return out, err
}
