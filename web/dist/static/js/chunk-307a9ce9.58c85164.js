(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-307a9ce9","chunk-2d21a844","chunk-2d0b9e22","chunk-2d0aa631","chunk-2d0a3741","chunk-2d0e9955"],{"01f0":function(t,e,r){t.exports=r.p+"static/img/video.09ec4416.png"},"0564":function(t,e,r){"use strict";var a=r("6347"),n=r.n(a);n.a},1194:function(t,e,r){t.exports=r.p+"static/img/logo_login.7918108b.png"},"359c":function(t,e,r){t.exports=r.p+"static/img/github.5d298198.png"},6347:function(t,e,r){},"8dc6":function(t,e,r){t.exports=r.p+"static/img/yuque.5236e459.png"},bc9f:function(t,e,r){t.exports=r.p+"static/img/docs.20eecda0.png"},e49c:function(t,e,r){"use strict";r.r(e);var a=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"user-layout-wrapper",attrs:{id:"userLayout"}},[r("div",{staticClass:"container"},[t._m(0),r("div",{staticClass:"main"},[r("el-form",{ref:"loginForm",attrs:{model:t.loginForm,rules:t.rules},nativeOn:{keyup:function(e){return!e.type.indexOf("key")&&t._k(e.keyCode,"enter",13,e.key,"Enter")?null:t.submitForm(e)}}},[r("el-form-item",{attrs:{prop:"username"}},[r("el-input",{attrs:{placeholder:"请输入用户名"},model:{value:t.loginForm.username,callback:function(e){t.$set(t.loginForm,"username",e)},expression:"loginForm.username"}},[r("i",{staticClass:"el-input__icon el-icon-user",attrs:{slot:"suffix"},slot:"suffix"})])],1),r("el-form-item",{attrs:{prop:"password"}},[r("el-input",{attrs:{type:"lock"===t.lock?"password":"text",placeholder:"请输入密码"},model:{value:t.loginForm.password,callback:function(e){t.$set(t.loginForm,"password",e)},expression:"loginForm.password"}},[r("i",{class:"el-input__icon el-icon-"+t.lock,attrs:{slot:"suffix"},on:{click:t.changeLock},slot:"suffix"})])],1),r("el-form-item",{staticStyle:{position:"relative"}},[r("el-input",{staticStyle:{width:"60%"},attrs:{name:"logVerify",placeholder:"请输入验证码"},model:{value:t.loginForm.captcha,callback:function(e){t.$set(t.loginForm,"captcha",e)},expression:"loginForm.captcha"}}),r("div",{staticClass:"vPic"},[t.picPath?r("img",{attrs:{src:t.picPath,width:"100%",height:"100%",alt:"请输入验证码"},on:{click:function(e){return t.loginVefify()}}}):t._e()])],1),r("el-form-item",[r("el-button",{staticStyle:{width:"100%"},attrs:{type:"primary"},on:{click:t.submitForm}},[t._v("登 录")])],1)],1)],1),r("div",{staticClass:"footer"},[t._m(1),r("div",{staticClass:"copyright"},[t._v("Copyright © "+t._s(t.curYear)+" 💖flipped-aurora")])])])])},n=[function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"top"},[a("div",{staticClass:"desc"},[a("img",{staticClass:"logo_login",attrs:{src:r("1194"),alt:""}})]),a("div",{staticClass:"header"},[a("a",{attrs:{href:"/"}},[a("span",{staticClass:"title"},[t._v("slotsevo-web")])])])])},function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"links"},[a("a",{attrs:{href:"http://doc.henrongyi.top/"}},[a("img",{staticClass:"link-icon",attrs:{src:r("bc9f")}})]),a("a",{attrs:{href:"https://www.yuque.com/flipped-aurora/"}},[a("img",{staticClass:"link-icon",attrs:{src:r("8dc6")}})]),a("a",{attrs:{href:"https://github.com/flipped-aurora/slotsevo-web"}},[a("img",{staticClass:"link-icon",attrs:{src:r("359c")}})]),a("a",{attrs:{href:"https://space.bilibili.com/322210472"}},[a("img",{staticClass:"link-icon",attrs:{src:r("01f0")}})])])}],i=(r("96cf"),r("1da1")),s=r("5530"),o=r("5880"),c=r("c24f"),l={name:"Login",data:function(){var t=function(t,e,r){if(e.length<5)return r(new Error("请输入正确的用户名"));r()},e=function(t,e,r){if(e.length<6)return r(new Error("请输入正确的密码"));r()};return{curYear:0,lock:"lock",loginForm:{username:"xiaoliang",password:"123456",captcha:"",captchaId:""},rules:{username:[{validator:t,trigger:"blur"}],password:[{validator:e,trigger:"blur"}]},logVerify:"",picPath:""}},created:function(){this.loginVefify(),this.curYear=(new Date).getFullYear()},methods:Object(s["a"])(Object(s["a"])({},Object(o["mapActions"])("user",["LoginIn"])),{},{login:function(){var t=this;return Object(i["a"])(regeneratorRuntime.mark((function e(){return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,t.LoginIn(t.loginForm);case 2:return e.abrupt("return",e.sent);case 3:case"end":return e.stop()}}),e)})))()},submitForm:function(){var t=this;return Object(i["a"])(regeneratorRuntime.mark((function e(){return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,t.$refs.loginForm.validate(function(){var e=Object(i["a"])(regeneratorRuntime.mark((function e(r){var a;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:if(!r){e.next=7;break}return e.next=3,t.login();case 3:a=e.sent,a||t.loginVefify(),e.next=10;break;case 7:return t.$message({type:"error",message:"请正确填写登录信息",showClose:!0}),t.loginVefify(),e.abrupt("return",!1);case 10:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}());case 2:case"end":return e.stop()}}),e)})))()},changeLock:function(){"lock"===this.lock?this.lock="unlock":this.lock="lock"},loginVefify:function(){var t=this;Object(c["captcha"])({}).then((function(e){t.picPath=e.data.picPath,t.loginForm.captchaId=e.data.captchaId}))}})},u=l,p=(r("0564"),r("2877")),f=Object(p["a"])(u,a,n,!1,null,"aee01afc",null);e["default"]=f.exports}}]);