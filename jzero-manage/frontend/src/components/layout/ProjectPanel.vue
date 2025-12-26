<template>
  <div class="project-panel">
    <!-- 顶部工具栏 -->
    <div class="panel-header">
      <n-space justify="space-between" align="center">
        <n-text strong>项目列表</n-text>
        <n-button-group>
          <n-button size="small" @click="handleRefresh">
            <template #icon>
              <n-icon><RefreshIcon /></n-icon>
            </template>
            刷新
          </n-button>
        </n-button-group>
      </n-space>
    </div>

    <!-- 项目搜索 -->
    <div class="panel-search">
      <n-input
        v-model:value="searchKeyword"
        placeholder="搜索项目..."
        clearable
        size="small"
      >
        <template #prefix>
          <n-icon><SearchIcon /></n-icon>
        </template>
      </n-input>
    </div>

    <!-- 项目树 -->
    <div class="panel-content">
      <n-spin :show="projectStore.loading">
        <n-empty v-if="filteredProjects.length === 0 && !projectStore.loading" description="暂无项目" />
        <n-tree
          v-else
          :data="filteredProjects"
          :node-props="nodeProps"
          :render-prefix="renderPrefix"
          :render-suffix="renderSuffix"
          key-field="id"
          children-field="children"
          expand-on-click
          selectable
          @update:selected-keys="handleSelectProject"
        />
      </n-spin>
    </div>

    <!-- 项目右键菜单 -->
    <n-dropdown
      :show="contextMenuVisible"
      :options="contextMenuOptions"
      :x="contextMenuX"
      :y="contextMenuY"
      placement="bottom-start"
      @clickoutside="contextMenuVisible = false"
      @select="handleContextMenuSelect"
    />
  </div>
</template>

<script setup>
import { ref, computed, h } from 'vue'
import { NIcon, NTag, useMessage, useDialog } from 'naive-ui'
import { FolderOutlined, ExtensionOutlined, CheckCircleOutlined } from '@vicons/material'
import { useProjectStore } from '@/stores/project'
import { useFileStore } from '@/stores/file'
import { CommandGen } from '@/api/command'

const projectStore = useProjectStore()
const fileStore = useFileStore()
const message = useMessage()
const dialog = useDialog()

// 图标
const RefreshIcon = () => h('svg', { viewBox: '0 0 24 24', fill: 'currentColor' }, [
  h('path', { d: 'M17.65 6.35C16.2 4.9 14.21 4 12 4c-4.42 0-7.99 3.58-7.99 8s3.57 8 7.99 8c3.73 0 6.84-2.55 7.73-6h-2.08c-.82 2.33-3.04 4-5.65 4-3.31 0-6-2.69-6-6s2.69-6 6-6c1.66 0 3.14.69 4.22 1.78L13 11h7V4l-2.35 2.35z' })
])
const SearchIcon = () => h('svg', { viewBox: '0 0 24 24', fill: 'currentColor' }, [
  h('path', { d: 'M15.5 14h-.79l-.28-.27C15.41 12.59 16 11.11 16 9.5 16 5.91 13.09 3 9.5 3S3 5.91 3 9.5 5.91 16 9.5 16c1.61 0 3.09-.59 4.23-1.57l.27.28v.79l5 4.99L20.49 19l-4.99-5zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z' })
])

// 搜索
const searchKeyword = ref('')
const contextMenuVisible = ref(false)
const contextMenuX = ref(0)
const contextMenuY = ref(0)
const selectedProjectNode = ref(null)

// 过滤项目
const filteredProjects = computed(() => {
  if (!searchKeyword.value) {
    return projectStore.projectTree
  }
  // 简单的搜索过滤
  return filterTree(projectStore.projectTree, searchKeyword.value)
})

// 递归过滤树
const filterTree = (nodes, keyword) => {
  const result = []
  for (const node of nodes) {
    if (node.name.toLowerCase().includes(keyword.toLowerCase())) {
      result.push(node)
    } else if (node.children && node.children.length > 0) {
      const filtered = filterTree(node.children, keyword)
      if (filtered.length > 0) {
        result.push({ ...node, children: filtered })
      }
    }
  }
  return result
}

