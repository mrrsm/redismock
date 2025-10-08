package redismock

import (
	"errors"

	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

func operationAggregateCmd(base baseMock, expected func() *ExpectedAggregate, actual func() *redis.AggregateCmd) {
	var (
		setErr             = errors.New("aggregate cmd error")
		val                *redis.FTAggregateResult
		err                error
		uninitalizedResult *redis.FTAggregateResult
	)

	base.ClearExpect()
	expected().SetErr(setErr)
	val, err = actual().Result()
	Expect(err).To(Equal(setErr))
	Expect(val).To(Equal(uninitalizedResult))

	base.ClearExpect()
	expected()
	val, err = actual().Result()
	Expect(err).To(HaveOccurred())
	Expect(val).To(Equal(uninitalizedResult))

	base.ClearExpect()
	expected().SetVal(&redis.FTAggregateResult{Total: 1})
	val, err = actual().Result()
	Expect(err).NotTo(HaveOccurred())
	Expect(val).To(Equal(&redis.FTAggregateResult{Total: 1}))
}

func operationFTInfoCmd(base baseMock, expected func() *ExpectedFTInfo, actual func() *redis.FTInfoCmd) {
	var (
		setErr = errors.New("FTInfo cmd error")
		val    redis.FTInfoResult
		err    error
	)

	base.ClearExpect()
	expected().SetErr(setErr)
	val, err = actual().Result()
	Expect(err).To(Equal(setErr))
	Expect(val).To(Equal(redis.FTInfoResult{}))

	base.ClearExpect()
	expected()
	val, err = actual().Result()
	Expect(err).To(HaveOccurred())
	Expect(val).To(Equal(redis.FTInfoResult{}))

	base.ClearExpect()
	expected().SetVal(redis.FTInfoResult{IndexName: "test"})
	val, err = actual().Result()
	Expect(err).NotTo(HaveOccurred())
	Expect(val).To(Equal(redis.FTInfoResult{IndexName: "test"}))
}

func operationFTSpellCheckCmd(base baseMock, expected func() *ExpectedFTSpellCheck, actual func() *redis.FTSpellCheckCmd) {
	var (
		setErr = errors.New("FTSpellCheck cmd error")
		val    []redis.SpellCheckResult
		err    error
	)

	base.ClearExpect()
	expected().SetErr(setErr)
	val, err = actual().Result()
	Expect(err).To(Equal(setErr))
	Expect(val).To(Equal([]redis.SpellCheckResult(nil)))

	base.ClearExpect()
	expected()
	val, err = actual().Result()
	Expect(err).To(HaveOccurred())
	Expect(val).To(Equal([]redis.SpellCheckResult(nil)))

	base.ClearExpect()
	expected().SetVal([]redis.SpellCheckResult{{Term: "test"}})
	val, err = actual().Result()
	Expect(err).NotTo(HaveOccurred())
	Expect(val).To(Equal([]redis.SpellCheckResult{{Term: "test"}}))
}

func operationFTSearchCmd(base baseMock, expected func() *ExpectedFTSearch, actual func() *redis.FTSearchCmd) {
	var (
		setErr = errors.New("FTSearch cmd error")
		val    redis.FTSearchResult
		err    error
	)

	base.ClearExpect()
	expected().SetErr(setErr)
	val, err = actual().Result()
	Expect(err).To(Equal(setErr))
	Expect(val).To(Equal(redis.FTSearchResult{}))

	base.ClearExpect()
	expected()
	val, err = actual().Result()
	Expect(err).To(HaveOccurred())
	Expect(val).To(Equal(redis.FTSearchResult{}))

	base.ClearExpect()
	expected().SetVal(redis.FTSearchResult{Total: 5})
	val, err = actual().Result()
	Expect(err).NotTo(HaveOccurred())
	Expect(val).To(Equal(redis.FTSearchResult{Total: 5}))
}

func operationFTSynDumpCmd(base baseMock, expected func() *ExpectedFTSynDump, actual func() *redis.FTSynDumpCmd) {
	var (
		setErr = errors.New("FTSearch cmd error")
		val    []redis.FTSynDumpResult
		err    error
	)

	base.ClearExpect()
	expected().SetErr(setErr)
	val, err = actual().Result()
	Expect(err).To(Equal(setErr))
	Expect(val).To(Equal([]redis.FTSynDumpResult(nil)))

	base.ClearExpect()
	expected()
	val, err = actual().Result()
	Expect(err).To(HaveOccurred())
	Expect(val).To(Equal([]redis.FTSynDumpResult(nil)))

	base.ClearExpect()
	expected().SetVal([]redis.FTSynDumpResult{{Term: "test"}})
	val, err = actual().Result()
	Expect(err).NotTo(HaveOccurred())
	Expect(val).To(Equal([]redis.FTSynDumpResult{{Term: "test"}}))
}
