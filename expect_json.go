package redismock

import (
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type ExpectedIntPointerSlice struct {
	expectedBase

	val []*int64
}

func (cmd *ExpectedIntPointerSlice) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedIntPointerSlice) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

func (cmd *ExpectedIntPointerSlice) SetVal(val []*int64) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedIntPointerSlice) Val() []*int64 {
	return cmd.val
}

func (cmd *ExpectedIntPointerSlice) Result() ([]*int64, error) {
	return cmd.val, cmd.err
}

type ExpectedJSON struct {
	expectedBase

	val      string
	expanded interface{}
}

func (cmd *ExpectedJSON) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedJSON) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

func (cmd *ExpectedJSON) SetVal(val string) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedJSON) Val() string {
	if len(cmd.val) == 0 && cmd.expanded != nil {
		val, err := json.Marshal(cmd.expanded)
		if err != nil {
			cmd.SetErr(err)
			return ""
		}
		return string(val)

	} else {
		return cmd.val
	}
}

func (cmd *ExpectedJSON) Result() (string, error) {
	return cmd.Val(), cmd.cmd.Err()
}

func (cmd *ExpectedJSON) Expanded() (interface{}, error) {
	if len(cmd.val) != 0 && cmd.expanded == nil {
		err := json.Unmarshal([]byte(cmd.val), &cmd.expanded)
		if err != nil {
			return nil, err
		}
	}

	return cmd.expanded, nil
}

type ExpectedJSONSlice struct {
	expectedBase

	val []interface{}
}

func (cmd *ExpectedJSONSlice) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedJSONSlice) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

func (cmd *ExpectedJSONSlice) SetVal(val []interface{}) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedJSONSlice) Val() []interface{} {
	return cmd.val
}

func (cmd *ExpectedJSONSlice) Result() ([]interface{}, error) {
	return cmd.val, cmd.err
}
