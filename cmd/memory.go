package cmd

import (
	"fmt"
	"strings"

	"github.com/m99Tanishq/CLI/internal/api"
	"github.com/m99Tanishq/CLI/internal/config"
	"github.com/m99Tanishq/CLI/internal/memory"
	"github.com/spf13/cobra"
)

var memoryCmd = &cobra.Command{
	Use:   "memory",
	Short: "Manage codebase memory and indexing",
	Long:  `Index, store, and query the entire codebase for better AI context and analysis.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Memory management commands:")
		fmt.Println("  index <path>     - Index a codebase for memory")
		fmt.Println("  query <query>    - Query the indexed codebase")
		fmt.Println("  list             - List indexed codebases")
		fmt.Println("  clear            - Clear all indexed data")
		fmt.Println("  analyze <path>   - Analyze codebase with memory context")
	},
}

var indexCmd = &cobra.Command{
	Use:   "index [path]",
	Short: "Index a codebase for memory",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}

		// Load configuration
		cfg, err := config.Load()
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		if cfg.APIKey == "" {
			fmt.Println("Error: API key not configured")
			return
		}

		// Create memory manager
		memManager := memory.NewManager()

		fmt.Printf("Indexing codebase at: %s\n", path)
		fmt.Println("This may take a while for large codebases...")

		// Index the codebase
		index, err := memManager.IndexCodebase(path)
		if err != nil {
			fmt.Printf("Error indexing codebase: %v\n", err)
			return
		}

		fmt.Printf("✅ Successfully indexed %d files\n", len(index.Files))
		fmt.Printf("📊 Total lines of code: %d\n", index.TotalLines)
		fmt.Printf("📁 Directories scanned: %d\n", index.Directories)
		fmt.Printf("💾 Memory size: %s\n", formatBytes(index.MemorySize))
	},
}

var queryCmd = &cobra.Command{
	Use:   "query [query]",
	Short: "Query the indexed codebase",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]

		// Load configuration
		cfg, err := config.Load()
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		if cfg.APIKey == "" {
			fmt.Println("Error: API key not configured")
			return
		}

		// Create memory manager
		memManager := memory.NewManager()

		// Load indexed data
		index, err := memManager.LoadIndex()
		if err != nil {
			fmt.Printf("Error loading index: %v\n", err)
			fmt.Println("Please run 'CLI memory index' first")
			return
		}

		// Create API client
		client := api.NewClient(cfg.APIKey)
		client.BaseURL = cfg.BaseURL

		// Prepare query with context
		prompt := fmt.Sprintf(`You have access to an indexed codebase. Please answer the following query based on the codebase information:

Query: %s

Codebase Information:
- Total files: %d
- Total lines: %d
- Directories: %d
- Languages: %s

File Structure:
%s

Key Files and Their Purposes:
%s

Please provide a comprehensive answer based on the codebase structure and content. If the query is about specific functionality, explain how it's implemented in the codebase.`,
			query,
			len(index.Files),
			index.TotalLines,
			index.Directories,
			strings.Join(index.Languages, ", "),
			formatFileStructure(index.Files),
			formatKeyFiles(index.Files))

		req := api.ChatRequest{
			Model: cfg.Model,
			Messages: []api.Message{
				{
					Role:    "user",
					Content: prompt,
				},
			},
		}

		fmt.Printf("Querying codebase: %s\n", query)
		resp, err := client.SendChat(req)
		if err != nil {
			fmt.Printf("Error querying codebase: %v\n", err)
			return
		}

		if len(resp.Choices) > 0 {
			fmt.Println("\n=== Codebase Query Result ===")
			fmt.Println(resp.Choices[0].Message.Content)
		}
	},
}

var memoryListCmd = &cobra.Command{
	Use:   "list",
	Short: "List indexed codebases",
	Run: func(cmd *cobra.Command, args []string) {
		memManager := memory.NewManager()

		index, err := memManager.LoadIndex()
		if err != nil {
			fmt.Println("No indexed codebases found.")
			fmt.Println("Use 'CLI memory index' to index a codebase.")
			return
		}

		fmt.Println("=== Indexed Codebases ===")
		fmt.Printf("📁 Root Path: %s\n", index.RootPath)
		fmt.Printf("📊 Files Indexed: %d\n", len(index.Files))
		fmt.Printf("📈 Total Lines: %d\n", index.TotalLines)
		fmt.Printf("🗂️  Directories: %d\n", index.Directories)
		fmt.Printf("💾 Memory Size: %s\n", formatBytes(index.MemorySize))
		fmt.Printf("🕒 Last Updated: %s\n", index.LastUpdated.Format("2006-01-02 15:04:05"))

		fmt.Println("\n📋 Languages Found:")
		for _, lang := range index.Languages {
			fmt.Printf("  • %s\n", lang)
		}

		fmt.Println("\n📁 Top-Level Structure:")
		for _, file := range index.Files {
			if !strings.Contains(file.Path, "/") {
				if file.IsDir {
					fmt.Printf("  📁 %s/\n", file.Name)
				} else {
					fmt.Printf("  📄 %s\n", file.Name)
				}
			}
		}
	},
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all indexed data",
	Run: func(cmd *cobra.Command, args []string) {
		memManager := memory.NewManager()

		err := memManager.ClearIndex()
		if err != nil {
			fmt.Printf("Error clearing index: %v\n", err)
			return
		}

		fmt.Println("✅ All indexed data cleared successfully")
	},
}

