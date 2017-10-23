package main

import (
	"encoding/csv"
	"github.com/jhorowitz/goleveldb/leveldb/table"
	"io"
	"os"
)

func main() {
	db := table.NewWriter(os.Stdout, nil)
	defer db.Close()

	reader := csv.NewReader(os.Stdin)
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		if len(record) != 2 {
			panic("The record should only have 2 columns. A key and a value")
		}

		err = db.Append([]byte(record[0]), []byte(record[1]))
		if err != nil {
			panic(err)
		}
	}
}
