package datas

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func CreateFile(fileName, data string) {
	if checkFileIsExist(fileName) {
		if f, err := os.OpenFile(fileName, os.O_APPEND, 0666); err != nil {
			log.Println("open file error:", err)
		} else if _, err := io.WriteString(f, data); err != nil {
			log.Println("write file error:", err)
		}
	} else if err := ioutil.WriteFile(fileName, []byte(data), 0666); err != nil {
		log.Println("create file error:", err)
	}
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func NewDir(dirName string) {
	if checkFileIsExist(dirName) {
		return
	}
	var path string
	if os.IsPathSeparator('\\') { //前边的判断是否是系统的分隔符
		path = "\\"
	} else {
		path = "/"
	}
	if dir, err := os.Getwd(); err != nil {
		log.Println("get now dir error:", err)
	} else if err := os.Mkdir(dir+path+dirName, os.ModePerm); err != nil {
		log.Println("create dir error:", err)
	}
}

func StringToFloat(str string) float64 {
	var value float64
	fmt.Sscanf(str, "%f", &value)
	return value
}
