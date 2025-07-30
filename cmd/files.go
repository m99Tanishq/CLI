package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "Manipulate files and folders",
	Long:  `Create, read, write, and manipulate files and folders with AI assistance.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("File manipulation commands:")
		fmt.Println("  read <file>     - Read and display file contents")
		fmt.Println("  write <file>    - Write content to a file")
		fmt.Println("  create <file>   - Create a new file")
		fmt.Println("  list <dir>      - List files in directory")
		fmt.Println("  search <dir>    - Search for files by pattern")
		fmt.Println("  analyze <file>  - Analyze file with AI")
		fmt.Println("  fix <file>      - Fix code issues with AI")
	},
}

var readCmd = &cobra.Command{
	Use:   "read [file]",
	Short: "Read and display file contents",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}
		fmt.Printf("=== Contents of %s ===\n", filePath)
		fmt.Println(string(content))
	},
}

var writeCmd = &cobra.Command{
	Use:   "write [file] [content]",
	Short: "Write content to a file",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		content := strings.Join(args[1:], " ")

		err := os.WriteFile(filePath, []byte(content), 0600)
		if err != nil {
			fmt.Printf("Error writing file: %v\n", err)
			return
		}
		fmt.Printf("Successfully wrote to %s\n", filePath)
	},
}

var createCmd = &cobra.Command{
	Use:   "create [file]",
	Short: "Create a new file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]

		// Create directory if it doesn't exist
		dir := filepath.Dir(filePath)
		if dir != "." {
			if err := os.MkdirAll(dir, 0755); err != nil {
				fmt.Printf("Error creating directory: %v\n", err)
				return
			}
		}

		file, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
			return
		}
		defer file.Close()

		fmt.Printf("Successfully created %s\n", filePath)
	},
}

var listCmd = &cobra.Command{
	Use:   "list [directory]",
	Short: "List files in directory",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir := "."
		if len(args) > 0 {
			dir = args[0]
		}

		entries, err := os.ReadDir(dir)
		if err != nil {
			fmt.Printf("Error reading directory: %v\n", err)
			return
		}

		fmt.Printf("=== Contents of %s ===\n", dir)
		for _, entry := range entries {
			if entry.IsDir() {
				fmt.Printf("📁 %s/\n", entry.Name())
			} else {
				fmt.Printf("📄 %s\n", entry.Name())
			}
		}
	},
}

var searchCmd = &cobra.Command{
	Use:   "search [directory] [pattern]",
	Short: "Search for files by pattern",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]
		pattern := args[1]

		matches, err := filepath.Glob(filepath.Join(dir, pattern))
		if err != nil {
			fmt.Printf("Error searching: %v\n", err)
			return
		}

		if len(matches) == 0 {
			fmt.Printf("No files found matching pattern: %s\n", pattern)
			return
		}

		fmt.Printf("=== Files matching '%s' in %s ===\n", pattern, dir)
		for _, match := range matches {
			fmt.Println(match)
		}
	},
}

func init() {
	filesCmd.AddCommand(readCmd)
	filesCmd.AddCommand(writeCmd)
	filesCmd.AddCommand(createCmd)
	filesCmd.AddCommand(listCmd)
	filesCmd.AddCommand(searchCmd)
}
