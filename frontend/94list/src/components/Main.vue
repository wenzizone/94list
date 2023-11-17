<template>
  <el-space direction="vertical" fill=true>
    <el-card shadow="always">
      <template #header>
        <div class="card-header">
          <h2>前台解析中心 | 94list-vue-go</h2>
        </div>
      </template>
      <el-form label-width="auto">
        <div>
          <el-form-item label="链接">
            <el-input v-model="panUrl" @blur="handleBlur" />
          </el-form-item>
        </div>
        
        <el-form-item label="密码">
          <el-input :value=passWd v-model="passWd"/>
        </el-form-item>
        <el-form-item label="当前路径">
          <el-input v-model="path" value="/" disabled />
        </el-form-item>
        <el-form-item>
          <el-button type="primary">解析链接</el-button>
          <el-button type="primary">刷新列表</el-button>
          <el-button type="primary" disabled>批量下载</el-button>
          <el-button type="primary">复制当前地址</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  
    <el-card shadow="always">
      <el-table :data="tableData" border style="width: 100%">
        <el-table-column type="selection" width="40" />
        <el-table-column prop="fileName" label="文件名" width="380" />
        <el-table-column prop="mTime" label="修改时间" width="180" />
        <el-table-column prop="fSize" label="文件大小" width="180"/>
      </el-table>
    </el-card>
  </el-space>
</template>

<style>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}
</style>


<script lang="ts" setup>
import { ref } from 'vue'
let panUrl = ref('')
let passWd = ref('')
let path = ref('')

function handleBlur(e){
  //url = e.target.value
  console.log(e.target.value)
  const parseUrl = new URL(e.target.value)
  passWd = parseUrl.searchParams.get('pwd')
  console.log(parseUrl.searchParams.get('pwd'))
  console.log(parseUrl.origin + parseUrl.pathname)
  panUrl = parseUrl.origin + parseUrl.pathname
}
</script>