<script setup>
import { Icon } from '@iconify/vue';
import { h } from 'vue';
import { RouterLink } from 'vue-router';

function renderIcon(icon) {
  return () => h(Icon, { icon })
}

const options = [
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: 'survey',
          }
        },
        { default: () => '设备扫描' }
      ),
    key: 'scan',
    icon: renderIcon("mdi:magnify-scan")
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: 'action',
          }
        },
        { default: () => '基本控制' }
      ),
    key: 'base',
    icon: renderIcon("uim:apps")
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            name: 'config',
          }
        },
        { default: () => '设置' }
      ),
    key: 'config',
    icon: renderIcon("icon-park-outline:config")
  }
]

</script>

<template>
  <div>
    <n-menu :options="options" size="small" />
  </div>
</template>
