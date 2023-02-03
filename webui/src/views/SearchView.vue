<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
            username: localStorage.getItem('Username'),
            BearerToken: localStorage.getItem('BearerToken'),
			users: [],
			usersProfiles: [],
            userProfile: {
                fixedUsername: "",
                username: "",
                photoProfile: "",
                biography: "",
                dateOfCreation: "",
                numberOfPhotos: 0,
                numberFollowers: 0,
                numberFollowing: 0,
                name: "",
                surname: "",
                dateOfBirth: "",
                email: "",
                nationality: "", 
                gender: "",
            }
		}
	},
	methods: {
		load() {
			return load
		},

        // GetUsers Function
		async getUsers() {
			this.loading = true;
			this.errormsg = null;
			try {

                let response = await this.$axios.get("/users/", {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					}
				})

				this.users = response.data;

				for (let i = 0; i < this.users.length; i++) {

					try{

						let responseProfile = await this.$axios.get("/users/"+this.users[i], {
							headers: {
								Authorization: "Bearer " + localStorage.getItem("BearerToken")
							}
						})

						this.usersProfiles.push(responseProfile.data);

					} catch (e) {
						this.errormsg = e.toString();
					}

					this.loading = false;
				} 
				
			} catch (e) {
				this.errormsg = e.toString();
			}

			this.loading = false;
		},

        // GetUserProfile Function
        async getUserProfile(usernameToSearch) {
			this.loading = true;
			this.errormsg = null;
            // this.usernameToSearch = usernameToSearch;
			try {

                let response = await this.$axios.get("/users/"+usernameToSearch, {
						headers: {
							Authorization: "Bearer " + localStorage.getItem("BearerToken")
						}
					})

				this.userProfile = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},


		async newItem() {
			this.$router.push("/new");
		},
		async deleteFountain(id) {
			this.loading = true;
			this.errormsg = null;
			try {
				await this.$axios.delete("/users/" + id);

				await this.getUsers();
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		}
	},
	mounted() {
		this.getUsers()
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Users list</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="getUsers">
						Refresh
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="newItem">
						New
					</button>
				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<LoadingSpinner v-if="loading"></LoadingSpinner>

		<div class="card" v-if="users.length === 0">
			<div class="card-body">
				<p>No User in the database.</p>

				<a href="javascript:" class="btn btn-primary" @click="newItem">Create a new User</a>
			</div>
		</div>

		<div class="card" v-if="!loading" v-for="u in usersProfiles">
            <!-- {{ this.userProfile = getUserProfile(u) }} -->
			<div class="card-header">
				User
			</div>
			<div class="card-body">
				<p class="card-text">
					Username: {{ u.username }}<br />
                    Name: {{ u.name }}
				</p>
			</div>
		</div>
	</div>
</template>

<style scoped>
.card {
	margin-bottom: 20px;
}
</style>
