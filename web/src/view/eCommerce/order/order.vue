
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <!-- <el-form-item label="创建日期" prop="createdAt">
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
      </el-form-item> -->
      <el-form-item label="订单号" prop="orderId">
         <el-input v-model="searchInfo.orderId" placeholder="搜索条件" />
      </el-form-item>

        <el-form-item label="电商平台" prop="platformName">
         <el-input v-model="searchInfo.platformName" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="店铺ID" prop="shopId">
         <el-input v-model="searchInfo.shopId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="店铺名称" prop="shopName">
         <el-input v-model="searchInfo.shopName" placeholder="搜索条件" />

        </el-form-item>
           <el-form-item label="状态" prop="statusId">
            <el-select v-model="searchInfo.statusId" clearable placeholder="请选择" @clear="()=>{searchInfo.statusId=undefined}">
              <el-option v-for="(item,key) in order_statusOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
        <el-form-item label="运单号" prop="shippingOrderId">
         <el-input v-model="searchInfo.shippingOrderId" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="国家" prop="shippingCountry">
         <el-input v-model="searchInfo.shippingCountry" placeholder="搜索条件" />

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
        <el-table-column align="left" label="电商平台" prop="platformName" width="120" />
        <el-table-column align="left" label="店铺ID" prop="shopId" width="120" />
        <el-table-column align="left" label="店铺名称" prop="shopName" width="120" />
        <el-table-column align="left" label="订单号" prop="orderId" width="120" />
        <el-table-column align="left" label="总金额" prop="totalAmount" width="120" />
        <el-table-column align="left" label="折扣金额" prop="discount" width="120" />
        <el-table-column align="left" label="税额" prop="tax" width="120" />
        <el-table-column align="left" label="状态" prop="statusId" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.statusId,order_statusOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="客户ID" prop="customerId" width="120" />
        <el-table-column align="left" label="客户名称" prop="customerName" width="120" />
        <el-table-column align="left" label="客户电话" prop="customerTel" width="120" />
        <el-table-column align="left" label="客户邮箱" prop="customerEmail" width="120" />
        <el-table-column align="left" label="客户备注" prop="customerRemarks" width="120" />
        <el-table-column align="left" label="支付方式" prop="paymentMethod" width="120" />
         <el-table-column align="left" label="付款时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.paymentAt) }}</template>
         </el-table-column>
        <el-table-column align="left" label="运单号" prop="shippingOrderId" width="120" />
        <el-table-column align="left" label="订单运费" prop="shippingCost" width="120" />
        <el-table-column align="left" label="街道地址" prop="shippingStreetAddress" width="120" />
        <el-table-column align="left" label="城市" prop="shippingCity" width="120" />
        <el-table-column align="left" label="省/州" prop="shippingState" width="120" />
        <el-table-column align="left" label="邮政编码" prop="shippingPostalCode" width="120" />
        <el-table-column align="left" label="国家" prop="shippingCountry" width="120" />
         <el-table-column align="left" label="发货时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.shipAt) }}</template>
         </el-table-column>
         <el-table-column align="left" label="退款时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.refundAt) }}</template>
         </el-table-column>
        <el-table-column align="left" label="操作" min-width="120">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateOrderFunc(scope.row)">变更</el-button>
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
            <el-form-item label="电商平台:"  prop="platformName" >
              <el-input v-model="formData.platformName" :clearable="true"  placeholder="请输入电商平台" />
            </el-form-item>
            <el-form-item label="店铺ID:"  prop="shopId" >
              <el-input v-model="formData.shopId" :clearable="true"  placeholder="请输入店铺ID" />
            </el-form-item>
            <el-form-item label="店铺名称:"  prop="shopName" >
              <el-input v-model="formData.shopName" :clearable="true"  placeholder="请输入店铺名称" />
            </el-form-item>
            <el-form-item label="订单号:"  prop="orderId" >
              <el-input v-model="formData.orderId" :clearable="true"  placeholder="请输入订单号" />
            </el-form-item>
            <el-form-item label="总金额:"  prop="totalAmount" >
              <el-input-number v-model="formData.totalAmount"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="折扣金额:"  prop="discount" >
              <el-input-number v-model="formData.discount"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="税额:"  prop="tax" >
              <el-input-number v-model="formData.tax"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="状态:"  prop="statusId" >
              <el-select v-model="formData.statusId" placeholder="请选择状态" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in order_statusOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="客户ID:"  prop="customerId" >
              <el-input v-model="formData.customerId" :clearable="true"  placeholder="请输入客户ID" />
            </el-form-item>
            <el-form-item label="客户名称:"  prop="customerName" >
              <el-input v-model="formData.customerName" :clearable="true"  placeholder="请输入客户名称" />
            </el-form-item>
            <el-form-item label="客户电话:"  prop="customerTel" >
              <el-input v-model="formData.customerTel" :clearable="true"  placeholder="请输入客户电话" />
            </el-form-item>
            <el-form-item label="客户邮箱:"  prop="customerEmail" >
              <el-input v-model="formData.customerEmail" :clearable="true"  placeholder="请输入客户邮箱" />
            </el-form-item>
            <el-form-item label="客户备注:"  prop="customerRemarks" >
              <el-input v-model="formData.customerRemarks" :clearable="true"  placeholder="请输入客户备注" />
            </el-form-item>
            <el-form-item label="支付方式:"  prop="paymentMethod" >
              <el-input v-model="formData.paymentMethod" :clearable="true"  placeholder="请输入支付方式" />
            </el-form-item>
            <el-form-item label="付款时间:"  prop="paymentAt" >
              <el-date-picker v-model="formData.paymentAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="运单号:"  prop="shippingOrderId" >
              <el-input v-model="formData.shippingOrderId" :clearable="true"  placeholder="请输入运单号" />
            </el-form-item>
            <el-form-item label="订单运费:"  prop="shippingCost" >
              <el-input-number v-model="formData.shippingCost"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="街道地址:"  prop="shippingStreetAddress" >
              <el-input v-model="formData.shippingStreetAddress" :clearable="true"  placeholder="请输入街道地址" />
            </el-form-item>
            <el-form-item label="城市:"  prop="shippingCity" >
              <el-input v-model="formData.shippingCity" :clearable="true"  placeholder="请输入城市" />
            </el-form-item>
            <el-form-item label="省/州:"  prop="shippingState" >
              <el-input v-model="formData.shippingState" :clearable="true"  placeholder="请输入省/州" />
            </el-form-item>
            <el-form-item label="邮政编码:"  prop="shippingPostalCode" >
              <el-input v-model="formData.shippingPostalCode" :clearable="true"  placeholder="请输入邮政编码" />
            </el-form-item>
            <el-form-item label="国家:"  prop="shippingCountry" >
              <el-input v-model="formData.shippingCountry" :clearable="true"  placeholder="请输入国家" />
            </el-form-item>
            <el-form-item label="发货时间:"  prop="shipAt" >
              <el-date-picker v-model="formData.shipAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="退款时间:"  prop="refundAt" >
              <el-date-picker v-model="formData.refundAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
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
                <el-descriptions-item label="电商平台">
                        {{ formData.platformName }}
                </el-descriptions-item>
                <el-descriptions-item label="店铺ID">
                        {{ formData.shopId }}
                </el-descriptions-item>
                <el-descriptions-item label="店铺名称">
                        {{ formData.shopName }}
                </el-descriptions-item>
                <el-descriptions-item label="订单号">
                        {{ formData.orderId }}
                </el-descriptions-item>
                <el-descriptions-item label="总金额">
                        {{ formData.totalAmount }}
                </el-descriptions-item>
                <el-descriptions-item label="折扣金额">
                        {{ formData.discount }}
                </el-descriptions-item>
                <el-descriptions-item label="税额">
                        {{ formData.tax }}
                </el-descriptions-item>
                <el-descriptions-item label="状态">
                        {{ filterDict(formData.statusId,order_statusOptions) }}
                </el-descriptions-item>
                <el-descriptions-item label="客户ID">
                        {{ formData.customerId }}
                </el-descriptions-item>
                <el-descriptions-item label="客户名称">
                        {{ formData.customerName }}
                </el-descriptions-item>
                <el-descriptions-item label="客户电话">
                        {{ formData.customerTel }}
                </el-descriptions-item>
                <el-descriptions-item label="客户邮箱">
                        {{ formData.customerEmail }}
                </el-descriptions-item>
                <el-descriptions-item label="客户备注">
                        {{ formData.customerRemarks }}
                </el-descriptions-item>
                <el-descriptions-item label="支付方式">
                        {{ formData.paymentMethod }}
                </el-descriptions-item>
                <el-descriptions-item label="付款时间">
                      {{ formatDate(formData.paymentAt) }}
                </el-descriptions-item>
                <el-descriptions-item label="运单号">
                        {{ formData.shippingOrderId }}
                </el-descriptions-item>
                <el-descriptions-item label="订单运费">
                        {{ formData.shippingCost }}
                </el-descriptions-item>
                <el-descriptions-item label="街道地址">
                        {{ formData.shippingStreetAddress }}
                </el-descriptions-item>
                <el-descriptions-item label="城市">
                        {{ formData.shippingCity }}
                </el-descriptions-item>
                <el-descriptions-item label="省/州">
                        {{ formData.shippingState }}
                </el-descriptions-item>
                <el-descriptions-item label="邮政编码">
                        {{ formData.shippingPostalCode }}
                </el-descriptions-item>
                <el-descriptions-item label="国家">
                        {{ formData.shippingCountry }}
                </el-descriptions-item>
                <el-descriptions-item label="发货时间">
                      {{ formatDate(formData.shipAt) }}
                </el-descriptions-item>
                <el-descriptions-item label="退款时间">
                      {{ formatDate(formData.refundAt) }}
                </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createOrder,
  deleteOrder,
  deleteOrderByIds,
  updateOrder,
  findOrder,
  getOrderList
} from '@/api/order'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'Order'
})

// 自动化生成的字典（可能为空）以及字段
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
  const table = await getOrderList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    order_statusOptions.value = await getDictFunc('order_status')
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
            deleteOrderFunc(row)
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
      const res = await deleteOrderByIds({ ids })
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
const updateOrderFunc = async(row) => {
    const res = await findOrder({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.reorder
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteOrderFunc = async (row) => {
    const res = await deleteOrder({ ID: row.ID })
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
  const res = await findOrder({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.reorder
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>

