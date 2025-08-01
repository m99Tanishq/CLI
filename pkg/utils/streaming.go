package utils

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"
)

type StreamingUI struct {
	showProgress bool
	showTimer    bool
	startTime    time.Time
	ctx          context.Context
	cancel       context.CancelFunc
	ui           *ModernUI
}

func NewStreamingUI(showProgress, showTimer bool) *StreamingUI {
	ctx, cancel := context.WithCancel(context.Background())

	return &StreamingUI{
		showProgress: showProgress,
		showTimer:    showTimer,
		startTime:    time.Now(),
		ctx:          ctx,
		cancel:       cancel,
		ui:           NewModernUI(),
	}
}

func (ui *StreamingUI) Start() {
	if ui.showTimer {
		ui.startTime = time.Now()
	}
}

func (ui *StreamingUI) WriteChunk(chunk string) {
	ui.ui.PrintStreamingMessage(chunk)
}

func (ui *StreamingUI) WriteChunkWithProgress(chunk string, chunkCount int) {
	ui.WriteChunk(chunk)
}

func (ui *StreamingUI) End() {
	if ui.showTimer {
		duration := time.Since(ui.startTime)
		ui.ui.PrintStreamingComplete(duration)
	}
	fmt.Println()
}

func (ui *StreamingUI) Cancel() {
	ui.cancel()
}

func (ui *StreamingUI) IsCancelled() bool {
	select {
	case <-ui.ctx.Done():
		return true
	default:
		return false
	}
}

func (ui *StreamingUI) GetContext() context.Context {
	return ui.ctx
}

func (ui *StreamingUI) Cleanup() {}

type Spinner struct {
	frames []string
	index  int
	stop   chan bool
	ui     *ModernUI
	closed bool
}

func NewSpinner() *Spinner {
	return &Spinner{
		frames: []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		index:  0,
		stop:   make(chan bool),
		ui:     NewModernUI(),
	}
}

func (s *Spinner) Start(message string) {
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				s.ui.PrintSpinner(s.frames[s.index], message)
				s.index = (s.index + 1) % len(s.frames)
			case <-s.stop:
				return
			}
		}
	}()
}

func (s *Spinner) Stop() {
	if !s.closed {
		close(s.stop)
		s.closed = true
	}
}

type StreamingHandler struct {
	ui       *StreamingUI
	spinner  *Spinner
	response strings.Builder
}

func NewStreamingHandler(showProgress, showTimer bool) *StreamingHandler {
	return &StreamingHandler{
		ui:      NewStreamingUI(showProgress, showTimer),
		spinner: NewSpinner(),
	}
}

func (sh *StreamingHandler) HandleStream(chunkChan <-chan string, errChan <-chan error) (string, error) {
	sh.ui.Start()
	sh.spinner.Start("Thinking...")

	chunkCount := 0
	hasReceivedChunk := false

	for {
		select {
		case chunk, ok := <-chunkChan:
			if !ok {
				if !hasReceivedChunk {
					sh.spinner.Stop()
				}
				sh.ui.End()
				fmt.Println()
				os.Stdout.Sync()
				return CleanResponse(sh.response.String()), nil
			}

			if !hasReceivedChunk {
				sh.spinner.Stop()
				hasReceivedChunk = true
				sh.ui.ui.PrintPrompt("🤖 AI: ")
			}

			sh.response.WriteString(chunk)
			sh.ui.WriteChunkWithProgress(chunk, chunkCount)
			chunkCount++

		case err, ok := <-errChan:
			if ok {
				if !hasReceivedChunk {
					sh.spinner.Stop()
				}
				sh.ui.End()
				return CleanResponse(sh.response.String()), err
			}

		case <-sh.ui.GetContext().Done():
			PrintCancellationInfo()
			return "", fmt.Errorf("operation cancelled by user")
		}
	}
}

func (sh *StreamingHandler) Cleanup() {
	sh.ui.Cleanup()
}
