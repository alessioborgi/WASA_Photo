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

			// Initializing variable for handling the UserProfile retrieval.
			userProfile: { fixedUsername: "", username: "", photoProfile: "", biography: "", dateOfCreation: "", numberOfPhotos: 0, numberFollowers: 0, numberFollowing: 0, name: "", surname: "", dateOfBirth: "", email: "", nationality: "", gender: ""},
			
            // Initializing the variable that will take the photo.
            photo: null,
            previewImage: null,

            // Initializing the Background Text on the Image Box.
            photoBackgroundText: "CLICK HERE to CHOOSE A FILE",
        }
	},

	// Declaration of the methods that will be used.
	methods: {

        // GetUserProfile Function: It retrieves the whole profile of the Logged username.
        async getUserProfile() {

            // Re-initializing variables to their default value.
            this.errormsg = "";
            this.loading = true;

            try{

                // Retrieving the Profile from the Back-end.
                let responseProfile = await this.$axios.get("/users/"+this.username, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("BearerToken")
                    }
                })
                
                // Let's add up to the "userProfiles" array the response of the profile. Note that it will be an array with only this element.
                this.userProfile = responseProfile.data;

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

            // Setting again the Loading flag to false.
            this.loading = false;
        },

        async setUserProfile() {

            // Re-initializing variables to their default value.
            this.errormsg = "";
            this.loading = true;

            // Creation of a multipart/form data to send to the go server.
            const form = new FormData()
            form.append('username', this.userProfile.username)
            form.append('photoProfile', this.photo)
            form.append('biography', this.userProfile.biography)
            form.append('name', this.userProfile.name)
            form.append('surname', this.userProfile.surname)
            form.append('dateOfBirth', this.userProfile.dateOfBirth)
            form.append('email', this.userProfile.email)
            form.append('nationality', this.userProfile.nationality)
            form.append('gender', this.userProfile.gender)

            try {
                
                // In the case the result is positive, we set the userProfile updated to the GO page.

                await this.$axios.put(`/users/${this.username}`, form, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					},

                    // This is used for showing in the console the Upload Progress Percentage.
                    onUploadProgress: uploadEvent => {
                        console.log("Upload Progress: " + Math.round(uploadEvent.loaded / uploadEvent.total * 100) + "%")
                    }
				})

                // Setting the new username received as the new username saved in the local cache.
                localStorage.setItem('Username', this.userProfile.username),
                this.username = this.userProfile.username;
                                
                // Re-addressing the page to the personal profile page of a user.
                this.$router.replace({ path: '/users/'+this.username })

            } catch (e) {

                // In case of error, retrieve it.
                if (e.response && e.response.status === 400) {
					this.errormsg = "Request error, please Login before doing some action or ask to updathe the profile of a valid user." + e.toString();
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

        // This method will be triggered whenever we have to select a file to upload.
        onFileSelected (event) {

            // This will assign to photo the first selected file.
            this.photo = event.target.files[0]
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

        clearPhoto () {

            // Initializing the Background Text on the Image Box.
            this.photoBackgroundText= "CLICK HERE to CHOOSE A FILE";
            this.previewImage = null;
        },


	},

    mounted() {
        this.getUserProfile();
    }
}
</script>



<!-- Actual Page for handling the page setting. -->
<template>

	<div>
			<!-- Let's handle first the upper part that will be the static one. -->
			<h1 class="h1">{{ username }}'s PROFILE UPDATE</h1>

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

                    <!-- Username -->
                    <div class="form-group">
                        <label class="col-md-4 control-label"><h3><b>Username</b></h3></label>  
                        <div class="col-md-4 inputGroupContainer">
                            <div class="input-group">
                                <span class="input-group-addon"><i class="glyphicon glyphicon-user"></i></span>
                                <input  name="username" :placeholder=this.userProfile.username v-model="userProfile.username" class="form-control"  type="text">
                            </div>
                        </div>
                    </div>

                    <!-- Name -->
                    <div class="form-group">
                        <label class="col-md-4 control-label"><h3><b>Name</b></h3></label>  
                        <div class="col-md-4 inputGroupContainer">
                            <div class="input-group">
                                <span class="input-group-addon"><i class="glyphicon glyphicon-user"></i></span>
                                <input  name="first_name" :placeholder=this.userProfile.name v-model="userProfile.name" class="form-control"  type="text">
                            </div>
                        </div>
                    </div>

                    <!-- Surname -->
                    <!-- <br> -->
                    <div class="form-group">
                        <label class="col-md-4 control-label"><h3><b>Surname</b></h3></label> 
                        <div class="col-md-4 inputGroupContainer">
                            <div class="input-group">
                                <span class="input-group-addon"><i class="glyphicon glyphicon-user"></i></span>
                                <input name="last_name" :placeholder=this.userProfile.surname v-model="userProfile.surname" class="form-control"  type="text">
                            </div>
                        </div>
                    </div>

                    <!-- Email-->
                    <!-- <br> -->
                    <div class="form-group">
                        <label class="col-md-4 control-label"><h3><b>E-Mail</b></h3></label>  
                        <div class="col-md-4 inputGroupContainer">
                            <div class="input-group">
                                <span class="input-group-addon"><i class="glyphicon glyphicon-envelope"></i></span>
                                <input name="email" :placeholder=this.userProfile.email v-model="userProfile.email" class="form-control"  type="text">
                            </div>
                        </div>
                    </div>

                    <!-- Nationality -->

                    <br><br>
                    <div class="form-group"> 
                        <label class="col-md-4 control-label"><h3><b>Nationality</b></h3></label>
                        <div class="col-md-4 selectContainer">
                            <div class="input-group">
                                <span class="input-group-addon"><i class="glyphicon glyphicon-list"></i></span>
                                <select name="nationality" class="select" :placeholder=this.userProfile.nationality v-model="userProfile.nationality">
                                    <option :value=userProfile.nationality>{{userProfile.nationality}}</option>
                                    <option value="afghan">Afghan</option>
                                    <option value="albanian">Albanian</option>
                                    <option value="algerian">Algerian</option>
                                    <option value="american">American</option>
                                    <option value="andorran">Andorran</option>
                                    <option value="angolan">Angolan</option>
                                    <option value="antiguans">Antiguans</option>
                                    <option value="argentinean">Argentinean</option>
                                    <option value="armenian">Armenian</option>
                                    <option value="australian">Australian</option>
                                    <option value="austrian">Austrian</option>
                                    <option value="azerbaijani">Azerbaijani</option>
                                    <option value="bahamian">Bahamian</option>
                                    <option value="bahraini">Bahraini</option>
                                    <option value="bangladeshi">Bangladeshi</option>
                                    <option value="barbadian">Barbadian</option>
                                    <option value="barbudans">Barbudans</option>
                                    <option value="batswana">Batswana</option>
                                    <option value="belarusian">Belarusian</option>
                                    <option value="belgian">Belgian</option>
                                    <option value="belizean">Belizean</option>
                                    <option value="beninese">Beninese</option>
                                    <option value="bhutanese">Bhutanese</option>
                                    <option value="bolivian">Bolivian</option>
                                    <option value="bosnian">Bosnian</option>
                                    <option value="brazilian">Brazilian</option>
                                    <option value="british">British</option>
                                    <option value="bruneian">Bruneian</option>
                                    <option value="bulgarian">Bulgarian</option>
                                    <option value="burkinabe">Burkinabe</option>
                                    <option value="burmese">Burmese</option>
                                    <option value="burundian">Burundian</option>
                                    <option value="cambodian">Cambodian</option>
                                    <option value="cameroonian">Cameroonian</option>
                                    <option value="canadian">Canadian</option>
                                    <option value="cape verdean">Cape Verdean</option>
                                    <option value="central african">Central African</option>
                                    <option value="chadian">Chadian</option>
                                    <option value="chilean">Chilean</option>
                                    <option value="chinese">Chinese</option>
                                    <option value="colombian">Colombian</option>
                                    <option value="comoran">Comoran</option>
                                    <option value="congolese">Congolese</option>
                                    <option value="costa rican">Costa Rican</option>
                                    <option value="croatian">Croatian</option>
                                    <option value="cuban">Cuban</option>
                                    <option value="cypriot">Cypriot</option>
                                    <option value="czech">Czech</option>
                                    <option value="danish">Danish</option>
                                    <option value="djibouti">Djibouti</option>
                                    <option value="dominican">Dominican</option>
                                    <option value="dutch">Dutch</option>
                                    <option value="east timorese">East Timorese</option>
                                    <option value="ecuadorean">Ecuadorean</option>
                                    <option value="egyptian">Egyptian</option>
                                    <option value="emirian">Emirian</option>
                                    <option value="equatorial guinean">Equatorial Guinean</option>
                                    <option value="eritrean">Eritrean</option>
                                    <option value="estonian">Estonian</option>
                                    <option value="ethiopian">Ethiopian</option>
                                    <option value="fijian">Fijian</option>
                                    <option value="filipino">Filipino</option>
                                    <option value="finnish">Finnish</option>
                                    <option value="french">French</option>
                                    <option value="gabonese">Gabonese</option>
                                    <option value="gambian">Gambian</option>
                                    <option value="georgian">Georgian</option>
                                    <option value="german">German</option>
                                    <option value="ghanaian">Ghanaian</option>
                                    <option value="greek">Greek</option>
                                    <option value="grenadian">Grenadian</option>
                                    <option value="guatemalan">Guatemalan</option>
                                    <option value="guinea-bissauan">Guinea-Bissauan</option>
                                    <option value="guinean">Guinean</option>
                                    <option value="guyanese">Guyanese</option>
                                    <option value="haitian">Haitian</option>
                                    <option value="herzegovinian">Herzegovinian</option>
                                    <option value="honduran">Honduran</option>
                                    <option value="hungarian">Hungarian</option>
                                    <option value="icelander">Icelander</option>
                                    <option value="indian">Indian</option>
                                    <option value="indonesian">Indonesian</option>
                                    <option value="iranian">Iranian</option>
                                    <option value="iraqi">Iraqi</option>
                                    <option value="irish">Irish</option>
                                    <option value="israeli">Israeli</option>
                                    <option value="italian">Italian</option>
                                    <option value="ivorian">Ivorian</option>
                                    <option value="jamaican">Jamaican</option>
                                    <option value="japanese">Japanese</option>
                                    <option value="jordanian">Jordanian</option>
                                    <option value="kazakhstani">Kazakhstani</option>
                                    <option value="kenyan">Kenyan</option>
                                    <option value="kittian and nevisian">Kittian and Nevisian</option>
                                    <option value="kuwaiti">Kuwaiti</option>
                                    <option value="kyrgyz">Kyrgyz</option>
                                    <option value="laotian">Laotian</option>
                                    <option value="latvian">Latvian</option>
                                    <option value="lebanese">Lebanese</option>
                                    <option value="liberian">Liberian</option>
                                    <option value="libyan">Libyan</option>
                                    <option value="liechtensteiner">Liechtensteiner</option>
                                    <option value="lithuanian">Lithuanian</option>
                                    <option value="luxembourger">Luxembourger</option>
                                    <option value="macedonian">Macedonian</option>
                                    <option value="malagasy">Malagasy</option>
                                    <option value="malawian">Malawian</option>
                                    <option value="malaysian">Malaysian</option>
                                    <option value="maldivan">Maldivan</option>
                                    <option value="malian">Malian</option>
                                    <option value="maltese">Maltese</option>
                                    <option value="marshallese">Marshallese</option>
                                    <option value="mauritanian">Mauritanian</option>
                                    <option value="mauritian">Mauritian</option>
                                    <option value="mexican">Mexican</option>
                                    <option value="micronesian">Micronesian</option>
                                    <option value="moldovan">Moldovan</option>
                                    <option value="monacan">Monacan</option>
                                    <option value="mongolian">Mongolian</option>
                                    <option value="moroccan">Moroccan</option>
                                    <option value="mosotho">Mosotho</option>
                                    <option value="motswana">Motswana</option>
                                    <option value="mozambican">Mozambican</option>
                                    <option value="namibian">Namibian</option>
                                    <option value="nauruan">Nauruan</option>
                                    <option value="nepalese">Nepalese</option>
                                    <option value="new zealander">New Zealander</option>
                                    <option value="ni-vanuatu">Ni-Vanuatu</option>
                                    <option value="nicaraguan">Nicaraguan</option>
                                    <option value="nigerien">Nigerien</option>
                                    <option value="north korean">North Korean</option>
                                    <option value="northern irish">Northern Irish</option>
                                    <option value="norwegian">Norwegian</option>
                                    <option value="omani">Omani</option>
                                    <option value="pakistani">Pakistani</option>
                                    <option value="palauan">Palauan</option>
                                    <option value="panamanian">Panamanian</option>
                                    <option value="papua new guinean">Papua New Guinean</option>
                                    <option value="paraguayan">Paraguayan</option>
                                    <option value="peruvian">Peruvian</option>
                                    <option value="polish">Polish</option>
                                    <option value="portuguese">Portuguese</option>
                                    <option value="qatari">Qatari</option>
                                    <option value="romanian">Romanian</option>
                                    <option value="russian">Russian</option>
                                    <option value="rwandan">Rwandan</option>
                                    <option value="saint lucian">Saint Lucian</option>
                                    <option value="salvadoran">Salvadoran</option>
                                    <option value="samoan">Samoan</option>
                                    <option value="san marinese">San Marinese</option>
                                    <option value="sao tomean">Sao Tomean</option>
                                    <option value="saudi">Saudi</option>
                                    <option value="scottish">Scottish</option>
                                    <option value="senegalese">Senegalese</option>
                                    <option value="serbian">Serbian</option>
                                    <option value="seychellois">Seychellois</option>
                                    <option value="sierra leonean">Sierra Leonean</option>
                                    <option value="singaporean">Singaporean</option>
                                    <option value="slovakian">Slovakian</option>
                                    <option value="slovenian">Slovenian</option>
                                    <option value="solomon islander">Solomon Islander</option>
                                    <option value="somali">Somali</option>
                                    <option value="south african">South African</option>
                                    <option value="south korean">South Korean</option>
                                    <option value="spanish">Spanish</option>
                                    <option value="sri lankan">Sri Lankan</option>
                                    <option value="sudanese">Sudanese</option>
                                    <option value="surinamer">Surinamer</option>
                                    <option value="swazi">Swazi</option>
                                    <option value="swedish">Swedish</option>
                                    <option value="swiss">Swiss</option>
                                    <option value="syrian">Syrian</option>
                                    <option value="taiwanese">Taiwanese</option>
                                    <option value="tajik">Tajik</option>
                                    <option value="tanzanian">Tanzanian</option>
                                    <option value="thai">Thai</option>
                                    <option value="togolese">Togolese</option>
                                    <option value="tongan">Tongan</option>
                                    <option value="trinidadian or tobagonian">Trinidadian or Tobagonian</option>
                                    <option value="tunisian">Tunisian</option>
                                    <option value="turkish">Turkish</option>
                                    <option value="tuvaluan">Tuvaluan</option>
                                    <option value="ugandan">Ugandan</option>
                                    <option value="ukrainian">Ukrainian</option>
                                    <option value="uruguayan">Uruguayan</option>
                                    <option value="uzbekistani">Uzbekistani</option>
                                    <option value="venezuelan">Venezuelan</option>
                                    <option value="vietnamese">Vietnamese</option>
                                    <option value="welsh">Welsh</option>
                                    <option value="yemenite">Yemenite</option>
                                    <option value="zambian">Zambian</option>
                                    <option value="zimbabwean">Zimbabwean</option>
                                </select>
                            </div>
                        </div>
                    </div>

                    <!-- Gender -->
                    <br>
                    <div class="form-group"> 
                        <label class="col-md-4 control-label"><h3><b>Gender</b></h3></label>
                        <div class="col-md-4 selectContainer">
                            <div class="input-group">
                                <span class="input-group-addon"><i class="glyphicon glyphicon-list"></i></span>
                                <select name="nationality" class="select" :placeholder=this.userProfile.gender v-model="userProfile.gender">
                                    <option value="male">Male</option>
                                    <option value="female">Female</option>
                                    <option value="do not specify">Do Not Specify</option>
                                </select>
                            </div>
                        </div>
                    </div>


                    <!-- Biography -->

                    <br><br>
                    <div class="form-group">
                        <label class="col-md-4 control-label"><h3><b>Biography</b></h3></label>
                        <div class="col-md-4 inputGroupContainer">
                            <div class="input-group">
                                <span class="input-group-addon"><i class="glyphicon glyphicon-pencil"></i></span>
                                <textarea class="form-control2" name="comment" :placeholder=this.userProfile.biography v-model="userProfile.biography"></textarea>
                            </div>
                        </div>
                    </div>

                    <!-- Image Change -->
                    <br>
                    <div class="form-group">
                        <label class="col-md-4 control-label"><h3><b>Photo Profile</b></h3></label>
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
                </fieldset>
            </form>
        </div>

        <!-- Clear Button -->
        <div class="clearButton">
            <!-- <button class="btn btn-primary btn-block btn-large" @click="clearPhoto">Clear</button> -->
            <button @click="clearPhoto" style="display: block; background: #c2e9fc; border: 1px solid #FF4742; border-radius: 6px; box-shadow: rgba(0, 0, 0, 0.1) 1px 2px 4px; box-sizing: border-box; color: #c2e9fc; cursor: pointer; font-family: sans-serif; font-size: 16px; font-weight: 800; line-height: 16px; min-height: 40px; outline: 0; padding: 12px 14px; text-align: center; text-rendering: geometricprecision; text-transform: none; user-select: none; -webkit-user-select: none; touch-action: manipulation; vertical-align: middle; background-color: #c2e9fc; background-position: 0 0; color: black; ">
                Clear</button>
        </div>

        <!-- Send Button -->
        <div class="form-group2">
			<button type="login-button" class="btn btn-primary btn-block btn-large" @click="setUserProfile"> Update Profile </button>
		</div>
        
        <div>
            <!-- Let's report the Error Message(if any), and the Loading Spinner if needed. -->
			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        </div>
        
    </div><!-- /.container -->
</template>

<!-- Declaration of the style(scoped) to use. -->
<style scoped>
	@import '../assets/updateProfile.css';
</style>
