import { createRouter, createWebHashHistory } from 'vue-router'
import ViewItem1 from '@/views/item-1/default.vue'
import ViewItem2 from '@/views/item-2/default.vue'

const routes = [
    {
        path: '/',
        name: 'scan',
        component: ViewItem1,
    },
    {
        path: '/base',
        name: 'base',
        component: ViewItem2,
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