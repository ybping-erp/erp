<template>
  <el-form ref="vForm" :model="formData" :rules="rules" @keyup.enter="onSubmit">
    <div class="tab-container">
      <el-tabs v-model="crawTab" type="border-card">
        <el-tab-pane name="url-craw-tab" label="单品采集">
          <el-form-item prop="urls" class="required">
            <el-input type="textarea" v-model="formData.urls" placeholder="请填写产品网址" rows="3"></el-input>
          </el-form-item>
          <el-form-item class="button-container">
            <el-button type="primary" icon="search" @click="onSubmit('单品采集')">开始采集</el-button>
            <el-button icon="refresh" @click="onReset">重置</el-button>
          </el-form-item>
        </el-tab-pane>
        <el-tab-pane name="category-craw-tab" label="分类采集">
          <el-form-item prop="category" class="label-right-align">
            <el-input type="textarea" v-model="formData.category" placeholder="请填写店铺内的分类网址" rows="3"></el-input>
          </el-form-item>
          <el-form-item class="button-container">
            <el-button type="primary" icon="search" @click="onSubmit('分类采集')">开始采集</el-button>
            <el-button icon="refresh" @click="onReset">重置</el-button>
          </el-form-item>
        </el-tab-pane>
        <!-- Other tabs... -->
      </el-tabs>
    </div>
  </el-form>
</template>

<script setup>
import {ref, reactive }from 'vue'

defineOptions({
  name: 'ProductSpider'
})

// Set the default tab name
const crawTab = ref('url-craw-tab'); 
const formData = ref({
  urls: "",
  category: "",
})

// 验证规则
const rule = reactive({
})

// 重置采集
const onReset = () => {
  formData.value = {}
}

// 开始采集
const onSubmit = (tab) => {
  alert(tab +"开发中");
}
</script>