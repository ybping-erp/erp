<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="仓库名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="负责人:" prop="owner">
          <el-input v-model="formData.owner" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="联系电话:" prop="telephone">
          <el-input v-model="formData.telephone" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="仓库街道地址:" prop="streetAddress">
          <el-input v-model="formData.streetAddress" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="城市:" prop="city">
          <el-input v-model="formData.city" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="州/省:" prop="stateProvince">
          <el-input v-model="formData.stateProvince" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="邮政编码:" prop="postalCode">
          <el-input v-model="formData.postalCode" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="国家:" prop="country">
          <el-input v-model="formData.country" :clearable="true" placeholder="请输入" />
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
  createWarehouse,
  updateWarehouse,
  findWarehouse
} from '@/api/warehouse'

defineOptions({
    name: 'WarehouseForm'
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
            name: '',
            owner: '',
            telephone: '',
            streetAddress: '',
            city: '',
            stateProvince: '',
            postalCode: '',
            country: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWarehouse({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewarehouse
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
               res = await createWarehouse(formData.value)
               break
             case 'update':
               res = await updateWarehouse(formData.value)
               break
             default:
               res = await createWarehouse(formData.value)
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
