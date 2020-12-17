import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Admin from '../views/Admin.vue'
import Role from '../views/Role.vue'
import Welcome from '../views/Welcome.vue'
import Goods from '../views/Goods.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    redirect: '/welcome',
    children: [
      {
        name: 'Welcome',
        path: '/welcome',
        component: Welcome
      },
      {
        path: '/admin',
        name: 'Admin',
        component: Admin
      },
      {
        path: '/role',
        name: 'Role',
        component: Role
      },
      {
        path: '/goods',
        name: 'Goods',
        component: Goods
      }
    ]
  }
]

const router = new VueRouter({
  routes
})

export default router
