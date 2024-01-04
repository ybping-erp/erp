<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="商品SKU:" prop="goodsSku">
          <el-input v-model="formData.goodsSku" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="报关中文名:" prop="declarationChineseName">
          <el-input v-model="formData.declarationChineseName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="报关英文名:" prop="declarationEnglishName">
          <el-input v-model="formData.declarationEnglishName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="报关价格币种:" prop="declarationPriceCurrency">
          <el-input v-model="formData.declarationPriceCurrency" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="报关重量单位:" prop="declarationWeightUnit">
          <el-input v-model="formData.declarationWeightUnit" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="材质:" prop="material">
          <el-input v-model="formData.material" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用途:" prop="purpose">
          <el-input v-model="formData.purpose" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="海关编码:" prop="customsCode">
          <el-input v-model="formData.customsCode" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="危险运输品:" prop="specialGoods">
          <el-select v-model="formData.specialGoods" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in special_goodsOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="商品净重 (g):" prop="netWeight">
          <el-input-number v-model="formData.netWeight" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="允许称重误差 (g):" prop="weightErrorTolerance">
          <el-input-number v-model="formData.weightErrorTolerance" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="商品尺寸 - 长 (cm):" prop="lengthCm">
          <el-input-number v-model="formData.lengthCm" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="商品尺寸 - 宽 (cm):" prop="widthCm">
          <el-input-number v-model="formData.widthCm" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="商品尺寸 - 高 (cm):" prop="heightCm">
          <el-input-number v-model="formData.heightCm" :precision="2" :clearable="true"></el-input-number>
       </el-form-item>
        <el-form-item label="包装信息:" prop="packaging">
          <el-input v-model="formData.packaging" :clearable="true" placeholder="请输入" />
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
  createLogisticsPackaging,
  updateLogisticsPackaging,
  findLogisticsPackaging
} from '@/api/logistics_packaging'

defineOptions({
    name: 'LogisticsPackagingForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const special_goodsOptions = ref([])
const formData = ref({
            goodsSku: '',
            declarationChineseName: '',
            declarationEnglishName: '',
            declarationPriceCurrency: '',
            declarationWeightUnit: '',
            material: '',
            purpose: '',
            customsCode: '',
            specialGoods: undefined,
            netWeight: 0,
            weightErrorTolerance: 0,
            lengthCm: 0,
            widthCm: 0,
            heightCm: 0,
            packaging: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findLogisticsPackaging({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.relogisticsPackaging
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    special_goodsOptions.value = await getDictFunc('special_goods')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createLogisticsPackaging(formData.value)
               break
             case 'update':
               res = await updateLogisticsPackaging(formData.value)
               break
             default:
               res = await createLogisticsPackaging(formData.value)
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
