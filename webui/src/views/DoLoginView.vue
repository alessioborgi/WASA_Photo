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
    <!-- <div>
        <div
            class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
            <h1 class="h2">Login</h1>
        </div>
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

        <div class="mb-3">
            <label for="description" class="form-label">Username</label>
            <input type="string" class="form-control" id="username" v-model="Username" placeholder="please insert username">
        </div>

        <div>
            <button v-if="!loading" type="button" class="btn btn-primary" @click="LoginUser">
                Login
            </button>
            <LoadingSpinner v-if="loading"></LoadingSpinner>
        </div>
    </div> -->

    <div class='login'>
        <div class='login_title'>
            <span>LOGIN</span>
        </div>
        <div class='login_fields'>
            <div class='login_fields__user'>
                <div class='icon'>
                    <img src='https://s3-us-west-2.amazonaws.com/s.cdpn.io/217233/user_icon_copy.png'>
                </div>
                <input placeholder='Username' type='text'>
                    <div class='validation'>
                        <img src='https://s3-us-west-2.amazonaws.com/s.cdpn.io/217233/tick.png'>
                    </div>
            </div>
            <div class='login_fields__submit'>
                <input type='submit' value='Log In'>
                <div class='forgot'>
                    <a href='#'></a>
                </div>
            </div>
        </div>
        <div class='success'>
            <h2>Authentication Success</h2>
            <p>Welcome back</p>
        </div>
        <div class='disclaimer'>
            <p>@AlessioBorgi WASAPhoto Project</p>
        </div>
    </div>
    <div class='authent'>
        <img src='https://s3-us-west-2.amazonaws.com/s.cdpn.io/217233/puff.svg'>
        <p>Authenticating...</p>
    </div>


</template>

<style>
</style>






<!-- For doing the heart!!! -->
<!-- .love
  %p Made with <img src="https://s3-us-west-2.amazonaws.com/s.cdpn.io/217233/love_copy.png" /> -->