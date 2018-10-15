import Buefy from 'buefy';
import Vue from 'vue';
import VueSSE from 'vue-sse';

import App from './App.vue';

Vue.use(Buefy);
Vue.use(VueSSE);

Vue.config.productionTip = false;

new Vue({
  render: h => h(App),
}).$mount('#app');
