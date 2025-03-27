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
            name: 'scan',
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
            name: 'base',
          }
        },
        { default: () => '基本控制' }
      ),
    key: 'base',
    icon: renderIcon("uim:apps")
  },
]

</script>

<template>
  <div>
    <n-menu :options="options" />
  </div>
</template>
