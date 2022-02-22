package main

import (
	"runtime"
)

var (
	app       Application
	config    Config
	resources Resources
	service   Service
)

func initNetwork() (err error) {
	var a = &app
	var c = &config
	var r = &resources
	var s = &service

	switch runtime.GOOS {
	case supportedOS[0]:
		// 应用初始化
		*a, err = newApplicationFromGithub(linuxBins, linuxRepo, linuxRepoList, linuxNeedExtract, linuxBinsLocation)
		if err != nil {
			return err
		}
		// 配置初始化
		*c = newConfig(linuxConfigName, linuxConfigContent, linuxConfigLocation)

		// 资源初始化
		*r = newResources(linuxResources, linuxResourcesUrl, linuxResourcesLocation)

		// 服务初始化
		*s = newService(linuxServiceName, linuxServiceContent)

	case supportedOS[1]:
		// 应用初始化
		*a, err = newApplicationFromGithub(freebsdBins, freebsdRepo, freebsdRepoList, freebsdNeedExtract, freebsdBinsLocation)
		if err != nil {
			return err
		}
		// 配置初始化
		*c = newConfig(freebsdConfigName, freebsdConfigContent, freebsdConfigLocation)

		// 资源初始化
		*r = newResources(freebsdResources, freebsdResourcesUrl, freebsdResourcesLocation)

		// 服务初始化
		*s = newService(freebsdServiceName, freebsdServiceContent)
	}

	return nil
}

func initLocal() {
	var a = &app
	var c = &config
	var r = &resources
	var s = &service

	switch runtime.GOOS {
	case supportedOS[0]:
		// 应用初始化
		*a = newApplication(linuxBins, "", linuxNeedExtract, linuxBinsLocation)

		// 配置初始化
		*c = newConfig(linuxConfigName, linuxConfigContent, linuxConfigLocation)

		// 资源初始化
		*r = newResources(linuxResources, linuxResourcesUrl, linuxResourcesLocation)

		// 服务初始化
		*s = newService(linuxServiceName, linuxServiceContent)

	case supportedOS[1]:
		// 应用初始化
		*a = newApplication(freebsdBins, "", freebsdNeedExtract, freebsdBinsLocation)

		// 配置初始化
		*c = newConfig(freebsdConfigName, freebsdConfigContent, freebsdConfigLocation)

		// 资源初始化
		*r = newResources(freebsdResources, freebsdResourcesUrl, freebsdResourcesLocation)

		// 服务初始化
		*s = newService(freebsdServiceName, freebsdServiceContent)
	}
}
