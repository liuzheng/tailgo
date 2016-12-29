package main

import (
	"fmt"
	"os"
	"bufio"
	"syscall"
)

func main() {
	filename := "/Users/liuzheng/mytask.retry"
	f, _ := os.Open(filename)
	defer f.Close()
	fr := bufio.NewReader(f)
	status, _ := f.Stat()
	tmp := status
	var fstat *syscall.Stat_t
	for {
		//fstat := reflect.ValueOf(status.Sys()).Elem().FieldByName("Ino").Uint()
		if status == nil {
			ffs, _ := os.Stat(filename)
			if ffs == nil {
				continue
			}
			f.Close()
			f, _ = os.Open(filename)
			defer f.Close()
			fr = bufio.NewReader(f)
			status, _ = f.Stat()
		} else {
			fstat, _ = status.Sys().(*syscall.Stat_t)
			ffstat := fstat
			ffs, ok := os.Stat(filename)
			if ok == nil {
				ffstat, _ = ffs.Sys().(*syscall.Stat_t)
			}

			line, _, err := fr.ReadLine()
			if err != nil {
				fmt.Errorf("Err :%v", err)
			}
			if len(line) > 0 {
				fmt.Println(string(line))
			}
			if status.Size() < tmp.Size() {
				f.Seek(0, os.SEEK_SET)
				fr.Reset(f)
				fmt.Println(status)
			}

			if fstat.Ino != ffstat.Ino {
				// file changed
				f.Close()
				f, _ := os.Open(filename)
				defer f.Close()
				fr = bufio.NewReader(f)
			}
			tmp = status
			status, _ = f.Stat()

		}
	}

}
