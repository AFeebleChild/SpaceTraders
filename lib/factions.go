package lib

import (
	"context"
	"encoding/json"
)

type (
	Factions []Faction
)

func GetFactions(c *Client) (*Factions, error) {
	params := &GetFactionsParams{}

	body, err := HandleResp(c.GetFactions(context.Background(), params))
	if err != nil {
		return nil, err
	}

	factions := &Factions{}
	err = json.Unmarshal(body, factions)
	if err != nil {
		return nil, err
	}

	return factions, nil
}
