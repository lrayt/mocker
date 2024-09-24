package main

import (
	"github.com/lrayt/mocker/mocker"
	"log"
	"path/filepath"
)

func main() {
	workDir, err := filepath.Abs("")
	if err != nil {
		log.Fatalf("获取项目工作路径失败,err:%s\n", err.Error())
	}
	svr, err := mocker.NewMServer(workDir)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if err := svr.Setup(); err != nil {
		log.Fatalln(err.Error())
	}
}
