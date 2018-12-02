package batch_test

import (
	"errors"
	"testing"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/x/batch"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
)

func TestMsg(t *testing.T) {
	Convey("Test Validate", t, func() {
		msg := &mockMsg{}
		Convey("Test happy flow", func() {
			msg.On("MsgList").Return(make([]weave.Msg, 10), nil)
			So(batch.Validate(msg), ShouldBeNil)
			msg.AssertExpectations(t)
		})

		Convey("Test list too long", func() {
			msg.On("MsgList").Return(make([]weave.Msg, 11), nil)
			So(batch.Validate(msg), ShouldNotBeNil)
			msg.AssertExpectations(t)
		})

		Convey("Test error", func() {
			msg.On("MsgList").Return(make([]weave.Msg, 10), errors.New("whatever"))
			So(batch.Validate(msg), ShouldNotBeNil)
			msg.AssertExpectations(t)
		})
	})
}

var _ batch.Msg = (*mockMsg)(nil)

type mockMsg struct {
	mock.Mock
}

func (m *mockMsg) Marshal() ([]byte, error) {
	panic("implement me")
}

func (m *mockMsg) Unmarshal([]byte) error {
	panic("implement me")
}

func (m *mockMsg) Path() string {
	panic("implement me")
}

func (m *mockMsg) Validate() error {
	args := m.Mock.Called()
	return args.Error(0)
}

func (m *mockMsg) MsgList() ([]weave.Msg, error) {
	args := m.Mock.Called()
	return args.Get(0).([]weave.Msg), args.Error(1)
}
