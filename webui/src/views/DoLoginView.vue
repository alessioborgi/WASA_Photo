<!-- Declaration of the style(scoped) to use. -->
<style scoped>
        @import '../assets/login.css';
</style>

<!-- Starting of the actual Login Page Handling. -->
<script>


import LoadingSpinner from '../components/LoadingSpinner.vue'

// Declaration of the export set.
export default {

    // Including some components that will be used in the page.
    components: {
        LoadingSpinner,
    },

    // Describing what are the Return variables.
    data: function() {
        return {
            // Variables handling the eventual error message and the Loading of the action.
            errormessage: null,
            loading: false,

            // Logged user Info.
            loginUsername: {
                username: "",
                uuid: "",
            },
        }
    },

    // Declaration of the methods that will be used.
    methods: {

        // Declaration of the DoLogin page. 
        async doLogin() {

            // Declaring the variables responsible for the Loading handling of the Loading Spinner components and for the error message.
            this.loading = true;
            this.errormessage = null;

            try {
                
                // In the case the result is positive, we post the username received to the GO page.
                let response = await this.$axios.post("/session/", { username: this.username });
                
                // Setting the uuid (Bearer Token) received as response by the Post action.
                this.uuid  = response.data,
                localStorage.setItem('Authorization', this.uuid),

                // Re-addressing the page to the personal profile page of a user.
                this.$router.push({ path: '/users/'+ this.username })

            } catch (e) {

                // In case of error, retrieve it.
                this.errormessage = e.toString();
            }

            // Setting again the Loading flag to false.
            this.loading = false;
        }
    },

    mounted() {
		this.refresh()
	}
}

</script>

<!-- Actual Page for handling the page setting. -->
<template>
    
    <div class="login">
        <h1>WASA Photo</h1>
        <img src="./img/wasa-logo.png" alt="">
        <!-- <ErrorMessage v-if="errormessage" :msg="errormessage"></ErrorMessage> -->

        <!-- Creation of the form for the Login. -->
        <form method="post" class="login-form">

            <!-- Creation of the place where to type the Username. -->
            <input type="text" id="username" v-model="username" placeholder="Insert Username..." required="required" class="form-control">

            <!-- Creation of the Login Button linked to the doLogin action. -->
            <button type="submit" class="btn btn-primary btn-block btn-large" @click="doLogin">Login</button>
            <!-- <LoadingSpinner v-if="loading"></LoadingSpinner> -->

        </form>

        <!-- Insertion of the Copyright footer -->
        <!-- <footer class="text-center card-footer fixed-bottom">
            <p>&copy Alessio Borgi</p>
        </footer> -->

    </div>

</template>


<!-- For doing the heart!!! -->
<!-- .love
  %p Made with <img src="https://s3-us-west-2.amazonaws.com/s.cdpn.io/217233/love_copy.png" /> -->


