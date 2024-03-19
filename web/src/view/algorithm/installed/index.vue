<script setup>

import { computed, onMounted, ref } from 'vue'
// import detail from '../detail.vue'
import { queryOngoingTask } from '@/api/task'
import { queryAllAlgorithm } from '@/api/algorithm'
import TaskInput from './components/taskInput.vue'
import axios from "axios";

const installed = ref([])
const dialogVisible = ref(false)
const taskRef = ref(null)
const currentAlgoName = ref('defaultName')
const currentAlgoId = ref(-1)
const propsLoaded = ref(false)
const addBtnText = ref('新建任务')

const serverAddress = ref('http://192.168.67.212:8888')

const terminalForm = ref({
  cpu_framework: 'X86',
  cpu_model: 'AMD Ryzen™ Embedded V2748 with Radeon™ Graphics',
  ram: 1024,
  rom: 8192,
  arithmetic: 'yolov5s',
  arithmetic_Version: '1.0',
})

const LatestAlgorithm = ref({
  algorithmId: null,
  algorithmName: 'yolov5s',
  algorithmVersion: 'v1.6',
  algorithmDownloadlink: null,
  algorithmSize: null,
  algorithmDescription: null,
  algorithmMd5: null,
  createTime: null,
  cpuFramwork: null,
  cpuModel: null,
  ram: null,
  rom: null,
})

const TerminAlgorithmUpdate = ref(true)
const ari_details = ref(false)
const title = ref('')

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
}

const loginCloud = async() => {
  console.log(serverAddress.value+'/BoxesUserLogin')
  const res = await axios.post(serverAddress.value+'/BoxesUserLogin', {
    username: 'zhangsan',
    password: '123456',
    code: null,
    uuid: null,
    terminalForm: terminalForm,
  }).catch(err => console.log(err))

  console.log('sssssss', TerminAlgorithmUpdate.value)

  if (res.data.LatestAlgorithm!=null){
    if (terminalForm.value.arithmetic_Version < res.data.LatestAlgorithm.algorithmVersion) {
      // this.TerminArithmUpdateTip(res.LatestAlgorithm)
      LatestAlgorithm.value = res.data.LatestAlgorithm
      LatestAlgorithm.value.cpuModel = res.data.LatestAlgorithm.cpuModel.split(',')
      TerminAlgorithmUpdate.value = true
    }
  }

}

const latestAlgorithmDetails = () => {
  ari_details.value = true
  title.value = "算法详情——" + LatestAlgorithm.value.algorithmName
}

const cancel = () => {
  dialogVisible.value = false
}


const atiDetailsComfrim = () => {
  ari_details.value = false
}

onMounted(() => {
  getInstalledAlgorithm()
  getTaskNumber()
  loginCloud()
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

    <!--    算法更新提示框-->
    <el-dialog title="新算法已发布" v-model="TerminAlgorithmUpdate" width="400px" append-to-body
               >
      <div style="text-align: center">
        <label
            v-text="'您的设备当前算法及版本为:  ' + terminalForm.arithmetic + '-' + terminalForm.arithmetic_Version"></label>
        <br/>
        <label
            v-text="'云端已发布该算法新版本:  ' + LatestAlgorithm.algorithmName +'-'+ LatestAlgorithm.algorithmVersion"></label>
        <el-button
            size="mini"
            type="text"
            icon="Search"
            @click="latestAlgorithmDetails"
            v-hasPermi="['']"
        >查看详情
        </el-button>
        <br/>
        <br/>
        <label v-text="'是否立刻更新算法？'"></label>
      </div>
      <div slot="footer" class="dialog-footer align-right" style="margin-top: 20px">
        <el-button @click="cancel">以后再说</el-button>

        <el-button @click="updateComfirm" type="primary">立刻更新</el-button>
      </div>
    </el-dialog>

    <!-- 查看算法详情对话框 -->
    <el-dialog :title="title" v-model="ari_details" width="550px" append-to-body>
      <el-form label-width="150px">
        <el-form-item label="算法名:" prop="algorithmName">
          <el-tag type="primary" v-text="LatestAlgorithm.algorithmName" style="margin-left: 10px"/>
        </el-form-item>
        <el-form-item label="算法版本:" prop="algorithmVersion">
          <el-tag type="primary" v-text="LatestAlgorithm.algorithmVersion" style="margin-left: 10px"/>
        </el-form-item>
        <el-form-item label="算法模型文件大小:" prop="algorithmSize">
          <el-tag type="primary" v-text="LatestAlgorithm.algorithmSize" style="margin-left: 10px"/>
        </el-form-item>
        <el-form-item label="发布时间:" prop="algorithmSize">
          <el-tag type="primary" v-text="LatestAlgorithm.createTime" style="margin-left: 10px"/>
        </el-form-item>
        <el-form-item label="所支持cpu框架:" prop="cpuFramwork">
          <el-tag type="primary" v-text="LatestAlgorithm.cpuFramwork" style="margin-left: 10px"/>
        </el-form-item>
        <el-form-item label="所支持cpu型号:" prop="cpuModel">
          <el-tag type="primary" v-for="item in LatestAlgorithm.cpuModel" style="margin-left: 10px">{{ item }}</el-tag>
        </el-form-item>
        <el-form-item label="最低ram要求:" prop="ram">
          <el-tag type="primary" v-text="LatestAlgorithm.ram + 'MB'" style="margin-left: 10px"/>
        </el-form-item>
        <el-form-item label="最低rom要求:" prop="rom">
          <el-tag type="primary" v-text="LatestAlgorithm.rom + 'MB'" style="margin-left: 10px"/>
        </el-form-item>
        <el-form-item label="算法描述:" prop="algorithmDescription">
          <label v-text="LatestAlgorithm.algorithmDescription" style="margin-left: 10px"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer align-right">
        <el-button @click="atiDetailsComfrim">确定</el-button>
      </div>
    </el-dialog>

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
              <el-button type="danger" icon="documentDelete">卸载模型</el-button>
              <!--              <el-button type="primary" class="el-button" @click="goToDetail">管理</el-button>-->
              <el-button type="primary" class="el-button" @click="showTaskDialog(algorithm)">{{ addBtnText }}</el-button>
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
