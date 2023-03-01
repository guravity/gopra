package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value){
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i)) // 再帰呼び出し
		}
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i)) // 再帰呼び出し
		}
	case reflect.Map:
		for _, key := range val.MapKeys(){
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv(){ // for 初期化文; 条件文; 後処理{処理}→条件文がtrueのときに「処理」が実行される
			walk(v.Interface(), fn)
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult{
			walk(res.Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}
