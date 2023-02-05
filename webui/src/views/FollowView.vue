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

			// Initializing two variables that will be used to Handle the Specific Search for a User.
			// usernameToSearch: "",
			// usernameToSearchBool: true,

			// Initializing four arrays for handling the list of respectively Followers and Followings and the two same arrays but with profiles. 
			followersList: [],
			followingsList: [],
            followersListProfiles: [],									// userProfile: { fixedUsername: "", username: "", photoProfile: "", biography: "", dateOfCreation: "", numberOfPhotos: 0, numberFollowers: 0, numberFollowing: 0, name: "", surname: "", dateOfBirth: "", email: "", nationality: "", gender: ""}
            followingsListProfiles: [],									// userProfile: { fixedUsername: "", username: "", photoProfile: "", biography: "", dateOfCreation: "", numberOfPhotos: 0, numberFollowers: 0, numberFollowing: 0, name: "", surname: "", dateOfBirth: "", email: "", nationality: "", gender: ""}
            
            // Initializing a flag indicating whether to update:
            //   - followersList (value: true) or 
            //   - followingsList (value: false)
            flagFollow: false,
        }
	},

	// Declaration of the methods that will be used.
	methods: {

        // GetUserProfile Function: It retrieves the whole profile of a username.
        async getUserProfile(i) {

            this.errormsg = "";
		    this.loading = true;

            // Checking whether the followFlag is:
            //  - true  (meaning that I need to update followersListProfiles).
            //  - false (meaning that I need to update followingsListProfiles).
            if (this.flagFollow === true) {

                // Here I operate onto the followersListProfiles.
                try{

                    // Retrieving the Profile from the Back-end.
                    let responseProfile = await this.$axios.get("/users/"+this.followersList[i], {
                        headers: {
                            Authorization: "Bearer " + localStorage.getItem("BearerToken")
                        }
                })

                // Add the profile retrieved to the "usersProfiles" array.
                this.followersListProfiles.push(responseProfile.data);

                } catch (e) {

                    // If an error is encountered, display it!
                    this.errormsg = e.toString();
                }
            } else {

                // Here I operate on the followingsListProfiles.
                try{

                    // Retrieving the Profile from the Back-end.
                    let responseProfile = await this.$axios.get("/users/"+this.followingsList[i], {
                        headers: {
                            Authorization: "Bearer " + localStorage.getItem("BearerToken")
                        }
                    })

                    // Add the profile retrieved to the "usersProfiles" array.
                    this.followingsListProfiles.push(responseProfile.data);

                } catch (e) {

                    // If an error is encountered, display it!
                    this.errormsg = e.toString();
                }
            }

		    this.loading = false;            
        },  

		// followUser: Given in input the username of the person to follow, it will let the Username Logged to follow the user passed.
		// async followUser(followingUsername){
			
		// },

		// getFollowings: It returns the list of usernames of the people I am following.
		async getFollowings() {

			// Re-initializing variables to their default value.
			this.errormsg = "";
			this.loading = true;

			this.followingsList = [];
            this.followingsListProfiles = [];

            // Set the flagFollow to false (meaning that I need to work on followingsListProfiles)
            this.flagFollow = false;
			
			try {

				// Getting the list of Users from the Back-End.
                let response = await this.$axios.get("/users/" + this.username +"/followings/", {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					}
				})

				// Saving the response in the "users" array.
				this.followingsList = response.data;

                // Retrieving for every username, its Profile.
                for (let i = 0; i < this.followingsList.length; i++) {
                    this.getUserProfile(i)
                }

			} catch (e) {

				// If an error is encountered, display it!
				this.errormsg = e.toString();
			}

			// Once the entire operation has finished, re-set the "loading" flag to false, in such a way to continue.
			this.loading = false;
		},

		// getFollowers: It returns the list of usernames of the people that are following me.
		async getFollowers(){

			// Re-initializing variables to their default value.
			this.errormsg = "";
			this.loading = true;

			this.followersList = [];
            this.followersListProfiles = [];

            // Set the flagFollow to true (meaning that I need to work on followersListProfiles)
            this.flagFollow = true;
			
			try {

				// Getting the list of Users from the Back-End.
                let response = await this.$axios.get("/users/" + this.username +"/followers/", {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					}
				})

				// Saving the response in the "users" array.
				this.followersList = response.data;

                // Retrieving for every username, its Profile.
				for (let i = 0; i < this.followersList.length; i++) {
                    this.getUserProfile(i, this.followersListProfiles)
                }

			} catch (e) {

				// If an error is encountered, display it!
				this.errormsg = e.toString();
			}

			// Once the entire operation has finished, re-set the "loading" flag to false, in such a way to continue.
			this.loading = false;
		},

	},
}
</script>

<!-- Actual Page for handling the page setting. -->
<template>

	<div>
			<!-- Let's handle first the upper part that will be the static one. -->
			<h1 class="h1">WASA Photo FOLLOW</h1>

			<div class="topMenu">

				<!-- "Users List" Button -->
				<div class="topMenuButtons">
					<button type="login-button" class="btn btn-primary btn-block btn-large" v-if="!loading" @click="getFollowings"> Users List </button>
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

				<!-- If the Username to search was not found, report the error. -->
				<div class="card" v-if="usernameToSearchBool === false">
					<div class="card-body">
						<p>No User present in the Database with the {{ this.usernameToSearch }} username.</p>
					</div>
				</div>


                <!-- ------------------------ FOLLOWINGS PART ------------------------  -->
                <!-- If the followingsListProfiles is empty after the computation, refer this fact. -->
				<div class="card" v-if="followingsListProfiles.length == 0">
					<div class="card-body">
						<p>Unfortunately you are not following still anyone! </p>
					</div>
				</div>
                
                <!-- If instead, it is all ok, Display a sort of card for each of the User Profiles in the followingsList -->
				<div class="card" v-if="!loading && flagFollow == false" v-for="u in followingsListProfiles">

					<div class="card-header">
						<div class="usernameLabel">
							<b> USERNAME: {{ u.username }} </b>
						</div>
						

						<div class="buttons-menu">
							<div class = "buttonsFollowBan">
								<svg class="feather" v-if="!loading" @click="getFollowings" ><use href="/feather-sprite-v4.29.0.svg#user-check"/></svg>
								<!-- <use href="/feather-sprite-v4.29.0.svg#user-plus"/> -->

							</div>
							<div class = "buttonsFollowBan">	
								<svg class="feather" v-if="!loading" @click="getFollowings" ><use href="/feather-sprite-v4.29.0.svg#unlock"/></svg>
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



                <!-- ------------------------ FOLLOWERS PART ------------------------  -->
                <!-- If the followersListProfiles is empty after the computation, refer this fact. -->
				<div class="card" v-if="followersListProfiles.length === 0">
					<div class="card-body">
						<p>Unfortunately you are not followed still by anyone! </p>
					</div>
				</div>
                
                <!-- In alternative, Display a sort of card for each of the User Profiles in the followersList -->
                <div class="card" v-if="!loading && followFlag === true" v-for="u in followersListProfiles">

					<div class="card-header">
						<div class="usernameLabel">
							<b> USERNAME: {{ u.username }} </b>
						</div>
						

						<div class="buttons-menu">
							<div class = "buttonsFollowBan">
								<svg class="feather" v-if="!loading" @click="followUser(u.username)" ><use href="/feather-sprite-v4.29.0.svg#user-check"/></svg>
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
