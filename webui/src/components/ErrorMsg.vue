<!-- This component will be used to handle the different Errors we encounter during the execution of the WebApp. -->

<!-- <script>
export default {
	props: ['msg']
}
</script>

<template>
	<div class="alert alert-danger" role="alert" >
		{{ msg }}
	</div>
</template>

<style>

.alert-danger{
	font-size: 25px;
}
</style>
 -->

<script>
 export default {
	props: ['msg']
}
</script>

<template>

	<head>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
	</head>
	
	
	<div class="alert alert-danger alert-white rounded">
    	<button type="button" class="close" data-dismiss="alert" aria-hidden="true">×</button>
    	<div class="icon"><i class="fa fa-times-circle"></i></div>
    	<strong>Error!</strong> {{ msg }}.
	</div>
	
</template>

<style>
@import url(https://fonts.googleapis.com/css?family=Open+Sans:400,700);

body {
    font-family:'Open Sans',sans-serif;
    line-height:normal;
    background:#e9f0f4;
    padding:0;
    margin:0;   
}

.content {
    padding:0;
    margin:10% 15%;   
}

.close {
    float: right;
    font-size: 21px;
    font-weight: bold;
    line-height: 1;
    color: #000;
    text-shadow: 0 1px 0 #fff;
    opacity: .2;
}

.close:hover,.close:focus {
    color: #000;
    text-decoration: none;
    cursor: pointer;
    opacity: .5;
}

button.close {
    padding: 0;
    cursor: pointer;
    background: transparent;
    border: 0;
}

.alert {
    padding: 15px;
    margin-bottom: 20px;
    border: 1px solid transparent;
    border-radius: 4px;
}

.alert h4 {
    margin-top: 0;
    color: inherit;
}

.alert .alert-link {
    font-weight: bold;
}

.alert>p,.alert>ul {
    margin-bottom: 0;
}

.alert>p+p {
    margin-top: 5px;
}

.alert-dismissable {
    padding-right: 35px;
}

.alert-dismissable .close {
    position: relative;
    top: -2px;
    right: -21px;
    color: inherit;
}

.alert-success {
    background-color: #dff0d8;
    border-color: #d6e9c6;
    color: #3c763d;
}

.alert-success hr {
    border-top-color: #c9e2b3;
}

.alert-success .alert-link {
    color: #2b542c;
}

.alert-info {
    background-color: #d9edf7;
    border-color: #bce8f1;
    color: #31708f;
}

.alert-info hr {
    border-top-color: #a6e1ec;
}

.alert-info .alert-link {
    color: #245269;
}

.alert-warning {
    background-color: #fcf8e3;
    border-color: #faebcc;
    color: #8a6d3b;
}

.alert-warning hr {
    border-top-color: #f7e1b5;
}

.alert-warning .alert-link {
    color: #66512c;
}

.alert-danger {
    background-color: #f2dede;
    border-color: #ebccd1;
    color: #a94442;
	font-size: 20px;
}

.alert-danger hr {
    border-top-color: #e4b9c0;
}

.alert-danger .alert-link {
    color: #843534;
}

.alert {
    border-radius: 0;
    -webkit-border-radius: 0;
    box-shadow: 0 1px 2px rgba(0,0,0,0.11);
}

.alert .sign {
    font-size: 20px;
    vertical-align: middle;
    margin-right: 5px;
    text-align: center;
    width: 25px;
    display: inline-block;
}

.alert-success {
    background-color: #dbf6d3;
    border-color: #aed4a5;
    color: #569745;
}

.alert-info {
    background-color: #d9edf7;
    border-color: #98cce6;
    color: #3a87ad;
}

.alert-warning {
    background-color: #fcf8e3;
    border-color: #f1daab;
    color: #c09853;
}

.alert-danger {
    background-color: #f2dede;
    border-color: #e0b1b8;
    color: #b94a48;
}

.alert-white {
    background-image: linear-gradient(to bottom,#FFFFFF,#F9F9F9);
    border-top-color: #d8d8d8;
    border-bottom-color: #bdbdbd;
    border-left-color: #cacaca;
    border-right-color: #cacaca;
    color: #404040;
    padding-left: 61px;
    position: relative;
}

.alert-white .icon {
    text-align: center;
    width: 45px;
    height: 100%;
    position: absolute;
    top: -1px;
    left: -1px;
    border: 1px solid #bdbdbd;
}

.alert-white .icon:after {
    display: block;
    content: '';
    width: 10px;
    height: 10px;
    border: 1px solid #bdbdbd;
    position: absolute;
    border-left: 0;
    border-bottom: 0;
    top: 50%;
    right: -6px;
    margin-top: -5px;
    background: #fff;
}

.alert-white.rounded {
    border-radius: 3px;
    -webkit-border-radius: 3px;
}

.alert-white.rounded .icon {
    border-radius: 3px 0 0 3px;
    -webkit-border-radius: 3px 0 0 3px;
}

.alert-white .icon i {
    font-size: 20px;
    color: #FFF;
    left: 12px;
    margin-top: -10px;
    position: absolute;
    top: 50%;
}

.alert-white.alert-danger .icon,.alert-white.alert-danger .icon:after {
    border-color: #ca452e;
    background: #da4932;
}

.alert-white.alert-info .icon,.alert-white.alert-info .icon:after {
    border-color: #3a8ace;
    background: #4d90fd;
}

.alert-white.alert-warning .icon,.alert-white.alert-warning .icon:after {
    border-color: #d68000;
    background: #fc9700;
}

.alert-white.alert-success .icon,.alert-white.alert-success .icon:after {
    border-color: #54a754;
    background: #60c060;
}
</style>

