package creation

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	// "gopkg.in/mgo.v2/bson"
)

func CreatingData(arg string) {
	dir := filepath.Dir(".")
	path := filepath.Join(dir, "databases")
	fmt.Println(dir)
	creationPath := filepath.Join(path, arg)
	if _, err := os.Stat(creationPath); err == nil {
		fmt.Println("The database already exists")
		return
	}
	err := os.Mkdir(creationPath, 0755)
	if err != nil {
		fmt.Println("Could not create the database: ", err)
	} else {
		fmt.Printf("\nDatabase %s created successfully\n", arg)
	}
}

func CreatingTable(arg string, db string) {
	if db == "" {
		fmt.Println("No database specified!")
		return
	}
	if _, err := os.Stat(db); err != nil {
		fmt.Println(db)
		fmt.Println("The database does not exist")
		return
	}
	fileName := arg + ".json"
	fileLoc := db + string(os.PathSeparator) + fileName

	file, err := os.Create(fileLoc)
	data := []interface{}{}
	initialData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	os.WriteFile(fileLoc, initialData, 0644)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}
	defer file.Close()
	fmt.Println("File created successfully!")
}

func StringToMap(data string) map[string]interface{} {
	var jsonData map[string]interface{}
	err := json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return jsonData
	}

	fmt.Println(jsonData)
	return jsonData
}

func AddingData(database, fileName, data string) {
	fileName = fileName + ".json"
	dbPath := filepath.Join(database, fileName)
	if _, err := os.Stat(dbPath); err != nil {
		fmt.Println("Wrong directory!")
		return
	}
	jsonData := StringToMap(data)
	if jsonData == nil {
		return
	}
	file, err := os.ReadFile(dbPath)
	if err != nil {
		fmt.Println("Error reading the file")
		return
	}
	inputData := []interface{}{}
	err = json.Unmarshal(file, &inputData)
	fmt.Println(inputData)
	inputData = append(inputData, jsonData)
	fmt.Println(inputData)

	updatedFile, err := json.MarshalIndent(inputData, "", "  ")
	if err != nil {
		fmt.Println("Error converting to json:", err)
		return
	}

	err = os.WriteFile(dbPath, updatedFile, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Written to file successfully!")
}
