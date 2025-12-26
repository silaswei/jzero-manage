package model

// JZeroConfig JZero配置结构
type JZeroConfig struct {
	Hooks *HooksConfig `yaml:"hooks,omitempty"`
	New   *NewConfig   `yaml:"new,omitempty"`
	Gen   *GenConfig   `yaml:"gen,omitempty"`
}

// HooksConfig 钩子配置
type HooksConfig struct {
	Before []string `yaml:"before,omitempty"`
	After  []string `yaml:"after,omitempty"`
}

// NewConfig 新项目配置
type NewConfig struct {
	Home string `yaml:"home,omitempty"`
}

// GenConfig 生成配置
type GenConfig struct {
	Hooks           *HooksConfig `yaml:"hooks,omitempty"`
	Style           string       `yaml:"style,omitempty"`
	GitChange       bool         `yaml:"git-change,omitempty"`
	Route2Code      bool         `yaml:"route2Code,omitempty"`
	ModelDriver     string       `yaml:"model-driver,omitempty"`
	ModelCache      bool         `yaml:"model-cache,omitempty"`
	ModelIgnoreCols []string     `yaml:"model-ignore-columns,omitempty"`
	ModelSchema     string       `yaml:"model-schema,omitempty"`
}
