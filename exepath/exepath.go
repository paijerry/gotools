//Package exepath - 取得build出的exe檔路徑，可用來讀取設定檔
//參考：http://studygolang.com/topics/52
package exepath

import (
	"os"
	"os/exec"
	"path"
)

//Get path for builded exe (can't used for run)
func Get() string {
	// file, _ := os.Getwd()
	// log.Println("current path:", file)

	file, _ := exec.LookPath(os.Args[0])
	//log.Println("exec path:", file)

	dir, _ := path.Split(file)
	//log.Println("exec folder relative path:", dir)

	os.Chdir(dir)
	wd, _ := os.Getwd()
	//log.Println("exec folder absolute path:", wd)

	return wd
}
