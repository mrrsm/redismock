package redismock

import "github.com/redis/go-redis/v9"

func (m *mock) ExpectJSONArrAppend(key, path string, values ...interface{}) *ExpectedIntSliceCmd {
	e := &ExpectedIntSliceCmd{}
	e.cmd = m.factory.JSONArrAppend(m.ctx, key, path, values)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONArrIndex(key, path string, value ...interface{}) *ExpectedIntSliceCmd {
	e := &ExpectedIntSliceCmd{}
	e.cmd = m.factory.JSONArrIndex(m.ctx, key, path, value)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONArrIndexWithArgs(key, path string, options *redis.JSONArrIndexArgs, value ...interface{}) *ExpectedIntSliceCmd {
	e := &ExpectedIntSliceCmd{}
	e.cmd = m.factory.JSONArrIndexWithArgs(m.ctx, key, path, options, value)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONArrInsert(key, path string, index int64, values ...interface{}) *ExpectedIntSliceCmd {
	e := &ExpectedIntSliceCmd{}
	e.cmd = m.factory.JSONArrInsert(m.ctx, key, path, index, values)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONArrLen(key, path string) *ExpectedIntSliceCmd {
	e := &ExpectedIntSliceCmd{}
	e.cmd = m.factory.JSONArrLen(m.ctx, key, path)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONArrPop(key, path string, index int) *ExpectedStringSliceCmd {
	e := &ExpectedStringSliceCmd{}
	e.cmd = m.factory.JSONArrPop(m.ctx, key, path, index)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONArrTrim(key, path string) *ExpectedIntSliceCmd {
	e := &ExpectedIntSliceCmd{}
	e.cmd = m.factory.JSONArrTrim(m.ctx, key, path)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONArrTrimWithArgs(key, path string, options *redis.JSONArrTrimArgs) *ExpectedIntSliceCmd {
	e := &ExpectedIntSliceCmd{}
	e.cmd = m.factory.JSONArrTrimWithArgs(m.ctx, key, path, options)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONClear(key, path string) *ExpectedIntCmd {
	e := &ExpectedIntCmd{}
	e.cmd = m.factory.JSONClear(m.ctx, key, path)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONDebugMemory(key, path string) *ExpectedIntCmd {
	e := &ExpectedIntCmd{}
	e.cmd = m.factory.JSONDebugMemory(m.ctx, key, path)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONDel(key, path string) *ExpectedIntCmd {
	e := &ExpectedIntCmd{}
	e.cmd = m.factory.JSONDel(m.ctx, key, path)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONForget(key, path string) *ExpectedIntCmd {
	e := &ExpectedIntCmd{}
	e.cmd = m.factory.JSONForget(m.ctx, key, path)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONGet(key string, paths ...string) *ExpectedJSONCmd {
	e := &ExpectedJSONCmd{}
	e.cmd = m.factory.JSONGet(m.ctx, key, paths...)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONGetWithArgs(key string, options *redis.JSONGetArgs, paths ...string) *ExpectedJSONCmd {
	e := &ExpectedJSONCmd{}
	e.cmd = m.factory.JSONGetWithArgs(m.ctx, key, options, paths...)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONMerge(key, path string, value string) *ExpectedStatusCmd {
	e := &ExpectedStatusCmd{}
	e.cmd = m.factory.JSONMerge(m.ctx, key, path, value)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONMSetArgs(docs []redis.JSONSetArgs) *ExpectedStatusCmd {
	e := &ExpectedStatusCmd{}
	e.cmd = m.factory.JSONMSetArgs(m.ctx, docs)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONMSet(params ...interface{}) *ExpectedStatusCmd {
	e := &ExpectedStatusCmd{}
	e.cmd = m.factory.JSONMSet(m.ctx, params...)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONMGet(path string, keys ...string) *ExpectedJSONSliceCmd {
	e := &ExpectedJSONSliceCmd{}
	e.cmd = m.factory.JSONMGet(m.ctx, path, keys...)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONNumIncrBy(key, path string, value float64) *ExpectedJSONCmd {
	e := &ExpectedJSONCmd{}
	e.cmd = m.factory.JSONNumIncrBy(m.ctx, key, path, value)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONObjKeys(key, path string) *ExpectedSliceCmd {
	e := &ExpectedSliceCmd{}
	e.cmd = m.factory.JSONObjKeys(m.ctx, key, path)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONObjLen(key, path string) *ExpectedIntPointerSliceCmd {
	e := &ExpectedIntPointerSliceCmd{}
	e.cmd = m.factory.JSONObjLen(m.ctx, key, path)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONSet(key, path string, value interface{}) *ExpectedStatusCmd {
	e := &ExpectedStatusCmd{}
	e.cmd = m.factory.JSONSet(m.ctx, key, path, value)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONSetMode(key, path string, value interface{}, mode string) *ExpectedStatusCmd {
	e := &ExpectedStatusCmd{}
	e.cmd = m.factory.JSONSetMode(m.ctx, key, path, value, mode)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONStrAppend(key, path, value string) *ExpectedIntPointerSliceCmd {
	e := &ExpectedIntPointerSliceCmd{}
	e.cmd = m.factory.JSONStrAppend(m.ctx, key, path, value)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONStrLen(key, path string) *ExpectedIntPointerSliceCmd {
	e := &ExpectedIntPointerSliceCmd{}
	e.cmd = m.factory.JSONStrLen(m.ctx, key, path)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONToggle(key, path string) *ExpectedIntPointerSliceCmd {
	e := &ExpectedIntPointerSliceCmd{}
	e.cmd = m.factory.JSONToggle(m.ctx, key, path)
	m.pushExpect(e)
	return e
}

func (m *mock) ExpectJSONType(key, path string) *ExpectedJSONSliceCmd {
	e := &ExpectedJSONSliceCmd{}
	e.cmd = m.factory.JSONType(m.ctx, key, path)
	m.pushExpect(e)
	return e
}
