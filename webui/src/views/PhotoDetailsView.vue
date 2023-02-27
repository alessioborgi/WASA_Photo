<!-- Starting of the actual Search Page Handling. -->
<script>

import ErrorMsg from '../components/ErrorMsg.vue'
import LoadingSpinner from '../components/LoadingSpinner.vue'
import PhotoCard from '../components/PhotoCard.vue'
// import EventBus from '../event-bus'

// Declaration of the export set.
export default {

	components: {
		ErrorMsg,
		LoadingSpinner,
		PhotoCard,
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
			colorPosts: '',
            photo: null,
		}
	},

	// Declaration of the methods that will be used.
	methods: {	

	},
	// mounted() {
    //     EventBus.$on(photo, function (payLoad) {
    //         this.photo = payLoad;
    //     });
	// }
}
</script>

<!-- Actual Page for handling the page setting. -->
<template>

	<div>
            {{ this.photo }}
			<!-- Let's handle first the upper part that will be the static one. -->
			<h1 class="h1"> {{ this.username + "'s"}} Photo</h1>
			<img src="./img/wasa-logo.png" alt="" class="img">

			<!-- Let's now handle the dynamic part. -->
			
			<!-- Let's report the Error Message(if any), and the Loading Spinner if needed. -->
			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
			<LoadingSpinner v-if="loading"></LoadingSpinner>

			<!-- If instead, it is all ok, Display a sort of card for each of the User Photo(Depending on we are asking the whole list or just one). -->
			

		</div>

</template>

<!-- Declaration of the style(scoped) to use. -->
<style scoped>
	@import '../assets/personalProfile.css';
</style>
