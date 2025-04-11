<script setup>
import { ref, h, defineExpose, defineProps } from 'vue';
import { NFlex, NIcon } from 'naive-ui';
import { Icon } from '@iconify/vue';
import { useStoreActions } from '@/store';

const store = useStoreActions();
const show = ref(false);
const nfsServer = ref('');
const nfsPath = ref('');

const open = () => {
  show.value = true;
}

const close = () => {
  show.value = false;
  nfsServer.value = '';
  nfsPath.value = '';
}

// 确认按钮点击事件
const handleConfirm = () => {
  store.form.nfs.ip = nfsServer.value;
  store.form.nfs.path = nfsPath.value;
  store.SubmitForm();
  close();
}

defineExpose({
  open,
})
</script>

<template>
  <div>
    <n-drawer v-model:show="show" :width="502">
      <n-drawer-content title="NFS远程共享" closable>
        <n-flex vertical>
          <n-alert title="注意" type="info">
            需要提供你的NFS共享服务器的地址以及文件所在位置
          </n-alert>
          <n-input v-model:value="nfsServer" placeholder="NFS服务器地址" />
          <n-input v-model:value="nfsPath" placeholder="文件所在位置" />
        </n-flex>
        <template #footer>
          <n-space>
            <n-button @click="close" size="small">
              取消
            </n-button>
            <n-button type="primary" @click="handleConfirm" size="small">
              确认
            </n-button>
          </n-space>
        </template>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>