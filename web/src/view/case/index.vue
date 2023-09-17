<script setup>
// import { createCase } from '@/api/case'
import { deleteCase } from '@/api/case'
import { ref } from 'vue'

const caseData = ref([])
</script>

<template>
  <div class="case">
    <div class="gva-btn-list">
      <el-button type="primary" icon="plus">新增案例</el-button>
    </div>
    <el-table
      :data="caseData"
      row-key="caseID"
      style="width: 100%"
    >
      <el-table-column label="案件ID" min-width="180" prop="caseID" />
      <el-table-column align="left" label="发起人UUID" min-width="180" prop="uuid" />
      <el-table-column align="left" label="案例标题" min-width="180" prop="title" />
      <el-table-column align="left" label="提交时间" min-width="180" prop="submitDate" />
      <el-table-column align="left" label="紧急等级" min-width="180" prop="severity" />
      <el-table-column align="left" label="案例状态" min-width="180" prop="status" />
      <el-table-column align="left" label="完结时间" min-width="180" prop="dateClosed" />
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
</template>

<style scoped lang="scss">

</style>