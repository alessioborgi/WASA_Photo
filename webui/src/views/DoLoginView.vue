<style scoped>
        @import '../assets/login.css';
</style>

<script>
export default {
    data: function() {
        return {
            errormsg: null,
            detailedmsg: null,
            loading: false,
            id : 10,
            User: {
                uuid: null,
                username: null,
            }
        }
    },
    methods: {
        LoginUser: async function () {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios.post("/session/", {
                    username: this.username,
                });
				this.uuid  = response.data,
                localStorage.setItem('Authorization', this.uuid),
                this.$router.push({ path: '/users/'+this.username })
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        }
    },
    mounted() {
		this.refresh()
	}
}


</script>


<template>
    
    <div class="login">
        <h1>WASA Photo</h1>
        <img src="./img/wasa-logo.png" alt="">
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

        <form method="post" class="login-form">

            <input type="string" id="username" name="u" v-model="username" placeholder="Username..." required="required" />
            <button type="submit" class="btn btn-primary btn-block btn-large" @click="LoginUser">Login</button>
            <LoadingSpinner v-if="loading"></LoadingSpinner>

        </form>

        <footer class="text-center card-footer fixed-bottom">
            <p>&copy Alessio Borgi</p>
        </footer>
        
    </div>

</template>


<!-- For doing the heart!!! -->
<!-- .love
  %p Made with <img src="https://s3-us-west-2.amazonaws.com/s.cdpn.io/217233/love_copy.png" /> -->


