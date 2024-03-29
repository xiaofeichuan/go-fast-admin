<template>
	<el-dialog :title="state.dialog.title" v-model="state.dialog.isShowDialog" width="750px" @close="closeDialog">
		<el-form ref="dataFormRef" :model="state.dataForm" :rules="state.rules" size="default" label-width="90px">
			<el-row :gutter="35">
				<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
					<el-form-item label="配置名称" prop="configName">
						<el-input v-model="state.dataForm.configName" placeholder="请输入配置名称" clearable></el-input>
					</el-form-item>
				</el-col>
                <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
					<el-form-item label="配置状态" prop="status">
						<el-switch
							v-model="state.dataForm.status"
							:active-value="0"
							:inactive-value="1"
							inline-prompt
							active-text="启用"
							inactive-text="禁用"
						></el-switch>
					</el-form-item>
				</el-col>
				<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
					<el-form-item label="配置键" prop="configKey">						
						<el-input v-model="state.dataForm.configKey" placeholder="请输入配置键" clearable></el-input>
					</el-form-item>
				</el-col>
				<el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
                    <el-form-item label="配置值" prop="configValue">						
						<el-input v-model="state.dataForm.configValue" placeholder="请输入配置值" clearable></el-input>
					</el-form-item>
				</el-col>
				
				<el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
					<el-form-item label="配置描述" prop="remark">
						<el-input v-model="state.dataForm.remark" type="textarea" placeholder="请输入配置描述" maxlength="150"></el-input>
					</el-form-item>
				</el-col>
			</el-row>
		</el-form>
		<template #footer>
			<span class="dialog-footer">
				<el-button @click="closeDialog" size="default">取 消</el-button>
				<el-button type="primary" @click="onSubmit" size="default">确 定</el-button>
			</span>
		</template>
	</el-dialog>
</template>

<script setup lang="ts">
import { reactive, ref ,nextTick} from 'vue';
import { ElMessage } from 'element-plus';
import configApi from '/@/api/system/config';

// 定义子组件向父组件传值/事件
const emit = defineEmits(['refresh']);

// 定义变量内容
const dataFormRef = ref();
const state = reactive({
	dataForm: {
		configName: '', // 配置名称
        status: 0, // 配置状态	
		configKey: '', // 配置键
		configValue: '', // 配置值
		remark: '', // 配置描述
	},
	rules: {
		configName: [{ required: true, message: '配置名称不能为空', trigger: 'blur' }],
		configKey: [{ required: true, message: '配置键不能为空', trigger: 'blur' }],
        configValue: [{ required: true, message: '配置值不能为空', trigger: 'blur' }],
	},
	menuProps: {
		children: 'children',
		label: 'label',
	},
	dialog: {
		isShowDialog: false,
		isEdit: false,
		title: '',
	},
});

// 打开弹窗
const openDialog = (row: any) => {
	if (row) {
		state.dialog.isEdit = true;
		state.dialog.title = '编辑配置';
		nextTick(() => {
			Object.assign(state.dataForm, row);
		});
	} else {
		state.dialog.isEdit = false;
		state.dialog.title = '添加配置';
	}
	state.dialog.isShowDialog = true;
};

// 关闭弹窗
const closeDialog = () => {
	state.dialog.isShowDialog = false;
	// 重置表单
	dataFormRef.value.resetFields();
	
};

// 提交
const onSubmit = () => {
	dataFormRef.value.validate(async (valid: boolean) => {
		if (!valid) {
			return;
		}

		if (state.dialog.isEdit) {
			//更新
			configApi.update(state.dataForm).then((res) => {
				if (res.success) {
					ElMessage.success('更新成功');
					emit('refresh'); //刷新页面
					closeDialog();
				}
			});
		} else {
			//添加
			configApi.add(state.dataForm).then((res) => {
				if (res.success) {
					ElMessage.success('添加成功');
					emit('refresh'); //刷新页面
					closeDialog();
				}
			});
		}
	});
};

// 暴露变量
defineExpose({
	openDialog,
});
</script>

