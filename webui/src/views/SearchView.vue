<!-- Starting of the actual Search Page Handling. -->
<script>

// Declaration of the export set.
export default {

	// Describing what are the Return variables.
	data: function() {
		return {

			// Initializing the two errormessage and loading variables.
			errormsg: null,
			loading: false,

			// Retrieving from the Cache the Username and the Bearer Authenticaiton Token.
            username: localStorage.getItem('Username'),
            BearerToken: localStorage.getItem('BearerToken'),

			// Initializing two arrays that will handle the list of users and their profiles.
			users: [],
			usersProfiles: [],									// userProfile: { fixedUsername: "", username: "", photoProfile: "", biography: "", dateOfCreation: "", numberOfPhotos: 0, numberFollowers: 0, numberFollowing: 0, name: "", surname: "", dateOfBirth: "", email: "", nationality: "", gender: ""}
			
			// Initializing two variables that will be used to Handle the Specific Search for a User.
			usernameToSearch: "",
			usernameToSearchBool: true,
		}
	},

	// Declaration of the methods that will be used.
	methods: {
		
        // GetUsers Function: It fills the "users" array with the usernames present in the DB.
		async getUsers() {

			// Re-initializing variables to their default value.
			this.errormsg = null;
			this.loading = true;

			this.users = [];
			this.usersProfiles = [];

			this.usernameToSearchBool = true;
			
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

				// Sorting the list of Profiles (newest to oldest) w.r.t. the dateOfCreation.
				this.usersProfiles.sort(function(a,b){
					return a.dateOfCreation - b.dateOfCreation;
				})

			} catch (e) {

				// If an error is encountered, display it!
				this.errormsg = e.toString();
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

				// Add the profile retrieved to the "usersProfiles" array.
				this.usersProfiles.push(responseProfile.data);

			} catch (e) {

				// If an error is encountered, display it!
				this.errormsg = e.toString();
			}
		},

		// SerchUsername: It will search for whether the Username inserted in the input is present.
		async searchUsername() {

			// Re-initializing variables to their default value.
			this.errormsg = null;
			this.loading = true;

			this.users = [];
			this.usersProfiles = [];

			this.usernameToSearchBool = true;
			
			try{

				// Let's retrieve the Profile of the Username we are searching for.
				let responseProfile = await this.$axios.get("/users/"+this.usernameToSearch, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					}
				})

				// Let's add up to the "userProfiles" array the response of the profile. Note that it will be an array with only this element.
				this.usersProfiles.push(responseProfile.data);

			} catch (e) {

				// If an error is encountered, display it! Moreover, here, put the "usernameToSearchBool" flag to false.
				this.usernameToSearchBool = false;
				this.errormsg = e.toString();
			}
	
			// Once the entire operation has finished, re-set the "loading" flag to false, in such a way to continue.
			this.loading = false;
			// this.usernameToSearchBool = true;
		},
	},
}
</script>

<!-- Actual Page for handling the page setting. -->
<template>

	<div>
			<!-- Let's handle first the upper part that will be the static one. -->
			<h1 class="h1">WASA Photo SEARCH</h1>

			<div class="topMenu">

				<!-- "Users List" Button -->
				<div class="topMenuButtons">
					<button type="login-button" class="btn btn-primary btn-block btn-large" v-if="!loading" @click="getUsers"> Users List </button>
				</div>

				<!-- WASA Photo Icon -->
				<div class="topMenuColumn">
					<img src="./img/wasa-logo.png" alt="" class="img">
				</div>

				<!-- "Search Username Field" -->
				<div class="topMenuButtons">
					<input type="text" id="usernameToSearch" v-model="usernameToSearch" placeholder="Search Username..." class="form-control">
					<button type="login-button" class="searchButton" v-if="!loading" @click="searchUsername"> 
						<!-- <img src="/feather-sprite-v4.29.0.svg#search"> -->
						<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
					</button>
				</div>

			</div>


			<!-- Let's now handle the dynamic part. -->
			<div class="result">

				<!-- Let's report the Error Message(if any), and the Loading Spinner if needed. -->
				<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
				<LoadingSpinner v-if="loading"></LoadingSpinner>

				<!-- If the Username to search was not found, report the error. -->
				<div class="card" v-if="usernameToSearchBool === false">
					<div class="card-body">
						<p>No User present in the Database with the {{ this.usernameToSearch }} username.</p>
					</div>
				</div>

				<!-- If instead, it is all ok, Display a sort of card for each of the User Profiles(Depending on we are asking the whole list or just one). -->
				<div class="card" v-if="!loading" v-for="u in usersProfiles">

					<div class="card-header">
						<b> USER {{ u.fixedUsername }} </b>
					</div>
					<div class="card-body">
						<p class="card-text">

							Photo: {{ u.photoProfile}} <br/>
							Username: {{ u.username }}<br/>
							Name: {{ u.name }} <br/>
							Biography: {{ u.biography }} <br/>
							DateOfCreation: {{ u.dateOfCreation}}

						</p>
					</div>
				</div>
			</div>
	</div>
</template>

