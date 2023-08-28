package neodb_go_client

import (
	"fmt"
	"testing"
)

func TestNeoClient_GetBookByUUID(t *testing.T) {
	client, _ := NewNeoClient(TestAccessToken)
	type args struct {
		uuid string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "黄金时代 - 王小波",
			args: args{"6zaED9nnuPMxGwsFCUMH0V"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetBookByUUID(tt.args.uuid)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(got)
		})
	}
}
