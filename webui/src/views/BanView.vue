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

			// Initializing two variables that will be used to Handle the Specific Search for a User among the Bans.
			usernameBanToSearch: "",
            

			// Initializing four arrays for handling the list of respectively Followers and Followings and the two same arrays but with profiles. 
			followersList: [],
			followingsList: [],
            			
			// Initializing two array for handling the list of Banned user of the Logged Username and the same list with profiles.
			bannedList: [],
            bannedListProfiles: [],
        }
	},

	// Declaration of the methods that will be used.
	methods: {

        async getData() {

            // Re-initializing variables to their default value.
            this.errormsg = "";
		    this.loading = true;

            this.followingsList = [];
            this.followersList = [];
            this.bannedList = [];
			
            // ----- Getting Followings. -----
			try {

				// Getting the list of Followings from the Back-End.
                let response = await this.$axios.get("/users/" + this.username +"/followings/", {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					}
				})

				// Saving the response in the "users" array.
				this.followingsList = response.data;

			} catch (e) {

				// If an error is encountered, display it!
				this.errormsg = e.toString();
			}

            // ----- Getting Followers. -----
            try {

                // Getting the list of Followers from the Back-End.
                let response = await this.$axios.get("/users/" + this.username +"/followers/", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("BearerToken")
                    }
                })

                // Saving the response in the "users" array.
                this.followersList = response.data;

            } catch (e) {

                // If an error is encountered, display it!
                this.errormsg = e.toString();
            }

            // ----- Getting Bans. -----
            try {

                // Getting the list of Bans from the Back-End.
                let response = await this.$axios.get("/users/" + this.username +"/bans/", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("BearerToken")
                    }
                })

                // Saving the response in the "users" array.
                this.bannedList = response.data;

            } catch (e) {

                // If an error is encountered, display it!
                this.errormsg = e.toString();
            }

            // ----- -----

			// Once the entire operation has finished, re-set the "loading" flag to false, in such a way to continue.
			this.loading = false;

        },

        // GetUserProfile Function: It retrieves the whole profile of a username.
        async getUserProfile(i) {

            this.errormsg = "";
		    this.loading = true;

            // Here I operate onto the bannedListProfiles. 
            try{

                // Retrieving the Profile from the Back-end.
                let responseProfile = await this.$axios.get("/users/"+this.bannedList[i], {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("BearerToken")
                    }
                })

                // Add the profile retrieved to the "usersProfiles" array with the flags of following, follower and ban.
				responseProfile.data.boolFollowing = this.followingsList.includes(responseProfile.data.username) ? true : false
				responseProfile.data.boolFollower = this.followersList.includes(responseProfile.data.username) ? true : false
				responseProfile.data.boolBanned = this.bannedList.includes(responseProfile.data.username) ? true : false
				
                // Add the profile retrieved to the "usersProfiles" array.
                this.bannedListProfiles.push(responseProfile.data);

            } catch (e) {

                // If an error is encountered, display it!
                this.errormsg = e.toString();
            }

            // Once the entire operation has finished, re-set the "loading" flag to false, in such a way to continue.
		    this.loading = false;            
        },  

		// getBans: It returns the list of usernames of the people I have banned.
		async getBans() {

			// Re-initializing variables to their default value.
			this.errormsg = "";
			this.loading = true;

            this.bannedListProfiles = [];
            this.bannedList = [];
	
            // Getting Data about Followers, Followings and Bans that will be used in Cards.
            await this.getData()

            // Get, for every user in the bannedList, the profile's data.
			try {

                // Retrieving for every username, its Profile.
                for (let i = 0; i < this.bannedList.length; i++) {
                    this.getUserProfile(i)
                }

                // Let's check whether you have banned some User.
                if (this.bannedList.length == 0){
                    this.errormsg = "Err: You don't have Banned any user!";
                }

			} catch (e) {

				// If an error is encountered, display it!
				this.errormsg = e.toString();
			}

			// Once the entire operation has finished, re-set the "loading" flag to false, in such a way to continue.
			this.loading = false;
		},

        // SerchUsername: It will search for whether the Username inserted in the input is present among the FollowingsList.
		async searchUsername() {

			// Re-initializing variables to their default value.
			this.errormsg = "";
			this.loading = true;

			this.bannedList = [];
            this.bannedListProfiles = [];

            // Let's get the data of Followers, Followings and Bans that will be used in the Cards.
            await this.getData()
			
            try{

                // Let's search for the username, only if it is > 0 (of course).
                if (this.usernameBanToSearch.length > 0) {
                        
                    // Let's check if the Username of the Banned Profile to search for is first of all among the Banned.
                    if (this.bannedList.includes(this.usernameBanToSearch)){

                        // Let's retrieve the Profile of the Username we are searching for.
                        let responseProfile = await this.$axios.get("/users/"+this.usernameBanToSearch, {
                            headers: {
                                Authorization: "Bearer " + localStorage.getItem("BearerToken")
                            }
                        })

                        // Add the profile retrieved to the "usersProfiles" array with the flags of following, follower and ban.
                        responseProfile.data.boolFollowing = this.followingsList.includes(responseProfile.data.username) ? true : false
                        responseProfile.data.boolFollower = this.followersList.includes(responseProfile.data.username) ? true : false
                        responseProfile.data.boolBanned = this.bannedList.includes(responseProfile.data.username) ? true : false
				
                        // Let's add up to the "followersListProfiles" array the response of the profile. Note that it will be an array with only this element.
                        this.bannedListProfiles.push(responseProfile.data);
                    } else {

                        // This means the username we are searching for is not a Followers of mine.
                        this.errormsg = "Err: The Username you are Searching for is not one of your Bans! Username: "+ this.usernameFollowersToSearch;
                    }
                } else {

                    // If an error is encountered, display it! 
                    this.errormsg = "Err: The Username to Search for among Bans cannot be empty!";
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
}
</script>



<!-- Actual Page for handling the page setting. -->
<template>

	<div>
			<!-- Let's handle first the upper part that will be the static one. -->
			<h1 class="h1">WASA Photo BAN</h1>

			<div class="topMenu">

				<!-- "Users List" Button -->
				<div class="topMenuButtons">
					<button type="login-button" class="btn btn-primary btn-block btn-large" v-if="!loading" @click="getBans"> Users List </button>
				</div>

				<!-- WASA Photo Icon -->
				<div class="topMenuColumn">
					<img src="./img/wasa-logo.png" alt="" class="img">
				</div>

				<!-- "Search Username Field" -->
				<div class="topMenuButtons">
					<div class="formControl">
						<input type="text" id="usernameToSearch" v-model="usernameBanToSearch" placeholder="Search Username..." class="form-control">
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
				<CardProfile v-if="!loading" v-for="u in bannedListProfiles" 
				    :user="u"> 
				</CardProfile>
			</div>
	</div>
</template>

<!-- Declaration of the style(scoped) to use. -->
<style scoped>
	@import '../assets/search.css';
</style>
