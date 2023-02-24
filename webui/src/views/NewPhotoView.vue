<!-- Starting of the actual Search Page Handling. -->
<script>

import ErrorMsg from '../components/ErrorMsg.vue'
import LoadingSpinner from '../components/LoadingSpinner.vue'


// Declaration of the export set.
export default {

	components: {
		ErrorMsg
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

            // Initializing the phrase of the photo.
            phrase: "",
			file: null,
        }
	},

	// Declaration of the methods that will be used.
	methods: {

        // This method will be triggered whenever we have to select a file to upload.
        onFileSelected (event) {

            // This will assign to file the first selected file.
            // this.file = event.target.files[0]
            this.file = this.$refs.file.files[0]
        },

        // uploadPhoto function: It has the role to add a new photo on the user profile.
        async uploadPhoto(){

            // Initializing the two errormessage and loading variables.
            this.errormsg= "";
            this.loading= true;

            try{

                // Creation of a multipart/form data to send to the go server.
                const form = new FormData()
                form.append('phrase', this.phrase)
                form.append('filename', this.file)
    
                // Adding the New Photo: /users/:username/photos/.
                await this.$axios.post(`/users/${this.username}/photos/`, form, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("BearerToken")
                    }
                })     

            } catch (e) {

                // If an error is encountered, display it!
                this.errormsg = e.toString();
            }
        
            // Once the entire operation has finished, re-set the "loading" flag to false, in such a way to continue.
            this.loading = false;
        },

    },
}
</script>



<!-- Actual Page for handling the page setting. -->
<template>
    <div>
        <!-- The @change will call the function and it will triggered whenever we select a new file -->
        <input type="file" @change="onFileSelected" ref="file">
        <button class="btn btn-success" @click="onFileSelected">Upload</button>
    </div>

	<div>
			<!-- Let's handle first the upper part that will be the static one. -->
			<h1 class="h1">{{ username }}'s NEW PHOTO</h1>

			<div class="topMenu">

				<!-- WASA Photo Icon -->
                <div class="topMenuButtons"></div>
				<div class="topMenuColumn">
					<img src="./img/wasa-logo.png" alt="" class="img">
				</div>
				<div class="topMenuButtons"></div>

            </div>
            <div class="formUpdate">

                <form class="well form-horizontal" action=" " method="post"  id="contact_form">
                <fieldset>

                    <!-- Phrase -->
                    <div class="form-group">
                        <label class="col-md-4 control-label"><h3><b>Phrase</b></h3></label>  
                        <div class="col-md-4 inputGroupContainer">
                            <div class="input-group">
                                <span class="input-group-addon"><i class="glyphicon glyphicon-user"></i></span>
                                <input  name="phrase" placeholder="Insert the Phrase..." class="form-control"  type="text">
                            </div>
                        </div>
                    </div>

                    <!-- Image Change -->
                    <!-- <br> -->
                    <div class="form-group">
                        <label class="col-md-4 control-label"><h3><b>New Photo</b></h3></label>
                        <div class="col-md-4 inputGroupContainer">
                            <div class="form-group">
                                <!-- The @change will call the function and it will triggered whenever we select a new file -->
                                <input type="file" class="btn btn-primary btn-block btn-large" @change="onFileSelected" ref="file">
                            </div>
                        </div>
                    </div>

                    <!-- Send Button -->
                    <div class="form-group2">
					    <button type="login-button" class="btn btn-primary btn-block btn-large" v-if="!loading" @click="uploadPhoto"> Upload New Photo </button>
				    </div>

                </fieldset>
            </form>
        </div>
    </div><!-- /.container -->
</template>

<!-- Declaration of the style(scoped) to use. -->
<style scoped>
	@import '../assets/updateProfile.css';
</style>
