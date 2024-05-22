package kutils

import (
	"encoding"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"

	"github.com/KyberNetwork/kutils/internal/json"
)

// StringUnmarshalHookFunc converts string values using json.Unmarshaler or encoding.TextUnmarshaler
// if the destination type implements it.
func StringUnmarshalHookFunc() mapstructure.DecodeHookFuncValue {
	return func(f, t reflect.Value) (any, error) {
		if f.Kind() != reflect.String {
			return f.Interface(), nil
		}
		data := f.String()

		var dest any
		if t.Kind() == reflect.Pointer {
			if t.IsNil() {
				dest = reflect.New(t.Type().Elem()).Interface()
			} else {
				dest = t.Interface()
			}
		} else if t.CanAddr() {
			dest = t.Addr().Interface()
		} else {
			dest = reflect.New(t.Type()).Interface()
		}

		switch dest := dest.(type) {
		case encoding.TextUnmarshaler:
			return dest, dest.UnmarshalText([]byte(data))
		case json.Unmarshaler:
			quoted, _ := json.Marshal(data)
			return dest, dest.UnmarshalJSON(quoted)
		}
		return data, nil
	}
}

// StringToTimeDurationHookFunc converts string to time.Duration. It also handles json.Number.
func StringToTimeDurationHookFunc() mapstructure.DecodeHookFuncValue {
	durationType := reflect.TypeOf(time.Duration(0))
	return func(f, t reflect.Value) (any, error) {
		if f.Kind() != reflect.String {
			return f.Interface(), nil
		}
		if t.Type() != durationType {
			return f.Interface(), nil
		}

		data := f.String()
		if data == "" {
			return time.Duration(0), nil
		}

		lastChar := data[len(data)-1]
		if lastChar == '.' || '0' <= lastChar && lastChar <= '9' {
			return Atoi[time.Duration](data)
		}
		return time.ParseDuration(data)
	}
}

// ConfigDecodeHook converts string to time.Duration and string to string slice like viper does.
// It also contains StringUnmarshalHookFunc to account for fields implementing Unmarshal interfaces.
// We also need special handling of string in StringToTimeDurationHookFunc hook in order to
// process json.Number produced by JSONUnmarshal which uses UseNumber option.
var ConfigDecodeHook = mapstructure.ComposeDecodeHookFunc(
	StringUnmarshalHookFunc(),
	StringToTimeDurationHookFunc(),
	mapstructure.StringToSliceHookFunc(","),
)

// DecodeConfig unmarshalls the raw config value bytes into the specified destination with custom logic hooks.
// It uses JSONUnmarshal which enables UseNumber option for maintaining number precision
// and mapstructure to make use of custom hooks including a hook for making use of custom text/json unmarshaler
// and hooks for time and string slice conversion hooks (same as viper config).
func DecodeConfig(data []byte, dest any) error {
	var cfgMap any
	if err := JSONUnmarshal(data, &cfgMap); err != nil {
		return err
	}
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:           dest,
		WeaklyTypedInput: true,
		DecodeHook:       ConfigDecodeHook,
	})
	if err != nil {
		return err
	}
	return decoder.Decode(cfgMap)
}
