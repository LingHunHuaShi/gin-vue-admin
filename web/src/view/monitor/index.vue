<script setup>
import {onMounted, ref} from "vue";

const webSocketUrl = ref('ws://192.168.6.158:9999/rtsp?url=')
const streamNumber = ref(1)
const rtspStreams = ref([
  'rtsp://192.168.6.209:8554/stream',
  'rtsp://192.168.6.209:8554/stream2',
])

const selectDialogVisible = ref(false)
const selectedStreamOrder = ref(0)

const setRTSP = (streamOrder = 0) => {
  console.log('rtsp set:', rtspStreams.value[streamOrder])
  new JSMpeg.Player(webSocketUrl.value + btoa(rtspStreams.value[streamOrder]), {
    canvas: document.getElementById('canvas-'+streamOrder ),
  })
  new JSMpeg.Player(webSocketUrl.value + btoa(rtspStreams.value[streamOrder]), {
    canvas: document.getElementById('canvas'),
  })
}

const openSelectDialog = () => {
  selectDialogVisible.value = true
}

const setStream = (streamOrder) => {
  selectedStreamOrder.value = streamOrder
  setRTSP(streamOrder)
  selectDialogVisible.value = false
}

onMounted(()=>{
  setRTSP(0)
})
</script>

<template>
  <div>

    <el-dialog v-model="selectDialogVisible" title="视频流选择">
      <el-card style="padding: 0; margin: 10pt auto; width: 80%" v-for='(stream, index) in rtspStreams' :label="stream" :key="index"
      shadow="hover"
      @click="setStream(index)">
        <canvas style="height: auto; width: 100%; margin: 0" :id="'canvas-' + index"/>
        <div class="body" style="padding: 15pt">
          <span>{{ stream }}</span>
        </div>
      </el-card>
    </el-dialog>

      <div class="gva-card-box">
        <div class="gva-card">
          <div class="card-header">
            <span></span>
          </div>
          <div class="monitor-box">
            <el-card>
              <template #header>
                <el-text size="large" style="">实时图像</el-text>
                <el-button type="primary" plain style="margin-left: 10pt" @click="openSelectDialog">切换视频流</el-button>
              </template>
              <el-row>
<!--                <el-button @click="setRTSP">-->
<!--                  开始播放-->
<!--                </el-button>-->
              </el-row>
              <el-row>
                <el-col :xs="24" :sm="12" :md="12" :lg="24/streamNumber">
                  <canvas id="canvas" style="height: 100%; width: 100%;"/>
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