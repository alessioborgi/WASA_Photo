<script setup>
import { ref } from 'vue';
import { RouterLink, RouterView } from 'vue-router'
</script>

<script>
export default {

	data: function() {
		return {

			// Retrieving from the Cache the Username and the Bearer Authenticaiton Token.
			usernameLogged: localStorage.getItem('Username'),
			username: localStorage.getItem('Username') == localStorage.getItem('usernameProfileToView') ? localStorage.getItem('Username') : localStorage.getItem('usernameProfileToView'),
            BearerToken: localStorage.getItem('BearerToken'),
		}
	},

	methods: {
		
		async goToLogin() {
			localStorage.clear();
			localStorage.setItem('usernameProfileToView', localStorage.getItem('Username'));
			this.$router.replace({ path: `/session/` })
		},
		async goToLogout() {
			localStorage.clear();
			localStorage.setItem('usernameProfileToView', localStorage.getItem('Username'));
			this.$router.replace({ path: `/session/` })
		},
		async goToMyProfile() {
			this.username = this.usernameLogged;
			localStorage.setItem('usernameProfileToView', localStorage.getItem('Username'));
			this.$router.push({ path: `/users/${this.username}` })
		},

		async goToMyStream() {
			localStorage.setItem('usernameProfileToView', localStorage.getItem('Username'));
			this.$router.push({ path: `/users/${this.username}/myStream/` })
		},
		async goToUpdateProfile() {
			localStorage.setItem('usernameProfileToView', localStorage.getItem('Username'));
			this.$router.push({ path: `/users/${this.username}/update/` })
		},
		async goToBan() {
			localStorage.setItem('usernameProfileToView', localStorage.getItem('Username'));
			this.$router.replace({ path: `/users/${this.username}/ban/` })
		},
		async goToFollow() {
			localStorage.setItem('usernameProfileToView', localStorage.getItem('Username'));
			this.$router.push({ path: `/users/${this.username}/follow/` })
		},
		async goToNewPhoto() {
			localStorage.setItem('usernameProfileToView', localStorage.getItem('Username'));
			this.$router.push({ path: `/users/${this.username}/photo/` })
		},

		async goToSearch() {
			localStorage.setItem('usernameProfileToView', localStorage.getItem('Username'));
			this.$router.push({ path: `/search/` })
		},
		async goToAnalytics() {
			localStorage.setItem('usernameProfileToView', localStorage.getItem('Username'));
			this.$router.push({ path: `/users/${this.username}/analytics/` })
		},

		async goToFutureWork() {
			localStorage.setItem('usernameProfileToView', localStorage.getItem('Username'));
			this.$router.push({ path: `/users/${this.username}/futureWork/` })
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
						<li class="nav-item" @click="goToLogin" style="margin-left: 10px; margin-top: 10px; margin-right: 10px;">
							<!-- <RouterLink to="/session/" class="nav-link"> -->
								<svg class="feather" style="color:#4a77d4; margin-right: 10px; "><use href="/feather-sprite-v4.29.0.svg#log-in"/></svg>
								<b>Login</b>
							<!-- </RouterLink> -->
						</li>
						<li class="nav-item" @click="goToLogout" style="margin-left: 10px; margin-top: 10px; margin-right: 10px;">
							<!-- <RouterLink to="/session/" class="nav-link"> -->
								<svg class="feather" style="color:#4a77d4; margin-right: 10px;"><use href="/feather-sprite-v4.29.0.svg#log-out"/></svg>
								<b>Logout</b>
							<!-- </RouterLink> -->
						</li>
					</ul>

				

					
					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-5 mb-1 text-muted text-uppercase">
						<span>PERSONAL PROFILE</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item" @click="goToMyProfile" style="margin-left: 10px; margin-top: 10px; margin-right: 10px;">
							<!-- <RouterLink :to="'/users/'+username" class="nav-link" > -->
								<svg class="feather" style="color:#4a77d4; margin-right: 10px;"><use href="/feather-sprite-v4.29.0.svg#instagram"/></svg>
								<b>My Profile</b>
							<!-- </RouterLink> -->
						</li>
						<li class="nav-item" @click="goToMyStream" style="margin-left: 10px; margin-top: 10px; margin-right: 10px;">
							<!-- <RouterLink :to="'/users/'+username+'/myStream/'" class="nav-link" > -->
								<svg class="feather" style="color:#4a77d4;margin-right: 10px; "><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
								<b>My Stream</b>
							<!-- </RouterLink> -->
						</li>
					</ul>



					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-5 mb-1 text-muted text-uppercase">
						<span>ACTIONS</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item" @click="goToUpdateProfile" style="margin-left: 10px; margin-top: 10px; margin-right: 10px;">
							<!-- <RouterLink :to="'/users/'+username+'/update/'" class="nav-link" > -->
								<svg class="feather" style="color:#4a77d4;  10px;margin-right: 10px;  "><use href="/feather-sprite-v4.29.0.svg#edit"/></svg>
								<b>Update Profile</b>
							<!-- </RouterLink> -->
						</li>
						<li class="nav-item" @click="goToBan" style="margin-left: 10px; margin-top: 10px; margin-right: 10px;">
							<!-- <RouterLink :to="'/users/'+username+'/ban/'" class="nav-link" > -->
								<svg class="feather" style="color:#4a77d4;  margin-right: 10px; "><use href="/feather-sprite-v4.29.0.svg#lock"/></svg>
								<b>Ban</b>
							<!-- </RouterLink> -->
						</li>
						<li class="nav-item" @click="goToFollow" style="margin-left: 10px; margin-top: 10px; margin-right: 10px;">
							<!-- <RouterLink :to="'/users/'+username+'/follow/'" class="nav-link" > -->
								<svg class="feather" style="color:#4a77d4; margin-right: 10px; "><use href="/feather-sprite-v4.29.0.svg#user"/></svg>
								<b>Follow</b>
							<!-- </RouterLink> -->
						</li>
						<li class="nav-item" @click="goToNewPhoto" style="margin-left: 10px; margin-top: 10px; margin-right: 10px;">
							<!-- <RouterLink :to="'/users/'+username+'/photo/'" class="nav-link" > -->
								<svg class="feather" style="color:#4a77d4; margin-right: 10px; "><use href="/feather-sprite-v4.29.0.svg#image"/></svg>
								<b>New Photo</b>
							<!-- </RouterLink> -->
						</li>
					</ul>

					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-5 mb-1 text-muted text-uppercase">
						<span>GENERAL</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item" @click="goToSearch" style="margin-left: 10px; margin-top: 10px; margin-right: 10px;">
							<!-- <RouterLink to="/search/" class="nav-link"> -->
								<svg class="feather" style="color:#4a77d4; margin-right: 10px; "><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
								<b>Search</b>
							<!-- </RouterLink> -->
						</li>
					</ul>

					<h5 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-5 mb-1 text-muted text-uppercase">
						<span>STATISTICS</span>
					</h5>
					<ul class="nav flex-column">
						<li class="nav-item" @click="goToAnalytics" style="margin-left: 10px; margin-top: 10px; margin-right: 10px;">
							<!-- <RouterLink :to="'/users/'+username+'/analytics/'" class="nav-link"> -->
								<svg class="feather" style="color:#4a77d4; margin-right: 10px; "><use href="/feather-sprite-v4.29.0.svg#bar-chart"/></svg>
								<b>Analytics</b>
							<!-- </RouterLink> -->
						</li>
					</ul>

					<h5 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-5 mb-1 text-muted text-uppercase">
						<span>FUTURE OF WORK</span>
					</h5>
					<ul class="nav flex-column">
						<li class="nav-item" @click="goToFutureWork" style="margin-left: 10px; margin-top: 10px; margin-right: 10px;">
							<!-- <RouterLink :to="'/users/'+username+'/futureWork/'" class="nav-link"> -->
								<!-- <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#briefcase"/></svg> -->
								<svg class="feather" style="color:#4a77d4; margin-right: 10px; "><use href="/feather-sprite-v4.29.0.svg#coffee"/></svg>
								<b>Works</b>
							<!-- </RouterLink> -->
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
