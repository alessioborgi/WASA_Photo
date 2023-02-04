<!-- This component will be used to handle the different Errors we encounter during the execution of the WebApp. -->
<!-- 
<script>

// Declaration of the Export set.
export default {

	// Declaration of the props needed.
	props: ['msg', 'details']
}

</script>

<template>
	<div class="alert alert-danger" role="alert">
		{{ msg }}
		<small v-if="details"><br /><br />Details: {{ details }}</small>
	</div>
</template>

<style>
</style> -->




<script>
export default {
	props: ['msg']
}
</script>

<template>
	<div class="alert alert-danger" role="alert">
		{{ msg }}
	</div>
</template>

<style>
</style>
