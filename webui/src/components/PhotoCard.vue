
<script>

export default {

    props: ['photo', 'userOwnerFlag', 'userProfile', 'photoListCurrent'],   //{ "photoid", "fixedUsername", "username", "filename", "uploadDate", "phrase", "numberLikes", "numberComments"}

    components: {
    },
    
	// Describing what are the Return variables.
	data: function() {
		return {	
            
            // Initializing the two errormessage and loading variables.
			errormsg: "",
			loading: false,

			// Retrieving from the Cache the Username and the Bearer Authenticaiton Token.
            // username: localStorage.getItem('Username'),
            username: localStorage.getItem('Username') == localStorage.getItem('usernameProfileToView') ? localStorage.getItem('Username') : localStorage.getItem('usernameProfileToView'),
            BearerToken: localStorage.getItem('BearerToken'),

            // Initializing variable for handling the deletion of the Photo.
            deletePhotoBool: false,

            // Initializing the newPhotoList.
            newPhotoList : [],

            // Initializing the image as a blob object, and declaring an object URL form it.
            photoBlobLink: "",
		}
	},

    methods: {

        // deletePhotoAlert: This method allows us to open an alert that will be used to alert about the Photo Deletion.
        async deletePhotoAlert() {

            if (confirm("Your Photo will be deleted from your Personal Profile. Are you sure?")){
                this.deletePhotoBool = true;
                alert("Photo Correctly Deleted");
            } else {
                this.deletePhotoBool = false;
            }

        },


        // deletePhoto: This method allows us to delete a Photo from the user Profile.
        async deletePhoto() {

            // Re-initializing variables to their default value.
            this.errormsg = "";
            this.loading = true;

            // Alerting the user to have the confirmation that he/she wants to Delete the photo.
            await this.deletePhotoAlert();

            try {
                
                // In the case the result is positive, we post the username received to the GO page.
                if (this.deletePhotoBool == true){

                    // /users/:username/photos/:photoid
                    // await this.$axios.delete("/users/"+this.username+"/photos/"+this.photo.photoid, {
                    await this.$axios.delete(`/users/${this.username}/photos/${this.photo.photoid}`, {
                        headers: {
                        Authorization: "Bearer " + localStorage.getItem("BearerToken")
                    }})
                             
                    // Eliminate from the list the photo.
                    this.newPhotoList = await this.removeObjectWithId(this.photoListCurrent, this.photo.photoid);
                    // Re-addressing the page to the personal profile page of a user.
                    this.$router.replace({ path: `/users/${this.username}` })
                    this.userProfile.numberOfPhotos = this.userProfile.numberOfPhotos - 1
                    this.$emit('refreshProfile', this.userProfile);
                    this.$emit('refreshPhotos', this.newPhotoList);
                    // eventBus.$emit("refreshNumberPhotos", this.numberOfPhotos - 1);

                }

            } catch (e) {

                if (e.response && e.response.status === 400) {
					this.errormsg = "Request error, please Login before doing some action or ask to delete a valid photo." + e.toString();
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

        // goToViewPhotoDetails: This function is used for redirecting the user to the specific photo's view.
        async goToViewPhotoDetails(idPhoto) {

            // Saving on the Local Cache the IdPhoto that we have to retrieve the information then in the next view.
            localStorage.setItem('idPhoto', idPhoto),

            // Re-address the user to the right page.
            this.$router.push({ path: `/users/${this.username}/photo/${this.photo.photoid}`})
        },

        // removeObjectWithId: This function is used for removing from the list of photos a specific one.
        async removeObjectWithId(arr, id) {
            
            const objWithIdIndex = arr.findIndex((obj) => obj.photoid === id);

            if (objWithIdIndex > -1) {
                arr.splice(objWithIdIndex, 1);
            }
            return arr;
        },


        async getPhotoView() {

            // Re-initializing variables to their default value.
            this.errormsg = "";
            this.loading = true;

            try {
                
                // Getting the image view from the Back-End.
                // /users/:username/photos/:photoid/view
                let response = await this.$axios.get(`/users/${this.username}/photos/${this.photo.photoid}/view`, {
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
    <div class="card" id="div1" >

        <div class="usernameLabel">
            <!-- <b> FIXEDUSERNAME: </b>{{ this.photo.filename }}  -->
            <!-- <b> FIXEDUSERNAME: </b>{{ this.userOwnerFlag }}  -->

        </div>
        
        <div class="upperPart"> 
            <div class="imageLabel">
                <div class="profileImage">
                    <!-- In this way works -->
                    <!-- <img src="../../../tmp/u1-photo-0-photo-profile.jpg" alt="Person" class="card__image"/> -->
                    <img :src=this.photoBlobLink class="card__image" />
                </div>
            </div>
                
            <div class="rightUpperPart">

                <!-- Grid for containing number of likes and of comments. -->
                <div class="grid-container2" style="margin-left: 50px;" v-if="!loading">
                    <div class="grid-child-posts">
                        <svg class="feather" style="fill: #ff0000;"><use href="/feather-sprite-v4.29.0.svg#heart"/></svg>
                        <b> Likes</b> 
                        {{ photo.numberLikes }} 
                    </div>

                    <div class="grid-child-posts">
                        <svg class="feather" style="color:green; fill: green;"><use href="/feather-sprite-v4.29.0.svg#message-circle"/></svg>
                        <b> Comments</b> 
                        {{ photo.numberComments }} 
                    </div>
                </div>

                <!-- Upload Date -->
                <div class="grid-child-posts3" style="margin-top: 50px; margin-left: 50px;">
                    <b>Upload Date</b> {{ photo.uploadDate }} 
                </div>

                <!-- Phrase -->
                <div class="grid-child-posts3" style="margin-left: 50px; margin-top: 10px;">
                    <b>Phrase</b> {{ photo.phrase }} 
                </div>

                <!-- View Photo Details Button -->
                <div class="form-group2" style="margin-left: 50px;" v-if="!loading" >
                    <button type="login-button" class="btn btn-primary btn-block btn-large" 
                    @click="goToViewPhotoDetails(this.photo.photoid)" 
                    style="width: 250px; margin-top: 165px;"
                    :photo="this.photo"
                    > View Photo Details </button>
                </div>

                <!-- Deletion -->
                <div class="grid-child-posts3" v-if="!loading && userOwnerFlag !== true" >
                    <svg class="feather" 
                        @click="deletePhoto" 
                        style="margin-left: 450px; margin-top: -190px; color:midnightblue">
                        <use href="/feather-sprite-v4.29.0.svg#trash-2"/></svg>
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

.card__image {
  height: 350px;
  width: 350px;
  margin-top: 40px;
  border: 5px solid #272133;
  margin-left: -150px;
}

.grid-container2 {
margin-top: 50px;
  display: grid;
  grid-template-columns: 1fr 1fr 1fr 1fr;
  grid-gap: 50px;
  width: 600px;
  font-size: 18px;
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


