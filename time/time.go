package time

import "time"

//TimeStrToTimeStamp
//timeTemplate1 := "2006-01-02 15:04:05" //常规类型
//timeTemplate2 := "2006/01/02 15:04:05" //其他类型
//timeTemplate3 := "2006-01-02"          //其他类型
//timeTemplate4 := "15:04:05"            //其他类型
func TimeStrToTimeStamp(t1 string, timeType int) int64 {

	timeTemplate := "2006-01-02 15:04:05" //常规类型
	if timeType == 1 {
		timeTemplate = "2006-01-02 15:04:05" //常规类型
	} else if timeType == 2 {
		timeTemplate = "2006/01/02 15:04:05" //常规类型
	} else if timeType == 3 {
		timeTemplate = "2006-01-02" //常规类型
	} else if timeType == 4 {
		timeTemplate = "15:04:05" //常规类型
	} else {
		timeTemplate = "2006-01-02 15:04:05" //常规类型
	}
	stamp, _ := time.ParseInLocation(timeTemplate, t1, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	return stamp.Unix()
}

// TimeNowUnixStampSecond Time Now to UnixStamp (Second) 单位s,打印结果:1491888244
func TimeNowUnixStampSecond() int64 {
	return time.Now().Unix()
}

// TimeStampToTimeStr  TimeStamp To TimeStr
func TimeStampToTimeStr(timestamp int64,timeType int) string {

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
	timeNow:=time.Unix(timestamp, 0)
	timeString := timeNow.Format(timeTemplate) //2015-06-15 08:52:32
	return timeString
}
//2009-03-28 sep  to 2009  03  28
func TimeStrtoYearMonthDay(t1 string, timeType int) (year int,month time.Month,day int) {
	return time.Unix(TimeStrToTimeStamp(t1,timeType),0).Date()
}
