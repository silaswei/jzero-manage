<template>
  <div class="editor-tabs">
    <n-tabs
      v-model:value="activeTabId"
      type="card"
      closable
      tab-style="min-width: 120px; max-width: 200px;"
      @close="handleClose"
      @update:value="handleTabChange"
    >
      <n-tab
        v-for="tab in editorStore.tabs"
        :key="tab.id"
        :name="tab.id"
      >
        <template #prefix>
          <n-icon v-if="tab.modified" size="14" style="color: #f0a020;">
            <CircleIcon />
          </n-icon>
        </template>
        <span>{{ tab.title }}</span>
      </n-tab>
    </n-tabs>

    <!-- 标签页操作 -->
    <div class="tab-actions">
      <n-dropdown :options="tabMenuOptions" @select="handleTabMenuSelect">
        <n-button size="small" quaternary>
          <template #icon>
            <n-icon><MoreIcon /></n-icon>
          </template>
        </n-button>
      </n-dropdown>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, h } from 'vue'
import { useEditorStore } from '@/stores/editor'
import { useMessage } from 'naive-ui'
import CircleOutlined from '@vicons/material/es/CircleOutlined'

const editorStore = useEditorStore()
const message = useMessage()

const CircleIcon = CircleOutlined

const activeTabId = computed({
  get: () => editorStore.currentTab?.id || '',
  set: (value) => {
    editorStore.switchTab(value)
  }
})

const MoreIcon = () => h('svg', { viewBox: '0 0 24 24', fill: 'currentColor' }, [
  h('path', { d: 'M12 8c1.1 0 2-.9 2-2s-.9-2-2-2-2 .9-2 2 .9 2 2 2zm0 2c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2zm0 6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z' })
])

// 关闭标签页
const handleClose = (name) => {
  editorStore.closeTab(name)
}

// 标签页切换
const handleTabChange = (name) => {
  editorStore.switchTab(name)
}

// 标签页菜单
const tabMenuOptions = [
  { label: '关闭其他', key: 'close-others' },
  { label: '关闭所有', key: 'close-all' },
  { label: '保存所有', key: 'save-all' }
]

const handleTabMenuSelect = (key) => {
  switch (key) {
    case 'close-others':
      editorStore.closeOthers()
      break
    case 'close-all':
      editorStore.closeAll()
      break
    case 'save-all':
      editorStore.saveAll()
      message.success('所有文件已保存')
      break
  }
}
</script>

<style scoped>
.editor-tabs {
  display: flex;
  align-items: center;
  background: #1e1e2e;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.tab-actions {
  padding: 0 8px;
  background: #1e1e2e;
}
</style>