<!-- Declaration of the style(scoped) to use. -->
<style scoped>
@import url(https://fonts.googleapis.com/css?family=Open+Sans);

.card {
	margin-bottom: 20px;
	font-size: 20px;
  	font-family: sans-serif;
}

.img {
	display: block;
  	margin-left: auto;
  	margin-right: auto;
  	width: 50%;
    width: 200px;
    height: auto;
}

.topMenu{
	display: block;
  	margin-left: auto;
  	margin-right: auto;
	margin-top: 10px;
	margin-bottom: 10px;
  	width: 50%;
	width: 720px;
}

.h1 {
	display: block;
  	margin-left: auto;
  	margin-right: auto;
	margin-top: 60px;
}

.topMenuColumn {
  float: left;
  width: 33.33%;
}

.topMenuButtons{
	margin-top: 78px;
	margin-left: 10px;
	margin-right: 10px;
	float: left;
  	width: 30%;
}


.result{
	display: block;
  	margin-left: auto;
  	margin-right: auto;
	margin-top: 300px;
}

.btn { display: inline-block; *display: inline; *zoom: 1;   font-family: sans-serif; padding: 4px 10px 4px; margin-bottom: 0; font-size: 13px; line-height: 18px; color: #333333; text-align: center;text-shadow: 0 1px 1px rgba(255, 255, 255, 0.75); vertical-align: middle; background-color: #f5f5f5; background-image: -moz-linear-gradient(top, #ffffff, #e6e6e6); background-image: -ms-linear-gradient(top, #ffffff, #e6e6e6); background-image: -webkit-gradient(linear, 0 0, 0 100%, from(#ffffff), to(#e6e6e6)); background-image: -webkit-linear-gradient(top, #ffffff, #e6e6e6); background-image: -o-linear-gradient(top, #ffffff, #e6e6e6); background-image: linear-gradient(top, #ffffff, #e6e6e6); background-repeat: repeat-x; filter: progid:dximagetransform.microsoft.gradient(startColorstr=#ffffff, endColorstr=#e6e6e6, GradientType=0); border-color: #e6e6e6 #e6e6e6 #e6e6e6; border-color: rgba(0, 0, 0, 0.1) rgba(0, 0, 0, 0.1) rgba(0, 0, 0, 0.25); border: 1px solid #e6e6e6; -webkit-border-radius: 4px; -moz-border-radius: 4px; border-radius: 4px; -webkit-box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.2), 0 1px 2px rgba(0, 0, 0, 0.05); -moz-box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.2), 0 1px 2px rgba(0, 0, 0, 0.05); box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.2), 0 1px 2px rgba(0, 0, 0, 0.05); cursor: pointer; *margin-left: .3em; }
.btn:hover, .btn:active, .btn.active, .btn.disabled, .btn[disabled] { background-color: #e6e6e6; }
.btn-large { padding: 9px 14px; font-size: 20px; line-height: normal; -webkit-border-radius: 5px; -moz-border-radius: 5px; border-radius: 5px; }
.btn:hover { color: #333333; text-decoration: none; background-color: #e6e6e6; background-position: 0 -15px; -webkit-transition: background-position 0.1s linear; -moz-transition: background-position 0.1s linear; -ms-transition: background-position 0.1s linear; -o-transition: background-position 0.1s linear; transition: background-position 0.1s linear; }
.btn-primary, .btn-primary:hover { text-shadow: 0 -1px 0 rgba(0, 0, 0, 0.25); color: #ffffff; }
.btn-primary.active { color: rgba(255, 255, 255, 0.75); }
.btn-primary { background-color: #4a77d4; background-image: -moz-linear-gradient(top, #6eb6de, #4a77d4); background-image: -ms-linear-gradient(top, #6eb6de, #4a77d4); background-image: -webkit-gradient(linear, 0 0, 0 100%, from(#6eb6de), to(#4a77d4)); background-image: -webkit-linear-gradient(top, #6eb6de, #4a77d4); background-image: -o-linear-gradient(top, #6eb6de, #4a77d4); background-image: linear-gradient(top, #6eb6de, #4a77d4); background-repeat: repeat-x; filter: progid:dximagetransform.microsoft.gradient(startColorstr=#6eb6de, endColorstr=#4a77d4, GradientType=0);  border: 1px solid #3762bc; text-shadow: 1px 1px 1px rgba(0,0,0,0.4); box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.2), 0 1px 2px rgba(0, 0, 0, 0.5); }
.btn-primary:hover, .btn-primary:active, .btn-primary.active, .btn-primary.disabled, .btn-primary[disabled] { filter: none; background-color: #4a77d4; }
.btn-block { width: 100%; display:block; }

/* This CSS is dedicated to the input string for the "Username" */
input { 
  width: 100%; 
  margin-bottom: 10px; 
  background: rgb(199, 246, 255); 
  border: none;
  outline: none;
  padding: 10px;
  font-size: 15px;
  font-family: sans-serif;
  border: 1px solid rgba(31, 86, 135, 0.3);
  border-radius: 4px;
  box-shadow: inset 0 -5px 45px rgba(12, 125, 123, 0.2), 0 1px 1px rgba(10, 131, 161, 0.2);
  -webkit-transition: box-shadow .5s ease;
  -moz-transition: box-shadow .5s ease;
  -o-transition: box-shadow .5s ease;
  -ms-transition: box-shadow .5s ease;
  transition: box-shadow .5s ease;
}

.form-control {
	width: 0%;
	float: left;
  	width: 85%;
}

.searchButton{
	width: 0%;
	float: left;
  	width: 15%;
}

* { -webkit-box-sizing:border-box; -moz-box-sizing:border-box; -ms-box-sizing:border-box; -o-box-sizing:border-box; box-sizing:border-box; }


</style>
