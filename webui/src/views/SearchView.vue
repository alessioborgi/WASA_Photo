<!-- Starting of the actual Search Page Handling. -->
<script>

import ErrorMsg from '../components/ErrorMsg.vue'
import LoadingSpinner from '../components/LoadingSpinner.vue'
// Declaration of the export set.
export default {

	components: {
		ErrorMsg
	},

	// Describing what are the Return variables.
	data: function() {
		return {

			// Initializing the two errormessage and loading variables.
			errormsg: "",
			loading: false,

			// Retrieving from the Cache the Username and the Bearer Authenticaiton Token.
            username: localStorage.getItem('Username'),
            BearerToken: localStorage.getItem('BearerToken'),

			// Initializing two arrays that will handle the list of users and their profiles.
			users: [],
			usersProfiles: [],									// userProfile: { fixedUsername: "", username: "", photoProfile: "", biography: "", dateOfCreation: "", numberOfPhotos: 0, numberFollowers: 0, numberFollowing: 0, name: "", surname: "", dateOfBirth: "", email: "", nationality: "", gender: ""}
			
			// Initializing two variables that will be used to Handle the Specific Search for a User.
			usernameToSearch: "",
			usernameToSearchBool: true,

			// Initializing two arrays for handling the list of respectively Followers and Followings. 
			followersList: [],
			followingsList: [],

		}
	},

	// Declaration of the methods that will be used.
	methods: {
		
        // GetUsers Function: It fills the "users" array with the usernames present in the DB.
		async getUsers() {

			// Re-initializing variables to their default value.
			this.errormsg = "";
			this.loading = true;

			this.users = [];
			this.usersProfiles = [];

			this.usernameToSearchBool = true;
			
			try {

				// Getting the list of Users from the Back-End.
                let response = await this.$axios.get("/users/", {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					}
				})

				// Saving the response in the "users" array.
				this.users = response.data;

				// Retrieving for every username, its Profile.
				for (let i = 0; i < this.users.length; i++) {

					this.getUserProfile(i)
				}

				// Sorting the list of Profiles (newest to oldest) w.r.t. the dateOfCreation.
				this.usersProfiles.sort(function(a,b){
					return a.dateOfCreation - b.dateOfCreation;
				})

			} catch (e) {

				// If an error is encountered, display it!
				this.errormsg = e.toString();
			}

			// Once the entire operation has finished, re-set the "loading" flag to false, in such a way to continue.
			this.loading = false;
		},

        // GetUserProfile Function: It retrieves the whole profile of a username.
        async getUserProfile(i) {

			try{

				// Retrieving the Profile from the Back-end.
				let responseProfile = await this.$axios.get("/users/"+this.users[i], {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					}
				})

				// Add the profile retrieved to the "usersProfiles" array.
				this.usersProfiles.push(responseProfile.data);

			} catch (e) {

				// If an error is encountered, display it!
				this.errormsg = e.toString();
			}
		},

		// SerchUsername: It will search for whether the Username inserted in the input is present.
		async searchUsername() {

			// Re-initializing variables to their default value.
			this.errormsg = "";
			this.loading = true;

			this.users = [];
			this.usersProfiles = [];

			this.usernameToSearchBool = true;
			
			try{

				// Let's search for the username, only if it is > 0 (of course).
				if (this.usernameToSearch.length > 0) {

					// Let's retrieve the Profile of the Username we are searching for.
					let responseProfile = await this.$axios.get("/users/"+this.usernameToSearch, {
						headers: {
							Authorization: "Bearer " + localStorage.getItem("BearerToken")
						}
					})

					// Let's add up to the "userProfiles" array the response of the profile. Note that it will be an array with only this element.
					this.usersProfiles.push(responseProfile.data);
				} else {
					// If an error is encountered, display it! Moreover, here, put the "usernameToSearchBool" flag to false.
					this.errormsg = "Err: The Username to Search for cannot be empty!";
				}
				

			} catch (e) {

				// If an error is encountered, display it! Moreover, here, put the "usernameToSearchBool" flag to false.
				this.errormsg = e.toString();

				// Set the usernameToSearchBool flag to false, meaning that we have not found it.
				this.usernameToSearchBool = false;

				// Let's handle the cases when the Error Occurs.
				if (e.response && e.response.status === 400) {
                    this.errormsg = e.response.statusText + " You Have either typed a Username that is not respecting the Regex or the User is not Present in WASAPhoto! \n The typed USERNAME is: " + this.usernameToSearch;
				}
			}
	
			// Once the entire operation has finished, re-set the "loading" flag to false, in such a way to continue.
			this.loading = false;
		},

		// SerchUsername: It will search for whether the Username inserted in the input is present.
		async goToFollowView() {

			// Re-addressing the page to the personal profile page of a user.
			this.$router.push({ path: `/users/${this.username}/follow/` })
		},

	},

	// created() {
    //     this.$root.$refs.Search = this;
    // },
}
</script>

