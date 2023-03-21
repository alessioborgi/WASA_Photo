<!-- Starting of the actual Search Page Handling. -->
<script>


// Declaration of the export set.
export default {

	// Describing what are the Return variables.
	data: function() {
		return {

			// Initializing the two errormessage and loading variables.
			errormsg: "",
			loading: false,

			// Retrieving from the Cache the Username and the Bearer Authenticaiton Token.
			// Notice that here I need also to keep track of the username that is Logged since anyone(that is not banned, of course) can view others profiles.
            BearerToken: localStorage.getItem('BearerToken'),
			usernameLogged: localStorage.getItem('Username'),
			username: localStorage.getItem('Username') == localStorage.getItem('usernameProfileToView') ? localStorage.getItem('Username') : localStorage.getItem('usernameProfileToView'),

			// Initializing flag that allows to see whether the user that is accessing the page is the actual user owner or not.
			userOwnerFlag: localStorage.getItem('Username') == localStorage.getItem('usernameProfileToView') ? true : false,

			// Initializing variable for handling the UserProfile retrieval.
			userProfile: { fixedUsername: "", username: "", photoProfile: "", biography: "", dateOfCreation: "", numberOfPhotos: 0, numberFollowers: 0, numberFollowing: 0, name: "", surname: "", dateOfBirth: "", email: "", nationality: "", gender: ""},
			colorBackground: '',
			colorPosts: '',

			// Initializing a list that will handle the links to the photos.
			photoListLinks: [],

			// Initializing the variables for the Photos Analytics.
			totNumberLikes: 0,
			totNumberComments: 0,
			numberLikesPerPostMean: 0,
			numberCommentsPerPostMean: 0,
			maxNumberLikes: 0,
			minNumberLikes: 0,
			meanNumberLikes: 0,
			medianNumberLikes: 0,
			modeNumberLikes: 0,
			varianceNumberLikes: 0,
			stdNumberLikes: 0,
			rangeNumberLikes: 0,

			maxNumberComments: 0,
			minNumberComments: 0,
			meanNumberComments: 0,
			medianNumberComments: 0,
			modeNumberComments: 0,
			varianceNumberComments: 0,
			stdNumberComments: 0,

			sortedListByLikes: [],
			sortedListByComments: [],
			frequencyLikes: {},
			frequencyComments: {},
			maxFrequencyLikes: 0,
			maxFrequencyComments: 0,
			rangeNumberComments: 0,
        }
	},

	// Declaration of the methods that will be used.
	methods: {

		// GetUserProfile Function: It retrieves the whole profile of the Logged username.
        async getUserProfile() {

			// Re-initializing variables to their default value.
			this.errormsg = "";
			this.loading = true;

			try{

				// Retrieving the Profile from the Back-end.
				let responseProfile = await this.$axios.get("/users/"+this.username, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					}
				})

				
				// Let's add up to the "userProfiles" array the response of the profile. Note that it will be an array with only this element.
				this.userProfile = responseProfile.data;
	
			} catch (e) {

				if (e.response && e.response.status === 400) {
					this.errormsg = "Request error, please Login before doing some action or ask to get the profile of a valid user." + e.toString();
                } else if (e.response && e.response.status === 403) {
                    this.errormsg = "An Unauthorized Action has been blocked. You are not allowed to do this action because you are not the profile's owner." + e.toString();
                } else if (e.response && e.response.status === 204) {
                    this.errormsg = "In the Internal DB there is not anymore the content you have asked." + e.toString();
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later." + e.toString();
                } else {
                    this.errormsg = e.toString();
                }

			}

			// Setting again the Loading flag to false.
			this.loading = false;
		},

		async getPhotoLinks() {

			// Re-initializing variables to their default value.
			this.errormsg = "";
			this.loading = true;

			try{

				// Retrieving the Profile from the Back-end.
				let responsePhotoList = await this.$axios.get(`/users/${this.username}/photos/`, {
					headers: {
						Authorization: "Bearer " + localStorage.getItem("BearerToken")
					}
				})

				// Getting the list of photos.
				this.photoListLinks = responsePhotoList.data;

				// COMPUTING PHOTOS ANALYTICS.
				// Retrieving the sum of comments.
				this.totNumberComments = this.photoListLinks.reduce((accumulator, currentValue) => accumulator + currentValue.numberComments, 0);

				// Retrieving the sum of likes.
				this.totNumberLikes = this.photoListLinks.reduce((accumulator, currentValue) => accumulator + currentValue.numberLikes, 0);

				// Getting the Mean, Max and Min, Median of Likes and Comments.
				this.numberCommentsPerPostMean = this.totNumberComments / this.userProfile.numberOfPhotos;
				this.numberLikesPerPostMean = this.totNumberLikes / this.userProfile.numberOfPhotos;
				this.maxNumberComments = this.photoListLinks.reduce((accumulator, currentValue) => { 
					if (currentValue.numberComments > accumulator) {
    					return currentValue.numberComments;
  					} else {
    					return accumulator;
  					}
				}, 0);
				this.maxNumberLikes = this.photoListLinks.reduce((accumulator, currentValue) => { 
					if (currentValue.numberLikes > accumulator) {
    					return currentValue.numberLikes;
  					} else {
    					return accumulator;
  					}
				}, 0);
				this.minNumberComments = this.photoListLinks.reduce((accumulator, currentValue) => { 
					if (currentValue.numberComments < accumulator) {
    					return currentValue.numberComments;
  					} else {
    					return accumulator;
  					}
				}, 0);
				this.minNumberLikes = this.photoListLinks.reduce((accumulator, currentValue) => { 
					if (currentValue.numberLikes < accumulator) {
    					return currentValue.numberLikes;
  					} else {
    					return accumulator;
  					}
				}, 0);
				// Sort the list based on the counter attribute
				this.sortedListByComments = this.photoListLinks.sort((a, b) => a.numberComments - b.numberComments);

				let middleIndex = Math.floor(this.photoListLinks.length / 2);

				if (this.photoListLinks.length % 2 === 0) {
					// If the list has an even number of elements, take the average of the two middle values
					this.medianNumberComments = (this.photoListLinks[middleIndex - 1].numberComments + this.photoListLinks[middleIndex].numberComments) / 2;
				} else {
					// If the list has an odd number of elements, take the middle value
					this.medianNumberComments = this.photoListLinks[middleIndex].numberComments;
				}

				// Sort the list based on the counter attribute
				this.sortedListByLikes = this.photoListLinks.sort((a, b) => a.numberLikes - b.numberLikes);

				if (this.photoListLinks.length % 2 === 0) {
					// If the list has an even number of elements, take the average of the two middle values
					this.medianNumberLikes = (this.photoListLinks[middleIndex - 1].numberLikes + this.photoListLinks[middleIndex].numberLikes) / 2;
				} else {
					// If the list has an odd number of elements, take the middle value
					this.medianNumberLikes = this.photoListLinks[middleIndex].numberLikes;
				}

				// Getting Modes.
				this.modeNumberComments = this.photoListLinks.forEach((obj) => {
					// Count the frequency of each unique counter value
					this.frequencyComments[obj.numberComments] = (this.frequencyComments[obj.numberComments] || 0) + 1;

					// Update the mode value and max frequency if the current frequency is higher than the previous max frequency
					if (this.frequencyComments[obj.numberComments] > maxFrequencyComments) {
						this.maxFrequencyComments = this.frequencyComments[obj.numberComments];
						this.modeNumberComments = obj.numberComments;
					}
				});

				this.modeNumberLikes = this.photoListLinks.forEach((obj) => {
					// Count the frequency of each unique counter value
					this.frequencyLikes[obj.numberLikes] = (this.frequencyLikes[obj.numberLikes] || 0) + 1;

					// Update the mode value and max frequency if the current frequency is higher than the previous max frequency
					if (this.frequencyLikes[obj.numberLikes] > this.maxFrequencyLikes) {
						this.maxFrequencyLikes = this.frequencyLikes[obj.numberLikes];
						this.modeNumberLikes = obj.numberLikes;
					}
				});

				// Getting Variance.
				// Mean is already computed.
				// Calculate the sum of the squared differences between each counter value and the mean
				let sumSquaredDiffComments = this.photoListLinks.reduce((acc, obj) => acc + (obj.numberComments - this.meanNumberComments) ** 2, 0);

				// Calculate the variance by dividing the sum of squared differences by the number of elements in the list minus one
				this.varianceNumberComments = sumSquaredDiffComments / (this.photoListLinks.length - 1);

				// Calculate the sum of the squared differences between each counter value and the mean
				let sumSquaredDiffLikes = this.photoListLinks.reduce((acc, obj) => acc + (obj.numberLikes - this.meanNumberLikes) ** 2, 0);

				// Calculate the variance by dividing the sum of squared differences by the number of elements in the list minus one
				this.varianceNumberLikes = sumSquaredDiffLikes / (this.photoListLinks.length - 1);

				// Getting std.
				this.stdNumberComments = Math.sqrt(varianceNumberComments);
				this.stdNumberLikes = Math.sqrt(this.varianceNumberLikes);

				// Getting range.
				this.rangeNumberComments = this.maxNumberComments - this.minNumberComments;
				this.rangeNumberLikes = this.maxNumberLikes - this.minNumberLikes;


			} catch (e) {

				// If an error is encountered, display it!
				if (e.response && e.response.status === 400) {
					this.errormsg = "Request error, please Login before doing some action or ask to get the list of photos of a valid user." + e.toString();
                } else if (e.response && e.response.status === 403) {
                    this.errormsg = "An Unauthorized Action has been blocked. You are not allowed to do this action because you are not the profile's owner." + e.toString();
                } else if (e.response && e.response.status === 204) {
                    this.errormsg = "In the Internal DB there is not anymore the content you have asked." + e.toString();
                } else if (e.response && e.response.status === 500) {
                    this.errormsg = "An internal error occurred. We will be notified. Please try again later." + e.toString();
                } else {
                    this.errormsg = e.toString();
                }		

			}

			// Setting again the Loading flag to false.
			this.loading = false;
        },

	},

	mounted() {
		this.getUserProfile()
		this.getPhotoLinks()
	}
}
</script>



