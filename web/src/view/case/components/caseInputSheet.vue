<template>
  <el-form :model="formData" ref="vForm" :rules="rules" label-position="left" label-width="150px"
           size="medium" @submit.prevent
  >
    <el-form-item label="发起人UUID" prop="uuid" class="required label-right-align">
      <el-input v-model="formData.uuid" type="text" clearable :disabled="true"></el-input>
    </el-form-item>
    <el-form-item label="案例标题" prop="title" class="required label-right-align">
      <el-input v-model="formData.title" type="text" clearable></el-input>
    </el-form-item>
    <el-form-item label="案例内容" prop="content" class="required label-right-align">
      <el-input type="textarea" v-model="formData.content" rows="3"></el-input>
    </el-form-item>
    <el-form-item label="紧急等级" prop="severity" class="label-right-align">
      <el-radio-group v-model="formData.severity">
        <el-radio v-for="(item, index) in severityOptions" :key="index" :label="item.value"
                  :disabled="item.disabled" style="{display: inline}"
        >{{ item.label }}
        </el-radio>
      </el-radio-group>
    </el-form-item>
    <el-form-item label="案例状态" prop="status" class="label-right-align">
      <el-radio-group v-model="formData.status">
        <el-radio v-for="(item, index) in statusOptions" :key="index" :label="item.value"
                  :disabled="item.disabled" style="{display: inline}"
        >{{ item.label }}
        </el-radio>
      </el-radio-group>
    </el-form-item>
    <el-form-item label="解决方案" prop="solution" class="label-right-align">
      <el-input type="textarea" v-model="formData.solution" rows="3"></el-input>
    </el-form-item>
  </el-form>

</template>

<script>
import {
  defineComponent,
  toRefs,
  reactive,
  getCurrentInstance, onMounted,
}
  from 'vue'
import { createCase, findCaseByCaseID, updateCaseByCaseID } from '@/api/case'
import { getUserInfo } from '@/api/user'
import { ElMessage } from 'element-plus'

export default defineComponent({
  components: {},
  props: {
    key: {
      type: String,
      required: true,
      default: 'edit',
    },
    scope: {
      type: Object,
      required: true,
    }
  },
  setup(props) {
    const state = reactive({
      formData: {
        uuid: '',
        title: '',
        content: '',
        severity: 1,
        status: null,
        solution: '',
      },
      rules: {
        uuid: [{
          required: true,
          message: '字段值不可为空',
        }],
        title: [{
          required: true,
          message: '字段值不可为空',
        }],
        content: [{
          required: true,
          message: '字段值不可为空',
        }],
      },
      severityOptions: [{
        'label': '普通',
        'value': 0,
      }, {
        'label': '紧急',
        'value': 1,
      }, {
        'label': '非常紧急',
        'value': 2,
      }],
      statusOptions: [{
        'label': '未处理',
        'value': 0,
      }, {
        'label': '正在处理',
        'value': 1,
      }, {
        'label': '已完成',
        'value': 2,
      }, {
        'value': 3,
        'label': '异常',
      }],
    })
    const setCurrentUser = async() => {
      await getUserInfo().then((result) => {
        state.formData.uuid = result.data.userInfo.uuid
      })
    }
    const setEditInfo = async() => {
      if (props.key === 'edit') {
        await findCaseByCaseID({ ID: props.scope.caseID }).then(result => {
          state.formData.title = result.data.title
          state.formData.content = result.data.content
          state.formData.severity = result.data.severity
          state.formData.solution = result.data.solution
          state.formData.status = result.data.status
        }).catch(err => {
          ElMessage({
            type: 'error',
            message: err
          })
        })
      }
    }
    onMounted(() => {
      setCurrentUser()
      setEditInfo()
    })
    const instance = getCurrentInstance()
    const submitForm = () => {
      instance.ctx.$refs['vForm'].validate(async(valid) => {
        if (!valid) return
        // TODO: 提交表单
        // console.log(state.formData.uuid)
        if (props.key === 'add') {
          await createCase(state.formData).then(res => {
            if (res.code === 0) {
              ElMessage({
                type: 'success',
                message: '添加成功! 请刷新页面'
              })
            }
          }).catch(err => {
            ElMessage({
              type: 'error',
              message: err,
            })
          })
        }
        if (props.key === 'edit') {
          await updateCaseByCaseID(state.formData).then(res => {
            if (res.code === 0) {
              ElMessage({
                type: 'success',
                message: '更新成功！请刷新页面',
              })
            }
          }).catch(err => {
            ElMessage({
              type: 'error',
              message: err,
            })
          })
        }
      })
    }
    const resetForm = () => {
      instance.ctx.$refs['vForm'].resetFields()
    }
    const testFunc = () => {
      console.log('test function is called')
    }

    return {
      ...toRefs(state),
      submitForm,
      resetForm,
      testFunc,
      setCurrentUser,
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