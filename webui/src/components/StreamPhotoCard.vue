
<script>

import InfoMsg from './InfoMsg.vue'

export default {

    props: ['photo'],   //{ "photoid", "fixedUsername", "username", "filename", "uploadDate", "phrase", "numberLikes", "numberComments"}

    components: {
        InfoMsg
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

    methods: {

        // likeUnLikePhoto function: It has the role to add or delete a like depending on the boolLike value.
        async likeUnLikePhoto(){

            // Initializing the two errormessage and loading variables.
            this.errormsg= "";
            this.loading= true;

            // Let's handle first the case where we are currently liking the photo.
            // We must therefore delete the Like.
            if (this.photo.boolLike == true) {

                try{

                    // Deleting the Like: /users/:username/photos/:photoid/likes/:usernameLiker.
                    await this.$axios.delete("/users/"+this.photo.username+"/photos/"+this.photo.photoid+"/likes/"+this.username, {
                        headers: {
                            Authorization: "Bearer " + localStorage.getItem("BearerToken")
                        }
                    })

                    // Once we have done with it, we must simply update the flag.
                    this.photo.boolLike = false;
                    this.iconFollowing = '/feather-sprite-v4.29.0.svg#user-plus';
                    this.$emit('refreshLike', false);

                } catch (e) {

                    // If an error is encountered, display it!
                    this.errormsg = e.toString();
                }

            } else{

                // Let's handle first the case where the user is NOT Followed by us.
                // We must therefore add the Follow.
                try{

                    // Adding the Follow: /users/:username/followings/:usernameFollowing.
                    await this.$axios.put(`/users/${this.username}/followings/${this.user.username}`, {}, {
                        headers: {
                            Authorization: "Bearer " + localStorage.getItem("BearerToken")
                        }
                    })

                    // Once we have done with it, we must simply update the flag.
                    this.user.boolFollowing = true;
                    this.iconFollowing = '/feather-sprite-v4.29.0.svg#user-check';
                    this.$emit('refreshFollowing', true);
                    

                } catch (e) {

                    // If an error is encountered, display it!
                    this.errormsg = e.toString();
                    }
            }

            // Once the entire operation has finished, re-set the "loading" flag to false, in such a way to continue.
            this.loading = false;

        },

        async goToViewPhotoDetails() {

            // Re-address the user to the right page.
            // alert(`/users/${userToView}/photo/${this.photo.photoid}`)
            localStorage.setItem('usernameProfileToView', this.photo.username),
            this.$router.push({ path: `/users/${this.photo.username}/photo/${this.photo.photoid}`})
        },

        async goToProfile() {

            // Re-address the user to the right page.
            // alert(`/users/${userToView}/photo/${this.photo.photoid}`)
            localStorage.setItem('usernameProfileToView', this.photo.username),
            this.$router.push({ path: `/users/${this.photo.username}`})
        },
                  
    }
}

     
</script>
    

<template>

    <!-- If instead, it is all ok, Display a sort of card for each of the User Profiles(Depending on we are asking the whole list or just one). -->
    <div class="card" id="div1" >

        <div class="usernameLabel">
            <b> FIXEDUSERNAME: </b>{{ this.photo }} 
            <!-- <b> FIXEDUSERNAME: </b>{{ photo }}  -->

        </div>
        
        <div class="upperPart"> 
            <div class="imageLabel">
                <div class="profileImage">
                    <!-- In this way works -->
                    <!-- <img src="../../u1-photo-0.png" alt="Person" class="card__image"/> -->
                    
                    <!-- In this other way it does not :( -->
                    <!-- <img src="Users/alessioborgi/Documents/GitHub/WASA_Photo/service/api/photos/u1-photo-0.png" alt="Person" class="card__image"/> -->
                    <img src="https://lh3.googleusercontent.com/ytP9VP86DItizVX2YNA-xTYzV09IS7rh4WexVp7eilIcfHmm74B7odbcwD5DTXmL0PF42i2wnRKSFPBHlmSjCblWHDCD2oD1oaM1CGFcSd48VBKJfsCi4bS170PKxGwji8CPmehwPw=w200-h247-no" alt="Person" class="card__image">
                    
                    <!-- <img :src= user.photoProfile alt="Person" class="card__image"/> -->
                    <!-- <img src="http://localhost/WASA_Photo/service/api/photos" alt="Person" class="card__image"> -->
                </div>
                          
            </div>
            <div class="rightUpperPart">

                <!-- Grid for containing number of likes and of comments. -->
                <div class="grid-container2">
                    <div class="grid-child-posts">
                        <svg class="feather" v-if="!loading"
                        style="fill: #ff0000;">
                        <use href="/feather-sprite-v4.29.0.svg#heart"/></svg>
                        <b> Likes</b> 
                        {{ photo.numberLikes }} 
                    </div>

                    <div class="grid-child-posts">
                        <svg class="feather" v-if="!loading" 
                        style="color:green; fill:green"><use href="/feather-sprite-v4.29.0.svg#message-circle"/></svg>
                        <b> Comments</b> 
                        {{ photo.numberComments }} 
                    </div>
                </div>

                <!-- Phrase -->
                <div class="grid-child-posts3">
                    <b>Phrase</b> {{ photo.phrase }} 
                </div>

                <!-- Username -->
                <div class="grid-child-posts3" style="margin-top: 20px;">
                    <b>Username</b> {{ photo.username }} 
                </div>

                <!-- Upload Date -->
                <div class="grid-child-posts3" style="margin-top: 20px;">
                    <b>Upload Date</b> {{ photo.uploadDate }} 
                </div>

                <div class="grid-container2">
                    <div class="grid-child-posts">
                        <svg class="feather" v-if="!loading" 
                        @click="likeUnLikePhoto"
                        :style="{fill: this.photo.fillHeart}">
                        <use href="/feather-sprite-v4.29.0.svg#heart"/></svg>
                        <b> Put Like</b> 
                    </div>

                    <div class="grid-child-posts">
                        <svg class="feather" v-if="!loading" 
                        style="color:green; fill:white;">
                        <use href="/feather-sprite-v4.29.0.svg#message-circle"/></svg>
                        <b> Add Comment</b> 
                    </div>
                </div>

                <!-- View Photo Details Button -->
                <div class="form-group2">
                    <button type="login-button" class="btn btn-primary btn-block btn-large" v-if="!loading" 
                    @click="goToViewPhotoDetails()" 
                    style="width: 200px; margin-top: 20px;"
                    :photo="this.photo"
                    > View Photo Details </button>
                </div>

                <!-- View Profile Button -->
                <div class="form-group2">
                    <button type="login-button" class="btn btn-primary btn-block btn-large" v-if="!loading" 
                    @click="goToProfile" 
                    style="width: 250px; margin-left: 250px; margin-top: -43px;"
                    :photo="this.photo"
                    > View {{this.photo.username}}'s Profile </button>
                </div>
                
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

.grid-child-posts{
    width: 145px;
}
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
	color: #ff0000;
}

#div1{
    background: #c2e9fc;
}

.card {
  background-color: #ffffff; 
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
  text-align: center;
  font-size: 1.25em;
}

.card__image {
  height: 350px;
  width: 350px;
  margin-top: 40px;
  border: 5px solid #272133;
  margin-left: -200px;
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


