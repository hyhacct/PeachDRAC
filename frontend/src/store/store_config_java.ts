import { create } from 'zustand';
import WailsResp from '@/types/wails_resp';
import { ConfigJavaAddOrUpdate, ConfigJavaDelete, ConfigJavaGetList } from "@wails/go/apps/App";
import { model } from '@wails/go/models';
import { Notification, Modal } from '@douyinfe/semi-ui';

// 暴露接口
interface ConfigStore {
  form: model.TableJava;
  show: boolean;
  dataList: model.TableJava[];
  update: (update: Partial<ConfigStore>) => void;
  reset: () => void;
  resetForm: () => void;
  AddOrUpdate: () => void;
  Delete: (id: number) => void;
  GetList: () => void;
}

// 初始状态
const initialState: Pick<ConfigStore, "dataList"> = {
  dataList: [],
}

// 初始表单状态
const initialFormState: Pick<ConfigStore, "form" | "show"> = {
  form: <model.TableJava>{
    id: 0,
    title: '',
    path: '',
    allot: '',
    ips: '',
  },
  show: false,
}


// 创建store
const useConfigStore = create<ConfigStore>((set, get) => ({
  ...initialState, // 初始化状态
  ...initialFormState, // 初始化表单状态
  update: (update: Partial<ConfigStore>) => set((state) => ({ ...state, ...update })),
  reset: () => set(initialState),
  resetForm: () => set(initialFormState),

  // 定义API
  AddOrUpdate: async () => {
    try {
      const { form } = get();
      const resp: WailsResp = await ConfigJavaAddOrUpdate(form);
      if (!resp.Status) {
        throw new Error(resp.Msg);
      }
      Notification.success({
        title: '成功',
        content: resp.Msg,
      });
      get().resetForm(); // 使用resetForm替代set(initialState)
      get().GetList();
    } catch (error: unknown) {
      Notification.error({
        title: '错误',
        content: error instanceof Error ? error.message : '操作失败',
      });
    }
  },

  // 删除
  Delete: async (id: number) => {
    Modal.error({
      title: '危险',
      content: '确定吗？删除后将无法恢复此记录, 请谨慎操作！',
      onOk: async () => {
        try {
          const resp: WailsResp = await ConfigJavaDelete(id);
          if (!resp.Status) {
            throw new Error(resp.Msg);
          }
          Notification.success({
            title: '成功',
            content: resp.Msg,
          });
          get().GetList();
        } catch (error: unknown) {
          Notification.error({
            title: '错误',
            content: error instanceof Error ? error.message : '删除失败',
          });
        }
      }
    });
  },

  // 获取列表
  GetList: async () => {
    try {
      const resp: WailsResp = await ConfigJavaGetList();
      if (!resp.Status) {
        throw new Error(resp.Msg);
      }
      get().reset(); // 重置状态
      set({ dataList: resp.Data }); // 设置数据列表
    } catch (error: unknown) {
      Notification.error({
        title: '错误',
        content: error instanceof Error ? error.message : '获取列表失败',
      });
    }
  },
}));

export default useConfigStore;
