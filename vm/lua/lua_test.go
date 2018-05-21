package lua

import (
	"testing"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"fmt"
)

func TestLuaVM(t *testing.T) {
	Convey("Test of Lua VM", t, func() {
		Convey("Normal", func() {
			mockCtl := gomock.NewController(t)
			pool := core_mock.NewMockPool(mockCtl)

			var k state.Key
			var v state.Value

			pool.EXPECT().Put(gomock.Any(), gomock.Any()).AnyTimes().Do(func(key state.Key, value state.Value) error {
				k = key
				v = value
				return nil
			})
			pool.EXPECT().Copy().AnyTimes().Return(pool)
			main := NewMethod("main", 0, 1)
			lc := Contract{
				info: vm.ContractInfo{Prefix: "test", GasLimit: 11},
				code: `function main()
	Put("hello", "world")
	return "success"
end`,
				main: main,
			}

			lvm := VM{}
			lvm.Prepare(&lc, nil)
			lvm.Start()
			ret, _, err := lvm.Call(pool, "main")
			lvm.Stop()
			So(err, ShouldBeNil)
			So(ret[0].EncodeString(), ShouldEqual, "ssuccess")
			So(k, ShouldEqual, "testhello")
			So(v.EncodeString(), ShouldEqual, "sworld")

		})

		Convey("Transfer", func() {
			mockCtl := gomock.NewController(t)
			pool := core_mock.NewMockPool(mockCtl)

			var va, vb state.Value
			var eeee error

			pool.EXPECT().GetHM(gomock.Any(), gomock.Any()).AnyTimes().Return(state.MakeVFloat(10000), nil)

			pool.EXPECT().PutHM(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Do(func(key, field state.Key, value state.Value) error {
				if key == "iost" {
					if field == "a" {
						va = value
						return nil
					} else if field == "b" {
						vb = value
						return nil
					}
				}
				eeee = fmt.Errorf("bad")
				return nil
			})

	})

}
