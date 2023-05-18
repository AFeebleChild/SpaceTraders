package lib

import (
	"encoding/json"
	"fmt"
	"os"
)

func (a Agent) Save() error {
	path := "agents/" + a.Symbol + ".json"
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

	return err
}

func LoadAgent(symbol string) (*Agent, error) {
	file, err := os.Open("agents/" + symbol + ".json")
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

	return agent, err
}
