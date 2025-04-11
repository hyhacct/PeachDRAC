<script setup>
import { ref, h, defineExpose, defineProps } from 'vue';
import { NFlex, NIcon } from 'naive-ui';
import { Icon } from '@iconify/vue';
import { useStoreActions } from '@/store';

const store = useStoreActions();
const show = ref(false);
const speed = ref(80); // 风扇转速默认80%

const open = () => {
  show.value = true;
}

const close = () => {
  show.value = false;
}

function renderIcon(icon) {
  return () => h(Icon, { icon })
}

const renderMark = (color, value) => {
  return h(
    NFlex,
    { style: { width: "120px" } },
    {
      default: () => [
        h(NIcon, { size: 24, color, component: renderIcon('ix:device-fan') }),
        h("span", { style: { color } }, `${value}%`)
      ]
    }
  );
}

const customMarks = {
  20: () => renderMark("#e74c3c", 20),
  40: () => renderMark("#e67e22", 40),
  60: () => renderMark("#a3e4d7", 60),
  80: () => renderMark("#2ecc71", 80),
  100: () => renderMark("#0048BA", 100)
};

// 确认按钮点击事件
const handleConfirm = () => {
  store.form.fan.speed = speed.value;
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
      <n-drawer-content title="风扇转速" closable>
        <n-flex vertical>
          <n-alert title="警告" type="warning">
            如果设备处于高负载运行,建议拉高风扇转速,否则可能导致设备过热,引发故障.
          </n-alert>
          <n-slider v-model:value="speed" :marks="customMarks" step="mark" :max="100" />
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