<template>
  <div>
    <div class="gva-card-box">
      <div class="gva-card">
        <div class="card-header">
          <span>访问统计</span>
        </div>
        <div class="echart-box">
          <echarts-line />
        </div>
      </div>
    </div>
    <el-row :gutter="15" class="system_state">
      <el-col :span="12">
        <el-card v-if="state.os" class="card_item">
          <template #header>
            <div>盒子环境</div>
          </template>
          <div>
            <el-row :gutter="10">
              <el-col :span="12">操作系统:</el-col>
              <el-col :span="12" v-text="state.os.goos" />
            </el-row>
            <el-row :gutter="10">
              <el-col :span="12">CPU核心数:</el-col>
              <el-col :span="12" v-text="state.os.numCpu" />
            </el-row>
            <el-row :gutter="10">
              <el-col :span="12">编译器:</el-col>
              <el-col :span="12" v-text="state.os.compiler" />
            </el-row>
            <el-row :gutter="10">
              <el-col :span="12">Golang 版本:</el-col>
              <el-col :span="12" v-text="state.os.goVersion" />
            </el-row>
            <el-row :gutter="10">
              <el-col :span="12">Go 协程数:</el-col>
              <el-col :span="12" v-text="state.os.numGoroutine" />
            </el-row>
          </div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card v-if="state.disk" class="card_item">
          <template #header>
            <div>磁盘</div>
          </template>
          <div>
            <el-row :gutter="10">
              <el-col :span="12">
                <el-row :gutter="10">
                  <el-col :span="12">总计 (MB)</el-col>
                  <el-col :span="12" v-text="state.disk.totalMb" />
                </el-row>
                <el-row :gutter="10">
                  <el-col :span="12">已使用 (MB)</el-col>
                  <el-col :span="12" v-text="state.disk.usedMb" />
                </el-row>
                <el-row :gutter="10">
                  <el-col :span="12">总计 (GB)</el-col>
                  <el-col :span="12" v-text="state.disk.totalGb" />
                </el-row>
                <el-row :gutter="10">
                  <el-col :span="12">已使用 (GB)</el-col>
                  <el-col :span="12" v-text="state.disk.usedGb" />
                </el-row>
              </el-col>
              <el-col :span="12">
                <el-progress
                  type="dashboard"
                  :percentage="state.disk.usedPercent"
                  :color="colors"
                />
              </el-col>
            </el-row>
          </div>
        </el-card>
      </el-col>
    </el-row>
    <el-row :gutter="15" class="system_state">
      <el-col :span="12">
        <el-card
          v-if="state.cpu"
          class="card_item"
          :body-style="{ height: '180px', 'overflow-y': 'scroll' }"
        >
          <template #header>
            <div>CPU</div>
          </template>
          <div>
            <el-row :gutter="10">
              <el-col :span="12">物理核心数:</el-col>
              <el-col :span="12" v-text="state.cpu.cores" />
            </el-row>
            <el-row v-for="(item, index) in state.cpu.cpus" :key="index" :gutter="10">
              <el-col :span="12">核心 {{ index }}:</el-col>
              <el-col
                :span="12"
              ><el-progress
                type="line"
                :percentage="+item.toFixed(0)"
                :color="colors"
              /></el-col>
            </el-row>
          </div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card v-if="state.ram" class="card_item">
          <template #header>
            <div>内存</div>
          </template>
          <div>
            <el-row :gutter="10">
              <el-col :span="12">
                <el-row :gutter="10">
                  <el-col :span="12">总计 (MB)</el-col>
                  <el-col :span="12" v-text="state.ram.totalMb" />
                </el-row>
                <el-row :gutter="10">
                  <el-col :span="12">已使用 (MB)</el-col>
                  <el-col :span="12" v-text="state.ram.usedMb" />
                </el-row>
                <el-row :gutter="10">
                  <el-col :span="12">总计 (GB)</el-col>
                  <el-col :span="12" v-text="state.ram.totalMb / 1024" />
                </el-row>
                <el-row :gutter="10">
                  <el-col :span="12">已使用 (GB)</el-col>
                  <el-col
                    :span="12"
                    v-text="(state.ram.usedMb / 1024).toFixed(2)"
                  />
                </el-row>
              </el-col>
              <el-col :span="12">
                <el-progress
                  type="dashboard"
                  :percentage="state.ram.usedPercent"
                  :color="colors"
                />
              </el-col>
            </el-row>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { getSystemState } from '@/api/system'
import { onUnmounted, ref } from 'vue'
import EchartsLine from '@/view/dashboard/dashboardCharts/echartsLine.vue'
const timer = ref(null)
const state = ref({})
const colors = ref([
  { color: '#5cb87a', percentage: 20 },
  { color: '#e6a23c', percentage: 40 },
  { color: '#f56c6c', percentage: 80 }
])

const reload = async() => {
  const { data } = await getSystemState()
  state.value = data.server
}

reload()
timer.value = setInterval(() => {
  reload()
}, 1000 * 10)

onUnmounted(() => {
  clearInterval(timer.value)
  timer.value = null
})


</script>

<script>
export default {
  name: 'State',
}
</script>

<style>
.system_state {
  padding: 10px;
}

.card_item {
  height: 280px;
}
</style>
