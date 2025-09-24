package redismock

import (
	"encoding"
	"errors"
	"fmt"
	"net"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/redis/go-redis/v9"
)

var (
	decoders = []decoderFunc{
		reflect.Bool:          decodeBool,
		reflect.Int:           decodeInt,
		reflect.Int8:          decodeInt8,
		reflect.Int16:         decodeInt16,
		reflect.Int32:         decodeInt32,
		reflect.Int64:         decodeInt64,
		reflect.Uint:          decodeUint,
		reflect.Uint8:         decodeUint8,
		reflect.Uint16:        decodeUint16,
		reflect.Uint32:        decodeUint32,
		reflect.Uint64:        decodeUint64,
		reflect.Float32:       decodeFloat32,
		reflect.Float64:       decodeFloat64,
		reflect.Complex64:     decodeUnsupported,
		reflect.Complex128:    decodeUnsupported,
		reflect.Array:         decodeUnsupported,
		reflect.Chan:          decodeUnsupported,
		reflect.Func:          decodeUnsupported,
		reflect.Interface:     decodeUnsupported,
		reflect.Map:           decodeUnsupported,
		reflect.Ptr:           decodeUnsupported,
		reflect.Slice:         decodeSlice,
		reflect.String:        decodeString,
		reflect.Struct:        decodeUnsupported,
		reflect.UnsafePointer: decodeUnsupported,
	}
	globalStructMap = newStructMap()
)

func decodeBool(f reflect.Value, s string) error {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	f.SetBool(b)
	return nil
}

func decodeInt8(f reflect.Value, s string) error {
	return decodeNumber(f, s, 8)
}

func decodeInt16(f reflect.Value, s string) error {
	return decodeNumber(f, s, 16)
}

func decodeInt32(f reflect.Value, s string) error {
	return decodeNumber(f, s, 32)
}

func decodeInt64(f reflect.Value, s string) error {
	return decodeNumber(f, s, 64)
}

func decodeInt(f reflect.Value, s string) error {
	return decodeNumber(f, s, 0)
}

func decodeNumber(f reflect.Value, s string, bitSize int) error {
	v, err := strconv.ParseInt(s, 10, bitSize)
	if err != nil {
		return err
	}
	f.SetInt(v)
	return nil
}

func decodeUint8(f reflect.Value, s string) error {
	return decodeUnsignedNumber(f, s, 8)
}

func decodeUint16(f reflect.Value, s string) error {
	return decodeUnsignedNumber(f, s, 16)
}

func decodeUint32(f reflect.Value, s string) error {
	return decodeUnsignedNumber(f, s, 32)
}

func decodeUint64(f reflect.Value, s string) error {
	return decodeUnsignedNumber(f, s, 64)
}

func decodeUint(f reflect.Value, s string) error {
	return decodeUnsignedNumber(f, s, 0)
}

func decodeUnsignedNumber(f reflect.Value, s string, bitSize int) error {
	v, err := strconv.ParseUint(s, 10, bitSize)
	if err != nil {
		return err
	}
	f.SetUint(v)
	return nil
}

func decodeFloat32(f reflect.Value, s string) error {
	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return err
	}
	f.SetFloat(v)
	return nil
}

// although the default is float64, but we better define it.
func decodeFloat64(f reflect.Value, s string) error {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	f.SetFloat(v)
	return nil
}

func decodeString(f reflect.Value, s string) error {
	f.SetString(s)
	return nil
}

func decodeSlice(f reflect.Value, s string) error {
	// []byte slice ([]uint8).
	if f.Type().Elem().Kind() == reflect.Uint8 {
		f.SetBytes([]byte(s))
	}
	return nil
}

func decodeUnsupported(v reflect.Value, s string) error {
	return fmt.Errorf("redis.Scan(unsupported %s)", v.Type())
}

