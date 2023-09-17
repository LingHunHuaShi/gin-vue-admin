<template>
  <el-form :model="formData" ref="vForm" :rules="rules" label-position="left" label-width="150px"
           size="medium" @submit.prevent
  >
    <el-form-item label="发起人UUID" prop="inputUUID" class="required label-right-align">
      <el-input v-model="formData.inputUUID" type="text" clearable></el-input>
    </el-form-item>
    <el-form-item label="案例标题" prop="inputTitle" class="required label-right-align">
      <el-input v-model="formData.inputTitle" type="text" clearable></el-input>
    </el-form-item>
    <el-form-item label="案例内容" prop="textareaContent" class="required label-right-align">
      <el-input type="textarea" v-model="formData.textareaContent" rows="3"></el-input>
    </el-form-item>
    <el-form-item label="紧急等级" prop="radioServerity" class="label-right-align">
      <el-radio-group v-model="formData.radioServerity">
        <el-radio v-for="(item, index) in radioServerityOptions" :key="index" :label="item.value"
                  :disabled="item.disabled" style="{display: inline}"
        >{{ item.label }}
        </el-radio>
      </el-radio-group>
    </el-form-item>
    <el-form-item label="案例状态" prop="radioStatus" class="label-right-align">
      <el-radio-group v-model="formData.radioStatus">
        <el-radio v-for="(item, index) in radioStatusOptions" :key="index" :label="item.value"
                  :disabled="item.disabled" style="{display: inline}"
        >{{ item.label }}
        </el-radio>
      </el-radio-group>
    </el-form-item>
    <el-form-item label="解决方案" prop="textareaSolution" class="label-right-align">
      <el-input type="textarea" v-model="formData.textareaSolution" rows="3"></el-input>
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

const { expose } = useContext

export default defineComponent({
  components: {},
  props: {},
  setup() {
    const state = reactive({
      formData: {
        inputUUID: '',
        inputTitle: '',
        textareaContent: '',
        radioServerity: 1,
        radioStatus: null,
        textareaSolution: '',
      },
      rules: {
        inputUUID: [{
          required: true,
          message: '字段值不可为空',
        }],
        inputTitle: [{
          required: true,
          message: '字段值不可为空',
        }],
        textareaContent: [{
          required: true,
          message: '字段值不可为空',
        }],
      },
      radioServerityOptions: [{
        'label': '普通',
        'value': 1,
      }, {
        'label': '紧急',
        'value': 2,
      }, {
        'label': '非常紧急',
        'value': 3,
      }],
      radioStatusOptions: [{
        'label': '未处理',
        'value': 1,
      }, {
        'label': '正在处理',
        'value': 2,
      }, {
        'label': '已完成',
        'value': 3,
      }, {
        'value': 4,
        'label': '异常',
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