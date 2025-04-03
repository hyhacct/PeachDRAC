<script setup>
import { ref, defineExpose, defineProps } from 'vue';
import { ConfigAddOrUpdate } from '@wails/go/apps/App';
import { useMessage } from 'naive-ui';

const props = defineProps({
  getList: {
    type: Function,
    required: true,
  },
});

const message = useMessage();
const formRef = ref(null)
const state = ref({
  show: false,
  loading: false,
});

const form = ref({
  username: '',
  password: '',
  port: 623,
  manufacturer: null,
});

const rules = ref({
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
  manufacturer: [{ required: true, message: '请选择厂商', trigger: 'blur' }],
});

const open = (row) => {
  state.value.show = true;
  if (row) {
    form.value = row; // 如果存在行数据, 则将行数据赋值给表单
  }
}

const close = () => {
  state.value.show = false;
}

const handleSubmit = (e) => {
  e.preventDefault();

  formRef.value?.validate(async (errors) => {
    if (errors) {
      message.error('验证失败, 请检查输入内容');
      return;
    }

    try {
      state.value.loading = true;

      // 确保数据类型正确
      const submitData = {
        ...form.value,
        port: Number(form.value.port) || 623,
        manufacturer: form.value.manufacturer || null,
      };

      const res = await ConfigAddOrUpdate(submitData);

      if (res && res.status) {
        message.success("配置更新成功");
        props.getList();
        close();
      } else {
        message.error(res?.msg || '更新失败');
      }
    } catch (error) {
      console.error('提交失败:', error);
      message.error('提交失败，请检查数据格式是否正确');
    } finally {
      state.value.loading = false;
    }
  });
}

defineExpose({
  open,
  close,
})
</script>

<template>
  <div>
    <n-drawer v-model:show="state.show" :width="500">
      <n-drawer-content title="配置管理" closable>
        <n-flex vertical>
          <n-alert title="你知道吗?" type="success">
            你可以存在无数个密码组配置,自动登录的时候会将密码组中的账号密码全试一遍,直到耗尽或者登录成功.
          </n-alert>
          <n-form ref="formRef" :model="form" :rules="rules" size="small">
            <n-form-item path="username" label="用户名">
              <n-input v-model:value="form.username" placeholder="请输入用户名" />
            </n-form-item>
            <n-form-item path="password" label="密码">
              <n-input v-model:value="form.password" placeholder="请输入密码" />
            </n-form-item>
            <n-form-item path="port" label="端口">
              <n-input-number v-model:value="form.port" clearable :style="{ width: '100%' }" placeholder="请输入端口" />
            </n-form-item>
          </n-form>
        </n-flex>
        <template #footer>
          <n-flex>
            <n-button @click="close" size="small">
              取消
            </n-button>
            <n-button type="primary" @click="handleSubmit" size="small" :loading="state.loading">
              提交
            </n-button>
          </n-flex>
        </template>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>