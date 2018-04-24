package iostdb

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMember(t *testing.T) {
	Convey("Test of Member", t, func() {
		m, err := NewMember(nil)
		Convey("New member: ", func() {
			So(err, ShouldBeNil)
			So(len(m.PubKey), ShouldEqual, 33)
			So(len(m.SecKey), ShouldEqual, 32)
		})

		Convey("sign and verify: ", func() {
			info := []byte("hello world")
			sig := Sign(Sha256(info), m.SecKey)
			So(VerifySignature(Sha256(info), m.PubKey, sig), ShouldBeTrue)
		})
	})
}
