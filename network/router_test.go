package network

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"

)

func TestRouterImpl_Init(t *testing.T) {
	//broadcast(t)
	router, _ := RouterFactory("base")
	baseNet, _ := NewBaseNetwork(&NetConifg{ListenAddr: "0.0.0.0"})
	router.Init(baseNet, 30601)
	Convey("init", t, func() {
		So(router.(*RouterImpl).port, ShouldEqual, 30601)
	})
}
