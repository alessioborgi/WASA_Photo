import {createRouter, createWebHashHistory} from 'vue-router'
// import HomeView from '../views/HomeView.vue'
import DoLoginView from '../views/DoLoginView.vue'
import ProfileView from '../views/ProfileView.vue'
// import GetUsersView from '../views/GetUsersView.vue'


const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/session/', component: DoLoginView},
		{path: '/users/:username', component: ProfileView, props: true},
	]
})

export default router