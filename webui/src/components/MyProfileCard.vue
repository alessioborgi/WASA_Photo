
<script>

export default {

    props: ['user', 'userOwnerFlag'],

    components: {
    },
    
	// Describing what are the Return variables.
	data: function() {
		return {	
            
            // Initializing the two errormessage and loading variables.
			errormsg: "",
			loading: false,

			// Retrieving from the Cache the Username and the Bearer Authenticaiton Token.
			username: localStorage.getItem('Username') == localStorage.getItem('usernameProfileToView') ? localStorage.getItem('Username') : localStorage.getItem('usernameProfileToView'),
            BearerToken: localStorage.getItem('BearerToken'),

            // Creation of a Local Attribute to handle the "user" prop.
            userProfile: { fixedUsername: this.user.fixedUsername, username: this.user.username, photoProfile: this.user.photoProfile, 
                           biography: this.user.biography, dateOfCreation: this.user.dateOfCreation, numberOfPhotos: this.user.numberOfPhotos, 
                           numberFollowers: this.user.numberFollowers, numberFollowing: this.user.numberFollowing, name: this.user.name, 
                           surname: this.user.surname, dateOfBirth: this.user.dateOfBirth, email: this.user.email, nationality: this.user.nationality, 
                           gender: this.user.gender,  boolFollowing: this.user.boolFollowing, boolFollower: this.user.boolFollower, boolBanned: this.user.boolBanned },
            // userProfile: this.user,
            
            // Initializing variable for handling the deletion of the Profile.
            deleteProfileBool: false,

            // Initializing the photoIdView that will be the photoProfile photoID.
            photoIdView: 0,

            // Initializing the image as a blob object, and declaring an object URL form it.
            photoBlobLink: this.user.photoProfile == "" ? "https://lh3.googleusercontent.com/ytP9VP86DItizVX2YNA-xTYzV09IS7rh4WexVp7eilIcfHmm74B7odbcwD5DTXmL0PF42i2wnRKSFPBHlmSjCblWHDCD2oD1oaM1CGFcSd48VBKJfsCi4bS170PKxGwji8CPmehwPw=w200-h247-no" : this.user.photoProfile
		}
	},

    methods: {
        
        async deleteProfileAlert() {

            if (confirm("Your Profile" + this.username +" will be deleted from WASAPhoto. Are you sure?")){
                this.deleteProfileBool = true;
                alert("Profile Correctly Deleted");
            } else {
                this.deleteProfileBool = false;
                alert("Profile still Alive!")
            }
            
        },


        // deleteProfile: This functio will be used to delete the profile. 
        async deleteProfile() {

            // Re-initializing variables to their default value.
            this.errormsg = "";
            this.loading = true;

            await this.deleteProfileAlert()

            try {
                
                // In the case the result is positive, we post the username received to the GO page.
                if (this.deleteProfileBool == true){
                    await this.$axios.delete(`/users/${this.username}`, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("BearerToken")
                    }})

                    // Setting the uuid (Bearer Token) received as response by the Post action.
                    localStorage.clear();
                    this.username = "";
                    this.BearerToken = "";
                                    
                    // Re-addressing the page to the personal profile page of a user.
                    this.$router.replace({ path: '/session/' })
                }

            } catch (e) {

                // In case of error, retrieve it.
                if (e.response && e.response.status === 400) {
					this.errormsg = "Request error, please Login before doing some action" + e.toString();
				} else if (e.response && e.response.status === 403) {
					this.errormsg = "An Unauthorized Action has been blocked. You are not allowed to do this action because you are not the profile's owner." + e.toString();
                }else if (e.response && e.response.status === 500) {
					this.errormsg = "An internal error occurred. We will be notified. Please try again later." + e.toString();
				} else {
					this.errormsg = e.toString();
				}

            }

            // Setting again the Loading flag to false.
            this.loading = false;
        },

        async getPhotoView() {

            // Re-initializing variables to their default value.
			this.errormsg = "";
			this.loading = true;

            try {
                
                // Getting the image view from the Back-End.
                // /users/:username/photos/:photoid/view
                let response = await this.$axios.get(`/users/${this.username}/photos/0/view`, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("BearerToken")
                    },

                    responseType: 'blob'
                })

                var photoBlob = response.data;
                this.photoBlobLink = URL.createObjectURL(photoBlob);

            } catch (e) {

                // In case of error, retrieve it.
                if (e.response && e.response.status === 400) {
					this.errormsg = "Request error, please Login before doing some action" + e.toString();
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


		async changeUsernameAlert() {

            this.newUsername = "";
            this.newUsername = prompt("Please enter the new Username:");

            if (this.newUsername.length < 3) {
                alert("You cannot change the Username because you have inserted an empty or a less than 3 characters username. It does not respect the Regex!")
            } else{
                if (confirm("The username will be updated to: " + this.newUsername)){} 
            }
        },

        // setUsername: This method is used for changing the Username. 
        async setUsername() {

            // Re-initializing variables to their default value.
            this.errormsg = "";
            this.loading = true;

            await this.changeUsernameAlert();

            try {
                
                // In the case the result is positive, we post the username received to the GO page.

                await this.$axios.patch(`/users/${this.username}`, { username: this.newUsername}, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("BearerToken")
                    }
                })

                // Setting the new username received as the new username saved in the local cache.
                this.username = this.newUsername;
                this.userProfile.username = this.newUsername;
                localStorage.setItem('Username', this.newUsername),
                localStorage.setItem('usernameProfileToView', this.newUsername)
                // Re-addressing the page to the personal profile page of a user.
                this.newUsername = "";
                // Re-address the user to the right page.
                // await this.
                // alert("Redirecting!" + this.username)
                this.$router.replace({ path: `/users/${this.userProfile.username}`})

            } catch (e) {

                // In case of error, retrieve it.
                if (e.response && e.response.status === 400) {
                    this.errormsg = "Request error, please Login before doing some action or ask to update the username of a valid user." + e.toString();
                } else if (e.response && e.response.status === 403) {
                    this.errormsg = "An Unauthorized Action has been blocked. You are not allowed to do this action because you are not the profile's owner." + e.toString();
                } else if (e.response && e.response.status === 204) {
                    this.errormsg = "In the Internal DB there is not anymore the content you have asked." + e.toString();
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later." + e.toString();
                } else {
                    this.errormsg = "Please Login before with an Authorized profile to view this page. " + e.toString();
                }
            }

            // Setting again the Loading flag to false.
            this.loading = false;
        },

        async goToUpdate() {

            // Re-address the user to the right page.
            this.$router.push({ path: `/users/${this.username}/update/`})
        },

        async goToAnalytics() {

            // Re-address the user to the right page.
            this.$router.push({ path: `/users/${this.username}/analytics/`})
        }, 
        
        async goToNewPhoto() {

            // Re-address the user to the right page.
            this.$router.push({ path: `/users/${this.username}/photo/`})
        },

        async goToBan() {

            // Re-address the user to the right page.
            this.$router.push({ path: `/users/${this.username}/ban/`})
        },

        async goToFollow() {

            // Re-address the user to the right page.
            this.$router.push({ path: `/users/${this.username}/follow/`})
        }, 

        async goToSearch() {

            // Re-address the user to the right page.
            this.$router.push({ path: `/search/`})
        }, 
        
        async goToMyStream() {

            // Re-address the user to the right page.
            this.$router.push({ path: `/users/${this.username}/myStream/`})
        }, 

        
    },

    mounted() {

        // Getting first the photo, provided that it is not the Default Photo (i.e., there is no photo!).
        if (this.userProfile.profileImage != "" ){
		    this.getPhotoView()
        }

	}

}

     
</script>
    

