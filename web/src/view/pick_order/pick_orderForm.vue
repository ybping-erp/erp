<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="商品:" prop="goodsName">
          <el-input v-model="formData.goodsName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="SKU:" prop="goodsSku">
          <el-input v-model="formData.goodsSku" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="订单:" prop="orderId">
          <el-input v-model.number="formData.orderId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="拣货数量:" prop="quantity">
          <el-input v-model.number="formData.quantity" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="货架:" prop="rackId">
          <el-input v-model.number="formData.rackId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="仓库:" prop="warehouseId">
          <el-input v-model.number="formData.warehouseId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="库区:" prop="zoneId">
          <el-input v-model.number="formData.zoneId" :clearable="true" placeholder="请输入" />
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
  createPickOrder,
  updatePickOrder,
  findPickOrder
} from '@/api/pick_order'

defineOptions({
    name: 'PickOrderForm'
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
            goodsName: '',
            goodsSku: '',
            orderId: 0,
            quantity: 0,
            rackId: 0,
            warehouseId: 0,
            zoneId: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findPickOrder({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.repickOrder
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
               res = await createPickOrder(formData.value)
               break
             case 'update':
               res = await updatePickOrder(formData.value)
               break
             default:
               res = await createPickOrder(formData.value)
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
