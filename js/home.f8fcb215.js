(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["home"],{"1ac2":function(e,t,s){"use strict";s("1b5f")},"1b5f":function(e,t,s){},"279d":function(e,t,s){},"28f0":function(e,t,s){"use strict";var a=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("a",{staticClass:"btn-download-file",attrs:{href:e.url,download:"file",title:e.$locale.common.download}},[s("i",{staticClass:"fas fa-download fa-fw"})])},r=[],i={props:{url:{type:String}}},n=i,l=(s("1ac2"),s("0c7c")),o=Object(l["a"])(n,a,r,!1,null,"b330b7f6",null);t["a"]=o.exports},"4fcd":function(e,t,s){"use strict";s("b85c")},"65f6":function(e,t,s){"use strict";var a=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("span",{staticClass:"btn-rename",attrs:{title:e.$locale.common.rename},on:{click:function(t){return e.$emit("click")}}},[s("i",{staticClass:"fas fa-pen fa-fw"})])},r=[],i=(s("a44f"),s("0c7c")),n={},l=Object(i["a"])(n,a,r,!1,null,"57d98a88",null);t["a"]=l.exports},"9e8e":function(e,t,s){},a44f:function(e,t,s){"use strict";s("279d")},a507:function(e,t,s){"use strict";s("9e8e")},a985:function(e,t,s){"use strict";var a=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("span",{staticClass:"btn-sharing",attrs:{title:e.customTitle},on:{click:function(t){return e.$emit("click")}}},[e.forUpdate?s("i",{staticClass:"fas fa-sync-alt fa-fw"}):s("i",{staticClass:"fas fa-share-alt fa-fw"})])},r=[],i={props:{forUpdate:{type:Boolean,default(){return!1}}},computed:{customTitle(){return this.forUpdate?this.$locale.common.updateShare:this.$locale.common.sharing}}},n=i,l=(s("f7a1"),s("0c7c")),o=Object(l["a"])(n,a,r,!1,null,"7d44a390",null);t["a"]=o.exports},b251:function(e,t,s){},b85c:function(e,t,s){},bb51:function(e,t,s){"use strict";s.r(t);var a=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("Layout",[s("div",{staticClass:"home"},[s("TaskBar",{ref:"taskBar",attrs:{currentFilesCount:e.currentFiles.length,filesToDownload:e.filesToDownload,cutOptions:e.cutOptions,canNewFolder:0===e.searchKeyword.length},on:{search:e.handleSearch,newFolder:e.handleNewFolderClick,upload:e.handleUploadFile}}),s("div",{staticClass:"home-frame"},[s("Breadcrumb",{attrs:{checkedCount:e.checkedItemsAmount,allChecked:e.checkedItemsAmount===e.currentFiles.length}}),s("table",[s("thead",[s("tr",[s("td",{staticClass:"col-select"},[s("CheckSquare",{attrs:{prepared:e.checkedItemsAmount>0&&e.checkedItemsAmount<e.currentFiles.length,checked:e.checkedItemsAmount===e.currentFiles.length},on:{click:e.handleAllCheck}})],1),s("td",{on:{click:e.handleSortName}},[e._v(" "+e._s(e.$locale.home.filename)+" "),"Name"===e.sortBy?["DESC"===e.sortOrder?s("i",{staticClass:"fas fa-sort-alpha-up"}):s("i",{staticClass:"fas fa-sort-alpha-down"})]:e._e()],2),s("td",{staticClass:"col-size",on:{click:e.handleSortSize}},[e._v(" "+e._s(e.$locale.home.size)+" "),"Size"===e.sortBy?["DESC"===e.sortOrder?s("i",{staticClass:"fas fa-sort-numeric-up-alt"}):s("i",{staticClass:"fas fa-sort-numeric-down"})]:e._e()],2),s("td",{staticClass:"col-modified-at",on:{click:e.handleSortUpdatedAt}},[e._v(" "+e._s(e.$locale.home.modifiedAt)+" "),"UpdatedAt"===e.sortBy?["DESC"===e.sortOrder?s("i",{staticClass:"fas fa-sort-amount-up"}):s("i",{staticClass:"fas fa-sort-amount-down"})]:e._e()],2)])]),s("tbody",[null!=e.parentPath||e.searchKeyword?s("tr",[s("td",{staticClass:"col-select"}),s("td",{attrs:{colspan:"3"}},[s("div",{staticClass:"col-filename"},[e.searchKeyword?s("span",{staticClass:"col-filename-start back",on:{click:e.handleClearOpts}},[s("i",{staticClass:"fas fa-arrow-left fa-fw"}),e._v(" "+e._s(e.$locale.home.endSearch)+" ")]):s("router-link",{staticClass:"col-filename-start back",attrs:{to:{query:{currentPath:e.parentPath}}}},[s("i",{staticClass:"fas fa-reply fa-fw"}),e._v(" "+e._s(e.$locale.home.back2Parent)+" ")])],1)])]):e._e(),e.showNewFolderForm?s("tr",[s("td",{staticClass:"col-select"}),s("td",{attrs:{colspan:"3"}},[s("div",{staticClass:"col-filename"},[s("div",{staticClass:"col-filename-start"},[s("div",{staticClass:"form-rename"},[s("input",{directives:[{name:"model",rawName:"v-model",value:e.newFolderName,expression:"newFolderName"}],attrs:{type:"text"},domProps:{value:e.newFolderName},on:{keyup:function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:e.handleNewFolder.apply(null,arguments)},input:function(t){t.target.composing||(e.newFolderName=t.target.value)}}}),s("i",{staticClass:"fas fa-check fa-fw",on:{click:e.handleNewFolder}}),s("i",{staticClass:"fas fa-times fa-fw",on:{click:function(t){e.showNewFolderForm=!1}}})])])])])]):e._e(),e._l(e.currentFiles,(function(t){return s("tr",{key:t.ID},[s("td",{staticClass:"col-select"},[s("CheckSquare",{attrs:{fileId:t.ID,checked:t.isChecked},on:{click:e.handleSingleCheck}})],1),s("td",[s("div",{staticClass:"col-filename"},["directory"===t.Type?s("div",{staticClass:"col-filename-start"},[t.canRename?s("div",{staticClass:"form-rename"},[s("input",{directives:[{name:"model",rawName:"v-model",value:e.renameFileName,expression:"renameFileName"}],attrs:{type:"text"},domProps:{value:e.renameFileName},on:{keyup:function(s){return!s.type.indexOf("key")&&e._k(s.keyCode,"enter",13,s.key,"Enter")?null:e.handleRenameFile(t)},input:function(t){t.target.composing||(e.renameFileName=t.target.value)}}}),s("i",{staticClass:"fas fa-check fa-fw",on:{click:function(s){return e.handleRenameFile(t)}}}),s("i",{staticClass:"fas fa-times fa-fw",on:{click:function(e){t.canRename=!1}}})]):s("div",[s("router-link",{attrs:{to:{query:{currentPath:t.Path}}}},[s("i",{staticClass:"fas fa-folder fa-fw"}),e._v(" "+e._s(t.Name)+" "),e.searchKeyword?s("span",{staticClass:"hint"},[e._v(" "+e._s(t.Path)+" ")]):e._e()])],1)]):s("div",{staticClass:"col-filename-start"},[t.canRename?s("div",{staticClass:"form-rename"},[s("input",{directives:[{name:"model",rawName:"v-model",value:e.renameFileName,expression:"renameFileName"}],attrs:{type:"text"},domProps:{value:e.renameFileName},on:{keyup:function(s){return!s.type.indexOf("key")&&e._k(s.keyCode,"enter",13,s.key,"Enter")?null:e.handleRenameFile(t)},input:function(t){t.target.composing||(e.renameFileName=t.target.value)}}}),s("i",{staticClass:"fas fa-check fa-fw",on:{click:function(s){return e.handleRenameFile(t)}}}),s("i",{staticClass:"fas fa-times fa-fw",on:{click:function(e){t.canRename=!1}}})]):s("div",[s("i",{staticClass:"fas fa-file fa-fw"}),e._v(" "+e._s(t.Name)+" "),e.searchKeyword?s("span",{staticClass:"hint"},[e._v(" "+e._s(t.Path)+" ")]):e._e()])]),s("div",{staticClass:"col-filename-end"},["file"===t.Type?s("BtnDownloadFile",{attrs:{"data-check":t.isChecked,url:e.getFileDownloadLink(t)}}):e._e(),"file"===t.Type?s("BtnSharing",{on:{click:function(s){return e.handleSharingClick(t)}}}):e._e(),"file"===t.Type?s("BtnCut",{on:{click:function(s){return e.handleCutClick(t)}}}):e._e(),s("BtnRename",{on:{click:function(s){return e.handleRenameClick(t)}}}),s("BtnRecycling",{on:{click:function(s){return e.handleRecyclingClick(t)}}})],1)])]),"directory"===t.Type?s("td",[e._v("-")]):s("td",[e._v(e._s(e.$convertSpace2String(t.Size)))]),s("td",[e._v(e._s(e.$formatDate(t.UpdatedAt,"YYYY-MM-DD HH:mm")))])])})),e.offset<e.currentSortedFiles.length?s("tr",[s("td",{staticClass:"more",attrs:{colspan:"4"},on:{click:function(t){e.offset+=10}}},[e._v(" "+e._s(e.$locale.common.loadMore)+" ")])]):e._e()],2)])],1)],1)])},r=[],i=s("ebad"),n=s("db9f"),l=s("b9bc"),o=s("2fbc"),c=s("28f0"),h=s("65f6"),d=s("fe9b"),u=s("a985"),f=s("b747"),m={components:{Layout:i["a"],TaskBar:n["a"],CheckSquare:l["a"],Breadcrumb:o["a"],BtnDownloadFile:c["a"],BtnRename:h["a"],BtnCut:d["a"],BtnSharing:u["a"],BtnRecycling:f["a"]},watch:{$route:{handler(){this.handleClearOpts()}}},methods:{handleUploadFile(e){let t=new FormData;t.append("id",this.$user.id),t.append("token",this.$user.token),t.append("desId",this.$getCurrentPathId()),t.append("file",e),this.$setProgress(0),this.$startProgressing(),this.$http.post("/api/service/upload",t,{headers:{"Content-Type":"multipart/form-data;charset=utf-8"},transformRequest:[function(e){return e}],onUploadProgress:e=>{let t=e.loaded/e.total*100|0;this.$setProgress(t)}}).then(e=>{if(e.data.success){let t=e.data.data;this.$setUser({usedSpace:t.usedSpace,allSpace:t.allSpace}),this.$setFiles(t.files)}else this.$showError(this.$locale.common.errorMsg);this.$stopProgressing()}).catch(e=>{this.$showError(this.$locale.common.errorServer),this.$stopProgressing()})},handleRecyclingClick(e){this.$startLoading(),this.$http.post("/api/files/recycle",{id:this.$user.id,token:this.$user.token,fileId:e.ID}).then(e=>{e.data.success?this.$setFiles(e.data.data.files):this.$showError(this.$locale.common.errorMsg),this.$stopLoading()}).catch(e=>{this.$showError(this.$locale.common.errorServer),this.$stopLoading()})},handleNewFolder(){this.$startLoading(),this.$http.post("/api/files/newFolder",{id:this.$user.id,token:this.$user.token,folderName:this.newFolderName,desId:this.$getCurrentPathId()}).then(e=>{e.data.success?(this.$setFiles(e.data.data.files),this.showNewFolderForm=!1):this.$showError(this.$locale.common.errorMsg),this.$stopLoading()}).catch(e=>{this.$showError(this.$locale.common.errorServer),this.$stopLoading()})},handleNewFolderClick(){this.newFolderName=this.$locale.common.newFolder,this.showNewFolderForm=!0},handleSharingClick(e){this.$startLoading(),this.$http.post("/api/files/shareFile",{id:this.$user.id,token:this.$user.token,fileId:e.ID}).then(e=>{if(e.data.success){let t=e.data.data,s=`${this.$http.defaults.baseURL}/api/service/download?c=${t.c}&d=${t.d}&e=${t.e}`;this.$setFiles(t.files),this.$showShareNotice(s)}else this.$showError(this.$locale.common.errorMsg);this.$stopLoading()}).catch(e=>{this.$showError(this.$locale.common.errorServer),this.$stopLoading()})},handleCutClick(e){this.cutOptions.srcId=e.ID,this.cutOptions.srcPath=e.Path,this.cutOptions.canPaste=!0},handleRenameFile(e){this.$startLoading(),this.$http.post("/api/files/renameFile",{id:this.$user.id,token:this.$user.token,fileId:this.renameFileId,filename:this.renameFileName}).then(t=>{t.data.success?(e.Name=this.renameFileName,e.canRename=!1,this.renameFileId=0,this.renameFileName=""):this.$showError(this.$locale.common.errorMsg),this.$stopLoading()}).catch(e=>{this.$showError(this.$locale.common.errorServer),this.$stopLoading()})},handleRenameClick(e){this.currentFiles.forEach(e=>{e.canRename=!1}),e.canRename=!0,this.renameFileId=e.ID,this.renameFileName=e.Name},getFileDownloadLink(e){let t=this.$user.usedSpace%64;return`${this.$http.defaults.baseURL}/api/service/download?a=${t+1}&b=${this.$user.token.substr(t)}&c=${e.ShareCode}&d=${e.ID}`},handleAllCheck(e,t){t?this.currentFiles.forEach(e=>{e.isChecked=!0}):this.$clearChecklist()},handleSingleCheck(e,t){this.$setFileCheckById({fileId:e,isChecked:t})},handleSortName(){"Name"===this.sortBy?this.sortOrder="ASC"===this.sortOrder?"DESC":"ASC":(this.sortBy="Name",this.sortOrder="ASC")},handleSortSize(){"Size"===this.sortBy?this.sortOrder="DESC"===this.sortOrder?"ASC":"DESC":(this.sortBy="Size",this.sortOrder="DESC")},handleSortUpdatedAt(){"UpdatedAt"===this.sortBy?this.sortOrder="DESC"===this.sortOrder?"ASC":"DESC":(this.sortBy="UpdatedAt",this.sortOrder="DESC")},handleClearOpts(){this.offset=10,this.searchKeyword="",this.$refs.taskBar.clearSearchKeyword(),this.$clearChecklist()},handleSearch(e,t){this.searchKeyword=e,t()}},computed:{filesToDownload(){return this.currentUnsortedFiles.filter(e=>"file"===e.Type&&e.isChecked)},checkedItemsAmount(){return this.currentUnsortedFiles.filter(e=>e.isChecked).length},parentPath(){let e=this.$route.query.currentPath;if("string"===typeof e&&e.length>0){let t=e.lastIndexOf("\\");return-1!==t?e.substr(0,t):""}return null},currentFiles(){return[...this.currentSortedFiles.filter(e=>"directory"===e.Type),...this.currentSortedFiles.filter(e=>"file"===e.Type)].slice(0,this.offset)},currentSortedFiles(){return"DESC"===this.sortOrder?this.currentUnsortedFiles.sort((e,t)=>e[this.sortBy]>t[this.sortBy]?-1:e[this.sortBy]<t[this.sortBy]?1:0):this.currentUnsortedFiles.sort((e,t)=>e[this.sortBy]>t[this.sortBy]?1:e[this.sortBy]<t[this.sortBy]?-1:0)},currentUnsortedFiles(){let e=this.$route.query.currentPath;if("string"===typeof e&&e.length>0){let t=e.match(/\\/g),s=t?t.length:0;return this.searchKeyword.length>0?this.$files.filter(t=>{let a=t.Path.match(/\\/g),r=a?a.length:0;return!!(0===t.Recycled&&t.Path.startsWith(e)&&t.Path.replace(e,"").includes(this.searchKeyword)&&r>s)}):this.$files.filter(t=>{let a=t.Path.match(/\\/g),r=a?a.length:0;return!(0!==t.Recycled||!t.Path.startsWith(e)||r!==s+1)})}return this.searchKeyword.length>0?this.$files.filter(e=>"."!==e.Path&&1!==e.Recycled&&!!e.Path.includes(this.searchKeyword)):this.$files.filter(e=>"."!==e.Path&&1!==e.Recycled&&!e.Path.includes("\\"))}},data(){return{offset:10,isSelectAll:!1,searchKeyword:"",sortBy:"UpdatedAt",sortOrder:"DESC",renameFileId:0,renameFileName:"",newFolderName:"",cutOptions:{canPaste:!1,srcId:0,srcPath:""},showNewFolderForm:!1}},mounted(){this.$startLoading(),this.$http.post("/api/files",{id:this.$user.id,token:this.$user.token}).then(e=>{e.data.success?this.$setFiles(e.data.data.files):this.$router.push("/login"),this.$stopLoading()}).catch(e=>{this.$router.push("/login"),this.$stopLoading()})}},p=m,$=(s("4fcd"),s("0c7c")),C=Object($["a"])(p,a,r,!1,null,"32fe3c2a",null);t["default"]=C.exports},f7a1:function(e,t,s){"use strict";s("b251")},fe9b:function(e,t,s){"use strict";var a=function(){var e=this,t=e.$createElement,s=e._self._c||t;return s("span",{staticClass:"btn-cut",attrs:{title:e.$locale.common.cut},on:{click:function(t){return e.$emit("click")}}},[s("i",{staticClass:"fas fa-cut fa-fw"})])},r=[],i=(s("a507"),s("0c7c")),n={},l=Object(i["a"])(n,a,r,!1,null,"533746cb",null);t["a"]=l.exports}}]);
//# sourceMappingURL=home.f8fcb215.js.map