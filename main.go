package main

import (
	"flag"
	"fmt"
	"gek_toolbox"
	"log"
	"os"
	"runtime"
)

var (
	cliInstall   bool
	cliUninstall bool
	cliUpdate    bool
	cliReload    bool
	cliTest      bool
	cliLocalFile string
	cliConfig    string
	cliHelp      bool
	cliVersion   bool
)

func init() {
	flag.BoolVar(&cliInstall, "I", false, "install")
	flag.BoolVar(&cliUninstall, "R", false, "uninstall")
	flag.BoolVar(&cliUpdate, "U", false, "update")
	flag.BoolVar(&cliReload, "r", false, "reload")
	flag.BoolVar(&cliTest, "t", false, "test config")
	flag.StringVar(&cliLocalFile, "l", "", "use local file without download from network")
	flag.StringVar(&cliConfig, "c", "", "use local config")
	flag.BoolVar(&cliHelp, "h", false, "show help")
	flag.BoolVar(&cliVersion, "v", false, "show version")
	flag.Parse()

	// 重写显示用法函数
	flag.Usage = func() {
		var helpInfo = `Usage:
    proxyctl [Command] [Arguments]

Arguments:
	-l <bin.zip>      : use local file without download from network
	-c <config.json>  : use local config

Command:
	-I                : install
	-R                : uninstall
	-U                : update
	-r                : reload
	-t                : test config
	-h                : show help
	-v                : show version

Example:
1) proxyctl -I -l bins.zip -c config.json   : Install proxy and service
2) proxyctl -U -l bins.zip -c config.json   : Update proxy and resources
3) proxyctl -R                              : Uninstall proxy and service
4) proxyctl -t -c config.json               : Test config
5) proxyctl -r -c config.json               : Reload service`
		fmt.Println(helpInfo)
	}

	// 如果无 args 或者 指定 h 参数,打印用法后退出
	if len(os.Args) == 1 || cliHelp {
		flag.Usage()
		os.Exit(0)
	}

	// 打印版本信息
	if cliVersion {
		fmt.Println("v2.00")
		os.Exit(0)
	}

	// 检查运行库是否完整
	switch runtime.GOOS {
	case supportedOS[0]:
		err := gek_toolbox.CheckToolbox(linuxToolbox)
		if err != nil {
			log.Fatal(err)
		}
	case supportedOS[1]:
		err := gek_toolbox.CheckToolbox(freebsdToolbox)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func showChangelog() {
	var versionInfo = `Changelog:
  1.00:
    - First release
  2.00:
    - Modular rewrite code`
	fmt.Println(versionInfo)
}

func main() {
	if cliInstall {
		err := install(cliConfig)
		if err != nil {
			log.Fatal(err)
		}
	}
	if cliUpdate {
		err := update(cliConfig)
		if err != nil {
			log.Fatal(err)
		}
	}
	if cliUninstall {
		err := uninstall()
		if err != nil {
			log.Fatal(err)
		}
	}
	if cliReload {
		err := reload(cliConfig)
		if err != nil {
			log.Println(err)
		}
	}
	if cliTest {
		err := test(cliConfig)
		if err != nil {
			log.Println(err)
		}
	}
}
