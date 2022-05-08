package error

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
	"runtime/debug"
)

// 打印堆栈信息
//print error stack
func PrintStackTrace(err interface{}) string {
	buf := new(bytes.Buffer)
	_, _ = fmt.Fprintf(buf, "%v\n", err)
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		_, _ = fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
	}
	return buf.String()
}
func Debug1()  {
	//1. using runtime
	if p := recover(); p != nil {
		log.Printf("message push panic: %v", p)

		//打印调用栈信息
		buf := make([]byte, 2048)
		n := runtime.Stack(buf, false)
		stackInfo := fmt.Sprintf("%s", buf[:n])
		log.Printf("panic stack info %s", stackInfo)
	}

}
func Debug2()  {
	if p := recover(); p != nil {
		log.Printf("message push panic: %v", p)
		//打印调用栈信息
		debug.PrintStack()
	}

}