webpackJsonp([1],{0:function(t,e){},"9pe8":function(t,e){},NHnr:function(t,e,o){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var a=o("7+uW"),l={render:function(){var t=this,e=t.$createElement,o=t._self._c||e;return o("div",{attrs:{id:"app"}},[o("el-menu",{staticClass:"el-menu-demo",attrs:{mode:"horizontal","background-color":"#545c64","text-color":"#fff","active-text-color":"#ffd04b"}},[o("el-menu-item",{attrs:{index:"1"}},[o("router-link",{attrs:{to:"/"}},[t._v("Topic查阅")])],1),t._v(" "),o("el-menu-item",{attrs:{index:"2"}},[o("router-link",{attrs:{to:"/consumeGroup"}},[t._v("ConsumerGroup查阅")])],1),t._v(" "),o("el-menu-item",{attrs:{index:"3"}},[o("router-link",{attrs:{to:"/consumeTemplate"}},[t._v("消费者模板")])],1),t._v(" "),o("el-menu-item",{attrs:{index:"4"}},[o("router-link",{attrs:{to:"/consumer"}},[t._v("消费者")])],1),t._v(" "),o("el-menu-item",{attrs:{index:"5"}},[t._v("监控")])],1),t._v(" "),o("router-view")],1)},staticRenderFns:[]};var n=o("VU/8")({name:"App",data:function(){return{}},methods:{}},l,!1,function(t){o("9pe8")},null,null).exports,i=o("/ocq"),s={name:"Index",data:function(){return{tableData:[]}},methods:{loadTopics:function(){var t=this;this.$http.get(this.GLOBAL.httpAddress+"/api/getTopics").then(function(e){0===e.data.code?t.tableData=e.data.data:alert(e.data.message)},function(){console.log("请求失败处理")})}},mounted:function(){this.loadTopics()}},r={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",[[e("el-table",{staticStyle:{width:"100%"},attrs:{data:this.tableData}},[e("el-table-column",{attrs:{prop:"name",label:"Topic名称",width:"1000px"}})],1)]],2)},staticRenderFns:[]};var c=o("VU/8")(s,r,!1,function(t){o("aiE5")},null,null).exports,u={data:function(){return{tableData:[]}},methods:{loadTopics:function(){var t=this;this.$http.get(this.GLOBAL.httpAddress+"/api/getConsumerGroups").then(function(e){0===e.data.code?t.tableData=e.data.data:alert(e.data.message)},function(){console.log("请求失败处理")})}},mounted:function(){this.loadTopics()}},d={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",[[e("el-table",{staticStyle:{width:"100%"},attrs:{data:this.tableData}},[e("el-table-column",{attrs:{prop:"name",label:"ConsumeGroup名称",width:"1000px"}})],1)]],2)},staticRenderFns:[]};var m=o("VU/8")(u,d,!1,function(t){o("WEIK")},null,null).exports,p={data:function(){return{tableData:[],dialogFormVisible:!1,dialogConsumerVisible:!1,dialogConsumerListVisible:!1,form:{topic_name:"",callback_url:""},formTwo:{consumer_group:"",topic_callback_id:"",topic_name:"",callback_url:""},options:[],consumerList:[],consumeGroupsOptions:[],formLabelWidth:"120px"}},methods:{loadConsumeTemplates:function(){var t=this;this.$http.get(this.GLOBAL.httpAddress+"/api/getConsumerTemplates").then(function(e){0===e.data.code?t.tableData=e.data.data:alert(e.data.message)},function(){console.log("请求失败处理")}),this.$http.get(this.GLOBAL.httpAddress+"/api/getTopics").then(function(e){0===e.data.code?t.options=e.data.data:alert(e.data.message)},function(){console.log("请求失败处理")}),this.$http.get(this.GLOBAL.httpAddress+"/api/getConsumerGroups").then(function(e){0===e.data.code?t.consumeGroupsOptions=e.data.data:alert(e.data.message)},function(){console.log("请求失败处理")})},createTopicCallbackTemplate:function(){this.dialogFormVisible=!1;this.$http.post(this.GLOBAL.httpAddress+"/api/createConsumerTemplate",this.form,{emulateJSON:!0}).then(function(t){0===t.data.code?(alert("成功"),this.loadConsumeTemplates()):alert(t.data.message)},function(t){})},NewConsumer:function(t){this.formTwo.topic_callback_id=t.id,this.formTwo.callback_url=t.callback_url,this.formTwo.topic_name=t.topic,this.dialogConsumerVisible=!0},GetConsumerList:function(t){var e=this;this.$http.get(this.GLOBAL.httpAddress+"/api/getAllConsumersOfTemplate?template_id="+t.id).then(function(t){0===t.data.code?e.consumerList=t.data.data:alert(t.data.message)},function(){console.log("请求失败处理")}),e.dialogConsumerListVisible=!0},DeleteTemplate:function(t){var e=this;this.$http.get(this.GLOBAL.httpAddress+"/api/deleteTemplate?template_id="+t.id).then(function(t){0===t.data.code?(alert("成功删除"),e.loadConsumeTemplates()):alert(t.data.message)},function(){console.log("请求失败处理")})},startConsumer:function(){this.dialogConsumerVisible=!1;this.$http.post(this.GLOBAL.httpAddress+"/api/startConsumer",this.formTwo,{emulateJSON:!0}).then(function(t){0===t.body.code?alert("成功"):alert(t.body.message)},function(t){}),this.formTwo.consumer_group=""}},mounted:function(){this.loadConsumeTemplates()}},f={render:function(){var t=this,e=t.$createElement,o=t._self._c||e;return o("div",[[o("el-tag",[t._v("主题回调模板")]),t._v(" "),o("el-row",{staticStyle:{float:"right"}},[o("el-button",{attrs:{type:"primary",size:"middle"},on:{click:function(e){t.dialogFormVisible=!0}}},[t._v("新建模板")])],1),t._v(" "),o("el-dialog",{attrs:{title:"新建模板",visible:t.dialogFormVisible},on:{"update:visible":function(e){t.dialogFormVisible=e}}},[o("el-form",{attrs:{model:t.form}},[o("el-form-item",{attrs:{label:"topic名称","label-width":t.formLabelWidth}},[[o("el-select",{attrs:{placeholder:"请选择",filterable:"",clearable:"","allow-create":""},model:{value:t.form.topic_name,callback:function(e){t.$set(t.form,"topic_name",e)},expression:"form.topic_name"}},t._l(t.options,function(t){return o("el-option",{key:t.name,attrs:{label:t.name,value:t.name}})}))]],2),t._v(" "),o("el-form-item",{attrs:{label:"回调http地址","label-width":t.formLabelWidth}},[o("el-input",{attrs:{autocomplete:"off"},model:{value:t.form.callback_url,callback:function(e){t.$set(t.form,"callback_url",e)},expression:"form.callback_url"}})],1)],1),t._v(" "),o("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[o("el-button",{on:{click:function(e){t.dialogFormVisible=!1}}},[t._v("取 消")]),t._v(" "),o("el-button",{attrs:{type:"primary"},on:{click:t.createTopicCallbackTemplate}},[t._v("确 定")])],1)],1),t._v(" "),o("el-table",{staticStyle:{width:"100%"},attrs:{data:t.tableData}},[o("el-table-column",{attrs:{prop:"id",label:"模板id",width:"120"}}),t._v(" "),o("el-table-column",{attrs:{prop:"topic",label:"Topic名称",width:"360"}}),t._v(" "),o("el-table-column",{attrs:{prop:"callback_url",label:"回调地址",width:"360"}}),t._v(" "),o("el-table-column",{scopedSlots:t._u([{key:"default",fn:function(e){return[o("el-button",{attrs:{type:"danger",size:"mini"},on:{click:function(o){t.NewConsumer(e.row)}}},[t._v("新开消费者\n        ")])]}}])}),t._v(" "),o("el-table-column",{scopedSlots:t._u([{key:"default",fn:function(e){return[o("el-button",{attrs:{type:"danger",size:"mini"},on:{click:function(o){t.GetConsumerList(e.row)}}},[t._v("查看消费者列表\n          ")])]}}])}),t._v(" "),o("el-table-column",{scopedSlots:t._u([{key:"default",fn:function(e){return[o("el-button",{attrs:{type:"danger",size:"mini"},on:{click:function(o){t.DeleteTemplate(e.row)}}},[t._v("删除\n          ")])]}}])})],1),t._v(" "),o("el-dialog",{attrs:{title:"新建消费者协程",visible:t.dialogConsumerVisible},on:{"update:visible":function(e){t.dialogConsumerVisible=e}}},[o("el-form",{attrs:{model:t.formTwo}},[o("el-form-item",{attrs:{label:"消费者组","label-width":t.formLabelWidth}},[[o("el-select",{attrs:{placeholder:"请选择",filterable:"",clearable:"","allow-create":""},model:{value:t.formTwo.consumer_group,callback:function(e){t.$set(t.formTwo,"consumer_group",e)},expression:"formTwo.consumer_group"}},t._l(t.consumeGroupsOptions,function(t){return o("el-option",{key:t.name,attrs:{label:t.name,value:t.name}})}))]],2),t._v(" "),o("el-form-item",{attrs:{label:"topic名称","label-width":t.formLabelWidth}},[o("el-input",{attrs:{autocomplete:"off",disabled:""},model:{value:t.formTwo.topic_name,callback:function(e){t.$set(t.formTwo,"topic_name",e)},expression:"formTwo.topic_name"}})],1),t._v(" "),o("el-form-item",{attrs:{label:"回调http地址","label-width":t.formLabelWidth}},[o("el-input",{attrs:{autocomplete:"off",disabled:""},model:{value:t.formTwo.callback_url,callback:function(e){t.$set(t.formTwo,"callback_url",e)},expression:"formTwo.callback_url"}})],1)],1),t._v(" "),o("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[o("el-button",{on:{click:function(e){t.dialogConsumerVisible=!1}}},[t._v("取 消")]),t._v(" "),o("el-button",{attrs:{type:"primary"},on:{click:t.startConsumer}},[t._v("确 定")])],1)],1),t._v(" "),o("el-dialog",{attrs:{title:"消费者列表",visible:t.dialogConsumerListVisible},on:{"update:visible":function(e){t.dialogConsumerListVisible=e}}},[o("el-table",{attrs:{data:t.consumerList}},[o("el-table-column",{attrs:{property:"consumer_id",label:"consumer_id",width:"300"}}),t._v(" "),o("el-table-column",{attrs:{property:"consumer_group",label:"consumer_group",width:"200"}})],1)],1)]],2)},staticRenderFns:[]};var h=o("VU/8")(p,f,!1,function(t){o("Xjc2")},null,null).exports,b={name:"Index",data:function(){return{tableData:[],form:{consumer_id:""}}},methods:{loadTopics:function(){var t=this;this.$http.get(this.GLOBAL.httpAddress+"/api/getConsumers").then(function(e){0===e.data.code?t.tableData=e.data.data:alert(e.data.message)},function(){console.log("请求失败处理")})},closeConsumer:function(t){this.form.consumer_id=t;this.$http.post(this.GLOBAL.httpAddress+"/api/closeConsumer",this.form,{emulateJSON:!0}).then(function(t){0===t.data.code?(alert("成功"),this.loadTopics()):alert(t.data.message)},function(t){}),this.form.consumer_id=""}},mounted:function(){this.loadTopics()}},_={render:function(){var t=this,e=t.$createElement,o=t._self._c||e;return o("div",[[o("el-table",{staticStyle:{width:"100%"},attrs:{data:t.tableData}},[o("el-table-column",{attrs:{prop:"ConsumerId",label:"ConsumerId",width:"280px"}}),t._v(" "),o("el-table-column",{attrs:{label:"状态",width:"50px"},scopedSlots:t._u([{key:"default",fn:function(t){return[o("el-row",[t.row.Running?o("el-button",{attrs:{type:"success",icon:"el-icon-check",circle:""}}):o("el-button",{attrs:{type:"danger",icon:"el-icon-check",circle:""}})],1)]}}])}),t._v(" "),o("el-table-column",{attrs:{prop:"Topic",label:"Topic名称",width:"240px"}}),t._v(" "),o("el-table-column",{attrs:{prop:"ConsumerGroup",label:"ConsumerGroup",width:"300px"}}),t._v(" "),o("el-table-column",{attrs:{prop:"Callback",label:"Callback地址",width:"240px"}}),t._v(" "),o("el-table-column",{attrs:{label:"操作"},scopedSlots:t._u([{key:"default",fn:function(e){return[o("el-button",{attrs:{size:"mini",type:"danger"},on:{click:function(o){t.closeConsumer(e.row.ConsumerId)}}},[t._v("关闭并删除")])]}}])})],1)]],2)},staticRenderFns:[]};var v=o("VU/8")(b,_,!1,function(t){o("SazK")},null,null).exports;a.default.use(i.a);var g=new i.a({routes:[{path:"/",name:"Index",component:c},{path:"/consumeGroup",name:"ConsumeGroup",component:m},{path:"/consumeTemplate",name:"ConsumeTemplate",component:h},{path:"/consumer",name:"Consumer",component:v}]}),w=o("zL8q"),T=o.n(w),C=(o("tvR6"),{httpAddress:window.location.host}),k=o("VU/8")(C,null,!1,null,null,null).exports,L=o("8+8L");a.default.use(T.a),a.default.use(L.a),a.default.config.productionTip=!1,a.default.prototype.GLOBAL=k,new a.default({el:"#app",router:g,components:{App:n},template:"<App/>"})},SazK:function(t,e){},WEIK:function(t,e){},Xjc2:function(t,e){},aiE5:function(t,e){},tvR6:function(t,e){}},["NHnr"]);
//# sourceMappingURL=app.f06e3a16d31c7c5fb51b.js.map