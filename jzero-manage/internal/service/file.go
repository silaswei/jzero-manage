package service

import (
	"os"
	"path/filepath"

	"jzero-manage/internal/types"
)

// FileService 文件服务
type FileService struct{}

// NewFileService 创建文件服务
func NewFileService() *FileService {
	return &FileService{}
}

// GetTree 获取文件树
func (s *FileService) GetTree(projectPath string, dir string) ([]types.FileNode, error) {
	targetDir := dir
	if targetDir == "" {
		targetDir = filepath.Join(projectPath, "desc")
	}

	fullPath := filepath.Join(projectPath, targetDir)
	nodes, err := s.buildTree(fullPath, projectPath)
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

// buildTree 构建文件树
func (s *FileService) buildTree(dirPath string, basePath string) ([]types.FileNode, error) {
	nodes := make([]types.FileNode, 0)

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		// 如果目录不存在，返回空列表
		if os.IsNotExist(err) {
			return nodes, nil
		}
		return nil, err
	}

	for _, entry := range entries {
		name := entry.Name()
		fullPath := filepath.Join(dirPath, name)

		// 跳过隐藏文件
		if len(name) > 0 && name[0] == '.' {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			continue
		}

		relPath, _ := filepath.Rel(basePath, fullPath)
		ext := filepath.Ext(name)

		node := types.FileNode{
			ID:       relPath,
			Name:     name,
			Path:     relPath,
			IsLeaf:   !entry.IsDir(),
			Size:     info.Size(),
			Modified: info.ModTime().Unix(),
		}

		if entry.IsDir() {
			node.Type = "directory"
			children, err := s.buildTree(fullPath, basePath)
			if err == nil {
				node.Children = children
			}
		} else {
			node.Type = "file"
			node.Ext = ext
		}

		nodes = append(nodes, node)
	}

	return nodes, nil
}

// Read 读取文件内容
func (s *FileService) Read(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// Write 写入文件
func (s *FileService) Write(filePath string, content string) error {
	// 确保目录存在
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(filePath, []byte(content), 0644)
}

// Create 创建文件或目录
func (s *FileService) Create(filePath string, isDir bool) error {
	if isDir {
		return os.MkdirAll(filePath, 0755)
	}

	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

// Delete 删除文件或目录
func (s *FileService) Delete(filePath string) error {
	return os.RemoveAll(filePath)
}

// Rename 重命名文件
func (s *FileService) Rename(oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
}

// GetInfo 获取文件信息
func (s *FileService) GetInfo(filePath string) (*types.FileInfo, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	name := filepath.Base(filePath)
	ext := filepath.Ext(name)

	return &types.FileInfo{
		Name:     name,
		Path:     filePath,
		Type:     getFileType(info),
		Size:     info.Size(),
		Modified: info.ModTime(),
		Ext:      ext,
		IsHidden: len(name) > 0 && name[0] == '.',
	}, nil
}

// Exists 检查文件是否存在
func (s *FileService) Exists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

// getFileType 获取文件类型
func getFileType(info os.FileInfo) string {
	if info.IsDir() {
		return "directory"
	}
	return "file"
}
