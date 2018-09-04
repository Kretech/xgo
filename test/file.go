package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

var fileCache map[string][]byte

func init() {
	fileCache = make(map[string][]byte)
}

func readFile(filename string) ([]byte, error) {
	file := fileCache[filename]
	if file != nil && len(file) > 0 {
		return file, nil
	}

	file, err := ioutil.ReadFile(filename)

	fileCache[filename] = file

	return file, err
}

func showFile(filename string, line int) string {
	file, _ := readFile(filename)

	buf := bytes.NewBufferString(``)

	size := len(file)
	begin := line - 2
	end := line + 3

	//	@todo line超过文件时
	found := 1
	for idx := 0; idx < size; idx++ {
		if found >= begin && found < end {
			buf.WriteByte(file[idx])
		}

		if file[idx] == '\n' {
			found++

			if found >= begin && found < end {
				tag := ' '
				if found == line {
					tag = '→'
				}
				buf.WriteString(fmt.Sprintf("%2c%4d |\t", tag, found))
			}

			if found >= end {
				break
			}
		}
	}

	return buf.String()
}
