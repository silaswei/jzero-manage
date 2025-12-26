package types

import "time"

// FileNode 文件树节点
type FileNode struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Path     string     `json:"path"`
	Type     string     `json:"type"` // "file" | "directory"
	Ext      string     `json:"ext,omitempty"`
	Size     int64      `json:"size,omitempty"`
	Modified int64      `json:"modified,omitempty"`
	Children []FileNode `json:"children,omitempty"`
	IsLeaf   bool       `json:"isLeaf"`
}

// FileInfo 文件信息
type FileInfo struct {
	Name     string    `json:"name"`
	Path     string    `json:"path"`
	Type     string    `json:"type"` // "file" | "directory"
	Size     int64     `json:"size"`
	Modified time.Time `json:"modified"`
	Ext      string    `json:"ext"`
	IsHidden bool      `json:"isHidden"`
}

// FileContent 文件内容
type FileContent struct {
	Path     string `json:"path"`
	Content  string `json:"content"`
	Encoding string `json:"encoding"` // "utf-8" | "binary"
	ReadOnly bool   `json:"readOnly"`
}

// FileTypeMap 文件类型映射
var FileTypeMap = map[string]string{
	".api":   "api",
	".proto": "proto",
	".sql":   "sql",
	".yaml":  "yaml",
	".yml":   "yaml",
	".go":    "go",
	".js":    "javascript",
	".ts":    "typescript",
	".vue":   "vue",
	".json":  "json",
	".md":    "markdown",
}

// GetFileType 获取文件类型
func GetFileType(filename string) string {
	for ext, fileType := range FileTypeMap {
		if len(filename) >= len(ext) && filename[len(filename)-len(ext):] == ext {
			return fileType
		}
	}
	return "unknown"
}
