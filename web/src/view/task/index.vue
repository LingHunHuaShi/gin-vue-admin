<script setup>
// import { createCase } from '@/api/case'
import { deleteCase, findCaseByCaseID } from '@/api/case'
import { onMounted, ref } from 'vue'
import TaskInputSheet from './components/taskInputSheet.vue'
import { queryOngoingTask } from '@/api/task'
import { findAlgorithmById } from '@/api/algorithm'
import { ElMessage } from 'element-plus'
import { map } from 'core-js/internals/array-iteration'

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

const getAlgorithmNameById = async(id) => {
  // console.log({ algorithmID: id })
  const result = await findAlgorithmById({ id: id })
  console.log(result)
}

const getTaskList = async() => {
  const result = await queryOngoingTask()
  // console.log('result:', result.data)
  taskData.value = result.data.map(item => ({
    taskID: item.taskID,
    uuid: item.uuid,
    source: item.videoSource,
    CreatedAt: item.CreatedAt.substring(0, 10) + ' ' + item.CreatedAt.substring(11, 19),
    resolution: getResolution(item.resolution),
    // algorithmName: getAlgorithmNameById(item.algorithmId)
    algorithmID: item.algorithmId,
    intensity: getIntensity(item.intensity),
  }))
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
        row-key="caseID"
        style="width: 100%"
      >
        <el-table-column label="任务ID" min-width="180" prop="taskID" />
        <el-table-column align="left" label="发起人UUID" min-width="180" prop="uuid" />
        <el-table-column align="left" label="视频源" min-width="180" prop="videoSource" />
        <el-table-column align="left" label="创建时间" min-width="180" prop="CreatedAt" />
        <el-table-column align="left" label="分辨率" min-width="180" prop="resolution" />
        <el-table-column align="left" label="任务容器算法ID" min-width="180" prop="algorithmID" />
<!--        <el-table-column align="left" label="任务容器算法名称" min-width="180" prop="algorithmName" />-->
        <el-table-column align="left" label="任务粒度" min-width="180" prop="intensity" />
        <el-table-column align="left" lebal="任务状态" min-width="180" prop="status" />
        <el-table-column align="left" label="操作" width="460">
          <template #default="scope">
            <el-button
              icon="view"
              type="primary"
              link
            >查看详情
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