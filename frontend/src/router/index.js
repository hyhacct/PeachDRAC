import { createRouter, createWebHashHistory } from 'vue-router'
import ViewSurvey from '@/views/survey/default.vue'
import ViewAction from '@/views/action/default.vue'
import ViewConfig from '@/views/config/default.vue'

const routes = [
    {
        path: '/',
        name: 'survey',
        component: ViewSurvey,
    },
    {
        path: '/action',
        name: 'action',
        component: ViewAction,
    },
    {
        path: '/config',
        name: 'config',
        component: ViewConfig,
    },
]



const router = createRouter({
    history: createWebHashHistory(),
    routes,
})


// router.beforeEach((to, from, next) => {
//     loadingBar.start();
//     next(); // 继续导航
// });

// router.afterEach(() => {
//     loadingBar.finish(); // 导航完成后停止loading
// });



export default router