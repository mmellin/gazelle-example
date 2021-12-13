package myapp

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	examplepb "example.com/gazelle-example/myapp/proto/v1"
	example_mock "example.com/gazelle-example/myapp/proto/v1/mock"
)

func TestMyService(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	exampleClient := example_mock.NewMockMyServiceClient(mockController)
	ctx := context.Background()
	req := &examplepb.GetRequest{UserId: 1}
	want := &examplepb.Mandate{DataId: 2}
	exampleClient.EXPECT().GetNewData(ctx, req).Return(want, nil)
	got, _ := exampleClient.GetNewData(ctx, req)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("TestMyService failed: \ngot %v, \nwant %v", got, want)
	}
}
