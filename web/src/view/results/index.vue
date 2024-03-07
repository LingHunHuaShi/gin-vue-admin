<template>
  <div v-loading.fullscreen.lock="fullscreenLoading">
    <div class="gva-table-box">
      <el-table :data="tableData">
        <el-table-column align="left" label="预览" width="100">
          <template #default="scope">
<!--            <CustomPic pic-type="file" :pic-src="scope.row.url" preview/>-->
            <CustomPic pic-type="file" :pic-src="getPath(scope.row)" preview/>

          </template>
        </el-table-column>
        <el-table-column align="left" label="任务ID" prop="taskID" width="180">
          <template #default="scope">
            <div>{{ scope.row.taskID }}</div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="算法名" prop="name" min-width="180">
          <template #default="scope">
            <div class="name">{{ scope.row.algorithmName }}</div>
          </template>
        </el-table-column>
<!--        <el-table-column align="left" label="更新时间" prop="updatedAt" min-width="180">-->
<!--          <template #default="scope">-->
<!--            <div>{{ formatDate(scope.row.updatedAt) }}</div>-->
<!--          </template>-->
<!--        </el-table-column>-->
        <el-table-column align="left" label="操作" width="160">
          <template #default="scope">
            <el-button icon="download" type="primary" link @click="downloadFile(scope.row)">下载</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script setup>
import { getFileList, deleteFile, editFileName } from '@/api/fileUploadAndDownload'
import { downloadImage } from '@/utils/downloadImg'
import CustomPic from '@/components/customPic/index.vue'
import UploadImage from '@/components/upload/image.vue'
import UploadCommon from '@/components/upload/common.vue'
import { formatDate } from '@/utils/format'
import WarningBar from '@/components/warningBar/warningBar.vue'

import {onBeforeUnmount, onMounted, ref} from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {queryOngoingTask} from "@/api/task";
import {findAlgorithmById} from "@/api/algorithm";


const path = ref(import.meta.env.VITE_BASE_API)
const timer = ref(null)
const picPath = "images/rickroll.png"

// const imageUrl = ref('')
// const imageCommon = ref('')

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const search = ref({})
const tableData = ref([])


// 查询
const getTableData = async() => {
  // const table = await getFileList({ page: page.value, pageSize: pageSize.value, ...search.value })
  // if (table.code === 0) {
  //   tableData.value = table.data.list
  //   total.value = table.data.total
  //   page.value = table.data.page
  //   pageSize.value = table.data.pageSize
  // }
  // console.log(tableData.value)

  const table = await queryOngoingTask();
  if (table.code === 0) {
    tableData.value = await Promise.all(table.data.map(async(item) => {
      const algorithmName = await findAlgorithmById({ ID: item.algorithmId }).then((res) => {
        return res.data.algorithmName
      })
      const updatedAt = item.UpdatedAt
      const taskID = item.ID


      return {
        taskID: taskID,
        algorithmName: algorithmName,
        updatedAt: updatedAt,
      }
    }))
    total.value = table.data.length
  }


}
getTableData()


const downloadFile = (row) => {
  if (row.url.indexOf('http://') > -1 || row.url.indexOf('https://') > -1) {
    downloadImage(row.url, row.name)
  } else {
    debugger
    downloadImage(path.value + '/' + row.url, row.name)
  }
}

const getPath = (row) => {
  return "images/" +
      row.algorithmName.replace(/\s/g, "").toLowerCase() +
      "/" + row.taskID + "/output.jpg"
}

onMounted(() => {
  timer.value = setInterval(getTableData, 1000)
})

onBeforeUnmount(() => {
  clearInterval(timer.value)
})


</script>

<script>

export default {
  name: 'Upload',
}
</script>
<style scoped>
.name {
  cursor: pointer;
}

.upload-btn + .upload-btn {
  margin-left: 12px;
}
</style>
