package display

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ShowDatabase() {
	path := filepath.Join(".", "databases")

	dir, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return
	}
	defer dir.Close()
	files, err := dir.ReadDir(0)
	if err != nil {
		fmt.Println("Error reading directory contents:", err)
		return
	}

	var folders []string
	for _, file := range files {
		if file.IsDir() {
			folders = append(folders, file.Name())
		}
	}

	for _, folder := range folders {
		fmt.Println(folder)
	}
}

func ShowTable(fileName, db string) {
	fileName = fileName + ".json"
	dbPath := filepath.Join(db, fileName)
	if _, err := os.Stat(dbPath); err != nil {
		fmt.Println("No such table exist")
		return
	}
	file, err := os.Open(dbPath)
	if err != nil {
		fmt.Println("Error opening the file")
		return
	}
	defer file.Close()

	jsonData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	answer := []interface{}{}
	// var data map[string]interface{}
	err = json.Unmarshal(jsonData, &answer)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}
	for i := 0; i < len(answer); i++ {
		ans := answer[i]
		collection, err := json.MarshalIndent(ans, "", "  ")
		if err != nil {
			fmt.Println("Some error occured: ", err)
		}
		fmt.Println(string(collection))
	}
}
