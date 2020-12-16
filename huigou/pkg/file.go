package pkg

import (
	"os"
)

//查看文件/文件夹是否存在
func FileOrDirIsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil{
		if os.IsExist(err){
			return true
		}
		return false
	}
	return true
}
//创建文件
func MakeFile(filename string) bool {
	if !FileOrDirIsExist(filename) {
		file, err := os.Create(filename)
		defer file.Close()
		if err != nil{
			return false
		}
	}
	return true
}