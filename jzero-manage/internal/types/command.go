package types

import "time"

// CommandResult 命令执行结果
type CommandResult struct {
	Success   bool      `json:"success"`
	ExitCode  int       `json:"exitCode"`
	Output    string    `json:"output"`
	Error     string    `json:"error,omitempty"`
	Duration  int64     `json:"duration"` // 毫秒
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}
