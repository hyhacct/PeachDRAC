import { create } from 'zustand';

// 暴露接口
interface SurveyStore {
  ipmi: string; // IPMI
  paragraph: string; // 网段
  filter: string; // 过滤规则
  data: any[]; // 表格数据
  update: (Update: Partial<SurveyStore>) => void;
  reset: () => void;
}

// 初始状态
const initialState: Pick<SurveyStore, "ipmi" | "paragraph" | "filter" | "data"> = {
  ipmi: '',
  paragraph: '',
  filter: '',
  data: [],
}

// 创建store
const useSurveyStore = create<SurveyStore>((set) => ({
  ...initialState,
  update: (update: Partial<SurveyStore>) => set((state) => ({ ...state, ...update })),
  reset: () => set(initialState),
}));

export default useSurveyStore;
