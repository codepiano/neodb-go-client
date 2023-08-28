package neodb_go_client

import (
	"fmt"
	"testing"
)

func TestNeoClient_GetMyShelfByTypeAndCategory(t *testing.T) {
	client, _ := NewNeoClient(TestAccessToken)
	type args struct {
		req *ShelfTypeRequest
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "book first page",
			args: args{
				&ShelfTypeRequest{
					ShelfType: ShelfTypeComplete,
					Category:  CategoryBook,
					Page:      1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetMyShelfByTypeAndCategory(tt.args.req)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(got)
		})
	}
}
