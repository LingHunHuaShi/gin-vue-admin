<script>

export default {
  data() {
    return {
      cloud: [ // 在云端可部署的算法
        {
          name: 'Yolo V8',
          introduction: 'Introduction for Yolo V8.Introduction for Yolo V8.Introduction for Yolo V8.Introduction for ' +
            'Yolo V8.Introduction for Yolo V8.Introduction for Yolo V8.Introduction for Yolo V8.'
        },
        {
          name: 'Yolo V8',
          introduction: 'Introduction for Yolo V8.Introduction for Yolo V8.Introduction for Yolo V8.Introduction for ' +
            'Yolo V8.Introduction for Yolo V8.Introduction for Yolo V8.Introduction for Yolo V8.'
        },
        {
          name: 'Yolo V8',
          introduction: 'Introduction for Yolo V8.Introduction for Yolo V8.Introduction for Yolo V8.Introduction for ' +
            'Yolo V8.Introduction for Yolo V8.Introduction for Yolo V8.Introduction for Yolo V8.'
        },
        {
          name: 'Yolo V8',
          introduction: 'Introduction for Yolo V8.Introduction for Yolo V8.Introduction for Yolo V8.Introduction for ' +
            'Yolo V8.Introduction for Yolo V8.Introduction for Yolo V8.Introduction for Yolo V8.'
        }
      ],
      dialogFormVisible: false,
      formLabelWidth: '120px',
      newAlgorithm: {
        VideoAddress: ''
      }
    }
  }
}
</script>

<template>
  <div class="algorithm-container">
    <el-dialog v-model="dialogFormVisible" title="基础配置">
      <el-form :model="newAlgorithm">
        <el-form-item label="视频流地址" :label-width="formLabelWidth">
          <el-input v-model="newAlgorithm.VideoAddress" autocomplete="off" type="url" placeholder="http://example.com" pattern="http://.*" size="30" required />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer align-right">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="dialogFormVisible = false">确定部署</el-button>
      </div>
    </el-dialog>
    <el-row class="algorithm-row">
      <h4 class="card-title">云端算法</h4>
      <el-divider></el-divider>
      <el-col v-for="algorithm in cloud" :key="algorithm" class="algorithm-col" :span="8">
        <el-card class="cloud-card card">
          <div slot="header">
            <span class="title">{{ algorithm.name }}</span>
          </div>
          <div class="content-box algorithm-text">
            <span class="algorithm-introduction">
              {{ algorithm.introduction }}
            </span>
            <br>
            <div class="align-right">
              <el-button type="primary" class="el-button" @click="dialogFormVisible = true">点击部署</el-button>
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
}
</style>
