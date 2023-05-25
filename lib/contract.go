package lib

import (
	"context"
	"encoding/json"
)

type (
	Contracts []Contract
)

func GetContracts(c *Client) (*Contracts, error) {
	params := &GetContractsParams{}

	body, err := HandleResp(c.GetContracts(context.Background(), params))
	if err != nil {
		return nil, err
	}

	contracts := &Contracts{}
	err = json.Unmarshal(body, contracts)
	if err != nil {
		return nil, err
	}

	return contracts, nil
}
