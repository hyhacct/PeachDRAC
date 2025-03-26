import { createRouter, createWebHashHistory } from 'vue-router'
import ViewItem1 from '@/views/item-1/default.vue'

const routes = [
    {
        path: '/',
        name: 'item-1',
        component: ViewItem1,
    },
    {
        path: '/',
        name: 'item-2',
        component: ViewItem1,
    }, {
        path: '/',
        name: 'item-3',
        component: ViewItem1,
    }, {
        path: '/',
        name: 'item-4',
        component: ViewItem1,
    }
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