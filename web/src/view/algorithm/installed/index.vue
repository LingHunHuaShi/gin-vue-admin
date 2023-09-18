<script setup>
import { queryAllAlgorithm, updateAlgorithm } from '@/api/algorithm'
import { computed, onMounted, reactive, ref } from 'vue'
import detail from '../detail.vue'

const installed = ref([])

// const installed = reactive([{
//   name: 'Yolo V5 Lite-g',
//   CPU_rate: '5%',
//   NPU_rate: '37%',
//   run_time: '6:35:09',
//   stream_number: 1,
// },
// {
//   name: 'Yolo V5 Lite-g',
//   CPU_rate: '5%',
//   NPU_rate: '37%',
//   run_time: '6:35:09',
//   stream_number: 1,
// }])

const getInstalledAlgorithm = async() => {
  await queryAllAlgorithm().then((result) => {
    installed.value = result.data
    console.log(result)
  })
}

onMounted(() => {
  getInstalledAlgorithm()
})

const usedNPUCore = computed(() => {
  let totalCore = 0
  for (let i = 0; i < installed.length; i++) {
    totalCore += installed[i].stream_number
  }
  return totalCore
})

const testAlgorithm = {
  algorithmID: 1,
  algorithmName: 'test',
  algorithmVersion: '1.0',
  description: 'None',
  size: 190.8,
  downloadLink: 'None',
  MD5: 'None'
}


const dialogVisible = ref(false)
</script>

<template>
  <div class="algorithm-container">
    <el-dialog v-model="dialogVisible" title="算法配置">
      <detail />
      <div slot="footer" class="dialog-footer align-right">
        <el-button @click="dialogVisible = false">取消更改</el-button>
      </div>
    </el-dialog>
    <el-row class="algorithm-row">
      <el-col :span="8" class="algorithm-col">
        <el-card class="conclusion-card card" :gutter="20" type="flex">
          <div slot="header">
            <span>核心使用情况</span>
            <br>
            <br>
          </div>
          <div class="content-box">
            <span style="line-height: 35px">已使用:{{ usedNPUCore }} / 3</span>
            <br>
            <div v-for="algorithm in installed" :key="algorithm" class="algorithm-text">
              {{ algorithm.name }} 使用核心数：{{ algorithm.stream_number }}
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row class="algorithm-row">
      <span class="title">已部署算法</span>
      <el-divider />
      <el-col v-for="algorithm in installed" :key="algorithm" class="algorithm-col" :span="8">
        <el-card class="installed-card card">
          <div slot="header">
            <span class="title algorithm-title">{{ algorithm.algorithmName }}</span>
          </div>
          <div class="content-box algorithm-text">
            <span>算法版本：{{ algorithm.algorithmVersion }}</span>
            <br>
            <span>NPU占用率：{{ algorithm.NPU_rate }}</span>
            <br>
            <span>运行时间：{{ algorithm.run_time }}</span>
            <br>
            <div class="align-right">
              <!--              <el-button type="primary" class="el-button" @click="goToDetail">管理</el-button>-->
              <el-button type="primary" class="el-button" @click="dialogVisible = true">管理</el-button>
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
