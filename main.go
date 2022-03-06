package main

import (
	"flag"
	"fmt"
	"gek_app"
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
	flag.BoolVar(&cliInstall, "install", false, "install")
	flag.BoolVar(&cliUninstall, "uninstall", false, "uninstall")
	flag.BoolVar(&cliUpdate, "update", false, "update")
	flag.BoolVar(&cliReload, "reload", false, "reload")
	flag.BoolVar(&cliTest, "test", false, "test config")
	flag.StringVar(&cliLocalFile, "l", "", "use local file without Download from network")
	flag.StringVar(&cliConfig, "c", "", "use local config")
	flag.BoolVar(&cliHelp, "help", false, "show help")
	flag.BoolVar(&cliVersion, "v", false, "show version")
	flag.Parse()

	// 重写显示用法函数
	flag.Usage = func() {
		var helpInfo = `Usage:
    proxyctl [Command] [Arguments]

Command:
	-install          : install
	-uninstall        : uninstall
	-update           : update
	-reload           : reload
	-test             : test config
	-h                : show help
	-v                : show version

Arguments:
	-l <bin.zip>      : use local file without Download from network
	-c <config.json>  : use local config

Example:
1) proxyctl -install   -c config.json -l bins.zip  : Install proxy and service
2) proxyctl -update    -c config.json -l bins.zip  : Update proxy and resources
3) proxyctl -uninstall                             : Uninstall proxy and service
4) proxyctl -test      -c config.json              : Test config
5) proxyctl -reload    -c config.json              : Reload service`
		fmt.Println(helpInfo)
	}

	// 如果无 args 或者 指定 h 参数,打印用法后退出
	if len(os.Args) == 1 || cliHelp {
		flag.Usage()
		os.Exit(0)
	}

	// 打印版本信息
	if cliVersion {
		fmt.Println("v2.03")
		os.Exit(0)
	}

	// 检查运行库是否完整
	switch runtime.GOOS {
	case gek_app.SupportedOS[0]:
		err := gek_toolbox.CheckToolbox(linuxToolbox)
		if err != nil {
			log.Panicln(err)
		}
	case gek_app.SupportedOS[1]:
		err := gek_toolbox.CheckToolbox(freebsdToolbox)
		if err != nil {
			log.Panicln(err)
		}
	default:
		log.Panicf("%s is not supported", runtime.GOOS)
	}

	// 初始化
	if cliInstall && cliLocalFile == "" || cliUpdate && cliLocalFile == "" {
		err := initNetwork()
		if err != nil {
			log.Panicln(err)
		}
	} else {
		initLocal()
	}
}

func showChangelog() {
	var versionInfo = `Changelog:
  1.00:
    - First release
  2.00:
    - Modular rewrite code
  2.01:
    - Add local file support
  2.02:
    - Fix bug
  2.03:
    - Fix permission bug`
	fmt.Println(versionInfo)
}

func main() {
	if cliInstall {
		if cliLocalFile != "" {
			err := installFromLocal(cliConfig, cliLocalFile)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := install(cliConfig)
			if err != nil {
				log.Panicln(err)
			}
		}
	}
	if cliUpdate {
		if cliLocalFile != "" {
			err := updateFromLocal(cliConfig, cliLocalFile)
			if err != nil {
				log.Panicln(err)
			}
		} else {
			err := update(cliConfig)
			if err != nil {
				log.Panicln(err)
			}
		}
	}
	if cliUninstall {
		err := uninstall()
		if err != nil {
			log.Panicln(err)
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
