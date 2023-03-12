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
		}
	},

	// Declaration of the methods that will be used.
	methods: {
		
        

	},

}
</script>

<!-- Actual Page for handling the page setting. -->
<template>

	<div>
			<!-- Let's handle first the upper part that will be the static one. -->
			<h1 class="h1">FUTURE OF WORK</h1>

			<div class="topMenu">

				<!-- "Users List" Button -->
				<div class="topMenuButtons"></div>

				<!-- WASA Photo Icon -->
				<div class="topMenuColumn">
					<img src="./img/wasa-logo.png" alt="" class="img">
				</div>

				<!-- "Search Username Field" -->
				<div class="topMenuButtons"></div>

			</div>


			<!-- Let's now handle the dynamic part. -->
			<div class="result">

				<!-- Let's report the Error Message(if any), and the Loading Spinner if needed. -->
				<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
				<LoadingSpinner v-if="loading"></LoadingSpinner>

				<!-- If instead, it is all ok, Display a sort of card for each of the User Profiles(Depending on we are asking the whole list or just one). -->
				<CardProfile 
					v-if="!loading" 
					v-for="u in usersProfiles" 
					:key="u.fixedUsername"
				    :user="u" 
					@refreshFollowing = "u.boolFollowing = $event"
					@refreshNumberFollowers = "u.numberFollowers = $event"
					@refreshBan = "u.boolBanned = $event"
				></CardProfile>
			</div>
	</div>
</template>

<!-- Declaration of the style(scoped) to use. -->
<style scoped>
	@import '../assets/search.css';
</style>
