
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="商品SPU:" prop="spuId">
          <el-input v-model.number="formData.spuId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品SKU:" prop="sku">
          <el-input v-model="formData.sku" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品分类:" prop="categoryId">
          <el-input v-model.number="formData.categoryId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="中文名称:" prop="chineseName">
          <el-input v-model="formData.chineseName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品编码:" prop="code">
          <el-input v-model="formData.code" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="英文名称:" prop="englishName">
          <el-input v-model="formData.englishName" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="识别码:" prop="identifierCode">
          <el-input v-model="formData.identifierCode" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="销售方式:" prop="salesMethod">
          <el-input v-model.number="formData.salesMethod" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="来源:" prop="sourceUrls">
          <el-input v-model="formData.sourceUrls" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建者:" prop="createdBy">
          <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="删除者:" prop="deletedBy">
          <el-input v-model.number="formData.deletedBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="更新者:" prop="updatedBy">
          <el-input v-model.number="formData.updatedBy" :clearable="true" placeholder="请输入" />
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
  createGoods,
  updateGoods,
  findGoods
} from '@/api/goods'

defineOptions({
    name: 'GoodsForm'
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
            spuId: 0,
            sku: '',
            categoryId: 0,
            chineseName: '',
            code: '',
            englishName: '',
            identifierCode: '',
            salesMethod: 0,
            status: 0,
            sourceUrls: '',
            createdBy: 0,
            deletedBy: 0,
            updatedBy: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findGoods({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.regoods
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
               res = await createGoods(formData.value)
               break
             case 'update':
               res = await updateGoods(formData.value)
               break
             default:
               res = await createGoods(formData.value)
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