// 右键菜单选项
const contextMenuOptions = computed(() => [
  {
    label: '刷新项目',
    key: 'refresh',
    icon: () => h('svg', { viewBox: '0 0 24 24', fill: 'currentColor' }, [
      h('path', { d: 'M17.65 6.35C16.2 4.9 14.21 4 12 4c-4.42 0-7.99 3.58-7.99 8s3.57 8 7.99 8c3.73 0 6.84-2.55 7.73-6h-2.08c-.82 2.33-3.04 4-5.65 4-3.31 0-6-2.69-6-6s2.69-6 6-6c1.66 0 3.14.69 4.22 1.78L13 11h7V4l-2.35 2.35z' })
    ])
  },
  {
    label: '生成代码',
    key: 'generate',
    icon: () => h('svg', { viewBox: '0 0 24 24', fill: 'currentColor' }, [
      h('path', { d: 'M9.4 16.6L4.8 12l4.6-4.6L8 6l-6 6 6 6 1.4-1.4zm5.2 0l4.6-4.6-4.6-4.6L16 6l6 6-6 6-1.4-1.4z' })
    ])
  },
  {
    label: '编辑配置',
    key: 'config',
    disabled: !selectedProjectNode.value?.hasConfig,
    icon: () => h('svg', { viewBox: '0 0 24 24', fill: 'currentColor' }, [
      h('path', { d: 'M19.14 12.94c.04-.3.06-.61.06-.94 0-.32-.02-.64-.07-.94l2.03-1.58c.18-.14.23-.41.12-.61l-1.92-3.32c-.12-.22-.37-.29-.59-.22l-2.39.96c-.5-.38-1.03-.7-1.62-.94l-.36-2.54c-.04-.24-.24-.41-.48-.41h-3.84c-.24 0-.43.17-.47.41l-.36 2.54c-.59.24-1.13.57-1.62.94l-2.39-.96c-.22-.08-.47 0-.59.22L5.09 8.87c-.12.21-.08.47.12.61l2.03 1.58c-.05.3-.09.63-.09.94s.02.64.07.94l-2.03 1.58c-.18.14-.23.41-.12.61l1.92 3.32c.12.22.37.29.59.22l2.39-.96c.5.38 1.03.7 1.62.94l.36 2.54c.05.24.24.41.48.41h3.84c.24 0 .44-.17.47-.41l.36-2.54c.59-.24 1.13-.56 1.62-.94l2.39.96c.22.08.47 0 .59-.22l1.92-3.32c.12-.22.07-.47-.12-.61l-2.01-1.58zM12 15.6c-1.98 0-3.6-1.62-3.6-3.6s1.62-3.6 3.6-3.6 3.6 1.62 3.6 3.6-1.62 3.6-3.6 3.6z' })
    ])
  }
])

// 节点属性
const nodeProps = ({ option }) => {
  return {
    onContextmenu(e) {
      e.preventDefault()
      selectedProjectNode.value = option
      contextMenuX.value = e.clientX
      contextMenuY.value = e.clientY
      contextMenuVisible.value = true
    }
  }
}

// 渲染前缀（图标）
const renderPrefix = ({ option }) => {
  if (option.type === 'plugin') {
    return h(NIcon, null, { default: () => h(ExtensionOutlined) })
  }
  return h(NIcon, null, { default: () => h(FolderOutlined) })
}

// 渲染后缀（标签）
const renderSuffix = ({ option }) => {
  if (option.hasConfig) {
    return h(NIcon, { style: { color: '#18a058' } }, { default: () => h(CheckCircleOutlined) })
  }
  return null
}

// 选择项目
const handleSelectProject = (keys) => {
  if (keys.length > 0) {
    const projectId = keys[0]
    projectStore.selectProject(projectId)

    // 加载文件树
    if (projectStore.currentProject) {
      fileStore.loadFileTree(projectStore.currentProject.path)
    }
  }
}

// 刷新
const handleRefresh = () => {
  projectStore.scanProjects('/Users/Apple/GolandProjects')
}

// 右键菜单选择
const handleContextMenuSelect = (key) => {
  const project = selectedProjectNode.value
  if (!project) return

  switch (key) {
    case 'refresh':
      projectStore.scanProjects('/Users/Apple/GolandProjects')
      break
    case 'generate':
      dialog.warning({
        title: '确认生成',
        content: '确定要生成代码吗？此操作将覆盖已生成的文件。',
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: async () => {
          try {
            const result = await CommandGen(project.path)
            if (result.success) {
              dialog.success({
                title: '生成成功',
                content: result.output || '代码生成完成',
                positiveText: '确定'
              })
            } else {
              dialog.error({
                title: '生成失败',
                content: result.error || '代码生成失败',
                positiveText: '确定'
              })
            }
          } catch (e) {
            message.error('代码生成异常: ' + e.message)
          }
        }
      })
      break
    case 'config':
      // 打开配置编辑器
      window.dispatchEvent(new CustomEvent('open-config-editor', { detail: { projectPath: project.path } }))
      break
  }

  contextMenuVisible.value = false
}
</script>

<style scoped>
.project-panel {
  display: flex;
  flex-direction: column;
  height: 100%;
  background: #16213e;
}

.panel-header {
  padding: 12px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  background: #1a1a2e;
}

.panel-search {
  padding: 8px 12px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  background: #16213e;
}

.panel-content {
  flex: 1;
  overflow: auto;
  padding: 8px;
  background: #16213e;
}
</style>
