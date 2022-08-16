import {createRouter, createWebHashHistory} from 'vue-router'
// import hello from './../components/HelloWorld'

const routes = [
    {path: '/', component: "<h1>hello</h1>>"},
]

const router = createRouter({
    history: createWebHashHistory("/"),
    routes,
})


export default router