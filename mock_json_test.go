package redismock

import (
	"errors"

	. "github.com/onsi/gomega"
	"github.com/redis/go-redis/v9"
)

func operationJSONCmd(base baseMock, expected func() *ExpectedJSON, actual func() *redis.JSONCmd) {
	var (
		setErr = errors.New("JSON cmd error")
		val    string
		err    error
	)

	base.ClearExpect()
	expected().SetErr(setErr)
	val, err = actual().Result()
	Expect(err).To(Equal(setErr))
	Expect(val).To(Equal(""))

	base.ClearExpect()
	expected()
	val, err = actual().Result()
	Expect(err).To(HaveOccurred())
	Expect(val).To(Equal(""))

	base.ClearExpect()
	expected().SetVal("test")
	val, err = actual().Result()
	Expect(err).NotTo(HaveOccurred())
	Expect(val).To(Equal("test"))
}

func operationJSONSliceCmd(base baseMock, expected func() *ExpectedJSONSlice, actual func() *redis.JSONSliceCmd) {
	var (
		setErr = errors.New("string slice cmd error")
		val    []interface{}
		err    error
	)

	base.ClearExpect()
	expected().SetErr(setErr)
	val, err = actual().Result()
	Expect(err).To(Equal(setErr))
	Expect(val).To(Equal([]interface{}(nil)))

	base.ClearExpect()
	expected()
	val, err = actual().Result()
	Expect(err).To(HaveOccurred())
	Expect(val).To(Equal([]interface{}(nil)))

	base.ClearExpect()
	expected().SetVal([]interface{}{"redis", "mock"})
	val, err = actual().Result()
	Expect(err).NotTo(HaveOccurred())
	Expect(val).To(Equal([]interface{}{"redis", "mock"}))
}
