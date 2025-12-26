import { defineStore } from 'pinia'
import { ref } from 'vue'
import { FileTree, FileRead, FileWrite } from '@/api/file'

export const useFileStore = defineStore('file', () => {
  // 状态
  const fileTree = ref([])
  const currentPath = ref('')
  const selectedFile = ref(null)
  const loading = ref(false)
  const error = ref(null)

  // Actions
  const setFileTree = (tree) => {
    fileTree.value = tree
  }

  const selectFile = (filePath) => {
    selectedFile.value = filePath
  }

  const loadFileTree = async (projectPath) => {
    loading.value = true
    error.value = null
    try {
      const tree = await FileTree(projectPath, 'desc')
      fileTree.value = tree
      return tree
    } catch (e) {
      error.value = e.message
      throw e
    } finally {
      loading.value = false
    }
  }

  const openFile = async (filePath) => {
    loading.value = true
    error.value = null
    try {
      const content = await FileRead(filePath)
      return content
    } catch (e) {
      error.value = e.message
      throw e
    } finally {
      loading.value = false
    }
  }

  const saveFile = async (filePath, content) => {
    loading.value = true
    error.value = null
    try {
      await FileWrite(filePath, content)
      return true
    } catch (e) {
      error.value = e.message
      throw e
    } finally {
      loading.value = false
    }
  }

  const navigateTo = (path) => {
    currentPath.value = path
  }

  return {
    fileTree,
    currentPath,
    selectedFile,
    loading,
    error,
    setFileTree,
    selectFile,
    loadFileTree,
    openFile,
    saveFile,
    navigateTo
  }
})
