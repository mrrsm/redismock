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

func (cmd *ExpectedMapMapStringInterface) SetVal(val map[string]interface{}) {
	cmd.setVal = true
	cmd.val = val
}

type ExpectedFTSearch struct {
	expectedBase

	val redis.FTSearchResult
}

func (cmd *ExpectedFTSearch) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedFTSearch) SetVal(val redis.FTSearchResult) {
	cmd.setVal = true
	cmd.val = val
}

type ExpectedFTSpellCheck struct {
	expectedBase

	val []redis.SpellCheckResult
}

func (cmd *ExpectedFTSpellCheck) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedFTSpellCheck) SetVal(val []redis.SpellCheckResult) {
	cmd.setVal = true
	cmd.val = val
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

type ExpectedFTInfo struct {
	expectedBase

	val redis.FTInfoResult
}

func (cmd *ExpectedFTInfo) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedFTInfo) SetVal(val redis.FTInfoResult) {
	cmd.setVal = true
	cmd.val = val
}

type ExpectedFTSynDump struct {
	expectedBase

	val []redis.FTSynDumpResult
}

func (cmd *ExpectedFTSynDump) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedFTSynDump) SetVal(val []redis.FTSynDumpResult) {
	cmd.setVal = true
	cmd.val = val
}
