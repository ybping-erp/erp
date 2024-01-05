<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="电商平台:" prop="platformName">
          <el-input v-model="formData.platformName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="店铺ID:" prop="shopId">
          <el-input v-model="formData.shopId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="店铺名称:" prop="shopName">
          <el-input v-model="formData.shopName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="店铺订单号:" prop="orderId">
          <el-input v-model="formData.orderId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="订单总金额:" prop="totalAmount">
          <el-input-number v-model="formData.totalAmount" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="订单折扣金额:" prop="discount">
          <el-input-number v-model="formData.discount" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="订单税额:" prop="tax">
          <el-input-number v-model="formData.tax" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="状态ID:" prop="statusId">
          <el-select v-model="formData.statusId" placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in order_statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="客户ID:" prop="customerId">
          <el-input v-model="formData.customerId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="客户名称:" prop="customerName">
          <el-input v-model="formData.customerName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="客户电话:" prop="customerTel">
          <el-input v-model="formData.customerTel" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="客户邮箱:" prop="customerEmail">
          <el-input v-model="formData.customerEmail" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="客户备注:" prop="customerRemarks">
          <el-input v-model="formData.customerRemarks" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="支付方式:" prop="paymentMethod">
          <el-input v-model="formData.paymentMethod" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="付款时间:" prop="paymentAt">
          <el-date-picker v-model="formData.paymentAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="运单号:" prop="shippingOrderId">
          <el-input v-model="formData.shippingOrderId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="运费:" prop="shippingCost">
          <el-input-number v-model="formData.shippingCost" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="街道地址:" prop="shippingStreetAddress">
          <el-input v-model="formData.shippingStreetAddress" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="城市:" prop="shippingCity">
          <el-input v-model="formData.shippingCity" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="省/州:" prop="shippingState">
          <el-input v-model="formData.shippingState" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="邮政编码:" prop="shippingPostalCode">
          <el-input v-model="formData.shippingPostalCode" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="国家:" prop="shippingCountry">
          <el-input v-model="formData.shippingCountry" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="发货时间:" prop="shipAt">
          <el-date-picker v-model="formData.shipAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="退款时间:" prop="refundAt">
          <el-date-picker v-model="formData.refundAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
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
  createOrder,
  updateOrder,
  findOrder
} from '@/api/order'

defineOptions({
    name: 'OrderForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const order_statusOptions = ref([])
const formData = ref({
            platformName: '',
            shopId: '',
            shopName: '',
            orderId: '',
            totalAmount: 0,
            discount: 0,
            tax: 0,
            statusId: undefined,
            customerId: '',
            customerName: '',
            customerTel: '',
            customerEmail: '',
            customerRemarks: '',
            paymentMethod: '',
            paymentAt: new Date(),
            shippingOrderId: '',
            shippingCost: 0,
            shippingStreetAddress: '',
            shippingCity: '',
            shippingState: '',
            shippingPostalCode: '',
            shippingCountry: '',
            shipAt: new Date(),
            refundAt: new Date(),
        })
// 验证规则
const rule = reactive({
               statusId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findOrder({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reorder
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    order_statusOptions.value = await getDictFunc('order_status')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createOrder(formData.value)
               break
             case 'update':
               res = await updateOrder(formData.value)
               break
             default:
               res = await createOrder(formData.value)
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
