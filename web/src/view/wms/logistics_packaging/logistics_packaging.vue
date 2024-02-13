<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
        <el-form-item label="商品SKU" prop="goodsSku">
         <el-input v-model="searchInfo.goodsSku" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="危险运输品" prop="specialGoods">
            <el-select v-model="searchInfo.specialGoods" clearable placeholder="请选择" @clear="()=>{searchInfo.specialGoods=undefined}">
              <el-option v-for="(item,key) in special_goodsOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" :disabled="!multipleSelection.length" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="商品SKU" prop="goodsSku" width="120" />
        <el-table-column align="left" label="报关中文名" prop="declarationChineseName" width="120" />
        <el-table-column align="left" label="报关英文名" prop="declarationEnglishName" width="120" />
        <el-table-column align="left" label="报关价格币种" prop="declarationPriceCurrency" width="120" />
        <el-table-column align="left" label="报关重量单位" prop="declarationWeightUnit" width="120" />
        <el-table-column align="left" label="材质" prop="material" width="120" />
        <el-table-column align="left" label="用途" prop="purpose" width="120" />
        <el-table-column align="left" label="海关编码" prop="customsCode" width="120" />
        <el-table-column align="left" label="危险运输品" prop="specialGoods" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.specialGoods,special_goodsOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="商品净重 (g)" prop="netWeight" width="120" />
        <el-table-column align="left" label="允许称重误差 (g)" prop="weightErrorTolerance" width="120" />
        <el-table-column align="left" label="商品尺寸 - 长 (cm)" prop="lengthCm" width="120" />
        <el-table-column align="left" label="商品尺寸 - 宽 (cm)" prop="widthCm" width="120" />
        <el-table-column align="left" label="商品尺寸 - 高 (cm)" prop="heightCm" width="120" />
        <el-table-column align="left" label="包装信息" prop="packaging" width="120" />
        <el-table-column align="left" label="操作" min-width="120">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateLogisticsPackagingFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'添加':'修改'" destroy-on-close>
      <el-scrollbar height="500px">
          <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="商品SKU:"  prop="goodsSku" >
              <el-input v-model="formData.goodsSku" :clearable="true"  placeholder="请输入商品SKU" />
            </el-form-item>
            <el-form-item label="报关中文名:"  prop="declarationChineseName" >
              <el-input v-model="formData.declarationChineseName" :clearable="true"  placeholder="请输入报关中文名" />
            </el-form-item>
            <el-form-item label="报关英文名:"  prop="declarationEnglishName" >
              <el-input v-model="formData.declarationEnglishName" :clearable="true"  placeholder="请输入报关英文名" />
            </el-form-item>
            <el-form-item label="报关价格币种:"  prop="declarationPriceCurrency" >
              <el-input v-model="formData.declarationPriceCurrency" :clearable="true"  placeholder="请输入报关价格币种" />
            </el-form-item>
            <el-form-item label="报关重量单位:"  prop="declarationWeightUnit" >
              <el-input v-model="formData.declarationWeightUnit" :clearable="true"  placeholder="请输入报关重量单位" />
            </el-form-item>
            <el-form-item label="材质:"  prop="material" >
              <el-input v-model="formData.material" :clearable="true"  placeholder="请输入材质" />
            </el-form-item>
            <el-form-item label="用途:"  prop="purpose" >
              <el-input v-model="formData.purpose" :clearable="true"  placeholder="请输入用途" />
            </el-form-item>
            <el-form-item label="海关编码:"  prop="customsCode" >
              <el-input v-model="formData.customsCode" :clearable="true"  placeholder="请输入海关编码" />
            </el-form-item>
            <el-form-item label="危险运输品:"  prop="specialGoods" >
              <el-select v-model="formData.specialGoods" placeholder="请选择危险运输品" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in special_goodsOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="商品净重 (g):"  prop="netWeight" >
              <el-input-number v-model="formData.netWeight"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="允许称重误差 (g):"  prop="weightErrorTolerance" >
              <el-input-number v-model="formData.weightErrorTolerance"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="商品尺寸 - 长 (cm):"  prop="lengthCm" >
              <el-input-number v-model="formData.lengthCm"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="商品尺寸 - 宽 (cm):"  prop="widthCm" >
              <el-input-number v-model="formData.widthCm"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="商品尺寸 - 高 (cm):"  prop="heightCm" >
              <el-input-number v-model="formData.heightCm"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="包装信息:"  prop="packaging" >
              <el-input v-model="formData.packaging" :clearable="true"  placeholder="请输入包装信息" />
            </el-form-item>
          </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="详情" destroy-on-close>
      <el-scrollbar height="550px">
        <el-descriptions column="1" border>
                <el-descriptions-item label="商品SKU">
                        {{ formData.goodsSku }}
                </el-descriptions-item>
                <el-descriptions-item label="报关中文名">
                        {{ formData.declarationChineseName }}
                </el-descriptions-item>
                <el-descriptions-item label="报关英文名">
                        {{ formData.declarationEnglishName }}
                </el-descriptions-item>
                <el-descriptions-item label="报关价格币种">
                        {{ formData.declarationPriceCurrency }}
                </el-descriptions-item>
                <el-descriptions-item label="报关重量单位">
                        {{ formData.declarationWeightUnit }}
                </el-descriptions-item>
                <el-descriptions-item label="材质">
                        {{ formData.material }}
                </el-descriptions-item>
                <el-descriptions-item label="用途">
                        {{ formData.purpose }}
                </el-descriptions-item>
                <el-descriptions-item label="海关编码">
                        {{ formData.customsCode }}
                </el-descriptions-item>
                <el-descriptions-item label="危险运输品">
                        {{ filterDict(formData.specialGoods,special_goodsOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="商品净重 (g)">
                        {{ formData.netWeight }}
                </el-descriptions-item>
                <el-descriptions-item label="允许称重误差 (g)">
                        {{ formData.weightErrorTolerance }}
                </el-descriptions-item>
                <el-descriptions-item label="商品尺寸 - 长 (cm)">
                        {{ formData.lengthCm }}
                </el-descriptions-item>
                <el-descriptions-item label="商品尺寸 - 宽 (cm)">
                        {{ formData.widthCm }}
                </el-descriptions-item>
                <el-descriptions-item label="商品尺寸 - 高 (cm)">
                        {{ formData.heightCm }}
                </el-descriptions-item>
                <el-descriptions-item label="包装信息">
                        {{ formData.packaging }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createLogisticsPackaging,
  deleteLogisticsPackaging,
  deleteLogisticsPackagingByIds,
  updateLogisticsPackaging,
  findLogisticsPackaging,
  getLogisticsPackagingList
} from '@/api/logistics_packaging'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'LogisticsPackaging'
})

// 自动化生成的字典（可能为空）以及字段
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

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getLogisticsPackagingList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    special_goodsOptions.value = await getDictFunc('special_goods')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteLogisticsPackagingFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteLogisticsPackagingByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateLogisticsPackagingFunc = async(row) => {
    const res = await findLogisticsPackaging({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.relogisticsPackaging
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteLogisticsPackagingFunc = async (row) => {
    const res = await deleteLogisticsPackaging({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)


// 详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findLogisticsPackaging({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.relogisticsPackaging
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
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
          }
}


// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
