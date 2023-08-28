package neodb_go_client

import (
	"time"
)

type ExternalResource struct {
	Url string `json:"url"`
}

type Item struct {
	UUID              string              `json:"uuid"`
	Url               string              `json:"url"`
	ApiUrl            string              `json:"api_url"`
	Category          Category            `json:"category"`
	ParentUUID        string              `json:"parent_uuid"`
	DisplayTitle      string              `json:"display_title"`
	ExternalResources []*ExternalResource `json:"external_resources"`
	Title             string              `json:"title"`
	Brief             string              `json:"brief"`
	CoverImageUrl     string              `json:"cover_image_url"`
	Rating            float32             `json:"rating"`
	RatingCount       int                 `json:"rating_count"`
}

type Mark struct {
	ShelfType      ShelfType  `json:"shelf_type"`
	Visibility     int        `json:"visibility"`
	Item           *Item      `json:"item"`
	CreatedTimeStr string     `json:"created_time"`
	CreatedTime    *time.Time `json:"-"`
	CommentText    string     `json:"comment_text"`
	RatingGrade    int        `json:"rating_grade"`
	Tags           []string   `json:"tags"`
}

type PagedMark struct {
	Data  []*Mark `json:"data"`
	Pages int     `json:"pages"`
	Count int     `json:"count"`
}
