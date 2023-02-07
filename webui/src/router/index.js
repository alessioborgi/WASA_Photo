import {createRouter, createWebHashHistory} from 'vue-router'
import SearchView from '../views/SearchView.vue'
import DoLoginView from '../views/DoLoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import BanView from '../views/BanView.vue'
import FollowView from '../views/FollowView.vue'


const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/session/', component: DoLoginView},
		{path: '/search/', component: SearchView},
		{path: '/users/:username', component: ProfileView},
		{path: '/users/:username/ban/', component: BanView},
		{path: '/users/:username/follow/', component: FollowView},
	]
})

export default router