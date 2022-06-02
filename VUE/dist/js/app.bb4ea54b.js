(function(){"use strict";var t={709:function(t,e,n){var r=n(8935),i=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{attrs:{id:"app"}},[n("router-view")],1)},o=[],a=n(1001),u={},s=(0,a.Z)(u,i,o,!1,null,null,null),c=s.exports,l=n(2809),p=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"home-container"},[r("img",{staticClass:"home-img",attrs:{src:n(1125),alt:""}}),r("SearchInput")],1)},d=[],f={methods:{}},m=f,h=(0,a.Z)(m,p,d,!1,null,"6d2d928c",null),g=h.exports,v=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"result-container"},[n("SearchInput"),n("div",{staticClass:"res-content"},t._l(t.documents,(function(t){return n("Tip",{key:t.id,attrs:{imgsrc:t.document.Url,textcontent:t.text}})})),1),n("div",{staticClass:"pagination"},[n("el-pagination",{ref:"pagination",attrs:{background:"",layout:"prev, pager, next",total:t.total,"current-page":t.currentPage},on:{"current-change":t.currentChange,"prev-click":t.prevClick,"next-click":t.nextClick}})],1)],1)},y=[],k=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"tip-container"},[n("div",{staticClass:"tip-img"},[n("img",{attrs:{src:t.imgsrc,alt:""}})]),n("div",{staticClass:"text"},[t._v(t._s(t.textcontent))])])},b=[],w={props:["imgsrc","textcontent"]},C=w,x=(0,a.Z)(C,k,b,!1,null,"388f9dc1",null),_=x.exports,S=n(4665),E={data(){return{currentPage:this.page}},methods:{currentChange(){const t=this.$refs.pagination.internalCurrentPage;this.$store.dispatch("pageJump",{targetPage:t,input:this.keyword})},prevClick(){const t=this.page-1;this.$store.dispatch("pageJump",{targetPage:t,input:this.keyword})},nextClick(){const t=this.page+1;this.$store.dispatch("pageJump",{targetPage:t,input:this.keyword})}},computed:{...(0,S.rn)(["documents","pageCount","total","page","keyword"])},mounted(){this.keyword=this.$route.query.input},components:{Tip:_}},T=E,P=(0,a.Z)(T,v,y,!1,null,"645e0fa2",null),O=P.exports;r["default"].use(l.Z);const $=[{path:"/",redirect:"/home"},{path:"/home",component:g,meta:{title:"零度搜索"}},{path:"/result",component:O,meta:{title:"零度搜索"}}],j=new l.Z({routes:$});j.beforeEach(((t,e,n)=>{t.meta.title&&(document.title=t.meta.title),n()}));var Z=j,U=(n(1703),n(6166)),I=n.n(U),R=n(9879),q=n.n(R);const J=I().create({baseURL:"http://leiyu.icu:5678/api",timeout:5e3});J.interceptors.request.use((t=>(q().start(),t))),J.interceptors.response.use((t=>(q().done(),t.data)),(t=>Promise.reject(new Error("faile"))));const L=(t,e=1)=>{const n={query:t,page:e,limit:10,order:"desc"},r=JSON.stringify(n);return J({url:"/query",method:"post",data:r})};r["default"].use(S.ZP);var G=new S.ZP.Store({state:{keyword:"",documents:[],limit:0,page:0,total:0,pageCount:0,time:0},getters:{},mutations:{SETINPUT(t,e){t.keyword=e},GETRESULT(t,e){t.documents=e.documents,t.limit=e.limit,t.page=e.page,t.total=e.total,t.pageCount=e.pageCount,t.time=e.time}},actions:{async getResult(t,e){const n=await L(e);(n.state="true")&&t.commit("GETRESULT",n.data)},async pageJump(t,e){const n=await L(e.input,e.targetPage);(n.state="true")&&t.commit("GETRESULT",n.data)}},modules:{}}),M=n(4549),N=n.n(M),F=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"searchInput-container",on:{keyup:function(e){return!e.type.indexOf("key")&&t._k(e.keyCode,"enter",13,e.key,"Enter")?null:t.goSearch(t.input)}}},[n("el-input",{attrs:{placeholder:"请输入内容",clearable:""},model:{value:t.input,callback:function(e){t.input=e},expression:"input"}}),n("el-button",{attrs:{type:"primary"},on:{click:function(e){return t.goSearch(t.input)}}},[t._v("搜索")])],1)},z=[],A={data(){return{input:""}},methods:{goSearch(t){this.$store.commit("SETINPUT",t),this.$store.dispatch("getResult",t),"/home"===this.$route.path&&this.$router.push({path:"result",query:{input:this.input}})}},computed:{...(0,S.rn)({keyword:t=>t.keyword})},mounted(){this.input=this.keyword}},B=A,D=(0,a.Z)(B,F,z,!1,null,"1cceb747",null),H=D.exports;r["default"].config.productionTip=!1,r["default"].use(N()),r["default"].component("SearchInput",H),new r["default"]({router:Z,store:G,render:t=>t(c)}).$mount("#app")},1125:function(t,e,n){t.exports=n.p+"img/leiyu.5f127551.png"}},e={};function n(r){var i=e[r];if(void 0!==i)return i.exports;var o=e[r]={exports:{}};return t[r].call(o.exports,o,o.exports,n),o.exports}n.m=t,function(){var t=[];n.O=function(e,r,i,o){if(!r){var a=1/0;for(l=0;l<t.length;l++){r=t[l][0],i=t[l][1],o=t[l][2];for(var u=!0,s=0;s<r.length;s++)(!1&o||a>=o)&&Object.keys(n.O).every((function(t){return n.O[t](r[s])}))?r.splice(s--,1):(u=!1,o<a&&(a=o));if(u){t.splice(l--,1);var c=i();void 0!==c&&(e=c)}}return e}o=o||0;for(var l=t.length;l>0&&t[l-1][2]>o;l--)t[l]=t[l-1];t[l]=[r,i,o]}}(),function(){n.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return n.d(e,{a:e}),e}}(),function(){n.d=function(t,e){for(var r in e)n.o(e,r)&&!n.o(t,r)&&Object.defineProperty(t,r,{enumerable:!0,get:e[r]})}}(),function(){n.g=function(){if("object"===typeof globalThis)return globalThis;try{return this||new Function("return this")()}catch(t){if("object"===typeof window)return window}}()}(),function(){n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)}}(),function(){n.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})}}(),function(){n.p="/"}(),function(){var t={143:0};n.O.j=function(e){return 0===t[e]};var e=function(e,r){var i,o,a=r[0],u=r[1],s=r[2],c=0;if(a.some((function(e){return 0!==t[e]}))){for(i in u)n.o(u,i)&&(n.m[i]=u[i]);if(s)var l=s(n)}for(e&&e(r);c<a.length;c++)o=a[c],n.o(t,o)&&t[o]&&t[o][0](),t[o]=0;return n.O(l)},r=self["webpackChunkleiyu"]=self["webpackChunkleiyu"]||[];r.forEach(e.bind(null,0)),r.push=e.bind(null,r.push.bind(r))}();var r=n.O(void 0,[998],(function(){return n(709)}));r=n.O(r)})();
//# sourceMappingURL=app.bb4ea54b.js.map