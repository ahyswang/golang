package samples

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"
)

func TestCSV(t *testing.T) {
	file, err := os.OpenFile("test.csv", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("open file error", err)
		return
	}
	w := csv.NewWriter(file)
	w.Write([]string{"123", "234", "abc"})
	w.Write([]string{"1", "2", "3"})
	w.Flush()
	file.Close()

	rfile, err := os.Open("test.csv")
	if err != nil {
		fmt.Println("open file error", err)
		return
	}
	r := csv.NewReader(rfile)
	strs, err := r.Read()
	fmt.Println(strs)
}
