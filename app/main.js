import Vue from 'vue'
import store from './store'
import App from './App'

import Json from './Json' //测试用数据
import axios from "axios"


const msg = (title, duration=1500, mask=false, icon='none')=>{
	//统一提示方便全局修改
	if(Boolean(title) === false){
		return;
	}
	uni.showToast({
		title,
		duration,
		mask,
		icon
	});
}
const json = type=>{
	//模拟异步请求数据
	return new Promise(resolve=>{
		setTimeout(()=>{
			resolve(Json[type]);
		}, 500)
	})
}

const prePage = ()=>{
	let pages = getCurrentPages();
	let prePage = pages[pages.length - 2];
	// #ifdef H5
	return prePage;
	// #endif
	return prePage.$vm;
}


Vue.config.productionTip = false
Vue.prototype.$fire = new Vue();
Vue.prototype.$store = store;
Vue.prototype.$api = {msg, json, prePage};


axios.defaults.baseURL = "http://api.huigoumall.cn/"; 
axios.defaults.headers.post["Content-Type"] = "application/json";
Vue.config.productionTip = false;
Vue.prototype.$http = axios;


App.mpType = 'app'

const app = new Vue({
    ...App
})
app.$mount()