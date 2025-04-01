<script setup>
import ComponentHeader from '@/components/header.vue'
import ComponentSider from '@/components/sider.vue'
</script>


<template>
    <div style="height: 100vh; position: relative">
        <n-layout position="absolute">
            <n-layout-header style="height: 64px; padding: 15px; --wails-draggable:drag" bordered>
                <ComponentHeader />
            </n-layout-header>
            <n-layout has-sider position="absolute" style="top: 64px; bottom: 64px">
                <n-layout-sider bordered width="170">
                    <ComponentSider />
                </n-layout-sider>
                <n-layout content-style="padding: 10px;">
                    <router-view />
                </n-layout>
            </n-layout>
            <n-layout-footer bordered position="absolute" style="height: 64px; padding: 24px">
                <n-flex justify="center">
                    hello world
                </n-flex>
            </n-layout-footer>
        </n-layout>
    </div>
</template>
