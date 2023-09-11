<script setup>
import { reactive } from 'vue'
import { createAlgorithm } from '@/api/algorithm'
import { ElMessage } from 'element-plus'

const algorithm = reactive({
  algorithmName: '',
  algorithmVersion: 1.3,
  description: '',
  size: 0,
  downloadLink: '',
  MD5: ''
})

const createNewAlgorithm = async() => {
  const res = await createAlgorithm(algorithm)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '添加成功!'
    })
  }
}
</script>

<template>
  <div class="test-container">
    <h3>算法api测试</h3>
    <el-form v-model="algorithm">
      <el-form-item label="算法名称">
        <el-input v-model="algorithm.algorithmName" />
      </el-form-item>
      <el-form-item label="算法版本">
        <el-input v-model="algorithm.algorithmVersion" />
      </el-form-item>
      <el-form-item label="算法描述">
        <el-input v-model="algorithm.description" autosize type="textarea" />
      </el-form-item>
      <el-form-item label="下载链接">
        <el-input v-model="algorithm.downloadLink" />
      </el-form-item>
      <el-form-item label="MD5">
        <el-input v-model="algorithm.MD5" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="createNewAlgorithm">创建新算法</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<!--<style scoped lang="scss">-->

<!--</style>-->
