package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var persistStorePath = "/tmp/store"

func persist() {
	data, err := json.Marshal(m)
	if len(m) != 0 {
		checkTargetDirectory()

		err := os.WriteFile(persistStorePath+"/data-"+time.Now().String(), data, 0644)
		check(err)
	} else {
		fmt.Println("Store is empty!")
	}
	if err != nil {
		fmt.Println("Error on marshalling...")
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func FetchLastPersistedData() {
	err := checkTargetDirectory()

	if err == nil {
		files, err := ioutil.ReadDir(persistStorePath)
		if err != nil {
			fmt.Println(err)
		}

		if len(files) != 0 {
			// find newest persisted store data
			lastPersistDate := files[0].ModTime()
			lastPersistedStorePath := ""
			for _, file := range files {
				if file.ModTime().After(lastPersistDate) {
					lastPersistedStorePath = file.Name()
				}
			}
			lastPersistedStore, err := os.ReadFile(persistStorePath + "/" + lastPersistedStorePath)

			err = json.Unmarshal(lastPersistedStore, &m)
			if err != nil {
				fmt.Println("an error occurred while unmarshalling...")
			}
		}
	}
}

// check target directory for store persist
// create if not exists
func checkTargetDirectory() error {
	_, err := os.ReadDir(persistStorePath)
	if err != nil {
		fmt.Println("persist store directory not found. creating...")
		err := os.Mkdir(persistStorePath, 0777)
		if err != nil {
			fmt.Println("an error with creating directory")
		}
	}

	err = os.Chmod(persistStorePath, 0777)
	if err != nil {
		fmt.Println("an error on chmod")
	}

	return err
}

// Routine trigger channel for persist data with given interval
func Routine() {
	fmt.Println("Routine started...")
	ticker := time.NewTicker(60 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				persist()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
