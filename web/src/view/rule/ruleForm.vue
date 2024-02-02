<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="动作代码:" prop="actionCode">
          <el-input v-model.number="formData.actionCode" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="动作描述:" prop="actionDesc">
          <el-input v-model="formData.actionDesc" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="领域:" prop="domain">
          <el-input v-model="formData.domain" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="规则名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="规则条件:" prop="ruleStr">
          <el-input v-model="formData.ruleStr" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="规则状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in goods_statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="系统默认规则:" prop="systemDefault">
          <el-input v-model.number="formData.systemDefault" :clearable="true" placeholder="请输入" />
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
  createRule,
  updateRule,
  findRule
} from '@/api/rule'

defineOptions({
    name: 'RuleForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const goods_statusOptions = ref([])
const formData = ref({
            actionCode: 0,
            actionDesc: '',
            domain: '',
            name: '',
            ruleStr: '',
            status: undefined,
            systemDefault: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findRule({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rerule
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    goods_statusOptions.value = await getDictFunc('goods_status')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createRule(formData.value)
               break
             case 'update':
               res = await updateRule(formData.value)
               break
             default:
               res = await createRule(formData.value)
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
