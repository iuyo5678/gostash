// +build ignore

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"runtime"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
    versionRe = regexp.MustCompile(`-[0-9]{1,3}-g[0-9a-f]{5,10}`)
    goarch    string
    goos      string
    gocc      string
    gocxx     string
    cgo       string
    pkgArch   string
    version   string = "v1"
    // deb & rpm does not support semver so have to handle their version a little differently
    linuxPackageVersion   string = "v1"
    linuxPackageIteration string = ""
    race                  bool
    phjsToRelease         string
    workingDir            string
)

const minGoVersion = 1.7

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
	
	ensureGoPath()
	
	log.Printf("Version: %s, Linux Version: %s, Package Iteration: %s\n", version, linuxPackageVersion, linuxPackageIteration)
	
	flag.StringVar(&goarch, "goarch", runtime.GOARCH, "GOARCH")
	flag.StringVar(&goos, "goos", runtime.GOOS, "GOOS")
	flag.StringVar(&gocc, "cc", "", "CC")
	flag.StringVar(&gocxx, "cxx", "", "CXX")
	flag.StringVar(&cgo, "cgo-enabled", "", "CGO_ENABLED")
	flag.StringVar(&pkgArch, "pkg-arch", "", "PKG ARCH")
	flag.StringVar(&phjsToRelease, "phjs", "", "PhantomJS binary")
	flag.BoolVar(&race, "race", race, "Use race detector")
	flag.Parse()
	
	if flag.NArg() == 0 {
		log.Println("Usage: go run build.go build")
		return
	}
	
	workingDir, _ = os.Getwd()
	
	for _, cmd := range flag.Args() {
		switch cmd {
		case "setup":
			setup()
			
		default:
			log.Fatalf("Unknown command %q", cmd)
		}
	}
}

func ensureGoPath() {
	if os.Getenv("GOPATH") == "" {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		gopath := filepath.Clean(filepath.Join(cwd, "../../../../"))
		log.Println("GOPATH is", gopath)
		os.Setenv("GOPATH", gopath)
	}
}

func ChangeWorkingDir(dir string) {
	os.Chdir(dir)
}


func setup() {
	runPrint("go", "get", "-v", "github.com/Shopify/sarama")
	runPrint("go", "get", "-v", "gopkg.in/olivere/elastic.v3")
	runPrint("go", "get", "-v", "github.com/spf13/viper")
}


func clean() {
	rmr("dist")
	rmr("tmp")
}


func readVersionFromPackageJson() {
    reader, err := os.Open("package.json")
    if err != nil {
        log.Fatal("Failed to open package.json")
        return
    }
    defer reader.Close()
	
    jsonObj := map[string]interface{}{}
    jsonParser := json.NewDecoder(reader)
	
    if err := jsonParser.Decode(&jsonObj); err != nil {
        log.Fatal("Failed to decode package.json")
    }
	
    version = jsonObj["version"].(string)
    linuxPackageVersion = version
    linuxPackageIteration = ""
	
    // handle pre version stuff (deb / rpm does not support semver)
    parts := strings.Split(version, "-")
	
    if len(parts) > 1 {
        linuxPackageVersion = parts[0]
        linuxPackageIteration = parts[1]
    }
	
    // add timestamp to iteration
    linuxPackageIteration = fmt.Sprintf("%d%s", time.Now().Unix(), linuxPackageIteration)
}


func runPrint(cmd string, args ...string) {
    log.Println(cmd, strings.Join(args, " "))
    ecmd := exec.Command(cmd, args...)
    ecmd.Stdout = os.Stdout
    ecmd.Stderr = os.Stderr
    err := ecmd.Run()
    if err != nil {
        log.Fatal(err)
    }
}

func rmr(paths ...string) {
    for _, path := range paths {
        log.Println("rm -r", path)
        os.RemoveAll(path)
    }
}
