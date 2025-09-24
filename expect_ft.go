package redismock

import (
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type ExpectedFTSearchCmd struct {
	expectedBase

	val redis.FTSearchResult
}

func (cmd *ExpectedFTSearchCmd) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedFTSearchCmd) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

func (cmd *ExpectedFTSearchCmd) SetVal(val redis.FTSearchResult) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedFTSearchCmd) Result() (redis.FTSearchResult, error) {
	return cmd.val, cmd.err
}

func (cmd *ExpectedFTSearchCmd) Val() redis.FTSearchResult {
	return cmd.val
}

func (cmd *ExpectedFTSearchCmd) RawVal() interface{} {
	return cmd.rawVal
}

func (cmd *ExpectedFTSearchCmd) RawResult() (interface{}, error) {
	return cmd.rawVal, cmd.err
}

type ExpectedFTSpellCheckCmd struct {
	expectedBase

	val []redis.SpellCheckResult
}

func (cmd *ExpectedFTSpellCheckCmd) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedFTSpellCheckCmd) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

func (cmd *ExpectedFTSpellCheckCmd) SetVal(val []redis.SpellCheckResult) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedFTSpellCheckCmd) Result() ([]redis.SpellCheckResult, error) {
	return cmd.val, cmd.err
}

func (cmd *ExpectedFTSpellCheckCmd) Val() []redis.SpellCheckResult {
	return cmd.val
}

func (cmd *ExpectedFTSpellCheckCmd) RawVal() interface{} {
	return cmd.rawVal
}

func (cmd *ExpectedFTSpellCheckCmd) RawResult() (interface{}, error) {
	return cmd.rawVal, cmd.err
}

type ExpectedMapStringInterfaceCmd struct {
	expectedBase

	val map[string]interface{}
}

func (cmd *ExpectedMapStringInterfaceCmd) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedMapStringInterfaceCmd) SetVal(val map[string]interface{}) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedMapStringInterfaceCmd) Val() map[string]interface{} {
	return cmd.val
}

func (cmd *ExpectedMapStringInterfaceCmd) Result() (map[string]interface{}, error) {
	return cmd.val, cmd.err
}

func (cmd *ExpectedMapStringInterfaceCmd) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

type ExpectedStringCmd struct {
	expectedBase

	val string
}

func (cmd *ExpectedStringCmd) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedStringCmd) SetVal(val string) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedStringCmd) Val() string {
	return cmd.val
}

func (cmd *ExpectedStringCmd) Result() (string, error) {
	return cmd.val, cmd.err
}

func (cmd *ExpectedStringCmd) Bytes() ([]byte, error) {
	return StringToBytes(cmd.val), cmd.err
}

func (cmd *ExpectedStringCmd) Bool() (bool, error) {
	if cmd.err != nil {
		return false, cmd.err
	}
	return strconv.ParseBool(cmd.val)
}

func (cmd *ExpectedStringCmd) Int() (int, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return strconv.Atoi(cmd.Val())
}

func (cmd *ExpectedStringCmd) Int64() (int64, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return strconv.ParseInt(cmd.Val(), 10, 64)
}

func (cmd *ExpectedStringCmd) Uint64() (uint64, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return strconv.ParseUint(cmd.Val(), 10, 64)
}

func (cmd *ExpectedStringCmd) Float32() (float32, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	f, err := strconv.ParseFloat(cmd.Val(), 32)
	if err != nil {
		return 0, err
	}
	return float32(f), nil
}

func (cmd *ExpectedStringCmd) Float64() (float64, error) {
	if cmd.err != nil {
		return 0, cmd.err
	}
	return strconv.ParseFloat(cmd.Val(), 64)
}

func (cmd *ExpectedStringCmd) Time() (time.Time, error) {
	if cmd.err != nil {
		return time.Time{}, cmd.err
	}
	return time.Parse(time.RFC3339Nano, cmd.Val())
}

func (cmd *ExpectedStringCmd) Scan(val interface{}) error {
	if cmd.err != nil {
		return cmd.err
	}
	return Scan([]byte(cmd.val), val)
}

func (cmd *ExpectedStringCmd) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

type ExpectedAggregateCmd struct {
	expectedBase

	val *redis.FTAggregateResult
}

func (cmd *ExpectedAggregateCmd) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedAggregateCmd) SetVal(val *redis.FTAggregateResult) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedAggregateCmd) Val() *redis.FTAggregateResult {
	return cmd.val
}

func (cmd *ExpectedAggregateCmd) Result() (*redis.FTAggregateResult, error) {
	return cmd.val, cmd.err
}

func (cmd *ExpectedAggregateCmd) RawVal() interface{} {
	return cmd.rawVal
}

func (cmd *ExpectedAggregateCmd) RawResult() (interface{}, error) {
	return cmd.rawVal, cmd.err
}

func (cmd *ExpectedAggregateCmd) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

type ExpectedFTInfoCmd struct {
	expectedBase

	val redis.FTInfoResult
}

func (cmd *ExpectedFTInfoCmd) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedFTInfoCmd) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

func (cmd *ExpectedFTInfoCmd) SetVal(val redis.FTInfoResult) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedFTInfoCmd) Result() (redis.FTInfoResult, error) {
	return cmd.val, cmd.err
}

func (cmd *ExpectedFTInfoCmd) Val() redis.FTInfoResult {
	return cmd.val
}

func (cmd *ExpectedFTInfoCmd) RawVal() interface{} {
	return cmd.rawVal
}

func (cmd *ExpectedFTInfoCmd) RawResult() (interface{}, error) {
	return cmd.rawVal, cmd.err
}

type ExpectedFTSynDumpCmd struct {
	expectedBase

	val []redis.FTSynDumpResult
}

func (cmd *ExpectedFTSynDumpCmd) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedFTSynDumpCmd) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

func (cmd *ExpectedFTSynDumpCmd) SetVal(val []redis.FTSynDumpResult) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedFTSynDumpCmd) Val() []redis.FTSynDumpResult {
	return cmd.val
}

func (cmd *ExpectedFTSynDumpCmd) Result() ([]redis.FTSynDumpResult, error) {
	return cmd.val, cmd.err
}

func (cmd *ExpectedFTSynDumpCmd) RawVal() interface{} {
	return cmd.rawVal
}

func (cmd *ExpectedFTSynDumpCmd) RawResult() (interface{}, error) {
	return cmd.rawVal, cmd.err
}
