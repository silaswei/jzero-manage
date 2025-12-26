<template>
  <div class="welcome-page">
    <div class="welcome-content">
      <div class="logo-section">
        <h1 class="title">JZero 项目管理工具</h1>
        <p class="subtitle">基于 Wails + Vue 3 + Naive UI 构建</p>
      </div>

      <n-space vertical size="large" style="margin-top: 40px;">
        <n-card hoverable class="action-card" @click="handleOpenProject">
          <template #header>
            <n-space align="center">
              <n-icon size="24" color="#18a058">
                <FolderOpenIcon />
              </n-icon>
              <span class="card-title">打开项目目录</span>
            </n-space>
          </template>
          <n-text depth="3">选择包含JZero项目的父目录，自动扫描所有项目</n-text>
        </n-card>

        <n-card hoverable class="action-card" @click="handleRecentProjects">
          <template #header>
            <n-space align="center">
              <n-icon size="24" color="#2080f0">
                <HistoryIcon />
              </n-icon>
              <span class="card-title">最近的项目</span>
            </n-space>
          </template>
          <n-text depth="3">查看最近打开的项目（功能开发中）</n-text>
        </n-card>

        <n-card hoverable class="action-card" @click="handleOpenFile">
          <template #header>
            <n-space align="center">
              <n-icon size="24" color="#f0a020">
                <FileIcon />
              </n-icon>
              <span class="card-title">打开单个文件</span>
            </n-space>
          </template>
          <n-text depth="3">直接打开API、Proto或SQL文件进行编辑（功能开发中）</n-text>
        </n-card>
      </n-space>

      <div class="tips-section">
        <n-divider />
        <n-text depth="3">
          💡 提示：JZero项目应包含 go.mod 和 desc 目录或 .jzero.yaml 配置文件
        </n-text>
      </div>
    </div>
  </div>
</template>

<script setup>
import { h } from 'vue'
import { NIcon, useMessage } from 'naive-ui'
import { FolderOpen, History, Description } from '@vicons/material'

const message = useMessage()

const FolderOpenIcon = () => h('svg', {
  viewBox: '0 0 24 24',
  fill: 'currentColor'
}, [
  h('path', { d: 'M20 6h-8l-2-2H4c-1.1 0-1.99.9-1.99 2L2 18c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V8c0-1.1-.9-2-2-2zm0 12H4V8h16v10z' })
])

const HistoryIcon = () => h('svg', {
  viewBox: '0 0 24 24',
  fill: 'currentColor'
}, [
  h('path', { d: 'M13 3c-4.97 0-9 4.03-9 9H1l3.89 3.89.07.14L9 12H6c0-3.87 3.13-7 7-7s7 3.13 7 7-3.13 7-7 7c-1.93 0-3.68-.79-4.94-2.06l-1.42 1.42C8.27 19.99 10.51 21 13 21c4.97 0 9-4.03 9-9s-4.03-9-9-9zm-1 5v5l4.28 2.54.72-1.21-3.5-2.08V8H12z' })
])

const FileIcon = () => h('svg', {
  viewBox: '0 0 24 24',
  fill: 'currentColor'
}, [
  h('path', { d: 'M14 2H6c-1.1 0-1.99.9-1.99 2L4 20c0 1.1.89 2 1.99 2H18c1.1 0 2-.9 2-2V8l-6-6zm2 16H8v-2h8v2zm0-4H8v-2h8v2zm-3-5V3.5L18.5 9H13z' })
])

// 打开项目目录
const handleOpenProject = () => {
  // 触发全局事件，由App.vue或MainLayout处理
  window.dispatchEvent(new CustomEvent('open-project-dialog'))
}

// 最近的项目
const handleRecentProjects = () => {
  message.info('最近项目功能开发中...')
}

// 打开单个文件
const handleOpenFile = () => {
  message.info('打开文件功能开发中...')
}
</script>

<style scoped>
.welcome-page {
  width: 100%;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #1a1a2e 0%, #16213e 100%);
}

.welcome-content {
  max-width: 600px;
  padding: 40px;
}

.logo-section {
  text-align: center;
  margin-bottom: 20px;
}

.title {
  font-size: 36px;
  font-weight: 600;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin: 0;
}

.subtitle {
  font-size: 16px;
  color: #a0a0b0;
  margin-top: 12px;
}

.action-card {
  cursor: pointer;
  transition: all 0.3s ease;
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.action-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 24px rgba(102, 126, 234, 0.2);
  border-color: rgba(102, 126, 234, 0.3);
}

.card-title {
  font-size: 16px;
  font-weight: 500;
}

.tips-section {
  margin-top: 40px;
  text-align: center;
}
</style>
