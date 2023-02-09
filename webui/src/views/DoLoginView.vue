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
            errormessage: "",
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
                
                // In the case the result is positive, we post the username received to the GO page.
                let response = await this.$axios.post("/session/", { 
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
        }
    }
}

</script>

<!-- Actual Page for handling the page setting. -->
<template>
    
    <div class="login">

        <ErrorMsg v-if="errormessage" :msg="errormessage"></ErrorMsg>

        <h1>WASA Photo LOGIN</h1>
        <img src="./img/wasa-logo.png" alt="">

        <!-- Creation of the form for the Login. -->
        <form method="post" class="login-form">

            <!-- Creation of the place where to type the Username. -->
            <input type="text" id="usernameLabel" v-model="username" placeholder="Insert Username..." required="required" class="form-control">

            <!-- Creation of the Login Button linked to the doLogin action. -->
            <button type="login-button" v-if="!loading" class="btn btn-primary btn-block btn-large" @click="doLogin">Login</button>
            <LoadingSpinner v-if="loading"></LoadingSpinner>
        </form>

    </div>

</template>

<!-- Declaration of the style(scoped) to use. -->
<style scoped>
    @import '../assets/login.css';
</style>


<!-- For doing the heart!!! -->
<!-- .love
  %p Made with <img src="https://s3-us-west-2.amazonaws.com/s.cdpn.io/217233/love_copy.png" /> -->


