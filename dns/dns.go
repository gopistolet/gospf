package dns

import (
	"errors"
	_ "fmt"
	"net"
)

func IsSPF(record string) bool {

	if record[0:5] == "v=spf" {
		return true
	}

	return false

}

func IsSupportedProtocol(record string) bool {

	if record[0:6] == "v=spf1" {
		return true
	}

	return false

}

func GetSPFRecord(name string) (string, error) {

	records, err := net.LookupTXT(name)
	if err != nil {
		return "", err
	}

	for _, record := range records {
		if !IsSPF(record) {
			continue
		}
		if !IsSupportedProtocol(record) {
			return "", errors.New("Unsupported SPF record: " + record)
		}
	}

	return "", nil

}