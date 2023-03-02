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
			// responseLikeList: [],
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

				// Saving the list of links in the variable.
				this.streamListLinks = responsePhotoList.data;

				// // Retrieving for every photo, its Likes.
				for (let i = 0; i < this.streamListLinks.length; i++) {

					this.getLikes(i)
				}

			} catch (e) {

				// If an error is encountered, display it!
				this.errormsg = e.toString();
			}

			// Setting again the Loading flag to false.
            this.loading = false;
		},

		// getLikes: It returns the list of likes of a determinate photo.
		async getLikes(i){

			// Re-initializing variables to their default value.
			this.errormsg = "";
			this.loading = true;

			// ----- Getting Likes. -----
			try {

				// Getting the list of Likes from the Back-End.
				// /users/:username/photos/:photoid/likes/
				let responseLikeList = await this.$axios.get(`/users/${this.streamListLinks[i].username}/photos/${this.streamListLinks[i].photoid}/likes/`, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					}
				})

				// Add the boolLike to the "streamListLinks" array. 
				this.streamListLinks[i].boolLike = responseLikeList.data.includes(this.username) ? true : false

				// Check whether it is true or not that it is Liking the Photo and changing, by consequence, the Like that is filled or not.
				if (this.streamListLinks[i].boolLike == true){
					this.streamListLinks[i].fillHeart = "red"
				} else {
					this.streamListLinks[i].fillHeart = "white"
				}

			} catch (e) {

				// If an error is encountered, display it!
				this.errormsg = e.toString();
			}

			// Once the entire operation has finished, re-set the "loading" flag to false, in such a way to continue.
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
			<h1 class="h1"> {{ this.username + "'s"}} MyStream</h1>
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
					@refreshLike = "p.boolLike = $event"
					@refreshLikeFill = "p.fillHeart = $event"
					@refreshLikeNumber = "p.numberLikes = $event"
					@refreshNumberComments = "p.numberComments = $event"
				></StreamPhotoCard>
			</div>

		</div>

</template>

<!-- Declaration of the style(scoped) to use. -->
<style scoped>
	@import '../assets/personalProfile.css';
</style>
