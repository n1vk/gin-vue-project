import {createRouter, createWebHistory} from "vue-router";

const router = createRouter({
    history: createWebHistory(),
    routes: [{
        path: '/',
        component: () => import('/src/components/Login.vue'),//普通路由
    }, {
        path: '/register',
        component: () => import('/src/components/Register.vue'),//普通路由
    }]
})

export {router};
