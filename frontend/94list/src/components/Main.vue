<template>
  <el-space direction="vertical">
    <el-card shadow="always">
      <template #header>
        <div class="card-header">
          <h2>前台解析中心 | 94list-vue-go</h2>
        </div>
      </template>
      <el-form label-width="auto">
        <div>
          <el-form-item label="链接">
            <el-input />
          </el-form-item>
        </div>
        
        <el-form-item label="密码">
          <el-input />
        </el-form-item>
        <el-form-item label="当前路径">
          <el-input value="/" disabled />
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
        <el-table-column prop="date" label="Date" width="180" />
        <el-table-column prop="name" label="Name" width="180" />
        <el-table-column prop="address" label="Address" />
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