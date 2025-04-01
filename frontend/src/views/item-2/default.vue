<script setup>
import { ref } from 'vue';
import { Icon } from '@iconify/vue';
import { pagination } from '@/constants/pagination';
import { buttons } from '@/constants/buttons';
import { useMessage } from 'naive-ui';

const message = useMessage();

// 状态
const state = ref({
  ipList: '', // 输入的ip列表
  ipListArray: [], // 将ipList按行分割成数组
  data: [], // 表格数据
  loading: false, // 是否正在加载
});

// 表格列
const columns = [
  {
    title: '地址',
    key: 'ip',
  },
  {
    title: '设备',
    key: 'model',
  },
  {
    title: '状态',
    key: 'status',
  },
  {
    title: '操作',
    key: 'action',
  },
  {
    title: '结果',
    key: 'result',
  },
];


// 开始处理IP列表
const handleClick = (action) => {
  // 删除空字符
  state.value.ipList = state.value.ipList.replace(/\s/g, '');

  if (state.value.ipList.length === 0) {
    message.error('请输入IP地址');
    return;
  }

  // 将ipList按行分割成数组，并过滤掉空行
  state.value.ipListArray = state.value.ipList
    .split('\n')
    .map(ip => ip.replace(/\s/g, ''))
    .filter(ip => ip.length > 0);

  // 验证所有IP地址
  for (const ip of state.value.ipListArray) {
    if (!isValidIp(ip)) {
      message.error(`IP地址 ${ip} 不合法`);
      return;
    }
  }
  // // 使用临时变量构建带换行符的字符串
  // let result = '';
  // for (let i = 0; i < state.value.ipListArray.length; i++) {
  //   result += state.value.ipListArray[i];
  //   if (i < state.value.ipListArray.length - 1) {
  //     result += '\n';
  //   }
  // }
  // state.value.ipList = result;
}

</script>

<template>
  <div>
    <n-flex vertical>
      <n-card size="small" :bordered="false">
        <n-grid x-gap="12" :cols="2">
          <n-gi>
            <n-flex>
              <n-button v-for="item in buttons" :color="item.color" size="small" @click="handleClick(item.action)">
                <template #icon>
                  <Icon :icon="item.icon" />
                </template>
                {{ item.name }}
              </n-button>
            </n-flex>
          </n-gi>
          <n-gi>
            <n-input placeholder="请输入IPMI地址,一行一个" type="textarea" size="small" :autosize="{
              minRows: 5,
              maxRows: 10,
            }" v-model:value="state.ipList" />
          </n-gi>
        </n-grid>
      </n-card>
      <n-card size="small" :bordered="false">
        <n-data-table :columns="columns" :data="state.data" :bordered="false" size="small" :pagination="pagination" />
      </n-card>
    </n-flex>
  </div>
</template>