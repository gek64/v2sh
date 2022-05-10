package main

import (
	"gek_app"
	"runtime"
)

var (
	app         App
	config      gek_app.Config
	resources   Res
	service     gek_app.Service
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
		a.Application = gek_app.NewApplication(cc.Application.File, "", cc.Application.Location, cc.Application.UninstallDeleteLocation)
	} else {
		a.Application, err = gek_app.NewApplicationFromGithub(cc.Application.File, cc.Application.Repo, cc.Application.RepoList, cc.Application.Location, cc.Application.UninstallDeleteLocation, tempFolder)
		if err != nil {
			return err
		}
	}

	// 配置初始化
	*c = gek_app.NewConfig(cc.Config.Name, cc.Config.Content, cc.Config.Location, cc.Config.UninstallDeleteLocation)

	// 资源初始化
	r.Resources = gek_app.NewResources(cc.Resources.File, cc.Resources.URL, cc.Resources.Location, cc.Application.UninstallDeleteLocation)

	// 服务初始化
	switch runtime.GOOS {
	case gek_app.SupportedOS[0]:
		bytes, err := container.ReadFile("service/v2ray.service")
		if err != nil {
			return err
		}
		*s = gek_app.NewService(cc.Service.Name, string(bytes))
	case gek_app.SupportedOS[1]:
		bytes, err := container.ReadFile("service/v2ray")
		if err != nil {
			return err
		}
		*s = gek_app.NewService(cc.Service.Name, string(bytes))
	}

	return nil
}
