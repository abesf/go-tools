package file

import (
	"bufio"
	"io"
	"os"
)

/**read file path use bufio
 *return file size and file content
 */
func Read(filename string) (int, []byte) {
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	buf := make([]byte, 4096)
	var nbytes int
	chunks := make([]byte, 0)
	rd := bufio.NewReader(fi)
	for {
		n, err := rd.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		nbytes += n
		chunks = append(chunks, buf[:n]...)
	}
	return nbytes, chunks
}