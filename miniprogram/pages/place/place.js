var Api = require('../../utils/api.js');
var wxRequest = require('../../utils/wxRequest.js')
Page({
    data: {
        page: 1,
        search: '',
        swipe_nav: []
    },
    onLoad: function (options) {
        var swipeNav = wxRequest.getRequest(Api.getSwipeNav());
        swipeNav.then(swipeNav => {
            console.warn(swipeNav.data.data.list)
            this.setData({
                swipe_nav: swipeNav.data.data.list
            })
        })
    },
})