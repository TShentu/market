<template>
	<div v-if="label" class="tag-view q-py-xs q-px-md text-caption text-bold">
		{{ label }}
	</div>
</template>

<script lang="ts" setup>
defineProps({
	label: {
		type: String,
		required: false
	}
});
</script>

<style scoped lang="scss">
.tag-view {
	font-family: Roboto;
	height: 20px;
	max-width: 120px;
	white-space: nowrap;
	overflow: hidden;
	text-overflow: ellipsis;
	border-radius: 20px;
	text-transform: capitalize;
	background: $grey-1;
	color: $blue-6;
}
</style>
