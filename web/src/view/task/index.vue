<script setup>
// import { createCase } from '@/api/case'
import { deleteCase, findCaseByCaseID } from '@/api/case'
import { onMounted, ref } from 'vue'
import TaskInputSheet from './components/taskInputSheet.vue'
import { queryOngoingTask } from '@/api/task'
import { findAlgorithmById } from '@/api/algorithm'
import { ElMessage } from 'element-plus'
import { map } from 'core-js/internals/array-iteration'
import { findNickNameByUuid } from '@/api/user'
import axios from 'axios'

const taskData = ref([])
const taskDialogVisible = ref(false)
const taskTitle = ref('新增任务')
const taskRef = ref(null)

const showInputDialog = () => {
  taskDialogVisible.value = true
}

const closeDialog = () => {
  taskRef.value.resetForm()
  taskDialogVisible.value = false
}

const submitDialog = () => {
  taskRef.value.submitForm()
  taskDialogVisible.value = false
}

const getResolution = (num) => {
  switch (num) {
    case 1:
      return '360P'
    case 2:
      return '480P'
    case 3:
      return '540P'
    case 4:
      return '720P'
    case 5:
      return '1080P'
    default:
      return 'Unknown'
  }
}

const getIntensity = (num) => {
  switch (num) {
    case 1:
      return '粗'
    case 2:
      return '中'
    case 3:
      return '细'
    default:
      return 'Unknown'
  }
}

// TODO: fix bug that backend gets EOF error
const getAlgorithmNameById = async(id) => {
  const result = await findAlgorithmById(id)
  console.log('res:', result)
  console.log(result.data.algorithmName)
}

const getNickNameByUuid = async(uuid) => {
  const res = await findNickNameByUuid({ uuid: uuid })
  console.log(res.data.NickName)
  return res.data.NickName
}

const processTaskList = async(result) => {
  taskData.value = await Promise.all(result.data.map(async(item) => {
    const nickName = await getNickNameByUuid(item.uuid)
    const resolution = getResolution(item.resolution)
    const intensity = getIntensity(item.intensity)
    const algorithmName = await getAlgorithmNameById(item.algorithmID)

    return {
      taskID: item.ID,
      uuid: item.uuid,
      nickName,
      videoSource: item.videoSource,
      CreatedAt: item.CreatedAt.substring(0, 10) + ' ' + item.CreatedAt.substring(11, 19),
      resolution,
      algorithmID: item.ID,
      intensity,
      algorithmName,
    }
  }))
}

const getTaskList = async() => {
  const result = await queryOngoingTask()
  console.log('result:', result.data)
  // taskData.value = result.data.map(item => ({
  //   taskID: item.ID,
  //   uuid: item.uuid,
  //   nickName: getNickNameByUuid(item.uuid),
  //   videoSource: item.videoSource,
  //   CreatedAt: item.CreatedAt.substring(0, 10) + ' ' + item.CreatedAt.substring(11, 19),
  //   resolution: getResolution(item.resolution),
  //   // algorithmName: getAlgorithmNameById(item.algorithmId)
  //   algorithmID: item.ID,
  //   intensity: getIntensity(item.intensity),
  // }))
  await processTaskList(result)
  console.log('taskData:', taskData.value)
}

onMounted(() => {
  getTaskList()
})
</script>

<template>
  <div class="case">
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="showInputDialog">新增任务</el-button>
      </div>
      <el-table
        :data="taskData"
        row-key="taskID"
        style="width: 100%"
      >
        <el-table-column label="任务ID" min-width="180" prop="taskID" />
        <el-table-column align="left" label="发起人" min-width="180" prop="nickName" />
        <el-table-column align="left" label="视频源" min-width="180" prop="videoSource" />
        <el-table-column align="left" label="创建时间" min-width="180" prop="CreatedAt" />
        <el-table-column align="left" label="分辨率" min-width="180" prop="resolution" />
<!--        <el-table-column align="left" label="任务容器算法ID" min-width="180" prop="algorithmID" />-->
        <el-table-column align="left" label="任务容器算法名称" min-width="180" prop="algorithmName" />
        <el-table-column align="left" label="任务粒度" min-width="180" prop="intensity" />
        <el-table-column align="left" lebal="任务状态" min-width="180" prop="status" />
        <el-table-column align="left" label="操作" width="460">
          <template #default="scope">
            <el-button
              icon="edit"
              type="primary"
              link
            >编辑
            </el-button>
            <el-button
              icon="delete"
              type="primary"
              link
              @click="deleteCase(scope.row)"
            >删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-dialog v-model="taskDialogVisible" :title="taskTitle">
      <task-input-sheet ref="taskRef" />
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="submitDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<style scoped lang="scss">

</style>