<template>

    <!-- If instead, it is all ok, Display a sort of card for each of the User Profiles(Depending on we are asking the whole list or just one). -->
    <div class="card" id="div1">

        <div class="usernameLabel">
            <!-- {{ this.user }}
            {{ this.userProfile }} -->
            <!-- <b> FIXEDUSEkRNAME: </b>{{ this.user.photoProfile }}
            <b> FIXEDUSERNAME: </b>{{ this.photoIdView }} 
            <b> FIXEDUSERNAME: </b>{{ this.photoBlobLink }}  -->
            
            <!-- <b> FIXEDUSERNAME: </b>{{ this.user.fixedUsername }}  -->

        </div>
        <div class="upperPart"> 
            <div class="imageLabel">
                <div class="profileImage">
                    <!-- In this way works -->
                    <!-- <img src="../../../../tmp/u1-photo-0-photo-profile.jpg" alt="Person" class="card__image"/> -->
                    <img :src=this.photoBlobLink class="card__image" />
                </div>
                <div class="profileLabel">
                    <p class="card__name" > <b>{{ this.user.username }}</b></p>
                </div>            
            </div>
            <div class="rightUpperPart">

                <div class="grid-container2">
                    <div class="grid-child-posts">
                        <b>Posts</b> {{ this.user.numberOfPhotos }}
                    </div>

                    <div class="grid-child-posts">
                        <b>Followings</b> {{ this.user.numberFollowing }} 
                    </div>

                    <div class="grid-child-posts">
                        <b>Followers</b> {{ this.user.numberFollowers }} 
                    </div>
                </div>


                <div class="grid-child-posts3">
                    <b>Biography</b> {{ this.user.biography }} 
                </div>
            </div>
                    
        </div>
            
            
        <div class="grid-container">

            <div class="grid-child-posts">
                <b>Name</b> {{ this.user.name }}
            </div>

            <div class="grid-child-posts">
                <b>Surname</b> {{ this.user.surname }} 
            </div>

            <div class="grid-child-posts">
                <b>Nationality</b> {{ this.user.nationality }} 
            </div>

            <div class="grid-child-posts">
                <b>DateOfBirth</b> {{ this.user.dateOfBirth }} 
            </div>

            <div class="grid-child-posts">
                <b>Email</b> {{ this.user.email }} 
            </div>

            <div class="grid-child-posts">
                <b>Gender</b> {{ this.user.gender }} 
            </div>

            <div class="grid-child-posts">
                <b>DateOfCreation</b> {{ this.user.dateOfCreation }} 
            </div>

            <div class="grid-child-posts" style="margin-left: -365px;" v-if="!loading && userOwnerFlag === true">
                        <!-- <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#settings" @click="replaceLogin"/></svg> -->
                        <nav>
                            <menu >
                                <menuitem id="demo1">
                                    <a><svg class="feather"><use href="/feather-sprite-v4.29.0.svg#list"/></svg></a>
                                    <menu>

                                        <menuitem>
                                            <!-- <button @click="showAlert"> -->
                                                <button type="login-button" class="btn btn-primary btn-block btn-large" 
                                                    @click="goToNewPhoto">                                                
                                                    <b>New Photo</b>
                                            </button>
                                        </menuitem>

                                        <menuitem>
                                            <button type="login-button" class="btn btn-primary btn-block btn-large" 
                                                @click="goToBan">
                                                <b>Ban Dashboard</b>
                                            </button>
                                        </menuitem>

                                        <menuitem>
                                            <button type="login-button" class="btn btn-primary btn-block btn-large" 
                                                @click="goToFollow">
                                                <b>Follow Dashboard</b>
                                            </button>
                                        </menuitem>

                                        <menuitem>
                                            <button type="login-button" class="btn btn-primary btn-block btn-large" 
                                                @click="goToSearch">
                                                <b>Search User</b>
                                            </button>
                                        </menuitem>

                                    </menu>
                                </menuitem>
                            
                            </menu>
	                    </nav>     

                        <a><svg class="feather" style="margin-left:410px; margin-top: -100px;" @click="goToMyStream">
                            <use href="/feather-sprite-v4.29.0.svg#home"/></svg>
                        </a>              
            </div>
                                               

            <div class="grid-child-posts" style="margin-left: -90px;" v-if="!loading && userOwnerFlag === true">
                        <!-- <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#settings" @click="replaceLogin"/></svg> -->
                        <nav>
                            <menu >
                                <menuitem id="demo1">
                                    <a><svg class="feather"><use href="/feather-sprite-v4.29.0.svg#settings"/></svg></a>
                                    <menu>

                                        <menuitem>
                                            <!-- <button @click="showAlert"> -->
                                                <button type="login-button" class="btn btn-primary btn-block btn-large" 
                                                    @click="setUsername">                                                
                                                    <b>Set Username</b>
                                            </button>
                                        </menuitem>

                                        <menuitem>
                                            <button type="login-button" class="btn btn-primary btn-block btn-large" 
                                                @click="goToUpdate">
                                                <b>Update Profile</b>
                                            </button>
                                        </menuitem>

                                        <menuitem>
                                            <button type="login-button" class="btn btn-primary btn-block btn-large" 
                                                @click="goToAnalytics">
                                                <b>See Analytics</b>
                                            </button>
                                        </menuitem>

                                        <menuitem>
                                            <button type="login-button" class="btn btn-primary btn-block btn-large" 
                                                @click="deleteProfile">
                                                <b>Delete Profile</b>
                                            </button>
                                        </menuitem>

                                    </menu>
                                </menuitem>
                            
                            </menu>
	                    </nav>                  
            </div>

        </div>

        

    </div>

