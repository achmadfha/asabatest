import { createRouter, createWebHistory } from 'vue-router'
import { getUser } from '@/utils'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import("views/Home.vue"),
    redirect: '/index',
    children: [
      {
        path: '/index',
        name: 'Index',
        component: () => import("views/home/Hello.vue")
      },
      {
        path: '/items',
        name: 'Items',
        component: () => import("views/items/Items.vue")
      },
      {
        path: '/404',
        name: '404',
        component: () => import('views/others/404.vue')
      },
      {
        path: '/:pathMatch(.*)',
        redirect: '/404'
      }
    ]
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import("views/auth/Login.vue")
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.WEAVE_BASE),
  routes
})

router.beforeEach((to, from, next) => {
  let isAuthenticated = getUser();

  if (!isAuthenticated && to.name !== 'Login') {
    next({ name: 'Login' });
  } else if (isAuthenticated && to.name === 'Login') {
    next({ name: 'Index' });
  } else {
    next();
  }
});


export default router