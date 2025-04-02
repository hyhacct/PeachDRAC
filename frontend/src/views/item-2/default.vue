<script setup>
import { ref, h } from 'vue';
import { Icon } from '@iconify/vue';
import { pagination } from '@/constants/pagination';
import { buttons } from '@/constants/buttons';
import { useMessage, NTag } from 'naive-ui';
import { CommonAction } from '@wails/go/apps/App';
import IPTraverser from '@/utils/ip';
import { EventsOn } from '@wails/runtime/runtime';

const message = useMessage();
const ipTraverser = new IPTraverser();

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
    width: 200,
  },
  {
    title: '状态',
    key: 'status',
    defaultSortOrder: 'ascend',
    sorter(rowA, rowB) {
      return rowA.status - rowB.status
    },
    width: 100,
    render(row) {
      return h(
        NTag,
        {
          bordered: false,
          type: row.status === null ? 'info' : row.status ? 'success' : 'error',
          size: 'small',
        },
        { default: () => row.status === null ? '就绪' : row.status ? '成功' : '失败' }
      )
    },
  },
  {
    title: '型号',
    key: 'productName',
    width: 130,
    render(row) {
      return h(
        NTag,
        {
          bordered: false,
          type: 'default',
          size: 'small',
        },
        { default: () => row.productName }
      )
    },
  },
  {
    title: '操作',
    key: 'action',
    width: 140,
  },
  {
    title: '结果',
    key: 'result',
    width: 100,
    ellipsis: {
      tooltip: true
    }
  },
];


// 开始处理IP列表
const handleClick = async (action) => {

  state.value.data = [];

  // 将ipList按行分割成数组，并过滤掉空行
  state.value.ipListArray = state.value.ipList
    .split('\n')
    .map(ip => ip.replace(/\s/g, ''))
    .filter(ip => ip.length > 0);

  // 验证所有IP地址
  for (const ip of state.value.ipListArray) {
    if (!ipTraverser.isValidIp(ip)) {
      message.error(`IP地址 ${ip} 不合法`);
      return;
    }

    // 将IP地址添加到表格中
    state.value.data.push({
      ip: ip,
      status: null,
      productName: '',
      action: action,
      result: '',
    });
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

  try {
    state.value.loading = true;
    const resp = await CommonAction({
      Action: action,
      IPs: state.value.ipListArray,
      Fan: {
        Speed: -1,
      },
    });
  } catch (error) {
    message.error("执行失败", error.message);
  } finally {
    state.value.loading = false;
  }
}



EventsOn("actions", (data) => {
  // 根据IP修改表格里面元素的状态
  state.value.data.forEach(item => {
    if (item.ip === data.ip) {
      item.status = data.status;
      item.productName = data.productName;
      item.result = data.result;
    }
  })
})
</script>

<template>
  <div>
    <n-flex vertical>
      <n-card size="small" :bordered="false">
        <n-grid x-gap="12" :cols="2">
          <n-gi>
            <n-flex>
              <n-button v-for="item in buttons" :color="item.color" size="small" @click="handleClick(item.action)"
                :loading="state.loading">
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
            }" v-model:value="state.ipList" :loading="state.loading" />
          </n-gi>
        </n-grid>
      </n-card>
      <n-card size="small" :bordered="false">
        <n-data-table :columns="columns" :data="state.data" :bordered="false" size="small" :pagination="pagination" />
      </n-card>
    </n-flex>
  </div>
</template>