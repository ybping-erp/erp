
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="公司地址:" prop="companyAddress">
          <el-input v-model="formData.companyAddress" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="公司网址:" prop="companyWebsite">
          <el-input v-model="formData.companyWebsite" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="联系人:" prop="contactName">
          <el-input v-model="formData.contactName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="联系电话:" prop="contactTelephone">
          <el-input v-model="formData.contactTelephone" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="银行名称:" prop="bankName">
          <el-input v-model="formData.bankName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="银行卡号:" prop="bankAccountNumber">
          <el-input v-model="formData.bankAccountNumber" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="收款人:" prop="bankAccountName">
          <el-input v-model="formData.bankAccountName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="付款方式:" prop="paymentMethod">
          <el-select v-model="formData.paymentMethod" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in payment_methodOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="结算方式:" prop="settlementMethod">
          <el-select v-model="formData.settlementMethod" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in settlement_methodOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="备注信息:" prop="remarks">
          <el-input v-model="formData.remarks" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建者:" prop="createdBy">
          <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="更新者:" prop="updatedBy">
          <el-input v-model.number="formData.updatedBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="删除者:" prop="deletedBy">
          <el-input v-model.number="formData.deletedBy" :clearable="true" placeholder="请输入" />
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
  createSupplier,
  updateSupplier,
  findSupplier
} from '@/api/supplier'

defineOptions({
    name: 'SupplierForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const payment_methodOptions = ref([])
const settlement_methodOptions = ref([])
const formData = ref({
            name: '',
            companyAddress: '',
            companyWebsite: '',
            contactName: '',
            contactTelephone: '',
            bankName: '',
            bankAccountNumber: '',
            bankAccountName: '',
            paymentMethod: undefined,
            settlementMethod: undefined,
            remarks: '',
            createdBy: 0,
            updatedBy: 0,
            deletedBy: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSupplier({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.resupplier
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    payment_methodOptions.value = await getDictFunc('payment_method')
    settlement_methodOptions.value = await getDictFunc('settlement_method')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createSupplier(formData.value)
               break
             case 'update':
               res = await updateSupplier(formData.value)
               break
             default:
               res = await createSupplier(formData.value)
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
