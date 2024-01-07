<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
        <el-form-item label="商品SKU" prop="sku">
         <el-input v-model="searchInfo.sku" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="销售方式" prop="salesMethod">
           <el-select v-model="searchInfo.stockStatus" clearable placeholder="请选择" @clear="()=>{searchInfo.salesMethod=undefined}">
              <el-option v-for="(item,key) in sales_methodOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
        </el-form-item>
        <el-form-item label="商品状态" prop="status">
         <el-select v-model="searchInfo.status" clearable placeholder="请选择" @clear="()=>{searchInfo.salesMethod=undefined}">
              <el-option v-for="(item,key) in goods_statusOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
        </el-form-item>
        <el-form-item label="中文名称" prop="chineseName">
         <el-input v-model="searchInfo.chineseName" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="英文名称" prop="englishName">
         <el-input v-model="searchInfo.englishName" placeholder="搜索条件" />

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
        <el-table-column align="left" label="商品编码" prop="code" width="120" />
        <el-table-column align="left" label="商品SPU" prop="spu" width="120" />
        <el-table-column align="left" label="属性信息" prop="spuAttributes" width="120" />
        <el-table-column align="left" label="商品SKU" prop="sku" width="120" />
        <el-table-column align="left" label="组合商品" prop="childrenSku" width="120" />
        <el-table-column align="left" label="需要加工" prop="needAdditionalProcess" width="120" />
        <el-table-column align="left" label="销售方式" prop="salesMethod" width="120" />
        <el-table-column align="left" label="商品状态" prop="status" width="120">
          <template #default="scope">
            {{ filterDict(scope.row.status,goods_statusOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="商品分类" prop="categoryId" width="120" />
        <el-table-column align="left" label="中文名称" prop="chineseName" width="120" />
        <el-table-column align="left" label="英文名称" prop="englishName" width="120" />
        <el-table-column align="left" label="识别码" prop="identifierCode" width="120" />
        <el-table-column align="left" label="来源URL列表" prop="sourceUrls" width="120" />
        <el-table-column align="left" label="操作" min-width="120">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
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
          <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="商品编码:"  prop="code" >
              <el-input v-model="formData.code" :clearable="true"  placeholder="请输入商品编码" />
            </el-form-item>
            <el-form-item label="商品SPU:"  prop="spu" >
              <el-input v-model="formData.spu" :clearable="true"  placeholder="请输入商品SPU" />
            </el-form-item>
            <el-form-item label="属性信息:"  prop="spuAttributes" >
              <el-input v-model="formData.spuAttributes" :clearable="true"  placeholder="请输入属性信息" />
            </el-form-item>
            <el-form-item label="商品SKU:"  prop="sku" >
              <el-input v-model="formData.sku" :clearable="true"  placeholder="请输入商品SKU" />
            </el-form-item>
            <el-form-item label="组合商品:"  prop="childrenSku" >
              <el-input v-model="formData.childrenSku" :clearable="true"  placeholder="请输入组合商品" />
            </el-form-item>
            <el-form-item label="需要加工:"  prop="needAdditionalProcess" >
              <el-input v-model.number="formData.needAdditionalProcess" :clearable="true" placeholder="请输入需要加工" />
            </el-form-item>
            <el-form-item label="销售方式:"  prop="salesMethod" >
              <el-select v-model="formData.salesMethod" placeholder="请选择销售方式" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in sales_methodOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="商品状态:"  prop="status" >
              <el-select v-model="formData.status" placeholder="请选择库存状态" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in goods_statusOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="商品分类:"  prop="categoryId" >
              <el-cascader
                v-model="formData.categoryId"
                style="width:100%"
                :options="goods_categoryOptions"
                :show-all-levels="false"
                :props="{ multiple:false,checkStrictly: true,label:'categoryName',value:'categoryId',disabled:'disabled',emitPath:false}"
                :clearable="true"
              />
            </el-form-item>
            <el-form-item label="中文名称:"  prop="chineseName" >
              <el-input v-model="formData.chineseName" :clearable="true"  placeholder="请输入中文名称" />
            </el-form-item>
            <el-form-item label="英文名称:"  prop="englishName" >
              <el-input v-model="formData.englishName" :clearable="true"  placeholder="请输入英文名称" />
            </el-form-item>
            <el-form-item label="识别码:"  prop="identifierCode" >
              <el-input v-model="formData.identifierCode" :clearable="true"  placeholder="请输入识别码" />
            </el-form-item>
            <el-form-item label="来源URL列表:"  prop="sourceUrls" >
              <el-input v-model="formData.sourceUrls" :clearable="true"  placeholder="请输入来源URL列表" />
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

    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="查看详情" destroy-on-close>
      <el-scrollbar height="550px">
        <el-descriptions column="1" border>
                <el-descriptions-item label="商品编码">
                        {{ formData.code }}
                </el-descriptions-item>
                <el-descriptions-item label="商品SPU">
                        {{ formData.spu }}
                </el-descriptions-item>
                <el-descriptions-item label="属性信息">
                        {{ formData.spuAttributes }}
                </el-descriptions-item>
                <el-descriptions-item label="商品SKU">
                        {{ formData.sku }}
                </el-descriptions-item>
                <el-descriptions-item label="组合商品">
                        {{ formData.childrenSku }}
                </el-descriptions-item>
                <el-descriptions-item label="需要加工">
                        {{ formData.needAdditionalProcess }}
                </el-descriptions-item>
                <el-descriptions-item label="销售方式">
                        {{ filterDict(formData.salesMethod,sales_methodOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="商品状态">
                        {{ filterDict(formData.status,goods_statusOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="商品分类">
                        {{ formData.categoryId }}
                </el-descriptions-item>
                <el-descriptions-item label="中文名称">
                        {{ formData.chineseName }}
                </el-descriptions-item>
                <el-descriptions-item label="英文名称">
                        {{ formData.englishName }}
                </el-descriptions-item>
                <el-descriptions-item label="识别码">
                        {{ formData.identifierCode }}
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

import {
  getCategoryList
} from '@/api/category'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'Goods'
})

// 自动化生成的字典（可能为空）以及字段
const sales_methodOptions = ref([])
const goods_statusOptions = ref([])
const formData = ref({
        code: '',
        spu: '',
        spuAttributes: '',
        sku: '',
        childrenSku: '',
        needAdditionalProcess: 0,
        categoryId: 0,
        chineseName: '',
        englishName: '',
        identifierCode: '',
        sourceUrls: '',
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
  const table = await getGoodsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const goods_categoryOptions = ref([])
const setGoodsCategoryOptions = (GoodsCategoryData, optionsData) => {
  GoodsCategoryData && GoodsCategoryData.forEach(item => {
    if (item.children && item.children.length) {
      const option = {
        categoryId: item.ID,
        categoryName: item.name,
        children: []
      }
      setGoodsCategoryOptions(item.children, option.children)
      optionsData.push(option)
    } else {
      const option = {
        categoryId: item.ID,
        categoryName: item.name,
      }
      optionsData.push(option)
    }
  })
}
// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
  sales_methodOptions.value = await getDictFunc('sales_method')
  goods_statusOptions.value = await getDictFunc('goods_status')

  // 构建商品类别字典
  const res = await getCategoryList({ page: 1, pageSize: 999, domain: "Goods"})
  goods_categoryOptions.value = []
  setGoodsCategoryOptions(res.data.list, goods_statusOptions.value)
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
      const res = await deleteGoodsByIds({ ids })
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


// 查看详情控制标记
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
          code: '',
          spu: '',
          spuAttributes: '',
          sku: '',
          childrenSku: '',
          needAdditionalProcess: 0,
          categoryId: 0,
          chineseName: '',
          englishName: '',
          identifierCode: '',
          sourceUrls: '',
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
        code: '',
        spu: '',
        spuAttributes: '',
        sku: '',
        childrenSku: '',
        needAdditionalProcess: 0,
        categoryId: 0,
        chineseName: '',
        englishName: '',
        identifierCode: '',
        sourceUrls: '',
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
