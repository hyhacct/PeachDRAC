<script setup>
import ComponentHeader from '@/components/header.vue'
</script>


<template>
    <div>
        <n-layout>
            <n-layout-header style="--wails-draggable:drag">
                <ComponentHeader />
            </n-layout-header>

            <n-layout-content content-style="padding: 24px;">
                <router-view />
            </n-layout-content>
            <!-- <n-layout-footer>成府路</n-layout-footer> -->
        </n-layout>
    </div>
</template>
