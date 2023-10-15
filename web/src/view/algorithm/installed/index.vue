<script setup>

import { computed, onMounted, ref } from 'vue'
// import detail from '../detail.vue'
import { queryOngoingTask } from '@/api/task'
import { queryAllAlgorithm } from '@/api/algorithm'
import TaskInput from './components/taskInput.vue'

const installed = ref([])
const dialogVisible = ref(false)
const taskRef = ref(null)
const currentAlgoName = ref('defaultName')
const currentAlgoId = ref(-1)
const propsLoaded = ref(false)
const addDisabled = ref(false)
const maxTaskNumber = 3
const addBtnText = ref('新建任务')

const taskData = ref([])
const getInstalledAlgorithm = async() => {
  await queryAllAlgorithm().then((result) => {
    installed.value = result.data
    console.log(result.data)
  })
}

const submitDialog = () => {
  taskRef.value.submitForm()
}

const showTaskDialog = (algo) => {
  currentAlgoName.value = algo.algorithmName
  currentAlgoId.value = algo.ID
  propsLoaded.value = true
  dialogVisible.value = true
}

const cancelDialog = () => {
  taskRef.value.resetForm()
  currentAlgoId.value = null
  currentAlgoId.value = null
  dialogVisible.value = false
  propsLoaded.value = false
}


const getTaskNumber = async() => {
  await queryOngoingTask().then(res => {
    taskData.value = res.data
  })
  if (taskData.value.length >= maxTaskNumber) {
    addDisabled.value = true
    addBtnText.value = '任务数量达到上限'
  }
}

onMounted(() => {
  getInstalledAlgorithm()
  getTaskNumber()
})

// const usedNPUCore = computed(() => {
//   let totalCore = 0
//   for (let i = 0; i < installed.length; i++) {
//     totalCore += installed[i].stream_number
//   }
//   return totalCore
// })

// TODO: fix本地算法添加的任务无法查询

</script>

<template>
  <div class="algorithm-container">
    <el-dialog v-if="propsLoaded" v-model="dialogVisible" title="新建任务" >
<!--      <detail />-->
      <task-input ref="taskRef" :algo-name="currentAlgoName" :algo-id="currentAlgoId" />
      <div slot="footer" class="dialog-footer align-right">
        <el-button type="primary" @click="submitDialog">确 定</el-button>
        <el-button @click="cancelDialog">取消</el-button>
      </div>
    </el-dialog>
<!--    <el-row class="algorithm-row">-->
<!--      <el-col :span="8" class="algorithm-col">-->
<!--        <el-card class="conclusion-card card" :gutter="20" type="flex">-->
<!--          <div slot="header">-->
<!--            <span>核心使用情况</span>-->
<!--            <br>-->
<!--            <br>-->
<!--          </div>-->
<!--          <div class="content-box">-->
<!--            <span style="line-height: 35px">已使用:{{ usedNPUCore }} / 3</span>-->
<!--            <br>-->
<!--            <div v-for="algorithm in installed" :key="algorithm" class="algorithm-text">-->
<!--              {{ algorithm.name }} 使用核心数：{{ algorithm.stream_number }}-->
<!--            </div>-->
<!--          </div>-->
<!--        </el-card>-->
<!--      </el-col>-->
<!--    </el-row>-->

    <el-row class="algorithm-row">
      <span class="title">已部署算法</span>
      <el-divider />
      <el-col v-for="algorithm in installed" :key="algorithm" class="algorithm-col" :md="12" :sm="24" :lg="12" :xl="8">
        <el-card class="installed-card card">
          <div slot="header">
            <span class="title algorithm-title">{{ algorithm.algorithmName }}</span>
          </div>
          <div class="content-box algorithm-text">
            <span>算法版本：{{ algorithm.algorithmVersion }}</span>
            <br>
            <span>算法描述：{{ algorithm.description }}</span>
            <br>
            <span>更新时间：{{ algorithm.UpdatedAt.substring(0, 10) + " " + algorithm.UpdatedAt.substring(11, 19) }}</span>
            <br>
            <div class="align-right">
              <!--              <el-button type="primary" class="el-button" @click="goToDetail">管理</el-button>-->
              <el-button :disabled="addDisabled" type="primary" class="el-button" @click="showTaskDialog(algorithm)">{{ addBtnText }}</el-button>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<style scoped lang="scss">
.algorithm {
  &-text {
    line-height: 35px;
    font-size: 15px;
  }
  &-title {
    font-size: 16px;
  }
}

.align-right {
  text-align: right;
}

.card {
  margin-right: 20px;
  margin-bottom: 20px;
}

.title {
  font-size: 20px;
}

</style>
