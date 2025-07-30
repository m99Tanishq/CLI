package memory

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FileInfo represents information about a file in the codebase
type FileInfo struct {
	Path     string `json:"path"`
	Name     string `json:"name"`
	IsDir    bool   `json:"is_dir"`
	Language string `json:"language"`
	Lines    int    `json:"lines"`
	Size     int64  `json:"size"`
	Purpose  string `json:"purpose"`
}

// CodebaseIndex represents the indexed information about a codebase
type CodebaseIndex struct {
	RootPath    string     `json:"root_path"`
	Files       []FileInfo `json:"files"`
	TotalLines  int        `json:"total_lines"`
	Directories int        `json:"directories"`
	Languages   []string   `json:"languages"`
	MemorySize  int64      `json:"memory_size"`
	LastUpdated time.Time  `json:"last_updated"`
}

// Manager handles codebase indexing and memory operations
type Manager struct {
	indexPath string
}

// NewManager creates a new memory manager
func NewManager() *Manager {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}

	return &Manager{
		indexPath: filepath.Join(homeDir, ".CLI", "memory"),
	}
}

// IndexCodebase indexes a codebase and stores the information
func (m *Manager) IndexCodebase(rootPath string) (*CodebaseIndex, error) {
	// Create memory directory if it doesn't exist
	if err := os.MkdirAll(m.indexPath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create memory directory: %w", err)
	}

	index := &CodebaseIndex{
		RootPath:    rootPath,
		Files:       []FileInfo{},
		Languages:   []string{},
		LastUpdated: time.Now(),
	}

	// Walk through the codebase
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip hidden files and directories (but allow some important ones)
		if strings.HasPrefix(info.Name(), ".") {
			// Allow some important hidden files
			allowedHidden := []string{".gitignore", ".env", ".dockerignore"}
			isAllowed := false
			for _, allowed := range allowedHidden {
				if info.Name() == allowed {
					isAllowed = true
					break
				}
			}

			if !isAllowed {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}

		// Skip common directories to ignore
		skipDirs := []string{".git", "node_modules", "vendor", "bin", "obj", "build", "dist"}
		if info.IsDir() {
			for _, skipDir := range skipDirs {
				if info.Name() == skipDir {
					return filepath.SkipDir
				}
			}
		}

		// Calculate relative path
		relPath, err := filepath.Rel(rootPath, path)
		if err != nil {
			relPath = path
		}

		fileInfo := FileInfo{
			Path:  relPath,
			Name:  info.Name(),
			IsDir: info.IsDir(),
			Size:  info.Size(),
		}

		if !info.IsDir() {
			// Determine language and count lines
			fileInfo.Language = m.detectLanguage(info.Name())
			fileInfo.Lines = m.countLines(path)
			fileInfo.Purpose = m.determinePurpose(info.Name(), relPath)

			// Add language to list if not already present
			if fileInfo.Language != "" {
				found := false
				for _, lang := range index.Languages {
					if lang == fileInfo.Language {
						found = true
						break
					}
				}
				if !found {
					index.Languages = append(index.Languages, fileInfo.Language)
				}
			}

			index.TotalLines += fileInfo.Lines
		} else {
			index.Directories++
		}

		index.Files = append(index.Files, fileInfo)
		index.MemorySize += info.Size()

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk codebase: %w", err)
	}

	// Save the index
	if err := m.saveIndex(index); err != nil {
		return nil, fmt.Errorf("failed to save index: %w", err)
	}

	return index, nil
}

// LoadIndex loads the stored codebase index
func (m *Manager) LoadIndex() (*CodebaseIndex, error) {
	indexFile := filepath.Join(m.indexPath, "codebase_index.json")

	data, err := os.ReadFile(indexFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read index file: %w", err)
	}

	var index CodebaseIndex
	if err := json.Unmarshal(data, &index); err != nil {
		return nil, fmt.Errorf("failed to unmarshal index: %w", err)
	}

	return &index, nil
}

// ClearIndex removes all indexed data
func (m *Manager) ClearIndex() error {
	return os.RemoveAll(m.indexPath)
}

