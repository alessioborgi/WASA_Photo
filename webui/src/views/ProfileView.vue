<!-- Starting of the actual Search Page Handling. -->
<script>

import ErrorMsg from '../components/ErrorMsg.vue'
import LoadingSpinner from '../components/LoadingSpinner.vue'
import MyProfileCard from '../components/MyProfileCard.vue'

// Declaration of the export set.
export default {

	components: {
		ErrorMsg,
		LoadingSpinner,
		MyProfileCard,
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
			colorBackground: '',

			// Initializing a list that will handle the links to the photos.
			photoListLinks: [],
			photoListRaw: [],
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
				
				// Let's add up to the "userProfiles" array the response of the profile. Note that it will be an array with only this element.
				this.userProfile = responseProfile.data;
				this.userProfile.photoProfile = "../../../service/api/photos/u1-photo-1.jpg";

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

				// Retrieving every photo, "/users/:username/photos/:photoid/view"		
				// for (let i = 0; i < this.photoListLinks; i++) {

				// 	try{

				// 		// Retrieving the Photo from the Back-end.
				// 		let responsePhoto = await this.$axios.get(`/users/${this.username}/photos/${this.photoListLinks[i].id}/view`, {
				// 			headers: {
				// 				Authorization: "Bearer " + localStorage.getItem("BearerToken")
				// 			}
				// 		})

				// 		// Let's add up to the "userProfiles" array the response of the profile. Note that it will be an array with only this element.
				// 		this.photoListRaw.push(responsePhoto.data);

				// 	} catch (e) {

				// 		// If an error is encountered, display it!
				// 		this.errormsg = e.toString();
				// 	}
				// }

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
			<h1 class="h1"> {{ this.userProfile.username + "'s"}} PERSONAL PROFILE</h1>
			<img src="./img/wasa-logo.png" alt="" class="img">

			<!-- Let's now handle the dynamic part. -->
			<div class="result">

				<!-- Let's report the Error Message(if any), and the Loading Spinner if needed. -->
				<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
				<LoadingSpinner v-if="loading"></LoadingSpinner>

				<!-- If instead, it is all ok, Display a sort of card for each of the User Profiles(Depending on we are asking the whole list or just one). -->
				<MyProfileCard v-if="!loading" :user=this.userProfile :style="{backgroundColor: this.colorBackground}" 
					@refreshProfile = "this.userProfile.username = $event"
				></MyProfileCard>
			</div>

			<!-- Divider Profile-Photos -->
			<div class="divider">
				<span></span><span>Posts</span><span></span>
			</div>

			<div class="photos">

			<!-- Let's report the Error Message(if any), and the Loading Spinner if needed. -->
			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
			<LoadingSpinner v-if="loading"></LoadingSpinner>

			<!-- If instead, it is all ok, Display a sort of card for each of the User Profiles(Depending on we are asking the whole list or just one). -->
			<div v-if="!loading" v-for="link in photoListLinks"> 
				<img :src=link alt="Image" />
			</div>

			</div>

			
	</div>
</template>

<!-- Declaration of the style(scoped) to use. -->
<style scoped>
	@import '../assets/personalProfile.css';
</style>
