<template>
	<div
		class="header-bar row justify-start items-center"
		v-if="show"
		:style="[{ boxShadow: shadow ? '0px 2px 4px 0px #0000001A' : 'none' }]"
	>
		<q-icon
			class="application_bar_return cursor-pointer"
			v-if="showBack"
			name="sym_r_arrow_back_ios_new"
			size="20px"
			@click="onReturn"
		/>
		<transition name="fade">
			<div
				v-if="showTitle"
				class="header-title text-subtitle2 text-color-title"
			>
				{{ title }}
			</div>
		</transition>
	</div>
</template>

<script lang="ts" setup>
import { useAppStore } from 'src/stores/app';

defineProps({
	show: {
		type: Boolean,
		require: true,
		default: false
	},
	showBack: {
		type: Boolean,
		default: true
	},
	showTitle: {
		type: Boolean,
		default: true
	},
	shadow: {
		type: Boolean,
		default: false
	},
	title: {
		type: String,
		require: true,
		default: ''
	}
});

const emit = defineEmits(['onReturn']);
const appStore = useAppStore();
const onReturn = () => {
	emit('onReturn');
};
</script>

<style scoped lang="scss">
.header-bar {
	width: 100%;
	height: 56px;
	background: transparent;
	text-align: center;
	color: #0a0b27;
	z-index: 9999;

	.application_bar_return {
		margin-left: 18px;
	}

	.header-title {
		margin-left: 6px;
	}
}
</style>
