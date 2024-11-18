```
 ▄▄   ▄▄ ▄▄▄▄▄▄▄ ▄▄   ▄▄ 
█  █▄█  █       █  █ █  █
█       █  ▄▄▄▄▄█  █▄█  █
█       █ █▄▄▄▄▄█       █
 █     ██▄▄▄▄▄  █   ▄   █
█   ▄   █▄▄▄▄▄█ █  █ █  █
█▄▄█ █▄▄█▄▄▄▄▄▄▄█▄▄█ █▄▄█
```

- Make `xray` easier to install, uninstall and reload
- The application is installed in `/usr/local/bin/xray`, and the application data is stored
  in `/usr/local/etc/xray/` to avoid problems caused by future system changes.

## Usage

```
NAME:
   xsh - xray quick install tool

USAGE:
   xsh [global options] command [command options] 

VERSION:
   v3.01

COMMANDS:
   install, i  Install xray
   uninstall   Remove config,cache and uninstall xray
   update      Update xray
   reload      Reload service
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Install

```sh
# For example, the host is a linux kernel based system with amd64 architecture
# Download the app
curl -Lo /usr/local/bin/xsh https://github.com/gek64/xsh/releases/latest/download/xsh-linux-amd64
# Give the app execute permission
chmod +x /usr/local/bin/xsh
# Show help
/usr/local/bin/xsh -h
```

## Compile

```sh
# Download application source code
git clone https://github.com/gek64/xsh.git
# Compile the source code
cd xsh
export CGO_ENABLED=0
go build -v -trimpath -ldflags "-s -w"
```

## License

- **GPL-3.0 License**
- See `LICENSE` for details

## Credits

- [goland](https://www.jetbrains.com/go/)
- [vscode](https://code.visualstudio.com/)
