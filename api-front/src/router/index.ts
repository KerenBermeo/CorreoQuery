import type { RouteRecordRaw } from "vue-router";
import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import SearchList from '../components/SearchList.vue'

const routes: RouteRecordRaw[] = [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/searchList/:query',
      name: 'searchList',
      component: SearchList
    }
]
const router = createRouter({
    history: createWebHistory(),
    routes
});


export default router;
