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
        console.warn(swipeNav)
        this.setData({
            swipe_nav: [
                {
                    "type": "apppage",
                    "jumptype": "redirect",
                    "appid": "",
                    "image": "https:\/\/www.watch-life.net\/images\/2021\/02\/\u5fae\u4fe1\u56fe\u7247_20210201093932-1-scaled.jpg",
                    "path": "\/pages\/detail\/detail?id=1827",
                    "url": "",
                    "enable": "yes",
                    "title": "\u5fae\u6155\u5c0f\u7a0b\u5e8f\u4e13\u4e1a\u7248V3.6-\u89c6\u9891\u53f7\u6269\u5c55"
                }, {
                    "type": "apppage",
                    "jumptype": "redirect",
                    "appid": "",
                    "image": "https:\/\/www.watch-life.net\/images\/2019\/07\/miapper-plus2.jpg",
                    "path": "\/pages\/detail\/detail?id=1773",
                    "url": "",
                    "enable": "yes",
                    "title": "\u5fae\u6155WordPress\u5c0f\u7a0b\u5e8f\u589e\u5f3a\u7248v2.0\u53d1\u5e03"
                }, {
                    "type": "apppage",
                    "jumptype": "redirect",
                    "appid": "",
                    "image": "https:\/\/www.watch-life.net\/images\/2019\/11\/minapper.jpg",
                    "path": "\/pages\/detail\/detail?id=1840",
                    "url": "",
                    "enable": "yes",
                    "title": "\u5fae\u6155\u5c0f\u7a0b\u5e8f\u5f00\u6e90\u7248v4.0\u53d1\u5e03"
                }, {
                    "type": "miniapp",
                    "jumptype": "redirect",
                    "appid": "wx074a216a7aabc30c",
                    "image": "https:\/\/www.watch-life.net\/images\/2021\/03\/4.jpg",
                    "path": "plugin-private:\/\/wx2b03c6e691cd7370\/pages\/live-player-plugin?room_id=15",
                    "url": "",
                    "enable": "yes",
                    "title": "\u4e13\u4e1a\u7248\u76f4\u64ad"
                }, {
                    "type": "apppage",
                    "jumptype": "redirect",
                    "appid": "",
                    "image": "https:\/\/www.watch-life.net\/images\/2022\/11\/tmp_841a2a6e549491c15ddd9ca807b98cb2.jpg",
                    "path": "\/pages\/my\/profile\/index",
                    "url": "",
                    "title": "\u6388\u6743",
                    "enable": "no"
                }
            ]
        })
        console.warn(this.data.swipe_nav)
    },
})