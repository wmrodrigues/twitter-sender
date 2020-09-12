package loader

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/wmrodrigues/twitter-sender/internal/structs"
	"io"
	"log"
	"os"
)

// LoadFromCsvFile
func LoadFromCsvFile(filePath string) ([]structs.Recipient, error)  {
	var recipients []structs.Recipient

	csvFile, err := os.Open(filePath)

	if err != nil {
		err = fmt.Errorf("error opening csv file, %s", err.Error())
		log.Println(err)
		return nil, err
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			err = fmt.Errorf("some awkward error just happened, %s", err.Error())
			log.Fatal(err)
		}

		name := line[0]
		treatment := line[1]
		username := ""

		if len(line) > 2 {
			username = line[2]
		}

		recipients = append(recipients, structs.Recipient{Name: name,
			Treatment: treatment,
			Username:    username})
	}

	return recipients, nil
}