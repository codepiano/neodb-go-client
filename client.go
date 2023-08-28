package neodb_go_client

import (
	"errors"
	"io"
	"net/http"
	"net/url"
)

type NeoClient struct {
	client      *http.Client
	accessToken string
}

func NewNeoClient(token string) (*NeoClient, error) {
	if token == "" {
		return nil, NeoDBTokenErr
	}
	return &NeoClient{
		client:      &http.Client{},
		accessToken: token,
	}, nil
}

func (n *NeoClient) get(reqUrl string) ([]byte, error) {
	apiUrl, err := url.Parse(reqUrl)
	if err != nil {
		return nil, errors.Join(InternalErr, err)
	}
	return n.request(&http.Request{
		Method: http.MethodGet,
		URL:    apiUrl,
	})
}

func (n *NeoClient) request(r *http.Request) ([]byte, error) {
	if r.Header == nil {
		r.Header = http.Header{}
	}
	r.Header.Set("Authorization", "Bearer "+n.accessToken)
	response, err := n.client.Do(r)
	if err != nil {
		return nil, errors.Join(NeoDBAPIRequestErr, err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Join(NeoDBAPIRequestErr, err)
	}
	if response.StatusCode != http.StatusOK {
		return nil, NeoDBAPIRequestErr
	}
	return body, nil
}
