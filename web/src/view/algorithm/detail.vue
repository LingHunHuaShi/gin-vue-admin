<script>
export default {
  data() {
    return {
      fits: ['fill', 'contain', 'cover', 'none', 'scale-down'],
      url: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
      states: [
        {
          name: '当前帧推理时间',
          number: 0
        },
        {
          name: '识别物体数量',
          number: 0
        },
        {
          name: '分辨率',
          number: 0
        },
        {
          name: 'FPS',
          number: 0
        },
        {
          name: 'CPU占用率',
          number: 0
        },
        {
          name: 'NPU占用率',
          number: 0
        },
        {
          name: '内存占用',
          number: 0
        },
        {
          name: '当前算法版本',
          number: 0
        }
      ],
      options: [
        {
          value: '选项1',
          label: '240P'
        }, {
          value: '选项2',
          label: '480P'
        }, {
          value: '选项3',
          label: '720P'
        }, {
          value: '选项4',
          label: '1080P'
        }
      ],
      value: '',
      // 让选项显示在修改框中
      interval: 0,
      inputAddress: 'localhost:1002'
    }
  }
}
</script>

<template>
  <div class="detail-container">
    <el-row class="detail-row" :gutter="20" type="flex">
      <el-col class="detail-col" :span="16">
        <el-card class="pic-card">
          <div class="content-box">
            <el-image :src="url" :fit="fits" style="width: 400px; height: 400px" />
          </div>
        </el-card>
      </el-col>
      <el-col class="detail-col" :span="8">
        <el-card class="state-card" style="height: 100%">
          <div slot="header">
            <span>推理数据</span>
          </div>
          <div v-for="state in states" :key="state" class="content-box detail-text">
            {{ state.name }}: {{ state.number }}
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row class="detail-row" :gutter="20" type="flex">
      <el-col class="detail-col" :span="12">
        <el-card class="adjustment-card" style="height: 100%">
          <div slot="header">
            <span>参数修改</span>
          </div>
          <div class="content-box">
            <span>分辨率</span>
            <br>
            <div class="resolution-choose">
              <el-select v-model="value" placeholder="请选择">
                <el-option
                  v-for="item in options"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
            </div>
            <br>
            <span>识别间隔(帧): {{ interval }}</span>
            <br>
            <div class="interval-choose">
              <el-slider
                v-model="interval"
                :max="30"
                :step="5"
                :show-stops="true"
                :show-tooltip="false"
              />
            </div>
            <el-button type="primary">应用更改</el-button>
            <el-button>恢复默认</el-button>
          </div>
        </el-card>
      </el-col>
      <el-col class="detail-col" :span="12">
        <el-card class="configuration-card" style="height: 100%">
          <div slot="header">
            <span>算法配置</span>
          </div>
          <div class="content-box">
            <span>视频流地址</span>
            <el-input model="inputAddress" :placeholder="inputAddress" class="address-input" />
            <el-button type="primary">应用更改</el-button>
          </div>
        </el-card>
      </el-col>

    </el-row>
  </div>

</template>

<style scoped lang="scss">
.detail {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 18px;
    line-height: 35px;
  }
  &-row {
    margin-bottom: 20px;
  }
}

.resolution-choose {
  margin: 10px 0
}

.interval-choose {
  margin: 10px 3px
}

.address-input {
  margin: 10px 0
}

</style>