<!-- Actual Page for handling the page setting. -->
<template>

	<div>
			<!-- Let's handle first the upper part that will be the static one. -->
			<h1 class="h1">{{ username }}'s ANALYTICS</h1>
			<!-- {{ this.userProfile }} -->
			<!-- {{ this.photoListLinks }} -->
			<!-- {{ this.totNumberComments }} -->
			<div class="topMenu">

				<!-- "Users List" Button -->
				<div class="topMenuButtons"></div>

				<!-- WASA Photo Icon -->
				<div class="topMenuColumn">
					<img src="./img/wasa-logo.png" alt="" class="img">
				</div>

				<!-- "Search Username Field" -->
				<div class="topMenuButtons"></div>

			</div>

			<div class="divider" style="margin-top: 300px;">
				<span></span><span>Posts</span><span></span>
			</div>
			
			<div class="card" id="div1" style=" background-color:peachpuff; margin-top: 31px;" v-if="!loading">
				<div class="grid-child-posts3" style=" font-size: 15px; margin-left:20px; margin-top: 30px;">
					<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#activity"/></svg>                
					<b> POSTS ANALYTICS </b> <br><br><br> 
					<b> NUMBER OF POSTS </b> {{ this.userProfile.numberOfPhotos }}
					<br><br><br>

					<b> TOTAL NUMBER OF LIKES </b> {{ this.totNumberLikes }}
					<br><br>
					<b> MEAN NUMBER OF LIKES </b> {{ this.numberLikesPerPostMean }}
					<br><br>
					<b> MAX NUMBER OF LIKES </b> {{ this.maxNumberLikes }}
					<br><br>
					<b> MIN NUMBER OF LIKES </b> {{ this.minNumberLikes }}
					<br><br>
					<b> MEDIAN NUMBER OF LIKES </b> {{ this.medianNumberLikes }}
					<br><br>
					<b> MODE NUMBER OF LIKES </b> {{ this.modeNumberLikes }}		
					<br><br>
					<b> VARIANCE NUMBER OF LIKES </b> {{ this.varianceNumberLikes }}			
					<br><br>
					<b> STANDARD DEVIATION NUMBER OF LIKES </b> {{ this.stdNumberLikes }}				
					<br><br>
					<b> RANGE NUMBER OF LIKES </b> {{ this.rangeNumberLikes }}					
					<br><br><br>

					<b> TOTAL NUMBER OF COMMENTS </b> {{ this.totNumberComments }}
					<br><br>
					<b> MEAN NUMBER OF COMMENTS </b> {{ this.numberCommentsPerPostMean }}
					<br><br>
					<b> MAX NUMBER OF COMMENTS </b> {{ this.maxNumberComments }}
					<br><br>
					<b> MIN NUMBER OF COMMENTS </b> {{ this.minNumberComments }}
					<br><br>
					<b> MEDIAN NUMBER OF COMMENTS </b> {{ this.medianNumberComments }}
					<br><br>
					<b> MODE NUMBER OF COMMENTS </b> {{ this.modeNumberComments }}	
					<br><br>
					<b> VARIANCE NUMBER OF COMMENTS </b> {{ this.varianceNumberComments }}		
					<br><br>
					<b> STANDARD DEVIATION NUMBER OF COMMENTS </b> {{ this.stdNumberComments }}				
					<br><br>
					<b> RANGE NUMBER OF COMMENTS </b> {{ this.rangeNumberComments }}		
				</div>
			</div>

    </div><!-- /.container -->
