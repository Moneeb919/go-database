package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Moneeb919/go-database/creation"
	"github.com/Moneeb919/go-database/display"
	"github.com/spf13/cobra"
)

const version = "1.0.0"

var globalDatabase string

var rootCmd = &cobra.Command{
	Use:   "go-database",
	Short: "My Database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the go database. Enter commands to run the database. Press Ctrl+C to exit.")
	},
}

var createCommand = &cobra.Command{
	Use:   "create",
	Short: "Creating entities",
}

var showCommand = &cobra.Command{
	Use:   "show",
	Short: "Displaying all the databases",
}

var dbCommand = &cobra.Command{
	Use:   "dbs",
	Short: "Display all dbs",
	Run: func(cmd *cobra.Command, args []string) {
		display.ShowDatabase()
	},
}

var tableName string
var tableCommand = &cobra.Command{
	Use:   "table <arg>",
	Short: "Creating tables",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index := 0
		for i := 0; i < len(args); i++ {
			if args[i] == "table" {
				index = i
			}
		}
		tableName = args[index+1]
		creation.CreatingTable(tableName, globalDatabase)
	},
}

var databaseName string
var databaseCommand = &cobra.Command{
	Use:   "database <arg>",
	Short: "Working with databases",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index := 0
		for i := 0; i < len(args); i++ {
			if args[i] == "database" {
				index = i
			}
		}
		databaseName = args[index+1]
		creation.CreatingData(databaseName)
	},
}

var useCommand = &cobra.Command{
	Use:   "use <db>",
	Short: "Using a particular database",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index := 0
		for i := 0; i < len(args); i++ {
			if args[i] == "use" {
				index = i
			}
		}
		useDatabase(args[index+1])
	},
}

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Adding data to tables",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		data := ""
		for i := 2; i < len(args); i++ {
			data += args[i]
		}
		creation.AddingData(globalDatabase, args[1], data)
	},
}

var displayCommand = &cobra.Command{
	Use:   "display <file>",
	Short: "Displaying table",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		display.ShowTable(args[1], globalDatabase)
	},
}

func useDatabase(dbName string) {
	dir := filepath.Join(".", "databases")
	dbPath := filepath.Join(dir, dbName)
	if globalDatabase == dbPath {
		fmt.Printf("Already using database %s\n", dbName)
		return
	}
	if _, err := os.Stat(dbPath); err != nil {
		fmt.Printf("No such database with name %s\n", dbName)
		return
	}
	globalDatabase = dbPath
	fmt.Println(globalDatabase)
	fmt.Printf("Database shifted to %s\n", dbName)
	return
}

func init() {
	dir := filepath.Dir(".")
	globalDatabase = filepath.Join(dir, "databases")
	createCommand.AddCommand(databaseCommand)
	createCommand.AddCommand(tableCommand)
	showCommand.AddCommand(dbCommand)

	rootCmd.AddCommand(createCommand)
	rootCmd.AddCommand(useCommand)
	rootCmd.AddCommand(showCommand)
	rootCmd.AddCommand(addCommand)
	rootCmd.AddCommand(displayCommand)
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		err := executeCommand(input)
		if err != nil {
			fmt.Println("Error executing commands")
		}
	}

}

func executeCommand(input string) error {

	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil
	}

	cmd, _, err := rootCmd.Find(parts)
	if err != nil {
		return err
	}

	if cmd != nil {
		cmd.Run(cmd, parts)
		return nil
	}

	return fmt.Errorf("Command not found: %s", parts[0])
}
