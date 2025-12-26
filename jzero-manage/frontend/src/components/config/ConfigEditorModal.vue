<template>
  <n-modal
    v-model:show="visible"
    :mask-closable="false"
    preset="card"
    title="项目配置编辑"
    size="huge"
    style="width: 800px;"
  >
    <n-spin :show="loading">
      <n-form
        ref="formRef"
        :model="formData"
        label-placement="left"
        label-width="140px"
      >
        <!-- Hooks配置 -->
        <n-divider>Hooks配置</n-divider>
        <n-form-item label="Before Hooks">
          <n-dynamic-tags v-model:value="formData.hooks.before" />
        </n-form-item>
        <n-form-item label="After Hooks">
          <n-dynamic-tags v-model:value="formData.hooks.after" />
        </n-form-item>

        <!-- New配置 -->
        <n-divider>New配置</n-divider>
        <n-form-item label="Template Home">
          <n-input v-model:value="formData.new.home" placeholder=".template/jzero-admin-plugin" />
        </n-form-item>

        <!-- Gen配置 -->
        <n-divider>Gen配置</n-divider>
        <n-form-item label="Style">
          <n-select
            v-model:value="formData.gen.style"
            :options="styleOptions"
            placeholder="选择代码风格"
          />
        </n-form-item>
        <n-form-item label="Git Change">
          <n-switch v-model:value="formData.gen.gitChange" />
        </n-form-item>
        <n-form-item label="Route to Code">
          <n-switch v-model:value="formData.gen.route2Code" />
        </n-form-item>
        <n-form-item label="Model Driver">
          <n-select
            v-model:value="formData.gen.modelDriver"
            :options="driverOptions"
            placeholder="选择数据库驱动"
          />
        </n-form-item>
        <n-form-item label="Model Cache">
          <n-switch v-model:value="formData.gen.modelCache" />
        </n-form-item>
        <n-form-item label="Model Schema">
          <n-input v-model:value="formData.gen.modelSchema" placeholder="数据库名称" />
        </n-form-item>
        <n-form-item label="Ignore Columns">
          <n-dynamic-tags v-model:value="formData.gen.modelIgnoreCols" />
        </n-form-item>
      </n-form>
    </n-spin>

    <template #footer>
      <n-space justify="end">
        <n-button @click="visible = false">取消</n-button>
        <n-button type="primary" @click="handleSave" :loading="saving">
          保存
        </n-button>
      </n-space>
    </template>
  </n-modal>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useMessage, useDialog } from 'naive-ui'
import { ConfigRead, ConfigWrite } from '@/api/config'

const message = useMessage()
const dialog = useDialog()

const visible = ref(false)
const loading = ref(false)
const saving = ref(false)
const formRef = ref(null)
const currentProjectPath = ref('')

const formData = reactive({
  hooks: {
    before: [],
    after: []
  },
  new: {
    home: ''
  },
  gen: {
    style: 'go_zero',
    gitChange: true,
    route2Code: true,
    modelDriver: 'mysql',
    modelCache: false,
    modelSchema: '',
    modelIgnoreCols: []
  }
})

const styleOptions = [
  { label: 'Go Zero', value: 'go_zero' },
  { label: 'Go Zero Plus', value: 'go_zero_plus' }
]

const driverOptions = [
  { label: 'MySQL', value: 'mysql' },
  { label: 'PostgreSQL', value: 'postgres' },
  { label: 'SQLite', value: 'sqlite' }
]

// 保存配置
const handleSave = async () => {
  try {
    saving.value = true
    await ConfigWrite(currentProjectPath.value, formData)
    message.success('配置保存成功')
    visible.value = false
  } catch (error) {
    message.error('配置保存失败: ' + error.message)
  } finally {
    saving.value = false
  }
}

// 打开编辑器
const open = async (projectPath) => {
  currentProjectPath.value = projectPath
  visible.value = true
  loading.value = true
  try {
    const config = await ConfigRead(projectPath)

    // 更新表单数据
    if (config.Hooks) {
      formData.hooks.before = config.Hooks.Before || []
      formData.hooks.after = config.Hooks.After || []
    }
    if (config.New) {
      formData.new.home = config.New.Home || ''
    }
    if (config.Gen) {
      formData.gen.style = config.Gen.Style || 'go_zero'
      formData.gen.gitChange = config.Gen.GitChange !== undefined ? config.Gen.GitChange : true
      formData.gen.route2Code = config.Gen.Route2Code !== undefined ? config.Gen.Route2Code : true
      formData.gen.modelDriver = config.Gen.ModelDriver || 'mysql'
      formData.gen.modelCache = config.Gen.ModelCache || false
      formData.gen.modelSchema = config.Gen.ModelSchema || ''
      formData.gen.modelIgnoreCols = config.Gen.ModelIgnoreCols || []
    }
  } catch (error) {
    message.error('加载配置失败: ' + error.message)
    visible.value = false
  } finally {
    loading.value = false
  }
}

// 监听自定义事件
onMounted(() => {
  window.addEventListener('open-config-editor', (e) => {
    open(e.detail.projectPath)
  })
})

defineExpose({
  open
})
</script>
