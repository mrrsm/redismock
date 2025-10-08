package redismock

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

var _ = Describe("FTCommands", func() {
	var (
		clientMock baseMock
		client     mockCmdable
	)

	callCommandTest := func() {
		It("ExpectFT_List", func() {
			operationStringSliceCmd(clientMock, func() *ExpectedStringSlice {
				return clientMock.ExpectFT_List()
			}, func() *redis.StringSliceCmd {
				return client.FT_List(ctx)
			})
		})

		It("ExpectFTAggregate", func() {
			operationMapStringInterfaceCmd(clientMock, func() *ExpectedMapStringInterface {
				return clientMock.ExpectFTAggregate("index", "query")
			}, func() *redis.MapStringInterfaceCmd {
				return client.FTAggregate(ctx, "index", "query")
			})
		})

		// ExpectFTAggregateWithArgs(index string, query string, options *redis.FTAggregateOptions) *ExpectedAggregate
		It("ExpectFTAggregateWithArgs", func() {
			operationAggregateCmd(clientMock, func() *ExpectedAggregate {
				return clientMock.ExpectFTAggregateWithArgs("index", "query", &redis.FTAggregateOptions{})
			}, func() *redis.AggregateCmd {
				return client.FTAggregateWithArgs(ctx, "index", "query", &redis.FTAggregateOptions{})
			})
		})

		It("ExpectFTAliasAdd", func() {
			operationStatusCmd(clientMock, func() *ExpectedStatus {
				return clientMock.ExpectFTAliasAdd("index", "alias")
			}, func() *redis.StatusCmd {
				return client.FTAliasAdd(ctx, "index", "alias")
			})
		})

		It("ExpectFTAliasDel", func() {
			operationStatusCmd(clientMock, func() *ExpectedStatus {
				return clientMock.ExpectFTAliasDel("alias")
			}, func() *redis.StatusCmd {
				return client.FTAliasDel(ctx, "alias")
			})
		})

		It("ExpectFTAliasUpdate", func() {
			operationStatusCmd(clientMock, func() *ExpectedStatus {
				return clientMock.ExpectFTAliasUpdate("index", "alias")
			}, func() *redis.StatusCmd {
				return client.FTAliasUpdate(ctx, "index", "alias")
			})
		})

		It("ExpectFTAlter", func() {
			operationStatusCmd(clientMock, func() *ExpectedStatus {
				return clientMock.ExpectFTAlter("index", false, []interface{}(nil))
			}, func() *redis.StatusCmd {
				return client.FTAlter(ctx, "index", false, []interface{}(nil))
			})
		})

		It("ExpectFTConfigGet", func() {
			operationMapMapStringInterfaceCmd(clientMock, func() *ExpectedMapMapStringInterface {
				return clientMock.ExpectFTConfigGet("option")
			}, func() *redis.MapMapStringInterfaceCmd {
				return client.FTConfigGet(ctx, "option")
			})
		})

		It("ExpectFTConfigSet", func() {
			operationStatusCmd(clientMock, func() *ExpectedStatus {
				return clientMock.ExpectFTConfigSet("option", interface{}(nil))
			}, func() *redis.StatusCmd {
				return client.FTConfigSet(ctx, "option", interface{}(nil))
			})
		})

		It("ExpectFTCreate", func() {
			operationStatusCmd(clientMock, func() *ExpectedStatus {
				return clientMock.ExpectFTCreate("index", &redis.FTCreateOptions{}, &redis.FieldSchema{FieldName: "test", FieldType: redis.SearchFieldTypeGeo})
			}, func() *redis.StatusCmd {
				return client.FTCreate(ctx, "index", &redis.FTCreateOptions{}, &redis.FieldSchema{FieldName: "test", FieldType: redis.SearchFieldTypeGeo})
			})
		})

		It("ExpectFTCursorDel", func() {
			operationStatusCmd(clientMock, func() *ExpectedStatus {
				return clientMock.ExpectFTCursorDel("index", 0)
			}, func() *redis.StatusCmd {
				return client.FTCursorDel(ctx, "index", 0)
			})
		})

		It("ExpectFTCursorRead", func() {
			operationMapStringInterfaceCmd(clientMock, func() *ExpectedMapStringInterface {
				return clientMock.ExpectFTCursorRead("index", 0, 1)
			}, func() *redis.MapStringInterfaceCmd {
				return client.FTCursorRead(ctx, "index", 0, 1)
			})
		})

		It("ExpectFTDictAdd", func() {
			operationIntCmd(clientMock, func() *ExpectedInt {
				return clientMock.ExpectFTDictAdd("dict", "term")
			}, func() *redis.IntCmd {
				return client.FTDictAdd(ctx, "dict", "term")
			})
		})

		It("ExpectFTDictDel", func() {
			operationIntCmd(clientMock, func() *ExpectedInt {
				return clientMock.ExpectFTDictDel("dict", "term")
			}, func() *redis.IntCmd {
				return client.FTDictDel(ctx, "dict", "term")
			})
		})

		It("ExpectFTDictDump", func() {
			operationStringSliceCmd(clientMock, func() *ExpectedStringSlice {
				return clientMock.ExpectFTDictDump("dict")
			}, func() *redis.StringSliceCmd {
				return client.FTDictDump(ctx, "dict")
			})
		})

		It("ExpectFTDropIndex", func() {
			operationStatusCmd(clientMock, func() *ExpectedStatus {
				return clientMock.ExpectFTDropIndex("index")
			}, func() *redis.StatusCmd {
				return client.FTDropIndex(ctx, "index")
			})
		})

		It("ExpectFTDropIndexWithArgs", func() {
			operationStatusCmd(clientMock, func() *ExpectedStatus {
				return clientMock.ExpectFTDropIndexWithArgs("index", &redis.FTDropIndexOptions{})
			}, func() *redis.StatusCmd {
				return client.FTDropIndexWithArgs(ctx, "index", &redis.FTDropIndexOptions{})
			})
		})

		It("ExpectFTExplain", func() {
			operationStringCmd(clientMock, func() *ExpectedString {
				return clientMock.ExpectFTExplain("key", "query")
			}, func() *redis.StringCmd {
				return client.FTExplain(ctx, "key", "query")
			})
		})

		It("ExpectFTExplainWithArgs", func() {
			operationStringCmd(clientMock, func() *ExpectedString {
				return clientMock.ExpectFTExplainWithArgs("key", "query", &redis.FTExplainOptions{})
			}, func() *redis.StringCmd {
				return client.FTExplainWithArgs(ctx, "key", "query", &redis.FTExplainOptions{})
			})
		})

		It("ExpectFTInfo", func() {
			operationFTInfoCmd(clientMock, func() *ExpectedFTInfo {
				return clientMock.ExpectFTInfo("index")
			}, func() *redis.FTInfoCmd {
				return client.FTInfo(ctx, "index")
			})
		})

		It("ExpectFTSpellCheck", func() {
			operationFTSpellCheckCmd(clientMock, func() *ExpectedFTSpellCheck {
				return clientMock.ExpectFTSpellCheck("index", "query")
			}, func() *redis.FTSpellCheckCmd {
				return client.FTSpellCheck(ctx, "index", "query")
			})
		})

		It("ExpectFTSpellCheckWithArgs", func() {
			operationFTSpellCheckCmd(clientMock, func() *ExpectedFTSpellCheck {
				return clientMock.ExpectFTSpellCheckWithArgs("index", "query", &redis.FTSpellCheckOptions{})
			}, func() *redis.FTSpellCheckCmd {
				return client.FTSpellCheckWithArgs(ctx, "index", "query", &redis.FTSpellCheckOptions{})
			})
		})

		It("ExpectFTSearch", func() {
			operationFTSearchCmd(clientMock, func() *ExpectedFTSearch {
				return clientMock.ExpectFTSearch("index", "query")
			}, func() *redis.FTSearchCmd {
				return client.FTSearch(ctx, "index", "query")
			})
		})

		It("ExpectFTSearchWithArgs", func() {
			operationFTSearchCmd(clientMock, func() *ExpectedFTSearch {
				return clientMock.ExpectFTSearchWithArgs("index", "query", &redis.FTSearchOptions{})
			}, func() *redis.FTSearchCmd {
				return client.FTSearchWithArgs(ctx, "index", "query", &redis.FTSearchOptions{})
			})
		})

		It("ExpectFTSynDump", func() {
			operationFTSynDumpCmd(clientMock, func() *ExpectedFTSynDump {
				return clientMock.ExpectFTSynDump("index")
			}, func() *redis.FTSynDumpCmd {
				return client.FTSynDump(ctx, "index")
			})
		})

		It("ExpectFTSynUpdate", func() {
			operationStatusCmd(clientMock, func() *ExpectedStatus {
				return clientMock.ExpectFTSynUpdate("index", interface{}(nil), []interface{}{})
			}, func() *redis.StatusCmd {
				return client.FTSynUpdate(ctx, "index", interface{}(nil), []interface{}{})
			})
		})

		It("ExpectFTSynUpdateWithArgs", func() {
			operationStatusCmd(clientMock, func() *ExpectedStatus {
				return clientMock.ExpectFTSynUpdateWithArgs("index", interface{}(nil), &redis.FTSynUpdateOptions{}, []interface{}{})
			}, func() *redis.StatusCmd {
				return client.FTSynUpdateWithArgs(ctx, "index", interface{}(nil), &redis.FTSynUpdateOptions{}, []interface{}{})
			})
		})

		It("ExpectFTTagVals", func() {
			operationStringSliceCmd(clientMock, func() *ExpectedStringSlice {
				return clientMock.ExpectFTTagVals("index", "field")
			}, func() *redis.StringSliceCmd {
				return client.FTTagVals(ctx, "index", "field")
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
