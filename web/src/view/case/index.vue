<script setup>
// import { createCase } from '@/api/case'
import { deleteCase, queryAllCases } from '@/api/case'
import { onMounted, ref } from 'vue'
import CaseInputSheet from '@/view/case/components/caseInputSheet.vue'

const caseData = ref([])
const caseDialogVisible = ref(false)
const dialogTitle = ref('新增案例')
const caseRef = ref(null)

const showInputDialog = () => {
  caseDialogVisible.value = true
  caseRef.setCurrentUser()
}

const closeDialog = () => {
  caseDialogVisible.value = false
}

const submitDialog = () => {
  caseRef.value.submitForm()
  caseDialogVisible.value = false
}

const getCaseData = async() => {
  await queryAllCases().then((result) => {npm
    caseData.value = result.data
    console.log(caseData)
  })
}

onMounted(() => {
  getCaseData()
})

</script>

<template>
  <div class="case">
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="showInputDialog">新增案例</el-button>
      </div>
      <el-table
        :data="caseData"
        row-key="caseID"
        style="width: 100%"
      >
        <el-table-column label="案件ID" min-width="180" prop="caseID"/>
        <el-table-column align="left" label="发起人UUID" min-width="180" prop="uuid"/>
        <el-table-column align="left" label="案例标题" min-width="180" prop="title"/>
        <el-table-column align="left" label="提交时间" min-width="180" prop="submitDate"/>
        <el-table-column align="left" label="紧急等级" min-width="180" prop="severity"/>
        <el-table-column align="left" label="案例状态" min-width="180" prop="status"/>
        <el-table-column align="left" label="完结时间" min-width="180" prop="dateClosed"/>
        <el-table-column align="left" label="操作" width="460">
          <template #default="scope">
            <el-button
              icon="view"
              type="primary"
              link
            >查看详情
            </el-button>
            <el-button
              icon="delete"
              type="primary"
              link
              @click="deleteCase(scope.row)"
            >删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-dialog v-model="caseDialogVisible" :title="dialogTitle">
      <case-input-sheet ref="caseRef" />
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="submitDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<style scoped lang="scss">

</style>