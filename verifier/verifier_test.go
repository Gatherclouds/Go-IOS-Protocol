package verifier

import (
	"testing"
	"github.com/golang/mock/gomock"
)

func TestGenesisVerify(t *testing.T) {
	Convey("Test of Genesis verify", t, func() {
		Convey("Parse Contract", func() {
			mockCtl := gomock.NewController(t)
			pool := core_mock.NewMockPool(mockCtl)
			var count int
			var k, f, k2 state.Key
			var v, v2 state.Value
			pool.EXPECT().PutHM(gomock.Any(),gomock.Any(),gomock.Any()).Times(2).Do(func(key, field state.Key, value state.Value) error {
				k ,f ,v = key, field, value
				count ++
				return nil
			})
			pool.EXPECT().Put(gomock.Any(), gomock.Any()).Do(func(key state.Key, value state.Value){
				k2, v2 = key, value
			})
			pool.EXPECT().Copy().Return(pool)
			contract := vm_mock.NewMockContract(mockCtl)
			contract.EXPECT().Code().Return(`
-- @PutHM iost abc f10000
-- @PutHM iost def f1000
-- @Put hello sworld
`)
			_, err := ParseGenesis(contract, pool)
			So(err, ShouldBeNil)
			So(count, ShouldEqual, 2)
			So(k, ShouldEqual, state.Key("iost"))
			So(v2.EncodeString(), ShouldEqual, "sworld")

		})
	})
}

