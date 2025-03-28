<script setup>
import { ref, h } from 'vue';
import { deviceType } from '@/constants/deviceType';
import { statusType } from '@/constants/statusType';
import { NTag, useMessage } from 'naive-ui';
import IPTraverser from '@/utils/ip';
import { paragraph } from '@/constants/paragraph';
import { CommonSurvey } from '@wails/go/apps/App'

const startAddress = ref('');
const paragraphValue = ref(null);
const message = useMessage();
const ipTraverser = new IPTraverser();
const loading = ref(false);
const data = ref([]);
const onlineNumber = ref(0); // 在线设备数量

const columns = [
  {
    title: '地址',
    key: 'ip'
  },
  {
    title: '设备',
    key: 'model',
    render(row) {
      return h(
        NTag,
        {
          type: deviceType(row.model),
          size: 'small',
        },
        { default: () => row.model }
      )
    }
  },
];

// 开始处理
const handleStart = async () => {
  // 验证起始地址和结束地址是否合法
  if (!ipTraverser.isValidIp(startAddress.value)) {
    message.error('请输入正确的IP地址');
    return;
  }

  if ((!paragraphValue.value) || (paragraphValue.value < 0 || paragraphValue.value > 3)) {
    message.error('请选择正确的网段');
    return;
  }

  const ipList = ipTraverser.traverse(startAddress.value, paragraphValue.value);
  data.value = ipList.map(ip => ({
    ip: ip,
    model: '就绪',
  }));

  loading.value = true;
  // 探测指定IP范围内的设备，并且自动识别型号
  try {
    const result = await CommonSurvey(ipList);
    console.log("扫描设备 = ", result);
    if (result.length <= 0) {
      data.value = [];
      return;
    }
    data.value = result;
    message.success('扫描完成');
    // 剔除离线设备
    onlineNumber.value = data.value.filter(item => item.model !== '未知/离线').length;
  } catch (error) {
    console.log("扫描失败 = ", error);
    onlineNumber.value = 0;
    message.error('扫描失败');
  } finally {
    loading.value = false;
  }
}
</script>

<template>
  <div>
    <n-flex vertical>
      <n-card size="small" hoverable title="探测">
        <n-flex justify="space-between">
          <div>
            <n-flex>
              <n-input v-model:value="startAddress" type="text" placeholder="起始地址" size="small"
                :style="{ width: '200px' }" />
              <n-select v-model:value="paragraphValue" :options="paragraph" :style="{ width: '200px' }" size="small"
                placeholder="选择网段" />
              <n-button strong secondary size="small" :loading="loading">
                <template #icon>
                  <Icon icon="mdi:reload" />
                </template>
                重置
              </n-button>
              <n-button strong secondary size="small" type="info" @click="handleStart" :loading="loading">
                <template #icon>
                  <Icon icon="streamline:startup-solid" />
                </template>
                开始
              </n-button>
            </n-flex>
          </div>

          <div>
            <!-- 留空 -->
          </div>

          <div>
            <n-tag type="info">
              在线：{{ onlineNumber }} 台
            </n-tag>
          </div>
        </n-flex>
      </n-card>

      <n-card size="small" hoverable>
        <n-data-table size="small" :columns="columns" :data="data" :bordered="false" :loading="loading" />
      </n-card>
    </n-flex>
  </div>
</template>