package api

import (
	"net/http"
)

// Handler that returns an instance of httprouter.Router that handle APIs registered here.
func (rt *_router) Handler() http.Handler {

	// -----
	// MANDATORY ROUTES:
	// -----

	// USER-RELATED
	// Register the doLogin API.
	rt.router.POST("/session/", rt.wrap(rt.doLogin))

	// Register the getUserProfile API.
	rt.router.GET("/users/:username", rt.wrap(rt.getUserProfile))

	// Register the setMyUserName API.
	rt.router.PATCH("/users/:username", rt.wrap(rt.setMyUserName))

	// FOLLOW-RELATED
	// Register the followUser API.
	rt.router.PUT("/users/:username/followings/:usernameFollowing", rt.wrap(rt.followUser))

	// Register the unfollowUser API.
	rt.router.DELETE("/users/:username/followings/:usernameFollowing", rt.wrap(rt.unfollowUser))

	// BAN-RELATED
	// Register the banUser API.
	rt.router.PUT("/users/:username/bans/:usernameBanned", rt.wrap(rt.banUser))

	// Register the unbanUser API.
	rt.router.DELETE("/users/:username/bans/:usernameBanned", rt.wrap(rt.unbanUser))

	// PHOTO-RELATED
	//Register the uploadPhoto API.
	rt.router.POST("/users/:username/photos/", rt.wrap(rt.uploadPhoto))

	// Register the deletePhoto API.
	rt.router.DELETE("/users/:username/photos/:photoid", rt.wrap(rt.deletePhoto))

	// Register the likePhoto API.
	rt.router.PUT("/users/:username/photos/:photoid/likes/:usernameLiker", rt.wrap(rt.likePhoto))

	// Register the unlikePhoto API.
	rt.router.DELETE("/users/:username/photos/:photoid/likes/:usernameLiker", rt.wrap(rt.unlikePhoto))

	// Register the commentPhoto API.
	rt.router.POST("/users/:username/photos/:photoid/comments/", rt.wrap(rt.commentPhoto))

	// Register the uncommentPhoto API.
	rt.router.DELETE("/users/:username/photos/:photoid/comments/:commentid", rt.wrap(rt.uncommentPhoto))

	// Register the getMyStream API.
	rt.router.GET("/users/:username/myStream/", rt.wrap(rt.getMyStream))

	// -----
	// OPTIONAL ROUTES
	// -----

	// Register the getUsers API.
	rt.router.GET("/users/", rt.wrap(rt.GetUsers))

	// Register the deleteUsername API.
	rt.router.DELETE("/users/:username", rt.wrap(rt.deleteUser))

	// Register the setUser API.
	rt.router.PUT("/users/:username", rt.wrap(rt.setUser))

	// Register the getPhotos API.
	rt.router.GET("/users/:username/photos/", rt.wrap(rt.getPhotos))

	// Register the getPhoto API.
	rt.router.GET("/users/:username/photos/:photoid", rt.wrap(rt.getPhoto))

	// Register the getPhotoComments API.
	// rt.router.GET("/users/:username/photos/:photoid/comments", rt.wrap(rt.getPhotoComments))

	// Register the getPhotoLikes API.
	// rt.router.GET("/users/:username/photos/:photoid/likes/", rt.wrap(rt.getPhotoLikes))

	// Register the getBannedUsers API.
	rt.router.GET("/users/:username/bans/", rt.wrap(rt.getBannedUsers))

	// Register the getFollowers API.
	rt.router.GET("/users/:username/followers/", rt.wrap(rt.getFollowers))

	// Register the getFollowings API.
	rt.router.GET("/users/:username/followings/", rt.wrap(rt.getFollowings))

	// -----
	// SPECIAL ROUTES
	// -----

	// Register the liveness API.
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
