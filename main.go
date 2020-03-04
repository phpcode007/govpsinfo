package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var str string

func main() {

	resp, err := http.Get("http://ifconfig.co/ip")

	if err != nil {
		log.Fatal("获取ip地址出错")
	}

	defer resp.Body.Close()

	s,err:=ioutil.ReadAll(resp.Body)


	// 去除换行符
	str = strings.Replace(string(s[:]), "\n", "", -1)

	fmt.Print("在浏览器打开 http://" + string(str) + ":19002")
	// 设置需要释放的目录
	dirs := []string{"shell"}

	isSuccess := true

	for _, dir := range dirs {
		// 解压shell目录到/tmp目录
		if err := RestoreAssets("/tmp", dir); err != nil {
			isSuccess = false
			break
		}
	}
	if !isSuccess {
		for _, dir := range dirs {
			os.RemoveAll(filepath.Join("/tmp", dir))
		}
	}

	http.HandleFunc("/", SystemInfo)
	//浏览器自带的请求,不响应
	http.HandleFunc("/favicon.ico", func(writer http.ResponseWriter, request *http.Request) {
		
	})

	err = http.ListenAndServe(":19002", nil)

	if err != nil {
		log.Fatal("启动服务器出错"+ err.Error())
	}

}

func SystemInfo(w http.ResponseWriter, req *http.Request) {

	var (
		cmd *exec.Cmd
		output []byte
		err error
	)

	shellpath := "/tmp/shell/getsysteminfo.sh"

	cmd = exec.Command("/bin/bash",shellpath,string(str))

	// 执行了命令, 捕获了子进程的输出( pipe )
	if output, err = cmd.Output(); err != nil {
		fmt.Println(err.Error())
		return
	}

	// 打印子进程的输出
	fmt.Println(string(output))

	w.Write(output)
}















