package main

import (
	"embed"
	"errors"
	"net/url"
	"os/exec"
	"runtime"
)

//go:embed configs/*
var container embed.FS

func GetDownloadURL(tagName string) (downloadURL string, err error) {
	var baseURL = ""
	var assetName = ""

	if tagName != "" {
		baseURL = "https://github.com/v2fly/v2ray-core/releases/download/" + tagName + "/"
	} else {
		baseURL = "https://github.com/v2fly/v2ray-core/releases/latest/download/"
	}

	switch runtime.GOOS {
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			assetName = "v2ray-linux-64.zip"
		case "386":
			assetName = "v2ray-linux-32.zip"
		case "arm":
			assetName = "v2ray-linux-arm32-v7a.zip"
		case "arm64":
			assetName = "v2ray-linux-arm64-v8a.zip"
		case "loong64":
			assetName = "v2ray-linux-loong64.zip"
		case "mips":
			assetName = "v2ray-linux-mips32.zip"
		case "mipsle":
			assetName = "v2ray-linux-mips32le.zip"
		case "mips64":
			assetName = "v2ray-linux-mips64.zip"
		case "mips64le":
			assetName = "v2ray-linux-mips64le.zip"
		case "riscv64":
			assetName = "v2ray-linux-riscv64.zip"
		}
	case "windows":
		switch runtime.GOARCH {
		case "amd64":
			assetName = "v2ray-windows-64.zip"
		case "386":
			assetName = "v2ray-windows-32.zip"
		case "arm":
			assetName = "v2ray-windows-arm32-v7a.zip"
		case "arm64":
			assetName = "v2ray-windows-arm64-v8a.zip"
		}
	case "darwin":
		switch runtime.GOARCH {
		case "amd64":
			assetName = "v2ray-macos-64.zip"
		case "arm64":
			assetName = "v2ray-macos-arm64-v8a.zip"
		}
	case "freebsd":
		switch runtime.GOARCH {
		case "amd64":
			assetName = "v2ray-freebsd-64.zip"
		case "386":
			assetName = "v2ray-freebsd-32.zip"
		case "arm":
			assetName = "v2ray-freebsd-arm32-v7a.zip"
		case "arm64":
			assetName = "v2ray-freebsd-arm64-v8a.zip"
		}
	case "openbsd":
		switch runtime.GOARCH {
		case "amd64":
			assetName = "v2ray-openbsd-64.zip"
		case "386":
			assetName = "v2ray-openbsd-32.zip"
		case "arm":
			assetName = "v2ray-openbsd-arm32-v7a.zip"
		case "arm64":
			assetName = "v2ray-openbsd-arm64-v8a.zip"
		}
	case "dragonfly":
		switch runtime.GOARCH {
		case "amd64":
			assetName = "v2ray-dragonfly-64.zip"
		}
	case "android":
		switch runtime.GOARCH {
		case "arm64":
			assetName = "v2ray-android-arm64-v8a.zip"
		}
	}
	return url.JoinPath(baseURL, assetName)
}

func GetService() (initSystem string, serviceContent []byte, err error) {
	serviceFile := ""
	// systemd
	_, err = exec.LookPath("systemctl")
	if err == nil {
		serviceFile = "configs/v2ray.service"
		initSystem = "systemd"
	}
	// alpine openrc
	_, err = exec.LookPath("openrc")
	if err == nil {
		serviceFile = "configs/v2ray.openrc"
		initSystem = "openrc"
	}
	// openwrt procd
	_, err = exec.LookPath("opkg")
	if err == nil {
		serviceFile = "configs/v2ray.procd"
		initSystem = "procd"
	}
	// freebsd rc.d
	_, err = exec.LookPath("rcorder")
	if err == nil {
		serviceFile = "configs/v2ray.rcd"
		initSystem = "rc.d"
	}
	// 找不到初始化系统返回错误
	if initSystem == "" {
		return "", nil, errors.New("init system not found")
	}
	// 读取文件并返回文件内容
	bytes, err := container.ReadFile(serviceFile)
	if err != nil {
		return initSystem, nil, err
	}
	return initSystem, bytes, nil
}
