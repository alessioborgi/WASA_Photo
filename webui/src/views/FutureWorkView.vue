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
            username: localStorage.getItem('Username'),
            BearerToken: localStorage.getItem('BearerToken'),
		}
	},

	// Declaration of the methods that will be used.
	methods: {
		
        

	},

}
</script>

<!-- Actual Page for handling the page setting. -->
<template>

	<div>
			<!-- Let's handle first the upper part that will be the static one. -->
			<h1 class="h1">FUTURE OF WORK</h1>

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

			<!-- If instead, it is all ok, Display a sort of card for each of the User Profiles(Depending on we are asking the whole list or just one). -->
			<div class="card" id="div1" style=" background-color:lightsalmon">
				<div class="grid-child-posts3" style=" font-size: 15px; margin-left:20px; margin-top: 30px;">
					<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#pen-tool"/></svg>                
					<b> HUMAN COMPUTER INTERACTION FOCUS (FEEDBACKS): </b> <br><br><br> 
					<br> This section is dedicated to a <b>HCI (Human Computer Interaction)</b> focus. Indeed, I decided to link this project also to some of the concepts I have studied in this other course. <br><br>
					I decided to go over the Project and to think at it as if it was a startup-project I was launching. For this reason, at the end of the prototyping part, I decided to ask to some <b>Pilots</b>, that in my case were my parents and a couple of friends to describe their experience with the Web App and to give the more feedbacks and advices they can. 
					<br><br>
					The Query techniques I have choosen are therefore very informal, but I gathered lots of advices that has helped me out to improve the User experience in the app. Indeed, even if the users tested (Pilots) were people I was very confident with, I decided to not give any input on the app and I asked to them to be the more transparent they could.
					Of course, this is not a complete collection, but it helped me out to solve some problems I was not aware of. 
					 
				</div>
			</div>
			
			<div class="divider">
				<span></span><span>FEEDBACKS & ADVICES</span><span></span>
			</div>

			<div class="card" id="div1" style=" background-color:peachpuff; margin-top: 30px;">
				<div class="grid-child-posts3" style=" font-size: 15px; margin-left:20px; margin-top: 30px;">
					<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#pen-tool"/></svg>                
					<b> PILOT andrea71 </b> <br><br><br> 
					<b> FEEDBACK & ADVICES </b>
					<br>"It is a very good implementation. If I had to suggest an Improvement to do, I would advise you to put the link to every actions the user can do also directly in the MyProfile page. Sometimes, providing redundancy in your application is not a bad thing, as in this case. I think that is very important for the user to be in its Profile Page and it can do all from the Card you have for your Profile. You could put two menus, and one button for the Stream."
					<br><br>
					<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#check"/></svg>   
					<br><br>
					<b> ACTIONS & CHANGES AFTER THIS FEEDBACK </b>
					<br>I proceeded in creating a new menu Item containing all the remaining main actions, in such a way that the user can do whatever action from its Profile Page.
				</div>
			</div>

			<div class="card" id="div1" style=" background-color:peachpuff; margin-top: 31px;">
				<div class="grid-child-posts3" style=" font-size: 15px; margin-left:20px; margin-top: 30px;">
					<svg class="feather" ><use href="/feather-sprite-v4.29.0.svg#pen-tool"/></svg>                
					<b> PILOT anna69 </b> <br><br><br> 
					<b> FEEDBACK & ADVICES </b>
					<br>"It's very similar to the Instagram App! As further implementation, what about considering the possibility to change the photo's caption? In my case, it happen that I post a photo and that I would like to change also the caption, also because maybe I misspell a word. In your implementation instead, you have to eliminate and reinsert the post."
					<br><br>
					<svg class="feather" ><use href="/feather-sprite-v4.29.0.svg#loader"/></svg>                
					<br><br>
					<b> ACTIONS & CHANGES AFTER THIS FEEDBACK </b>
					<br>
				</div>
			</div>

			<div class="card" id="div1" style=" background-color:peachpuff; margin-top: 31px;">
				<div class="grid-child-posts3" style=" font-size: 15px; margin-left:20px; margin-top: 30px;">
					<svg class="feather" ><use href="/feather-sprite-v4.29.0.svg#pen-tool"/></svg>                
					<b> PILOT alice05 </b> <br><br><br> 
					<b> FEEDBACK & ADVICES </b>
					<br>"It's very cool! One thing I don't like so much is the color you set by-default depending on the sex. That could be improved by adding up the possibility to change the color to the user, in such a way he/she can put whatever it wants. It would be also a feature Instagram doesn't have, the customization. This could be a reason in plus, to download your app instead of Instagram!"
					<br><br>
					<svg class="feather" ><use href="/feather-sprite-v4.29.0.svg#loader"/></svg>                
					<br><br>
					<b> ACTIONS & CHANGES AFTER THIS FEEDBACK </b>
					<br>I proceeded to create a "do not specify" option in such a way to also provide the possibility to choice to not define the sex. I put it as default. 
				</div>
			</div>

			<div class="card" id="div1" style=" background-color:peachpuff; margin-top: 31px;">
				<div class="grid-child-posts3" style=" font-size: 15px; margin-left:20px; margin-top: 30px;">
					<svg class="feather" ><use href="/feather-sprite-v4.29.0.svg#pen-tool"/></svg>                
					<b> PILOT samantha01 </b>  <br><br><br> 
					<b> FEEDBACK & ADVICES </b>
					<br>"That's a very good achievement. What I would prefer is the fact that in the Follow Dashboard, if you have asked for the followings list and you decide to unfollow one of them, you get immediately deleted from the list the user. It could happen tha you press the button wrongly, and for re-following the user, you should necessary do the Search of the user again and press the following button again. The same if for the Ban Dashboard."
					<br><br>
					<svg class="feather" ><use href="/feather-sprite-v4.29.0.svg#check"/></svg>                
					<br><br>
					<b> ACTIONS & CHANGES AFTER THIS FEEDBACK </b>
					<br>I proceeded in eliminating the immediate elimination of the user card from the ban and the follow dashboard. This was a thing I was not considering and, now, we can eliminate from the list only if we reload the List.
				</div>
			</div>

			<div class="card" id="div1" style=" background-color:peachpuff; margin-top: 31px;">
				<div class="grid-child-posts3" style=" font-size: 15px; margin-left:20px; margin-top: 30px;">
					<svg class="feather" ><use href="/feather-sprite-v4.29.0.svg#pen-tool"/></svg>                
					<b> PILOT chicca78 </b> <br><br><br> 
					<b> FEEDBACK & ADVICES </b>
					<br>"Very good! One thing I would change is the fact that when you have to update the profile, I may want to not change necessarily also the username. Therefore, eliminate this constraint! Indeed, in most of the cases I modify the profile, I never change the username, or at least in my case. I havent's changed it never."
					<br><br>
					<svg class="feather" ><use href="/feather-sprite-v4.29.0.svg#check"/></svg>                
					<br><br>
					<b> ACTIONS & CHANGES AFTER THIS FEEDBACK </b>
					<br>I proceeded in eliminating this constraint. I relised this only when I arrived at the front-end part. This was a dumb constraint because in most of the cases we may want to change only other things and also because if we wanted to chang the username, we go for the setUsername change.
				</div>
			</div>

			<div class="card" id="div1" style=" background-color:peachpuff; margin-top: 31px;">
				<div class="grid-child-posts3" style=" font-size: 15px; margin-left:20px; margin-top: 30px;">
					<svg class="feather" ><use href="/feather-sprite-v4.29.0.svg#pen-tool"/></svg>                
					<b> PILOT aurora46 </b> <br><br><br> 
					<b> FEEDBACK & ADVICES </b>
					<br>"That's impressive. Couldn't you add up a page where you explain how to do the things? For example, for every type of action, try to put a set of steps that you have to follow, in such a way to achieve the task. This would help also people like me that are less technological :) What you could do is also one sort of animation that explain you how to use the app at the very start, like in almost all new apps nowadays."
					<br><br>
					<svg class="feather" ><use href="/feather-sprite-v4.29.0.svg#loader"/></svg>                
					<br><br>
					<b> ACTIONS & CHANGES AFTER THIS FEEDBACK </b> 
				</div>
			</div>

			<div class="card" id="div1" style=" background-color:peachpuff; margin-top: 31px;">
				<div class="grid-child-posts3" style=" font-size: 15px; margin-left:20px; margin-top: 30px;">
					<svg class="feather" ><use href="/feather-sprite-v4.29.0.svg#pen-tool"/></svg>                
					<b> PILOT federico01 </b> <br><br><br> 
					<b> FEEDBACK & ADVICES </b>
					<br>"Very good! One thing I would change is the fact that the change of the user Profile information and the Photo Profile could be separate. I don't know whether is is difficult, but I would do it in two different pages. Indeed, if I change one, this doesn't mean I want also to change the other. Changing photo profile is also more often done wrt the change of profile information."
					<br><br>
					<svg class="feather" ><use href="/feather-sprite-v4.29.0.svg#loader"/></svg>                
					<br><br>
					<b> ACTIONS & CHANGES AFTER THIS FEEDBACK </b>
					<br>
				</div>
			</div>
			
		</div>

</template>

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
	height: 31rem;
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
