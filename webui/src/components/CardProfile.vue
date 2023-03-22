
<script>

export default {

    props: ['user'],

    
	// Describing what are the Return variables.
	data: function() {
		return {

            // Initializing the two errormessage and loading variables.
            errormsg: "",
			loading: false,

            // Retrieving from the Cache the Username and the Bearer Authenticaiton Token.
            username: localStorage.getItem('Username'),
            BearerToken: localStorage.getItem('BearerToken'),

            // Creation of a Local Attribute to handle the "user" prop.
            userProfile: this.user,

            // Initializing colorBackground of the Card depending on the Gender.
            colorBackground: this.user.gender == "male" ? '#c2e9fc' : this.user.gender == "female" ? '#fbd3f0' : '#cff6cc',

			// Initializing iconFollow, that can receive two values:
            //   Follow(true):    /feather-sprite-v4.29.0.svg#user-check
            //   NotFollow(false): /feather-sprite-v4.29.0.svg#user-plus
			iconFollowing: this.user.boolFollowing == true ? '/feather-sprite-v4.29.0.svg#user-check' : '/feather-sprite-v4.29.0.svg#user-plus',
            colorIconFollowing: this.user.boolFollowing == true ? 'green' : 'red',
            iconFollower: this.user.boolFollower == true ? '/feather-sprite-v4.29.0.svg#user-check' : '/feather-sprite-v4.29.0.svg#user-x',
            colorIconFollower: this.user.boolFollower == true ? 'green' : 'red',
            iconBanned: this.user.boolBanned == true ? '/feather-sprite-v4.29.0.svg#lock' : '/feather-sprite-v4.29.0.svg#unlock',
            colorIconBanned: this.user.boolBanned == true ? 'red' : 'green',

            // Initializing the image as a blob object, and declaring an object URL form it.
            photoBlobLink: this.user.photoProfile == "" ? "https://lh3.googleusercontent.com/ytP9VP86DItizVX2YNA-xTYzV09IS7rh4WexVp7eilIcfHmm74B7odbcwD5DTXmL0PF42i2wnRKSFPBHlmSjCblWHDCD2oD1oaM1CGFcSd48VBKJfsCi4bS170PKxGwji8CPmehwPw=w200-h247-no" : this.user.photoProfile
        }
	},

    methods: {

        // banUnbanUser function: It has the role to add or delete a ban depending on the boolBanned value.
        async banUnbanUser(){

            // Initializing the two errormessage and loading variables.
			this.errormsg= "";
			this.loading= true;

            // Let's handle first the case where the user is Banned.
            // We must therefore delete the Ban.
            if (this.userProfile.boolBanned == true) {

                try{

                    // Deleting the Ban: /users/:username/bans/:usernameBanned.
                    await this.$axios.delete(`/users/${this.username}/bans/${this.userProfile.username}`, {
                        headers: {
                            Authorization: "Bearer " + localStorage.getItem("BearerToken")
                        }
                    })

                    // Once we have done with it, we must simply update the flag.
                    this.userProfile.boolBanned = false;
                    this.iconBanned = '/feather-sprite-v4.29.0.svg#unlock';
                    this.colorIconBanned = 'green';
                    this.$emit('refreshBan', false);

                } catch (e) {

                    if (e.response && e.response.status === 400) {
					this.errormsg = "Request error, please Login before doing some action or ask to delete the ban of a valid user." + e.toString();
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
            } else{

                // Let's handle first the case where the user is NOT Banned.
                // We must therefore add the Ban.
                try{

                    // Deleting the Ban: /users/:username/bans/:usernameBanned.
                    await this.$axios.put(`/users/${this.username}/bans/${this.userProfile.username}`, {}, {
                        headers: {
                            Authorization: "Bearer " + localStorage.getItem("BearerToken")
                        }
                    })

                    // Once we have done with it, we must simply update the flag.
                    this.userProfile.boolBanned = true;
                    this.iconBanned = '/feather-sprite-v4.29.0.svg#lock';
                    this.colorIconBanned = 'red';
                    this.$emit('refreshBan', true);

                } catch (e) {

                    if (e.response && e.response.status === 400) {
					this.errormsg = "Request error, please Login before doing some action or ask to ban a valid user." + e.toString();
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
            }

            // Once the entire operation has finished, re-set the "loading" flag to false, in such a way to continue.
			this.loading = false;

        },


        // followUnfollowUser function: It has the role to add or delete a follow depending on the boolFollowing value.
        async followUnfollowUser(){

            // Initializing the two errormessage and loading variables.
            this.errormsg= "";
            this.loading= true;

            // Let's handle first the case where we are currently following the user.
            // We must therefore delete the Follow.
            if (this.userProfile.boolFollowing == true) {

                try{

                    // Deleting the Follow: /users/:username/followings/:usernameFollowing.
                    await this.$axios.delete("/users/"+this.username+"/followings/"+this.userProfile.username, {
                        headers: {
                            Authorization: "Bearer " + localStorage.getItem("BearerToken")
                        }
                    })

                    // Once we have done with it, we must simply update the flag.
                    this.userProfile.boolFollowing = false;
                    this.iconFollowing = '/feather-sprite-v4.29.0.svg#user-plus';
                    this.colorIconFollowing = 'red';
                    this.$emit('refreshFollowing', false);
                    this.$emit('refreshNumberFollowers', this.userProfile.numberFollowers - 1);

                } catch (e) {

                    if (e.response && e.response.status === 400) {
					this.errormsg = "Request error, please Login before doing some action or ask to delete the follow of a valid user." + e.toString();
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

            } else{

                // Let's handle first the case where the user is NOT Followed by us.
                // We must therefore add the Follow.
                try{

                    // Adding the Follow: /users/:username/followings/:usernameFollowing.
                    await this.$axios.put(`/users/${this.username}/followings/${this.userProfile.username}`, {}, {
                        headers: {
                            Authorization: "Bearer " + localStorage.getItem("BearerToken")
                        }
                    })

                    // Once we have done with it, we must simply update the flag.
                    this.userProfile.boolFollowing = true;
                    this.iconFollowing = '/feather-sprite-v4.29.0.svg#user-check';
                    this.colorIconFollowing = 'green';
                    this.$emit('refreshFollowing', true);
                    this.$emit('refreshNumberFollowers', this.userProfile.numberFollowers + 1);

                } catch (e) {

                    if (e.response && e.response.status === 400) {
					this.errormsg = "Request error, please Login before doing some action or ask to follow a valid user." + e.toString();
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
            }

            // Once the entire operation has finished, re-set the "loading" flag to false, in such a way to continue.
            this.loading = false;
        },

        async goToProfileView(userToView) {

            // Re-address the user to the right page.
            localStorage.setItem('usernameProfileToView', userToView),
            this.$router.push({ path: `/users/${userToView}`})
        },

        async getPhotoView() {

            // Re-initializing variables to their default value.
            this.errormsg = "";
            this.loading = true;

            try {
                
                // Getting the image view from the Back-End.
                // /users/:username/photos/:photoid/view
                let response = await this.$axios.get(`/users/${this.userProfile.username}/photos/0/view`, {
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
    },

    mounted() {
        this.getPhotoView()
    }  
}    
</script>
    

<template>

    <!-- If instead, it is all ok, Display a sort of card for each of the User Profiles(Depending on we are asking the whole list or just one). -->
    <div class="card" id="div1" :style="{backgroundColor: this.colorBackground}">
        <!-- <div class="card" id="div1" :style="{backgroundColor: this.colorBackground}"> -->

        <div class="usernameLabel">
            <!-- <b> FIXEDUSERNAME: </b>{{ this.userProfile }}  -->
            <!-- <b> FIXEDUSERNAME: </b>{{ this.userProfile.fixedUsername }}  -->

        </div>
        
        <!--  <div class="grid-container2"> -->

            <div class="upperPart"> 
                <div class="imageLabel">
                    <div class="profileImage">
                        <img :src=this.photoBlobLink alt="Person" class="card__image">
                    </div>
                    <div class="profileLabel">
                        <p class="card__name" > <b>{{ this.userProfile.username }}</b></p>
                    </div>            
                </div>

                <div class="rightUpperPart">

                    <div class="grid-container2">
                        <div class="grid-child-posts">
                        <b>Posts</b> {{ this.userProfile.numberOfPhotos }}
                        </div>

                        <div class="grid-child-posts">
                            <b>Followings</b> {{ this.userProfile.numberFollowing }} 
                        </div>

                        <div class="grid-child-posts">
                            <b>Followers</b> {{ this.userProfile.numberFollowers }} 
                        </div>
                    </div>


                    <div class="grid-container2">
                        <div class="grid-child-posts2">
                            <b>Is it Banned? </b> 
                            <svg class="feather" 
                                @click="banUnbanUser" 
                                :style="{color: this.colorIconBanned}">
                                <use :href="this.iconBanned"/>
                            </svg>
                        </div> 

                        <div class="grid-child-posts2">
                            <b>Am I Following it?</b>
                            <svg class="feather" 
                                @click="followUnfollowUser" 
                                :style="{color: this.colorIconFollowing}">
                                <use :href="this.iconFollowing"/>
                            </svg>
                        </div>

                        <div class="grid-child-posts2">
                            <b>Is it my Follower?</b>
                            <svg class="feather" 
                                :style="{color: this.colorIconFollower}">
                                <use :href="this.iconFollower"/>
                            </svg>
                        </div>                          
                    </div>
                </div>
                    
            </div>
            
            
        <div class="grid-container">

            <div class="grid-child-posts">
                <b>Name</b> {{ this.userProfile.name }}
            </div>

            <div class="grid-child-posts">
                <b>Surname</b> {{ this.userProfile.surname }} 
            </div>

            <div class="grid-child-posts">
                <b>Nationality</b> {{ this.userProfile.nationality }} 
            </div>

            <div class="grid-child-posts">
                <b>DateOfBirth</b> {{ this.userProfile.dateOfBirth }} 
            </div>

            <div class="grid-child-posts">
                <b>Email</b> {{ this.userProfile.email }} 
            </div>

            <div class="grid-child-posts">
                <b>Gender</b> {{ this.userProfile.gender }} 
            </div>

            <div class="grid-child-posts">
                <b>DateOfCreation</b> {{ this.userProfile.dateOfCreation }} 
            </div>

            <div class="grid-child-posts">
                <b>Biography</b> {{ this.userProfile.biography }} 
            </div>

        </div>

        <!-- View Photo Details Button -->
        <div class="form-group2" style="margin-left: 50px;">
                    <button type="login-button" class="btn btn-primary btn-block btn-large"  
                    @click="goToProfileView(this.userProfile.username)" 
                    style="width: 250px; margin-left: 600px; margin-top: 10px; "
                    > View {{this.user.username + "'s "}}Profile </button>
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
  grid-template-columns: 1fr 1fr 1fr;
  grid-gap: 20px;
  font-size: 18px;
  /* background-color: greenyellow; */
}

/* .grid-container3 {
  margin-top: 80px;
  margin-left: 120px;
  display: grid;
  grid-template-columns: 1fr 1fr;
  grid-gap: 0px;
  font-size: 15px;
  background-color: grey;
} */

.grid-child-posts2{
    font-size:12px;
    width: 109px;
}


</style>


