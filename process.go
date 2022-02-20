package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

var (
	app       Application
	config    Config
	resources Resources
	service   Service
)

func init() {
	var a = &app
	var c = &config
	var r = &resources
	var s = &service
	var err error

	switch runtime.GOOS {
	case supportedOS[0]:
		// 应用初始化
		*a, err = newApplicationFromGithub(linuxBins, linuxRepo, linuxRepoList, linuxNeedExtract, linuxBinsLocation)
		if err != nil {
			log.Panicln(err)
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
			log.Panicln(err)
		}
		// 配置初始化
		*c = newConfig(freebsdConfigName, freebsdConfigContent, freebsdConfigLocation)

		// 资源初始化
		*r = newResources(freebsdResources, freebsdResourcesUrl, freebsdResourcesLocation)

		// 服务初始化
		*s = newService(freebsdServiceName, freebsdServiceContent)
	}
}

// 安装
func install(configFile string) (err error) {
	// 配置装载内容
	if configFile != "" {
		switch runtime.GOOS {
		case supportedOS[0]:
			config, err = newConfigFromFile(linuxConfigName, configFile, linuxConfigLocation)
			if err != nil {
				return err
			}
		case supportedOS[1]:
			config, err = newConfigFromFile(freebsdConfigName, configFile, freebsdConfigLocation)
			if err != nil {
				return err
			}
		}
	}

	// 应用安装
	err = app.install()
	if err != nil {
		return err
	}
	// 配置安装
	err = config.install()
	if err != nil {
		return err
	}
	// 资源安装
	err = resources.install()
	if err != nil {
		return err
	}
	// 服务安装
	err = service.install()
	if err != nil {
		return err
	}

	return nil
}

// 安装(从本地文件,无需网络下载)
func installFromLocal(configFile string, localFile string) (err error) {
	// 检查本地文件是否存在
	_, err = os.Stat(localFile)
	if os.IsNotExist(err) {
		return fmt.Errorf("%s is not exist", localFile)
	}

	// 配置装载内容
	if configFile != "" {
		switch runtime.GOOS {
		case supportedOS[0]:
			config, err = newConfigFromFile(linuxConfigName, configFile, linuxConfigLocation)
			if err != nil {
				return err
			}
		case supportedOS[1]:
			config, err = newConfigFromFile(freebsdConfigName, configFile, freebsdConfigLocation)
			if err != nil {
				return err
			}
		}
	}

	// 应用安装
	err = app.installFromLocal(localFile)
	if err != nil {
		return err
	}
	// 配置安装
	err = config.install()
	if err != nil {
		return err
	}
	// 资源安装
	err = resources.installFromLocal(localFile)
	if err != nil {
		return err
	}
	// 服务安装
	err = service.install()
	if err != nil {
		return err
	}

	return nil
}

// 卸载
func uninstall() (err error) {
	// 停止应用+卸载服务
	err = service.uninstall()
	if err != nil {
		log.Println(err)
	}
	// 卸载配置文件
	err = config.uninstall()
	if err != nil {
		log.Println(err)
	}
	// 卸载资源
	err = resources.uninstall()
	if err != nil {
		log.Println(err)
	}
	// 卸载应用
	err = app.uninstall()
	if err != nil {
		log.Println(err)
	}

	return nil
}

// 更新
func update(configFile string) (err error) {
	if configFile != "" {
		switch runtime.GOOS {
		case supportedOS[0]:
			// 配置装载内容
			config, err = newConfigFromFile(linuxConfigName, configFile, linuxConfigLocation)
			if err != nil {
				return err
			}
		case supportedOS[1]:
			// 配置装载内容
			config, err = newConfigFromFile(freebsdConfigName, configFile, freebsdConfigLocation)
			if err != nil {
				return err
			}
		}
		// 配置卸载
		err = config.uninstall()
		if err != nil {
			return err
		}
		// 配置安装
		err = config.install()
		if err != nil {
			return err
		}
	}

	// 应用卸载
	err = app.uninstall()
	if err != nil {
		return err
	}
	// 应用安装
	err = app.install()
	if err != nil {
		return err
	}

	// 资源卸载
	err = resources.uninstall()
	if err != nil {
		return err
	}
	// 资源安装
	err = resources.install()
	if err != nil {
		return err
	}

	// 服务重启
	err = service.restart()
	if err != nil {
		return err
	}

	return nil
}

// 更新(从本地文件,无需网络下载)
func updateFromLocal(configFile string, localFile string) (err error) {
	// 检查本地文件是否存在
	_, err = os.Stat(localFile)
	if os.IsNotExist(err) {
		return fmt.Errorf("%s is not exist", localFile)
	}

	if configFile != "" {
		switch runtime.GOOS {
		case supportedOS[0]:
			// 配置装载内容
			config, err = newConfigFromFile(linuxConfigName, configFile, linuxConfigLocation)
			if err != nil {
				return err
			}
		case supportedOS[1]:
			// 配置装载内容
			config, err = newConfigFromFile(freebsdConfigName, configFile, freebsdConfigLocation)
			if err != nil {
				return err
			}
		}
		// 配置卸载
		err = config.uninstall()
		if err != nil {
			return err
		}
		// 配置安装
		err = config.install()
		if err != nil {
			return err
		}
	}

	// 应用卸载
	err = app.uninstall()
	if err != nil {
		return err
	}
	// 应用安装
	err = app.installFromLocal(localFile)
	if err != nil {
		return err
	}

	// 资源卸载
	err = resources.uninstall()
	if err != nil {
		return err
	}
	// 资源安装
	err = resources.installFromLocal(localFile)
	if err != nil {
		return err
	}

	// 服务重启
	err = service.restart()
	if err != nil {
		return err
	}

	return nil
}

// 重载
func reload(configFile string) (err error) {
	if configFile != "" {
		switch runtime.GOOS {
		case supportedOS[0]:
			// 配置装载内容
			config, err = newConfigFromFile(linuxConfigName, configFile, linuxConfigLocation)
			if err != nil {
				return err
			}
		case supportedOS[1]:
			// 配置装载内容
			config, err = newConfigFromFile(freebsdConfigName, configFile, freebsdConfigLocation)
			if err != nil {
				return err
			}
		}
		// 配置卸载
		err = config.uninstall()
		if err != nil {
			return err
		}
		// 配置安装
		err = config.install()
		if err != nil {
			return err
		}
	}

	// 服务重启
	err = service.restart()
	if err != nil {
		return err
	}

	return nil
}

// 测试
func test(configFile string) (err error) {
	return app.test(configFile)
}
