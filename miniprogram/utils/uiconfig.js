//引入框架的方法函数库
import ColorUI from '../mp-cu/main'
export const colorUI = new ColorUI({
   config: {
       theme: 'auto',
       main: 'blue',
       text: 1,
       footer: false,
       share: true,
       shareTitle: 'MP CU**（ ColorUI3.x 原生小程序版）**',
       homePath: '/pages/home/home',
       tabBar: [{
          title: '首页',
          icon: 'cicon-home-sm-o',
          curIcon: 'cicon-home-line',
          url: '/pages/index/index',
          type: 'tab'
      },
      {
          title: '我的',
          icon: 'cicon-my-o',
          curIcon: 'cicon-my',
          url: '/pages/my/my',
          type: 'tab'
      }],
   },
   data: {
      //全局data
   }, 
   methods: {
      //全局函数
   }
})