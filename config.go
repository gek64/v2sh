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
		baseURL = "https://github.com/XTLS/Xray-core/releases/download/" + tagName + "/"
	} else {
		baseURL = "https://github.com/XTLS/Xray-core/releases/latest/download/"
	}

	switch runtime.GOOS {
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			assetName = "Xray-linux-64.zip"
		case "386":
			assetName = "Xray-linux-32.zip"
		case "arm":
			assetName = "Xray-linux-arm32-v7a.zip"
		case "arm64":
			assetName = "Xray-linux-arm64-v8a.zip"
		case "loong64":
			assetName = "Xray-linux-loong64.zip"
		case "mips":
			assetName = "Xray-linux-mips32.zip"
		case "mipsle":
			assetName = "Xray-linux-mips32le.zip"
		case "mips64":
			assetName = "Xray-linux-mips64.zip"
		case "mips64le":
			assetName = "Xray-linux-mips64le.zip"
		case "riscv64":
			assetName = "Xray-linux-riscv64.zip"
		}
	case "windows":
		switch runtime.GOARCH {
		case "amd64":
			assetName = "Xray-windows-64.zip"
		case "386":
			assetName = "Xray-windows-32.zip"
		case "arm":
			assetName = "Xray-windows-arm32-v7a.zip"
		case "arm64":
			assetName = "Xray-windows-arm64-v8a.zip"
		}
	case "darwin":
		switch runtime.GOARCH {
		case "amd64":
			assetName = "Xray-macos-64.zip"
		case "arm64":
			assetName = "Xray-macos-arm64-v8a.zip"
		}
	case "freebsd":
		switch runtime.GOARCH {
		case "amd64":
			assetName = "Xray-freebsd-64.zip"
		case "386":
			assetName = "Xray-freebsd-32.zip"
		case "arm":
			assetName = "Xray-freebsd-arm32-v7a.zip"
		case "arm64":
			assetName = "Xray-freebsd-arm64-v8a.zip"
		}
	case "openbsd":
		switch runtime.GOARCH {
		case "amd64":
			assetName = "Xray-openbsd-64.zip"
		case "386":
			assetName = "Xray-openbsd-32.zip"
		case "arm":
			assetName = "Xray-openbsd-arm32-v7a.zip"
		case "arm64":
			assetName = "Xray-openbsd-arm64-v8a.zip"
		}
	case "dragonfly":
		switch runtime.GOARCH {
		case "amd64":
			assetName = "Xray-dragonfly-64.zip"
		}
	case "android":
		switch runtime.GOARCH {
		case "arm64":
			assetName = "Xray-android-arm64-v8a.zip"
		}
	}
	return url.JoinPath(baseURL, assetName)
}

func GetService() (initSystem string, serviceContent []byte, err error) {
	serviceFile := ""
	// systemd
	_, err = exec.LookPath("systemctl")
	if err == nil {
		serviceFile = "configs/xray.service"
		initSystem = "systemd"
	}
	// alpine openrc
	_, err = exec.LookPath("openrc")
	if err == nil {
		serviceFile = "configs/xray.openrc"
		initSystem = "openrc"
	}
	// openwrt procd
	_, err = exec.LookPath("opkg")
	if err == nil {
		serviceFile = "configs/xray.procd"
		initSystem = "procd"
	}
	// freebsd rc.d
	_, err = exec.LookPath("rcorder")
	if err == nil {
		serviceFile = "configs/xray.rcd"
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
