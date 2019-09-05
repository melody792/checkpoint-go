import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import {checkoutLogin} from './utils/utils'

Vue.use(Router)

export default new Router({
    routes: [
        {
            path: '/login',
            name: 'login',
            component: () => import('./views/Login.vue')
        },
        {
            path: '/register',
            name: 'register',
            component: () => import('./views/Register.vue')
        },
        {
            path: '/',
            name: 'home',
            component: Home,
            meta: {title: "首页"},
            beforeEnter: (to, from, next) => {
                if (!checkoutLogin()) {
                    next('/login')
                } else next()
            }
        },
        {
            path: '/personalPage',
            name: 'personalPage',
            component: () => import('./views/PersonalPage.vue'),
            beforeEnter: (to, from, next) => {
                if (!checkoutLogin()) {
                    next('/login')
                } else next()
            }
        },
    ]
})
