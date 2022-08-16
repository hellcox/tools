import 'core-js/stable';
import 'regenerator-runtime/runtime';
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import Vue from 'vue';
import App from './App.vue';

Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from '@wailsapp/runtime';

Wails.Init(() => {
	Vue.use(ElementUI);
	new Vue({
		render: h => h(App)
	}).$mount('#app');
});
