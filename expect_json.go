package redismock

import (
	"github.com/redis/go-redis/v9"
)

type ExpectedIntPointerSlice struct {
	expectedBase

	val []*int64
}

func (cmd *ExpectedIntPointerSlice) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedIntPointerSlice) SetVal(val []*int64) {
	cmd.setVal = true
	cmd.val = val
}

type ExpectedJSON struct {
	expectedBase

	val      string
	expanded interface{}
}

func (cmd *ExpectedJSON) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedJSON) SetVal(val string) {
	cmd.setVal = true
	cmd.val = val
}

type ExpectedJSONSlice struct {
	expectedBase

	val []interface{}
}

func (cmd *ExpectedJSONSlice) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedJSONSlice) SetVal(val []interface{}) {
	cmd.setVal = true
	cmd.val = val
}
