import { RouteRecordRaw } from 'vue-router'


const routes: RouteRecordRaw[] = [
    { path: '/', redirect: 'rule' ,name:'entry'},
    { path: '/rule', component: () => import('../views/rules.vue'), name: 'rule' },
    { path: '/raw', component: () => import('../views/raw.vue'), name: 'raw'},
    { path: '/setting', component: () => import('../views/setting.vue'), name: 'setting' }
]

export default routes