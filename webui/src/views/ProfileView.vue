<!-- Starting of the actual Search Page Handling. -->
<script>
import ErrorMsg from '../components/ErrorMsg.vue'
import LoadingSpinner from '../components/LoadingSpinner.vue'
import MyProfileCard from '../components/MyProfileCard.vue'
import PhotoCard from '../components/PhotoCard.vue'

// Declaration of the export set.
export default {

	components: {
		ErrorMsg,
		LoadingSpinner,
		MyProfileCard,
		PhotoCard,
	},

	// Describing what are the Return variables.
	data: function() {
		return {

			// Initializing the two errormessage and loading variables.
			errormsg: "",
			loading: false,

			// Retrieving from the Cache the Username and the Bearer Authenticaiton Token.
			// Notice that here I need also to keep track of the username that is Logged since anyone(that is not banned, of course) can view others profiles.
            BearerToken: localStorage.getItem('BearerToken'),
			usernameLogged: localStorage.getItem('Username'),
			username: localStorage.getItem('Username') == localStorage.getItem('usernameProfileToView') ? localStorage.getItem('Username') : localStorage.getItem('usernameProfileToView'),

			// Initializing flag that allows to see whether the user that is accessing the page is the actual user owner or not.
			userOwnerFlag: localStorage.getItem('Username') == localStorage.getItem('usernameProfileToView') ? true : false,

			// Initializing variable for handling the UserProfile retrieval.
			userProfile: { fixedUsername: "", username: "", photoProfile: "", biography: "", dateOfCreation: "", numberOfPhotos: 0, numberFollowers: 0, numberFollowing: 0, name: "", surname: "", dateOfBirth: "", email: "", nationality: "", gender: ""},
			colorBackground: '',
			colorPosts: '',

			// Initializing a list that will handle the links to the photos.
			photoListLinks: [],

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
				let responseProfile = await this.$axios.get(`/users/${this.username}`, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					}
				})

				responseProfile.data.boolFollowing = false;
				responseProfile.data.boolFollower = false;
				responseProfile.data.boolBanned = false;
				this.colorBackground = responseProfile.data.gender == "male" ? '#c2e9fc' : responseProfile.data.gender == "female" ? '#fbd3f0' : '#cff6cc',
				this.colorPosts = '#ffffff'
				
				// Let's add up to the "userProfiles" array the response of the profile. Note that it will be an array with only this element.
				this.userProfile = responseProfile.data;
				// this.userProfile.photoProfile = (this.userProfile.photoProfile.split('-')[2]).split('.')[0]
				// this.userProfile.photoProfile = "../../../../tmp/u1-photo-0.jpg";

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

		// getPhotoLinks Function: It retrieves the whole photolist(os links) of the username.
		async getPhotoLinks() {

			// Re-initializing variables to their default value.
            this.errormsg = "";
            this.loading = true;

			try{

				// Retrieving the Profile from the Back-end.
				let responsePhotoList = await this.$axios.get(`/users/${this.username}/photos/`, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					}
				})

				this.photoListLinks = responsePhotoList.data;

			} catch (e) {

				// If an error is encountered, display it!
				if (e.response && e.response.status === 400) {
					this.errormsg = "Request error, please Login before doing some action or ask to get the list of photos of a valid user." + e.toString();
                } else if (e.response && e.response.status === 403) {
                    this.errormsg = "An Unauthorized Action has been blocked. You are not allowed to do this action because you are not the profile's owner." + e.toString();
                } else if (e.response && e.response.status === 204) {
                    this.errormsg = "In the Internal DB there is not anymore the content you have asked." + e.toString();
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later." + e.toString();
                } else {
                    this.errormsg = "Please Login before with an Authorized profile to view this page. " + e.toString();
                }
			}

			// Setting again the Loading flag to false.
            this.loading = false;
		},
	},
	mounted() {
		this.getUserProfile()
		this.getPhotoLinks()
	},

	watch: {
    	'$route.params.username': function(newUsername) {
			this.username= newUsername;
			this.userOwnerFlag = true;
			// localStorage.setItem('usernameProfileToView', this.newUsername);
			// localStorage.setItem('Username', this.newUsername);
			// alert(localStorage.getItem('usernameProfileToView'), localStorage.getItem('Username'))
	  		this.getUserProfile()
	  		this.getPhotoLinks()
    	},
		'$route.params.usernameLogged': function(newUsername) {
			this.usernameLogged= newUsername;
    	},
  	},
}
</script>

<!-- Actual Page for handling the page setting. -->
<template>

	<div>
			<!-- Let's handle first the upper part that will be the static one. -->
			<div>
				<h3 class="h3"> <b> User Logged:</b> {{ this.usernameLogged }} </h3>
			</div>

			<h1 class="h1"> {{ this.userProfile.username + "'s"}} PERSONAL PROFILE</h1>
			<img src="./img/wasa-logo.png" alt="" class="img">

			<!-- Let's now handle the dynamic part. -->
			<div class="result" v-if="!loading">

				<!-- If instead, it is all ok, Display a sort of card for each of the User Profiles(Depending on we are asking the whole list or just one). -->
				<MyProfileCard  
					:user=this.userProfile 
					:userOwnerFlag = "this.userOwnerFlag"
					:style="{backgroundColor: this.colorBackground}" 
				></MyProfileCard>
			</div>

			<!-- Divider Profile-Photos -->
			<br><br><br><br><br><br><br><br><br><br><br><br>
			<div class="divider">
				<span></span><span>Posts</span><span></span>
			</div>

			<!-- Let's report the Error Message(if any), and the Loading Spinner if needed. -->
			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
			<LoadingSpinner v-if="loading"></LoadingSpinner>

			<!-- {{ this.userProfile }} -->
			<!-- If instead, it is all ok, Display a sort of card for each of the User Photo(Depending on we are asking the whole list or just one). -->
			<div class="photoList" v-if="!loading" > 
				<PhotoCard  
					v-for="p in photoListLinks" 
					:key="p.photoid"
					style="background-color:papayawhip; margin-top:80px;"
					:photo="p"
					:userOwnerFlag = "!this.userOwnerFlag"
					:usernameLogged = "this.usernameLogged"
					:userProfile = "this.userProfile"
					:photoListCurrent = "this.photoListLinks"
					@refreshProfile = "this.userProfile = $event"
					@refreshPhotos = "this.photoListLinks = $event"
				></PhotoCard>
				<!-- :photo="p" -->
			</div>

		</div>

</template>

<!-- Declaration of the style(scoped) to use. -->
<style scoped>
	@import '../assets/personalProfile.css';
</style>
