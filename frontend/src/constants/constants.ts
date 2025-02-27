export interface AppStoreInfo {
	id: string;
	name: string;
	cfgType: CFG_TYPE;
	chartName: string;
	icon: string;
	desc: string;
	appid: string;
	title: string;
	version: string;
	curVersion: string;
	needUpdate: boolean;
	categories: string[];
	versionName: string;
	fullDescription: string;
	upgradeDescription: string;
	promoteImage: string[];
	promoteVideo: string;
	supportArch: string[];
	subCategory: string;
	developer: string;
	requiredMemory: string;
	requiredDisk: string;
	supportClient: {
		edge: string;
		android: string;
		ios: string;
		windows: string;
		mac: string;
		linux: string;
		chrome: string;
	};
	requiredGpu: string;
	requiredCpu: string;
	rating: number;
	target: string;
	permission: {
		appData: boolean;
		appCache: boolean;
		userData: string[];
		sysData: SysDataCfg[];
	};
	entrances: Entrance[];
	middleware: {
		MongoDB: MiddleWareCfg;
		Redis: MiddleWareCfg;
		Postgres: MiddleWareCfg;
		ZincSearch: MiddleWareCfg;
	};
	options: {
		analytics: {
			enable: boolean;
		};
		dependencies: Dependency[];
		policies: Policy[];
		appScope: {
			clusterScoped: boolean;
			appRef: string[];
		};
		websocket: {
			port: number;
			url: string;
		};
	};
	lastCommitHash: string;
	createTime: number;
	updateTime: number;
	installTime: string;
	uid: string;
	status: APP_STATUS;
	language: string[];
	submitter: string;
	doc: string;
	website: string;
	featuredImage: string;
	sourceCode: string;
	license: License[];
	legal: null;
	appLabels: string[];
	source: string;
	modelSize: string;

	//local
	progress: string;
	preflightError: string[];
}

export interface Entrance {
	authLevel: string;
	host: string;
	icon: string;
	name: string;
	port: 0;
	title: string;
	invisible: boolean;
}

export interface SysDataCfg {
	group: string;
	dataType: string;
	version: string;
	ops: string[];
}

export interface MiddleWareCfg {
	database: any;
	username: string;
	password: string;
}

export interface Dependency {
	name: string;
	type: DEPENDENCIES_TYPE;
	version: string;
}

export interface Policy {
	entranceName: string;
	description: string;
	level: string;
	oneTime: boolean;
	uriRegex: string;
	validDuration: string;
}

export interface License {
	text: string;
	url: string;
}

export interface VersionRecord {
	appName: string;
	mergedAt: string;
	version: string;
	versionName: string;
	upgradeDescription: string;
}

export interface PermissionNode {
	label: string;
	icon?: string;
	children: PermissionNode[];
}

export enum CFG_TYPE {
	APPLICATION = 'app',
	WORK_FLOW = 'recommend',
	MODEL = 'model'
}

export enum SOURCE_TYPE {
	Market = 'market',
	Development = 'custom'
}

export interface Token {
	access_token: string;
	token_type: string;
	refresh_token: string;
	expires_in: number;
	expires_at: number;
}

export interface Resource {
	total: number;
	usage: number;
	ratio: number;
	unit: string;
}

export interface UserResource {
	cpu: Resource;
	memory: Resource;
	disk: Resource;
	gpu: Resource;
}

export interface TerminusResource {
	metrics: {
		cpu: Resource;
		memory: Resource;
		disk: Resource;
		gpu: Resource;
	};
	nodes: string[];
}

export interface TopicInfo {
	name: string;
	introduction: string;
	des: string;
	iconimg: string;
	detailimg: string;
	richtext: string;
	apps: AppStoreInfo[];

	//local index
	id: number;
}

export interface User {
	role: string;
	username: string;
}

export enum ROLE_TYPE {
	Admin = 'platform-admin'
}

export enum TRANSACTION_PAGE {
	App = 'App',
	List = 'List',
	Discover = 'Discover',
	Preview = 'Preview',
	Version = 'Version',
	Log = 'log',
	Update = 'update'
}

export const MENU_TYPE = {
	Application: {
		Home: 'Home',
		SocialNetwork: 'Social Network',
		Entertainment: 'Entertainment',
		Productivity: 'Productivity',
		Utilities: 'Utilities',
		Blockchain: 'Blockchain',
		Recommendation: 'Recommend',
		Models: 'Model'
	},
	MyTerminus: 'MyTerminus'
};

export const CATEGORIES_TYPE = {
	LOCAL: {
		TOP: 'Top',
		LATEST: 'Latest',
		RECOMMENDS: 'Recommends',
		ALL: 'All'
	},
	SERVER: {
		SocialNetwork: 'Social Network',
		Entertainment: 'Entertainment',
		Productivity: 'Productivity',
		Utilities: 'Utilities',
		Blockchain: 'Blockchain',
		News: 'News',
		LifeStyle: 'Lifestyle',
		Sports: 'Sports'
	}
};

export enum DEPENDENCIES_TYPE {
	application = 'application',
	system = 'system'
}

export enum APP_STATUS {
	//server status
	installing = 'installing',
	pending = 'pending',
	running = 'running',
	resuming = 'resuming',
	suspend = 'suspend',
	uninstalling = 'uninstalling',
	upgrading = 'upgrading',
	//model installed maybe not running
	installed = 'installed',

	//backend add status
	uninstalled = 'uninstalled',

	// fronted add status
	installable = 'installable',
	waiting = 'waiting',
	preflightFailed = 'preflightFailed'
}

export enum OPERATE_STATUS {
	pending = 'pending',
	processing = 'processing',
	canceled = 'canceled',
	failed = 'failed',
	completed = 'completed',
	suspend = 'suspend'
}

export enum OPERATE_ACTION {
	pending = 'pending',
	install = 'install',
	uninstall = 'uninstall',
	upgrade = 'upgrade',
	suspend = 'suspend',
	resume = 'resume',
	cancel = 'cancel'
}

export enum CLIENT_TYPE {
	ios = 'ios',
	android = 'android',
	edge = 'edge',
	windows = 'windows',
	mac = 'mac',
	linux = 'linux',
	chrome = 'chrome'
}

export enum PERMISSION_SYSDATA_GROUP {
	service_bfl = 'service.bfl',
	message_dispatcher_system_server = 'message-disptahcer.system-server',
	service_desktop = 'service.desktop',
	service_did = 'service.did',
	api_intent = 'api.intent',
	service_intent = 'service.intent',
	service_message = 'service.message',
	service_notification = 'service.notification',
	service_search = 'service.search',
	secret_infisical = 'secret.infisical',
	secret_vault = 'secret.vault'
}
