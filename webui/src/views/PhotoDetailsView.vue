<!-- Starting of the actual Search Page Handling. -->
<script>

import ErrorMsg from '../components/ErrorMsg.vue'
import LoadingSpinner from '../components/LoadingSpinner.vue'
import PhotoCardDetail from '../components/PhotoCardDetail.vue'
import Comment from '../components/Comment.vue'
import Like from '../components/Like.vue'


// Declaration of the export set.
export default {

	components: {
		ErrorMsg,
		LoadingSpinner,
		PhotoCardDetail,
		Comment,
		Like,
	},

	// Describing what are the Return variables.
	data: function() {
		return {

			// Initializing the two errormessage and loading variables.
			errormsg: "",
			loading: false,

			// Retrieving from the Cache the Username and the Bearer Authenticaiton Token.
            BearerToken: localStorage.getItem('BearerToken'),
			usernameLogged: localStorage.getItem('Username'),
			username: localStorage.getItem('Username') == localStorage.getItem('usernameProfileToView') ? localStorage.getItem('Username') : localStorage.getItem('usernameProfileToView'),

			// Initializing flag that allows to see whether the user that is accessing the page is the actual user owner or not.
			userOwnerFlag: localStorage.getItem('Username') == localStorage.getItem('usernameProfileToView') ? true : false,

			// Retrieving also the photo information from the idPhoto saved in the Local Cache.
			idPhoto: localStorage.getItem('idPhoto'),

			// Initializing the Photo variable.
			photoData: {Photoid: 0, FixedUsername:"", Filename:"", UploadDate:"", Phrase:"", NumberLikes:0, NumberComments:0},
			colorPosts: '#ffffff',

			// Initializing a flag indicating whether to update:
            //   - likesList (value: true) or 
            //   - commentsList (value: false)
            flagCommentsLikes: false,

			// Initializing the Comments and the Like List.
			commentsList: [],
			likesList: [],

			// Initializing the two lists that will keep track of both the users and the usersProfiles.
			users: [],
			usersProfiles: [],
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
				if (e.response && e.response.status === 400) {
					this.errormsg = "Request error, please Login before doing some action or ask to get the photo of a valid user." + e.toString();
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

		// getComments: It returns the list of comments of a determinate photo.
		async getComments(){

			// Re-initializing variables to their default value.
			this.errormsg = "";
			this.loading = true;

			this.commentsList = [];

			// Set the flagCommentsLikes to false (meaning that I need to work on commentsList)
			this.flagCommentsLikes = false;

			// ----- Getting Comments. -----
			try {

				// Getting the list of Comments from the Back-End.
				// /users/:username/photos/:photoid/comments/.
				let response = await this.$axios.get(`/users/${this.username}/photos/${this.photoData.photoid}/comments/`, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					}
				})

				// Saving the response in the "users" array.
				this.commentsList = response.data;

				if (this.commentsList.length > 0){
					// Sorting the list of Profiles (newest to oldest) w.r.t. the dateOfCreation.
					this.commentsList.sort(function(a,b){

						return new Date(b.UploadDate) - new Date(a.UploadDate);
					})
				} else {
					this.errormsg = "There are no Comments to the Photo yet!";
				}

			} catch (e) {

				// If an error is encountered, display it!
				if (e.response && e.response.status === 400) {
					this.errormsg = "Request error, please Login before doing some action or ask to get the list of comments of a valid photo." + e.toString();
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


		// getLikes: It returns the list of likes of a determinate photo.
		async getLikes(){

			// Re-initializing variables to their default value.
			this.errormsg = "";
			this.loading = true;

			this.likesList = [];

			// Set the flagCommentsLikes to true (meaning that I need to work on likesList)
			this.flagCommentsLikes = true;

			// ----- Getting Likes. -----
			try {

				// Getting the list of Comments from the Back-End.
				// /users/:username/photos/:photoid/likes/
				let response = await this.$axios.get(`/users/${this.username}/photos/${this.photoData.photoid}/likes/`, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					}
				})

				// Saving the response in the "users" array.
				this.likesList = response.data;

				if (this.likesList.length == 0){
					this.errormsg = "There are no Likes to the Photo yet!";
				}

			} catch (e) {

				// If an error is encountered, display it!
				if (e.response && e.response.status === 400) {
					this.errormsg = "Request error, please Login before doing some action or ask to get the list of likes of a valid photo." + e.toString();
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

		// GetUsers Function: It fills the "users" array with the usernames present in the DB.
		async getUsers() {

			// Re-initializing variables to their default value.
			this.errormsg = "";
			this.loading = true;

			this.users = [];
			this.usersProfiles = [];
						
			try {

				// Getting the list of Users from the Back-End.
				let response = await this.$axios.get("/users/", {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					}
				})

				// Saving the response in the "users" array.
				this.users = response.data;

				// Retrieving for every username, its Profile.
				for (let i = 0; i < this.users.length; i++) {

					this.getUserProfile(i)
				}

				await this.getUserProfilesLogged();

			} catch (e) {

				// If an error is encountered, display it!
				if (e.response && e.response.status === 400) {
					this.errormsg = "Request error, please Login before doing some action or ask to get the list of users." + e.toString();
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

		// GetUserProfile Function: It retrieves the whole profile of a username.
		async getUserProfile(i) {

			try{

				// Retrieving the Profile from the Back-end.
				let responseProfile = await this.$axios.get("/users/"+this.users[i], {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					}
				})

				// Let's add up to the "userProfiles" array the response of the profile. Note that it will be an array with only this element.
				this.usersProfiles.push([responseProfile.data.username, responseProfile.data.fixedUsername]);

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
		},

		// GetUserProfile Function: It retrieves the whole profile of a username.
		async getUserProfilesLogged() {

			try{

				// Retrieving the Profile from the Back-end.
				let responseProfile = await this.$axios.get("/users/"+this.usernameLogged, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					}
				})

				// Let's add up to the "userProfiles" array the response of the profile. Note that it will be an array with only this element.
				this.usersProfiles.push([responseProfile.data.username, responseProfile.data.fixedUsername]);

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
		},

	},

	mounted() {
		this.getPhoto()
		this.getUsers()
	}
}
</script>

<!-- Actual Page for handling the page setting. -->
<template>

	<div>
		<!-- {{ this.photoData }} -->
			<!-- Let's handle first the upper part that will be the static one. -->
			<div>
				<h3 class="h3"> <b> User Logged:</b> {{ this.usernameLogged }} </h3>
			</div>
			
			<h1 class="h1">{{ username }}'s Photo</h1>

			<div class="topMenu">

                <!-- Followings Menu left-Part -->
                <div class="followingsMenu">
                    
                    <h2 class="h2" style="margin-left:80px; margin-top: 50px; margin-bottom: -50px;">LIKES</h2>

                    <!-- "Users List" Button -->
                    <div class="topMenuButtons">
                        <button type="login-button" class="btn btn-primary btn-block btn-large" 
							@click="getLikes">  
							Likes List 
						</button>
                    </div>

                    <!-- "Search Username Field" -->
                    
                </div>

				<!-- WASA Photo Icon -->
				<div class="topMenuColumn">
					<img src="./img/wasa-logo.png" alt="" class="img">
				</div>

                <!-- Followings Menu left-Part -->
                <div class="followingsMenu">
                    
                    <h2 class="h2" style="margin-left:530px; margin-top: 25px; margin-bottom: -50px;"
						>COMMENTS
					</h2>

                    <!-- "Users List" Button -->
                    <div class="topMenuButtons">
                        <button type="login-button" class="btn btn-primary btn-block btn-large" 
							@click="getComments">  
							Comments List 
						</button>
                    </div>

                </div>
			</div>
		</div>

		<!-- Photo Cards -->
		<div>
			<div class="result" style="margin-top: 300px;">
				<!-- If instead, it is all ok, Display a sort of card for each of the User Photo(Depending on we are asking the whole list or just one). -->
				<div class="photoList" v-if="!loading" > 
					<PhotoCardDetail 
						:photo="this.photoData"
						:style="{backgroundColor: this.colorPosts}" style="background-color:papayawhip; margin-top:80px;"
						:userOwnerFlag = "!this.userOwnerFlag"
						@refreshNumberComments = "this.photoData.numberComments = $event"
						@refreshCommentsList = "this.commentsList = $event"
					></PhotoCardDetail>
				</div>
			</div>
		</div>

		<!-- {{ this.usersProfiles }} -->
		<!-- Divider Profile-Photos -->
		<br><br><br><br><br><br><br><br>
		<div class="divider" style="margin-top: -50px;">
			<span></span><span>Likes & Comments</span><span></span>
		</div>

		<!-- Let's report the Error Message(if any), and the Loading Spinner if needed. -->
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<LoadingSpinner v-if="loading"></LoadingSpinner>

		<!-- {{ this.commentsList }} -->
		<!-- Comments List -->
		<div class="commentsList" v-if="!loading && flagCommentsLikes == false" >  
			<Comment 
				v-for="c in commentsList" 
				:key="c.fixedUsername"
				:style="{backgroundColor: this.colorPosts}" 
				style="background-color: #c0f3b2 ; margin-top:80px;"
				:comment="c"
				:userOwnerFlag = this.userOwnerFlag
				:photoid="this.photoData.photoid"
				:numberComments="this.photoData.numberComments"
				:usersProfiles="this.usersProfiles"
				@refreshNumberCommentsFromComment = "this.photoData.numberComments = $event"
				@refreshCommentsListFromComment = "this.commentsList = $event"
			></Comment>
		</div>

		<!-- {{ this.likesList }} -->
		<div class="likesList" v-if="!loading && flagCommentsLikes == true" >  
			<Like 
				v-for="l in likesList" 
				:key="l.fixedUsername"
				:style="{backgroundColor: this.colorPosts}" 
				style="background-color:#f3c3b2 ; margin-top:80px;"
				:like="l"
				:userOwnerFlag = !this.userOwnerFlag
			></Like>
		</div>


	
	

</template>

<!-- Declaration of the style(scoped) to use. -->
<style scoped>
	@import '../assets/PhotoDetailsView.css';
</style>
