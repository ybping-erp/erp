<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
        <el-form-item label="仓库" prop="warehouseId">
             <el-select v-model.number="searchInfo.warehouseId" placeholder="请选择" style="width:100%" :clearable="true" >
                <el-option v-for="item in my_wms_warehouses" :key="item.ID" :label="item.name" :value="item.ID" />
              </el-select>
        </el-form-item>
        <el-form-item label="库区" prop="stockStatus">
            <el-select v-model="searchInfo.rackId" clearable placeholder="请选择" @clear="()=>{searchInfo.rackId=undefined}">
              <el-option v-for="(item,key) in wms_zoneOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="商品SKU" prop="goodsSku">
         <el-input v-model="searchInfo.goodsSku" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="库存状态" prop="stockStatus">
            <el-select v-model="searchInfo.stockStatus" clearable placeholder="请选择" @clear="()=>{searchInfo.stockStatus=undefined}">
              <el-option v-for="(item,key) in stock_statusOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <!-- <el-table-column type="selection" width="55" /> -->
        <el-table-column align="left" label="商品SKU" prop="goodsSku" width="120" />
        <el-table-column align="left" label="仓库" prop="warehouseId" width="120" />
        <el-table-column align="left" label="总数量" prop="quantity" width="120" />
        <el-table-column align="left" label="预留" prop="safetyQuantity" width="120" />
        <el-table-column align="left" label="在途" prop="onOrderQuantity" width="120" />
        <el-table-column align="left" label="状态" prop="stockStatus" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.stockStatus,stock_statusOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="操作" min-width="120">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateInventoryFunc(scope.row)">变更</el-button>
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
            <el-form-item label="仓库:"  prop="warehouseId" >
              <el-select v-model.number="formData.warehouseId" placeholder="请选择" :clearable="true" >
                <el-option v-for="item in my_wms_warehouses" :key="item.ID" :label="item.name" :value="item.ID" />
              </el-select>
            </el-form-item>
            <el-form-item label="库存数量:"  prop="quantity" >
              <el-input v-model.number="formData.quantity" :clearable="true" placeholder="请输入库存数量" />
            </el-form-item>
            <el-form-item label="预留库存数量:"  prop="safetyQuantity" >
              <el-input v-model.number="formData.safetyQuantity" :clearable="true" placeholder="请输入预留库存数量" />
            </el-form-item>
            <el-form-item label="在途库存数量:"  prop="onOrderQuantity" >
              <el-input v-model.number="formData.onOrderQuantity" :clearable="true" placeholder="请输入在途库存数量" />
            </el-form-item>
            <el-form-item label="库存状态:"  prop="stockStatus" >
              <el-select v-model="formData.stockStatus" placeholder="请选择库存状态" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in stock_statusOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="最后更新库存时间:"  prop="lastStockUpdate" >
              <el-date-picker v-model="formData.lastStockUpdate" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
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
                <el-descriptions-item label="仓库">
                        {{ formData.warehouseId }}
                </el-descriptions-item>
                <el-descriptions-item label="库存数量">
                        {{ formData.quantity }}
                </el-descriptions-item>
                <el-descriptions-item label="预留库存数量">
                        {{ formData.safetyQuantity }}
                </el-descriptions-item>
                <el-descriptions-item label="在途库存数量">
                        {{ formData.onOrderQuantity }}
                </el-descriptions-item>
                <el-descriptions-item label="库存状态">
                        {{ filterDict(formData.stockStatus,stock_statusOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="最后更新库存时间">
                      {{ formatDate(formData.lastStockUpdate) }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createInventory,
  deleteInventory,
  deleteInventoryByIds,
  updateInventory,
  findInventory,
  getInventoryList
} from '@/api/inventory'

import {
  getWarehouseList 
} from '@/api/warehouse'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'Inventory'
})

// 自动化生成的字典（可能为空）以及字段
const stock_statusOptions = ref([])
const wms_zoneOptions = ref([])
const my_wms_warehouses = ref([])
const formData = ref({
        goodsSku: '',
        warehouseId: undefined,
        quantity: 0,
        safetyQuantity: 0,
        onOrderQuantity: 0,
        stockStatus: undefined,
        lastStockUpdate: new Date(),
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
  const table = await getInventoryList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
  stock_statusOptions.value = await getDictFunc('stock_status')
  wms_zoneOptions.value = await getDictFunc('wms_zone')

  const res = await getWarehouseList({ page: 1, pageSize: 1000})
  if (res.code === 0) {
    my_wms_warehouses.value = res.data.list
  }
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
            deleteInventoryFunc(row)
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
      const res = await deleteInventoryByIds({ ids })
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
const updateInventoryFunc = async(row) => {
    const res = await findInventory({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reinventory
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteInventoryFunc = async (row) => {
    const res = await deleteInventory({ ID: row.ID })
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
  const res = await findInventory({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.reinventory
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
          goodsSku: '',
          warehouseId: 0,
          quantity: 0,
          safetyQuantity: 0,
          onOrderQuantity: 0,
          stockStatus: undefined,
          lastStockUpdate: new Date(),
          }
}


// 打开弹窗
// const openDialog = () => {
//     type.value = 'create'
//     dialogFormVisible.value = true
// }

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        goodsSku: '',
        warehouseId: 0,
        quantity: 0,
        safetyQuantity: 0,
        onOrderQuantity: 0,
        stockStatus: undefined,
        lastStockUpdate: new Date(),
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
