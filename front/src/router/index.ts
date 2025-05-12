import { createRouter, createWebHistory } from 'vue-router'

import UserView from '../views/user.vue'
import login from '../views/login.vue'
import Home from '../views/Home.vue'
import Index from '../views/Index.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
      children:[
        {
          path: '/user',
          name: 'user',
          component: UserView
        },
        {
          path: '/',
          name: 'index',
          component: Index
        },
        {
          path: '/discover/:query?',
          name: 'discover',
          component: () => import('../views/Discover.vue'),
        },
        {
          path: '/notification',
          name: 'notification',
          component: () => import('../views/Notification.vue')
        },
      ]
    },
    {
      path: '/test',
      name: 'test',
      component: () => import('../views/Test.vue')
    },
    {
      path: '/testTheme',
      name: 'testTheme',
      component: () => import('../views/testTheme.vue')
    },
    {
      path: '/repo/:owner/:repo',
      name: 'repository',
      component: () => import('../views/RepoView.vue')
    },
    
    {
      path: '/repo/:owner/:repo/commits/:branch', //Commits列表
      name: 'repository-commits',
      component: () => import('../views/Commits.vue')
    },
    // {
    //   path: '/repo/:owner/:repo/commit/:branch', //具体的Commit
    //   name: 'repository',
    //   component: () => import('../views/RepoView.vue')
    // },
    {
      path: '/repo/:owner/:repo/tree/:branch/:path*',
      name: 'RepoTree',
      component: () => import('../views/RepoTree.vue')
    },
    {
      path: '/repo/:owner/:repo/blob/:path',
      name: 'RepoBlob',
      component: () => import('../views/RepoBlob.vue')
    },
    {
      path: '/login',
      name: 'login',
      component: login
    }
  ]
})

export default router
