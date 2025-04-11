<script setup>
import { ref, h } from 'vue';
import { pagination } from '@/constants/pagination';
import { NTag } from 'naive-ui';
import { EventsOn } from '@wails/runtime/runtime';
import { useStoreActions } from '@/store';
import { buttons } from '@/constants/buttons';
import DrawerFan from './drawer_fan.vue';
import DrawerNfs from './drawer_nfs.vue';
import IPTraverser from '@/utils/ip';


const store = useStoreActions();
const ipTraverser = new IPTraverser();
const drawerFan = ref(null);
const drawerNfs = ref(null);

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


const handleClick = async (args) => {
  store.item = args; // 当前操作项
  store.state.data = [];
  store.form.action = args.action;
  store.form.ips = store.state.ipList
    .split('\n')
    .map(ip => ip.replace(/\s/g, ''))
    .filter(ip => ip.length > 0);
  for (const ip of store.form.ips) {
    if (!ipTraverser.isValidIp(ip)) {
      message.error(`IP地址 ${ip} 不合法`);
      return;
    }
    store.state.data.push({
      ip: ip,
      status: null,
      productName: '',
      action: store.form.action, // 使用实际执行的操作
      result: '',
    });
  }

  // 需要打开抽屉
  if (args.action === 'fanAdjust') {
    drawerFan.value.open(); // 打开抽屉
    return;
  }
  if (args.action === 'mountNFS') {
    drawerNfs.value.open(); // 打开抽屉
    return;
  }

  store.SubmitForm();
};

EventsOn("actions", (data) => {
  // 根据IP修改表格里面元素的状态
  store.state.data.forEach(item => {
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
    <DrawerFan ref="drawerFan" />
    <DrawerNfs ref="drawerNfs" />
  </div>
  <div>
    <n-flex vertical>
      <n-card size="small" :bordered="false">
        <n-grid x-gap="12" :cols="2">
          <n-gi>
            <n-flex>
              <n-button v-for="item in buttons" :color="item.color" size="small" @click="handleClick(item)"
                :loading="store.state.isLoading">
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
            }" v-model:value="store.state.ipList" :loading="store.state.isLoading" />
          </n-gi>
        </n-grid>
      </n-card>
      <n-card size="small" :bordered="false">
        <n-data-table :columns="columns" :data="store.state.data" :bordered="false" size="small"
          :pagination="pagination" />
      </n-card>
    </n-flex>
  </div>
</template>