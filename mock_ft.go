package redismock

import "github.com/redis/go-redis/v9"

func (m *mock) ExpectFT_List() *ExpectedStringSliceCmd {
	e := &ExpectedStringSliceCmd{}
	e.cmd = m.factory.FT_List(m.ctx)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTAggregate(index string, query string) *ExpectedMapStringInterfaceCmd {
	e := &ExpectedMapStringInterfaceCmd{}
	e.cmd = m.factory.FTAggregate(m.ctx, index, query)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTAggregateWithArgs(index string, query string, options *redis.FTAggregateOptions) *ExpectedAggregateCmd {
	e := &ExpectedAggregateCmd{}
	e.cmd = m.factory.FTAggregateWithArgs(m.ctx, index, query, options)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTAliasAdd(index string, alias string) *ExpectedStatusCmd {
	e := &ExpectedStatusCmd{}
	e.cmd = m.factory.FTAliasAdd(m.ctx, index, alias)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTAliasDel(alias string) *ExpectedStatusCmd {
	e := &ExpectedStatusCmd{}
	e.cmd = m.factory.FTAliasDel(m.ctx, alias)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTAliasUpdate(index string, alias string) *ExpectedStatusCmd {
	e := &ExpectedStatusCmd{}
	e.cmd = m.factory.FTAliasUpdate(m.ctx, index, alias)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTAlter(index string, skipInitialScan bool, definition []interface{}) *ExpectedStatusCmd {
	e := &ExpectedStatusCmd{}
	e.cmd = m.factory.FTAlter(m.ctx, index, skipInitialScan, definition)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTConfigGet(option string) *ExpectedMapMapStringInterfaceCmd {
	e := &ExpectedMapMapStringInterfaceCmd{}
	e.cmd = m.factory.FTConfigGet(m.ctx, option)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTConfigSet(option string, value interface{}) *ExpectedStatusCmd {
	e := &ExpectedStatusCmd{}
	e.cmd = m.factory.FTConfigSet(m.ctx, option, value)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTCreate(index string, options *redis.FTCreateOptions, schema ...*redis.FieldSchema) *ExpectedStatusCmd {
	e := &ExpectedStatusCmd{}
	e.cmd = m.factory.FTCreate(m.ctx, index, options, schema...)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTCursorDel(index string, cursorId int) *ExpectedStatusCmd {
	e := &ExpectedStatusCmd{}
	e.cmd = m.factory.FTCursorDel(m.ctx, index, cursorId)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTCursorRead(index string, cursorId int, count int) *ExpectedMapStringInterfaceCmd {
	e := &ExpectedMapStringInterfaceCmd{}
	e.cmd = m.factory.FTCursorRead(m.ctx, index, cursorId, count)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTDictAdd(dict string, term ...interface{}) *ExpectedIntCmd {
	e := &ExpectedIntCmd{}
	e.cmd = m.factory.FTDictAdd(m.ctx, dict, term)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTDictDel(dict string, term ...interface{}) *ExpectedIntCmd {
	e := &ExpectedIntCmd{}
	e.cmd = m.factory.FTDictDel(m.ctx, dict, term)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTDictDump(dict string) *ExpectedStringSliceCmd {
	e := &ExpectedStringSliceCmd{}
	e.cmd = m.factory.FTDictDump(m.ctx, dict)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTDropIndex(index string) *ExpectedStatusCmd {
	e := &ExpectedStatusCmd{}
	e.cmd = m.factory.FTDropIndex(m.ctx, index)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTDropIndexWithArgs(index string, options *redis.FTDropIndexOptions) *ExpectedStatusCmd {
	e := &ExpectedStatusCmd{}
	e.cmd = m.factory.FTDropIndexWithArgs(m.ctx, index, options)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTExplain(index string, query string) *ExpectedStringCmd {
	e := &ExpectedStringCmd{}
	e.cmd = m.factory.FTExplain(m.ctx, index, query)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTExplainWithArgs(index string, query string, options *redis.FTExplainOptions) *ExpectedStringCmd {
	e := &ExpectedStringCmd{}
	e.cmd = m.factory.FTExplainWithArgs(m.ctx, index, query, options)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTInfo(index string) *ExpectedFTInfoCmd {
	e := &ExpectedFTInfoCmd{}
	e.cmd = m.factory.FTInfo(m.ctx, index)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTSpellCheck(index string, query string) *ExpectedFTSpellCheckCmd {
	e := &ExpectedFTSpellCheckCmd{}
	e.cmd = m.factory.FTSpellCheck(m.ctx, index, query)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTSpellCheckWithArgs(index string, query string, options *redis.FTSpellCheckOptions) *ExpectedFTSpellCheckCmd {
	e := &ExpectedFTSpellCheckCmd{}
	e.cmd = m.factory.FTSpellCheckWithArgs(m.ctx, index, query, options)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTSearch(index string, query string) *ExpectedFTSearchCmd {
	e := &ExpectedFTSearchCmd{}
	e.cmd = m.factory.FTSearch(m.ctx, index, query)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTSearchWithArgs(index string, query string, options *redis.FTSearchOptions) *ExpectedFTSearchCmd {
	e := &ExpectedFTSearchCmd{}
	e.cmd = m.factory.FTSearchWithArgs(m.ctx, index, query, options)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTSynDump(index string) *ExpectedFTSynDumpCmd {
	e := &ExpectedFTSynDumpCmd{}
	e.cmd = m.factory.FTSynDump(m.ctx, index)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTSynUpdate(index string, synGroupId interface{}, terms []interface{}) *ExpectedStatusCmd {
	e := &ExpectedStatusCmd{}
	e.cmd = m.factory.FTSynUpdate(m.ctx, index, synGroupId, terms)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTSynUpdateWithArgs(index string, synGroupId interface{}, options *redis.FTSynUpdateOptions, terms []interface{}) *ExpectedStatusCmd {
	e := &ExpectedStatusCmd{}
	e.cmd = m.factory.FTSynUpdateWithArgs(m.ctx, index, synGroupId, options, terms)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectFTTagVals(index string, field string) *ExpectedStringSliceCmd {
	e := &ExpectedStringSliceCmd{}
	e.cmd = m.factory.FTTagVals(m.ctx, index, field)
	m.pushExpect(e)
	return e
}
