// index.js
import config from '../../utils/config.js'
var Api = require('../../utils/api.js');
var wxRequest = require('../../utils/wxRequest.js')
var util = require('../../utils/util.js');
var webSiteName = config.getWebsiteName;
var domain = config.getDomain;
var pageCount = config.getPageCount;

Page({
  data: {
    postsList: [],
    postsstickyList: [],
    postsShowSwiperList: [],
    isLastPage: false,
    page: 1,
    search: '',
    categories: 0,
    showerror: "none",
    showCategoryName: "",
    categoryName: "",
    floatDisplay: "none",
    listAdsuccess: true,
    webSiteName: webSiteName,
    domain: domain,
    isFirst: true, // 是否第一次打开,
    isLoading: true,
    swipe_nav: [],
    selected_nav: []
  },
  onLoad: function (options) {
    var self = this;
    self.fetchTopFivePosts();
    this.getHomeconfig();
    this.fetchPostsData();
  },
  getHomeconfig() {
    //获取扩展设置
    var self = this;

    var getHomeconfig = wxRequest.getRequest(Api.get_homeconfig());
    getHomeconfig.then(res => {
      // console.log(res.data);
      let expand = res.data.expand;
      let swipe_nav = expand.swipe_nav;
      let selected_nav = expand.selected_nav;
      let _d = res.data.downloadfileDomain
      let _b = res.data.businessDomain

      let zanImageurl = res.data.zanImageurl
      let logoImageurl = res.data.logoImageurl

      let downloadfileDomain = _d.length ? _d.split(',') : []
      let businessDomain = _b.length ? _b.split(',') : []
      self.setData({
        swipe_nav: swipe_nav,
        selected_nav,
      });
      wx.setStorageSync('downloadfileDomain', downloadfileDomain);
      wx.setStorageSync('businessDomain', businessDomain);
      wx.setStorageSync('zanImageurl', zanImageurl);
      wx.setStorageSync('logoImageurl', logoImageurl);
    });
  },
  fetchPostsData: function (data) {
    var self = this;
    if (!data) data = {};
    if (!data.page) data.page = 1;
    if (!data.categories) data.categories = 0;
    if (!data.search) data.search = '';
    if (data.page === 1) {
      self.setData({
        postsList: []
      });
    };
    self.setData({
      isLoading: true
    })
    var getCategoriesRequest = wxRequest.getRequest(Api.getCategoriesIds());
    getCategoriesRequest.then(res => {
      if (!res.data.Ids == "") {
        data.categories = res.data.Ids;
        self.setData({
          categories: res.data.Ids
        })

      }

      var getPostsRequest = wxRequest.getRequest(Api.getPosts(data));
      getPostsRequest
        .then(response => {
          if (response.statusCode === 200) {
            if (response.data.length) {
              if (response.data.length < pageCount) {
                self.setData({
                  isLastPage: true,
                  isLoading: false
                });
              }
              self.setData({
                floatDisplay: "block",
                postsList: self.data.postsList.concat(response.data.map(function (item) {

                  var strdate = item.date

                  item.categoryImage = "";


                  if (item.post_medium_image == null || item.post_medium_image == '') {
                    item.post_medium_image = "../../images/logo700.png";
                  }
                  item.date = util.cutstr(strdate, 10, 1);
                  return item;
                })),

              });

            } else {
              if (response.data.code == "rest_post_invalid_page_number") {
                self.setData({
                  isLastPage: true,
                  isLoading: false
                });
                wx.showToast({
                  title: '没有更多内容',
                  mask: false,
                  duration: 1500
                });
              } else {
                wx.showToast({
                  title: response.data.message,
                  duration: 1500
                })
              }
            }
          }
        })
        .catch(function (response) {
          if (data.page == 1) {

            self.setData({
              showerror: "block",
              floatDisplay: "none"
            });

          } else {
            wx.showModal({
              title: '加载失败',
              content: '加载数据失败,请重试.',
              showCancel: false,
            });
            self.setData({
              page: data.page - 1
            });
          }

        })
        .finally(function (response) {
          wx.hideLoading();
          self.setData({
            isLoading: false
          })
          wx.stopPullDownRefresh();
        });

    })


  },
  fetchTopFivePosts: function (data) {
    var self = this;
    var getCategoriesRequest = wxRequest.getRequest(Api.getCategoriesIds());
    getCategoriesRequest.then(res => {


      var getPostsRequest = wxRequest.getRequest(Api.getStickyPosts(data));
      getPostsRequest
        .then(response => {
          if (response.statusCode === 200) {
            if (response.data.length) {

              self.setData({
                floatDisplay: "block",
                postsstickyList: self.data.postsstickyList.concat(response.data.map(function (item) {

                  var strdate = item.date

                  item.categoryImage = "";


                  if (item.post_medium_image == null || item.post_medium_image == '') {
                    item.post_medium_image = "../../images/logo700.png";
                  }
                  item.date = util.cutstr(strdate, 10, 1);
                  return item;
                })),

              });
            } else {
              if (response.data.code == "rest_post_invalid_page_number") {
                self.setData({
                  isLastPage: true,
                  isLoading: false
                });
                wx.showToast({
                  title: '没有更多内容',
                  mask: false,
                  duration: 1500
                });
              } else {
                wx.showToast({
                  title: response.data.message,
                  duration: 1500
                })
              }
            }
          }
        })



    })


  },
  loadMore: function (e) {

    var self = this;
    if (!self.data.isLastPage) {
      self.setData({
        page: self.data.page + 1
      });
      //console.log('当前页' + self.data.page);
      this.fetchPostsData(self.data);
    } else {
      wx.showToast({
        title: '没有更多内容',
        mask: false,
        duration: 1000
      });
    }
  },
  onPullDownRefresh: function () {
    var self = this;
    self.setData({
      showerror: "none",
      floatDisplay: "none",
      isLastPage: false,
      page: 1,
      postsShowSwiperList: [],
      listAdsuccess: true

    });
    console.warn(123123)
    this.getHomeconfig();
    this.fetchPostsData(self.data);


  },
  onReachBottom: function () {

    var self = this;
    console.warn(123123123)
    if (!self.data.isLastPage) {
      self.setData({
        page: self.data.page + 1
      });
      console.log('当前页' + self.data.page);
      this.fetchPostsData(self.data);
    } else {
      console.log('最后一页');
    }

  },
  formSubmit: function (e) {
    var url = '../list/list';
    var key = '';
    if (e.currentTarget.id == 'search-input') {
      key = e.detail.value;
    } else {
      key = e.detail.value.input;
    }
    if (key != '') {
      url = url + '?search=' + key;
      wx.navigateTo({
        url: url,
      })
    } else {
      wx.showModal({
        title: '提示',
        content: '请输入内容',
        showCancel: false
      })
    }
  },
  redictAppDetail: function (e) {
    let {
      type,
      appid,
      url,
      path,
      jumptype
    } = e.currentTarget.dataset;
    console.warn(this.data.swipe_nav)
    console.warn(e.currentTarget.dataset)
    if (type === 'apppage') { // 小程序页面
      wx.navigateTo({
        url: path
      })
    }
    if (type === 'webpage') { // web-view页面
      url = '../webpage/webpage?url=' + encodeURIComponent(url)
      wx.navigateTo({
        url: url
      })
    }
    if (type === 'miniapp') { // 其他小程序
      if (jumptype == 'embedded') {
        wx.openEmbeddedMiniProgram({
          appId: appid,
          path: path
        })

      }
      else {
        wx.navigateToMiniProgram({
          appId: appid,
          path: path
        })
      }

    }
  }
})
