<script setup>
import {getCurrentInstance, onBeforeMount, onMounted, ref, watch} from 'vue'
import {
  deleteTask,
  queryOngoingTask,
  killProcess,
  awakeProcess,
  pauseProcess,
  startProcess,
  createTask, updateTask
} from '@/api/task'
import {findAlgorithmById, queryAllAlgorithm} from '@/api/algorithm'
import { ElMessage } from 'element-plus'
import {findNickNameByUuid, getUserInfo} from '@/api/user'

const taskData = ref([])
const taskDialogVisible = ref(false)
const taskTitle = ref('新增任务')

const algoNumTextType = ref('default')
const taskNumber = ref(0)
const addDisabled = ref(false)

const dialogKey = ref('add')
const formData = ref({})

const statusDialogVisible = ref(false)
const currentTask = ref({})
const isLoading = ref(true)

const rules = {
  videoSource: [{
    required: true,
    message: '字段值不可为空',
  }],

}
const resolutionOptions = [{
  'label': '360P',
  'value': 1,
}, {
  'label': '480P',
  'value': 2,
}, {
  'label': '540P',
  'value': 3,
}, {
  'value': 4,
  'label': '720P',
}, {
  'value': 5,
  'label': '1080P',
}]
const algorithmOptions = ref([])
const intensityOptions = [{
  'label': '粗',
  'value': 1,
}, {
  'label': '中',
  'value': 2,
}, {
  'label': '细',
  'value': 3,
}]


const getLocalAlgorithm = async() => {
  await queryAllAlgorithm().then((result) => {
    console.log('result:', result.data)
    algorithmOptions.value = result.data.map(item => {
      return {
        label: item.algorithmName,
        value: item.ID,
      }
    })
  })
}

const setCurrentUuid = async() => {
  await getUserInfo().then((result) => {
    formData.value.uuid = result.data.userInfo.uuid
  })
}

const showInputDialog = () => {
  if (dialogKey.value === 'add') {
    setCurrentUuid()
    console.log(formData.value.uuid)
  }
  getLocalAlgorithm()
  taskDialogVisible.value = true
}

const instance = getCurrentInstance()

const closeDialog = () => {
  resetTable()
  taskDialogVisible.value = false
}

const submitDialog = () => {
  instance.ctx.$refs['vForm'].validate(async(valid) => {
    if (!valid)
      return
    if (dialogKey.value === 'add') {
      await createTask({
        uuid: formData.value.uuid,
        videoSource: formData.value.videoSource,
        resolution: formData.value.resolution,
        intensity: formData.value.intensity,
        algorithmId: formData.value.algorithmId,
      }).then(res => {
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '添加成功！'
          })
          location.reload()
        }
      }).catch(err => {
        console.error('Error:', err)
      })
    }
    if (dialogKey.value === 'edit') {
      await updateTask({
        id: formData.value.id,
        uuid: formData.value.uuid,
        videoSource: formData.value.videoSource,
        resolution: formData.value.resolution,
        intensity: formData.value.intensity,
        algorithmId: formData.value.algorithmId,
      }).then(res => {
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '更新成功！',
          })
          location.reload()
        }
      }).catch(err => {
        ElMessage({
          type: 'error',
          message: err,
        })
      })
    }
    taskDialogVisible.value = false
  })
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

const getStatus = (num) => {
  switch (num) {
    case 0:
      return '未在运行'
    case 1:
      return '正在运行'
  }
}

// TODO: fix bug that backend gets EOF error
const getAlgorithmNameById = async(id) => {
  const res = await findAlgorithmById({ ID: id })
  return res.data.algorithmName
}

const getNickNameByUuid = async(uuid) => {
  const res = await findNickNameByUuid({ uuid: uuid })
  // console.log(res.data.NickName)
  return res.data.NickName
}

const processTaskList = async(result) => {
  taskData.value = await Promise.all(result.data.map(async(item) => {
    const nickName = await getNickNameByUuid(item.uuid)
    const resolutionText = getResolution(item.resolution)
    const intensityText = getIntensity(item.intensity)
    const algorithmName = await getAlgorithmNameById(item.algorithmId)
    const statusText = getStatus(item.status)

    return {
      id: item.ID,
      uuid: item.uuid,
      nickName,
      videoSource: item.videoSource,
      CreatedAt: item.CreatedAt.substring(0, 10) + ' ' + item.CreatedAt.substring(11, 19),
      resolutionText,
      algorithmId: item.algorithmId,
      intensityText,
      algorithmName,
      statusText,
      intensity: item.intensity,
      status: item.status,
      resolution: item.resolution,
    }
  }))
}

const getTaskList = async() => {
  const result = await queryOngoingTask()
  console.log('result:', result.data)
  // if (result.data.length === 0)
  //   return
  await processTaskList(result)
  console.log('taskData:', taskData.value)
  taskNumber.value = taskData.value.length

}

const deleteTaskInForm = async(taskID) => {
  await deleteTask({ id: taskID }).then(res => {
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功！'
      })
      location.reload()
    }
  }).catch(err => {
    ElMessage({
      type: 'error',
      message: err
    })
  })
}