<!-- Actual Page for handling the page setting. -->
<template>

	<div>
			<!-- Let's handle first the upper part that will be the static one. -->
			<h1 class="h1">WASA Photo SEARCH</h1>

			<div class="topMenu">

				<!-- "Users List" Button -->
				<div class="topMenuButtons">
					<button type="login-button" class="btn btn-primary btn-block btn-large" v-if="!loading" @click="getUsers"> Users List </button>
				</div>

				<!-- WASA Photo Icon -->
				<div class="topMenuColumn">
					<img src="./img/wasa-logo.png" alt="" class="img">
				</div>

				<!-- "Search Username Field" -->
				<div class="topMenuButtons">
					<div class="formControl">
						<input type="text" id="usernameToSearch" v-model="usernameToSearch" placeholder="Search Username..." class="form-control">
					</div>
					<div class= "searchButton">
						<svg class="feather" v-if="!loading" @click="searchUsername" ><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
					</div>
				</div>

			</div>


			<!-- Let's now handle the dynamic part. -->
			<div class="result">

				<!-- Let's report the Error Message(if any), and the Loading Spinner if needed. -->
				<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
				<LoadingSpinner v-if="loading"></LoadingSpinner>

				<!-- If instead, it is all ok, Display a sort of card for each of the User Profiles(Depending on we are asking the whole list or just one). -->
				<div class="card" v-if="!loading" v-for="u in usersProfiles">

					<div class="card-header">
						<div class="usernameLabel">
							<b> USERNAME: {{ u.username }} </b>
						</div>

						<div class="buttons-menu">
							<div class = "buttonsFollowBan">
								<svg class="feather" v-if="!loading" @click="goToFollowView" ><use href="/feather-sprite-v4.29.0.svg#user-check"/></svg>
								<!-- <svg class="feather" v-if="!loading" @click="followUser(u.username)" ><use href="/feather-sprite-v4.29.0.svg#user-check"/></svg> -->
								<!-- <use href="/feather-sprite-v4.29.0.svg#user-plus"/> -->

							</div>
							<div class = "buttonsFollowBan">	
								<svg class="feather" v-if="!loading" @click="getUsers" ><use href="/feather-sprite-v4.29.0.svg#unlock"/></svg>
								<!-- <use href="/feather-sprite-v4.29.0.svg#lock"/> -->

							</div>

						</div>

					</div>
					<div class="card-body">
						<p class="card-text">

							<p><b>PHOTO:</b> {{ u.photoProfile}} <br/> </p>
							<p>
								<b>NAME:</b> {{ u.name }}
								<b>SURNAME:</b> {{ u.surname }}  
								<b>FIXEDUSERNAME:</b> {{ u.fixedUsername }}<br/>
							</p>
							<p><b>BIOGRAPHY:</b> {{ u.biography }} <br/></p>
							
							<!-- DateOfCreation: {{ u.dateOfCreation}} -->

						</p>
					</div>
				</div>
			</div>
	</div>
</template>

<!-- Declaration of the style(scoped) to use. -->
<style scoped>
	@import '../assets/search.css';
</style>
