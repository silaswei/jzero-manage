package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"jzero-manage/internal/model"
	"jzero-manage/internal/service"
	"jzero-manage/internal/types"
	"jzero-manage/internal/util"

	"gopkg.in/yaml.v3"
)

// App struct
type App struct {
	ctx      context.Context
	scanner  *util.ProjectScanner
	fileSrv  *service.FileService
	cmdSrv   *service.CommandService
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		scanner: util.NewProjectScanner(),
		fileSrv: service.NewFileService(),
		cmdSrv:  service.NewCommandService(),
	}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// ==================== 项目管理API ====================

// ProjectScan 扫描项目
func (a *App) ProjectScan(rootPath string) ([]types.ProjectNode, error) {
	projects, err := a.scanner.Scan(rootPath)
	if err != nil {
		return nil, err
	}

	return util.BuildProjectTree(projects), nil
}

// ProjectInfo 获取项目信息
func (a *App) ProjectInfo(projectPath string) (*types.ProjectInfo, error) {
	info := &types.ProjectInfo{
		Name:     filepath.Base(projectPath),
		Path:     projectPath,
		HasConfig: a.scanner.HasConfigFile(projectPath),
	}

	// 统计desc目录下的文件
	descPath := filepath.Join(projectPath, "desc")
	apiPath := filepath.Join(descPath, "api")
	protoPath := filepath.Join(descPath, "proto")
	sqlPath := filepath.Join(descPath, "sql")

	info.DescDirs = []string{}

	if _, err := os.Stat(apiPath); err == nil {
		info.DescDirs = append(info.DescDirs, "api")
		files, _ := os.ReadDir(apiPath)
		info.APICount = len(files)
	}

	if _, err := os.Stat(protoPath); err == nil {
		info.DescDirs = append(info.DescDirs, "proto")
		files, _ := os.ReadDir(protoPath)
		info.ProtoCount = len(files)
	}

	if _, err := os.Stat(sqlPath); err == nil {
		info.DescDirs = append(info.DescDirs, "sql")
		files, _ := os.ReadDir(sqlPath)
		info.SQLCount = len(files)
	}

	return info, nil
}

// ProjectList 获取所有项目
func (a *App) ProjectList() ([]types.ProjectInfo, error) {
	// 这里可以返回缓存的列表
	return []types.ProjectInfo{}, nil
}

// ProjectGen 生成代码
func (a *App) ProjectGen(projectPath string) (string, error) {
	result, err := a.cmdSrv.Gen(projectPath)
	if err != nil {
		return "", err
	}

	if result.Success {
		return result.Output, nil
	}
	return "", fmt.Errorf(result.Error)
}

// ==================== 文件管理API ====================

// FileTree 获取文件树
func (a *App) FileTree(projectPath string, dir string) ([]types.FileNode, error) {
	return a.fileSrv.GetTree(projectPath, dir)
}

// FileList 获取文件列表
func (a *App) FileList(projectPath string, dir string) ([]types.FileInfo, error) {
	targetDir := dir
	if targetDir == "" {
		targetDir = filepath.Join(projectPath, "desc")
	}

	fullPath := filepath.Join(projectPath, targetDir)
	entries, err := os.ReadDir(fullPath)
	if err != nil {
		return nil, err
	}

	files := make([]types.FileInfo, 0)
	for _, entry := range entries {
		name := entry.Name()
		if len(name) > 0 && name[0] == '.' {
			continue
		}

		info, _ := entry.Info()
		filePath := filepath.Join(fullPath, name)
		ext := filepath.Ext(name)

		fileType := "file"
		if entry.IsDir() {
			fileType = "directory"
		}

		files = append(files, types.FileInfo{
			Name:     name,
			Path:     filePath,
			Type:     fileType,
			Size:     info.Size(),
			Modified: info.ModTime(),
			Ext:      ext,
			IsHidden: false,
		})
	}

	return files, nil
}

// FileRead 读取文件
func (a *App) FileRead(filePath string) (string, error) {
	return a.fileSrv.Read(filePath)
}

// FileWrite 写入文件
func (a *App) FileWrite(filePath string, content string) error {
	return a.fileSrv.Write(filePath, content)
}

// FileCreate 创建文件或目录
func (a *App) FileCreate(filePath string, isDir bool) error {
	return a.fileSrv.Create(filePath, isDir)
}

// FileDelete 删除文件
func (a *App) FileDelete(filePath string) error {
	return a.fileSrv.Delete(filePath)
}

// FileRename 重命名文件
func (a *App) FileRename(oldPath string, newPath string) error {
	return a.fileSrv.Rename(oldPath, newPath)
}

// FileInfo 获取文件信息
func (a *App) FileInfo(filePath string) (*types.FileInfo, error) {
	return a.fileSrv.GetInfo(filePath)
}

// FileExists 检查文件是否存在
func (a *App) FileExists(filePath string) bool {
	return a.fileSrv.Exists(filePath)
}

// ==================== 配置管理API ====================

// ConfigRead 读取配置文件
func (a *App) ConfigRead(projectPath string) (*model.JZeroConfig, error) {
	configPath := filepath.Join(projectPath, ".jzero.yaml")

	content, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config model.JZeroConfig
	if err := yaml.Unmarshal(content, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// ConfigWrite 写入配置文件
func (a *App) ConfigWrite(projectPath string, config *model.JZeroConfig) error {
	configPath := filepath.Join(projectPath, ".jzero.yaml")

	content, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, content, 0644)
}

// ConfigValidate 验证配置
func (a *App) ConfigValidate(projectPath string) (bool, []string, error) {
	config, err := a.ConfigRead(projectPath)
	if err != nil {
		return false, []string{err.Error()}, err
	}

	errors := make([]string, 0)

	// 验证逻辑
	if config.Gen != nil {
		if config.Gen.Style == "" {
			errors = append(errors, "gen.style is required")
		}
	}

	return len(errors) == 0, errors, nil
}

// ==================== 命令执行API ====================

// CommandExecute 执行命令
func (a *App) CommandExecute(workDir string, command string, args []string) (*types.CommandResult, error) {
	return a.cmdSrv.Execute(workDir, command, args)
}

// CommandGen 执行jzero gen
func (a *App) CommandGen(projectPath string) (*types.CommandResult, error) {
	return a.cmdSrv.Gen(projectPath)
}

// CommandSwagger 执行jzero gen swagger
func (a *App) CommandSwagger(projectPath string) (*types.CommandResult, error) {
	return a.cmdSrv.Swagger(projectPath)
}

// ==================== 旧API（兼容） ====================

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
