package main

import (
	"fmt"
	"gek_app"
	"gek_exec"
	"os/exec"
	"path/filepath"
	"runtime"
)

type Application struct {
	gek_app.Application
}

// 测试应用
func (a Application) test(configFile string) (err error) {
	// 查看app是否存在
	exist, app, _ := gek_exec.Exist(filepath.Join(a.Location, a.AppFiles[0]))
	if !exist {
		return fmt.Errorf("can not find app")
	}

	// 分系统运行不同的命令
	var c string
	switch runtime.GOOS {
	case gek_app.SupportedOS[0]:
		c = linuxConfigLocation
	case gek_app.SupportedOS[1]:
		c = freebsdConfigLocation
	}

	if configFile != "" {
		err = gek_exec.Run(exec.Command(app, "-test", "-config", configFile))
	} else {
		err = gek_exec.Run(exec.Command(app, "-test", "-confdir", c))
	}

	return err
}
