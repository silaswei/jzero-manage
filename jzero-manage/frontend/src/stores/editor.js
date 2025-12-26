import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useEditorStore = defineStore('editor', () => {
  // 状态
  const tabs = ref([])
  const currentTabId = ref(null)

  // 计算属性
  const currentTab = computed(() => {
    return tabs.value.find(tab => tab.id === currentTabId.value)
  })

  // Actions
  const openTab = (tab) => {
    // 检查是否已打开
    const existingTab = tabs.value.find(t => t.id === tab.id)
    if (existingTab) {
      currentTabId.value = tab.id
      return
    }

    // 添加新标签页
    tabs.value.push(tab)
    currentTabId.value = tab.id
  }

  const closeTab = (tabId) => {
    const index = tabs.value.findIndex(tab => tab.id === tabId)
    if (index === -1) return

    // 如果关闭的是当前标签页
    if (tabId === currentTabId.value) {
      // 切换到相邻标签页
      if (tabs.value.length > 1) {
        currentTabId.value = tabs.value[index + 1]?.id || tabs.value[index - 1]?.id
      } else {
        currentTabId.value = null
      }
    }

    tabs.value.splice(index, 1)
  }

  const switchTab = (tabId) => {
    currentTabId.value = tabId
  }

  const updateTab = (tab) => {
    const index = tabs.value.findIndex(t => t.id === tab.id)
    if (index !== -1) {
      tabs.value[index] = tab
    }
  }

  const closeOthers = () => {
    if (!currentTabId.value) return
    tabs.value = tabs.value.filter(tab => tab.id === currentTabId.value)
  }

  const closeAll = () => {
    tabs.value = []
    currentTabId.value = null
  }

  const saveAll = async () => {
    for (const tab of tabs.value) {
      if (tab.modified) {
        // 保存文件逻辑
        tab.modified = false
      }
    }
  }

  return {
    tabs,
    currentTabId,
    currentTab,
    openTab,
    closeTab,
    switchTab,
    updateTab,
    closeOthers,
    closeAll,
    saveAll
  }
})
