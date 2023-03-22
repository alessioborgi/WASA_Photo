<!-- Starting of the actual Search Page Handling. -->
<script>

import ErrorMsg from '../components/ErrorMsg.vue'

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

			// Initializing variable for handling the UserProfile retrieval.
			userProfile: { fixedUsername: "", username: "", photoProfile: "", biography: "", dateOfCreation: "", numberOfPhotos: 0, numberFollowers: 0, numberFollowing: 0, name: "", surname: "", dateOfBirth: "", email: "", nationality: "", gender: ""},
			
            // Initializing the variable that will take the new username.
            newUsername: "",

        }
	},

	// Declaration of the methods that will be used.
	methods: {

        // GetUserProfile Function: It retrieves the whole profile of the Logged username.
        async getUserProfile() {

            // Re-initializing variables to their default value.
            this.errormsg = "";
            this.loading = true;

            try{

                // Retrieving the Profile from the Back-end.
                let responseProfile = await this.$axios.get("/users/"+this.username, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("BearerToken")
                    }
                })
                
                // Let's add up to the "userProfiles" array the response of the profile. Note that it will be an array with only this element.
                this.userProfile = responseProfile.data;

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

            // Setting again the Loading flag to false.
            this.loading = false;
        },

        // setUsername: This method is used for changing the Username. 
        async setUsername() {

            // Re-initializing variables to their default value.
            this.errormsg = "";
            this.loading = true;

            try {
                
                // In the case the result is positive, we post the username received to the GO page.

                await this.$axios.patch(`/users/${this.username}`, { username: this.newUsername}, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("BearerToken")
                    }
                })

                // Setting the new username received as the new username saved in the local cache.
                localStorage.setItem('Username', this.newUsername),
                localStorage.setItem('usernameProfileToView', this.newUsername)
                this.username = this.newUsername;
                this.userProfile.username = this.newUsername;
                                
                this.$emit('refreshProfile', this.username);

                // Re-addressing the page to the personal profile page of a user.
                this.$router.push({ path: `/users/${this.username}` })
                this.newUsername = "";


            } catch (e) {

                // In case of error, retrieve it.
                if (e.response && e.response.status === 400) {
					this.errormsg = "Request error, please Login before doing some action or ask to update the username of a valid user." + e.toString();
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

            // Setting again the Loading flag to false.
            this.loading = false;
        },

	},

    mounted() {
        this.getUserProfile();
    }
}
</script>



<!-- Actual Page for handling the page setting. -->
<template>

	<div>
			<!-- Let's handle first the upper part that will be the static one. -->
			<h1 class="h1">{{ username }}'s USERNAME UPDATE</h1>

			<div class="topMenu">

				<!-- WASA Photo Icon -->
                <div class="topMenuButtons"></div>
				<div class="topMenuColumn">
					<img src="./img/wasa-logo.png" alt="" class="img">
				</div>
				<div class="topMenuButtons"></div>

            </div>
            <div class="formUpdate">

                <form class="well form-horizontal" action=" " method="post"  id="contact_form">
                <fieldset>

                    <!-- Username -->
                    <div class="form-group">
                        <label class="col-md-4 control-label"><h3><b>Username</b></h3></label>  
                        <div class="col-md-4 inputGroupContainer">
                            <div class="input-group">
                                <span class="input-group-addon"><i class="glyphicon glyphicon-user"></i></span>
                                <input  name="username" :placeholder=this.userProfile.username v-model=this.newUsername class="form-control"  type="text">
                            </div>
                        </div>
                    </div>

                   

                    <!-- Send Button -->
                    <div class="form-group2">
					    <button type="login-button" class="btn btn-primary btn-block btn-large" @click="setUsername"> Update Username </button>
				    </div>

                    <div>
                        <!-- Let's report the Error Message(if any), and the Loading Spinner if needed. -->
			            <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
                    </div>

                </fieldset>
            </form>
        </div>
    </div><!-- /.container -->
</template>

<!-- Declaration of the style(scoped) to use. -->
<style scoped>
	@import '../assets/newUsername.css';
</style>
