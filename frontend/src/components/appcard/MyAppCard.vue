<template>
	<div v-if="skeleton" class="my-application-root column items-center">
		<app-featured-image :skeleton="true" />
		<div class="my-application-info column justify-start items-start">
			<div class="my-application-title-layout row justify-start items-center">
				<q-skeleton width="80px" height="24px" />
				<q-skeleton
					v-if="version"
					class="my-application-version"
					width="40px"
					height="20px"
				/>
			</div>

			<q-skeleton width="200px" height="32px" style="margin-top: 2px" />

			<div
				class="my-application-bottom-layout row justify-between items-center"
			>
				<q-skeleton width="88px" height="32px" />
			</div>
		</div>
	</div>
	<div
		v-else
		class="my-application-root column items-center"
		@click="goAppDetails"
		@mouseover="hoverRef = true"
		@mouseleave="hoverRef = false"
	>
		<app-featured-image :src="item.featuredImage">
			<app-icon
				:src="item.icon"
				:skeleton="false"
				:size="108"
				:cs-app="clusterScopedApp"
			/>
		</app-featured-image>
		<div class="my-application-info column justify-start items-start">
			<div class="my-application-title-layout row justify-start items-center">
				<div class="my-application-title">{{ item.title }}</div>
				<div v-if="version" class="my-application-version">
					{{ item.versionName }}
				</div>
			</div>

			<div class="my-application-content">
				{{ item.desc }}
			</div>

			<div
				class="my-application-bottom-layout row justify-between items-center"
			>
				<install-button
					v-if="item && !appStore.isPublic"
					style="width: auto"
					:development="development"
					:item="item"
					:larger="true"
					:is-update="isUpdate"
					:manager="manager"
				/>
				<div
					v-if="tag === CFG_TYPE.WORK_FLOW"
					class="row justify-end items-center"
				>
					<app-tag
						:label="
							item.categories && item.categories.length > 0
								? item.categories[0]
								: item.subCategory
						"
						class="text-green-5"
					/>
					<app-tag
						:label="
							item.language && item.language.length > 0
								? convertLanguageCodeToName(item.language[0])
								: ''
						"
						class="second-tag q-ml-sm"
					/>
				</div>
				<div
					v-else-if="tag === CFG_TYPE.MODEL"
					class="row justify-end items-center"
				>
					<app-tag
						:label="
							item.categories && item.categories.length > 0
								? item.categories[0]
								: item.subCategory
						"
						class="text-green-5"
					/>
					<app-tag :label="item.modelSize" class="second-tag q-ml-sm" />
				</div>
				<div
					v-else
					class="row justify-end items-center"
					style="max-width: calc(100% - 120px)"
				>
					<app-tag
						:label="
							item.cfgType === CFG_TYPE.WORK_FLOW
								? t('main.recommend')
								: item.cfgType
						"
						:class="
							item.cfgType === CFG_TYPE.WORK_FLOW
								? 'text-orange-6'
								: item.cfgType === CFG_TYPE.MODEL
									? 'text-red-6'
									: ''
						"
					/>
					<app-tag
						:label="
							item.categories && item.categories.length > 0
								? item.categories[0]
								: item.subCategory
						"
						class="second-tag text-green-5 q-ml-sm"
					/>
				</div>
			</div>
		</div>
	</div>
</template>

<script lang="ts" setup>
import { onMounted, PropType, ref } from 'vue';
import {
	AppStoreInfo,
	CFG_TYPE,
	SOURCE_TYPE,
	TRANSACTION_PAGE
} from 'src/constants/constants';
import InstallButton from 'src/components/appcard/InstallButton.vue';
import { useRouter } from 'vue-router';
import { useAppStore } from 'src/stores/app';
import AppFeaturedImage from 'components/appcard/AppFeaturedImage.vue';
import AppIcon from 'components/appcard/AppIcon.vue';
import AppTag from 'src/components/appcard/AppTag.vue';
import { useI18n } from 'vue-i18n';
import { convertLanguageCodeToName } from 'src/utils/utils';

const router = useRouter();
const appStore = useAppStore();
const hoverRef = ref(false);
const { t } = useI18n();

const props = defineProps({
	item: {
		type: Object as PropType<AppStoreInfo>,
		required: false
	},
	development: {
		type: Boolean,
		required: false
	},
	version: {
		type: Boolean,
		default: false
	},
	disabled: {
		type: Boolean,
		default: false
	},
	isUpdate: {
		type: Boolean,
		required: false,
		default: false
	},
	manager: {
		type: Boolean,
		required: false,
		default: false
	},
	skeleton: {
		type: Boolean,
		default: false
	},
	tag: {
		type: String,
		default: ''
	}
});

const clusterScopedApp = ref(false);

onMounted(() => {
	if (props.item && props.item.options && props.item.options.appScope) {
		clusterScopedApp.value = props.item.options.appScope.clusterScoped;
	} else {
		clusterScopedApp.value = false;
	}
});

function goAppDetails() {
	if (props.disabled) {
		return;
	}
	if (!props.item) {
		console.log('go app details failure');
		return;
	}

	if (props.item.source === SOURCE_TYPE.Development) {
		console.log('go app details refuse');
		return;
	}
	appStore.setAppItem(props.item);
	router.push({
		name: TRANSACTION_PAGE.App,
		params: {
			name: props.item.name,
			type: props.item.cfgType
		}
	});
}
</script>
<style lang="scss" scoped>
.my-application-root {
	width: 100%;
	height: auto;
	overflow: hidden;

	.my-application-info {
		width: 100%;
		height: 112px;

		.my-application-title-layout {
			width: 100%;
			text-overflow: ellipsis;
			white-space: nowrap;
			overflow: hidden;
			margin-top: 12px;

			.my-application-title {
				font-family: Roboto;
				font-size: 16px;
				font-weight: 700;
				line-height: 24px;
				letter-spacing: 0;
				text-align: left;
				color: $title;
				max-width: 70%;
			}

			.my-application-version {
				font-family: Roboto;
				font-size: 14px;
				font-weight: 700;
				line-height: 20px;
				letter-spacing: 0;
				text-align: left;
				color: $blue;
				max-width: 25%;
				margin-left: 8px;
			}
		}

		.my-application-content {
			font-family: Roboto;
			font-size: 12px;
			font-weight: 400;
			height: 32px;
			line-height: 16px;
			letter-spacing: 0;
			text-align: left;
			overflow: hidden;
			text-overflow: ellipsis;
			display: -webkit-box;
			-webkit-line-clamp: 2;
			-webkit-box-orient: vertical;
			color: var(--Grey-05, #adadad);
		}

		.my-application-bottom-layout {
			width: 100%;
			height: 32px;
			margin-top: 10px;

			@media (max-width: 1200px) and (min-width: 1024px) {
				.second-tag {
					display: none;
				}
			}
		}
	}
}
</style>
