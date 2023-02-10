
<script>

import InfoMsg from '../components/InfoMsg.vue'

export default {

    props: ['user'],

    components: {
        InfoMsg
    },

    
	// Describing what are the Return variables.
	data: function() {
		return {	
            
            errormsg: "",
            loading: true,
		}
	},

    methods: {
        // showAlert() {
        //     alert('test')
        // }
        createAlert() {
            var input = prompt("Please enter your name:", "");
            if (input != null) {
                alert("Hello " + input + "! How are you today?");
            }
        },

        // Declaration of the DoLogin page. 
        async setUsername() {

            // Re-initializing variables to their default value.
            this.errormsg = "";
            this.loading = true;

            try {
                
                // In the case the result is positive, we post the username received to the GO page.
                let response = await this.$axios.put("/session/", { 
                    username: this.username 
                });
                
                // Setting the uuid (Bearer Token) received as response by the Post action.
                this.uuid  = response.data,
                localStorage.setItem('BearerToken', this.uuid),
                localStorage.setItem('Username', this.username),

                // Re-addressing the page to the personal profile page of a user.
                
                // Re-addressing the page to the personal profile page of a user.
                this.$router.replace({ path: `/users/${this.username}` })

            } catch (e) {

                // In case of error, retrieve it.
                this.errormessage = e.toString();
            }

            // Setting again the Loading flag to false.
            this.loading = false;
        },
                
                
    }
}

     
</script>
    

<template>

    <!-- If instead, it is all ok, Display a sort of card for each of the User Profiles(Depending on we are asking the whole list or just one). -->
    <div class="card" id="div1">

        <div class="usernameLabel">
            <!-- <b> FIXEDUSERNAME: </b>{{ user.gender }}  -->
            <!-- <b> FIXEDUSERNAME: </b>{{ user.fixedUsername }}  -->

        </div>
        
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


                <div class="grid-child-posts3">
                    <b>Biography</b> {{ user.biography }} 
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

            <div class="grid-child-posts"></div>

            <div class="grid-child-posts">
                        <!-- <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#settings" @click="replaceLogin"/></svg> -->
                        <nav>
                            <menu>
                                <menuitem id="demo1">
                                    <a><svg class="feather"><use href="/feather-sprite-v4.29.0.svg#settings"/></svg></a>
                                    <menu>

                                        <menuitem>
                                            <!-- <button @click="showAlert"> -->
                                                <button @click="createAlert">
                                                <b>Set Username</b>
                                            </button>
                                        </menuitem>

                                        <menuitem>
                                            <button @click="showAlert">
                                                <b>Set Profile</b>
                                            </button>
                                        </menuitem>

                                        <menuitem><a><b>Delete Profile</b></a></menuitem>
                                    </menu>
                                </menuitem>
                            
                            </menu>
	                    </nav>                  
            </div>

        </div>

        

    </div>

</template>





<style scoped>


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
   background:#c2e9fc;
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
   border-top: 10px solid white;
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


