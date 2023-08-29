package neodb_go_client

import (
	"encoding/json"
	"errors"
	"fmt"
)

type ShelfTypeRequest struct {
	ShelfType ShelfType `json:"shelf_ty"`
	Category  Category  `json:"category,omitempty"`
	Page      int       `json:"page"`
}

const ShelfTypeUrl = NeoDBUrl + "/api/me/shelf/%s?page=%d"

func (n *NeoClient) GetMyShelfByTypeAndCategory(req *ShelfTypeRequest) (*PagedMark, error) {
	if req.ShelfType == "" {
		return nil, NoShelfTypeErr
	}
	if req.Page == 0 {
		req.Page = 1
	}
	url := fmt.Sprintf(ShelfTypeUrl, req.ShelfType, req.Page)
	if req.Category != "" {
		url = fmt.Sprintf("%s&category=%s", url, req.Category)
	}

	response, err := n.get(url)
	if err != nil {
		return nil, err
	}
	result := &PagedMark{}
	err = json.Unmarshal(response, result)
	if err != nil {
		return nil, errors.Join(InternalErr, err)
	}
	for _, datum := range result.Data {
		t, err := parseDate(datum.CreatedTimeStr)
		if err != nil {
			return nil, errors.Join(InternalErr, err)
		}
		datum.CreatedTime = t
	}
	return result, nil
}
