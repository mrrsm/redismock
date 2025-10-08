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

type ExpectedIntPointerSlice struct {
	expectedBase

	val []*int64
}

func (cmd *ExpectedIntPointerSlice) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedIntPointerSlice) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

func (cmd *ExpectedIntPointerSlice) SetVal(val []*int64) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedIntPointerSlice) Val() []*int64 {
	return cmd.val
}

func (cmd *ExpectedIntPointerSlice) Result() ([]*int64, error) {
	return cmd.val, cmd.err
}

type ExpectedJSON struct {
	expectedBase

	val      string
	expanded interface{}
}

func (cmd *ExpectedJSON) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedJSON) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

func (cmd *ExpectedJSON) SetVal(val string) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedJSON) Val() string {
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

func (cmd *ExpectedJSON) Result() (string, error) {
	return cmd.Val(), cmd.cmd.Err()
}

func (cmd *ExpectedJSON) Expanded() (interface{}, error) {
	if len(cmd.val) != 0 && cmd.expanded == nil {
		err := json.Unmarshal([]byte(cmd.val), &cmd.expanded)
		if err != nil {
			return nil, err
		}
	}

	return cmd.expanded, nil
}

type ExpectedJSONSlice struct {
	expectedBase

	val []interface{}
}

func (cmd *ExpectedJSONSlice) inflow(c redis.Cmder) {
	inflow(c, "val", cmd.val)
}

func (cmd *ExpectedJSONSlice) String() string {
	return cmdString(cmd.cmd, cmd.val)
}

func (cmd *ExpectedJSONSlice) SetVal(val []interface{}) {
	cmd.setVal = true
	cmd.val = val
}

func (cmd *ExpectedJSONSlice) Val() []interface{} {
	return cmd.val
}

func (cmd *ExpectedJSONSlice) Result() ([]interface{}, error) {
	return cmd.val, cmd.err
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