</template>


<!-- Declaration of the style(scoped) to use. -->
<!-- Declaration of the style(scoped) to use. -->
<style scoped>
	.h3 {
	width: 250px;
}
.photos{
	margin-top: 150px;
}
.img {
	display: block;
  	margin-left: auto;
  	margin-right: auto;
  	width: 50%;
    width: 200px;
    height: auto;
}

.card {
	background-color: #ffffff; 
	/* background-color: yellow; */
	margin-top: 300px;
	margin-bottom: 20px;
	height: 83rem;
	width: auto;
	border-radius: 5px;
	align-items: center;
	margin-left: auto;
	flex-direction: column;
	align-items: center;
	box-shadow: rgba(0, 0, 0, 0.7);
	color: black;
  
  }
.topMenu{
	display: block;
  	margin-left: auto;
  	margin-right: auto;
	margin-top: 10px;
	margin-bottom: 10px;
  	width: 50%;
	width: 720px;
}

.h1 {
	display: block;
  	margin-left: auto;
  	margin-right: auto;
	margin-top: 60px;
}

.topMenuColumn {
  float: left;
  width: 33.33%;
}

.topMenuButtons{
	margin-top: 78px;
	margin-left: 10px;
	margin-right: 10px;
	float: left;
  	width: 30%;
}

