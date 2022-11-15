package main

import (
	"github.com/gek64/gek/gApp"
	"runtime"
)

var (
	app         App
	config      gApp.Config
	resources   Res
	service     gApp.Service
	tempFolder  = "/tmp/proxy"
	needExtract = true
	cc          CC
)

func initApp(local bool) (err error) {
	var a = &app
	var c = &config
	var r = &resources
	var s = &service

	err = initConf()
	if err != nil {
		return err
	}

	// 应用初始化
	if local {
		a.Application = gApp.NewApplication(cc.Application.File, "", cc.Application.Location, cc.Application.UninstallDeleteLocation)
	} else {
		a.Application, err = gApp.NewApplicationFromGithub(cc.Application.File, cc.Application.Repo, cc.Application.RepoList, cc.Application.Location, cc.Application.UninstallDeleteLocation, tempFolder)
		if err != nil {
			return err
		}
	}

	// 配置初始化
	*c = gApp.NewConfig(cc.Config.Name, cc.Config.Content, cc.Config.Location, cc.Config.UninstallDeleteLocation)

	// 资源初始化
	r.Resources = gApp.NewResources(cc.Resources.File, cc.Resources.URL, cc.Resources.Location, cc.Application.UninstallDeleteLocation)

	// 服务初始化
	switch runtime.GOOS {
	case gApp.SupportedOS[0]:
		bytes, err := container.ReadFile("configs/v2ray.service")
		if err != nil {
			return err
		}
		*s = gApp.NewService(cc.Service.Name, string(bytes))
	case gApp.SupportedOS[1]:
		bytes, err := container.ReadFile("configs/v2ray")
		if err != nil {
			return err
		}
		*s = gApp.NewService(cc.Service.Name, string(bytes))
	}

	return nil
}
