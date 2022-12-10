package launcher

import (
	"api"
	"fileserver"
	"fmt"
	"net/http"
	"runtime"
	"sync"
)

func StartFrontendAndBackendWindows(frontendPath string, launcher ChromeLauncher) {
	fmt.Println("Attempting to start on: " + runtime.GOOS + ", " + runtime.GOARCH)
	var waitgroup *sync.WaitGroup
	waitgroup = &sync.WaitGroup{}
	waitgroup.Add(1)
	// Start Frontend
	launched, waitgroup := launcher.launchChromeForWindows(waitgroup)
	if launched {
		// Start Backend
		err := StartServer(frontendPath)
		if err != nil {
			fmt.Println(err)
		}
		waitgroup.Done()
	} else {
		waitgroup.Done()
	}
}

func StartFrontendAndBackendLinux(frontendPath string, launcher ChromiumLauncher) {
	fmt.Println("Attempting to start on: " + runtime.GOOS + ", " + runtime.GOARCH)
	var waitgroup *sync.WaitGroup
	waitgroup = &sync.WaitGroup{}
	waitgroup.Add(1)
	// Start Frontend
	launched, waitgroup := launcher.launchChromiumForLinux(waitgroup)
	if launched {
		// Start Backend
		err := StartServer(frontendPath)
		if err != nil {
			fmt.Println(err)
		}
		waitgroup.Done()
	} else {
		waitgroup.Done()
	}
}

func StartServer(frontendPath string) error {
	fileserver.FrontendPath = frontendPath

	// HTTP HANDLERS
	http.HandleFunc("/", fileserver.ServeFileServer) // address: /
	http.HandleFunc("/api/", api.ServeAPIUseGZip)    // address: /api/
	return fileserver.GracefulStart()
}