// saveIndex saves the codebase index to disk
func (m *Manager) saveIndex(index *CodebaseIndex) error {
	indexFile := filepath.Join(m.indexPath, "codebase_index.json")

	jsonData, err := json.MarshalIndent(index, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal index: %w", err)
	}

	if err := os.WriteFile(indexFile, jsonData, 0600); err != nil {
		return fmt.Errorf("failed to write index file: %w", err)
	}

	return nil
}

// detectLanguage determines the programming language based on file extension
func (m *Manager) detectLanguage(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))

	languageMap := map[string]string{
		".go":         "Go",
		".js":         "JavaScript",
		".ts":         "TypeScript",
		".py":         "Python",
		".java":       "Java",
		".cpp":        "C++",
		".c":          "C",
		".cs":         "C#",
		".php":        "PHP",
		".rb":         "Ruby",
		".rs":         "Rust",
		".swift":      "Swift",
		".kt":         "Kotlin",
		".scala":      "Scala",
		".html":       "HTML",
		".css":        "CSS",
		".scss":       "SCSS",
		".sass":       "Sass",
		".json":       "JSON",
		".xml":        "XML",
		".yaml":       "YAML",
		".yml":        "YAML",
		".toml":       "TOML",
		".ini":        "INI",
		".conf":       "Config",
		".sh":         "Shell",
		".bash":       "Bash",
		".zsh":        "Zsh",
		".fish":       "Fish",
		".sql":        "SQL",
		".md":         "Markdown",
		".txt":        "Text",
		".log":        "Log",
		".dockerfile": "Dockerfile",
		".makefile":   "Makefile",
	}

	if lang, exists := languageMap[ext]; exists {
		return lang
	}

	// Special cases for files without extensions
	if filename == "Dockerfile" {
		return "Dockerfile"
	}
	if filename == "Makefile" {
		return "Makefile"
	}
	if filename == "README" {
		return "Markdown"
	}

	return "Unknown"
}

// countLines counts the number of lines in a file
func (m *Manager) countLines(filepath string) int {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return 0
	}

	lines := strings.Split(string(data), "\n")
	return len(lines)
}

// determinePurpose determines the purpose of a file based on its name and path
func (m *Manager) determinePurpose(filename, path string) string {
	lowerName := strings.ToLower(filename)
	lowerPath := strings.ToLower(path)

	// Configuration files
	if strings.Contains(lowerName, "config") || strings.Contains(lowerPath, "config") {
		return "Configuration"
	}
	if strings.Contains(lowerName, "settings") || strings.Contains(lowerPath, "settings") {
		return "Settings"
	}

	// Build and dependency files
	if strings.Contains(lowerName, "package.json") || strings.Contains(lowerName, "go.mod") {
		return "Dependencies"
	}
	if strings.Contains(lowerName, "makefile") || strings.Contains(lowerName, "build") {
		return "Build"
	}

	// Documentation
	if strings.Contains(lowerName, "readme") || strings.Contains(lowerName, "docs") {
		return "Documentation"
	}
	if strings.Contains(lowerName, "license") {
		return "License"
	}

	// Source code
	if strings.Contains(lowerPath, "src") || strings.Contains(lowerPath, "lib") {
		return "Source Code"
	}
	if strings.Contains(lowerPath, "test") || strings.Contains(lowerPath, "spec") {
		return "Tests"
	}

	// Main entry points
	if lowerName == "main.go" || lowerName == "main.py" || lowerName == "index.js" {
		return "Entry Point"
	}

	// Database
	if strings.Contains(lowerName, "migration") || strings.Contains(lowerPath, "db") {
		return "Database"
	}

	// API
	if strings.Contains(lowerPath, "api") || strings.Contains(lowerPath, "routes") {
		return "API"
	}

	// Models
	if strings.Contains(lowerPath, "model") || strings.Contains(lowerPath, "entity") {
		return "Data Model"
	}

	// Utilities
	if strings.Contains(lowerPath, "util") || strings.Contains(lowerPath, "helper") {
		return "Utility"
	}

	return "General"
}
