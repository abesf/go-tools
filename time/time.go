package time

import "time"

//TimeStrToTimeStamp
//timeTemplate1 := "2006-01-02 15:04:05" //常规类型
//timeTemplate2 := "2006/01/02 15:04:05" //其他类型
//timeTemplate3 := "2006-01-02"          //其他类型
//timeTemplate4 := "15:04:05"            //其他类型
func TimeStrToTimeStamp(t1 string,timeType int) int64 {

	timeTemplate := "2006-01-02 15:04:05" //常规类型
	if timeType==1{
		timeTemplate = "2006-01-02 15:04:05" //常规类型
	} else if timeType==2{
		timeTemplate = "2006/01/02 15:04:05" //常规类型
	} else if timeType==3{
		timeTemplate = "2006-01-02" //常规类型
	} else if timeType==4{
		timeTemplate = "15:04:05" //常规类型
	}else {
		timeTemplate = "2006-01-02 15:04:05" //常规类型
	}
	stamp, _ := time.ParseInLocation(timeTemplate, t1, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	return stamp.Unix()
}
