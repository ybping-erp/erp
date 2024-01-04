<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="商品SKU:" prop="goodsSku">
          <el-input v-model="formData.goodsSku" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="关联仓库的标识符:" prop="warehouseId">
          <el-input v-model.number="formData.warehouseId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="库存数量:" prop="quantity">
          <el-input v-model.number="formData.quantity" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="预留库存数量:" prop="reservedQuantity">
          <el-input v-model.number="formData.reservedQuantity" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="在途库存数量:" prop="onOrderQuantity">
          <el-input v-model.number="formData.onOrderQuantity" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="库存状态:" prop="stockStatus">
          <el-select v-model="formData.stockStatus" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in stock_statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="最后更新库存时间:" prop="lastStockUpdate">
          <el-date-picker v-model="formData.lastStockUpdate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
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
  createInventory,
  updateInventory,
  findInventory
} from '@/api/inventory'

defineOptions({
    name: 'InventoryForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const stock_statusOptions = ref([])
const formData = ref({
            goodsSku: '',
            warehouseId: 0,
            quantity: 0,
            reservedQuantity: 0,
            onOrderQuantity: 0,
            stockStatus: undefined,
            lastStockUpdate: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findInventory({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reinventory
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    stock_statusOptions.value = await getDictFunc('stock_status')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createInventory(formData.value)
               break
             case 'update':
               res = await updateInventory(formData.value)
               break
             default:
               res = await createInventory(formData.value)
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
