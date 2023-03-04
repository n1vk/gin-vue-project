import {createRouter, createWebHistory} from "vue-router";

const router = createRouter({
    history: createWebHistory(),
    routes: [{
        path: '/',
        component: () => import('/src/components/Login.vue'),
    }, {
        path: '/register',
        component: () => import('/src/components/Register.vue'),
    }, {
        path: '/home',
        component: () => import('/src/components/home.vue'),
    }]
})

export {router};
