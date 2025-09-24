package redismock

import (
	"encoding"
	"encoding/json"
	"fmt"
	"reflect"
	"sync"

	"github.com/redis/go-redis/v9"
)

type ExpectedMapMapStringInterfaceCmd struct {
	expectedBase

	val map[string]interface{}
}

func (cmd *ExpectedMapMapStringInterfaceCmd) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedMapMapStringInterfaceCmd) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

func (cmd *ExpectedMapMapStringInterfaceCmd) SetVal(val map[string]interface{}) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedMapMapStringInterfaceCmd) Result() (map[string]interface{}, error) {
	return cmd.val, cmd.err
}

func (cmd *ExpectedMapMapStringInterfaceCmd) Val() map[string]interface{} {
	return cmd.val
}

type ExpectedIntSliceCmd struct {
	expectedBase

	val []int64
}

func (cmd *ExpectedIntSliceCmd) SetVal(val []int64) {
	cmd.setVal = true
	cmd.val = make([]int64, len(val))
	copy(cmd.val, val)
}

func (cmd *ExpectedIntSliceCmd) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedIntSliceCmd) Result() ([]int64, error) {
	return cmd.val, cmd.err
}

func (cmd *ExpectedIntSliceCmd) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

func (cmd *ExpectedIntSliceCmd) Val() []int64 {
	return cmd.val
}

type ExpectedStatusCmd struct {
	expectedBase

	val string
}

func (cmd *ExpectedStatusCmd) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedStatusCmd) SetVal(val string) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedStatusCmd) Val() string {
	return cmd.val
}

func (cmd *ExpectedStatusCmd) Result() (string, error) {
	return cmd.val, cmd.err
}

func (cmd *ExpectedStatusCmd) Bytes() ([]byte, error) {
	return StringToBytes(cmd.val), cmd.err
}

func (cmd *ExpectedStatusCmd) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

type ExpectedIntCmd struct {
	expectedBase

	val int64
}

func (cmd *ExpectedIntCmd) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedIntCmd) SetVal(val int64) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedIntCmd) Val() int64 {
	return cmd.val
}

func (cmd *ExpectedIntCmd) Result() (int64, error) {
	return cmd.val, cmd.err
}

func (cmd *ExpectedIntCmd) Uint64() (uint64, error) {
	return uint64(cmd.val), cmd.err
}

func (cmd *ExpectedIntCmd) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

type ExpectedIntPointerSliceCmd struct {
	expectedBase

	val []*int64
}

func (cmd *ExpectedIntPointerSliceCmd) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedIntPointerSliceCmd) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

func (cmd *ExpectedIntPointerSliceCmd) SetVal(val []*int64) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedIntPointerSliceCmd) Val() []*int64 {
	return cmd.val
}

func (cmd *ExpectedIntPointerSliceCmd) Result() ([]*int64, error) {
	return cmd.val, cmd.err
}

type ExpectedJSONCmd struct {
	expectedBase

	val      string
	expanded interface{}
}

func (cmd *ExpectedJSONCmd) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedJSONCmd) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

func (cmd *ExpectedJSONCmd) SetVal(val string) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedJSONCmd) Val() string {
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

func (cmd *ExpectedJSONCmd) Result() (string, error) {
	return cmd.Val(), cmd.cmd.Err()
}

func (cmd *ExpectedJSONCmd) Expanded() (interface{}, error) {
	if len(cmd.val) != 0 && cmd.expanded == nil {
		err := json.Unmarshal([]byte(cmd.val), &cmd.expanded)
		if err != nil {
			return nil, err
		}
	}

	return cmd.expanded, nil
}

type ExpectedJSONSliceCmd struct {
	expectedBase

	val []interface{}
}

func (cmd *ExpectedJSONSliceCmd) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedJSONSliceCmd) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

func (cmd *ExpectedJSONSliceCmd) SetVal(val []interface{}) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedJSONSliceCmd) Val() []interface{} {
	return cmd.val
}

func (cmd *ExpectedJSONSliceCmd) Result() ([]interface{}, error) {
	return cmd.val, cmd.err
}

type ExpectedStringSliceCmd struct {
	expectedBase

	val []string
}

func (cmd *ExpectedStringSliceCmd) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedStringSliceCmd) SetVal(val []string) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedStringSliceCmd) Val() []string {
	return cmd.val
}

func (cmd *ExpectedStringSliceCmd) Result() ([]string, error) {
	return cmd.val, cmd.err
}

func (cmd *ExpectedStringSliceCmd) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

func (cmd *ExpectedStringSliceCmd) ScanSlice(container interface{}) error {
	return ScanSlice(cmd.Val(), container)
}

type ExpectedSliceCmd struct {
	expectedBase

	val []interface{}
}

func (cmd *ExpectedSliceCmd) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedSliceCmd) SetVal(val []interface{}) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedSliceCmd) Val() []interface{} {
	return cmd.val
}

func (cmd *ExpectedSliceCmd) Result() ([]interface{}, error) {
	return cmd.val, cmd.err
}

func (cmd *ExpectedSliceCmd) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

// Scan scans the results from the map into a destination struct. The map keys
// are matched in the Redis struct fields by the `redis:"field"` tag.
func (cmd *ExpectedSliceCmd) Scan(dst interface{}) error {
	if cmd.err != nil {
		return cmd.err
	}

	// Pass the list of keys and values.
	// Skip the first two args for: HMGET key
	var args []interface{}
	if cmd.args()[0] == "hmget" {
		args = cmd.args()[2:]
	} else {
		// Otherwise, it's: MGET field field ...
		args = cmd.args()[1:]
	}

	return HScan(dst, args, cmd.val)
}

type decoderFunc func(reflect.Value, string) error

type structField struct {
	index int
	fn    decoderFunc
}

type structSpec struct {
	m map[string]*structField
}

func (s *structSpec) set(tag string, sf *structField) {
	s.m[tag] = sf
}

type StructValue struct {
	spec  *structSpec
	value reflect.Value
}

type Scanner interface {
	ScanRedis(s string) error
}

func (s StructValue) Scan(key string, value string) error {
	field, ok := s.spec.m[key]
	if !ok {
		return nil
	}

	v := s.value.Field(field.index)
	isPtr := v.Kind() == reflect.Ptr

	if isPtr && v.IsNil() {
		v.Set(reflect.New(v.Type().Elem()))
	}
	if !isPtr && v.Type().Name() != "" && v.CanAddr() {
		v = v.Addr()
		isPtr = true
	}

	if isPtr && v.Type().NumMethod() > 0 && v.CanInterface() {
		switch scan := v.Interface().(type) {
		case Scanner:
			return scan.ScanRedis(value)
		case encoding.TextUnmarshaler:
			return scan.UnmarshalText(StringToBytes(value))
		}
	}

	if isPtr {
		v = v.Elem()
	}

	if err := field.fn(v, value); err != nil {
		t := s.value.Type()
		return fmt.Errorf("cannot scan redis.result %s into struct field %s.%s of type %s, error-%s",
			value, t.Name(), t.Field(field.index).Name, t.Field(field.index).Type, err.Error())
	}
	return nil
}

type structMap struct {
	m sync.Map
}

func newStructMap() *structMap {
	return new(structMap)
}

func (s *structMap) get(t reflect.Type) *structSpec {
	if v, ok := s.m.Load(t); ok {
		return v.(*structSpec)
	}

	spec := newStructSpec(t, "redis")
	s.m.Store(t, spec)
	return spec
}
