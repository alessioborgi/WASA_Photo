<!-- Starting of the actual Search Page Handling. -->
<script>

import ErrorMsg from '../components/ErrorMsg.vue'
import LoadingSpinner from '../components/LoadingSpinner.vue'
import CardProfile from '../components/CardProfile.vue'

// Declaration of the export set.
export default {

	components: {
		ErrorMsg,
		LoadingSpinner,
		CardProfile,
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

			// Initializing an array for handling the list of Banned user of the Logged Username.
			bannedList: [],
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

				// First, make this three calls to gather information about the followings, followers and bans of the Logged User.
				await this.getFollowings()
				await this.getFollowers()
				await this.getBans()

				// Retrieving for every username, its Profile.
				for (let i = 0; i < this.users.length; i++) {

					this.getUserProfile(i)
				}

				// Sorting the list of Profiles (newest to oldest) w.r.t. the dateOfCreation.
				this.usersProfiles.sort(function(a,b){
					
					return new Date(a.dateOfCreation) - new Date(b.dateOfCreation);
				})

				// Let's check whether there are actually some user that are following you.
                if (this.users.length == 0){
                    this.errormsg = "Err: There are still no User except you yet!";
                }

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

				// Add the profile retrieved to the "usersProfiles" array with the flags of following, follower and ban.
				responseProfile.data.boolFollowing = this.followingsList.includes(responseProfile.data.username) ? true : false
				responseProfile.data.boolFollower = this.followersList.includes(responseProfile.data.username) ? true : false
				responseProfile.data.boolBanned = this.bannedList.includes(responseProfile.data.username) ? true : false
				
				// Let's add up to the "userProfiles" array the response of the profile. Note that it will be an array with only this element.
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

					// First, make this three calls to gather information about the followings, followers and bans of the Logged User.
					await this.getFollowings()
					await this.getFollowers()
					await this.getBans()

					// Let's retrieve the Profile of the Username we are searching for.
					let responseProfile = await this.$axios.get("/users/"+this.usernameToSearch, {
						headers: {
							Authorization: "Bearer " + localStorage.getItem("BearerToken")
						}
					})

					// Add the profile retrieved to the "usersProfiles" array.
					responseProfile.data.boolFollowing = this.followingsList.includes(responseProfile.data.username) ? true : false
					responseProfile.data.boolFollower = this.followersList.includes(responseProfile.data.username) ? true : false
					responseProfile.data.boolBanned = this.bannedList.includes(responseProfile.data.username) ? true : false

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


		// getFollowings: It returns the list of usernames of the people I am following.
		async getFollowings() {

			// Re-initializing variables to their default value.
			this.errormsg = "";
			this.loading = true;

			this.followingsList = [];

			try {

				// Getting the list of Users from the Back-End.
				let response = await this.$axios.get("/users/" + this.username +"/followings/", {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					}
				})

				// Saving the response in the "followingsList" array.
				this.followingsList = response.data;

			} catch (e) {

				// If an error is encountered, display it!
				this.errormsg = e.toString();
			}

			// Once the entire operation has finished, re-set the "loading" flag to false, in such a way to continue.
			this.loading = false;
		},

		// getFollowers: It returns the list of usernames of the people that are following me.
		async getFollowers() {

			// Re-initializing variables to their default value.
			this.errormsg = "";
			this.loading = true;

			this.followersList = [];

			try {

				// Getting the list of Users from the Back-End.
				let response = await this.$axios.get("/users/" + this.username +"/followers/", {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					}
				})

				// Saving the response in the "followingsList" array.
				this.followersList = response.data;

			} catch (e) {

				// If an error is encountered, display it!
				this.errormsg = e.toString();
			}

			// Once the entire operation has finished, re-set the "loading" flag to false, in such a way to continue.
			this.loading = false;
		},

		// getBans: It returns the list of usernames of the people that has been banned by me.
		async getBans() {

			// Re-initializing variables to their default value.
			this.errormsg = "";
			this.loading = true;

			this.bannedList = [];

			try {

				// Getting the list of Users from the Back-End.
				let response = await this.$axios.get("/users/" + this.username +"/bans/", {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					}
				})

				// Saving the response in the "bannedList" array.
				this.bannedList = response.data;

			} catch (e) {

				// If an error is encountered, display it!
				this.errormsg = e.toString();
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
			<h1 class="h1">{{ username }}'s SEARCH</h1>

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
						<svg class="feather" v-if="!loading" 
							@click="searchUsername" 
							style="color:floralwhite"><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
					</div>
				</div>

			</div>


			<!-- Let's now handle the dynamic part. -->
			<div class="result">

				<!-- Let's report the Error Message(if any), and the Loading Spinner if needed. -->
				<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
				<LoadingSpinner v-if="loading"></LoadingSpinner>

				<!-- If instead, it is all ok, Display a sort of card for each of the User Profiles(Depending on we are asking the whole list or just one). -->
				<CardProfile 
					v-if="!loading" 
					v-for="u in usersProfiles" 
					:key="u.fixedUsername"
				    :user="u" 
					@refreshFollowing = "u.boolFollowing = $event"
					@refreshNumberFollowers = "u.numberFollowers = $event"
					@refreshBan = "u.boolBanned = $event"
				></CardProfile>
			</div>
	</div>
</template>

<!-- Declaration of the style(scoped) to use. -->
<style scoped>
	@import '../assets/search.css';
</style>
