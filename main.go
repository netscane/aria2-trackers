package main

import (
	"bufio"
	"flag"
	"net/http"
	"os"
	"path/filepath"
    "io"
    "bytes"
    "fmt"
    "strings"
)

var (
    trackerUrl = "https://cdn.jsdelivr.net/gh/ngosang/trackerslist/trackers_best_ip.txt"
    //trackerUrl = "https://raw.githubusercontent.com/ngosang/trackerslist/master/trackers_best.txt" //need http proxy
)

func getTrackers() ([]string, error) {
    client := &http.Client{}
    request, err := http.NewRequest("GET", trackerUrl, nil)
    if err != nil {
        return nil, err
    }
    //add header
    request.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	rd := bufio.NewReader(resp.Body)
	trackers := []string{}
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			break
		}
		if len(line) > 0 {
			trackers = append(trackers, string(line))
		}
	}
	return trackers, nil
}

func getConfFilePath() (string, error) {
    var (
        home string
        confPath string
        err error
    )
	home, err = os.UserHomeDir()
	if err != nil {
		return "", err
	}
	confPath = filepath.Join(home, "aria2", "aria2.conf")
    return confPath, nil
}

func updateConfFile(confPath string, trackers []string) error {
    var (
        err error
        file *os.File
        lineb []byte
        line string
        reader *bufio.Reader
        bufw *bytes.Buffer
    )

    file, err = os.OpenFile(confPath, os.O_RDWR, 0666)
    if err != nil {
        fmt.Println("Open conf file error!", err)
        return err
    }
    defer file.Close()

    reader = bufio.NewReader(file)
    bufw = new(bytes.Buffer)
    for {
        lineb, _, err = reader.ReadLine()
        if err != nil {
            if err == io.EOF {
                break
            } else {
                fmt.Println("Read file error!", err)
                return err
            }
        }
        line = string(lineb)
        if strings.HasPrefix(line, "bt-tracker") &&
            !strings.HasPrefix(line, "bt-tracker-") {
            line = "bt-tracker=" + strings.Join(trackers, ",") + "\n" 
        }
        if bufw.Len() > 0 {
            bufw.WriteString("\n")
        }
        bufw.WriteString(line)
    }
    file.Seek(0, 0)
    _, err = file.WriteString(bufw.String())
    if err != nil {
        fmt.Println("Write conf file error!", err)
        return err
    }
    return nil
}

func main() {
	flag.Parse()
	if len(flag.Arg(0)) > 0 {
		trackerUrl = flag.Arg(0)
	}
    var (
        trackers []string
        err error
        confPath string
    )
	trackers, err = getTrackers()
	if err != nil {
		panic(err.Error())
	}
    confPath, err = getConfFilePath()
	if err != nil {
		panic(err.Error())
	}
	err = updateConfFile(confPath, trackers)
	if err != nil {
		panic(err.Error())
	}
    fmt.Println("Update tracker success!")
}
