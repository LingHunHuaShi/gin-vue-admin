<template>
  <el-form :model="formData" ref="vForm" :rules="rules" label-position="left" label-width="150px"
           size="medium" @submit.prevent
  >
    <el-form-item label="发起人UUID" prop="uuid" class="required label-right-align">
      <el-input v-model="formData.uuid" type="text" clearable :disabled="true"></el-input>
    </el-form-item>
    <el-form-item label="视频源" prop="source" class="required label-right-align">
      <el-input v-model="formData.source" type="text" clearable></el-input>
    </el-form-item>
    <el-form-item label="分辨率" prop="resolution" class="label-right-align">
      <el-select v-model="formData.resolution" class="full-width-input" clearable>
        <el-option v-for="(item, index) in resolutionOptions" :key="index" :label="item.label"
                   :value="item.value" :disabled="item.disabled"
        ></el-option>
      </el-select>
    </el-form-item>
    <el-form-item label="任务容器算法" prop="algorithm" class="label-right-align">
      <el-text> {{ algoName }} </el-text>
    </el-form-item>
    <el-form-item label="粒度" prop="intensity" class="label-right-align">
      <el-radio-group v-model="formData.intensity">
        <el-radio v-for="(item, index) in intensityOptions" :key="index" :label="item.value"
                  :disabled="item.disabled" style="{display: inline}"
        >{{ item.label }}
        </el-radio>
      </el-radio-group>
    </el-form-item>
  </el-form>
</template>

<script>
import {
  defineComponent,
  toRefs,
  reactive,
  getCurrentInstance, onMounted, ref, watch,
}
  from 'vue'
import { createTask } from '@/api/task'
import { getUserInfo } from '@/api/user'
import { queryAllAlgorithm } from '@/api/algorithm'
import async from 'async'
import { ElMessage } from 'element-plus'

export default defineComponent({
  name: 'TaskInput',
  components: {},
  props: {
    algoName: {
      type: String,
      required: true,
      default: 'defaultName',
    },
    algoId: {
      type: Number,
      required: true,
      default: -1,
    },
  },
  setup(props) {
    console.log('props initial:', props)
    const state = reactive({
      formData: {
        uuid: '',
        source: '',
        resolution: '',
        algorithm: '',
        intensity: null,
        algoName: '1111111',
        algoId: '',
      },
      rules: {
        uuid: [{
          required: true,
          message: '字段值不可为空',
        }],
        source: [{
          required: true,
          message: '字段值不可为空',
        }],
      },
      resolutionOptions: [{
        'label': '360P',
        'value': 1,
      }, {
        'label': '480P',
        'value': 2,
      }, {
        'label': '540P',
        'value': 3,
      }, {
        'value': 4,
        'label': '720P',
      }, {
        'value': 5,
        'label': '1080P',
      }],
      intensityOptions: [{
        'label': '粗',
        'value': 1,
      }, {
        'label': '中',
        'value': 2,
      }, {
        'label': '细',
        'value': 3,
      }],
    })
    const setCurrentUser = async() => {
      await getUserInfo().then((result) => {
        state.formData.uuid = result.data.userInfo.uuid
      })
    }

    onMounted(() => {
      setCurrentUser()
    })

    const instance = getCurrentInstance()
    const submitForm = () => {
      instance.ctx.$refs['vForm'].validate(async(valid) => {
        if (!valid) return
        // TODO: 提交表单
        console.log('formData:' + state.formData.source)
        const taskData = {
          uuid: state.formData.uuid,
          source: state.formData.source,
          resolution: state.formData.resolution,
          algorithm: state.formData.algoId,
          intensity: state.formData.intensity,
        }
        await createTask(taskData).then(res => {
          if (res.code === 0) {
            ElMessage({
              type: 'success',
              message: '添加成功！'
            })
          }
        }).catch(err => {
          console.error('Error:', err)
        })
      })
    }
    const resetForm = () => {
      instance.ctx.$refs['vForm'].resetFields()
    }

    return {
      ...toRefs(state),
      submitForm,
      resetForm,
    }
  },
})

</script>

<style lang="scss">
.el-input-number.full-width-input,
.el-cascader.full-width-input {
  width: 100% !important;
}

.el-form-item--medium {
  .el-radio {
    line-height: 36px !important;
  }

  .el-rate {
    margin-top: 8px;
  }
}

.el-form-item--small {
  .el-radio {
    line-height: 32px !important;
  }

  .el-rate {
    margin-top: 6px;
  }
}

.el-form-item--mini {
  .el-radio {
    line-height: 28px !important;
  }

  .el-rate {
    margin-top: 4px;
  }
}

.clear-fix:before,
.clear-fix:after {
  display: table;
  content: "";
}

.clear-fix:after {
  clear: both;
}

.float-right {
  float: right;
}

</style>

<style lang="scss" scoped>
div.table-container {
  table.table-layout {
    width: 100%;
    table-layout: fixed;
    border-collapse: collapse;

    td.table-cell {
      display: table-cell;
      height: 36px;
      border: 1px solid #e1e2e3;
    }
  }
}

div.tab-container {
}

.label-left-align :deep(.el-form-item__label) {
  text-align: left;
}

.label-center-align :deep(.el-form-item__label) {
  text-align: center;
}

.label-right-align :deep(.el-form-item__label) {
  text-align: right;
}

.custom-label {
}

.static-content-item {
  min-height: 20px;
  display: flex;
  align-items: center;

  :deep(.el-divider--horizontal) {
    margin: 0;
  }
}

</style>