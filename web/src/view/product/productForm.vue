<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="产品所属的类别标识符:" prop="categoryId">
          <el-input v-model.number="formData.categoryId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="产品描述:" prop="description">
          <el-input v-model="formData.description" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="产品名称:" prop="productName">
          <el-input v-model="formData.productName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="产品的唯一标识符:" prop="sku">
          <el-input v-model="formData.sku" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="产品摘要:" prop="summary">
          <el-input v-model="formData.summary" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="产品的单价:" prop="unitPrice">
          <el-input-number v-model="formData.unitPrice" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createProduct,
  updateProduct,
  findProduct
} from '@/api/product'

defineOptions({
    name: 'ProductForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            categoryId: 0,
            description: '',
            productName: '',
            sku: '',
            summary: '',
            unitPrice: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findProduct({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reproduct
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createProduct(formData.value)
               break
             case 'update':
               res = await updateProduct(formData.value)
               break
             default:
               res = await createProduct(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
