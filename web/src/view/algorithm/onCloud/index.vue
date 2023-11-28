<script setup>
import { onMounted, ref } from 'vue'
import { queryAllAlgorithm } from '@/api/algorithm'
import axios from 'axios'

// 拉取算法的地址
const getAlgoAddress = 'http://192.168.6.132:8888/Algorithm/algorithm/list'
const downloadAddress = ref('')
const dialogVisible = ref(false)

const cloud = ref([])



// TODO: 将下面的函数改为从服务器获取云端算法列表
const getOnCloudAlgorithms = async() => {
  const res = await axios.get(getAlgoAddress)
  cloud.value = res.data.rows
  console.log(res.data)
}

const downloadAlgorithm = () => {
  // TODO:download algorithm here
  location.href = 'http://192.168.6.132:8888/files/1699237799841'
  dialogVisible.value = false
}

onMounted(() => {
  getOnCloudAlgorithms()
})

</script>

<template>
  <div class="algorithm-container">
    <el-dialog
        v-model="dialogVisible"
        title="Tips"
        width="30%"
    >
      <span>确定要下载这个算法模型吗？</span>
      <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="downloadAlgorithm">
          确定
        </el-button>
      </span>
      </template>
    </el-dialog>
    <el-row class="algorithm-row">
      <span class="card-title">云端算法</span>
      <el-divider></el-divider>
      <el-col v-for="algorithm in cloud" :key="algorithm" class="algorithm-col" :span="8">
        <el-card class="cloud-card card">
          <div slot="header">
            <span class="title">{{ algorithm.algorithmName }}</span>
          </div>
          <div class="content-box algorithm-text">
            <span class="algorithm-version">
              算法版本：{{ algorithm.algorithmVersion }}
            </span>
            <br>
            <span class="algorithm-size">
              模型体积：{{ algorithm.algorithmSize }}
            </span>
            <br>
            <span class="algorithm-introduction">
              算法描述：{{ algorithm.algorithmDescription }}
            </span>
            <br>
            <span class="algorithm-upedateDate">
              更新日期：{{ algorithm.updateTime }}
            </span>
            <br>
            <div class="align-right">
              <el-button type="primary" class="el-button" @click="dialogVisible = true">点击下载到本地</el-button>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<style scoped lang="scss">
.algorithm {
  &-container {
    margin: 30px;
  }
  &-text {
    line-height: 35px;
    font-size: 15px;
  }
  &-row {
    margin-bottom: 20px;
  }
}
.align-right {
  text-align: right;
}
.card {
  margin-right: 20px;
  margin-bottom: 20px;
  &-title {
    font-size: 40px;
  }
}
</style>
