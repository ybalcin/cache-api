package inmemorystore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	fileNameBase string = "-data-*.json"
	fileRegex    string = `\d+-data-\d+.json`
)

func getCacheFromFile() storage {
	c := readFile()
	store := storage{}

	if c == nil || len(c.Data) <= 0 {
		return store
	}

	for _, item := range c.Data {
		store[item.Key] = item.Value
	}

	return store
}

// StartSaveTask starts a task that saves cache to file in a specified interval time of seconds.
func (c *Client) StartSaveTask() {
	if c.interval <= 0 {
		return
	}

	startSaveTask(c.interval)
}

func startSaveTask(interval int) {
	finish := make(chan struct{}, 1)
	finish <- struct{}{}

	now := time.Now().UTC()

	go func(t int, n time.Time) {
		saveTime := n.Add(time.Second * time.Duration(t))

		for {
			n = time.Now().UTC()
			if saveTime.Before(n) {
				select {
				case <-finish:
					saveCacheToFile(finish)
					n = time.Now().UTC()
					saveTime = n.Add(time.Second * time.Duration(t))
				}
			}
		}
	}(interval, now)
}

func readFile() *cacheItems {
	file := getLatestFile()
	if file == nil {
		return nil
	}
	defer closeFile(file)

	c := cacheItems{}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&c); err != nil {
		panic(err)
	}

	return &c
}

func getLatestFile() *os.File {
	files, err := ioutil.ReadDir(os.TempDir())
	if err != nil {
		panic(err)
	}

	var names []string
	regex := regexp.MustCompile(fileRegex)
	for _, f := range files {
		if !regex.MatchString(f.Name()) {
			continue
		}

		if f.Mode().IsRegular() {
			names = append(names, f.Name())
		}
	}

	if len(names) <= 0 {
		return nil
	}

	var latestStamp int
	var latestFileName string
	for _, n := range names {
		values := strings.Split(n, "-")
		v := values[0]
		if v == "" {
			continue
		}

		t, err := strconv.Atoi(v)
		if err != nil {
			continue
		}

		if t > latestStamp {
			latestStamp = t
			latestFileName = n
		}
	}

	file, err := os.Open(fmt.Sprintf("%s/%s", os.TempDir(), latestFileName))
	if err != nil {
		panic(err)
	}

	return file
}

func saveCacheToFile(finish chan struct{}) {
	cacheItems := mapCacheToStruct()
	if cacheItems == nil {
		finish <- struct{}{}
		return
	}

	filename := fmt.Sprintf("%d%s", time.Now().Unix(), fileNameBase)
	file, err := ioutil.TempFile(os.TempDir(), filename)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFile(file)

	if err := writeToFile(file, cacheItems); err != nil {
		defer removeFile(file)
	}
	defer log.Printf("file saved: %s", file.Name())
	finish <- struct{}{}
}

func writeToFile(file *os.File, c *cacheItems) error {
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}
	if _, err := fmt.Fprintln(file, string(b)); err != nil {
		return err
	}

	return nil
}

func mapCacheToStruct() *cacheItems {
	snapshot := createCacheSnapshot()
	if snapshot == nil {
		return nil
	}

	var c []cacheItem
	for key, value := range snapshot {
		c = append(c, cacheItem{
			Key:   key,
			Value: value,
		})
	}

	return &cacheItems{
		Data: c,
	}
}

func createCacheSnapshot() storage {
	lock.Lock()
	defer lock.Unlock()

	if len(cache) <= 0 {
		return nil
	}

	target := storage{}

	for key, value := range cache {
		if key == "" || value == "" {
			continue
		}

		target[key] = value
	}

	return target
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(err)
	}
}

func removeFile(f *os.File) {
	err := os.Remove(f.Name())
	if err != nil {
		panic(err)
	}
}
