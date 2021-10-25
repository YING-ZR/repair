import { createRouter, createWebHashHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Login from '../views/Login.vue'
import House_split from '../views/House_split'
import Select from '../views/Select.vue'
import Home_Consume from '../views/Home_Consume.vue'
import Home_diaodu from '../views/Home_diaodu.vue'
import Home_kuguan from '../views/Home_kuguan.vue'

const routes = [
    // {
    //     path: '/Home_Consume',
    //     name: 'Home_Consume',
    //     component: Home_Consume,
    //     meta: {
    //         isShow: false,
    //     },
    //     children: [
    //         // {
    //         //     meta: {
    //         //         isShow: true,
    //         //         title: '分房申请'
    //         //     },
    //         //     path: '/House_split',
    //         //     name: 'House_split',
    //         //     component: House_split
    //         //         // component: () =>
    //         //         //     import ( /* webpackChunkName: "login" */ '../views/House_split.vue'), //路由的软加载

    //         // },
    //         {
    //             path: '/House_query',
    //             name: 'House_query',
    //             component: () =>
    //                 import ( /* webpackChunkName: "login" */ '../views/House_query.vue'), //路由的软加载
    //             meta: {
    //                 isShow: true,
    //                 title: '查看修改个人信息'
    //             }
    //         },
    //         {
    //             path: '/House_change',
    //             name: 'House_change',
    //             component: () =>
    //                 import ( /* webpackChunkName: "login" */ '../views/House_change.vue'), //路由的软加载
    //             meta: {
    //                 isShow: true,
    //                 title: '填写维修信息'
    //             }
    //         },
    //         {
    //             path: '/Housing',
    //             name: '/Housing',
    //             component: () =>
    //                 import ( /* webpackChunkName: "login" */ '../views/Housing.vue'), //路由的软加载
    //             meta: {
    //                 isShow: true,
    //                 title: '查看维修单'
    //             }
    //         },
    //         {
    //             path: '/State',
    //             name: '/State',
    //             component: () =>
    //                 import ( /* webpackChunkName: "login" */ '../views/State.vue'), //路由的软加载
    //             meta: {
    //                 isShow: true,
    //                 title: '订单状态'
    //             }
    //         },
    //         // {
    //         //     path: '/House_out',
    //         //     name: 'House_out',
    //         //     component: () =>
    //         //         import ( /* webpackChunkName: "login" */ '../views/House_out.vue'), //路由的软加载
    //         //     meta: {
    //         //         isShow: true,
    //         //         title: '退房申请'
    //         //     }
    //         // },


    //     ]
    // },
    {
        path: '/Login',
        name: 'Login',
        component: Login,
        meta: {
            isShow: false,
        }
        // component: () => import(/* webpackChunkName: "login" */ '..views/login.vue')//路由的软加载
    },
    {
        path: '/',
        name: 'Select',
        component: Select,
        meta: {
            isShow: false,
        },
        // component: () => import(/* webpackChunkName: "login" */ '..views/YonghuLogin.vue')//路由的软加载
    },
    {
        path: '/Shenfen',
        name: 'Shenfen',
        // component: YonghuLogin,
        meta: {
            isShow: false,
        },
        component: () =>
            import ('../views/Shenfen.vue')
            // component: () => import(/* webpackChunkName: "login" */ '..views/login.vue')//路由的软加载
    },
    {
        path: '/YonghuLogin',
        name: 'YonghuLogin',
        // component: YonghuLogin,
        meta: {
            isShow: false,
        },
        component: () =>
            import ('../views/YonghuLogin.vue')
            // component: () => import(/* webpackChunkName: "login" */ '..views/login.vue')//路由的软加载
    },
    {
        path: '/GuanliyuanLogin',
        name: 'GuanliyuanLogin',
        // component: YonghuLogin,
        meta: {
            isShow: false,
        },
        component: () =>
            import ('../views/GuanliyuanLogin.vue')
            // component: () => import(/* webpackChunkName: "login" */ '..views/login.vue')//路由的软加载
    },
    {
        path: '/Kefu_login',
        name: 'Kefu_login',
        // component: YonghuLogin,
        meta: {
            isShow: false,
        },
        component: () =>
            import ('../views/Kefu_login.vue')
            // component: () => import(/* webpackChunkName: "login" */ '..views/login.vue')//路由的软加载
    },
    {
        path: '/Caiwu_login',
        name: 'Caiwu_login',
        // component: YonghuLogin,
        meta: {
            isShow: false,
        },
        component: () =>
            import ('../views/Caiwu_login.vue')
            // component: () => import(/* webpackChunkName: "login" */ '..views/login.vue')//路由的软加载
    },
    {
        path: '/Kuguan_login',
        name: 'Kuguan_login',
        // component: YonghuLogin,
        meta: {
            isShow: false,
        },
        component: () =>
            import ('../views/Kuguan_login.vue')
            // component: () => import(/* webpackChunkName: "login" */ '..views/login.vue')//路由的软加载
    },
    {
        path: '/Renwudiaodu_login',
        name: 'Renwudiaodu_login',
        // component: YonghuLogin,
        meta: {
            isShow: false,
        },
        component: () =>
            import ('../views/Renwudiaodu_login.vue')
            // component: () => import(/* webpackChunkName: "login" */ '..views/login.vue')//路由的软加载
    },
    {
        path: '/Jishu_login',
        name: 'Jishu_login',
        // component: YonghuLogin,
        meta: {
            isShow: false,
        },
        component: () =>
            import ('../views/Jishu_login.vue')
            // component: () => import(/* webpackChunkName: "login" */ '..views/login.vue')//路由的软加载
    },
    {
        path: '/YonghuZhuce',
        name: 'YonghuZhuce',
        // component: YonghuLogin,
        meta: {
            isShow: false,
        },
        component: () =>
            import ('../views/YonghuZhuce.vue')
            // component: () => import(/* webpackChunkName: "login" */ '..views/login.vue')//路由的软加载
    },
    // {
    //     path: '/engineer',
    //     name: 'engineer',
    //     // component: YonghuLogin,
    //     meta: {
    //         isShow: false,
    //     },
    //     component: () => import('../views/engineerlogin.vue')
    //     // component: () => import(/* webpackChunkName: "login" */ '..views/login.vue')//路由的软加载
    // },
    // {
    //     path: '/diaodu',
    //     name: 'diaodu',
    //     // component: YonghuLogin,
    //     meta: {
    //         isShow: false,
    //     },
    //     component: () => import('../views/diaodulogin.vue')
    //     // component: () => import(/* webpackChunkName: "login" */ '..views/login.vue')//路由的软加载
    // },
    // {
    //     path: '/kuguan',
    //     name: 'kuguan',
    //     // component: YonghuLogin,
    //     meta: {
    //         isShow: false,
    //     },
    //     component: () => import('../views/kuguanlogin.vue')
    //     // component: () => import(/* webpackChunkName: "login" */ '..views/login.vue')//路由的软加载
    // },
    // {
    //     path: '/Home',
    //     name: 'Home',
    //     component: Home,
    //     meta: {
    //         isShow: false,
    //     },
    //     children: [{
    //                 meta: {
    //                     isShow: true,
    //                     title: '填写维修信息'
    //                 },
    //                 path: '/Admin_control',
    //                 name: 'Admin_control',
    //                 component: () =>
    //                     import ( /* webpackChunkName: "login" */ '../views/Admin_control.vue'), //路由的软加载
    //                 // component: () =>
    //                 //     import ( /* webpackChunkName: "login" */ '../views/House_split.vue'), //路由的软加载

    //             },
    //             {
    //                 path: '/Score_correct',
    //                 name: 'Score_correct',
    //                 component: () =>
    //                     import ( /* webpackChunkName: "login" */ '../views/Score_correct.vue'), //路由的软加载
    //                 meta: {
    //                     isShow: true,
    //                     title: '查看维修任务'
    //                 }
    //             },
    // {
    //     path: '/Admin_split',
    //     name: 'Admin_split',
    //     component: () =>
    //         import ( /* webpackChunkName: "login" */ '../views/Admin_split.vue'), //路由的软加载
    //     meta: {
    //         isShow: true,
    //         title: '分房'
    //     }
    // },

    // ]
    // component: () => import(/* webpackChunkName: "login" */ '..views/login.vue')//路由的软加载
    //  },
    {
        path: '/Kefu_Home',
        name: 'Kefu_Home',
        meta: { title: '首页' },
        // redirect: '/index',
        component: () =>
            import ('../views/Kefu_Home/index.vue'),
        children: [{
                path: '/Kefu_index',
                name: 'Kefu_index',
                meta: { title: '主页面' },
                component: () =>
                    import ('../views/Kefu_Home/index/index.vue')
            },
            {
                path: '/Kefu_kehuxinxi',
                name: 'Kefu_kehuxinxi',
                meta: { title: '客户信息管理' },
                component: () =>
                    import ('../views/Kefu_Home/kehuxinxi/index.vue')
            },
            {
                path: '/Kefu_weixiuxinxi',
                name: 'Kefu_weixiuxinxi',
                meta: { title: '客户维修信息管理' },
                component: () =>
                    import ('../views/Kefu_Home/weixiuxinxi/index.vue'),
            },
            {
                path: '/Kefu_hesuanxinxi',
                name: 'Kefu_hesuanxinxi',
                meta: { title: '结算信息单' },
                component: () =>
                    import ('../views/Kefu_Home/hesuanxinxi/index.vue'),
            }
        ]
    },
    {
        path: '/Caiwu_Home',
        name: 'Caiwu_Home',
        meta: { title: '首页' },
        // redirect: '/index',
        component: () =>
            import ('../views/Caiwu_Home/index.vue'),
        children: [{
                path: '/Caiwu_index',
                name: 'Caiwu_index',
                meta: { title: '主页面' },
                component: () =>
                    import ('../views/Caiwu_Home/index/index.vue')
            },
            {
                path: '/Caiwu_hesuanxinxi',
                name: 'Caiwu_hesuanxinxi',
                meta: { title: '结算信息管理' },
                component: () =>
                    import ('../views/Caiwu_Home/hesuanxinxi/index.vue')
            }
        ]
    },



    //   binbin的路由
    {
        path: '/engineerhome', //布局页
        name: 'engineerhome',
        meta: { title: '工程师首页' },
        redirect: '/index3',
        component: () =>
            import ('../views/engineerhome/index3.vue'),
        children: [{
                path: '/index3', //kefu首页
                name: 'index3',
                meta: { title: '工程师首页' },
                component: () =>
                    import ('../views/engineerhome/index3/index.vue')
            },
            {
                path: '/all_function', //
                name: 'all_function',
                meta: { title: '操作维修信息' },
                component: () =>
                    import ('../views/engineerhome/all_function/index.vue')
            },
            // {
            //   path: '/maintance information', //
            //   name: 'maintance information',
            //   meta:{ title:'填写维修信息' },
            //   component: () => import( '../views/engineerhome/maintance information/index.vue')
            // },
            // {
            //   path: '/view maintance task', //查看维修信息
            //   name: 'view maintance task',
            //   meta:{ title:'查看维修任务' },
            //   component: () => import( '../views/engineerhome/view maintance task/index.vue')
            // },

        ]
    },
    {
        path: '/renwudiaoduhome', //布局页
        name: 'renwudiaoduhome',
        meta: { title: '任务调度首页' },
        redirect: '/index6',
        component: () =>
            import ('../views/renwudiaoduhome/index6.vue'),
        children: [{
                path: '/index6', //任务调度首页
                name: 'index6',
                meta: { title: '任务调度首页' },
                component: () =>
                    import ('../views/renwudiaoduhome/index6/index.vue')
            },
            {
                path: '/assign_task', //查看未分配的维修任务
                name: 'assign_task',
                meta: { title: '查看未分配的维修任务' },
                component: () =>
                    import ('../views/renwudiaoduhome/assign_task/index.vue')
            },
            {
                path: '/assigned_task_situation', //查看未分配的维修任务
                name: 'assigned_task_situation',
                meta: { title: '已分配的任务情况' },
                component: () =>
                    import ('../views/renwudiaoduhome/assigned_task_situation/index.vue')
            },


        ]
    },

    {
        path: '/kuguanhome', //布局页
        name: 'kuguanhome',
        meta: { title: '库管首页' },
        redirect: '/index5',
        component: () =>
            import ('../views/kuguanhome/index5.vue'),
        children: [{
                path: '/index5', //任务调度首页
                name: 'index5',
                meta: { title: '库管首页' },
                component: () =>
                    import ('../views/kuguanhome/index5/index.vue')
            },
            {
                path: '/add_kucun_information', //查看未分配的维修任务
                name: 'add_kucun_information',
                meta: { title: '入库' },
                component: () =>
                    import ('../views/kuguanhome/add_kucun_information/index.vue')
            },
            {
                path: '/update', //查看未分配的维修任务
                name: 'update',
                meta: { title: '添加备件' },
                component: () =>
                    import ('../views/kuguanhome/update/index.vue')
            },
            {
                path: '/search_kucun_information', //查看未分配的维修任务
                name: 'search_kucun_information',
                meta: { title: '查看库存信息' },
                component: () =>
                    import ('../views/kuguanhome/search_kucun_information/index.vue')
            },

        ]
    },
    {
        path: '/yonghuhome',
        name: 'yonghuhome',
        meta: { title: '首页' },
        // redirect: '/index',
        component: () =>
            import ('../views/yonghuhome/index9.vue'),
        children: [{
                path: '/index9',
                name: 'index9',
                meta: { title: '主页面' },
                component: () =>
                    import ('../views/yonghuhome/index9/index.vue')
            },
            {
                path: '/check_personal_information',
                name: 'check_personal_information',
                meta: { title: '查看修改个人信息' },
                component: () =>
                    import ('../views/yonghuhome/check_personal_information/index.vue')
            },
            {
                path: '/write_maintance_information',
                name: 'write_maintance_information',
                meta: { title: '填写维修信息' },
                component: () =>
                    import ('../views/yonghuhome/write_maintance_information/index.vue'),
            },
            {
                path: '/view_maintance_form',
                name: 'view_maintance_form',
                meta: { title: '查看维修单' },
                component: () =>
                    import ('../views/yonghuhome/view_maintance_form/index.vue'),
            },
            {
                path: '/state',
                name: 'state',
                meta: { title: '订单状态' },
                component: () =>
                    import ('../views/yonghuhome/state/index.vue'),
            }
        ]
    },

]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router