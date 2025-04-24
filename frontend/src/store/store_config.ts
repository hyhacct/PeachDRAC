import { create } from 'zustand';
import WailsResp from '@/types/wails_resp';
import { ConfigPassAddOrUpdate, ConfigPassDelete, ConfigPassGetList } from '@/wailsjs/go/apps/App';
import { Notification, Modal } from '@douyinfe/semi-ui';

// 暴露接口
interface ConfigStore {
  form: {
    id: number;
    priority: number;
    username: string;
    password: string;
    port: string;
    status: boolean;
  }
  show: boolean; // 是否显示侧边栏
  dataList: any[]; // 数据列表
  update: (update: Partial<ConfigStore>) => void;
  reset: () => void;
  AddOrUpdate: () => void;
  Delete: (id: number) => void;
  GetList: () => void;
}

// 初始状态
const initialState: Pick<ConfigStore, "form" | "show" | "dataList"> = {
  form: {
    id: 0,
    priority: 100,
    username: '',
    password: '',
    port: '',
    status: false,
  },
  show: false,
  dataList: [],
}

// 创建store
const useConfigStore = create<ConfigStore>((set, get) => ({
  ...initialState,
  update: (update: Partial<ConfigStore>) => set((state) => ({ ...state, ...update })),
  reset: () => set(initialState),

  // 定义API
  AddOrUpdate: async () => {
    try {
      const { form } = get();
      const resp: WailsResp = await ConfigPassAddOrUpdate(form.id, form.username, form.password, form.port);
      if (!resp.Status) {
        throw new Error(resp.Msg);
      }
      Notification.success({
        title: '成功',
        content: resp.Msg,
      });
      set(initialState); // 重置状态
      get().GetList(); // 刷新列表
    } catch (error: any) {
      Notification.error({
        title: '错误',
        content: error?.message,
      });
    }
  },

  // 删除
  Delete: async (id: number) => {
    try {
      Modal.error({
        title: '危险', content: '确定要删除密码组吗？删除后将无法恢复此记录, 请谨慎操作！',
        okButtonProps: {
          onClick: async () => {
            const resp: WailsResp = await ConfigPassDelete(id);
            if (!resp.Status) {
              throw new Error(resp.Msg);
            }
            Notification.success({
              title: '成功',
              content: resp.Msg,
            });
            get().GetList(); // 刷新列表
          }
        }
      });
    } catch (error: any) {
      Notification.error({
        title: '错误',
        content: error?.message,
      });
    }
  },

  // 获取列表
  GetList: async () => {
    try {
      const resp: WailsResp = await ConfigPassGetList();
      if (!resp.Status) {
        throw new Error(resp.Msg);
      }
      set({ dataList: resp.Data });
    } catch (error: any) {
      Notification.error({
        title: '错误',
        content: error?.message,
      });
    }
  }
}));

export default useConfigStore;
