package lib

import (
	"encoding/json"
	"fmt"
	"os"
)

type Token struct {
	Token  string `json:"token"`
	Symbol string `json:"symbol"`
}

func NewToken(token string, symbol string) *Token {
	return &Token{
		Token:  token,
		Symbol: symbol,
	}
}

func (t *Token) Save() error {
	path := "agents/" + t.Symbol + "/token.json"

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("cannot create token file: %v", err)
	}
	defer file.Close()

	d, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot marshal token: %v", err)
	}
	_, err = file.Write(d)
	if err != nil {
		return fmt.Errorf("cannot write token to file: %v", err)
	}

	return nil
}

func LoadToken(symbol string) (*Token, error) {
	path := "agents/" + symbol + "/token.json"

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open token file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	token := &Token{}
	err = decoder.Decode(token)

	if err != nil {
		return nil, fmt.Errorf("cannot get token from file: %v", err)
	}

	return token, nil
}
