import { create } from "zustand";
import { EventsOn } from "@wails/runtime/runtime";
import { ActionsStart, ActionsStop } from "@wails/go/apps/App";
import { Notification } from "@douyinfe/semi-ui";
import { verify_ip } from "@/utils/address";

interface ActionsData {
  ipmi: string;
  status: string;
  model: string;
  manufacturer: string;
  sn: string;
  action: string;
  message: string;
}

interface ActionsStore {
  ipmi: string;
  action: string;
  dataList: ActionsData[];
  loading: boolean;
  formIpmiInput: string; // 输入的IPMI
  formIpmiList: string[]; // 输入的IPMI列表
  formFan: number; // 输入的转速百分比
  formNfs: string; // 输入的挂载路径
  isExiting: boolean; // 是否正在停止

  update: (update: Partial<ActionsStore>) => void;
  reset: () => void;
  Start: (actions: string) => void;
  Stop: () => void;
  onTask: () => void;
}

const initialState: Pick<
  ActionsStore,
  | "ipmi"
  | "action"
  | "dataList"
  | "loading"
  | "formIpmiInput"
  | "formIpmiList"
  | "formFan"
  | "formNfs"
  | "isExiting"
> = {
  ipmi: "",
  action: "",
  dataList: [],
  loading: false,
  formIpmiInput: "",
  formIpmiList: [],
  formFan: 80,
  formNfs: "",
  isExiting: false,
};

const useActionsStore = create<ActionsStore>((set, get) => ({
  ...initialState,
  update: (update: Partial<ActionsStore>) =>
    set((state) => ({ ...state, ...update })),
  reset: () => set(initialState),
  Start: async (actions: string) => {
    const ipmiList = get().formIpmiInput.split("\n");
    let array: string[] = [];
    let arrayData: ActionsData[] = [];
    ipmiList.map(async (ipmi) => {
      if (!verify_ip(ipmi.trim())) {
        return; // 跳过不合法的IP
      }
      array.push(ipmi.trim());
      arrayData.push({
        ipmi: ipmi.trim(),
        status: "ready",
        model: "",
        manufacturer: "",
        sn: "",
        action: actions,
        message: "", // 消息
      });
    });
    set({ formIpmiList: array, dataList: arrayData });
    set({ dataList: arrayData });

    try {
      set({ loading: true });
      const resp = await ActionsStart(
        array,
        actions,
        get().formFan,
        get().formNfs
      );
      if (!resp.Status) {
        throw new Error(resp.Msg);
      }
      Notification.success({
        title: "成功",
        content: resp.Msg,
      });
    } catch (error) {
      Notification.error({
        title: "错误",
        content: error instanceof Error ? error.message : "操作失败",
      });
    } finally {
      set({ loading: false });
    }
  },

  onTask: () => {
    EventsOn("wails_actions", (task: ActionsData) => {
      console.log("wails_actions ===>", task);
      try {
        set((state) => {
          let newDataList = state.dataList.map((item) =>
            item.ipmi === task.ipmi
              ? {
                  ...item, // 创建新对象
                  status: task.status,
                  model: task.model,
                  manufacturer: task.manufacturer,
                  sn: task.sn,
                  message: task.message,
                }
              : item
          );
          return { dataList: newDataList };
        });
      } catch (error) {
        console.error("Failed to process wails_actions:", error);
      }
    });

    return () => {};
  },
  Stop: async () => {
    try {
      set({ isExiting: true });
      const resp = await ActionsStop();
      if (!resp.Status) {
        throw new Error(resp.Msg);
      }
      Notification.success({
        title: "成功",
        content: resp.Msg,
      });
    } catch (error) {
      Notification.error({
        title: "错误",
        content: error instanceof Error ? error.message : "操作失败",
      });
    } finally {
      set({ isExiting: false });
      set({ loading: false });
    }
  },
}));

export default useActionsStore;
