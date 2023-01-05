package datastorage_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/AndrewSukhobok95/yagometrics.git/internal/datastorage"
	"github.com/stretchr/testify/assert"
)

const TestBackUpFile = `[
{"delta":10,"id":"C1","type":"counter"},
{"id":"G1","type":"gauge","value":5.67}
]`

func TestUpdateBackUpFile(t *testing.T) {
	// Test Settings
	storeInterval := time.Duration(2 * time.Second)
	sleepInterval := time.Duration(3 * time.Second)
	// Init Storage
	memStorage := datastorage.NewMemStorage()
	memStorage.InsertCounterMetric("C1", 10)
	memStorage.InsertGaugeMetric("G1", 5.67)
	// Creating BackUp File
	file, _ := os.CreateTemp(".", "test-backup-file-*.json")
	defer file.Close()
	defer os.Remove(file.Name())
	// Update BackUp File
	go datastorage.UpdateBackUpFile(memStorage, file, storeInterval)
	time.Sleep(sleepInterval)
	// Asserting
	data, err := os.ReadFile(file.Name())
	if err != nil {
		log.Println("Error while trying to intialise the storage:")
		log.Printf(err.Error() + "\n\n")
		return
	}
	assert.Equal(t, TestBackUpFile, string(data), "Wrong content of the saved json")
}

func TestFillStorageFromBackUpFile(t *testing.T) {
	// Init Storage
	memStorage := datastorage.NewMemStorage()
	// Create BackUp File
	file, _ := os.CreateTemp(".", "test-backup-file-*.json")
	defer file.Close()
	defer os.Remove(file.Name())
	file.WriteString(TestBackUpFile)
	// Fill Storage
	datastorage.FillStorageFromBackUpFile(memStorage, file.Name())
	// Asserting
	delta, err := memStorage.GetCounterMetric("C1")
	assert.NoError(t, err)
	assert.Equal(t, int64(10), delta, "Delta mismatch in counter C1")
	value, err := memStorage.GetGaugeMetric("G1")
	assert.NoError(t, err)
	assert.Equal(t, 5.67, value, "Value mismatch in gauge G1")
}
