package array

import "reflect"
//get int64 column from Two dimensional array
//从二维数组里的int64列 组成一个独立一维数组

func GetIds(lists interface{}, field string) []int64 {
	t := reflect.TypeOf(lists)
	v := reflect.ValueOf(lists)
	data := make([]int64, v.Len())
	if t.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			temp := v.Index(i).FieldByName(field)
			if temp.IsValid() {
				switch temp.Type().Kind() {
				case reflect.Int64:
					data[i] = temp.Int()
				}
			}
		}
	}
	return data
}
//get string column from Two dimensional array
//从二维数组里的string列 组成一个独立一维数组
func GetStrArr(lists interface{}, field string) []string {
	t := reflect.TypeOf(lists)
	v := reflect.ValueOf(lists)
	data := make([]string, v.Len())
	if t.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			temp := v.Index(i).FieldByName(field)
			if temp.IsValid() {
				switch temp.Type().Kind() {
				case reflect.String:
					data[i] = temp.String()
				}
			}
		}
	}
	return data
}