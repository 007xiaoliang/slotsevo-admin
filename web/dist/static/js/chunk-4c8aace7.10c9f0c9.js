(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-4c8aace7"],{5004:function(s,e,a){"use strict";a.r(e);var t=function(){var s=this,e=s.$createElement,a=s._self._c||e;return a("div",[a("el-row",[a("el-col",{attrs:{span:6}},[a("div",{staticClass:"fl-left avatar-box"},[a("div",{staticClass:"user-card"},[a("div",{staticClass:"user-headpic-update",style:{"background-image":"url("+(s.userInfo.headerImg&&"http"!==s.userInfo.headerImg.slice(0,4)?s.path+s.userInfo.headerImg:s.userInfo.headerImg)+")","background-repeat":"no-repeat","background-size":"cover"}},[a("span",{staticClass:"update",on:{click:s.openChooseImg}},[a("i",{staticClass:"el-icon-edit"}),s._v(" 重新上传")])]),a("div",{staticClass:"user-personality"},[a("p",{staticClass:"nickname"},[s._v(s._s(s.userInfo.nickName))]),a("p",{staticClass:"person-info"},[s._v("这个家伙很懒，什么都没有留下")])]),a("div",{staticClass:"user-information"},[a("ul",[a("li",[a("i",{staticClass:"el-icon-user"}),s._v(s._s(s.userInfo.nickName)+" ")]),a("li",[a("i",{staticClass:"el-icon-data-analysis"}),s._v("北京反转极光科技有限公司-技术部-前端事业群 ")]),a("li",[a("i",{staticClass:"el-icon-video-camera-solid"}),s._v("中国·北京市·朝阳区 ")]),a("li",[a("i",{staticClass:"el-icon-medal-1"}),s._v("goLang/JavaScript/Vue/Gorm ")])])])])])]),a("el-col",{attrs:{span:18}},[a("div",{staticClass:"user-addcount"},[a("el-tabs",{on:{"tab-click":s.handleClick},model:{value:s.activeName,callback:function(e){s.activeName=e},expression:"activeName"}},[a("el-tab-pane",{attrs:{label:"账号绑定",name:"second"}},[a("ul",[a("li",[a("p",{staticClass:"title"},[s._v("密保手机")]),a("p",{staticClass:"desc"},[s._v(" 已绑定手机:1245678910 "),a("a",{attrs:{href:"#"}},[s._v("立即修改")])])]),a("li",[a("p",{staticClass:"title"},[s._v("密保邮箱")]),a("p",{staticClass:"desc"},[s._v(" 已绑定邮箱：slotsevo-web@google.com.cn "),a("a",{attrs:{href:"#"}},[s._v("立即修改")])])]),a("li",[a("p",{staticClass:"title"},[s._v("密保问题")]),a("p",{staticClass:"desc"},[s._v(" 未设置密保问题 "),a("a",{attrs:{href:"#"}},[s._v("去设置")])])]),a("li",[a("p",{staticClass:"title"},[s._v("修改密码")]),a("p",{staticClass:"desc"},[s._v(" 修改个人密码 "),a("a",{attrs:{href:"#"},on:{click:function(e){s.showPassword=!0}}},[s._v("修改密码")])])])])])],1)],1)])],1),a("el-dialog",{attrs:{visible:s.showPassword,title:"修改密码",width:"360px"},on:{"update:visible":function(e){s.showPassword=e},close:s.clearPassword}},[a("el-form",{ref:"modifyPwdForm",attrs:{model:s.pwdModify,rules:s.rules,"label-width":"80px"}},[a("el-form-item",{attrs:{minlength:6,label:"原密码",prop:"password"}},[a("el-input",{attrs:{"show-password":""},model:{value:s.pwdModify.password,callback:function(e){s.$set(s.pwdModify,"password",e)},expression:"pwdModify.password"}})],1),a("el-form-item",{attrs:{minlength:6,label:"新密码",prop:"newPassword"}},[a("el-input",{attrs:{"show-password":""},model:{value:s.pwdModify.newPassword,callback:function(e){s.$set(s.pwdModify,"newPassword",e)},expression:"pwdModify.newPassword"}})],1),a("el-form-item",{attrs:{minlength:6,label:"确认密码",prop:"confirmPassword"}},[a("el-input",{attrs:{"show-password":""},model:{value:s.pwdModify.confirmPassword,callback:function(e){s.$set(s.pwdModify,"confirmPassword",e)},expression:"pwdModify.confirmPassword"}})],1)],1),a("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[a("el-button",{on:{click:function(e){s.showPassword=!1}}},[s._v("取 消")]),a("el-button",{attrs:{type:"primary"},on:{click:s.savePassword}},[s._v("确 定")])],1)],1)],1)},r=[],o=(a("96cf"),a("1da1")),i=a("5530"),n=a("c24f"),l=a("5880"),c="/api",d={name:"Person",data:function(){var s=this;return{path:c,activeName:"second",showPassword:!1,pwdModify:{},rules:{password:[{required:!0,message:"请输入密码",trigger:"blur"},{min:6,message:"最少6个字符",trigger:"blur"}],newPassword:[{required:!0,message:"请输入新密码",trigger:"blur"},{min:6,message:"最少6个字符",trigger:"blur"}],confirmPassword:[{required:!0,message:"请输入确认密码",trigger:"blur"},{min:6,message:"最少6个字符",trigger:"blur"},{validator:function(e,a,t){a!==s.pwdModify.newPassword?t(new Error("两次密码不一致")):t()},trigger:"blur"}]}}},computed:Object(i["a"])({},Object(l["mapGetters"])("user",["userInfo","token"])),methods:Object(i["a"])(Object(i["a"])({},Object(l["mapMutations"])("user",["ResetUserInfo"])),{},{savePassword:function(){var s=this;this.$refs.modifyPwdForm.validate((function(e){if(!e)return!1;Object(n["changePassword"])({username:s.userInfo.userName,password:s.pwdModify.password,newPassword:s.pwdModify.newPassword}).then((function(e){0===e.code&&s.$message.success("修改密码成功！"),s.showPassword=!1}))}))},clearPassword:function(){this.pwdModify={password:"",newPassword:"",confirmPassword:""},this.$refs.modifyPwdForm.clearValidate()},openChooseImg:function(){this.$refs.chooseImg.open()},enterImg:function(s){var e=this;return Object(o["a"])(regeneratorRuntime.mark((function a(){var t;return regeneratorRuntime.wrap((function(a){while(1)switch(a.prev=a.next){case 0:return a.next=2,Object(n["setUserInfo"])({headerImg:s,ID:e.userInfo.ID});case 2:t=a.sent,0===t.code&&(e.ResetUserInfo({headerImg:s}),e.$message({type:"success",message:"设置成功"}));case 4:case"end":return a.stop()}}),a)})))()},handleClick:function(s,e){console.log(s,e)}})},u=d,w=(a("673e"),a("2877")),p=Object(w["a"])(u,t,r,!1,null,null,null);e["default"]=p.exports},"673e":function(s,e,a){"use strict";var t=a("ddb2"),r=a.n(t);r.a},ddb2:function(s,e,a){}}]);