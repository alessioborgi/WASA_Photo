
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

            // Initializing colorBackground of the Card depending on the Gender.
            colorBackground: this.user.gender == "male" ? '#c2e9fc' : this.user.gender == "female" ? '#fbd3f0' : '#cff6cc',
            // colorBackground: this.user.gender == "male" ? '#c2e9fc' : '#fbd3f0',


			// Initializing iconFollow, that can receive two values:
            //   Follow(true):    /feather-sprite-v4.29.0.svg#user-check
            //   NotFollow(false): /feather-sprite-v4.29.0.svg#user-plus
			iconFollowing: this.user.boolFollowing == true ? '/feather-sprite-v4.29.0.svg#user-check' : '/feather-sprite-v4.29.0.svg#user-plus',
            iconFollower: this.user.boolFollower == true ? '/feather-sprite-v4.29.0.svg#user-check' : '/feather-sprite-v4.29.0.svg#user-x',
            iconBanned: this.user.boolBanned == true ? '/feather-sprite-v4.29.0.svg#lock' : '/feather-sprite-v4.29.0.svg#unlock',
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
            if (this.user.boolBanned == true) {

                try{

                    // Deleting the Ban: /users/:username/bans/:usernameBanned.
                    await this.$axios.delete(`/users/${this.username}/bans/${this.user.username}`, {
                        headers: {
                            Authorization: "Bearer " + localStorage.getItem("BearerToken")
                        }
                    })

                    // Once we have done with it, we must simply update the flag.
                    this.user.boolBanned = false;
                    this.iconBanned = '/feather-sprite-v4.29.0.svg#unlock';
                    this.$emit('refreshBan', false);

                } catch (e) {

                    // If an error is encountered, display it!
                    this.errormsg = e.toString();
                }
            } else{

                // Let's handle first the case where the user is NOT Banned.
                // We must therefore add the Ban.
                try{

                    // Deleting the Ban: /users/:username/bans/:usernameBanned.
                    await this.$axios.put(`/users/${this.username}/bans/${this.user.username}`, {}, {
                        headers: {
                            Authorization: "Bearer " + localStorage.getItem("BearerToken")
                        }
                    })

                    // Once we have done with it, we must simply update the flag.
                    this.user.boolBanned = true;
                    this.iconBanned = '/feather-sprite-v4.29.0.svg#lock';
                    this.$emit('refreshBan', true);

                } catch (e) {

                    // If an error is encountered, display it!
                    this.errormsg = e.toString();
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
            if (this.user.boolFollowing == true) {

                try{

                    // Deleting the Follow: /users/:username/followings/:usernameFollowing.
                    await this.$axios.delete("/users/"+this.username+"/followings/"+this.user.username, {
                        headers: {
                            Authorization: "Bearer " + localStorage.getItem("BearerToken")
                        }
                    })

                    // Once we have done with it, we must simply update the flag.
                    this.user.boolFollowing = false;
                    this.iconFollowing = '/feather-sprite-v4.29.0.svg#user-plus';
                    this.$emit('refreshFollowing', false);

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

        async goToProfileView(userToView) {

            // Re-address the user to the right page.
            localStorage.setItem('usernameProfileToView', userToView),
            this.$router.push({ path: `/users/${userToView}`})
        },
    }

    
}    
</script>
    

<template>

    <!-- If instead, it is all ok, Display a sort of card for each of the User Profiles(Depending on we are asking the whole list or just one). -->
    <div class="card" id="div1" :style="{backgroundColor: this.colorBackground}">
        <!-- <div class="card" id="div1" :style="{backgroundColor: this.colorBackground}"> -->

        <div class="usernameLabel">
            <!-- <b> FIXEDUSERNAME: </b>{{ user }}  -->
            <!-- <b> FIXEDUSERNAME: </b>{{ user.fixedUsername }}  -->

        </div>
        
        <!--  <div class="grid-container2"> -->

            <div class="upperPart"> 
                <div class="imageLabel">
                    <div class="profileImage">
                        <img src="https://lh3.googleusercontent.com/ytP9VP86DItizVX2YNA-xTYzV09IS7rh4WexVp7eilIcfHmm74B7odbcwD5DTXmL0PF42i2wnRKSFPBHlmSjCblWHDCD2oD1oaM1CGFcSd48VBKJfsCi4bS170PKxGwji8CPmehwPw=w200-h247-no" alt="Person" class="card__image">
                    </div>
                    <div class="profileLabel">
                        <p class="card__name" > <b>{{ user.username }}</b></p>
                    </div>            
                </div>

                <div class="rightUpperPart">

                    <div class="grid-container2">
                        <div class="grid-child-posts">
                        <b>Posts</b> {{ user.numberOfPhotos }}
                        </div>

                        <div class="grid-child-posts">
                            <b>Followings</b> {{ user.numberFollowing }} 
                        </div>

                        <div class="grid-child-posts">
                            <b>Followers</b> {{ user.numberFollowers }} 
                        </div>
                    </div>


                    <div class="grid-container2">
                        <div class="grid-child-posts2">
                            <b>Is it Banned? </b> <svg class="feather" v-if="!loading" @click="banUnbanUser" ><use :href="this.iconBanned"/></svg>
                        </div> 

                        <div class="grid-child-posts2">
                            <b>Am I Following it?</b><svg class="feather" v-if="!loading" @click="followUnfollowUser" ><use :href="this.iconFollowing"/></svg>
                        </div>

                        <div class="grid-child-posts2">
                            <b>Is it my Follower?</b><svg class="feather" v-if="!loading"><use :href="this.iconFollower"/></svg>
                        </div>                          
                    </div>
                </div>
                    
            </div>
            
            
        <div class="grid-container">

            <div class="grid-child-posts">
                <b>Name</b> {{ user.name }}
            </div>

            <div class="grid-child-posts">
                <b>Surname</b> {{ user.surname }} 
            </div>

            <div class="grid-child-posts">
                <b>Nationality</b> {{ user.nationality }} 
            </div>

            <div class="grid-child-posts">
                <b>DateOfBirth</b> {{ user.dateOfBirth }} 
            </div>

            <div class="grid-child-posts">
                <b>Email</b> {{ user.email }} 
            </div>

            <div class="grid-child-posts">
                <b>Gender</b> {{ user.gender }} 
            </div>

            <div class="grid-child-posts">
                <b>DateOfCreation</b> {{ user.dateOfCreation }} 
            </div>

            <div class="grid-child-posts">
                <b>Biography</b> {{ user.biography }} 
            </div>

        </div>

        <!-- View Photo Details Button -->
        <div class="form-group2" style="margin-left: 50px;">
                    <button type="login-button" class="btn btn-primary btn-block btn-large" v-if="!loading" 
                    @click="goToProfileView(this.user.username)" 
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
  text-align: center;
  font-size: 1.25em;
}

.card__image {
  height: 160px;
  width: 160px;
  border-radius: 50%;
  border: 5px solid #272133;
  margin-top: 20px;
  box-shadow: 0 10px 50px rgb(25, 214, 235);
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


