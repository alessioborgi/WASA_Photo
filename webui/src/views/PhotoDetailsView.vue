<!-- Starting of the actual Search Page Handling. -->
<script>

import ErrorMsg from '../components/ErrorMsg.vue'
import LoadingSpinner from '../components/LoadingSpinner.vue'
import PhotoCardDetail from '../components/PhotoCardDetail.vue'

// Declaration of the export set.
export default {

	components: {
		ErrorMsg,
		LoadingSpinner,
		PhotoCardDetail,
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

			// Retrieving also the photo information from the idPhoto saved in the Local Cache.
			idPhoto: localStorage.getItem('idPhoto'),

			// Initializing flag that allows to see whether the user that is accessing the page is the actual user owner or not.
			userOwnerFlag: localStorage.getItem('Username') == localStorage.getItem('usernameProfileToView') ? true : false,

			// Initializing the Photo variable.
			photoData: {Photoid: 0, FixedUsername:"", Filename:"", UploadDate:"", Phrase:"", NumberLikes:0, NumberComments:0},
			colorPosts: '#ffffff',
		}
	},

	// Declaration of the methods that will be used.
	methods: {	

		// getPhoto Function: It retrieves the photoInformation given the IdPhoto.
		async getPhoto() {

			// Re-initializing variables to their default value.
			this.errormsg = "";
			this.loading = true;

			try{

				// Retrieving the Photo from the Back-end.
				// /users/:username/photos/:photoid
				let responsePhotoList = await this.$axios.get(`/users/${this.username}/photos/${this.idPhoto}`, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					}
				})

				this.photoData = responsePhotoList.data;
				this.photoData = {photoid: this.photoData.Photoid, fixedUsername:this.photoData.FixedUsername, filename:this.photoData.Filename, uploadDate:this.photoData.UploadDate, phrase:this.photoData.Phrase, numberLikes:this.photoData.NumberLikes, numberComments:this.photoData.NumberComments}

			} catch (e) {

				// If an error is encountered, display it!
				this.errormsg = e.toString();
			}

			// Setting again the Loading flag to false.
			this.loading = false;
		},
	},

	mounted() {
		this.getPhoto()
	}
}
</script>

<!-- Actual Page for handling the page setting. -->
<template>

	<div>
            {{ this.photoData }}
			<!-- Let's handle first the upper part that will be the static one. -->
			<h1 class="h1"> {{ this.username + "'s"}} Photo</h1>
			<img src="./img/wasa-logo.png" alt="" class="img">

			<!-- Let's now handle the dynamic part. -->
			
			<!-- Let's report the Error Message(if any), and the Loading Spinner if needed. -->
			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
			<LoadingSpinner v-if="loading"></LoadingSpinner>

			<!-- If instead, it is all ok, Display a sort of card for each of the User Photo(Depending on we are asking the whole list or just one). -->
			<div class="photoList"> 
				<PhotoCardDetail v-if="!loading" 
					:photo="this.photoData"
					:style="{backgroundColor: this.colorPosts}" style="background-color:papayawhip; margin-top:80px;"
					:userOwnerFlag = !this.userOwnerFlag
				></PhotoCardDetail>
			</div>

		</div>

</template>

<!-- Declaration of the style(scoped) to use. -->
<style scoped>
	@import '../assets/personalProfile.css';
</style>
