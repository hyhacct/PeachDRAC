import { create } from "zustand";
import { EventsOn, ClipboardSetText } from "@wails/runtime/runtime";
import { SurveyStart } from "@wails/go/apps/App";
import { Notification } from "@douyinfe/semi-ui";

interface WailsTask {
  id: string;
  done: boolean;
  exit: boolean;
  msg: string;
  args: string[];
  login: boolean;
}

interface SurveyData {
  ipmi: string;
  status: string;
  model: string;
  sn: string;
  msg: string;
}

interface SurveyStore {
  ipmi: string;
  paragraph: string;
  filter: string;
  dataList: SurveyData[];
  loading: boolean;
  update: (update: Partial<SurveyStore>) => void;
  reset: () => void;
  onTask: () => () => void; // 返回清理函数
  Start: (ips: string[]) => void;
  filterData: () => SurveyData[]; // 过滤后的数据
  onlineLength: () => number; // 在线数量
  copyOnline: () => Promise<boolean>; // 复制在线设备的IPMI，型号，SN
}

const initialState: Pick<
  SurveyStore,
  "ipmi" | "paragraph" | "filter" | "dataList" | "loading"
> = {
  ipmi: "",
  paragraph: "",
  filter: "",
  dataList: [],
  loading: false,
};

const useSurveyStore = create<SurveyStore>((set, get) => ({
  ...initialState,
  update: (update: Partial<SurveyStore>) =>
    set((state) => ({ ...state, ...update })),
  reset: () => set(initialState),
  onTask: () => {
    EventsOn("wails_task", (task: WailsTask) => {
      try {
        set((state) => {
          // 检查是否需要添加新记录
          const exists = state.dataList.some((item) => item.ipmi === task.id);
          let newDataList: SurveyData[];

          if (!exists) {
            newDataList = [
              ...state.dataList,
              {
                ipmi: task.id,
                status:
                  (task.done && task.exit) || task.login
                    ? "online"
                    : task.done && !task.exit
                    ? "completed"
                    : task.exit
                    ? "offline"
                    : "ready",
                model: task.args[0] || "",
                sn: task.args[1] || "",
                msg: task.msg,
              },
            ];
          } else {
            // 更新现有记录
            newDataList = state.dataList.map((item) =>
              item.ipmi === task.id
                ? {
                    ...item, // 创建新对象
                    status:
                      (task.done && task.exit) || task.login
                        ? "online"
                        : task.done && !task.exit
                        ? "completed"
                        : task.exit
                        ? "offline"
                        : "ready",
                    model:
                      task.done && task.exit
                        ? task.args[0] || item.model
                        : item.model,
                    sn:
                      task.done && task.exit
                        ? task.args[1] || item.sn
                        : item.sn,
                    msg: task.msg,
                  }
                : item
            );
          }
          return { dataList: newDataList };
        });
      } catch (error) {
        console.error("Failed to process wails_task:", error);
      }
    });

    return () => {};
  },
  Start: async (ips: string[]) => {
    try {
      set({ loading: true });
      const resp = await SurveyStart(ips);
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
  filterData: () => {
    const { dataList, filter } = get();
    // 循环表格数据，筛选出符合条件的记录
    if (!filter) {
      return dataList;
    }
    const filteredDataList = dataList.filter((item) => {
      if (filter === item?.status) {
        return item;
      }
    });
    return filteredDataList;
  },
  onlineLength: () => {
    const { dataList } = get();
    const onlineList = dataList.filter((item) => item.status === "online");
    return onlineList.length;
  },
  copyOnline: async () => {
    const { dataList } = get();
    const onlineList = dataList.filter((item) => item.status === "online");
    let text = onlineList
      .map((item) => `${item.ipmi} ${item.model} ${item.sn}`)
      .join("\n");
    return ClipboardSetText(text);
  },
}));

export default useSurveyStore;
