package neodb_go_client

import (
	"encoding/json"
	"errors"
	"fmt"
)

const BookDetailUrl = NeoDBUrl + "/api/book/%s"

func (n *NeoClient) GetBookByUUID(uuid string) (*Edition, error) {
	url := fmt.Sprintf(BookDetailUrl, uuid)
	response, err := n.get(url)
	if err != nil {
		return nil, err
	}
	result := &Edition{}
	err = json.Unmarshal(response, result)
	if err != nil {
		return nil, errors.Join(InternalErr, err)
	}
	return result, nil
}
