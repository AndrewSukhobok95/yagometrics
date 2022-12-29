package datastorage

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/serialization"
)

func BackUpToFile(storage Storage, filePath string, storeInterval time.Duration, restore bool, wg *sync.WaitGroup) {
	log.Printf("Path to back up file:" + filePath)

	if restore {
		log.Println("Filling storage from back up file")
		FillStorageFromBackUpFile(storage, filePath)
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Println("Error during creating the back up file:")
		log.Printf(err.Error() + "\n\n")
		return
	}

	log.Printf("Updating file every " + storeInterval.String() + "\n")
	wg.Add(1)
	go func() {
		UpdateBackUpFile(storage, file, storeInterval)
		file.Close()
	}()
}

func UpdateBackUpFile(storage Storage, file *os.File, storeInterval time.Duration) {
	ticker := time.NewTicker(storeInterval)
	for {
		<-ticker.C
		err := file.Truncate(0)
		if err != nil {
			log.Println("Back up file was not cleared:")
			log.Printf(err.Error() + "\n\n")
		}
		_, err = file.Seek(0, 0)
		if err != nil {
			log.Println("Couldn't find the begging of the file:")
			log.Printf(err.Error() + "\n\n")
		}
		file.WriteString(storage.ExportToJSONString())
		log.Println("Back up file successfully updated")
	}
}

func FillStorageFromBackUpFile(storage Storage, filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("Error while trying to intialise the storage:")
		log.Printf(err.Error() + "\n\n")
		return
	}
	var metrics []serialization.Metrics
	err = json.Unmarshal(data, &metrics)
	if err != nil {
		log.Println("Error while trying to unmarshal the read data:")
		log.Printf(err.Error() + "\n\n")
		return
	}
	for _, m := range metrics {
		switch m.MType {
		case "counter":
			storage.InsertCounterMetric(m.ID, *m.Delta)
		case "gauge":
			storage.InsertGaugeMetric(m.ID, *m.Value)
		}
	}
}
