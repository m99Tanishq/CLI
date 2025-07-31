package memory

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/m99Tanishq/CLI/internal/api"
	"github.com/m99Tanishq/CLI/internal/config"
)

type FileInfo struct {
	Path     string `json:"path"`
	Name     string `json:"name"`
	IsDir    bool   `json:"is_dir"`
	Language string `json:"language"`
	Lines    int    `json:"lines"`
	Size     int64  `json:"size"`
	Purpose  string `json:"purpose"`
}

type CodebaseIndex struct {
	RootPath    string     `json:"root_path"`
	Files       []FileInfo `json:"files"`
	TotalLines  int        `json:"total_lines"`
	Directories int        `json:"directories"`
	Languages   []string   `json:"languages"`
	MemorySize  int64      `json:"memory_size"`
	LastUpdated time.Time  `json:"last_updated"`
	Model       string     `json:"model"`
}

type Manager struct {
	indexPath string
}

func NewManager() *Manager {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}

	return &Manager{
		indexPath: filepath.Join(homeDir, ".CLI", "memory"),
	}
}

func (m *Manager) IndexCodebase(rootPath string, model string) (*CodebaseIndex, error) {
	if err := os.MkdirAll(m.indexPath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create memory directory: %w", err)
	}

	index := &CodebaseIndex{
		RootPath:    rootPath,
		Files:       []FileInfo{},
		Languages:   []string{},
		LastUpdated: time.Now(),
		Model:       model,
	}

	allowedHidden := map[string]bool{".gitignore": true, ".env": true, ".dockerignore": true}
	skipDirs := map[string]bool{".git": true, "node_modules": true, "vendor": true, "bin": true, "obj": true, "build": true, "dist": true}

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasPrefix(info.Name(), ".") {
			if !allowedHidden[info.Name()] {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}

		if info.IsDir() && skipDirs[info.Name()] {
			return filepath.SkipDir
		}

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
			fileInfo.Language = m.detectLanguage(info.Name())
			fileInfo.Lines = m.countLines(path)
			fileInfo.Purpose = m.determinePurpose(info.Name(), relPath)

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

	if err := m.saveIndex(index); err != nil {
		return nil, fmt.Errorf("failed to save index: %w", err)
	}

	return index, nil
}

func (m *Manager) LoadIndex() (*CodebaseIndex, error) {
	indexPath := filepath.Join(m.indexPath, "index.json")
	data, err := os.ReadFile(indexPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read index file: %w", err)
	}

	var index CodebaseIndex
	if err := json.Unmarshal(data, &index); err != nil {
		return nil, fmt.Errorf("failed to parse index file: %w", err)
	}

	return &index, nil
}

func (m *Manager) ClearIndex() error {
	indexPath := filepath.Join(m.indexPath, "index.json")
	return os.Remove(indexPath)
}

func (m *Manager) saveIndex(index *CodebaseIndex) error {
	indexPath := filepath.Join(m.indexPath, "index.json")
	data, err := json.MarshalIndent(index, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal index: %w", err)
	}

	if err := os.WriteFile(indexPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write index file: %w", err)
	}

	return nil
}

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

	if filename == "Dockerfile" || filename == "Makefile" || filename == "README" {
		return "Markdown"
	}

	return "Unknown"
}

func (m *Manager) countLines(filepath string) int {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return 0
	}

	lines := strings.Split(string(data), "\n")
	return len(lines)
}

func (m *Manager) determinePurpose(filename, path string) string {
	lowerName := strings.ToLower(filename)
	lowerPath := strings.ToLower(path)

	if strings.Contains(lowerName, "config") || strings.Contains(lowerPath, "config") {
		return "Configuration"
	}
	if strings.Contains(lowerName, "settings") || strings.Contains(lowerPath, "settings") {
		return "Settings"
	}

	if strings.Contains(lowerName, "package.json") || strings.Contains(lowerName, "go.mod") {
		return "Dependencies"
	}
	if strings.Contains(lowerName, "makefile") || strings.Contains(lowerName, "build") {
		return "Build"
	}

	if strings.Contains(lowerName, "readme") || strings.Contains(lowerName, "docs") {
		return "Documentation"
	}
	if strings.Contains(lowerName, "license") {
		return "License"
	}

	if strings.Contains(lowerPath, "test") || strings.Contains(lowerName, "test") {
		return "Testing"
	}
	if strings.Contains(lowerPath, "cmd") || strings.Contains(lowerName, "main") {
		return "Entry Point"
	}
	if strings.Contains(lowerPath, "internal") {
		return "Internal Logic"
	}
	if strings.Contains(lowerPath, "pkg") {
		return "Package"
	}

	return "Source Code"
}

func (m *Manager) QueryCodebase(query string) (string, error) {
	cfg, err := config.Load()
	if err != nil {
		return "", fmt.Errorf("failed to load config: %w", err)
	}

	client := api.NewClient(cfg.APIKey, cfg.BaseURL)

	index, err := m.LoadIndex()
	if err != nil {
		return "", fmt.Errorf("failed to load index: %w", err)
	}

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

Please provide a comprehensive answer based on the codebase structure and content.`,
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

	resp, err := client.SendChat(req)
	if err != nil {
		return "", fmt.Errorf("failed to query codebase: %w", err)
	}

	if len(resp.Choices) > 0 {
		return resp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response received")
}

func formatFileStructure(files []FileInfo) string {
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

func formatKeyFiles(files []FileInfo) string {
	var result strings.Builder
	for _, file := range files {
		if isKeyFile(file.Name) {
			result.WriteString(fmt.Sprintf("- %s: %s\n", file.Name, file.Purpose))
		}
	}
	return result.String()
}

func isKeyFile(filename string) bool {
	keyFiles := map[string]bool{
		"main.go": true, "go.mod": true, "go.sum": true, "Makefile": true,
		"README.md": true, "Dockerfile": true, "package.json": true,
		"requirements.txt": true, "Cargo.toml": true, "pom.xml": true,
		"build.gradle": true, "Gemfile": true, "composer.json": true,
		"pubspec.yaml": true,
	}
	return keyFiles[filename]
}
