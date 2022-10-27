package utils

import (
	"bufio"
	"io/ioutil"
	"os"
)

// PathExists 判断路径是否存在
func PathExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		return false
	}
	return info.IsDir()
}

// FileExists 判断文件是否存在
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) || err != nil || info == nil {
		return false
	}
	return !info.IsDir()
}

func ReadLines(filename string) (lines []string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return
}

func ReadFile(filename string) (data []byte, err error) {
	data, err = ioutil.ReadFile(filename)
	return
}

func WriteFile(filename string, data string) (err error) {
	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer f.Close()
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0666)
	writer := bufio.NewWriter(file)
	_, _ = writer.WriteString(data)
	_ = writer.Flush()
	return
}
