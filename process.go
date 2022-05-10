package main

import (
	"fmt"
	"gek_app"
	"log"
	"os"
)

// 安装
func install(configFile string) (err error) {
	// 配置装载内容
	if configFile != "" {
		config, err = gek_app.NewConfigFromFile(cc.Config.Name, configFile, cc.Config.UninstallDeleteLocation, cc.Config.Location)
		if err != nil {
			return err
		}
	}

	// 应用安装
	err = app.Install(tempFolder, true, needExtract)
	if err != nil {
		return err
	}
	// 配置安装
	err = config.Install()
	if err != nil {
		return err
	}
	// 资源安装,从互联网安装时使用resources.Install()
	err = resources.InstallFromLocation(tempFolder)
	if err != nil {
		return err
	}
	// 服务安装
	err = service.Install()
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
		config, err = gek_app.NewConfigFromFile(cc.Config.Name, configFile, cc.Config.UninstallDeleteLocation, cc.Config.Location)
		if err != nil {
			return err
		}
	}

	// 应用安装
	err = app.InstallFromLocal(tempFolder, localFile, needExtract)
	if err != nil {
		return err
	}
	// 配置安装
	err = config.Install()
	if err != nil {
		return err
	}
	// 资源安装
	err = resources.installFromLocalArchiveFile(localFile)
	if err != nil {
		return err
	}
	// 服务安装
	err = service.Install()
	if err != nil {
		return err
	}

	return nil
}

// 卸载
func uninstall() (err error) {
	// 停止应用+卸载服务
	err = service.Uninstall()
	if err != nil {
		log.Println(err)
	}
	// 卸载配置文件
	err = config.Uninstall()
	if err != nil {
		log.Println(err)
	}
	// 卸载资源
	err = resources.Uninstall()
	if err != nil {
		log.Println(err)
	}
	// 卸载应用
	err = app.Uninstall()
	if err != nil {
		log.Println(err)
	}

	return nil
}

// 更新
func update(configFile string) (err error) {
	if configFile != "" {
		// 配置装载内容
		config, err = gek_app.NewConfigFromFile(cc.Config.Name, configFile, cc.Config.UninstallDeleteLocation, cc.Config.Location)
		if err != nil {
			return err
		}
		// 配置卸载
		err = config.Uninstall()
		if err != nil {
			return err
		}
		// 配置安装
		err = config.Install()
		if err != nil {
			return err
		}
	}

	// 应用卸载
	err = app.Uninstall()
	if err != nil {
		return err
	}
	// 应用安装
	err = app.Install(tempFolder, true, needExtract)
	if err != nil {
		return err
	}

	// 资源卸载
	err = resources.Uninstall()
	if err != nil {
		return err
	}
	// 资源安装
	err = resources.InstallFromInternet()
	if err != nil {
		return err
	}

	// 服务重启
	err = service.Restart()
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
		// 配置装载内容
		config, err = gek_app.NewConfigFromFile(cc.Config.Name, configFile, cc.Config.UninstallDeleteLocation, cc.Config.Location)
		if err != nil {
			return err
		}
		// 配置卸载
		err = config.Uninstall()
		if err != nil {
			return err
		}
		// 配置安装
		err = config.Install()
		if err != nil {
			return err
		}
	}

	// 应用卸载
	err = app.Uninstall()
	if err != nil {
		return err
	}
	// 应用安装
	err = app.InstallFromLocal(tempFolder, localFile, needExtract)
	if err != nil {
		return err
	}

	// 资源卸载
	err = resources.Uninstall()
	if err != nil {
		return err
	}
	// 资源安装
	err = resources.installFromLocalArchiveFile(localFile)
	if err != nil {
		return err
	}

	// 服务重启
	err = service.Restart()
	if err != nil {
		return err
	}

	return nil
}

// 重载
func reload(configFile string) (err error) {
	if configFile != "" {
		// 配置装载内容
		config, err = gek_app.NewConfigFromFile(cc.Config.Name, configFile, cc.Config.UninstallDeleteLocation, cc.Config.Location)
		if err != nil {
			return err
		}
		// 配置卸载
		err = config.Uninstall()
		if err != nil {
			return err
		}
		// 配置安装
		err = config.Install()
		if err != nil {
			return err
		}
	}

	// 服务重启
	err = service.Restart()
	if err != nil {
		return err
	}

	return nil
}

// 测试
func test(configFile string) (err error) {
	return app.test(configFile)
}