</template>





<style scoped>

.btn { display: inline-block; *display: inline; *zoom: 1;   font-family: sans-serif; padding: 4px 10px 4px; margin-bottom: 0; font-size: 13px; line-height: 18px; color: #333333; text-align: center;text-shadow: 0 1px 1px rgba(255, 255, 255, 0.75); vertical-align: middle; background-color: #f5f5f5; background-image: -moz-linear-gradient(top, #ffffff, #e6e6e6); background-image: -ms-linear-gradient(top, #ffffff, #e6e6e6); background-image: -webkit-gradient(linear, 0 0, 0 100%, from(#ffffff), to(#e6e6e6)); background-image: -webkit-linear-gradient(top, #ffffff, #e6e6e6); background-image: -o-linear-gradient(top, #ffffff, #e6e6e6); background-image: linear-gradient(top, #ffffff, #e6e6e6); background-repeat: repeat-x; filter: progid:dximagetransform.microsoft.gradient(startColorstr=#ffffff, endColorstr=#e6e6e6, GradientType=0); border-color: #e6e6e6 #e6e6e6 #e6e6e6; border-color: rgba(0, 0, 0, 0.1) rgba(0, 0, 0, 0.1) rgba(0, 0, 0, 0.25); border: 1px solid #e6e6e6; -webkit-border-radius: 4px; -moz-border-radius: 4px; border-radius: 4px; -webkit-box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.2), 0 1px 2px rgba(0, 0, 0, 0.05); -moz-box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.2), 0 1px 2px rgba(0, 0, 0, 0.05); box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.2), 0 1px 2px rgba(0, 0, 0, 0.05); cursor: pointer; *margin-left: .3em; }
.btn:hover, .btn:active, .btn.active, .btn.disabled, .btn[disabled] { background-color: #e6e6e6; }
.btn-large { padding: 9px 14px; font-size: 20px; line-height: normal; -webkit-border-radius: 5px; -moz-border-radius: 5px; border-radius: 5px; }
.btn:hover { color: #333333; text-decoration: none; background-color: #e6e6e6; background-position: 0 -15px; -webkit-transition: background-position 0.1s linear; -moz-transition: background-position 0.1s linear; -ms-transition: background-position 0.1s linear; -o-transition: background-position 0.1s linear; transition: background-position 0.1s linear; }
.btn-primary, .btn-primary:hover { text-shadow: 0 -1px 0 rgba(0, 0, 0, 0.25); color: #ffffff; }
.btn-primary.active { color: rgba(255, 255, 255, 0.75); }
.btn-primary { background-color: #4a77d4; background-image: -moz-linear-gradient(top, #6eb6de, #4a77d4); background-image: -ms-linear-gradient(top, #6eb6de, #4a77d4); background-image: -webkit-gradient(linear, 0 0, 0 100%, from(#6eb6de), to(#4a77d4)); background-image: -webkit-linear-gradient(top, #6eb6de, #4a77d4); background-image: -o-linear-gradient(top, #6eb6de, #4a77d4); background-image: linear-gradient(top, #6eb6de, #4a77d4); background-repeat: repeat-x; filter: progid:dximagetransform.microsoft.gradient(startColorstr=#6eb6de, endColorstr=#4a77d4, GradientType=0);  border: 1px solid #3762bc; text-shadow: 1px 1px 1px rgba(0,0,0,0.4); box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.2), 0 1px 2px rgba(0, 0, 0, 0.5); }
.btn-primary:hover, .btn-primary:active, .btn-primary.active, .btn-primary.disabled, .btn-primary[disabled] { filter: none; background-color: #4a77d4; }
.btn-block { width: 100%; display:block; }

.settingsMenu{
    display: block;
    margin-top: -50px;
    margin-left: 1050px;
    width: 50%;
    width: 600px;
    float: left;
    /* background-color: yellow; */
    
}
.upperPart{
    display: block;
  	margin-left: auto;
  	margin-right: auto;
    width: 50%;
    width: 600px;
    height: 230;
    float: left;
    /* background-color: yellow; */
    
}
.imageLabel{
    display: block;
  	margin-left: auto;
  	margin-right: auto;
    float: left;
    height: auto;
  	width: 50%;
    width: 200px;
    height: 230;
    /* background-color: orange; */
}

.rightUpperPart{
    float: left;
  	width: 70%;
    width: 400px;
    /* background-color: purple; */
    height: 230px;
}
.buttons-menu{
    
	float: left;
  	width: 30%;
}

.usernameLabel{
    display: block;
	float: left;
  	width: 90%;
    margin-top: 10px;
    margin-left: -60px;
    font-size: 9px;
}


.buttonsFollowBan{
	float: left;
  	width: 50%;
}

.feather {
	color: #4a77d4;
}

#div1{
    background: #c2e9fc;
}

.card {
  /* background-color: #c2e9fc; */
  /* background-color: yellow; */
  margin-bottom: 20px;
  height: 45rem;
  width: auto;
  border-radius: 5px;
  align-items: center;
  margin-left: auto;
  flex-direction: column;
  align-items: center;
  box-shadow: rgba(0, 0, 0, 0.7);
  color: black;

}

.card__name {
  align-items: center;
  margin-right: auto;
  margin-top: 15px;
  margin-left: -125px;
  text-align: center;
  font-size: 2em;
}

.card__image {
  height: 180px;
  width: 180px;
  border-radius: 50%;
  border: 5px solid #272133;
  margin-top: 20px;
  margin-left: -60px;
}
.grid-container {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;
  grid-gap: 20px;
  font-size: 1.2em;
  margin-left: 50px;
  /* background-color: red; */
  font-size: 15px;
}

.grid-container2 {
margin-top: 50px;
  display: grid;
  grid-template-columns: 1fr 1fr 1fr 1fr;
  grid-gap: 50px;
  width: 600px;
  font-size: 18px;
  /* background-color: greenyellow; */
}

.grid-child-posts3{
    margin-top: 65px;
    font-size:15px;
    width: 500px;
}



html, body{
   padding:0px;
   margin:0px;
   background:#191A1D;
   font-family: 'Karla', sans-serif;
   width:100vw;
}
body * {
   margin:0;
   padding:0;
}

/* HTML Nav Styles */
/* HTML Nav Styles */
/* HTML Nav Styles */
nav menuitem {
   position:relative;
   display:block;
   opacity:0;
   cursor:pointer;
}

nav menuitem > menu {
   position: absolute;
   pointer-events:none;
}
nav > menu { display:flex; }

nav > menu > menuitem { pointer-events: all; opacity:1; }
menu menuitem a { white-space:nowrap; display:block; }
   
menuitem:hover > menu {
   pointer-events:initial;
}
menuitem:hover > menu > menuitem,
menu:hover > menuitem{
   opacity:1;
}
nav > menu > menuitem menuitem menu {
   transform:translateX(100%);
   top:0; right:0;
}
/* User Styles Below Not Required */
/* User Styles Below Not Required */
/* User Styles Below Not Required */

nav { 
   margin-top: 40px;
   margin-left: 40px;
}

nav a {
   background:#ffffff;
   color:#FFF;
   min-width:190px;
   transition: background 0.5s, color 0.5s, transform 0.5s;
   margin:0px 6px 6px 0px;
   padding:20px 40px;
   box-sizing:border-box;
   border-radius:3px;
   box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.5);
   position:relative;
   /* margin-top: -50px; */
}

nav a:hover:before {
   content: '';
   top:0;left:0;
   position:absolute;
   background:rgba(0, 0, 0, 0.2);
   width:100%;
   height:100%;
}

nav > menu > menuitem > a + menu:after{
   content: '';
   position:absolute;
   border:10px solid transparent;
   border-top: 10px solid #4a77d4;
   left:12px;
   top: -40px;  
}
nav menuitem > menu > menuitem > a + menu:after{ 
   content: '';
   position:absolute;
   border:10px solid transparent;
   border-left: 10px solid white;
   top: 20px;
   left:-180px;
   transition: opacity 0.6, transform 0s;
}

nav > menu > menuitem > menu > menuitem{
   transition: transform 0.6s, opacity 0.6s;
   transform:translateY(150%);
   opacity:0;
}
nav > menu > menuitem:hover > menu > menuitem,
nav > menu > menuitem.hover > menu > menuitem{
   transform:translateY(0%);
   opacity: 1;
}

menuitem > menu > menuitem > menu > menuitem{
   transition: transform 0.6s, opacity 0.6s;
   transform:translateX(195px) translateY(0%);
   opacity: 0;
} 
menuitem > menu > menuitem:hover > menu > menuitem,  
menuitem > menu > menuitem.hover > menu > menuitem{  
   transform:translateX(0) translateY(0%);
   opacity: 1;
}



</style>


