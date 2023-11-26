package env

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashibuto/mirage"
)

// Initialize initializes target from the environment
func Initialize(target any) error {
	r := mirage.Reflect(target, "env")
	ioObj := r.Io()
	for i, key := range r.Keys() {
		info, err := r.InfoByName(key)
		if err != nil {
			return err
		}
		if len(info.TagParts) == 0 {
			continue
		}
		envVar := info.TagParts[0]
		value, isSet := os.LookupEnv(envVar)
		if !isSet {
			if len(info.TagParts) > 1 {
				value = info.TagParts[1]
			} else {
				return fmt.Errorf("environment variable %s is required and no default was provided", envVar)
			}
		}

		var setVal any
		switch info.Kind {
		case reflect.String:
			setVal = value
		case reflect.Float32:
			setVal, err = strconv.ParseFloat(value, 32)
			if err != nil {
				return err
			}
			setVal = float32(setVal.(float64))
		case reflect.Float64:
			setVal, err = strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
		case reflect.Int:
			setVal, err = strconv.ParseInt(value, 0, 0)
			if err != nil {
				return err
			}
			setVal = int(setVal.(int64))
		case reflect.Int8:
			setVal, err = strconv.ParseInt(value, 0, 8)
			if err != nil {
				return err
			}
			setVal = int8(setVal.(int64))
		case reflect.Int16:
			setVal, err = strconv.ParseInt(value, 0, 16)
			if err != nil {
				return err
			}
			setVal = int16(setVal.(int64))
		case reflect.Int32:
			setVal, err = strconv.ParseInt(value, 0, 32)
			if err != nil {
				return err
			}
			setVal = int32(setVal.(int64))
		case reflect.Int64:
			setVal, err = strconv.ParseInt(value, 0, 64)
			if err != nil {
				return err
			}
		case reflect.Bool:
			v := strings.ToLower(value)
			if v == "true" || v == "t" || v == "1" {
				setVal = true
			} else if v == "false" || v == "f" || v == "0" {
				setVal = false
			} else {
				return fmt.Errorf("unable to determine boolean type from value: %s", v)
			}
		default:
			return fmt.Errorf("unsupported data type for field: %s", key)
		}

		err = ioObj.SetValueByIdx(i, setVal)
		if err != nil {
			return err
		}
	}

	return nil
}
