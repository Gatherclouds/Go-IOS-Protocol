package core

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMember(t *testing.T) {
	Convey("Test of Member", t, func() {
		m, err := NewMember(nil)
		Convey("New member: ", func() {
			So(err, ShouldBeNil)
			So(len(m.Pubkey), ShouldEqual, 33)
			So(len(m.Seckey), ShouldEqual, 32)
		})
	})
}
