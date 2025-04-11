import { defineStore } from "pinia";
import { ref } from "vue";
import { CommonAction } from "@wails/go/apps/App";

export const useStoreActions = defineStore("actions", () => {

  const state = ref({
    isLoading: false, // 是否正在加载
    data: [], // 数据
    ipList: '', // IP列表
    item: null, // 当前操作项
  });

  const form = ref({
    action: '', // 操作
    ips: [], // IP列表
    fan: {
      speed: -1, // 风扇转速
    },
    nfs: {
      ip: '',
      path: ''
    },
  });

  // 提交表单
  const SubmitForm = async () => {
    console.log("提交表单 === ", form.value);
    try {
      state.value.isLoading = true;
      const resp = await CommonAction(form.value);
    } finally {
      state.value.isLoading = false;
    }
  };

  return {
    state,
    form,
    SubmitForm,
  };
});