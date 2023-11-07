<script setup>
import {onMounted, ref} from "vue";

const streamNumber = ref(1)
const streams = ref([])
const rtsp1 = 'rtsp://127.0.0.1:8554/stream'


const setRTSP = () => {
  console.log('rtsp set')
  new JSMpeg.Player('ws://localhost:9999/rtsp?url=' + btoa(rtsp1), {
    canvas1: document.getElementById('canvas-1'),
    // canvas2: document.getElementById('canvas-2'),
  })
}

onMounted(()=>{

})
</script>

<template>
  <div>
      <div class="gva-card-box">
        <div class="gva-card">
          <div class="card-header">
            <span></span>
          </div>
          <div class="monitor-box">
            <el-card>
              <template #header>
                <el-text size="large" style="">实时图像</el-text>
              </template>
              <el-row>
                <el-button @click="setRTSP">
  <!--                <canvas id="canvas-2" style="height: 300px; width:100px;"/>-->
                </el-button>

                <el-col :sm="12" :md="6" :lg="6" :xs="12" :xl="4">
  <!--                <el-select-->
  <!--                    v-model="streams"-->
  <!--                    placeholder="选择展示的流"-->
  <!--                    multiple-->
  <!--                    collapse-tags-->
  <!--                    collapse-tags-tooltip-->
  <!--                >-->
  <!--&lt;!&ndash;                  <el-option v-for="(task, index) in taskList" :key="index" :label="'算法:'+task.algorithmName+' 视频源:'+task.videoSource" :value="task.ID" />&ndash;&gt;-->
  <!--                  <el-option>-->
  <!--                    <canvas id="canvas-2" style="max-height: 500px" />-->
  <!--                  </el-option>-->
  <!--                </el-select>-->

                </el-col>
              </el-row>
              <el-row>
  <!--              <el-col v-for="(stream, index) in streams" :key="index" :xs="24" :sm="12" :md="12" :lg="24/streamNumber">-->
  <!--                <canvas id="canvas-1" />-->
  <!--              </el-col>-->
                <el-col :xs="24" :sm="12" :md="12" :lg="24/streamNumber">
                  <canvas id="canvas-1" style="height: 10%; width: 100%;" />
                </el-col>
              </el-row>
            </el-card>
          </div>

        </div>
      </div>
  </div>
</template>

<style scoped lang="scss">

</style>