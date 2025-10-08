package redismock

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("JSONCommands", func() {
	var (
		clientMock baseMock
		client     mockCmdable
	)

	callCommandTest := func() {
		It("ExpectJSONArrAppend", func() {
			operationIntSliceCmd(clientMock, func() *ExpectedIntSlice {
				return clientMock.ExpectJSONArrAppend("key", "path", "value")
			}, func() *redis.IntSliceCmd {
				return client.JSONArrAppend(ctx, "key", "path", "value")
			})
		})

		It("ExpectJSONArrIndex", func() {
			operationIntSliceCmd(clientMock, func() *ExpectedIntSlice {
				return clientMock.ExpectJSONArrIndex("key", "path", "value")
			}, func() *redis.IntSliceCmd {
				return client.JSONArrIndex(ctx, "key", "path", "value")
			})
		})

		It("ExpectJSONArrIndexWithArgs", func() {
			operationIntSliceCmd(clientMock, func() *ExpectedIntSlice {
				return clientMock.ExpectJSONArrIndexWithArgs("key", "path", &redis.JSONArrIndexArgs{}, "value")
			}, func() *redis.IntSliceCmd {
				return client.JSONArrIndexWithArgs(ctx, "key", "path", &redis.JSONArrIndexArgs{}, "value")
			})
		})

		It("ExpectJSONArrInsert", func() {
			operationIntSliceCmd(clientMock, func() *ExpectedIntSlice {
				return clientMock.ExpectJSONArrInsert("key", "path", 1, "value")
			}, func() *redis.IntSliceCmd {
				return client.JSONArrInsert(ctx, "key", "path", 1, "value")
			})
		})

		It("ExpectJSONArrLen", func() {
			operationIntSliceCmd(clientMock, func() *ExpectedIntSlice {
				return clientMock.ExpectJSONArrLen("key", "path")
			}, func() *redis.IntSliceCmd {
				return client.JSONArrLen(ctx, "key", "path")
			})
		})

		It("ExpectJSONArrPop", func() {
			operationStringSliceCmd(clientMock, func() *ExpectedStringSlice {
				return clientMock.ExpectJSONArrPop("key", "path", 1)
			}, func() *redis.StringSliceCmd {
				return client.JSONArrPop(ctx, "key", "path", 1)
			})
		})

		It("ExpectJSONArrTrim", func() {
			operationIntSliceCmd(clientMock, func() *ExpectedIntSlice {
				return clientMock.ExpectJSONArrTrim("key", "path")
			}, func() *redis.IntSliceCmd {
				return client.JSONArrTrim(ctx, "key", "path")
			})
		})

		It("ExpectJSONArrTrimWithArgs", func() {
			operationIntSliceCmd(clientMock, func() *ExpectedIntSlice {
				return clientMock.ExpectJSONArrTrimWithArgs("key", "path", &redis.JSONArrTrimArgs{})
			}, func() *redis.IntSliceCmd {
				return client.JSONArrTrimWithArgs(ctx, "key", "path", &redis.JSONArrTrimArgs{})
			})
		})

		It("ExpectJSONClear", func() {
			operationIntCmd(clientMock, func() *ExpectedInt {
				return clientMock.ExpectJSONClear("key", "path")
			}, func() *redis.IntCmd {
				return client.JSONClear(ctx, "key", "path")
			})
		})

		It("ExpectJSONDebugMemory", func() {
			Skip("ExpectJSONDebugMemory just panics")
		})

		It("ExpectJSONDel", func() {
			operationIntCmd(clientMock, func() *ExpectedInt {
				return clientMock.ExpectJSONDel("key", "path")
			}, func() *redis.IntCmd {
				return client.JSONDel(ctx, "key", "path")
			})
		})

		It("ExpectJSONForget", func() {
			operationIntCmd(clientMock, func() *ExpectedInt {
				return clientMock.ExpectJSONForget("key", "path")
			}, func() *redis.IntCmd {
				return client.JSONForget(ctx, "key", "path")
			})
		})

		It("ExpectJSONGet", func() {
			operationJSONCmd(clientMock, func() *ExpectedJSON {
				return clientMock.ExpectJSONGet("key", "paths")
			}, func() *redis.JSONCmd {
				return client.JSONGet(ctx, "key", "paths")
			})
		})

		It("ExpectJSONGetWithArgs", func() {
			operationJSONCmd(clientMock, func() *ExpectedJSON {
				return clientMock.ExpectJSONGetWithArgs("key", &redis.JSONGetArgs{}, "paths")
			}, func() *redis.JSONCmd {
				return client.JSONGetWithArgs(ctx, "key", &redis.JSONGetArgs{}, "paths")
			})
		})

		It("ExpectJSONMerge", func() {
			operationStatusCmd(clientMock, func() *ExpectedStatus {
				return clientMock.ExpectJSONMerge("key", "path", "value")
			}, func() *redis.StatusCmd {
				return client.JSONMerge(ctx, "key", "path", "value")
			})
		})

		It("ExpectJSONMSetArgs", func() {
			operationStatusCmd(clientMock, func() *ExpectedStatus {
				return clientMock.ExpectJSONMSetArgs([]redis.JSONSetArgs{})
			}, func() *redis.StatusCmd {
				return client.JSONMSetArgs(ctx, []redis.JSONSetArgs{})
			})
		})

		It("ExpectJSONMSet", func() {
			operationStatusCmd(clientMock, func() *ExpectedStatus {
				return clientMock.ExpectJSONMSet("params")
			}, func() *redis.StatusCmd {
				return client.JSONMSet(ctx, "params")
			})
		})

		It("ExpectJSONMGet", func() {
			operationJSONSliceCmd(clientMock, func() *ExpectedJSONSlice {
				return clientMock.ExpectJSONMGet("path", "keys")
			}, func() *redis.JSONSliceCmd {
				return client.JSONMGet(ctx, "path", "keys")
			})
		})

		It("ExpectJSONNumIncrBy", func() {
			operationJSONCmd(clientMock, func() *ExpectedJSON {
				return clientMock.ExpectJSONNumIncrBy("key", "path", 0.1)
			}, func() *redis.JSONCmd {
				return client.JSONNumIncrBy(ctx, "key", "path", 0.1)
			})
		})

		It("ExpectJSONObjKeys", func() {
			operationSliceCmd(clientMock, func() *ExpectedSlice {
				return clientMock.ExpectJSONObjKeys("key", "path")
			}, func() *redis.SliceCmd {
				return client.JSONObjKeys(ctx, "key", "path")
			})
		})

		It("ExpectJSONObjLen", func() {
			operationIntPointerSliceCmd(clientMock, func() *ExpectedIntPointerSlice {
				return clientMock.ExpectJSONObjLen("key", "path")
			}, func() *redis.IntPointerSliceCmd {
				return client.JSONObjLen(ctx, "key", "path")
			})
		})

		It("ExpectJSONSet", func() {
			operationStatusCmd(clientMock, func() *ExpectedStatus {
				return clientMock.ExpectJSONSet("key", "path", "value")
			}, func() *redis.StatusCmd {
				return client.JSONSet(ctx, "key", "path", "value")
			})
		})

		It("ExpectJSONSetMode", func() {
			operationStatusCmd(clientMock, func() *ExpectedStatus {
				return clientMock.ExpectJSONSetMode("key", "path", "value", "NX")
			}, func() *redis.StatusCmd {
				return client.JSONSetMode(ctx, "key", "path", "value", "NX")
			})
		})

		It("ExpectJSONStrAppend", func() {
			operationIntPointerSliceCmd(clientMock, func() *ExpectedIntPointerSlice {
				return clientMock.ExpectJSONStrAppend("key", "path", "value")
			}, func() *redis.IntPointerSliceCmd {
				return client.JSONStrAppend(ctx, "key", "path", "value")
			})
		})

		It("ExpectJSONStrLen", func() {
			operationIntPointerSliceCmd(clientMock, func() *ExpectedIntPointerSlice {
				return clientMock.ExpectJSONStrLen("key", "path")
			}, func() *redis.IntPointerSliceCmd {
				return client.JSONStrLen(ctx, "key", "path")
			})
		})

		It("ExpectJSONToggle", func() {
			operationIntPointerSliceCmd(clientMock, func() *ExpectedIntPointerSlice {
				return clientMock.ExpectJSONToggle("key", "path")
			}, func() *redis.IntPointerSliceCmd {
				return client.JSONToggle(ctx, "key", "path")
			})
		})

		It("ExpectJSONType", func() {
			operationJSONSliceCmd(clientMock, func() *ExpectedJSONSlice {
				return clientMock.ExpectJSONType("key", "path")
			}, func() *redis.JSONSliceCmd {
				return client.JSONType(ctx, "key", "path")
			})
		})
	}

	Describe("client", func() {
		BeforeEach(func() {
			client, clientMock = NewClientMock()
		})

		AfterEach(func() {
			Expect(client.(*redis.Client).Close()).NotTo(HaveOccurred())
			Expect(clientMock.ExpectationsWereMet()).NotTo(HaveOccurred())
		})

		callCommandTest()
	})
})
