package untils

import "reflect"

type MyTools struct {
}

/*
功能：检查某个值是否为空
实现：
参数：
返回：为空返回true 不为空返回false
*/


func (this *MyTools) InterfaceIsNil(i interface{}) bool {
	ret := i == nil
	if !ret {
		defer func() {
			recover()
		}()
		ret = reflect.ValueOf(i).IsNil() // there will exit some panic
	}
	return ret
}
