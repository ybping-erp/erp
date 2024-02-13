
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
      
      <el-form-item label="类别" prop="categoryId">
             <el-cascader
                v-model="searchInfo.categoryId"
                style="width:100%"
                :options="product_categoryOptions"
                :show-all-levels="false"
                :props="{ multiple:false,checkStrictly: true,label:'label',value:'value',disabled:'disabled',emitPath:false}"
                :clearable="true"
              />
        </el-form-item>
           <el-form-item label="销售方式" prop="salesMethod">
            <el-select v-model="searchInfo.salesMethod" clearable placeholder="请选择" @clear="()=>{searchInfo.salesMethod=undefined}">
              <el-option v-for="(item,key) in sales_methodOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="商品SKU" prop="sku">
         <el-input v-model="searchInfo.sku" placeholder="搜索条件" />
        </el-form-item>
           <el-form-item label="商品状态" prop="status">
            <el-select v-model="searchInfo.status" clearable placeholder="请选择" @clear="()=>{searchInfo.status=undefined}">
              <el-option v-for="(item,key) in goods_statusOptions" :key="key" :label="item.label" :value="item.value" />
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
        <el-table-column align="left" label="商品SKU" prop="sku" width="120" />
        <el-table-column align="left" label="商品SPU" prop="spuId" width="120" />
        <el-table-column align="left" label="商品状态" prop="status" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.status,goods_statusOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="类别" prop="categoryId" width="120">
          <template #default="scope"> 
            {{ filterCascaderDict(scope.row.categoryId, product_categoryOptions) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="中文名称" prop="chineseName" width="120" />
        <el-table-column align="left" label="英文名称" prop="englishName" width="120" />
        <el-table-column align="left" label="销售方式" prop="salesMethod" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.salesMethod,sales_methodOptions) }}
            </template>
        </el-table-column>
        
        <el-table-column align="left" label="来源URL列表" prop="sourceUrls" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateGoodsFunc(scope.row)">变更</el-button>
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
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="类别:"  prop="categoryId" >
              <el-cascader
                v-model="formData.categoryId"
                style="width:100%"
                :options="product_categoryOptions"
                :show-all-levels="false"
                :props="{ multiple:false,checkStrictly: true,label:'label',value:'value',disabled:'disabled',emitPath:false}"
                :clearable="true"
              />
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
          <el-select v-model="formData.salesMethod" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in sales_methodOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="商品SKU:" prop="sku">
          <el-input v-model="formData.sku" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="来源URL列表:" prop="sourceUrls">
          <el-input v-model="formData.sourceUrls" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品SPU:" prop="spuId">
          <el-input v-model.number="formData.spuId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="商品状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in goods_statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
                <el-descriptions-item label="商品SPU">
                        {{ formData.spuId }}
                </el-descriptions-item>
                <el-descriptions-item label="商品SKU">
                        {{ formData.sku }}
                </el-descriptions-item>
                <el-descriptions-item label="商品状态">
                        {{ filterDict(formData.status,goods_statusOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="类别">
                  {{ filterCascaderDict(formData.categoryId, product_categoryOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="中文名称">
                        {{ formData.chineseName }}
                </el-descriptions-item>
                <el-descriptions-item label="商品编码">
                        {{ formData.code }}
                </el-descriptions-item>
                <el-descriptions-item label="英文名称">
                        {{ formData.englishName }}
                </el-descriptions-item>
                <el-descriptions-item label="识别码">
                        {{ formData.identifierCode }}
                </el-descriptions-item>
                <el-descriptions-item label="销售方式">
                        {{ filterDict(formData.salesMethod,sales_methodOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="来源URL列表">
                        {{ formData.sourceUrls }}
                </el-descriptions-item>
                
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createGoods,
  deleteGoods,
  deleteGoodsByIds,
  updateGoods,
  findGoods,
  getGoodsList
} from '@/api/goods'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, filterCascaderDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import {
  buildProductCategoryOptions
} from '@/api/category'

defineOptions({
    name: 'Goods'
})

// 自动化生成的字典（可能为空）以及字段
const sales_methodOptions = ref([])
const goods_statusOptions = ref([])
const product_categoryOptions = ref([])
const formData = ref({
        categoryId: 0,
        chineseName: '',
        code: '',
        englishName: '',
        identifierCode: '',
        salesMethod: undefined,
        sku: '',
        sourceUrls: '',
        spuId: 0,
        status: undefined,
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
  const table = await getGoodsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    sales_methodOptions.value = await getDictFunc('sales_method')
    goods_statusOptions.value = await getDictFunc('goods_status')
    product_categoryOptions.value = await buildProductCategoryOptions("Goods")
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
            deleteGoodsFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteGoodsByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateGoodsFunc = async(row) => {
    const res = await findGoods({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.regoods
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteGoodsFunc = async (row) => {
    const res = await deleteGoods({ ID: row.ID })
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
  const res = await findGoods({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.regoods
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          categoryId: 0,
          chineseName: '',
          code: '',
          englishName: '',
          identifierCode: '',
          salesMethod: undefined,
          sku: '',
          sourceUrls: '',
          spuId: 0,
          status: undefined,
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
        categoryId: 0,
        chineseName: '',
        code: '',
        englishName: '',
        identifierCode: '',
        salesMethod: undefined,
        sku: '',
        sourceUrls: '',
        spuId: 0,
        status: undefined,
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>

