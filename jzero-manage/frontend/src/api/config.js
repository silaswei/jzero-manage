// 配置相关API
export const ConfigRead = (projectPath) => window.go.main.App.ConfigRead(projectPath)
export const ConfigWrite = (projectPath, config) => window.go.main.App.ConfigWrite(projectPath, config)
export const ConfigValidate = (projectPath) => window.go.main.App.ConfigValidate(projectPath)
