package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
)

const ApiUrl = "https://api.spacetraders.io/v2/"

type Data struct {
	Data interface{} `json:"data"`
}

func NewClientFromCallsign(callsign string) (*Client, error) {
	token, err := LoadToken(callsign)
	if err != nil {
		return nil, err
	}
	provider, err := securityprovider.NewSecurityProviderBearerToken(token.Token)
	if err != nil {
		return nil, err
	}
	client, err := NewClient(ApiUrl, WithRequestEditorFn(provider.Intercept))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func HandleResp(resp *http.Response, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		JsonPrettyPrint(body)
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	// All returned data (that is not an error) is wrapped in a "data" object.
	// Doing a generic unmarshal to a struct with an interface{} removes the need
	// to create a struct for every data type.
	d := Data{}
	err = json.Unmarshal(body, &d)
	if err != nil {
		return nil, err
	}
	nd, err := json.Marshal(d.Data)
	if err != nil {
		return nil, err
	}

	return nd, nil
}
