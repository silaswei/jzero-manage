<template>
  <n-config-provider :theme="darkTheme" :theme-overrides="themeOverrides">
    <div class="main-layout">
      <!-- 欢迎页面 -->
      <WelcomePage v-if="showWelcome" />

      <!-- 主界面 -->
      <n-layout v-else has-sider position="absolute" style="height: 100vh">
        <!-- 左侧项目面板 (20%) -->
        <n-layout-sider
          width="20%"
          :native-scrollbar="false"
          bordered
          show-trigger="arrow-circle"
        >
          <ProjectPanel />
        </n-layout-sider>

        <!-- 中间文件浏览器 (30%) -->
        <n-layout-sider
          width="30%"
          :native-scrollbar="false"
          bordered
          show-trigger="arrow-circle"
        >
          <FileExplorer />
        </n-layout-sider>

        <!-- 右侧代码编辑器 (50%) -->
        <n-layout-content native-scrollbar="false">
          <CodeEditor />
        </n-layout-content>
      </n-layout>

      <!-- 配置编辑器模态框 -->
      <ConfigEditorModal ref="configEditorRef" />
    </div>
  </n-config-provider>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { darkTheme } from 'naive-ui'
import WelcomePage from './WelcomePage.vue'
import ProjectPanel from './ProjectPanel.vue'
import FileExplorer from './FileExplorer.vue'
import CodeEditor from './CodeEditor.vue'
import ConfigEditorModal from '../config/ConfigEditorModal.vue'
import { useProjectStore } from '@/stores/project'

const projectStore = useProjectStore()
const configEditorRef = ref(null)

// 是否显示欢迎页面
const showWelcome = computed(() => {
  return projectStore.projects.length === 0
})

// 主题覆盖
const themeOverrides = {
  common: {
    primaryColor: '#667eea',
    primaryColorHover: '#764ba2',
    primaryColorPressed: '#5a67d8',
    primaryColorSuppl: '#667eea',
  },
  Layout: {
    color: '#1a1a2e',
    siderColor: '#16213e',
    headerColor: '#1a1a2e',
    footerColor: '#1a1a2e',
  },
  Card: {
    color: '#1e1e2e',
    colorModal: '#1e1e2e',
  },
  Menu: {
    color: '#16213e',
  },
  Tree: {
    color: '#1e1e2e',
  },
  Tabs: {
    colorSegment: '#1e1e2e',
    tabGapSmall: '10px',
    tabGapMedium: '12px',
    tabGapLarge: '14px',
  },
}

// 监听打开项目对话框事件
onMounted(() => {
  window.addEventListener('open-project-dialog', handleOpenProjectDialog)

  // 如果之前有打开过项目，可以自动加载（从localStorage等）
  // 暂时不自动扫描，让用户主动选择
})

const handleOpenProjectDialog = async () => {
  // 使用 Wails 的目录选择对话框
  // 这里需要调用后端API
  try {
    // 暂时使用默认路径，后续可以实现目录选择对话框
    const projectPath = '/Users/Apple/GolandProjects'
    await projectStore.scanProjects(projectPath)
  } catch (error) {
    console.error('Failed to open project:', error)
  }
}

// 暴露给子组件调用
defineExpose({
  configEditorRef
})
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  background: #1a1a2e;
  color: #e0e0e0;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
}

#app {
  width: 100%;
  height: 100vh;
  overflow: hidden;
  background: #1a1a2e;
}
</style>

<style scoped>
.main-layout {
  width: 100%;
  height: 100vh;
  overflow: hidden;
}
</style>
