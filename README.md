# go-local-web-gui (Local Go Chrome Framework)

go-local-web (GOLW) is a simple framework made for developing localhosted software that can reuse chrome/chromium or embed chromium (in future releases). Both available in deployment for the applications.

This framework uses Chrome (Windows) or Chromium (Linux) as frontend by opening them with cmd/terminal and hosting a localhost webserver, while opening chrome/chromium with --app and --user-directory arguments. The frontend can be changed by the user in runtime, while the backend needs to be compiled/build. The API can be decoupled in future versions, so every part of the application is changeable - Sustainable development. Frontends is easy to change. Alternatives to this is embedding a chromium or webview framework into the project, which will require more space. I chose to depend on Chrome/Chromium, as they are market leaders and html/css/javascript technology frontrunners.

Feel free to use this piece of software, I will be happy to assist you

I am currently working on this project, it will be updated and maintained. 

I consider it production ready.

## Requirements to developers
Go 1.19+
Chrome (Windows) or Chromium (Linux)

## Requirements for users
Chrome (Windows) or Chromium (Linux)

## How to use (download example project)
The best way to start using the project is to download the example project at:
https://github.com/NineNineFive/go-local-web-example

This example project uses this package and combines it with a local api
Then the Go api is being developed and customized by you together with the frontend (JavaScript, HTML, CSS)

## How to use (with go get)
first run the following in CMD (with go installed)
<code>go get github.com/NineNineFive/go-local-web-gui/</code>
Example: how to add framework to main.go
<pre>
package main

import (
	"github.com/NineNineFive/go-local-web-gui/fileserver"
	"github.com/NineNineFive/go-local-web-gui/launcher"
	"net/http"
	"os"
	"runtime"
)

// For windows we need a organisation name and project name
var projectName = "NewProjectName"
var organisationName = "NewOrganisationName"

var frontendPath = "./frontend"

// remember to change the ports to something unique
var chromeLauncher = launcher.ChromeLauncher{
	Location:                "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe",
	LocationCMD:             "C:\\\"Program Files\"\\Google\\Chrome\\Application\\chrome.exe",
	FrontendInstallLocation: os.Getenv("localappdata") + "\\Google\\Chrome\\InstalledApps\\" + organisationName + "\\" + projectName,
	Domain:                  "localhost",
	PortMin:                 11430, 
	PreferredPort:           11451,
	PortMax:                 11500,
}

// remember to change the ports to something unique
var chromiumLauncher = launcher.ChromiumLauncher{
	Location:      "/var/lib/snapd/desktop/applications/chromium_chromium.desktop", // TODO: check if better location or can be customised
	Domain:        "localhost",
	PortMin:       11430,
	PreferredPort: 11451,
	PortMax:       11500,
}

func main() {
	launchApp()
}

func initHTTPHandlers() {
    // static fileserver
	http.HandleFunc("/", fileserver.ServeFileServer)

    // api (local api can be added)
	//http.HandleFunc("/api/", api.ServeAPIUseGZip)
}

func launchApp() {
	switch runtime.GOOS {
	case "windows":
		initHTTPHandlers()
		launcher.StartFrontendAndBackendWindows(frontendPath, chromeLauncher)
		return
	case "darwin": // "mac"
		panic("Darwin Not Supported Yet")
		return
	case "linux": // "linux"
		initHTTPHandlers()
		launcher.StartFrontendAndBackendLinux(frontendPath, chromiumLauncher)
		return
	default: // "freebsd", "openbsd", "netbsd"
		initHTTPHandlers()
		launcher.StartFrontendAndBackendLinux(frontendPath, chromiumLauncher)
		return
	}
}
</pre>

## How to run
<code>go run main.go</code>

## How to build
<code>go build -ldflags -H=windowsgui -o NewProjectName.exe</code>

## How to apply manifest and logo to executible
Use something like goversioninfo: https://github.com/josephspurrier/goversioninfo 

## How to make setup file and update functionality
Coming later