import { RouteRecordRaw } from 'vue-router'


const routes: RouteRecordRaw[] = [
    { path: '/', redirect: 'rule' },
    { path: '/rule', component: () => import('../views/rules.vue'), name: 'rule' },
    { path: '/raw', component: () => import('../views/raw.vue') }
]

export default routes