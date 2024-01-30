
<template>
  <el-form :model="formData" ref="vForm" :rules="rules" label-position="left" size="medium" @submit.prevent>
    <div class="tab-container">
      <el-tabs v-model="tab60554ActiveTab" type="border-card">
        <el-tab-pane name="url-craw-tab" label="单品采集">
          <el-form-item prop="urls" class="required">
            <el-input type="textarea" v-model="formData.urls" placeholder="请填写产品网址" rows="3">
            </el-input>
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane name="category-craw-tab" label="分类采集">
          <el-form-item prop="category" class="label-right-align">
            <el-input type="textarea" v-model="formData.category" placeholder="请填写店铺内的分类网址" rows="3"></el-input>
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane name="tab-pane-73949" label="关键字采集">
          等待需求说明
        </el-tab-pane>
      </el-tabs>
    </div>
  </el-form>
  
</template>

<script setup>
  import {ref, reactive }from 'vue'

  defineOptions({
    name: 'ProductSpider'
  })

const formData = ref({
          urls: "",
          category: "",
})

// 验证规则
const rule = reactive({
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
  
  div.tab-container {}
  
  .label-left-align :deep(.el-form-item__label) {
    text-align: left;
  }
  
  .label-center-align :deep(.el-form-item__label) {
    text-align: center;
  }
  
  .label-right-align :deep(.el-form-item__label) {
    text-align: right;
  }
  
  .custom-label {}
  
  .static-content-item {
    min-height: 20px;
    display: flex;
    align-items: center;

    :deep(.el-divider--horizontal) {
      margin: 0;
    }
  }
  
</style>