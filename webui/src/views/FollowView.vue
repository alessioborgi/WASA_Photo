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

			// Initializing two variables that will be used to Handle the Specific Search for a User among the Followings and Followers.
			usernameFollowingsToSearch: "",
            usernameFollowersToSearch: "",
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
			
			// Initializing an array for handling the list of Banned user of the Logged Username.
			bannedList: [],
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
				if (e.response && e.response.status === 400) {
					this.errormsg = "Request error, please Login before doing some action or ask to get the list of followings of a valid user." + e.toString();
                } else if (e.response && e.response.status === 403) {
                    this.errormsg = "An Unauthorized Action has been blocked. You are not allowed to do this action because you are not the profile's owner." + e.toString();
                } else if (e.response && e.response.status === 204) {
                    this.errormsg = "In the Internal DB there is not anymore the content you have asked." + e.toString();
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later." + e.toString();
                } else {
                    this.errormsg = e.toString();
                }
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
                if (e.response && e.response.status === 400) {
					this.errormsg = "Request error, please Login before doing some action or ask to get the list of followers of a valid user." + e.toString();
                } else if (e.response && e.response.status === 403) {
                    this.errormsg = "An Unauthorized Action has been blocked. You are not allowed to do this action because you are not the profile's owner." + e.toString();
                } else if (e.response && e.response.status === 204) {
                    this.errormsg = "In the Internal DB there is not anymore the content you have asked." + e.toString();
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later." + e.toString();
                } else {
                    this.errormsg = e.toString();
                }
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
                if (e.response && e.response.status === 400) {
					this.errormsg = "Request error, please Login before doing some action or ask to get the list of bans of a valid user." + e.toString();
                } else if (e.response && e.response.status === 403) {
                    this.errormsg = "An Unauthorized Action has been blocked. You are not allowed to do this action because you are not the profile's owner." + e.toString();
                } else if (e.response && e.response.status === 204) {
                    this.errormsg = "In the Internal DB there is not anymore the content you have asked." + e.toString();
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later." + e.toString();
                } else {
                    this.errormsg = e.toString();
                }
            }

            // ----- -----

			// Once the entire operation has finished, re-set the "loading" flag to false, in such a way to continue.
			this.loading = false;

        },

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

                // Add the profile retrieved to the "usersProfiles" array with the flags of following, follower and ban.
				responseProfile.data.boolFollowing = this.followingsList.includes(responseProfile.data.username) ? true : false
				responseProfile.data.boolFollower = this.followersList.includes(responseProfile.data.username) ? true : false
				responseProfile.data.boolBanned = this.bannedList.includes(responseProfile.data.username) ? true : false
				
                // Add the profile retrieved to the "usersProfiles" array.
                this.followersListProfiles.push(responseProfile.data);

                } catch (e) {

                    // If an error is encountered, display it!
                    if (e.response && e.response.status === 400) {
					this.errormsg = "Request error, please Login before doing some action or ask to get the user profile of a valid user." + e.toString();
                    } else if (e.response && e.response.status === 403) {
                        this.errormsg = "An Unauthorized Action has been blocked. You are not allowed to do this action because you are not the profile's owner." + e.toString();
                    } else if (e.response && e.response.status === 204) {
                        this.errormsg = "In the Internal DB there is not anymore the content you have asked." + e.toString();
                    } else if (e.response && e.response.status === 500) {
                        this.errormsg = "An internal error occurred. We will be notified. Please try again later." + e.toString();
                    } else {
                        this.errormsg = e.toString();
                    }
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

                    // Add the profile retrieved to the "usersProfiles" array with the flags of following, follower and ban.
                    responseProfile.data.boolFollowing = this.followingsList.includes(responseProfile.data.username) ? true : false
                    responseProfile.data.boolFollower = this.followersList.includes(responseProfile.data.username) ? true : false
                    responseProfile.data.boolBanned = this.bannedList.includes(responseProfile.data.username) ? true : false
                    
                    // Add the profile retrieved to the "usersProfiles" array.
                    this.followingsListProfiles.push(responseProfile.data);

                } catch (e) {

                    // If an error is encountered, display it!
                    if (e.response && e.response.status === 400) {
                        this.errormsg = "Request error, please Login before doing some action or ask to get the list of followings of a valid user." + e.toString();
                    } else if (e.response && e.response.status === 403) {
                        this.errormsg = "An Unauthorized Action has been blocked. You are not allowed to do this action because you are not the profile's owner." + e.toString();
                    } else if (e.response && e.response.status === 204) {
                        this.errormsg = "In the Internal DB there is not anymore the content you have asked." + e.toString();
                    } else if (e.response && e.response.status === 500) {
                        this.errormsg = "An internal error occurred. We will be notified. Please try again later." + e.toString();
                    } else {
                        this.errormsg = e.toString();
                    }
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

            this.followingsListProfiles = [];

            // Set the flagFollow to false (meaning that I need to work on followingsListProfiles)
            this.flagFollow = false;
			
            // Getting Data about Followers, Followings and Bans that will be used in Cards.
            await this.getData()

            // Get, for every user in the followingsList, the profile's data.
			try {

                // Retrieving for every username, its Profile.
                for (let i = 0; i < this.followingsList.length; i++) {
                    this.getUserProfile(i)
                }

                // Let's check whether you are actually following some User.
                if (this.followingsList.length == 0){
                    this.errormsg = "Err: You don't have followed any User yet!";
                }

			} catch (e) {

				// If an error is encountered, display it!
				if (e.response && e.response.status === 400) {
					this.errormsg = "Request error, please Login before doing some action or ask to get the profile of a valid user." + e.toString();
                } else if (e.response && e.response.status === 403) {
                    this.errormsg = "An Unauthorized Action has been blocked. You are not allowed to do this action because you are not the profile's owner." + e.toString();
                } else if (e.response && e.response.status === 204) {
                    this.errormsg = "In the Internal DB there is not anymore the content you have asked." + e.toString();
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later." + e.toString();
                } else {
                    this.errormsg = e.toString();
                }
			}

			// Once the entire operation has finished, re-set the "loading" flag to false, in such a way to continue.
			this.loading = false;
		},

		// getFollowers: It returns the list of usernames of the people that are following me.
		async getFollowers(){

			// Re-initializing variables to their default value.
			this.errormsg = "";
			this.loading = true;

            this.followersListProfiles = [];


            // Set the flagFollow to false (meaning that I need to work on followersListProfiles)
            this.flagFollow = true;
			
            // Getting Data about Followers, Followings and Bans that will be used in Cards.
            await this.getData()

            // Get, for every user in the followingsList, the profile's data.
			try {

                // Retrieving for every username, its Profile.
                for (let i = 0; i < this.followersList.length; i++) {
                    this.getUserProfile(i)
                }

                // Let's check whether there are actually some user that are following you.
                if (this.followingsList.length == 0){
                    this.errormsg = "Err: You don't have any Follower User yet!";
                }

			} catch (e) {

				// If an error is encountered, display it!
				if (e.response && e.response.status === 400) {
					this.errormsg = "Request error, please Login before doing some action or ask to get the list of followers of a valid user." + e.toString();
                } else if (e.response && e.response.status === 403) {
                    this.errormsg = "An Unauthorized Action has been blocked. You are not allowed to do this action because you are not the profile's owner." + e.toString();
                } else if (e.response && e.response.status === 204) {
                    this.errormsg = "In the Internal DB there is not anymore the content you have asked." + e.toString();
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later." + e.toString();
                } else {
                    this.errormsg = e.toString();
                }
			}

			// Once the entire operation has finished, re-set the "loading" flag to false, in such a way to continue.
			this.loading = false;
		},


        // SerchUsername: It will search for whether the Username inserted in the input is present among the FollowingsList.
		async searchUsername(flag) {

			// Re-initializing variables to their default value.
			this.errormsg = "";
			this.loading = true;

			this.followingsList = [];
            this.followingsListProfiles = [];
            this.followersList = [];
            this.followersListProfiles = [];

            // Set the flagFollow to true (meaning that I need to work on followersListProfiles)
            this.flagFollow = flag;

            // Let's get the data of Followers, Followings and Bans that will be used in the Cards.
            await this.getData()
			
            // Checking whether the followFlag is:
            //  - true  (meaning that I need to update followersListProfiles).
            //  - false (meaning that I need to update followingsListProfiles).
            if (this.flagFollow === true) {

                // Here I operate onto the followersListProfiles.
                try{

                    // Let's search for the username, only if it is > 0 (of course).
                    if (this.usernameFollowersToSearch.length > 0) {
                        
                        // Let's check if the Username of the Followers to search for is first of all among the Followers.
                        if (this.followersList.includes(this.usernameFollowersToSearch)){

                            // Let's retrieve the Profile of the Username we are searching for.
                            let responseProfile = await this.$axios.get("/users/"+this.usernameFollowersToSearch, {
                                headers: {
                                    Authorization: "Bearer " + localStorage.getItem("BearerToken")
                                }
                            })

                            // Add the profile retrieved to the "usersProfiles" array with the flags of following, follower and ban.
                            responseProfile.data.boolFollowing = this.followingsList.includes(responseProfile.data.username) ? true : false
                            responseProfile.data.boolFollower = this.followersList.includes(responseProfile.data.username) ? true : false
                            responseProfile.data.boolBanned = this.bannedList.includes(responseProfile.data.username) ? true : false
				
                            // Let's add up to the "followersListProfiles" array the response of the profile. Note that it will be an array with only this element.
                            this.followersListProfiles.push(responseProfile.data);
                        } else {

                            // This means the username we are searching for is not a Followers of mine.
                            this.errormsg = "Err: The Username you are Searching for is not one of your Followers! Username: "+ this.usernameFollowersToSearch;
                        }
                    } else {

                        // If an error is encountered, display it! 
                        this.errormsg = "Err: The Username to Search for among Followers cannot be empty!";
                    }
                    

                } catch (e) {

                    // If an error is encountered, display it! Moreover, here, put the "usernameToSearchBool" flag to false.
                    this.errormsg = e.toString();

                    // Set the usernameToSearchBool flag to false, meaning that we have not found it.
                    this.usernameToSearchBool = false;

                    // Let's handle the cases when the Error Occurs.
                    if (e.response && e.response.status === 400) {
                        this.errormsg = e.response.statusText + " You Have either typed a Username that is not respecting the Regex or the User is not Present in WASAPhoto! \n The typed USERNAME is: " + this.usernameToSearch;
                    } else if (e.response && e.response.status === 403) {
                        this.errormsg = "An Unauthorized Action has been blocked. You are not allowed to do this action because you are not the profile's owner." + e.toString();
                    } else if (e.response && e.response.status === 204) {
                        this.errormsg = "In the Internal DB there is not anymore the content you have asked." + e.toString();
                    } else if (e.response && e.response.status === 500) {
                        this.errormsg = "An internal error occurred. We will be notified. Please try again later." + e.toString();
                    } else {
                        this.errormsg = e.toString();
                    }
                }

            } else {

                // Here I operate onto the followingsListProfiles.
                try{

                    // Let's search for the username, only if it is > 0 (of course).
                    if (this.usernameFollowingsToSearch.length > 0) {
                        
                        // Let's check if the Username of the Following to search for is first of all among the Followings.
                        if (this.followingsList.includes(this.usernameFollowingsToSearch)){

                            // Let's retrieve the Profile of the Username we are searching for.
                            let responseProfile = await this.$axios.get("/users/"+this.usernameFollowingsToSearch, {
                                headers: {
                                    Authorization: "Bearer " + localStorage.getItem("BearerToken")
                                }
                            })

                            // Add the profile retrieved to the "usersProfiles" array with the flags of following, follower and ban.
                            responseProfile.data.boolFollowing = this.followingsList.includes(responseProfile.data.username) ? true : false
                            responseProfile.data.boolFollower = this.followersList.includes(responseProfile.data.username) ? true : false
                            responseProfile.data.boolBanned = this.bannedList.includes(responseProfile.data.username) ? true : false
				
                            // Let's add up to the "followingsListProfiles" array the response of the profile. Note that it will be an array with only this element.
                            this.followingsListProfiles.push(responseProfile.data);
                        } else {

                            // This means the username we are searching for is not a Following of mine.
                            this.errormsg = "Err: The Username you are Searching for is not one of your Following! Username: "+ this.usernameFollowingsToSearch;
                        }
                    } else {

                        // If an error is encountered, display it! 
                        this.errormsg = "Err: The Username to Search for among Followings cannot be empty!";
                    }
                    

                } catch (e) {

                    // If an error is encountered, display it! Moreover, here, put the "usernameToSearchBool" flag to false.
                    this.errormsg = e.toString();

                    // Set the usernameToSearchBool flag to false, meaning that we have not found it.
                    this.usernameToSearchBool = false;

                    // Let's handle the cases when the Error Occurs.
                    if (e.response && e.response.status === 400) {
                        this.errormsg = e.response.statusText + " You Have either typed a Username that is not respecting the Regex or the User is not Present in WASAPhoto! \n The typed USERNAME is: " + this.usernameToSearch;
                    }else if (e.response && e.response.status === 403) {
                        this.errormsg = "An Unauthorized Action has been blocked. You are not allowed to do this action because you are not the profile's owner." + e.toString();
                    } else if (e.response && e.response.status === 204) {
                        this.errormsg = "In the Internal DB there is not anymore the content you have asked." + e.toString();
                    } else if (e.response && e.response.status === 500) {
                        this.errormsg = "An internal error occurred. We will be notified. Please try again later." + e.toString();
                    } else {
                        this.errormsg = e.toString();
                    }
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
			<h1 class="h1">{{ username }}'s FOLLOW</h1>

			<div class="topMenu">

                <!-- Followings Menu left-Part -->
                <div class="followingsMenu" v-if="!loading">
                    
                    <h2 class="h2">FOLLOWINGS</h2>

                    <!-- "Users List" Button -->
                    <div class="topMenuButtons">
                        <button type="login-button" class="btn btn-primary btn-block btn-large" @click="getFollowings">  Followings List </button>
                    </div>

                    <!-- "Search Username Field" -->
                    <div class="topMenuButtons">
                        <div class="formControl">
                            <input type="text" id="usernameFollowingsToSearch" onfocus="this.value=''" v-model="usernameFollowingsToSearch" placeholder="Search Following..." class="form-control">
                        </div>
                        <div class= "searchButton">
                            <svg class="feather" @click="searchUsername(false)" ><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
                        </div>
                    </div>
                </div>

				<!-- WASA Photo Icon -->
				<div class="topMenuColumn">
					<img src="./img/wasa-logo.png" alt="" class="img">
				</div>

                <!-- Followings Menu left-Part -->
                <div class="followingsMenu" v-if="!loading">
                    
                    <h2 class="h2">FOLLOWERS</h2>

                    <!-- "Users List" Button -->
                    <div class="topMenuButtons">
                        <button type="login-button" class="btn btn-primary btn-block btn-large" @click="getFollowers">  Followers List </button>
                    </div>

                    <!-- "Search Username Field" -->
                    <div class="topMenuButtons">
                        <div class="formControl">
                            <input type="text" id="usernameFollowersToSearch" onfocus="{{ this.value='' }}" v-model="usernameFollowersToSearch" placeholder="Search Follower..." class="form-control">
                        </div>
                        <div class= "searchButton">
                            <svg class="feather" @click="searchUsername(true)"><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
                        </div>
                    </div>
                </div>
				

			</div>

			<!-- Let's now handle the dynamic part. -->
			<div class="result">

				<!-- Let's report the Error Message(if any), and the Loading Spinner if needed. -->
				<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

                <!-- ------------------------ FOLLOWINGS PART ------------------------  -->                
                <!-- If instead, it is all ok, Display a sort of card for each of the User Profiles in the followingsList -->
				
                <CardProfile 
                    v-if="!loading && flagFollow == false" 
                    v-for="u in followingsListProfiles" 
                    :key="u.fixedUsername"
                    :user="u"> 
                </CardProfile>
                
                
                <!-- ------------------------ FOLLOWERS PART ------------------------  -->                
                <!-- In alternative, Display a sort of card for each of the User Profiles in the followersList -->
                <CardProfile 
                    v-if="!loading && flagFollow == true" 
                    v-for="u in followersListProfiles" 
                    :key="u.fixedUsername"
                    :user="u"> 
                </CardProfile>

                <LoadingSpinner v-if="loading"></LoadingSpinner>

			</div>
	</div>
</template>

<!-- Declaration of the style(scoped) to use. -->
<style scoped>
	@import '../assets/follow.css';
</style>
