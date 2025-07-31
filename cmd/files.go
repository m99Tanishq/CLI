package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/m99Tanishq/CLI/pkg/utils"
	"github.com/spf13/cobra"
)

var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "Manipulate files and folders",
	Long:  `Create, read, write, and manipulate files and folders with AI assistance.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui := utils.NewModernUI()
		ui.PrintHeader("File Management")

		ui.PrintSection("Available Commands")
		ui.PrintList([]string{
			"read <file>     - Read and display file contents",
			"write <file>    - Write content to a file",
			"create <file>   - Create a new file",
			"list <dir>      - List files in directory",
			"search <dir>    - Search for files by pattern",
			"analyze <file>  - Analyze file with AI",
			"fix <file>      - Fix code issues with AI",
		}, "📁")

		ui.PrintInfo("Use --help for detailed information about each command")
	},
}

var readCmd = &cobra.Command{
	Use:   "read [file]",
	Short: "Read and display file contents",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ui := utils.NewModernUI()
		filePath := args[0]

		ui.PrintHeader("File Reader")
		ui.PrintInfo(fmt.Sprintf("Reading: %s", filePath))

		content, err := os.ReadFile(filePath)
		if err != nil {
			ui.PrintError(fmt.Sprintf("Error reading file: %v", err))
			return
		}

		ui.PrintSection("File Contents")
		ui.PrintCard(filePath, string(content))
	},
}

var writeCmd = &cobra.Command{
	Use:   "write [file] [content]",
	Short: "Write content to a file",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ui := utils.NewModernUI()
		filePath := args[0]
		content := strings.Join(args[1:], " ")

		ui.PrintHeader("File Writer")
		ui.PrintInfo(fmt.Sprintf("Writing to: %s", filePath))

		err := os.WriteFile(filePath, []byte(content), 0600)
		if err != nil {
			ui.PrintError(fmt.Sprintf("Error writing file: %v", err))
			return
		}

		ui.PrintSuccess(fmt.Sprintf("Successfully wrote to %s", filePath))
	},
}

var createCmd = &cobra.Command{
	Use:   "create [file]",
	Short: "Create a new file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ui := utils.NewModernUI()
		filePath := args[0]

		ui.PrintHeader("File Creator")
		ui.PrintInfo(fmt.Sprintf("Creating: %s", filePath))

		// Create directory if it doesn't exist
		dir := filepath.Dir(filePath)
		if dir != "." {
			if err := os.MkdirAll(dir, 0755); err != nil {
				ui.PrintError(fmt.Sprintf("Error creating directory: %v", err))
				return
			}
		}

		file, err := os.Create(filePath)
		if err != nil {
			ui.PrintError(fmt.Sprintf("Error creating file: %v", err))
			return
		}
		defer file.Close()

		ui.PrintSuccess(fmt.Sprintf("Successfully created %s", filePath))
	},
}

var listCmd = &cobra.Command{
	Use:   "list [directory]",
	Short: "List files in directory",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ui := utils.NewModernUI()
		dir := "."
		if len(args) > 0 {
			dir = args[0]
		}

		ui.PrintHeader("Directory Listing")
		ui.PrintInfo(fmt.Sprintf("Listing: %s", dir))

		entries, err := os.ReadDir(dir)
		if err != nil {
			ui.PrintError(fmt.Sprintf("Error reading directory: %v", err))
			return
		}

		ui.PrintSection("Contents")

		// Prepare table data
		headers := []string{"Type", "Name", "Size"}
		var rows [][]string

		for _, entry := range entries {
			entryType := "📄"
			if entry.IsDir() {
				entryType = "📁"
			}

			info, err := entry.Info()
			size := "N/A"
			if err == nil {
				if info.Size() < 1024 {
					size = fmt.Sprintf("%d B", info.Size())
				} else if info.Size() < 1024*1024 {
					size = fmt.Sprintf("%.1f KB", float64(info.Size())/1024)
				} else {
					size = fmt.Sprintf("%.1f MB", float64(info.Size())/(1024*1024))
				}
			}

			rows = append(rows, []string{entryType, entry.Name(), size})
		}

		ui.PrintTable(headers, rows)
	},
}

var searchCmd = &cobra.Command{
	Use:   "search [directory] [pattern]",
	Short: "Search for files by pattern",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ui := utils.NewModernUI()
		dir := args[0]
		pattern := args[1]

		ui.PrintHeader("File Search")
		ui.PrintInfo(fmt.Sprintf("Searching in: %s", dir))
		ui.PrintInfo(fmt.Sprintf("Pattern: %s", pattern))

		matches, err := filepath.Glob(filepath.Join(dir, pattern))
		if err != nil {
			ui.PrintError(fmt.Sprintf("Error searching: %v", err))
			return
		}

		if len(matches) == 0 {
			ui.PrintWarning(fmt.Sprintf("No files found matching pattern: %s", pattern))
			return
		}

		ui.PrintSection("Search Results")
		ui.PrintList(matches, "📄")
	},
}

func init() {
	filesCmd.AddCommand(readCmd)
	filesCmd.AddCommand(writeCmd)
	filesCmd.AddCommand(createCmd)
	filesCmd.AddCommand(listCmd)
	filesCmd.AddCommand(searchCmd)
}
