<script setup>
import { ref } from 'vue';

const valueAddress = ref('');
const valueNetwork = ref(null);

const options = [
    {
        label: '第一网段',
        value: 0
    },
    {
        label: '第二网段',
        value: 1
    },
    {
        label: '第三网段',
        value: 2
    },
    {
        label: '第四网段',
        value: 3
    },
];

const columns = [
    {
        title: 'IP地址',
        key: 'address'
    },

];
</script>

<template>
    <div>
        <n-flex vertical>
            <n-card size="small" hoverable title="搜索">
                <n-flex>
                    <n-input v-model:value="valueAddress" type="text" placeholder="输入IP地址" size="small"
                        :style="{ width: '200px' }" />
                    <n-select v-model:value="valueNetwork" :options="options" :default-value="2"
                        :style="{ width: '200px' }" size="small" placeholder="选择网段" />
                    <n-button strong secondary size="small" type="primary">
                        <template #icon>
                            <Icon icon="mdi:reload" />
                        </template>
                        重置
                    </n-button>
                    <n-button strong secondary size="small" type="info">
                        <template #icon>
                            <Icon icon="streamline:startup-solid" />
                        </template>
                        启动
                    </n-button>
                </n-flex>
            </n-card>

            <n-card size="small" hoverable>
                <n-data-table size="small" :columns="columns" :data="data" :bordered="false" />
            </n-card>
        </n-flex>
    </div>
</template>