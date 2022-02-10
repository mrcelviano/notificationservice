package commons

import "reflect"

func FindDbTags(obj interface{}, exceptions ...string) []string {
	return FindTags("db", obj, exceptions...)
}

func FindTags(tag string, obj interface{}, exceptions ...string) []string {
	exc := make(map[string]bool)
	for _, e := range exceptions {
		exc[e] = true
	}
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Struct {
		panic("obj is not struct")
	}
	t := v.Type()

	cols := make([]string, 0)

	for i := 0; i < t.NumField(); i++ {
		dbTag := t.Field(i).Tag.Get(tag)
		if dbTag == "" || exc[dbTag] {
			continue
		}
		cols = append(cols, dbTag)
	}
	return cols
}

func MakeJsonToDbTagMap(obj interface{}) map[string]string {
	return MakeTagMap(obj, "json", "db")
}

func MakeTagMap(obj interface{}, keyTag string, valTag string) map[string]string {
	res := map[string]string{}
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Struct {
		panic("obj is not struct")
	}
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		tags := t.Field(i).Tag
		keyTagVal := tags.Get(keyTag)
		if keyTagVal == "" || keyTagVal == "-" {
			continue
		}
		valTagVal := tags.Get(valTag)
		if valTagVal != "" && valTagVal != "-" {
			res[keyTagVal] = valTagVal
		}
	}
	return res
}
