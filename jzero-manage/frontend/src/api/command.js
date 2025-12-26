// 命令相关API
export const CommandExecute = (workDir, command, args) => window.go.main.App.CommandExecute(workDir, command, args)
export const CommandGen = (projectPath) => window.go.main.App.CommandGen(projectPath)
export const CommandSwagger = (projectPath) => window.go.main.App.CommandSwagger(projectPath)
