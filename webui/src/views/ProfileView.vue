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

				responseProfile.data.boolFollowing = false;
				responseProfile.data.boolFollower = false;
				responseProfile.data.boolBanned = false;
				this.colorBackground = responseProfile.data.gender == "male" ? '#c2e9fc' : responseProfile.data.gender == "female" ? '#fbd3f0' : '#cff6cc',
				this.colorPosts = '#ffffff'
				
				// Let's add up to the "userProfiles" array the response of the profile. Note that it will be an array with only this element.
				this.userProfile = responseProfile.data;
				this.userProfile.photoProfile = "../../../tmp/u1-photo-0-photo-profile.jpg";

			} catch (e) {

				// If an error is encountered, display it!
				this.errormsg = e.toString();
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
				this.errormsg = e.toString();
			}

			// Setting again the Loading flag to false.
            this.loading = false;
		},

	},
	mounted() {
		this.getUserProfile()
		this.getPhotoLinks()
	}
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
			<div class="result">

				<!-- Let's report the Error Message(if any), and the Loading Spinner if needed. -->
				<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
				<LoadingSpinner v-if="loading"></LoadingSpinner>

				<!-- If instead, it is all ok, Display a sort of card for each of the User Profiles(Depending on we are asking the whole list or just one). -->
				<MyProfileCard 
					v-if="!loading" 
					:user=this.userProfile 
					:style="{backgroundColor: this.colorBackground}" 
					:userOwnerFlag = "this.userOwnerFlag"
				></MyProfileCard>
			</div>

			<!-- Divider Profile-Photos -->
			<br><br><br><br><br><br><br><br>
			<div class="divider">
				<span></span><span>Posts</span><span></span>
			</div>

			<!-- Let's report the Error Message(if any), and the Loading Spinner if needed. -->
			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
			<LoadingSpinner v-if="loading"></LoadingSpinner>

			<!-- {{ this.photoListLinks }} -->
			<!-- If instead, it is all ok, Display a sort of card for each of the User Photo(Depending on we are asking the whole list or just one). -->
			<div class="photoList"> 
				<PhotoCard 
					v-if="!loading" 
					v-for="p in photoListLinks" 
					:key="p.photoid"
					:style="{backgroundColor: this.colorPosts}" style="background-color:papayawhip; margin-top:80px;"
					:photo="p"
					:userOwnerFlag = "!this.userOwnerFlag"
					:usernameLogged = "this.usernameLogged"
					:numberOfPhotos = "this.userProfile.numberOfPhotos"
					:photoListCurrent = "this.photoListLinks"
					@refreshNumberPhotos = "this.userProfile.numberOfPhotos = $event"
					@refreshPhotos = "this.photoListLinks = $event"
				></PhotoCard>
			</div>

		</div>

</template>

<!-- Declaration of the style(scoped) to use. -->
<style scoped>
	@import '../assets/personalProfile.css';
</style>