.usernameLabel{
	float: left;
  	width: 90%;
}
.buttons-menu{
	float: left;
  	width: 10%;
}

.buttonsFollowBan{
	float: left;
  	width: 50%;
}

.feather {
	color: #4a77d4;
}

.searchButton{
	float: left;
	margin-left: 0;
	margin-top: 8px;
}

.result{
	display: block;
  	margin-left: auto;
  	margin-right: auto;
	margin-top: 30px;
}

.btn { display: inline-block; *display: inline; *zoom: 1;   font-family: sans-serif; padding: 4px 10px 4px; margin-bottom: 0; font-size: 13px; line-height: 18px; color: #333333; text-align: center;text-shadow: 0 1px 1px rgba(255, 255, 255, 0.75); vertical-align: middle; background-color: #f5f5f5; background-image: -moz-linear-gradient(top, #ffffff, #e6e6e6); background-image: -ms-linear-gradient(top, #ffffff, #e6e6e6); background-image: -webkit-gradient(linear, 0 0, 0 100%, from(#ffffff), to(#e6e6e6)); background-image: -webkit-linear-gradient(top, #ffffff, #e6e6e6); background-image: -o-linear-gradient(top, #ffffff, #e6e6e6); background-image: linear-gradient(top, #ffffff, #e6e6e6); background-repeat: repeat-x; filter: progid:dximagetransform.microsoft.gradient(startColorstr=#ffffff, endColorstr=#e6e6e6, GradientType=0); border-color: #e6e6e6 #e6e6e6 #e6e6e6; border-color: rgba(0, 0, 0, 0.1) rgba(0, 0, 0, 0.1) rgba(0, 0, 0, 0.25); border: 1px solid #e6e6e6; -webkit-border-radius: 4px; -moz-border-radius: 4px; border-radius: 4px; -webkit-box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.2), 0 1px 2px rgba(0, 0, 0, 0.05); -moz-box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.2), 0 1px 2px rgba(0, 0, 0, 0.05); box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.2), 0 1px 2px rgba(0, 0, 0, 0.05); cursor: pointer; *margin-left: .3em; }
.btn:hover, .btn:active, .btn.active, .btn.disabled, .btn[disabled] { background-color: #e6e6e6; }
.btn-large { padding: 9px 14px; font-size: 20px; line-height: normal; -webkit-border-radius: 5px; -moz-border-radius: 5px; border-radius: 5px; }
.btn:hover { color: #333333; text-decoration: none; background-color: #e6e6e6; background-position: 0 -15px; -webkit-transition: background-position 0.1s linear; -moz-transition: background-position 0.1s linear; -ms-transition: background-position 0.1s linear; -o-transition: background-position 0.1s linear; transition: background-position 0.1s linear; }
.btn-primary, .btn-primary:hover { text-shadow: 0 -1px 0 rgba(0, 0, 0, 0.25); color: #ffffff; }
.btn-primary.active { color: rgba(255, 255, 255, 0.75); }
.btn-primary { background-color: #4a77d4; background-image: -moz-linear-gradient(top, #6eb6de, #4a77d4); background-image: -ms-linear-gradient(top, #6eb6de, #4a77d4); background-image: -webkit-gradient(linear, 0 0, 0 100%, from(#6eb6de), to(#4a77d4)); background-image: -webkit-linear-gradient(top, #6eb6de, #4a77d4); background-image: -o-linear-gradient(top, #6eb6de, #4a77d4); background-image: linear-gradient(top, #6eb6de, #4a77d4); background-repeat: repeat-x; filter: progid:dximagetransform.microsoft.gradient(startColorstr=#6eb6de, endColorstr=#4a77d4, GradientType=0);  border: 1px solid #3762bc; text-shadow: 1px 1px 1px rgba(0,0,0,0.4); box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.2), 0 1px 2px rgba(0, 0, 0, 0.5); }
.btn-primary:hover, .btn-primary:active, .btn-primary.active, .btn-primary.disabled, .btn-primary[disabled] { filter: none; background-color: #4a77d4; }
.btn-block { width: 100%; display:block; }

/* This CSS is dedicated to the input string for the "Username" */
input { 
  width: 100%; 
  margin-bottom: 10px; 
  background: rgb(199, 246, 255); 
  border: none;
  outline: none;
  padding: 10px;
  font-size: 15px;
  font-family: sans-serif;
  border: 1px solid rgba(31, 86, 135, 0.3);
  border-radius: 4px;
  box-shadow: inset 0 -5px 45px rgba(12, 125, 123, 0.2), 0 1px 1px rgba(10, 131, 161, 0.2);
  -webkit-transition: box-shadow .5s ease;
  -moz-transition: box-shadow .5s ease;
  -o-transition: box-shadow .5s ease;
  -ms-transition: box-shadow .5s ease;
  transition: box-shadow .5s ease;
}

.form-control {
	width: 0%;
	float: left;
  	width: 85%;
}

.searchButton{
	width: 0%;
	float: left;
  	width: 15%;
}

* { -webkit-box-sizing:border-box; -moz-box-sizing:border-box; -ms-box-sizing:border-box; -o-box-sizing:border-box; box-sizing:border-box; }

.divider {								/* minor cosmetics */
	display: table; 
	font-size: 24px; 
	text-align: center; 
	width: 75%; 						/* divider width */
	margin: 40px auto;					/* spacing above/below */
}
.divider span { 
	display: table-cell; 
	position: relative; 
}
.divider span:first-child, .divider span:last-child {
	width: 50%;
	top: 13px;							/* adjust vertical align */
	-moz-background-size: 100% 2px; 	/* line width */
	background-size: 100% 2px; 			/* line width */
	background-position: 0 0, 0 100%;
	background-repeat: no-repeat;
}
.divider span:first-child {				/* color changes in here */
	background-image: -webkit-gradient(linear, 0 0, 0 100%, from(transparent), to( rgb(199, 246, 255)));
	background-image: -webkit-linear-gradient(180deg, transparent,  rgb(199, 246, 255));
	background-image: -moz-linear-gradient(180deg, transparent,  rgb(199, 246, 255));
	background-image: -o-linear-gradient(180deg, transparent, rgb(199, 246, 255));
	background-image: linear-gradient(90deg, transparent,  rgb(199, 246, 255));
}
.divider span:nth-child(2) {
	color:  rgb(199, 246, 255); padding: 0px 5px; width: auto; white-space: nowrap;
}
.divider span:last-child {				/* color changes in here */
	background-image: -webkit-gradient(linear, 0 0, 0 100%, from( rgb(199, 246, 255)), to(transparent));
	background-image: -webkit-linear-gradient(180deg,  rgb(199, 246, 255), transparent);
	background-image: -moz-linear-gradient(180deg, rgb(199, 246, 255), transparent);
	background-image: -o-linear-gradient(180deg, rgb(199, 246, 255), transparent);
	background-image: linear-gradient(90deg, rgb(199, 246, 255), transparent);
}
</style>

