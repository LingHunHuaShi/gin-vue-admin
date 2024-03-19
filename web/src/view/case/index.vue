<script setup>
import {createCase, deleteCase, queryAllCases, updateCase} from '@/api/case'
import {getCurrentInstance, onMounted, ref} from 'vue'
import {findNickNameByUuid, getUserInfo} from '@/api/user'
import { ElMessage } from 'element-plus'

const caseData = ref([])
const caseDialogVisible = ref(false)
const dialogTitle = ref('新增案例')
const dialogKey = ref('add')
const formData = ref({})

const setCurrentUser = async() => {
  await getUserInfo().then((result) => {
    formData.value.uuid = result.data.userInfo.uuid
  })
  console.log(formData.value)
}


const showInputDialog = () => {
  caseDialogVisible.value = true
  setCurrentUser()
}

const closeDialog = () => {
  caseDialogVisible.value = false
}

const instance = getCurrentInstance()

const submitDialog = () => {
  instance.ctx.$refs['vForm'].validate(async(valid) => {
    if (!valid)
      return
    if (dialogKey.value === 'add') {
      await createCase(formData.value).then(res => {
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '添加成功! '
          })
          location.reload()
        }
      }).catch(err => {
        ElMessage({
          type: 'error',
          message: err,
        })
      })
    }
    if (dialogKey.value === 'edit') {
      await updateCase({
        id: formData.value.caseID,
        title: formData.value.title,
        severity: formData.value.severity,
        status: formData.value.status,
        solution: formData.value.solution,
        content: formData.value.content,
      }).then(res => {
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '更新成功！',
          })
          location.reload()
        }
      }).catch(err => {
        ElMessage({
          type: 'error',
          message: err,
        })
      })
    }
    caseDialogVisible.value = false
  })
}

const getCaseData = async() => {
  const result = await queryAllCases()
  await processCaseList(result)
}

const processCaseList = async(result) => {
  caseData.value = await Promise.all(result.data.map(async(item) => {
    const nickName = await getNickNameByUuid(item.uuid)
    const severityText = getSeverity(item.severity)
    const statusText = getStatus(item.status)
    return {
      caseID: item.ID,
      nickName: nickName,
      CreatedAt: item.CreatedAt.substring(0, 10) + ' ' + item.CreatedAt.substring(11, 19),
      content: item.content,
      title: item.title,
      severity: item.severity,
      status: item.status,
      severityText: severityText,
      statusText: statusText,
      dataClosed: '未完成',
      solution: item.solution,
    }
  }))
}

const severityOptions = [{
  'label': '普通',
  'value': 0,
}, {
  'label': '紧急',
  'value': 1,
}, {
  'label': '非常紧急',
  'value': 2,
}]
const statusOptions = [{
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
}]

const rules = {
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
}

const getSeverity = (num) => {
  switch (num) {
    case 0:
      return '普通'
    case 1:
      return '紧急'
    case 2:
      return '非常紧急'
    default:
      return 'Unknown'
  }
}

const getStatus = (num) => {
  switch (num) {
    case 0:
      return '未处理'
    case 1:
      return '正在处理'
    case 2:
      return '已完成'
    case 3:
      return '异常'
    default:
      return 'Unknown'
  }
}

const getNickNameByUuid = async(uuid) => {
  const res = await findNickNameByUuid({ uuid: uuid })
  // console.log(res.data.NickName)
  return res.data.NickName
}

const deleteCaseInForm = async(caseID) => {
  console.log('caseID:', caseID)
  await deleteCase({ ID: caseID }).then(res => {
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功！'
      })
      location.reload()
    }
  }).catch(err => {
    ElMessage({
      type: 'error',
      message: err
    })
  })
}

const addCase = () => {
  dialogKey.value = 'add'
  dialogTitle.value = '新增案例'
  formData.value = {}
  showInputDialog()
}

const editCase = (scope) => {
  dialogKey.value = 'edit'
  dialogTitle.value = '编辑案例'
  formData.value = scope
  showInputDialog()
}

onMounted(() => {
  getCaseData()
})

</script>

<template>
  <div class="case">
    <el-dialog v-model="caseDialogVisible" :title="dialogTitle">
<!--      <case-input-sheet ref="caseRef" :key="dialogKey" :scope="dialogData" />-->

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

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="submitDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="addCase">新增案例</el-button>
      </div>
      <el-table
        :data="caseData"
        row-key="caseID"
        style="width: 100%"
      >
        <el-table-column label="案件ID" min-width="180" prop="caseID"/>
        <el-table-column align="left" label="发起人" min-width="180" prop="nickName"/>
        <el-table-column align="left" label="案例标题" min-width="180" prop="title"/>
        <el-table-column align="left" label="提交时间" min-width="180" prop="CreatedAt"/>
        <el-table-column align="left" label="紧急等级" min-width="180" prop="severityText"/>
        <el-table-column align="left" label="案例状态" min-width="180" prop="statusText"/>
        <el-table-column align="left" label="完结时间" min-width="180" prop="dataClosed"/>
        <el-table-column align="left" label="操作" width="460">
          <template #default="scope">
            <el-button
              icon="edit"
              type="primary"
              link
              @click="editCase(scope.row)"
            >编辑
            </el-button>
            <el-button
              icon="delete"
              type="primary"
              link
              @click="deleteCaseInForm(scope.row.caseID)"
            >删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<style scoped lang="scss">

</style>