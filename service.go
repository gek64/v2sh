package main

import (
	"gek_service"
	"gek_service_freebsd"
	"os"
	"runtime"
)

type Service struct {
	Name    string
	Content string
}

// 新建服务
func newService(name string, content string) (s Service) {
	return Service{Name: name, Content: content}
}

// 新建服务(从文件)
func newServiceFromFile(name string, file string) (s Service, err error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return Service{}, err
	}
	return newService(name, string(bytes)), nil
}

// 安装服务
func (s Service) install() (err error) {
	// 分系统运行不同的命令
	switch runtime.GOOS {
	case supportedOS[0]:
		service := gek_service.NewService(s.Name, s.Content)
		// 安装服务
		err = service.Install()
		if err != nil {
			return err
		}
		// 查看服务状态
		err = service.Status()
		if err != nil {
			return err
		}
	case supportedOS[1]:
		service := gek_service_freebsd.NewService(s.Name, s.Content)
		// 安装服务
		err = service.Install()
		if err != nil {
			return err
		}
		// 查看服务状态
		err = service.Status()
		if err != nil {
			return err
		}
	}
	return nil
}

// 卸载服务
func (s Service) uninstall() (err error) {
	// 分系统运行不同的命令
	switch runtime.GOOS {
	case supportedOS[0]:
		service := gek_service.NewService(s.Name, s.Content)
		// 卸载服务
		err = service.Uninstall()
		if err != nil {
			return err
		}
	case supportedOS[1]:
		service := gek_service_freebsd.NewService(s.Name, s.Content)
		// 卸载服务
		err = service.Uninstall()
		if err != nil {
			return err
		}
	}
	return nil
}

// 重启服务
func (s Service) restart() (err error) {
	// 分系统运行不同的命令
	switch runtime.GOOS {
	case supportedOS[0]:
		service := gek_service.NewService(s.Name, s.Content)
		// 重载服务
		err = service.Reload()
		if err != nil {
			return err
		}
		// 查看服务状态
		err = service.Status()
		if err != nil {
			return err
		}
	case supportedOS[1]:
		service := gek_service_freebsd.NewService(s.Name, s.Content)
		// 重载服务
		err = service.Reload()
		if err != nil {
			return err
		}
		// 查看服务状态
		err = service.Status()
		if err != nil {
			return err
		}
	}
	return nil
}
