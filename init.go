package main

import (
	"gek_app"
	"runtime"
)

var (
	app         Application
	config      gek_app.Config
	resources   Resources
	service     gek_app.Service
	tempFolder  string = "/tmp/proxy"
	needExtract bool   = true
)

func initNetwork() (err error) {
	var a = &app
	var c = &config
	var r = &resources
	var s = &service

	switch runtime.GOOS {
	case gek_app.SupportedOS[0]:
		// 应用初始化
		a.Application, err = gek_app.NewApplicationFromGithub(linuxBins, linuxRepo, linuxRepoList, linuxBinsLocation, linuxBinsUninstallDeleteLoaction, tempFolder)
		if err != nil {
			return err
		}
		// 配置初始化
		*c = gek_app.NewConfig(linuxConfigName, linuxConfigContent, linuxConfigLocation, linuxConfigUninstallDeleteLoaction)

		// 资源初始化
		r.Resources = gek_app.NewResources(linuxResources, linuxResourcesUrl, linuxResourcesLocation, linuxResourcesUninstallDeleteLoaction)

		// 服务初始化
		*s = gek_app.NewService(linuxServiceName, linuxServiceContent)

	case gek_app.SupportedOS[1]:
		// 应用初始化
		a.Application, err = gek_app.NewApplicationFromGithub(freebsdBins, freebsdRepo, freebsdRepoList, freebsdBinsLocation, freebsdBinsUninstallDeleteLoaction, tempFolder)
		if err != nil {
			return err
		}
		// 配置初始化
		*c = gek_app.NewConfig(freebsdConfigName, freebsdConfigContent, freebsdConfigLocation, freebsdConfigUninstallDeleteLoaction)

		// 资源初始化
		r.Resources = gek_app.NewResources(freebsdResources, freebsdResourcesUrl, freebsdResourcesLocation, freebsdResourcesUninstallDeleteLoaction)

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
		a.Application = gek_app.NewApplication(linuxBins, "", linuxBinsLocation, linuxBinsUninstallDeleteLoaction)

		// 配置初始化
		*c = gek_app.NewConfig(linuxConfigName, linuxConfigContent, linuxConfigLocation, linuxConfigUninstallDeleteLoaction)

		// 资源初始化
		r.Resources = gek_app.NewResources(linuxResources, linuxResourcesUrl, linuxResourcesLocation, linuxResourcesUninstallDeleteLoaction)

		// 服务初始化
		*s = gek_app.NewService(linuxServiceName, linuxServiceContent)

	case gek_app.SupportedOS[1]:
		// 应用初始化
		a.Application = gek_app.NewApplication(freebsdBins, "", freebsdBinsLocation, freebsdBinsUninstallDeleteLoaction)

		// 配置初始化
		*c = gek_app.NewConfig(freebsdConfigName, freebsdConfigContent, freebsdConfigLocation, freebsdConfigUninstallDeleteLoaction)

		// 资源初始化
		r.Resources = gek_app.NewResources(freebsdResources, freebsdResourcesUrl, freebsdResourcesLocation, freebsdResourcesUninstallDeleteLoaction)

		// 服务初始化
		*s = gek_app.NewService(freebsdServiceName, freebsdServiceContent)
	}
}
