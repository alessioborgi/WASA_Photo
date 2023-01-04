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
    <!-- <header>
	<div class="container">

		<div class="profile">

			<div class="profile-image">

				<img src="./img/wasa-logo.png" alt="">

			</div>
        
        </div>

	</div>
-->

    <div class="login">
        <!-- <div class="profile-image"> -->
        <img src="./img/wasa-logo.png" alt="">
        <!-- </div> -->
        <!-- <h1>Login</h1> -->
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <form method="post" class="login-form">
            <input type="string" id="username" name="u" v-model="username" placeholder="Username..." required="required" />
            <button type="submit" class="btn btn-primary btn-block btn-large" @click="LoginUser">Login</button>
            <LoadingSpinner v-if="loading"></LoadingSpinner>
        </form>
    </div>
</template>


<!-- For doing the heart!!! -->
<!-- .love
  %p Made with <img src="https://s3-us-west-2.amazonaws.com/s.cdpn.io/217233/love_copy.png" /> -->


<!-- 

  <template>
    <header>
	<div class="container">

		<div class="profile">

			<div class="profile-image">

				<img src="./img/wasa-logo.png" alt="">

			</div>
        
        </div>

	</div>

</header>

    <div class="login">
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
        <form method="post">
            <input type="string" id="username" name="u" v-model="username" placeholder="Username..." required="required" />
            <button type="submit" class="btn btn-primary btn-block btn-large" @click="LoginUser">Login</button>
            <LoadingSpinner v-if="loading"></LoadingSpinner>
        </form>
    </div>
</template> -->
