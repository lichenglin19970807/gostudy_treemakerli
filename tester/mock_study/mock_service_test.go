package tester

import (
	"github.com/golang/mock/gomock"
	"gostudy_treemakerli/mocks"
	"testing"
)

var req int = 1

func TestMockService(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockService := mocks.NewMockMockService(controller)
	mockService.EXPECT().MethodOne(1).Return("a", nil)
	//mockService.EXPECT().MethodOne(2).Return("b", errors.New("bbb"))
	if resp, err := mockService.MethodOne(req); err != nil {
		t.Fatalf("1: %v, %v\n", resp, err)
	}
}
