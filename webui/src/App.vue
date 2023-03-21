<script setup>
import { ref } from 'vue';
import { RouterLink, RouterView } from 'vue-router'
</script>

<script>
export default {

	data: function() {
		return {

			// Retrieving from the Cache the Username and the Bearer Authenticaiton Token.
            username: localStorage.getItem('Username') == "" ? "NOT LOGGED" : localStorage.getItem('Username'),
            BearerToken: localStorage.getItem('BearerToken'),
		}
	},

	methods: {
		
		// replaceLogin Function: It fills the "users" array with the usernames present in the DB.
		async replaceLogin() {

				// Re-addressing the page to the personal profile page of a user.
                this.$router.replace({ path: '/session/' })
		},

		// setLocalStorage Function: It will free-up the two localStorage settings (username and bearerauth).
		async setLocalStorage() {
			localStorage.clear();
		},

		async setUsernameAgain() {
			localStorage.setItem('usernameProfileToView', localStorage.getItem('Username'));
		},
	},

}
</script>

<template>

	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-7" href="#/session/">WASA Photo</a>
		<!-- <button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button> -->
	</header>

	<div class="container-fluid">
		<div class="row">
			<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
				<div class="position-sticky pt-3 sidebar-sticky">
					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>Menu</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink to="/session/" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#log-in" @click="{{ this.$router.replace({ path: '/session/' }); }}"/></svg>
								Login
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/session/" class="nav-link">
								<svg class="feather" :style="{feather:'red'}"><use href="/feather-sprite-v4.29.0.svg#log-out" @click="setLocalStorage()"/></svg>
								Logout
							</RouterLink>
						</li>
					</ul>

				

					
					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-5 mb-1 text-muted text-uppercase">
						<span>PERSONAL PROFILE</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item" @click="setUsernameAgain">
							<RouterLink :to="'/users/'+username" class="nav-link" >
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#instagram"/></svg>
								My Profile
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink :to="'/users/'+username+'/myStream/'" class="nav-link" >
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
								My Stream
							</RouterLink>
						</li>
					</ul>



					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-5 mb-1 text-muted text-uppercase">
						<span>ACTIONS</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink :to="'/users/'+username+'/newUsername/'" class="nav-link" >
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#edit-2"/></svg>
								Set Username
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink :to="'/users/'+username+'/update/'" class="nav-link" >
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#edit"/></svg>
								Update Profile
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink :to="'/users/'+username+'/ban/'" class="nav-link" >
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#lock"/></svg>
								Ban
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink :to="'/users/'+username+'/follow/'" class="nav-link" >
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user"/></svg>
								Follow
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink :to="'/users/'+username+'/photo/'" class="nav-link" >
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#image"/></svg>
								New Photo
							</RouterLink>
						</li>
					</ul>

					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-5 mb-1 text-muted text-uppercase">
						<span>GENERAL</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink to="/search/" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
								Search
							</RouterLink>
						</li>
					</ul>

					<h5 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-5 mb-1 text-muted text-uppercase">
						<span>STATISTICS</span>
					</h5>
					<ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink :to="'/users/'+username+'/analytics/'" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#bar-chart"/></svg>
								Analytics
							</RouterLink>
						</li>
					</ul>

					<h5 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-5 mb-1 text-muted text-uppercase">
						<span>FUTURE OF WORK</span>
					</h5>
					<ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink :to="'/users/'+username+'/futureWork/'" class="nav-link">
								<!-- <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#briefcase"/></svg> -->
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#coffee"/></svg>
								Works
							</RouterLink>
						</li>
					</ul>
				</div>
			</nav>

			<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
				<RouterView />
			</main>
		</div>
	</div>
</template>

<style>
</style>
