module proxyctl

require (
	gek_app v0.0.0
	gek_exec v0.0.0
	gek_toolbox v0.0.0
)

replace (
	gek_app => ../gek/gek_app
	gek_downloader => ../gek/gek_downloader
	gek_exec => ../gek/gek_exec
	gek_file => ../gek/gek_file
	gek_github => ../gek/gek_github
	gek_json => ../gek/gek_json
	gek_math => ../gek/gek_math
	gek_service => ../gek/gek_service
	gek_service_freebsd => ../gek/gek_service_freebsd
	gek_toolbox => ../gek/gek_toolbox
)
