package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var _year int
var _month int
var _day int
var wait sync.WaitGroup

func init() {
	fmt.Println("clash start init")
	_year = time.Now().UTC().Year()
	_day = time.Now().UTC().Day()
	_month = int(time.Now().UTC().Month())
	fmt.Printf("current date: %v/%v/%v\n", _year, _month, _day)
}

func renameFile(source string) {
	fileName := strings.Split(source, ".tmp")[0]
	_, err := os.Stat(fileName)
	if os.IsExist(err) {
		os.Remove(fileName)
	}
	os.Rename(source, fileName)
}

func saveToFile(reader io.ReadCloser, fileName string, result func(string, bool)) {
// 	tmp := fileName + ".tmp"
// 	_, err := os.Stat(tmp)
// 	if os.IsExist(err) {
// 		os.Remove(tmp)
// 	}
// 	file, err := os.Create(tmp)
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("create tmp file err:" + tmp)
		return
	}
	defer file.Close()
	resultStatus := true
	buf := make([]byte, 2048)
	for {
		count, wErr := reader.Read(buf)
		if count > 0 {
			rcount, eErr := file.Write(buf[0:count])
			if eErr != nil {
				resultStatus = false
				break
			}
			if rcount != count {
				resultStatus = false
				break
			}
		}
		if wErr != nil {
			if wErr != io.EOF {
				resultStatus = false
			}
			break
		}
	}
	result(fileName, resultStatus)
}

func LoadNodeFreeV2ray() {
	url := "https://nodefree.org/dy/" + strconv.Itoa(_year) + strconv.Itoa(_month) + "/" + strconv.Itoa(_year) + strconv.Itoa(_month) + strconv.Itoa(_day) + ".txt"
	fmt.Println("start load node free v2ray:" + url)
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		fmt.Printf("error load node free v2ray: %v\n", err)
		return
	}
	defer resp.Body.Close()
	result := func(tmp string, status bool) {
		fmt.Printf("load node free v2ray result:%v, file:%v\n", status, tmp)
// 		if status {
// 			renameFile(tmp)
// 		}
		wait.Done()
	}
	saveToFile(resp.Body, "nodeFreeV2ray.txt", result)
}

func LoadNodeFreeClash() {
	url := "https://nodefree.org/dy/" + strconv.Itoa(_year) + strconv.Itoa(_month) + "/" + strconv.Itoa(_year) + strconv.Itoa(_month) + strconv.Itoa(_day) + ".yaml"
	fmt.Println("start load node free clash:" + url)
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		fmt.Printf("error load node free v2ray: %v\n", err)
		return
	}
	defer resp.Body.Close()
	result := func(tmp string, status bool) {
		fmt.Printf("load node free v2ray result:%v, file:%v\n", status, tmp)
// 		if status {
// 			renameFile(tmp)
// 		}
		wait.Done()
	}
	saveToFile(resp.Body, "nodeFreeClash.yaml", result)
}

func LoadClashNodeV2ray() {
	url := "https://clashnode.com/wp-content/uploads/" + strconv.Itoa(_year) + "/" + strconv.Itoa(_month) + "/" + strconv.Itoa(_year) + strconv.Itoa(_month) + strconv.Itoa(_day) + ".txt"
	fmt.Println("start load node free v2ray:" + url)
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		fmt.Printf("error load node free v2ray: %v\n", err)
		return
	}
	defer resp.Body.Close()
	result := func(tmp string, status bool) {
		fmt.Printf("load node free v2ray result:%v, file:%v\n", status, tmp)
// 		if status {
// 			renameFile(tmp)
// 		}
		wait.Done()
	}
	saveToFile(resp.Body, "clashNodeV2ray.txt", result)
}

func LoadClashNodeClash() {
	url := "https://clashnode.com/wp-content/uploads/" + strconv.Itoa(_year) + "/" + strconv.Itoa(_month) + "/" + strconv.Itoa(_year) + strconv.Itoa(_month) + strconv.Itoa(_day) + ".yaml"
	fmt.Println("start load node free clash:" + url)
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		fmt.Printf("error load node free v2ray: %v\n", err)
		return
	}
	defer resp.Body.Close()
	result := func(tmp string, status bool) {
		fmt.Printf("load node free v2ray result:%v, file:%v\n", status, tmp)
// 		if status {
// 			renameFile(tmp)
// 		}
		wait.Done()
	}
	saveToFile(resp.Body, "clashNodeClash.yaml", result)
}

func main() {
	wait.Add(1)
	go LoadNodeFreeClash()
	wait.Add(1)
	go LoadNodeFreeV2ray()
	wait.Add(1)
	go LoadClashNodeClash()
	wait.Add(1)
	go LoadClashNodeV2ray()

	wait.Wait()
}
