(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["signup"],{3041:function(s,e,a){},"34c3":function(s,e,a){"use strict";a.r(e);var t=function(){var s=this,e=s.$createElement,a=s._self._c||e;return a("div",{staticClass:"signup"},[a("Copyright"),a("div",{staticClass:"title"},[s._v(s._s(s.$locale.signup.title))]),a("div",{staticClass:"form"},[a("div",{staticClass:"form-item"},[a("label",{attrs:{for:"username"}},[s._v(s._s(s.$locale.signup.username))]),a("div",{staticClass:"custom-input"},[a("input",{directives:[{name:"model",rawName:"v-model",value:s.username,expression:"username"}],attrs:{id:"username",type:"text",placeholder:s.$locale.signup.usernamePlaceholder},domProps:{value:s.username},on:{keyup:function(e){return!e.type.indexOf("key")&&s._k(e.keyCode,"enter",13,e.key,"Enter")?null:s.handleSignup.apply(null,arguments)},input:function(e){e.target.composing||(s.username=e.target.value)}}}),s.username?a("div",{staticClass:"clean-input"},[a("i",{staticClass:"fas fa-times",on:{click:s.handleCleanUsername}})]):s._e()])]),a("div",{staticClass:"form-item"},[a("label",{attrs:{for:"account"}},[s._v(s._s(s.$locale.signup.account))]),a("div",{staticClass:"custom-input"},[a("input",{directives:[{name:"model",rawName:"v-model",value:s.account,expression:"account"}],attrs:{id:"account",type:"text",placeholder:s.$locale.signup.accountPlaceholder},domProps:{value:s.account},on:{keyup:function(e){return!e.type.indexOf("key")&&s._k(e.keyCode,"enter",13,e.key,"Enter")?null:s.handleSignup.apply(null,arguments)},input:function(e){e.target.composing||(s.account=e.target.value)}}}),s.account?a("div",{staticClass:"clean-input"},[a("i",{staticClass:"fas fa-times",on:{click:s.handleCleanAccount}})]):s._e()])]),a("div",{staticClass:"form-item"},[a("label",{attrs:{for:"password"}},[s._v(s._s(s.$locale.signup.password))]),a("div",{staticClass:"custom-input"},[a("input",{directives:[{name:"model",rawName:"v-model",value:s.password,expression:"password"}],attrs:{id:"password",type:"password",placeholder:s.$locale.signup.passwordPlaceholder},domProps:{value:s.password},on:{keyup:function(e){return!e.type.indexOf("key")&&s._k(e.keyCode,"enter",13,e.key,"Enter")?null:s.handleSignup.apply(null,arguments)},input:function(e){e.target.composing||(s.password=e.target.value)}}}),s.password?a("div",{staticClass:"clean-input"},[a("i",{staticClass:"fas fa-times",on:{click:s.handleCleanPassword}})]):s._e()])]),a("div",{staticClass:"form-item"},[a("label",{attrs:{for:"confirmPassword"}},[s._v(s._s(s.$locale.signup.confirmPassword))]),a("div",{staticClass:"custom-input"},[a("input",{directives:[{name:"model",rawName:"v-model",value:s.confirmPassword,expression:"confirmPassword"}],attrs:{id:"confirmPassword",type:"password",placeholder:s.$locale.signup.confirmPasswordPlaceholder},domProps:{value:s.confirmPassword},on:{keyup:function(e){return!e.type.indexOf("key")&&s._k(e.keyCode,"enter",13,e.key,"Enter")?null:s.handleSignup.apply(null,arguments)},input:function(e){e.target.composing||(s.confirmPassword=e.target.value)}}}),s.confirmPassword?a("div",{staticClass:"clean-input"},[a("i",{staticClass:"fas fa-times",on:{click:s.handleCleanConfirmPassword}})]):s._e()])]),a("div",{staticClass:"form-item"},[a("button",{staticClass:"btn-submit",on:{click:s.handleSignup}},[s._v(" "+s._s(s.$locale.signup.submit)+" ")])]),a("div",{staticClass:"form-item back-option"},[s._v(" "+s._s(s.$locale.signup.back)+" "),a("router-link",{attrs:{to:"/login"}},[s._v(s._s(s.$locale.signup.login))])],1)]),a("div",{staticClass:"float-panel"},[a("BtnLangSwap")],1)],1)},n=[],i=a("8053"),o=a("48e3"),r={components:{Copyright:i["a"],BtnLangSwap:o["a"]},data(){return{username:"",account:"",password:"",confirmPassword:""}},methods:{handleSignup(){this.password==this.confirmPassword?(this.$startLoading(),this.$http.post("/api/signup",{username:this.username,account:this.account,password:this.password}).then(s=>{if(s.data.success){let e=s.data.data;this.$setUser({id:e.userid,name:e.username,role:e.role,permission:e.permission,usedSpace:e.usedSpace,allSpace:e.allSpace,token:e.token}),this.$router.push("/"),this.$stopLoading()}else this.$showError(this.$locale.common.errorMsg),this.$stopLoading()}).catch(s=>{this.$showError(this.$locale.common.errorServer),this.$stopLoading()})):this.$showError(this.$locale.common.errorConfirmPassword)},handleCleanUsername(){this.username=""},handleCleanAccount(){this.account=""},handleCleanPassword(){this.password=""},handleCleanConfirmPassword(){this.confirmPassword=""}}},l=r,c=(a("eea4"),a("0c7c")),u=Object(c["a"])(l,t,n,!1,null,"6cea6997",null);e["default"]=u.exports},"48e3":function(s,e,a){"use strict";var t=function(){var s=this,e=s.$createElement,a=s._self._c||e;return a("div",{staticClass:"btn-lang-swap",attrs:{title:s.$locale.common.btnLangTips},on:{click:s.$swapLang}},[a("span",{staticClass:"text"},[s._v(s._s(s.$locale.common.langTag))])])},n=[],i=(a("f566"),a("0c7c")),o={},r=Object(i["a"])(o,t,n,!1,null,"726851f1",null);e["a"]=r.exports},8053:function(s,e,a){"use strict";var t=function(){var s=this,e=s.$createElement,a=s._self._c||e;return a("div",{staticClass:"copyright"},[s._v("© 2021 jinyaoMa")])},n=[],i=(a("f4d6"),a("0c7c")),o={},r=Object(i["a"])(o,t,n,!1,null,"43e0ccf1",null);e["a"]=r.exports},b4fc:function(s,e,a){},b506:function(s,e,a){},eea4:function(s,e,a){"use strict";a("3041")},f4d6:function(s,e,a){"use strict";a("b4fc")},f566:function(s,e,a){"use strict";a("b506")}}]);
//# sourceMappingURL=signup.b0d4f6fb.js.map