var memoryAnalyzeCmd = &cobra.Command{
	Use:   "analyze [path]",
	Short: "Analyze codebase with memory context",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}

		// Load configuration
		cfg, err := config.Load()
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		if cfg.APIKey == "" {
			fmt.Println("Error: API key not configured")
			return
		}

		// Create memory manager
		memManager := memory.NewManager()

		// Load or create index
		index, err := memManager.LoadIndex()
		if err != nil {
			fmt.Println("No existing index found. Creating new index...")
			index, err = memManager.IndexCodebase(path)
			if err != nil {
				fmt.Printf("Error indexing codebase: %v\n", err)
				return
			}
		}

		// Create API client
		client := api.NewClient(cfg.APIKey)
		client.BaseURL = cfg.BaseURL

		// Prepare analysis prompt
		prompt := fmt.Sprintf(`Please provide a comprehensive analysis of this codebase based on the indexed information:

Codebase Overview:
- Root Path: %s
- Total Files: %d
- Total Lines: %d
- Directories: %d
- Languages: %s

File Structure:
%s

Key Files Analysis:
%s

Please provide:
1. **Architecture Overview**: Describe the overall structure and design patterns
2. **Technology Stack**: Identify the technologies and frameworks used
3. **Code Quality Assessment**: Evaluate code organization, naming conventions, and structure
4. **Potential Issues**: Identify any architectural or organizational problems
5. **Improvement Suggestions**: Provide specific recommendations for better organization
6. **Security Considerations**: Highlight any potential security concerns
7. **Performance Insights**: Suggest performance optimizations if applicable
8. **Maintainability Score**: Rate the codebase maintainability (1-10)

Format your response with clear sections and actionable insights.`,
			index.RootPath,
			len(index.Files),
			index.TotalLines,
			index.Directories,
			strings.Join(index.Languages, ", "),
			formatFileStructure(index.Files),
			formatDetailedFiles(index.Files))

		req := api.ChatRequest{
			Model: cfg.Model,
			Messages: []api.Message{
				{
					Role:    "user",
					Content: prompt,
				},
			},
		}

		fmt.Printf("Analyzing codebase at: %s\n", path)
		resp, err := client.SendChat(req)
		if err != nil {
			fmt.Printf("Error analyzing codebase: %v\n", err)
			return
		}

		if len(resp.Choices) > 0 {
			fmt.Println("\n=== Codebase Analysis ===")
			fmt.Println(resp.Choices[0].Message.Content)
		}
	},
}

func init() {
	memoryCmd.AddCommand(indexCmd)
	memoryCmd.AddCommand(queryCmd)
	memoryCmd.AddCommand(memoryListCmd)
	memoryCmd.AddCommand(clearCmd)
	memoryCmd.AddCommand(memoryAnalyzeCmd)
}

// Helper functions for formatting
func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func formatFileStructure(files []memory.FileInfo) string {
	var result strings.Builder
	for _, file := range files {
		indent := strings.Repeat("  ", strings.Count(file.Path, "/"))
		if file.IsDir {
			result.WriteString(fmt.Sprintf("%s📁 %s/\n", indent, file.Name))
		} else {
			result.WriteString(fmt.Sprintf("%s📄 %s\n", indent, file.Name))
		}
	}
	return result.String()
}

func formatKeyFiles(files []memory.FileInfo) string {
	var result strings.Builder
	for _, file := range files {
		if !file.IsDir && isKeyFile(file.Name) {
			result.WriteString(fmt.Sprintf("- %s (%s, %d lines): %s\n",
				file.Path, file.Language, file.Lines, file.Purpose))
		}
	}
	return result.String()
}

func formatDetailedFiles(files []memory.FileInfo) string {
	var result strings.Builder
	for _, file := range files {
		if !file.IsDir {
			result.WriteString(fmt.Sprintf("- %s (%s, %d lines)\n",
				file.Path, file.Language, file.Lines))
		}
	}
	return result.String()
}

func isKeyFile(filename string) bool {
	keyFiles := []string{
		"main.go", "go.mod", "package.json", "requirements.txt", "Dockerfile",
		"README.md", "Makefile", "CMakeLists.txt", "pom.xml", "build.gradle",
		"setup.py", "Cargo.toml", "composer.json", "Gemfile", "pubspec.yaml",
	}
	for _, key := range keyFiles {
		if filename == key {
			return true
		}
	}
	return false
}
