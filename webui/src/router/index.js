import {createRouter, createWebHashHistory} from 'vue-router'
// import HomeView from '../views/HomeView.vue'
import DoLoginView from '../views/DoLoginView.vue'
import MyProfileView from '../views/MyProfileView.vue'
// import GetUsersView from '../views/GetUsersView.vue'


const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		// {path: '/', component: DoLoginView},
		// {path: '/doLogin', component: DoLoginView},
		{path: '/session/', component: DoLoginView},
		{path: '/users/:username', component: MyProfileView, props: true},
		// {path: '/users/:username/getMyStream', component: HomeView, props: true},
		// {path: '/some/:id/link', component: HomeView},
	]
})

export default router