<script setup>
// import { createCase } from '@/api/case'
import { deleteCase } from '@/api/case'
import { ref } from 'vue'

const taskData = ref([])
</script>

<template>
  <div class="case">
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus">新增任务</el-button>
      </div>
      <el-table
        :data="taskData"
        row-key="caseID"
        style="width: 100%"
      >
        <el-table-column label="任务ID" min-width="180" prop="taskID" />
        <el-table-column align="left" label="发起人UUID" min-width="180" prop="uuid" />
        <el-table-column align="left" label="视频源" min-width="180" prop="videoSource" />
        <el-table-column align="left" label="创建时间" min-width="180" prop="creationDate" />
        <el-table-column align="left" label="创建时间" min-width="180" prop="resolution" />
        <el-table-column align="left" label="任务容器算法ID" min-width="180" prop="algorithmId" />
        <el-table-column align="left" label="任务粒度" min-width="180" prop="intensity" />
        <el-table-column align="left" lebal="任务状态" min-width="180" prop="status" />
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
  </div>
</template>

<style scoped lang="scss">

</style>