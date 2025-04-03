<script setup>
import { Icon } from '@iconify/vue';
import Drawer from './drawer.vue';
import { ref } from 'vue';
import {
  ConfigGetAll,
  ConfigDelete,
  SystemFindJavaInstalls,
  ConfigJavaAddOrUpdate,
  ConfigJavaGetAll,
} from '@wails/go/apps/App';
import { useMessage, useDialog } from 'naive-ui';

const state = ref({
  loading: false,
  javaInstalls: [],

  javaVersionDell: null,
  javaPathDell: '',

  javaVersionInspur: null,
  javaPathInspur: '',
});

const drawerRef = ref(null);
const list = ref([]);
const message = useMessage();
const dialog = useDialog();

// 获取配置列表
const getList = async () => {
  const res = await ConfigGetAll();
  list.value = res?.status ? res.data : [];
};
getList();


// 删除配置
const deleteConfig = async (id) => {
  dialog.warning({
    title: '警告',
    content: `确定要删除${id}号配置吗？此操作不可恢复`,
    positiveText: '确定',
    negativeText: '取消',
    onPositiveClick: async () => {
      const res = await ConfigDelete(id);
      if (res?.status) {
        getList();
        message.success('删除成功');
      } else {
        message.error(res?.msg || '删除失败');
      }
    }
  });
};


// 获取系统中的Java安装
const handleGetJavaInstalls = async () => {
  state.value.loading = true;
  state.value.javaInstalls = []; // 清空选项
  try {
    const res = await SystemFindJavaInstalls();
    if (res.status) {
      res.data.map((item) => {
        state.value.javaInstalls.push({
          label: item.version,
          value: item.path,
        })
      })
    } else {
      message.error(res.msg);
    }
  } catch (error) {
    message.error(error);
  } finally {
    state.value.loading = false;
  }
}



// 更换Java版本 戴尔
const handleUpdateJavaDell = async (value, row) => {
  const res = await ConfigJavaAddOrUpdate({
    id: row.id,
    javaPath: row.value,
    javaVersion: row.label,
    manufacturer: '戴尔',
  });
  if (res.status) {
    message.success(`Java更改到 ${value}`);
    getJavaConfig();
  } else {
    message.error(res.msg);
  }
}

// 更换Java版本 浪潮
const handleUpdateJavaInspur = async (value, row) => {
  const res = await ConfigJavaAddOrUpdate({
    id: row.id,
    javaPath: row.value,
    javaVersion: row.label,
    manufacturer: '浪潮',
  });
  if (res.status) {
    message.success(`Java更改到 ${value}`);
    getJavaConfig();
  } else {
    message.error(res.msg);
  }
};

// 获取Java配置
const getJavaConfig = async () => {
  const res = await ConfigJavaGetAll();
  if (res?.status) {
    res.data.map((item) => {
      if (item.manufacturer === '戴尔') {
        state.value.javaVersionDell = item.javaVersion;
        state.value.javaPathDell = item.javaPath;
      } else if (item.manufacturer === '浪潮') {
        state.value.javaVersionInspur = item.javaVersion;
        state.value.javaPathInspur = item.javaPath;
      }
    })
  }
};
getJavaConfig();
</script>

<template>
  <div>
    <Drawer ref="drawerRef" :getList="getList" />
  </div>
  <div>
    <n-flex vertical>
      <n-card size="small" :bordered="false" title="Java配置">
        <div>
          <n-form label-placement="left" label-width="auto" size="small">
            <n-form-item label="戴尔服务器">
              <n-select v-model:value="state.javaVersionDell" :options="state.javaInstalls" placeholder="请选择控制台Java版本"
                @focus="handleGetJavaInstalls" :loading="state.loading" clearable @update:value="handleUpdateJavaDell">
                <template #empty>
                  <n-spin :show="state.loading">
                    <template #description>
                      正在获取Java安装...
                    </template>
                  </n-spin>
                </template>
              </n-select>
            </n-form-item>
            <n-form-item label="浪潮服务器">
              <n-select v-model:value="state.javaVersionInspur" :options="state.javaInstalls" placeholder="请选择控制台Java版本"
                @focus="handleGetJavaInstalls" :loading="state.loading" clearable
                @update:value="handleUpdateJavaInspur">
                <template #empty>
                  <n-spin :show="state.loading">
                    <template #description>
                      正在获取Java安装...
                    </template>
                  </n-spin>
                </template>
              </n-select>
            </n-form-item>
          </n-form>
        </div>
      </n-card>
      <n-card size="small" :bordered="false" title="密码组">
        <n-flex vertical>
          <n-flex justify="end">
            <n-button type="primary" dashed size="small" @click="drawerRef.open()">
              <template #icon>
                <Icon icon="solar:pen-2-bold" />
              </template>
              新增配置
            </n-button>
          </n-flex>
          <n-result v-if="list.length === 0" status="418" title="为什么是空的呢?" description="哦，原来是你还没创建配置~" />
          <n-list v-else>
            <n-list-item v-for="item in list" :key="item.id">
              <n-collapse>
                <n-collapse-item :name="item.id">
                  <template #header>
                    <n-tag :bordered="false" type="success" size="small">
                      {{ item.id }}号配置
                    </n-tag>
                  </template>
                  <template #header-extra>
                    <n-flex>
                      <n-button type="error" dashed size="small" @click="deleteConfig(item.id)">
                        <template #icon>
                          <Icon icon="solar:trash-bin-minimalistic-broken" />
                        </template>
                      </n-button>
                      <n-button type="info" dashed size="small" @click="drawerRef.open(item)">
                        <template #icon>
                          <Icon icon="solar:pen-2-bold" />
                        </template>
                      </n-button>
                    </n-flex>
                  </template>
                  <div>
                    <n-descriptions label-placement="left" size="small" :label-style="{ width: '100px' }" column="1">
                      <n-descriptions-item label="账号">
                        <n-tag size="small" :bordered="false">
                          {{ item.username }}
                        </n-tag>
                      </n-descriptions-item>
                      <n-descriptions-item label="密码">
                        <n-tag size="small" :bordered="false">
                          {{ item.password }}
                        </n-tag>
                      </n-descriptions-item>
                      <n-descriptions-item label="端口">
                        <n-tag size="small" :bordered="false">
                          {{ item.port }}
                        </n-tag>
                      </n-descriptions-item>
                      <n-descriptions-item label="创建时间">
                        <n-tag size="small" :bordered="false">
                          <template #icon>
                            <Icon icon="solar:alarm-sleep-bold-duotone" />
                          </template>
                          {{ item.createdAt }}
                        </n-tag>
                      </n-descriptions-item>
                    </n-descriptions>
                  </div>
                </n-collapse-item>
              </n-collapse>
            </n-list-item>
          </n-list>
        </n-flex>
      </n-card>
    </n-flex>
  </div>
</template>