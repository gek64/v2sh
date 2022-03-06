package main

import (
	"gek_app"
	"runtime"
)

var (
	app       Application
	config    gek_app.Config
	resources Resources
	service   gek_app.Service
)

func initNetwork() (err error) {
	var a = &app
	var c = &config
	var r = &resources
	var s = &service

	switch runtime.GOOS {
	case gek_app.SupportedOS[0]:
		// 应用初始化
		a.Application, err = gek_app.NewApplicationFromGithub(linuxBins, linuxRepo, linuxRepoList, linuxNeedExtract, linuxBinsLocation)
		if err != nil {
			return err
		}
		// 配置初始化
		*c = gek_app.NewConfig(linuxConfigName, linuxConfigContent, linuxConfigLocation)

		// 资源初始化
		r.Resources = gek_app.NewResources(linuxResources, linuxResourcesUrl, linuxResourcesLocation)

		// 服务初始化
		*s = gek_app.NewService(linuxServiceName, linuxServiceContent)

	case gek_app.SupportedOS[1]:
		// 应用初始化
		a.Application, err = gek_app.NewApplicationFromGithub(freebsdBins, freebsdRepo, freebsdRepoList, freebsdNeedExtract, freebsdBinsLocation)
		if err != nil {
			return err
		}
		// 配置初始化
		*c = gek_app.NewConfig(freebsdConfigName, freebsdConfigContent, freebsdConfigLocation)

		// 资源初始化
		r.Resources = gek_app.NewResources(freebsdResources, freebsdResourcesUrl, freebsdResourcesLocation)

		// 服务初始化
		*s = gek_app.NewService(freebsdServiceName, freebsdServiceContent)
	}

	return nil
}

func initLocal() {
	var a = &app
	var c = &config
	var r = &resources
	var s = &service

	switch runtime.GOOS {
	case gek_app.SupportedOS[0]:
		// 应用初始化
		a.Application = gek_app.NewApplication(linuxBins, "", linuxNeedExtract, linuxBinsLocation)

		// 配置初始化
		*c = gek_app.NewConfig(linuxConfigName, linuxConfigContent, linuxConfigLocation)

		// 资源初始化
		r.Resources = gek_app.NewResources(linuxResources, linuxResourcesUrl, linuxResourcesLocation)

		// 服务初始化
		*s = gek_app.NewService(linuxServiceName, linuxServiceContent)

	case gek_app.SupportedOS[1]:
		// 应用初始化
		a.Application = gek_app.NewApplication(freebsdBins, "", freebsdNeedExtract, freebsdBinsLocation)

		// 配置初始化
		*c = gek_app.NewConfig(freebsdConfigName, freebsdConfigContent, freebsdConfigLocation)

		// 资源初始化
		r.Resources = gek_app.NewResources(freebsdResources, freebsdResourcesUrl, freebsdResourcesLocation)

		// 服务初始化
		*s = gek_app.NewService(freebsdServiceName, freebsdServiceContent)
	}
}
