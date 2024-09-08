package file

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

const (
	FirstName = "first_name"
	LastName  = "last_name"
	Email     = "email"
	Gender    = "gender"
	IpAddress = "ip_address"
	filePath  = "customers.csv"
)

type FileReader struct {
}

func NewFileReader() *FileReader {
	return &FileReader{}
}

type HeaderIndexes map[string]int

// HandleHeader handles the header values and returns the header values map and any errors encountered.
func (f FileReader) HandleHeader(header []string) (HeaderIndexes, []error) {
	var (
		headerValues = make(HeaderIndexes, len(header))
		errors       []error
	)

	for i, value := range header {
		switch value {
		case FirstName:
			headerValues[FirstName] = i
		case LastName:
			headerValues[LastName] = i
		case Email:
			headerValues[Email] = i
		case Gender:
			headerValues[Gender] = i
		case IpAddress:
			headerValues[IpAddress] = i
		default:
			if value == "" {
				errors = append(errors, fmt.Errorf("empty header value with number: %v. Header value is requiered for correct proccessing", i+1))
			} else {
				errors = append(errors, fmt.Errorf("invalid header value: %v", value))
			}
		}
	}
	return headerValues, errors
}

func (f FileReader) ReadFile(recordsCh chan<- []string, errCh chan<- error) {
	file, err := os.Open(filePath)
	if err != nil {
		errCh <- fmt.Errorf("error opening file: %v", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			close(recordsCh)
			break
		} else if err != nil {
			errCh <- fmt.Errorf("error reading file: %v", err)
			break
		}

		recordsCh <- record
	}
}
