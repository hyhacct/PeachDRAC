<script setup>
import Layouts from "./layouts/index.vue"
import { darkTheme } from 'naive-ui'
// import 'element-plus/theme-chalk/dark/css-vars.css'
</script>


<template>
    <n-config-provider :theme="darkTheme">
        <Layouts />
    </n-config-provider>
</template>

<style></style>
