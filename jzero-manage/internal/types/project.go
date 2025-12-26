package types

// ProjectNode 项目树节点
type ProjectNode struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	Path       string         `json:"path"`
	Type       string         `json:"type"` // "project" | "plugin"
	IsLeaf     bool           `json:"isLeaf"`
	Children   []ProjectNode  `json:"children,omitempty"`
	ConfigPath string         `json:"configPath,omitempty"`
	HasConfig  bool           `json:"hasConfig"`
}

// ProjectInfo 项目信息
type ProjectInfo struct {
	Name      string   `json:"name"`
	Path      string   `json:"path"`
	RootPath  string   `json:"rootPath"`
	IsPlugin  bool     `json:"isPlugin"`
	Parent    string   `json:"parent,omitempty"`
	HasConfig bool     `json:"hasConfig"`
	DescDirs  []string `json:"descDirs"`
	APICount  int      `json:"apiCount"`
	ProtoCount int     `json:"protoCount"`
	SQLCount   int     `json:"sqlCount"`
}

// ScanConfig 扫描配置
type ScanConfig struct {
	RootPath       string   `json:"rootPath"`
	MaxDepth       int      `json:"maxDepth"`
	IncludePlugins bool     `json:"includePlugins"`
	ExcludeDirs    []string `json:"excludeDirs"`
}
