import { ViewLayout } from '@/layout';

function App() {
  document.body.setAttribute('theme-mode', 'dark');

  return (
    <div>
      <ViewLayout />
    </div>
  );
}

export default App;