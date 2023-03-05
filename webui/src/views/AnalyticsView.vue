<!-- Starting of the actual Search Page Handling. -->
<script>

import ErrorMsg from '../components/ErrorMsg.vue'
import LoadingSpinner from '../components/LoadingSpinner.vue'


// Declaration of the export set.
export default {

	components: {
		ErrorMsg,
		LoadingSpinner,
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
	methods: {},
}
</script>



<!-- Actual Page for handling the page setting. -->
<template>

	<div>
			<!-- Let's handle first the upper part that will be the static one. -->
			<h1 class="h1">{{ username }}'s ANALYTICS</h1>

			<div class="topMenu">

				<!-- WASA Photo Icon -->
                <div class="topMenuButtons"></div>
				<div class="topMenuColumn">
					<img src="./img/wasa-logo.png" alt="" class="img">
				</div>
				<div class="topMenuButtons"></div>

            </div>

			<div>
				<!-- Let's report the Error Message(if any), and the Loading Spinner if needed. -->
				<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
				<LoadingSpinner v-if="loading"></LoadingSpinner>
			</div>
        
    </div><!-- /.container -->
</template>

<!-- Declaration of the style(scoped) to use. -->
<style scoped>
	@import '../assets/updateProfile.css';
</style>
