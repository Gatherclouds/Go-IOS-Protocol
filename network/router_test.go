package network

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"

	"io/ioutil"
	"os"
	"strconv"
	"time"
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

func TestGetInstance(t *testing.T) {
	Convey("", t, func() {

		router, err := GetInstance(&NetConifg{NodeTablePath: "tale_test", ListenAddr: "127.0.0.1"}, "base", 30304)

		So(err, ShouldBeNil)
		So(router.(*RouterImpl).port, ShouldEqual, uint16(30304))
		So(Route.(*RouterImpl).port, ShouldEqual, uint16(30304))
	})
}

func initNetConf() *NetConifg {
	conf := &NetConifg{}
	conf.SetLogPath("iost_log_")
	tablePath, _ := ioutil.TempDir(os.TempDir(), "iost_node_table_"+strconv.Itoa(int(time.Now().UnixNano())))
	conf.SetNodeTablePath(tablePath)
	conf.SetListenAddr("0.0.0.0")
	return conf
}
