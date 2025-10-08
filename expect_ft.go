package redismock

import (
	"github.com/redis/go-redis/v9"
)

type ExpectedMapMapStringInterface struct {
	expectedBase

	val map[string]interface{}
}

func (cmd *ExpectedMapMapStringInterface) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedMapMapStringInterface) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

func (cmd *ExpectedMapMapStringInterface) SetVal(val map[string]interface{}) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedMapMapStringInterface) Result() (map[string]interface{}, error) {
	return cmd.val, cmd.err
}

func (cmd *ExpectedMapMapStringInterface) Val() map[string]interface{} {
	return cmd.val
}

type ExpectedFTSearch struct {
	expectedBase

	val redis.FTSearchResult
}

func (cmd *ExpectedFTSearch) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedFTSearch) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

func (cmd *ExpectedFTSearch) SetVal(val redis.FTSearchResult) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedFTSearch) Result() (redis.FTSearchResult, error) {
	return cmd.val, cmd.err
}

func (cmd *ExpectedFTSearch) Val() redis.FTSearchResult {
	return cmd.val
}

func (cmd *ExpectedFTSearch) RawVal() interface{} {
	return cmd.rawVal
}

func (cmd *ExpectedFTSearch) RawResult() (interface{}, error) {
	return cmd.rawVal, cmd.err
}

type ExpectedFTSpellCheck struct {
	expectedBase

	val []redis.SpellCheckResult
}

func (cmd *ExpectedFTSpellCheck) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedFTSpellCheck) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

func (cmd *ExpectedFTSpellCheck) SetVal(val []redis.SpellCheckResult) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedFTSpellCheck) Result() ([]redis.SpellCheckResult, error) {
	return cmd.val, cmd.err
}

func (cmd *ExpectedFTSpellCheck) Val() []redis.SpellCheckResult {
	return cmd.val
}

func (cmd *ExpectedFTSpellCheck) RawVal() interface{} {
	return cmd.rawVal
}

func (cmd *ExpectedFTSpellCheck) RawResult() (interface{}, error) {
	return cmd.rawVal, cmd.err
}

type ExpectedAggregate struct {
	expectedBase

	val *redis.FTAggregateResult
}

func (cmd *ExpectedAggregate) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedAggregate) SetVal(val *redis.FTAggregateResult) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedAggregate) Val() *redis.FTAggregateResult {
	return cmd.val
}

func (cmd *ExpectedAggregate) Result() (*redis.FTAggregateResult, error) {
	return cmd.val, cmd.err
}

func (cmd *ExpectedAggregate) RawVal() interface{} {
	return cmd.rawVal
}

func (cmd *ExpectedAggregate) RawResult() (interface{}, error) {
	return cmd.rawVal, cmd.err
}

func (cmd *ExpectedAggregate) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

type ExpectedFTInfo struct {
	expectedBase

	val redis.FTInfoResult
}

func (cmd *ExpectedFTInfo) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedFTInfo) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

func (cmd *ExpectedFTInfo) SetVal(val redis.FTInfoResult) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedFTInfo) Result() (redis.FTInfoResult, error) {
	return cmd.val, cmd.err
}

func (cmd *ExpectedFTInfo) Val() redis.FTInfoResult {
	return cmd.val
}

func (cmd *ExpectedFTInfo) RawVal() interface{} {
	return cmd.rawVal
}

func (cmd *ExpectedFTInfo) RawResult() (interface{}, error) {
	return cmd.rawVal, cmd.err
}

type ExpectedFTSynDump struct {
	expectedBase

	val []redis.FTSynDumpResult
}

func (cmd *ExpectedFTSynDump) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedFTSynDump) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

func (cmd *ExpectedFTSynDump) SetVal(val []redis.FTSynDumpResult) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedFTSynDump) Val() []redis.FTSynDumpResult {
	return cmd.val
}

func (cmd *ExpectedFTSynDump) Result() ([]redis.FTSynDumpResult, error) {
	return cmd.val, cmd.err
}

func (cmd *ExpectedFTSynDump) RawVal() interface{} {
	return cmd.rawVal
}

func (cmd *ExpectedFTSynDump) RawResult() (interface{}, error) {
	return cmd.rawVal, cmd.err
}
