<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="仓库ID:" prop="warehouseId">
          <el-input v-model.number="formData.warehouseId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="分区ID:" prop="zoneId">
          <el-select v-model="formData.zoneId" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in wms_zoneOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="货架编号:" prop="rackCode">
          <el-input v-model="formData.rackCode" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="货架状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in rack_statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="拣货权重:" prop="priority">
          <el-input v-model.number="formData.priority" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备注说明:" prop="remarks">
          <el-input v-model="formData.remarks" :clearable="true" placeholder="请输入" />
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
  createRack,
  updateRack,
  findRack
} from '@/api/rack'

defineOptions({
    name: 'RackForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const wms_zoneOptions = ref([])
const rack_statusOptions = ref([])
const formData = ref({
            warehouseId: 0,
            zoneId: undefined,
            rackCode: '',
            status: undefined,
            priority: 0,
            remarks: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findRack({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rerack
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    wms_zoneOptions.value = await getDictFunc('wms_zone')
    rack_statusOptions.value = await getDictFunc('rack_status')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createRack(formData.value)
               break
             case 'update':
               res = await updateRack(formData.value)
               break
             default:
               res = await createRack(formData.value)
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
