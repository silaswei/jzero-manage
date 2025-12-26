package util

import (
	"os"
	"path/filepath"
	"strings"

	"jzero-manage/internal/types"
)

// ScannedProject 扫描到的项目
type ScannedProject struct {
	Name      string
	Path      string
	IsPlugin  bool
	Parent    string
	HasConfig bool
}

// ProjectScanner 项目扫描器
type ProjectScanner struct {
	maxDepth       int
	includePlugins bool
	excludeDirs    []string
}

// NewProjectScanner 创建项目扫描器
func NewProjectScanner() *ProjectScanner {
	return &ProjectScanner{
		maxDepth:       3,
		includePlugins: true,
		excludeDirs:    []string{".git", "node_modules", "vendor", ".idea"},
	}
}

// Scan 扫描指定目录
func (s *ProjectScanner) Scan(rootPath string) ([]ScannedProject, error) {
	projects := make([]ScannedProject, 0)

	entries, err := os.ReadDir(rootPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		if s.shouldExclude(entry.Name()) {
			continue
		}

		fullPath := filepath.Join(rootPath, entry.Name())

		// 检查是否是JZero项目
		if s.isJZeroProject(fullPath) {
			project := ScannedProject{
				Name:      entry.Name(),
				Path:      fullPath,
				IsPlugin:  false,
				HasConfig: s.HasConfigFile(fullPath),
			}
			projects = append(projects, project)

			// 扫描plugins子项目
			if s.includePlugins {
				plugins := s.scanPlugins(fullPath)
				projects = append(projects, plugins...)
			}
		}
	}

	return projects, nil
}

// scanPlugins 扫描plugins目录
func (s *ProjectScanner) scanPlugins(projectPath string) []ScannedProject {
	plugins := make([]ScannedProject, 0)
	pluginsPath := filepath.Join(projectPath, "plugins")

	entries, err := os.ReadDir(pluginsPath)
	if err != nil {
		return plugins
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		fullPath := filepath.Join(pluginsPath, entry.Name())

		if s.isJZeroProject(fullPath) {
			plugin := ScannedProject{
				Name:      entry.Name(),
				Path:      fullPath,
				IsPlugin:  true,
				Parent:    projectPath,
				HasConfig: s.HasConfigFile(fullPath),
			}
			plugins = append(plugins, plugin)
		}
	}

	return plugins
}

// isJZeroProject 判断是否是JZero项目
func (s *ProjectScanner) isJZeroProject(path string) bool {
	// 检查是否有go.mod
	if _, err := os.Stat(filepath.Join(path, "go.mod")); os.IsNotExist(err) {
		return false
	}

	// 检查是否有desc目录或.jzero.yaml
	descPath := filepath.Join(path, "desc")
	configPath := filepath.Join(path, ".jzero.yaml")

	if _, err := os.Stat(descPath); err == nil {
		return true
	}

	if _, err := os.Stat(configPath); err == nil {
		return true
	}

	return false
}

// HasConfigFile 检查是否有配置文件
func (s *ProjectScanner) HasConfigFile(path string) bool {
	_, err := os.Stat(filepath.Join(path, ".jzero.yaml"))
	return err == nil
}

// shouldExclude 是否应该排除
func (s *ProjectScanner) shouldExclude(name string) bool {
	for _, exclude := range s.excludeDirs {
		if strings.HasPrefix(name, exclude) {
			return true
		}
	}
	return false
}

// BuildProjectTree 构建项目树
func BuildProjectTree(projects []ScannedProject) []types.ProjectNode {
	// 构建主项目和子项目的映射
	mainProjects := make(map[string]*types.ProjectNode)
	pluginProjects := make(map[string][]types.ProjectNode)

	for _, p := range projects {
		node := types.ProjectNode{
			ID:        generateID(p.Path),
			Name:      p.Name,
			Path:      p.Path,
			Type:      "project",
			IsLeaf:    true,
			HasConfig: p.HasConfig,
		}

		if p.IsPlugin {
			node.Type = "plugin"
			node.IsLeaf = true
			pluginProjects[p.Parent] = append(pluginProjects[p.Parent], node)
		} else {
			mainProjects[p.Path] = &node
		}
	}

	// 组装树形结构
	result := make([]types.ProjectNode, 0)
	for _, main := range mainProjects {
		if plugins, ok := pluginProjects[main.Path]; ok {
			main.Children = plugins
			main.IsLeaf = false
		}
		result = append(result, *main)
	}

	return result
}

// generateID 生成唯一ID
func generateID(path string) string {
	// 简单的ID生成：使用路径的hash或直接使用路径
	// 这里使用路径本身作为ID，实际使用时可以考虑hash
	return strings.ReplaceAll(path, "/", "_")
}
