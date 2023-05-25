package lib

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
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
	path := "agents/" + a.Symbol + "/" + a.Symbol + ".json"
	fmt.Println("Saving agent:", path)
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("cannot create agent file: %v", err)
	}
	defer file.Close()

	d, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot marshal agent: %v", err)
	}
	_, err = file.Write(d)
	if err != nil {
		return fmt.Errorf("cannot write agent to file: %v", err)
	}

	return nil
}

func LoadAgent(symbol string) (*Agent, error) {
	path := "agents/" + symbol + "/" + symbol + ".json"
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open agent file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	agent := &Agent{}
	err = decoder.Decode(agent)

	if err != nil {
		return nil, fmt.Errorf("cannot get agent from file: %v", err)
	}

	return agent, nil
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
