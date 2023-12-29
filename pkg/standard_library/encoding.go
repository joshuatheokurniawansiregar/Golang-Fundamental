package standard_library

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func RunningCSVReading() {
	var names string = "Joshua,Theo,Kurniawan,Siregar\n" +
		"Anastasya,Gabriella,Siregar\n" + "Jemima,Joy,Krisanti\n"
	reader := csv.NewReader(strings.NewReader(names))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}