const editTask = (scope) => {
  dialogKey.value = 'edit'
  taskTitle.value = '编辑任务'
  formData.value = scope
  showInputDialog()
}

const addTask = () => {
  dialogKey.value = 'add'
  taskTitle.value = '新增任务'
  showInputDialog()
}

const showStatusDialog = (task) => {
  currentTask.value = task
  statusDialogVisible.value = true
  console.log('cT:', currentTask.value)
}

const pauseTask = async() => {
  isLoading.value = true
  await pauseProcess({
    ID: currentTask.value.taskID
  })
  // isLoading.value = false
}

const awakeTask = async() => {
  await awakeProcess({
    ID: currentTask.value.taskID
  })
}

const startTask = async() => {
  await startProcess({
    ID: currentTask.value.taskID
  })
}

const killTask = async() => {
  await killProcess({
    ID: currentTask.value.taskID
  })
}

const resetTable = () => {
  formData.value.algorithmId = ''
  formData.value.videoSource = ''
  formData.value.intensity = ''
  formData.value.resolution = ''
}

onMounted(() => {
  getTaskList()
})

</script>

<template>
  <div class="case">
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="addTask">新增任务</el-button>
        <el-text size="large" :type="algoNumTextType">&nbsp; &nbsp; 当前已使用的算法数量：{{ taskNumber }}</el-text>
      </div>
      <el-table
        :data="taskData"
        row-key="taskID"
        style="width: 100%"
      >
        <el-table-column label="任务ID" min-width="120" prop="id" />
        <el-table-column align="left" label="发起人" min-width="150" prop="nickName" />
        <el-table-column align="left" label="视频源" min-width="180" prop="videoSource" />
        <el-table-column align="left" label="创建时间" min-width="180" prop="CreatedAt" />
        <el-table-column align="left" label="分辨率" min-width="150" prop="resolutionText" />
        <el-table-column align="left" label="任务容器算法名称" min-width="180" prop="algorithmName" />
        <el-table-column align="left" label="任务粒度" min-width="120" prop="intensityText" />
        <el-table-column align="left" label="任务状态" min-width="150" prop="statusText" />
        <el-table-column align="left" label="操作" width="280">
          <template #default="scope">
            <el-button
              icon="edit"
              type="primary"
              link
              @click="editTask(scope.row)"
            >编辑信息
            </el-button>
            <el-button
                link
                icon="open"
                type="primary"
                @click="showStatusDialog(scope.row)"
            >状态管理
            </el-button>
            <el-button
              icon="delete"
              type="danger"
              link
              @click="deleteTaskInForm(scope.row.id)"
            >删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-dialog v-loading="isLoading" v-model="statusDialogVisible" title="任务状态管理">
      <el-text size="large">当前状态：{{ currentTask.status === 1 ? "正在运行" : "未在运行" }}</el-text>
      <br>
      <template #footer>
        <el-button
            v-if="currentTask.status === 1 "
            icon="videoPause"
            type="warning"
            @click="pauseTask"
        >暂停
        </el-button>
        <el-button
            v-if="currentTask.status === 1 "
            icon="close"
            type="danger"
            @click="killTask"
        >终止
        </el-button>
        <el-button
            v-if="currentTask.status === 0 "
            icon="refreshLeft"
            type="primary"
            @click="awakeTask"
        >唤醒
        </el-button>
        <el-button
            v-if="currentTask.status === 0 "
            icon="refreshLeft"
            type="success"
            @click="startTask"
        >启动
        </el-button>
      </template>
    </el-dialog>
    <el-dialog v-model="taskDialogVisible" :title="taskTitle">

      <el-form :model="formData" ref="vForm" :rules="rules" label-position="left" label-width="150px"
               size="medium" @submit.prevent
      >
        <el-form-item label="发起人UUID" prop="uuid" class="required label-right-align">
          <el-input v-model="formData.uuid" type="text" clearable :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="视频源" prop="videoSource" class="required label-right-align">
          <el-input v-model="formData.videoSource" type="text" clearable></el-input>
        </el-form-item>
        <el-form-item label="分辨率" prop="resolution" class="label-right-align">
          <el-select v-model="formData.resolution" class="full-width-input" clearable>
            <el-option v-for="(item, index) in resolutionOptions" :key="index" :label="item.label"
                       :value="item.value" :disabled="item.disabled"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="任务容器算法" prop="algorithmId" class="label-right-align">
          <el-select v-model="formData.algorithmId" class="full-width-input" clearable>
            <el-option v-for="(item, index) in algorithmOptions" :key="index" :label="item.label"
                       :value="item.value" :disabled="item.disabled"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="粒度" prop="intensity" class="label-right-align">
          <el-radio-group v-model="formData.intensity">
            <el-radio v-for="(item, index) in intensityOptions" :key="index" :label="item.value"
                      :disabled="item.disabled" style="{display: inline}"
            >{{ item.label }}
            </el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>

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