func newStructSpec(t reflect.Type, fieldTag string) *structSpec {
	numField := t.NumField()
	out := &structSpec{
		m: make(map[string]*structField, numField),
	}

	for i := 0; i < numField; i++ {
		f := t.Field(i)

		tag := f.Tag.Get(fieldTag)
		if tag == "" || tag == "-" {
			continue
		}

		tag = strings.Split(tag, ",")[0]
		if tag == "" {
			continue
		}

		// Use the built-in decoder.
		kind := f.Type.Kind()
		if kind == reflect.Pointer {
			kind = f.Type.Elem().Kind()
		}
		out.set(tag, &structField{index: i, fn: decoders[kind]})
	}

	return out
}

func Struct(dst interface{}) (StructValue, error) {
	v := reflect.ValueOf(dst)

	// The destination to scan into should be a struct pointer.
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return StructValue{}, fmt.Errorf("redis.Scan(non-pointer %T)", dst)
	}

	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return StructValue{}, fmt.Errorf("redis.Scan(non-struct %T)", dst)
	}

	return StructValue{
		spec:  globalStructMap.get(v.Type()),
		value: v,
	}, nil
}

func HScan(dst interface{}, keys []interface{}, vals []interface{}) error {
	if len(keys) != len(vals) {
		return errors.New("args should have the same number of keys and vals")
	}

	strct, err := Struct(dst)
	if err != nil {
		return err
	}

	// Iterate through the (key, value) sequence.
	for i := 0; i < len(vals); i++ {
		key, ok := keys[i].(string)
		if !ok {
			continue
		}

		val, ok := vals[i].(string)
		if !ok {
			continue
		}

		if err := strct.Scan(key, val); err != nil {
			return err
		}
	}

	return nil
}

func cmdString(cmd redis.Cmder, val interface{}) string {
	b := make([]byte, 0, 64)

	for i, arg := range cmd.Args() {
		if i > 0 {
			b = append(b, ' ')
		}
		b = AppendArg(b, arg)
	}

	if err := cmd.Err(); err != nil {
		b = append(b, ": "...)
		b = append(b, err.Error()...)
	} else if val != nil {
		b = append(b, ": "...)
		b = AppendArg(b, val)
	}

	return BytesToString(b)
}

func AppendArg(b []byte, v interface{}) []byte {
	switch v := v.(type) {
	case nil:
		return append(b, "<nil>"...)
	case string:
		return appendUTF8String(b, StringToBytes(v))
	case []byte:
		return appendUTF8String(b, v)
	case int:
		return strconv.AppendInt(b, int64(v), 10)
	case int8:
		return strconv.AppendInt(b, int64(v), 10)
	case int16:
		return strconv.AppendInt(b, int64(v), 10)
	case int32:
		return strconv.AppendInt(b, int64(v), 10)
	case int64:
		return strconv.AppendInt(b, v, 10)
	case uint:
		return strconv.AppendUint(b, uint64(v), 10)
	case uint8:
		return strconv.AppendUint(b, uint64(v), 10)
	case uint16:
		return strconv.AppendUint(b, uint64(v), 10)
	case uint32:
		return strconv.AppendUint(b, uint64(v), 10)
	case uint64:
		return strconv.AppendUint(b, v, 10)
	case float32:
		return strconv.AppendFloat(b, float64(v), 'f', -1, 64)
	case float64:
		return strconv.AppendFloat(b, v, 'f', -1, 64)
	case bool:
		if v {
			return append(b, "true"...)
		}
		return append(b, "false"...)
	case time.Time:
		return v.AppendFormat(b, time.RFC3339Nano)
	default:
		return append(b, fmt.Sprint(v)...)
	}
}

func appendUTF8String(dst []byte, src []byte) []byte {
	dst = append(dst, src...)
	return dst
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func ScanSlice(data []string, slice interface{}) error {
	v := reflect.ValueOf(slice)
	if !v.IsValid() {
		return fmt.Errorf("redis: ScanSlice(nil)")
	}
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("redis: ScanSlice(non-pointer %T)", slice)
	}
	v = v.Elem()
	if v.Kind() != reflect.Slice {
		return fmt.Errorf("redis: ScanSlice(non-slice %T)", slice)
	}

	next := makeSliceNextElemFunc(v)
	for i, s := range data {
		elem := next()
		if err := Scan([]byte(s), elem.Addr().Interface()); err != nil {
			err = fmt.Errorf("redis: ScanSlice index=%d value=%q failed: %w", i, s, err)
			return err
		}
	}

	return nil
}

