package service

import (
	"os/exec"
	"time"

	"jzero-manage/internal/types"
)

// CommandService 命令服务
type CommandService struct{}

// NewCommandService 创建命令服务
func NewCommandService() *CommandService {
	return &CommandService{}
}

// Execute 执行命令
func (s *CommandService) Execute(workDir string, command string, args []string) (*types.CommandResult, error) {
	start := time.Now()

	cmd := exec.Command(command, args...)
	cmd.Dir = workDir

	output, err := cmd.CombinedOutput()

	result := &types.CommandResult{
		StartTime: start,
		EndTime:   time.Now(),
		Duration:  time.Since(start).Milliseconds(),
		Output:    string(output),
	}

	if err != nil {
		result.Success = false
		result.Error = err.Error()
		if exitError, ok := err.(*exec.ExitError); ok {
			result.ExitCode = exitError.ExitCode()
		}
	} else {
		result.Success = true
		result.ExitCode = 0
	}

	return result, nil
}

// Gen 执行 jzero gen
func (s *CommandService) Gen(projectPath string) (*types.CommandResult, error) {
	return s.Execute(projectPath, "jzero", []string{"gen"})
}

// Swagger 执行 jzero gen swagger
func (s *CommandService) Swagger(projectPath string) (*types.CommandResult, error) {
	return s.Execute(projectPath, "jzero", []string{"gen", "swagger"})
}
