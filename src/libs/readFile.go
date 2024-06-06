package libs

import (
	"fmt"
	"io"
	"os"
)

func ReadFile(filePath string) (num []int) {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", filePath, err))
	}
	defer fd.Close() // Ensure the file is closed when the function returns

	var line int
	for {
		_, err := fmt.Fscanf(fd, "%d\n", &line)
		if err != nil {
			if err == io.EOF {
				break // Exit the loop on EOF
			}
			panic(fmt.Sprintf("Scan Failed %s: %v", filePath, err))
		}
		num = append(num, line)
	}
	return
}