func Scan(b []byte, v interface{}) error {
	switch v := v.(type) {
	case nil:
		return fmt.Errorf("redis: Scan(nil)")
	case *string:
		*v = BytesToString(b)
		return nil
	case *[]byte:
		*v = b
		return nil
	case *int:
		var err error
		*v, err = strconv.Atoi(BytesToString(b))
		return err
	case *int8:
		n, err := strconv.ParseInt(BytesToString(b), 10, 8)
		if err != nil {
			return err
		}
		*v = int8(n)
		return nil
	case *int16:
		n, err := strconv.ParseInt(BytesToString(b), 10, 16)
		if err != nil {
			return err
		}
		*v = int16(n)
		return nil
	case *int32:
		n, err := strconv.ParseInt(BytesToString(b), 10, 32)
		if err != nil {
			return err
		}
		*v = int32(n)
		return nil
	case *int64:
		n, err := strconv.ParseInt(BytesToString(b), 10, 64)
		if err != nil {
			return err
		}
		*v = n
		return nil
	case *uint:
		n, err := strconv.ParseUint(BytesToString(b), 10, 64)
		if err != nil {
			return err
		}
		*v = uint(n)
		return nil
	case *uint8:
		n, err := strconv.ParseUint(BytesToString(b), 10, 8)
		if err != nil {
			return err
		}
		*v = uint8(n)
		return nil
	case *uint16:
		n, err := strconv.ParseUint(BytesToString(b), 10, 16)
		if err != nil {
			return err
		}
		*v = uint16(n)
		return nil
	case *uint32:
		n, err := strconv.ParseUint(BytesToString(b), 10, 32)
		if err != nil {
			return err
		}
		*v = uint32(n)
		return nil
	case *uint64:
		n, err := strconv.ParseUint(BytesToString(b), 10, 64)
		if err != nil {
			return err
		}
		*v = n
		return nil
	case *float32:
		n, err := strconv.ParseFloat(BytesToString(b), 32)
		if err != nil {
			return err
		}
		*v = float32(n)
		return err
	case *float64:
		var err error
		*v, err = strconv.ParseFloat(BytesToString(b), 64)
		return err
	case *bool:
		*v = len(b) == 1 && b[0] == '1'
		return nil
	case *time.Time:
		var err error
		*v, err = time.Parse(time.RFC3339Nano, BytesToString(b))
		return err
	case *time.Duration:
		n, err := strconv.ParseInt(BytesToString(b), 10, 64)
		if err != nil {
			return err
		}
		*v = time.Duration(n)
		return nil
	case encoding.BinaryUnmarshaler:
		return v.UnmarshalBinary(b)
	case *net.IP:
		*v = b
		return nil
	default:
		return fmt.Errorf(
			"redis: can't unmarshal %T (consider implementing BinaryUnmarshaler)", v)
	}
}

func makeSliceNextElemFunc(v reflect.Value) func() reflect.Value {
	elemType := v.Type().Elem()

	if elemType.Kind() == reflect.Ptr {
		elemType = elemType.Elem()
		return func() reflect.Value {
			if v.Len() < v.Cap() {
				v.Set(v.Slice(0, v.Len()+1))
				elem := v.Index(v.Len() - 1)
				if elem.IsNil() {
					elem.Set(reflect.New(elemType))
				}
				return elem.Elem()
			}

			elem := reflect.New(elemType)
			v.Set(reflect.Append(v, elem))
			return elem.Elem()
		}
	}

	zero := reflect.Zero(elemType)
	return func() reflect.Value {
		if v.Len() < v.Cap() {
			v.Set(v.Slice(0, v.Len()+1))
			return v.Index(v.Len() - 1)
		}

		v.Set(reflect.Append(v, zero))
		return v.Index(v.Len() - 1)
	}
}

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}
