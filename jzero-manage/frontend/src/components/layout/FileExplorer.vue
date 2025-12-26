<template>
  <div class="file-explorer">
    <!-- 顶部工具栏 -->
    <div class="explorer-header">
      <n-space justify="space-between" align="center">
        <n-text strong>文件资源管理器</n-text>
        <n-button-group size="small">
          <n-button @click="handleRefresh">
            <template #icon>
              <n-icon><RefreshIcon /></n-icon>
            </template>
          </n-button>
          <n-button @click="handleOpenFolder">
            <template #icon>
              <n-icon><FolderOpenIcon /></n-icon>
            </template>
          </n-button>
        </n-button-group>
      </n-space>
    </div>

    <!-- 当前项目信息 -->
    <div v-if="projectStore.currentProject" class="explorer-info">
      <n-text depth="3" style="font-size: 12px;">{{ projectStore.currentProject.name }}</n-text>
    </div>

    <!-- 文件树 -->
    <div class="explorer-content">
      <n-spin :show="fileStore.loading">
        <n-empty v-if="fileStore.fileTree.length === 0 && !fileStore.loading" description="暂无文件" />
        <n-tree
          v-else
          :data="fileStore.fileTree"
          :node-props="fileNodeProps"
          :render-prefix="renderFileIcon"
          key-field="id"
          children-field="children"
          expand-on-click
          selectable
          @update:selected-keys="handleSelectFile"
        />
      </n-spin>
    </div>

    <!-- 文件右键菜单 -->
    <n-dropdown
      :show="contextMenuVisible"
      :options="fileContextMenuOptions"
      :x="contextMenuX"
      :y="contextMenuY"
      placement="bottom-start"
      @clickoutside="contextMenuVisible = false"
      @select="handleFileContextMenu"
    />
  </div>
</template>

<script setup>
import { ref, computed, h, watch } from 'vue'
import { NIcon, useMessage } from 'naive-ui'
import {
  FolderOutlined,
  InsertDriveFileOutlined,
  DescriptionOutlined,
  CodeOutlined,
  StorageOutlined
} from '@vicons/material'
import { useFileStore } from '@/stores/file'
import { useProjectStore } from '@/stores/project'
import { useEditorStore } from '@/stores/editor'

const fileStore = useFileStore()
const projectStore = useProjectStore()
const editorStore = useEditorStore()
const message = useMessage()

// 图标
const RefreshIcon = () => h('svg', { viewBox: '0 0 24 24', fill: 'currentColor' }, [
  h('path', { d: 'M17.65 6.35C16.2 4.9 14.21 4 12 4c-4.42 0-7.99 3.58-7.99 8s3.57 8 7.99 8c3.73 0 6.84-2.55 7.73-6h-2.08c-.82 2.33-3.04 4-5.65 4-3.31 0-6-2.69-6-6s2.69-6 6-6c1.66 0 3.14.69 4.22 1.78L13 11h7V4l-2.35 2.35z' })
])
const FolderOpenIcon = () => h('svg', { viewBox: '0 0 24 24', fill: 'currentColor' }, [
  h('path', { d: 'M20 6h-8l-2-2H4c-1.1 0-1.99.9-1.99 2L2 18c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V8c0-1.1-.9-2-2-2zm0 12H4V8h16v10z' })
])

const contextMenuVisible = ref(false)
const contextMenuX = ref(0)
const contextMenuY = ref(0)
const selectedFileNode = ref(null)

// 右键菜单选项
const fileContextMenuOptions = computed(() => [
  {
    label: '打开',
    key: 'open',
    disabled: selectedFileNode.value?.type === 'directory'
  },
  {
    label: '删除',
    key: 'delete'
  }
])

// 节点属性
const fileNodeProps = ({ option }) => {
  return {
    onContextmenu(e) {
      e.preventDefault()
      selectedFileNode.value = option
      contextMenuX.value = e.clientX
      contextMenuY.value = e.clientY
      contextMenuVisible.value = true
    },
    onDblclick() {
      if (option.type === 'file') {
        openFile(option)
      }
    }
  }
}

// 渲染文件图标
const renderFileIcon = ({ option }) => {
  if (option.type === 'directory') {
    return h(NIcon, null, { default: () => h(FolderOutlined) })
  }

  // 根据文件扩展名返回不同图标
  const ext = option.ext?.toLowerCase()
  if (ext === '.api' || ext === '.proto') {
    return h(NIcon, { color: '#18a058' }, { default: () => h(CodeOutlined) })
  } else if (ext === '.sql') {
    return h(NIcon, { color: '#f0a020' }, { default: () => h(StorageOutlined) })
  } else if (ext === '.yaml' || ext === '.yml') {
    return h(NIcon, { color: '#2080f0' }, { default: () => h(DescriptionOutlined) })
  }

  return h(NIcon, null, { default: () => h(InsertDriveFileOutlined) })
}

// 选择文件
const handleSelectFile = (keys) => {
  if (keys.length > 0) {
    fileStore.selectFile(keys[0])
  }
}

// 打开文件
const openFile = async (fileNode) => {
  if (fileNode.type === 'directory') return

  try {
    const fullPath = fileNode.path
    const content = await fileStore.openFile(fullPath)

    // 在编辑器中打开
    editorStore.openTab({
      id: fullPath,
      title: fileNode.name,
      path: fullPath,
      content: content,
      modified: false,
      ext: fileNode.ext
    })
  } catch (e) {
    message.error('打开文件失败: ' + e.message)
  }
}

// 刷新
const handleRefresh = () => {
  if (projectStore.currentProject) {
    fileStore.loadFileTree(projectStore.currentProject.path)
  }
}

// 打开文件夹
const handleOpenFolder = () => {
  if (projectStore.currentProject) {
    message.info('功能开发中...')
  }
}

// 右键菜单选择
const handleFileContextMenu = (key) => {
  const file = selectedFileNode.value
  if (!file) return

  switch (key) {
    case 'open':
      openFile(file)
      break
    case 'delete':
      message.info('删除功能开发中...')
      break
  }

  contextMenuVisible.value = false
}

// 监听当前项目变化
watch(() => projectStore.currentProject, (newProject) => {
  if (newProject) {
    fileStore.loadFileTree(newProject.path)
  }
})
</script>

<style scoped>
.file-explorer {
  display: flex;
  flex-direction: column;
  height: 100%;
  background: #16213e;
}

.explorer-header {
  padding: 12px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  background: #1a1a2e;
}

.explorer-info {
  padding: 8px 12px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  background: #16213e;
}

.explorer-content {
  flex: 1;
  overflow: auto;
  padding: 8px;
  background: #16213e;
}
</style>
