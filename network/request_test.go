package network

import (
	"testing"
	"time"
	"bytes"
	"Go-IOS-Protocol/common"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRequest_Unpack(t *testing.T) {
	tim := time.Now().UnixNano()
	req := newRequest(Message, "0.0.0.0", common.Int64ToBytes(tim))

	Convey("test unpack packet splicing", t, func() {
		testData, err := req.Pack()
		So(err, ShouldEqual, nil)
		buf := new(bytes.Buffer)
		buf.Write(testData)
		buf.Write(testData)
		buf.Write(testData)

		readerCh := make(chan Request, 3)
		// scanner
		reader(buf, readerCh)
		i := 0
		for {
			select {
			case req := <-readerCh:
				if len(req.Body) > 0 {
					So(common.BytesToInt64(req.Body), ShouldEqual, tim)
					i++
				}
				if i == 3 {
					return
				}
			case <-time.After(1 * time.Second):
				So("timeout", ShouldEqual, "")
				break

			}
		}

	})
}
