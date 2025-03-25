<template>
	<div class="layout-padding">
		<el-card shadow="hover" :body-style="{ paddingBottom: 0 }">
			<el-form :model="state.tableData.param" ref="queryFormRef" :inline="true" label-width="70px">
				<el-form-item label="登录时间" prop="loginTime">
					<el-date-picker v-model="state.tableData.loginTime" style="width: 240px" type="daterange" range-separator="-" start-placeholder="开始日期" end-placeholder="结束日期"  value-format="YYYY-MM-DD HH:mm:ss" :default-time="[new Date(2000, 1, 1, 0, 0, 0), new Date(2000, 2, 1, 23, 59, 59),]" />
				</el-form-item>
				<el-form-item>
					<el-button @click="resetQueryForm()">
						<el-icon>
							<ele-Refresh />
						</el-icon>
						重置
					</el-button>
					<el-button type="primary" @click="handleQuery">
						<el-icon>
							<ele-Search />
						</el-icon>
						查询
					</el-button>
				</el-form-item>
			</el-form>
		</el-card>
		<el-card shadow="hover" style="margin-top: 10px">
			<!-- <div>
				<el-button type="primary" @click="openEditDialog()" v-permission="['system_loginLog_add']" plain>
					<el-icon>
						<ele-Plus />
					</el-icon>
					新增
				</el-button>
			</div> -->
			<el-table :data="state.tableData.data" v-loading="state.tableData.loading">
				<el-table-column prop="id" label="编号" width="60" />
				<el-table-column prop="userId" label="用户id" show-overflow-tooltip></el-table-column>
				<el-table-column prop="ip" label="IP地址" show-overflow-tooltip></el-table-column>
				<el-table-column prop="location" label="登录位置" show-overflow-tooltip></el-table-column>
				<el-table-column prop="browser" label="浏览器" show-overflow-tooltip></el-table-column>
				<el-table-column prop="os" label="操作系统" show-overflow-tooltip></el-table-column>
				<el-table-column prop="loginTime" label="登录时间" show-overflow-tooltip></el-table-column>
				<!-- <el-table-column label="操作" width="100">
					<template #default="scope">
						<el-button text type="primary" @click="openEditDialog(scope.row)" v-permission="['system_loginLog_update']">修改</el-button>
						<el-button text type="primary" @click="handleDel(scope.row)" v-permission="['system_loginLog_delete']">删除</el-button>
					</template>
				</el-table-column> -->
			</el-table>
			<el-pagination
				@size-change="handleSizeChange"
				@current-change="handleCurrentChange"
				:pager-count="5"
				:page-sizes="[10, 20, 30]"
				v-model:current-page="state.tableData.param.pageNum"
				background
				v-model:page-size="state.tableData.param.pageSize"
				layout="total, sizes, prev, pager, next, jumper"
				:total="state.tableData.total"
			>
			</el-pagination>
		</el-card>
		<EditDialog ref="editFormRef" @refresh="handleQuery()" />
	</div>
</template>

<script setup lang="ts" name="systemRole">
import { defineAsyncComponent, reactive, onMounted, ref } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import loginLogApi from '/@/api/system/loginLog';

// 引入组件
const EditDialog = defineAsyncComponent(() => import('./component/editDialog.vue'));

// 定义变量内容
const queryFormRef = ref();
const editFormRef = ref();
const state = reactive({
	tableData: {
		data: [],
		total: 0,
		loading: false,
		loginTime: [],
		param: {
			pageNum: 1,
			pageSize: 10,
			startTime: '',
			endTime: '',
		},
	},
});

// 页面加载时
onMounted(() => {
	handleQuery();
});

// 重置
const resetQueryForm = () => {
	debugger
	queryFormRef.value?.resetFields();
	handleQuery();
};

// 初始化表格数据
const handleQuery = () => {
	state.tableData.loading = true;
	state.tableData.param.startTime = state.tableData.loginTime[0]
	state.tableData.param.endTime = state.tableData.loginTime[1]
	loginLogApi.query(state.tableData.param).then((res) => {
		if (res.success) {
			state.tableData.data = res.data.list;
			state.tableData.total = res.data.totalCount;
			state.tableData.loading = false;
		}
	});
};
// // 打开编辑弹窗
// const openEditDialog = (row?: any) => {
// 	editFormRef.value.openDialog(row);
// };

// // 删除
// const handleDel = (row: any) => {
// 	ElMessageBox.confirm(`是否删除此记录?`, '提示', {
// 		confirmButtonText: '确认',
// 		cancelButtonText: '取消',
// 		type: 'warning',
// 	})
// 		.then(() => {
// 			loginLogApi.delete({ id: row.id }).then((res) => {
// 				if (res.success) {
// 					handleQuery();
// 					ElMessage.success('删除成功');
// 				}
// 			});
// 		})
// 		.catch(() => {});
// };

// 分页改变
const handleSizeChange = (val: number) => {
	state.tableData.param.pageSize = val;
	handleQuery();
};

// 分页改变
const handleCurrentChange = (val: number) => {
	state.tableData.param.pageNum = val;
	handleQuery();
};
</script>

