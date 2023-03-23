<!-- Starting of the actual Login Page Handling. -->
<script>

import ErrorMsg from '../components/ErrorMsg.vue'
import LoadingSpinner from '../components/LoadingSpinner.vue'

// Declaration of the export set.
export default {

    // Including some components that will be used in the page.
    components: {
        LoadingSpinner,
        ErrorMsg,
    },

    // Describing what are the Return variables.
    data: function() {
        return {
            // Variables handling the eventual error message and the Loading of the action.
            errormsg: "",
            loading: false,

            // Logged user Info.
            username: "",
            uuid: "",
        }
    },

    // Declaration of the methods that will be used.
    methods: {

        // Declaration of the DoLogin page. 
        async doLogin() {

            // Re-initializing variables to their default value.
			this.errormsg = "";
			this.loading = true;


            try {
                
                // if (this.username.length < 3 ) {
                //     this.errormsg = "Please insert a valid Username. It must be non-empty, it must be more than 3 letters, and it must be a different username w.r.t. the already present ones. It should respect the following Regex: [a-zA-Z0-9._]{5,20}$"
                // } else {
                    // In the case the result is positive, we post the username received to the GO page.
                    let response = await this.$axios.post("/session/", { 
                        username: this.username 
                    });
                    
                    // Setting the uuid (Bearer Token) received as response by the Post action.
                    this.uuid  = response.data,
                    localStorage.setItem('BearerToken', this.uuid),
                    localStorage.setItem('Username', this.username),
                    localStorage.setItem('usernameProfileToView', this.username),
                    
                    // Re-addressing the page to the personal profile page of a user.
                    this.$router.replace({ path: `/users/${this.username}` })
                // }

            } catch (e) {

                // In case of error, retrieve it.
                if (e.response && e.response.status === 400) {
					this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us. This may be due to an incorrect insertion. Be sure to resepct the rules!" + e.toString();
				} else if (e.response && e.response.status === 500) {
					this.errormsg = "An Internal Error occurred. We are sorry for the inconvenient. Please try again later." + + e.toString();
				} else {
					this.errormsg = e.toString();
				}

            }

            // Setting again the Loading flag to false.
            this.loading = false;
        }
    }
}

</script>

<!-- Actual Page for handling the page setting. -->
<template>
    
    <div class="login" v-if="!loading">

        <h1>WASA Photo LOGIN</h1>
        <img src="./img/wasa-logo.png" alt="">

        <!-- Creation of the form for the Login. -->
        <form method="post" class="login-form">

            <!-- Creation of the place where to type the Username. -->
            <input type="text" id="usernameLabel" v-model="username" placeholder="Insert Username..." required="required" class="form-control">

            <!-- Creation of the Login Button linked to the doLogin action. -->
            <button type="login-button" class="btn btn-primary btn-block btn-large" @click="doLogin">Login</button>
        </form>

        <div>
            <LoadingSpinner v-if="loading"></LoadingSpinner>
            <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        </div>
    </div>
    
</template>

<!-- Declaration of the style(scoped) to use. -->
<style scoped>
    @import '../assets/login.css';
</style>


<!-- For doing the heart!!! -->
<!-- .love
  %p Made with <img src="https://s3-us-west-2.amazonaws.com/s.cdpn.io/217233/love_copy.png" /> -->


