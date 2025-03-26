<script setup>
import { HomeOutline } from '@vicons/ionicons5'
import { ContourFinding, TextIndentMore } from '@vicons/carbon'
import { InfoShield20Regular } from '@vicons/fluent'
import { NIcon } from 'naive-ui'
import { h } from 'vue'
import { RouterLink } from 'vue-router'



function renderIcon(icon) {
    return () => h(NIcon, null, { default: () => h(icon) })
}


const menuOptions = [
    {
        label: () =>
            h(
                RouterLink,
                {
                    to: {
                        name: 'item-1',
                    }
                },
                { default: () => '网络扫描' }
            ),
        key: 'go-back-home',
        icon: renderIcon(ContourFinding)
    },
    {
        label: () =>
            h(
                RouterLink,
                {
                    to: {
                        name: 'item-2',
                    }
                },
                { default: () => '单机详情' }
            ),
        key: 'go-back-home',
        icon: renderIcon(InfoShield20Regular)
    },
    {
        label: () =>
            h(
                RouterLink,
                {
                    to: {
                        name: 'item-3',
                    }
                },
                { default: () => '批量操作' }
            ),
        key: 'go-back-home',
        icon: renderIcon(TextIndentMore)
    },
];
</script>


<template>
    <div>
        <n-menu mode="horizontal" :options="menuOptions" responsive />
    </div>
</template>
