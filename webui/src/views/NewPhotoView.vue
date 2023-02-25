<!-- Starting of the actual Search Page Handling. -->
<script>

import ErrorMsg from '../components/ErrorMsg.vue'

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

            // Initializing the phrase of the photo and the photo's variable.
            phrase: "",
			photo: null,

            // Initializing the PhotoId variable and the variable that will handle the preview of the Image.
            idPhoto: 0,
            previewImage: null,

            // Initializing the Background Text on the Image Box.
            photoBackgroundText: "CLICK HERE to CHOOSE A FILE",
        }
	},

	// Declaration of the methods that will be used.
	methods: {

        // This method will be triggered whenever we have to select a file to upload.
        onFileSelected (event) {

            // This will assign to photo the first selected file.
            this.photo = event.target.files[0]
            this.photoflag = true;
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
                form.append('filename', this.photo)
    
                // Adding the New Photo: /users/:username/photos/.
                let response = await this.$axios.post(`/users/${this.username}/photos/`, form, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("BearerToken")
                    }
                })    
                
                // Saving the Photo's Id got in response.
                this.idPhoto = response.data;

                // Re-addressing the page to the personal profile page of a user.
                this.$router.push({ path: `/users/${this.username}` })

            } catch (e) {

                // If an error is encountered, display it!
                this.errormsg = e.toString();
            }
        
            // Once the entire operation has finished, re-set the "loading" flag to false, in such a way to continue.
            this.loading = false;
        },

        // pickFile function: It will be used to simulate a click on the real "Choose file" that is hidden due to its uglyness.
        pickFile () {

            // Let's take from the fileInput the reference.
            let input = this.$refs.fileInput

            // Let's save the files.
            let file = input.files

            // Check whether the file contains something.
            if (file && file[0]) {

                // If it is so, create a new FileReader that will, onload, display the image in the photo box.
                let reader = new FileReader
                reader.onload = e => {
                    this.previewImage = e.target.result
                }

                reader.readAsDataURL(file[0])
            }
        },


        // selectImage is a function that is called whenever a user clicks on the box where it is written "CLICK HERE to CHOOSE a FILE".
        selectImage () {

            // First change the Background Text in such a way to eliminate it.
            this.photoBackgroundText = ""

            // Then, simulate the click on the Choose Photo button that is hidden due to its uglyness. 
            this.$refs.fileInput.click()
        },

    },
}
</script>



<!-- Actual Page for handling the page setting. -->
<template>
   
   <div>
			<!-- Let's handle first the upper part that will be the static one. -->
			<h1 class="h1">{{ username }}'s NEW PHOTO</h1>
            {{ this.file }}
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
                                <textarea  name="phrase" v-model="phrase" placeholder="Insert the Phrase..." class="form-control"  type="text"></textarea>
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
                                <!-- Since it is very ugly in its naive version, I will hide it and simulate a click on it using ref if we click on the button below. -->
                                <!-- <input type="file" @change="onFileSelected"> -->

                                <input type="file" @input="pickFile" ref="fileInput" @change="onFileSelected" style="display:none"> 
                                <div class="imagePreviewWrapper" 
                                     :style="{'background-image': `url(${previewImage})` }" @click="selectImage">
                                     <br><br><br><br><br><br><br><br><br><br><br><br><br><br><h3 style="color:#c2e9fc;">{{ this.photoBackgroundText }}</h3>
                                </div>
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
	@import '../assets/newPhoto.css';
</style>
