<script setup>
import Layouts from "./layouts/index.vue"
import { darkTheme } from 'naive-ui'
// import 'element-plus/theme-chalk/dark/css-vars.css'
</script>


<template>
    <n-config-provider :theme="darkTheme">
        <n-message-provider>
            <n-dialog-provider>
                <Layouts />
            </n-dialog-provider>
        </n-message-provider>
    </n-config-provider>
</template>
