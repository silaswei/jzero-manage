<template>
  <div class="code-editor">
    <!-- 标签页 -->
    <EditorTabs />

    <!-- Monaco编辑器 -->
    <div class="editor-container">
      <div v-if="editorStore.currentTab" ref="editorContainer" class="monaco-container"></div>
      <n-empty v-else description="请选择一个文件打开" style="height: 100%; display: flex; align-items: center; justify-content: center;" />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, watch, nextTick } from 'vue'
import { useEditorStore } from '@/stores/editor'
import { useFileStore } from '@/stores/file'
import { FileWrite } from '@/api/file'
import { useMessage } from 'naive-ui'

const editorStore = useEditorStore()
const fileStore = useFileStore()
const message = useMessage()

const editorContainer = ref(null)
let editorInstance = null
let monacoInstance = null

onMounted(async () => {
  await initMonaco()
})

onBeforeUnmount(() => {
  if (editorInstance) {
    editorInstance.dispose()
  }
})

// 初始化Monaco编辑器
const initMonaco = async () => {
  try {
    const monaco = await import('monaco-editor')
    monacoInstance = monaco

    // 监听自定义事件，用于保存文件
    window.addEventListener('save-current-file', handleSaveFile)
  } catch (e) {
    console.error('Failed to load Monaco Editor:', e)
  }
}

// 创建编辑器实例
const createEditor = () => {
  if (!editorContainer.value || !monacoInstance) return

  if (editorInstance) {
    editorInstance.dispose()
  }

  const currentTab = editorStore.currentTab
  const language = getLanguage(currentTab.path)

  editorInstance = monacoInstance.editor.create(editorContainer.value, {
    value: currentTab.content || '',
    language: language,
    theme: 'vs-dark',
    automaticLayout: true,
    fontSize: 14,
    tabSize: 4,
    minimap: { enabled: true },
    scrollBeyondLastLine: false,
    wordWrap: 'on',
    formatOnPaste: true,
    formatOnType: true,
  })

  // 监听内容变化
  editorInstance.onDidChangeModelContent(() => {
    const tab = editorStore.currentTab
    if (tab) {
      const value = editorInstance.getValue()
      tab.content = value
      tab.modified = value !== tab.originalContent
      editorStore.updateTab(tab)
    }
  })

  // 添加快捷键 Cmd+S / Ctrl+S
  editorInstance.addCommand(monacoInstance.KeyMod.CtrlCmd | monacoInstance.KeyCode.KeyS, () => {
    handleSaveFile()
  })
}

// 保存文件
const handleSaveFile = async () => {
  const tab = editorStore.currentTab
  if (!tab || !tab.modified) return

  try {
    await FileWrite(tab.path, tab.content)
    tab.originalContent = tab.content
    tab.modified = false
    editorStore.updateTab(tab)
    message.success('文件保存成功')
  } catch (e) {
    message.error('文件保存失败: ' + e.message)
  }
}

// 获取文件对应的语言
const getLanguage = (filePath) => {
  const ext = filePath.split('.').pop()
  const languageMap = {
    'go': 'go',
    'api': 'go',
    'proto': 'protobuf',
    'sql': 'sql',
    'yaml': 'yaml',
    'yml': 'yaml',
    'json': 'json',
    'js': 'javascript',
    'ts': 'typescript',
    'vue': 'html',
    'md': 'markdown'
  }
  return languageMap[ext] || 'plaintext'
}

// 监听当前标签页变化
watch(() => editorStore.currentTab, (newTab) => {
  if (newTab) {
    nextTick(() => {
      if (!editorInstance) {
        createEditor()
      } else {
        // 更新语言
        const language = getLanguage(newTab.path)
        const model = editorInstance.getModel()
        if (model) {
          monacoInstance.editor.setModelLanguage(model, language)
          editorInstance.setValue(newTab.content || '')
        }

        // 保存原始内容用于比较修改状态
        if (!newTab.hasOwnProperty('originalContent')) {
          newTab.originalContent = newTab.content
        }
      }
    })
  }
}, { deep: true })
</script>

<style scoped>
.code-editor {
  display: flex;
  flex-direction: column;
  height: 100%;
  background: #1e1e2e;
}

.editor-container {
  flex: 1;
  overflow: hidden;
  background: #1e1e2e;
}

.monaco-container {
  width: 100%;
  height: 100%;
}
</style>
