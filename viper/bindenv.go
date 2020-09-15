package viper

import (
	"github.com/spf13/viper"
	"reflect"
	"strings"
)

func BindEnvs(v *viper.Viper, iface interface{}, parts ...string) error {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		fieldv := ifv.Field(i)
		t := ift.Field(i)
		name := strings.ToLower(t.Name)
		tag, ok := t.Tag.Lookup("mapstructure")
		if ok {
			name = tag
		}
		path := append(parts, name)
		switch fieldv.Kind() {
		case reflect.Struct:
			err := BindEnvs(v, fieldv.Interface(), path...)
			if err != nil {
				return err
			}
		default:
			err := v.BindEnv(strings.Join(path, "."))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
