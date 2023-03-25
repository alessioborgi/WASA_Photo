import {createRouter, createWebHashHistory} from 'vue-router'
import SearchView from '../views/SearchView.vue'
import DoLoginView from '../views/DoLoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import BanView from '../views/BanView.vue'
import FollowView from '../views/FollowView.vue'
import UpdateView from '../views/UpdateProfile.vue'
import NewPhotoView from '../views/NewPhotoView.vue'
import PhotoDetailsView from '../views/PhotoDetailsView.vue'
import MyStreamView from '../views/MyStreamView.vue'
import AnalyticsView from '../views/AnalyticsView.vue'
import FutureWorkView from '../views/FutureWorkView.vue'



const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/session/', component: DoLoginView},
		{path: '/search/', component: SearchView},
		{path: '/users/:username', component: ProfileView},
		{path: '/users/:username/ban/', component: BanView},
		{path: '/users/:username/follow/', component: FollowView},
		{path: '/users/:username/update/', component: UpdateView},
		{path: '/users/:username/photo/', component: NewPhotoView},
		{path: '/users/:username/photo/:idphoto', component: PhotoDetailsView},
		{path: '/users/:username/myStream/', component: MyStreamView},
		{path: '/users/:username/analytics/', component: AnalyticsView},
		{path: '/futureWork/', component: FutureWorkView},
	]
})

export default router