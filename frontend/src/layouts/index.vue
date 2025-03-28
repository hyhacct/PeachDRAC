<script setup>
import ComponentHeader from '@/components/header.vue'
import ComponentSider from '@/components/sider.vue'
</script>


<template>
    <div style="height: 100vh; position: relative">
        <n-layout position="absolute">
            <n-layout-header style="height: 64px; padding: 24px" bordered>
                <ComponentHeader />
            </n-layout-header>
            <n-layout has-sider position="absolute" style="top: 64px; bottom: 64px">
                <n-layout-sider bordered :style="{ width: '200px' }">
                    <ComponentSider />
                </n-layout-sider>
                <n-layout content-style="padding: 10px;">
                    <router-view />
                </n-layout>
            </n-layout>
            <n-layout-footer bordered position="absolute" style="height: 64px; padding: 24px">
                城府路
            </n-layout-footer>
        </n-layout>
    </div>
</template>
