<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="平台名称:" prop="platformName">
          <el-input v-model="formData.platformName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="平台ID:" prop="shopId">
          <el-input v-model="formData.shopId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="店铺名称:" prop="shopName">
          <el-input v-model="formData.shopName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="API Client ID:" prop="clientId">
          <el-input v-model="formData.clientId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="API Client Cert:" prop="clientCert">
          <el-input v-model="formData.clientCert" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建人:" prop="creatorId">
          <el-input v-model.number="formData.creatorId" :clearable="true" placeholder="请输入" />
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
  createShop,
  updateShop,
  findShop
} from '@/api/shop'

defineOptions({
    name: 'ShopForm'
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
            platformName: '',
            shopId: '',
            shopName: '',
            clientId: '',
            clientCert: '',
            creatorId: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findShop({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reshop
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
               res = await createShop(formData.value)
               break
             case 'update':
               res = await updateShop(formData.value)
               break
             default:
               res = await createShop(formData.value)
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
