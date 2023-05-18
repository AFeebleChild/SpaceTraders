package lib

import (
	"fmt"
	"io"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
)

func NewClientFromCallsign(callsign string) (*Client, *Agent) {
	agent, err := LoadAgent(callsign)
	if err != nil {
		panic(err)
	}
	provider, err := securityprovider.NewSecurityProviderBearerToken(agent.Token)
	if err != nil {
		panic(err)
	}
	client, err := NewClient("https://api.spacetraders.io/v2/", WithRequestEditorFn(provider.Intercept))
	if err != nil {
		panic(err)
	}
	return client, agent
}

func HandleResp(resp *http.Response, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
