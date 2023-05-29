package lib

import (
	"context"
	"encoding/json"
	"fmt"
)

type (
	NewAgentResp struct {
		Token    string   `json:"token"`
		Agent    Agent    `json:"agent"`
		Ship     Ship     `json:"ship"`
		Contract Contract `json:"contract"`
		Faction  Faction  `json:"faction"`
	}
)

func (a Agent) Save() error {
	path := MakeJsonPath("agents", a.Symbol)
	fmt.Println("Saving agent:", path)

	return JsonFilePrettyPrint(path, a)
}

func LoadAgent(symbol string) (*Agent, error) {
	path := MakeJsonPath("agents", symbol)
	agent := &Agent{}
	err := JsonReadFile(path, agent)
	return agent, err
}

func NewAgent(client *Client, symbol, faction string) (*NewAgentResp, error) {
	params := RegisterJSONRequestBody{
		Symbol:  symbol,
		Faction: &faction,
	}

	body, err := HandleResp(client.Register(context.Background(), params))
	if err != nil {
		return nil, err
	}

	newAgent := &NewAgentResp{}
	err = json.Unmarshal(body, newAgent)
	if err != nil {
		return nil, err
	}

	return newAgent, nil
}

func GetAgent(c *Client) (*Agent, error) {
	body, err := HandleResp(c.GetMyAgent(context.Background()))
	if err != nil {
		return nil, err
	}

	agent := &Agent{}
	err = json.Unmarshal(body, agent)
	if err != nil {
		return nil, err
	}

	return agent, nil
}
