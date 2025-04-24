import { ViewLayout } from '@/layout';
import useSurveyStore from "@/store/store_survey"
import { useEffect } from "react"




function App() {
  document.body.setAttribute('theme-mode', 'dark');

  // 监听任务集中注册
  const { onTask: onTaskSurvey } = useSurveyStore();
  useEffect(() => {
    onTaskSurvey();
  }, []);

  return (
    <div>
      <ViewLayout />
    </div>
  );
}

export default App;