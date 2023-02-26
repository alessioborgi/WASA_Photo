<!-- Starting of the actual Search Page Handling. -->
<script>

import ErrorMsg from '../components/ErrorMsg.vue'
import LoadingSpinner from '../components/LoadingSpinner.vue'
import MyProfileCard from '../components/MyProfileCard.vue'
import StreamPhotoCard from '../components/StreamPhotoCard.vue'

// Declaration of the export set.
export default {

	components: {
		ErrorMsg,
		LoadingSpinner,
		MyProfileCard,
		StreamPhotoCard,
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

			// Initializing a list that will handle the links to the photos.
			streamListLinks: [],
		}
	},

	// Declaration of the methods that will be used.
	methods: {
		
		// getMyStream Function: It retrieves the whole photolist(os links) of the Stream.
		async getMyStream() {

			// Re-initializing variables to their default value.
            this.errormsg = "";
            this.loading = true;

			try{

				// Retrieving the Stream from the Back-end.
                // /users/:username/myStream/
				let responsePhotoList = await this.$axios.get(`/users/${this.username}/myStream/`, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					}
				})

				this.streamListLinks = responsePhotoList.data;

			} catch (e) {

				// If an error is encountered, display it!
				this.errormsg = e.toString();
			}

			// Setting again the Loading flag to false.
            this.loading = false;
		},

	},
	mounted() {
		this.getMyStream()
	}
}
</script>

<!-- Actual Page for handling the page setting. -->
<template>

	<div>
			<!-- Let's handle first the upper part that will be the static one. -->
			<h1 class="h1"> {{ this.username + "'s"}} Stream</h1>
			<img src="./img/wasa-logo.png" alt="" class="img">

			<!-- Divider Profile-Photos -->
			<br><br><br><br><br><br><br><br>
			<div class="divider">
				<span></span><span>Posts</span><span></span>
			</div>

			<!-- Let's report the Error Message(if any), and the Loading Spinner if needed. -->
			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
			<LoadingSpinner v-if="loading"></LoadingSpinner>

			{{ this.streamListLinks }}
			<!-- If instead, it is all ok, Display a sort of card for each of the User Photo(Depending on we are asking the whole list or just one). -->
			<div class="photoList"> 
				<StreamPhotoCard v-if="!loading" v-for="p in streamListLinks" :style="{backgroundColor: this.colorPosts}" style="margin-top:80px;"
					:photo="p"
				></StreamPhotoCard>
			</div>

		</div>

</template>

<!-- Declaration of the style(scoped) to use. -->
<style scoped>
	@import '../assets/personalProfile.css';
</style>
