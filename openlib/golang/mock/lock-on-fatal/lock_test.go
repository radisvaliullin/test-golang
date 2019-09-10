package mocklock

import (
	"testing"

	"github.com/golang/mock/gomock"
	mocklock_mock "github.com/radisvaliullin/test-golang/v2/openlib/golang/mock/lock-on-fatal/mocks"
)

func TestSome_Do(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocklock_mock.NewMockDo(ctrl)
	mr := mocklock_mock.NewMockRun(ctrl)

	// m.EXPECT().Do(gomock.Any).Return(4).AnyTimes()

	d := &Doer{}
	s := Some{doer: d}
	// i := s.Do2(2)
	// if i != 4 {
	// 	t.Fatal("wrong Do2 result")
	// }
	// t.Logf("do2 return %v", i)

	// //
	s.doer = m
	// i = s.Do3(2)
	// if i != 8 {
	// 	t.Fatal("wrong Do3 result")
	// }
	// t.Logf("do3 return %v", i)

	//
	s.runner = mr
	// r := &Runner{}
	// s.runner = r
	ch := make(chan int)
	go func() {
		i := s.Run(3)
		ch <- i
	}()
	i := <-ch
	if i != 81 {
		t.Fatal("wrong run result")
	}
	t.Logf("run return %v", i)
}
