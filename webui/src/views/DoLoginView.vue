<style>
  @import '/Users/alessioborgi/Documents/GitHub/WASA_Photo/webui/src/assets/login.css';
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
                UserID: null,
                Username: null,
            }
        }
    },
    methods: {
        LoginUser: async function () {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios.post("/session/", {
                    username: this.Username,
                });
				this.UserID  = response.data,
                localStorage.setItem('Authorization', this.UserID),
                this.$router.push({ path: '/users/'+this.Username })
            } catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        }
    }
}


</script>

<template>
    
    <div class="login">
        <h1>Login</h1>
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <form method="post">
            <input type="string" id="username" name="u" v-model="Username" placeholder="Username..." required="required" />
            <button type="submit" class="btn btn-primary btn-block btn-large" @click="LoginUser">Login</button>
            <LoadingSpinner v-if="loading"></LoadingSpinner>
        </form>
    </div>
</template>


<!-- For doing the heart!!! -->
<!-- .love
  %p Made with <img src="https://s3-us-west-2.amazonaws.com/s.cdpn.io/217233/love_copy.png" /> -->