/**
 * Created by xyp on 17-7-16.
 */
import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const Home = {template: '<div>home</div>'}
const Foo = {template: '<div>foo</div>'}
const Bar2 = {
  template: `
    <div>
      bar2
      <div style="height:500px"></div>
      <p id="bar2anchor" style="height:1000px">Bar2Anchor</p>
      <p id="bar2anchor2" style="height:1000px">Bar2Anchor2</p>
    </div>
  `
}
const Bar = {
  template: `
    <div>
      bar
      <div style="height:500px"></div>
      <p id="anchor" style="height:500px">Anchor</p>
      <p id="anchor2">Anchor2</p>
      <router-view></router-view>
    </div>
  `
}

// scrollBehavior:
// - only available in html5 history mode
// - defaults to no scroll behavior
// - return false to prevent scroll
const scrollBehavior = (to, from, savedPosition) => {
  if (savedPosition) {
    // savedPosition is only available for popstate navigations.
    return savedPosition
  } else {
    const position = {}
    // new navigation.
    // scroll to anchor by returning the selector
    if (to.hash) {
      position.selector = to.hash
    }
    // check if any matched route config has meta that requires scrolling to top
    if (to.matched.some(m => m.meta.scrollToTop)) {
      // cords will be used if no selector is provided,
      // or if the selector didn't match any element.
      position.x = 0
      position.y = 0
    }
    // if the returned position is falsy or an empty object,
    // will retain current scroll position.
    return position
  }
}

const router = new VueRouter({
  mode: 'history',
  base: __dirname,
  scrollBehavior,
  routes: [
    {path: '/', component: Home, meta: {scrollToTop: true}},
    {path: '/foo', component: Foo},
    {
      path: '/bar',
      component: Bar,
      meta: {scrollToTop: true},
      children: [{
        path: 'bar2',
        component: Bar2
      }
      ]
    }
  ]
})

new Vue({
  router,
  template: `
    <div id="app">
      <h1>Scroll Behavior</h1>
      <ul>
        <li><router-link to="/">/</router-link></li>
        <li><router-link to="/foo">/foo</router-link></li>
        <li><router-link to="/bar">/bar</router-link></li>
        <li><router-link to="/bar#anchor">/bar#anchor</router-link></li>
        <li><router-link to="/bar#anchor2">/bar#anchor2</router-link></li>
        <li><router-link to="#bar2anchor">/bar/bar2#anchor</router-link></li>
        <li><router-link to="#bar2anchor2">/bar/bar2#anchor2</router-link></li>
        <li><router-link to="/bar/bar2">/bar/bar2</router-link></li>
      </ul>
      <router-view class="view"></router-view>
    </div>
  `
}).$mount('#app')
