package teleBot

import (
	"fmt"
	"reflect"
)

func value2String(val reflect.Value) string {
	if val.Kind() == reflect.String && val.String() == "" {
		return "\"\""
	}
	return fmt.Sprint(val)
}

func StructToString(intf interface{}) string {
	val := reflect.ValueOf(intf)
	buf := make([]byte, 500)
	buf = append(buf, val.Type().PkgPath()+"/"+val.Type().Name()+" { "...)
	field := reflect.Value{}
	for i := 0; i < val.Type().NumField(); i++ {
		field = val.Field(i)
		buf = append(buf, val.Type().Field(i).Name+": "...)
		// important to use field.Interface() below, otherwise String method of reflect.Value will be found
		if mString, ok := reflect.TypeOf(field.Interface()).MethodByName("String"); ok {
			buf = append(buf, mString.Func.Call([]reflect.Value{field})[0].String()...)
		} else {
			buf = append(buf, value2String(reflect.Indirect(field))...)
		}
		if i < val.Type().NumField()-1 {
			buf = append(buf, ", "...)
		}
	}
	return string(append(buf, " }"...))
}

func MaxId(updates []Update) int {
	max := -1
	for _, upd := range updates {
		if upd.Id > max {
			max = upd.Id
		}
	}
	return max
}

func IsValidPointer(arg interface{}) bool {
	rv := reflect.ValueOf(arg)
	return rv.Kind() == reflect.Ptr && !rv.IsNil()
}

func GroupByChatId(updates *GetUpdatesResponse) map[int][]Update {
	grouped := make(map[int][]Update)
	for _, upd := range updates.Res {
		grouped[upd.Message.Chat.Id] = append(grouped[upd.Message.Chat.Id], upd)
	}
	return grouped
}
