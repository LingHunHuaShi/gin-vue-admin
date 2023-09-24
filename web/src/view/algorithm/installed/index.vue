<script setup>

import { computed, onMounted, reactive, ref } from 'vue'
import detail from '../detail.vue'
import { queryOngoingTask } from '@/api/task'
import { queryAllAlgorithm } from '@/api/algorithm'

const installed = ref([])

const getInstalledAlgorithm = async() => {
  await queryAllAlgorithm().then((result) => {
    installed.value = result.data
    console.log(result.data[0].UpdatedAt.toLocaleString())
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
