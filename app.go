package main

import (
	"fmt"
	"github.com/gek64/gek/gApp"
	"github.com/gek64/gek/gExec"
	"os/exec"
	"path/filepath"
)

type App struct {
	gApp.Application
}

// 测试应用
func (a App) test(configFile string) (err error) {
	// 查看app是否存在
	exist, app, _ := gExec.Exist(filepath.Join(a.Location, a.AppFiles[0]))
	if !exist {
		return fmt.Errorf("can not find app")
	}

	// 分系统运行不同的命令
	if configFile != "" {
		err = gExec.Run(exec.Command(app, "test", "-config", configFile))
	} else {
		err = gExec.Run(exec.Command(app, "test", "-confdir", cc.Config.Location))
	}

	return err
}
