package standard_library

import (
	"os"
)

func CreateNewFile(filename string, message string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_RDONLY|os.O_APPEND, 0777)

	if err != nil {
		file, err = os.Create(filename)
		file.WriteString(message)
		if err == nil {
			return
		}
	}
	defer file.Close()
	file.WriteString(message)
	// fmt.Fprintf(file, "%s", message)
}
