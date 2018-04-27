package iostdb

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSign(t *testing.T) {
	testData := "c6e193266883a500c6e51a117e012d96ad113d5f21f42b28eb648be92a78f92f"
	privKey := ParseHex(testData)
	var pubKey []byte

	Convey("Test of Crypto", t, func() {
		Convey("Sha256", func() {
			sha := "d4daf0546cb71d90688b45488a8fa000b0821ec14b73677b2fb7788739228c8b"
			So(ToHex(Sha256(privKey)), ShouldEqual, sha)
		})

		Convey("Calculate public key", func() {
			pub := "0314bf901a6640033ea07b39c6b3acb675fc0af6a6ab526f378216085a93e5c7a2"
			pubKey = CalcPubKey(privKey)
			So(ToHex(pubKey), ShouldEqual, pub)
		})

		Convey("Hash-160", func() {
			hash := "9c1185a5c5e9fc54612808977ee8f548b2258d31"
			So(ToHex(Hash160(CalcPubkey(privKey))), ShouldEqual, hash)
		})

	})
}
