<script setup>
import {onMounted, ref} from "vue";

import flvjs from 'flv.js'

const webSocketUrl = ref('ws://localhost:9999/rtsp?url=')
const streamNumber = ref(1)
const rtspStreams = ref([
  'rtsp://192.168.6.129:8554/stream',
  'rtsp://192.168.6.129:8554/stream2',
])

const mainFlvPlayer = ref(null)

const selectDialogVisible = ref(false)
const selectedStreamOrder = ref(0)


const initPlayers = () => {
  if (flvjs.isSupported()) {
      const flvPlayer = flvjs.createPlayer({
        isLive: true,
        type: 'flv',
        url: webSocketUrl + btoa(rtspStreams.value[selectedStreamOrder]),
        enableWorker: true,
        enableStashBuffer: false,
        stashInitialSize: 128,
      },{
        deferLoadAfterSourceOpen: false
      })
      const player = document.getElementById("main-player")
      flvPlayer.attachMediaElement(player)
    return flvPlayer
    }
    console.log('init players done')
}

const play = async() => {
  if (mainFlvPlayer.value != null) {
    mainFlvPlayer.value.destroy()
    mainFlvPlayer.value = null
  }
  mainFlvPlayer.value = initPlayers()
  if (flvjs.isSupported()) {
    const mainPlayer = flvjs.createPlayer({
      isLive: true,
      type: 'flv',
      url: webSocketUrl.value + btoa(rtspStreams.value[selectedStreamOrder.value]),
      enableWorker: true,
      enableStashBuffer: false,
      stashInitialSize: 128,
    },{
      deferLoadAfterSourceOpen: false
    })
    console.log('rtsp url:', rtspStreams.value[selectedStreamOrder.value])
    console.log('ws url:', webSocketUrl.value + btoa(rtspStreams.value[selectedStreamOrder.value]))
    const flv = document.getElementById('main-player')
    mainPlayer.attachMediaElement(flv)
    try {
      mainPlayer.load()

      let controller = mainPlayer._transmuxer._controller
      let wsLoader = controller._ioctl._loader
      var oldWsOnCompleteFunc = wsLoader._onComplete
      wsLoader._onComplete = function() {
        if(!controller._remuxer) {
          controller._remuxer = {
            flushStashedSamples: function () {
              // _this.loadingVisiable = false
              console.log("flushStashedSamples")
            }
          }
        }
        oldWsOnCompleteFunc()
      }

      mainPlayer.play()
      console.log('successfully created main player')
    } catch(e) {console.log(e)}
  }
}


const openSelectDialog = () => {
  selectDialogVisible.value = true
}

const setStream = (streamOrder) => {
  if (mainFlvPlayer.value != null) {
    mainFlvPlayer.value.destroy()
    mainFlvPlayer.value = null
  }
  selectedStreamOrder.value = streamOrder
  selectDialogVisible.value = false
}

onMounted(()=>{
  // initPlayers()
})

</script>

<template>
  <div>

    <el-dialog v-model="selectDialogVisible" title="视频流选择">
      <el-card style="padding: 0; margin: 10pt auto; width: 80%" v-for='(stream, index) in rtspStreams' :label="stream" :key="index"
      shadow="hover"
      @click="setStream(index)">
<!--        <canvas style="height: auto; width: 100%; margin: 0"  :id="'canvas-' + index"/>-->
<!--        <video ref="playerRefs" style="height: auto; width: 100%; margin: 0" muted controls loop></video>-->
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
                <el-button type="success" plain style="margin-left: 10pt" @click="play" icon="videoPlay"></el-button>
              </template>
              <el-row>
<!--                <el-button @click="setRTSP">-->
<!--                  开始播放-->
<!--                </el-button>-->
              </el-row>
              <el-row>
                <el-col :xs="24" :sm="24" :md="18" :lg="24">
<!--                  <canvas id="canvas" style="height: 100%; width: 100%;"/>-->
                  <video id="main-player" style="height: 90%; width: 90%;" controls loop></video>
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