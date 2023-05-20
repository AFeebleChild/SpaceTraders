package lib

import (
	"context"
	"encoding/json"
)

type (
	Contracts []Contract
)

func GetContracts(callsign string) (*Contracts, error) {
	client, err := NewClientFromCallsign(callsign)
	if err != nil {
		return nil, err
	}

	params := &GetContractsParams{}

	body, err := HandleResp(client.GetContracts(context.Background(), params))
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
