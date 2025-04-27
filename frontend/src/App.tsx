import { ViewLayout } from "@/layout";
import { useEffect } from "react";
import useSurveyStore from "@/store/store_survey";
import useActionsStore from "@/store/store_actions";

function App() {
  document.body.setAttribute("theme-mode", "dark");

  // 监听任务集中注册
  const { onTask: onTaskSurvey } = useSurveyStore();
  const { onTask: onTaskActions } = useActionsStore();
  useEffect(() => {
    onTaskSurvey();
    onTaskActions();
  }, []);

  return (
    <div>
      <ViewLayout />
    </div>
  );
}

export default App;
