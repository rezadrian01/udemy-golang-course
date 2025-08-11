package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadBalanceFile(filename string) (float64, error) {
	valueByte, err := os.ReadFile(filename)

	if err != nil {
		return 1000, errors.New("failed to read file")
	}

	valueText := string(valueByte)
	valueFloat, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse value")
	}

	return valueFloat, nil

}

func WriteBalanceFile(value float64, filename string) error {
	valueText := fmt.Sprint(value)
	return os.WriteFile(filename, []byte(valueText), 0644)
}