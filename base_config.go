package main

import (
	"fmt"
	"gek_file"
	"os"
	"path/filepath"
)

type Config struct {
	// 配置文件名称
	Name string
	// 配置文件内容
	Content string
	// 配置文件安装文件夹
	Location string
}

// 新建配置文件
func newConfig(name string, content string, location string) (c Config) {
	return Config{Name: name, Content: content, Location: location}
}

// 新建配置文件(从文件)
func newConfigFromFile(name string, file string, location string) (c Config, err error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return Config{}, err
	}
	return newConfig(name, string(bytes), location), nil
}

// 安装配置文件
func (c Config) install() (err error) {
	// 检测配置安装路径是否存在
	// 不存在则创建
	_, err = os.Stat(c.Location)
	if os.IsNotExist(err) {
		err = os.MkdirAll(c.Location, 755)
		if err != nil {
			return err
		}
	}

	_, err = gek_file.CreateFile(filepath.Join(c.Location, c.Name), c.Content)
	if err != nil {
		return err
	}

	// 赋权755
	err = chmodRecursive(c.Location, 755)
	if err != nil {
		return err
	}

	return nil
}

// 卸载配置文件,并删除配置文件安装文件夹
func (c Config) uninstall() (err error) {
	// 检测配置安装路径是否存在
	_, err = os.Stat(c.Location)
	if os.IsNotExist(err) {
		return fmt.Errorf("can't find config location %s", c.Location)
	}
	// 删除配置安装文件夹,及所有配置文件
	err = os.RemoveAll(c.Location)
	if err != nil {
		return err
	}
	return nil
}
