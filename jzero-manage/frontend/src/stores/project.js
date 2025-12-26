import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { ProjectScan, ProjectInfo } from '@/api/project'

export const useProjectStore = defineStore('project', () => {
  // 状态
  const projects = ref([])
  const currentProject = ref(null)
  const loading = ref(false)
  const error = ref(null)

  // 计算属性
  const projectTree = computed(() => {
    return projects.value
  })

  // Actions
  const setProjects = (projectList) => {
    projects.value = projectList
  }

  const selectProject = (projectId) => {
    const project = findProjectById(projects.value, projectId)
    if (project) {
      currentProject.value = project
    }
  }

  const scanProjects = async (rootPath) => {
    loading.value = true
    error.value = null
    try {
      const tree = await ProjectScan(rootPath)
      projects.value = tree
      return tree
    } catch (e) {
      error.value = e.message
      throw e
    } finally {
      loading.value = false
    }
  }

  const refreshProject = async () => {
    if (currentProject.value) {
      loading.value = true
      try {
        const info = await ProjectInfo(currentProject.value.path)
        currentProject.value.info = info
      } catch (e) {
        error.value = e.message
      } finally {
        loading.value = false
      }
    }
  }

  // 递归查找项目
  const findProjectById = (nodes, id) => {
    for (const node of nodes) {
      if (node.id === id) {
        return node
      }
      if (node.children && node.children.length > 0) {
        const found = findProjectById(node.children, id)
        if (found) {
          return found
        }
      }
    }
    return null
  }

  return {
    projects,
    currentProject,
    loading,
    error,
    projectTree,
    setProjects,
    selectProject,
    scanProjects,
    refreshProject
  }
})
