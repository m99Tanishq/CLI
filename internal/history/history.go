package history

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// ChatSession represents a chat session
type ChatSession struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Messages  []Message `json:"messages"`
}

// Message represents a chat message in history
type Message struct {
	Role      string    `json:"role"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

// Manager handles chat history operations
type Manager struct {
	historyPath string
}

// NewManager creates a new history manager
func NewManager() *Manager {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}

	return &Manager{
		historyPath: filepath.Join(homeDir, ".CLI", "history"),
	}
}

// SaveSession saves a chat session to history
func (m *Manager) SaveSession(session *ChatSession) error {
	if err := os.MkdirAll(m.historyPath, 0755); err != nil {
		return fmt.Errorf("failed to create history directory: %w", err)
	}

	filename := filepath.Join(m.historyPath, session.ID+".json")

	jsonData, err := json.MarshalIndent(session, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal session: %w", err)
	}

	if err := os.WriteFile(filename, jsonData, 0600); err != nil {
		return fmt.Errorf("failed to write session file: %w", err)
	}

	return nil
}

// LoadSession loads a chat session from history
func (m *Manager) LoadSession(sessionID string) (*ChatSession, error) {
	filename := filepath.Join(m.historyPath, sessionID+".json")

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read session file: %w", err)
	}

	var session ChatSession
	if err := json.Unmarshal(data, &session); err != nil {
		return nil, fmt.Errorf("failed to unmarshal session: %w", err)
	}

	return &session, nil
}

// ListSessions lists all available chat sessions
func (m *Manager) ListSessions() ([]ChatSession, error) {
	if err := os.MkdirAll(m.historyPath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create history directory: %w", err)
	}

	files, err := os.ReadDir(m.historyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read history directory: %w", err)
	}

	var sessions []ChatSession
	for _, file := range files {
		if filepath.Ext(file.Name()) != ".json" {
			continue
		}

		sessionID := file.Name()[:len(file.Name())-5] // Remove .json extension
		session, err := m.LoadSession(sessionID)
		if err != nil {
			continue // Skip corrupted files
		}

		sessions = append(sessions, *session)
	}

	return sessions, nil
}

// DeleteSession deletes a chat session from history
func (m *Manager) DeleteSession(sessionID string) error {
	filename := filepath.Join(m.historyPath, sessionID+".json")

	if err := os.Remove(filename); err != nil {
		return fmt.Errorf("failed to delete session file: %w", err)
	}

	return nil
}

// ClearHistory deletes all chat sessions
func (m *Manager) ClearHistory() error {
	if err := os.RemoveAll(m.historyPath); err != nil {
		return fmt.Errorf("failed to clear history: %w", err)
	}

	return nil
}
