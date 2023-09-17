<template>
  <el-form :model="formData" ref="vForm" :rules="rules" label-position="left" label-width="150px"
           size="medium" @submit.prevent
  >
    <div class="static-content-item">
      <div>单列表单</div>
    </div>
    <div class="static-content-item">
      <el-divider direction="horizontal"></el-divider>
    </div>
    <el-form-item label="发起人UUID" prop="inputUUID" class="required label-right-align">
      <el-input v-model="formData.inputUUID" type="text" clearable></el-input>
    </el-form-item>
    <el-form-item label="视频源" prop="inputSource" class="required label-right-align">
      <el-input v-model="formData.inputSource" type="text" clearable></el-input>
    </el-form-item>
    <el-form-item label="分辨率" prop="selectResolution" class="label-right-align">
      <el-select v-model="formData.selectResolution" class="full-width-input" clearable>
        <el-option v-for="(item, index) in selectResolutionOptions" :key="index" :label="item.label"
                   :value="item.value" :disabled="item.disabled"
        ></el-option>
      </el-select>
    </el-form-item>
    <el-form-item label="任务容器算法" prop="selectAlgorithm" class="label-right-align">
      <el-select v-model="formData.selectAlgorithm" class="full-width-input" clearable>
        <el-option v-for="(item, index) in selectAlgorithmOptions" :key="index" :label="item.label"
                   :value="item.value" :disabled="item.disabled"
        ></el-option>
      </el-select>
    </el-form-item>
    <el-form-item label="粒度" prop="radioIntensity" class="label-right-align">
      <el-radio-group v-model="formData.radioIntensity">
        <el-radio v-for="(item, index) in radioIntensityOptions" :key="index" :label="item.value"
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
  getCurrentInstance,
}
  from 'vue'

export default defineComponent({
  components: {},
  props: {},
  setup() {
    const state = reactive({
      formData: {
        inputUUID: '',
        inputSource: '',
        selectResolution: '',
        selectAlgorithm: '',
        radioIntensity: null,
      },
      rules: {
        inputUUID: [{
          required: true,
          message: '字段值不可为空',
        }],
        inputSource: [{
          required: true,
          message: '字段值不可为空',
        }],
      },
      selectResolutionOptions: [{
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
      selectAlgorithmOptions: [{
        'label': 'select 1',
        'value': 1,
      }, {
        'label': 'select 2',
        'value': 2,
      }, {
        'label': 'select 3',
        'value': 3,
      }],
      radioIntensityOptions: [{
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
    const instance = getCurrentInstance()
    const submitForm = () => {
      instance.ctx.$refs['vForm'].validate(valid => {
        if (!valid) return
        //TODO: 提交表单
      })
    }
    const resetForm = () => {
      instance.ctx.$refs['vForm'].resetFields()
    }
    const expose = { ...toRefs(state), submitForm, resetForm }
    defineExpose(
      expose
    )
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