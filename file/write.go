package file

import (
	"io/ioutil"
	"os"
)
//just wirte
func WriteFile(path string, content []byte) {
	err := ioutil.WriteFile(path, content, 0644)
	if err != nil {
		panic(err)
	}
	return
}
//append conent with sep
func Tracefile(filepath,str_content,sep string)  {
	fd,_:=os.OpenFile(filepath,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	//fd_time:=time.Now().Format("2006-01-02 15:04:05");
	//fd_content:=strings.Join([]string{"======",fd_time,"=====",str_content,"\n"},"")
	buf:=[]byte(str_content+sep)
	fd.Write(buf)
	fd.Close()
}