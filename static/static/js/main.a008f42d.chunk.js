(this["webpackJsonpdropship-ui"]=this["webpackJsonpdropship-ui"]||[]).push([[0],{108:function(e,t,n){},125:function(e,t){},180:function(e,t,n){},181:function(e,t,n){},204:function(e,t,n){"use strict";n.r(t);var a=n(0),r=n.n(a),s=n(79),o=n.n(s),c=(n(88),n(89),n(32)),i=n(11),l=n.n(i),u=n(19),d=n(5),m=n(6),p=n(9),h=n(7),f=n(8),v=n(22),g=n.n(v),E=(n(108),n(18)),b=n.n(E),y=n(80),j=n.n(y),O=n(17);b.a.registerTheme({color:{primary:"#03a9f4",secondary:"#795548",white:"#FFFFFF",gray:"#CCCCCC",default:"#999999"},size:{xg:24,lg:18,md:14,sm:12,xs:10},lineHeight:{xg:"60px",lg:"54px",md:"36px",sm:"24px",xs:"18px"},unit:4}),b.a.registerInterface(j.a);var w=O.withStyles,P=function(e){Object(f.a)(n,e);var t=Object(h.a)(n);function n(){return Object(d.a)(this,n),t.apply(this,arguments)}return Object(m.a)(n,[{key:"render",value:function(){var e=this.props,t=e.children,n=e.onPress,a=e.disabled;return r.a.createElement("button",Object.assign({disabled:a,onClick:n},Object(O.css)({margin:"1px"})),t)}}]),n}(a.PureComponent);P.defaultProps={onPress:function(){}};var x=P,k=n(82),S=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"Pending";return function(t){var n=t.displayName,a=t.name,s=n||a;function o(n){var a=n.isPending,s=Object(k.a)(n,["isPending"]);return a?e:r.a.createElement(t,s)}return o.displayName="withLoading(".concat(s,")"),o}},D=n(23),N=function(e){Object(f.a)(n,e);var t=Object(h.a)(n);function n(){return Object(d.a)(this,n),t.apply(this,arguments)}return Object(m.a)(n,[{key:"render",value:function(){var e=this.props,t=e.styles,n=e.children;return r.a.createElement("div",Object(O.css)(t.overlay),r.a.createElement("div",Object(O.css)(t.wrapper),r.a.createElement("div",Object(O.css)(t.container),n)))}}]),n}(a.PureComponent),C=w((function(e){var t=e.color;return{overlay:{position:"fixed",zIndex:9999,top:0,left:0,width:"100%",height:"100%",backgroundColor:"rgba(0, 0, 0, .5)"},wrapper:{verticalAlign:"middle"},container:{margin:"40px auto 0px",padding:4*e.unit,backgroundColor:t.white,width:400}}}))(N),F=r.a.createContext({}),M=F.Provider,I=F.Consumer;var z,H=function(e){Object(f.a)(n,e);var t=Object(h.a)(n);function n(){return Object(d.a)(this,n),t.apply(this,arguments)}return Object(m.a)(n,[{key:"render",value:function(){var e=this.props,t=e.children,n=e.styles,a=e.large,s=e.xlarge,o=e.small,c=e.xsmall,i=e.primary,l=e.secondary;return r.a.createElement("span",Object(O.css)(n.default,c&&n.xsmall,o&&n.small,a&&n.large,s&&n.xlarge,l&&n.secondary,i&&n.primary),t)}}]),n}(a.PureComponent),_=w((function(e){var t=e.color,n=e.size;return{default:{color:t.default,fontSize:n.md},xlarge:{fontSize:n.xg},large:{fontSize:n.lg},small:{fontSize:n.sm},xsmall:{fontSize:n.xs},primary:{color:t.primary},secondary:{color:t.secondary}}}))(H);var T=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return function(t){Object(f.a)(a,t);var n=Object(h.a)(a);function a(e){var t;return Object(d.a)(this,a),(t=n.call(this,e)).state={showModal:!1},t.handleClose=t.handleClose.bind(Object(p.a)(t)),t.handleOpen=t.handleOpen.bind(Object(p.a)(t)),t}return Object(m.a)(a,[{key:"handleOpen",value:function(e,t){this.contentId=e,this.modalProps=t,this.setState({showModal:!0})}},{key:"handleClose",value:function(){this.setState({showModal:!1})}},{key:"render",value:function(){var t=this.props.children,n=this.state.showModal,a=e[this.contentId];return r.a.createElement(M,{value:{openModal:this.handleOpen,closeModal:this.handleClose}},t,n&&a&&r.a.createElement(C,null,r.a.createElement(a,this.modalProps)))}}]),a}(a.PureComponent)}((z={},Object(D.a)(z,"start_modal",(function(e){e.id;var t=e.project,n=e.startF;return r.a.createElement(I,null,(function(e){var a=e.closeModal;return r.a.createElement("div",null,r.a.createElement("div",null,r.a.createElement(_,null,t,"\uc744 \uc815\ub9d0\ub85c \uc2dc\uc791 \ud558\uc2dc\uaca0\uc2b5\ub2c8\uae4c?")),r.a.createElement(x,{primary:!0,onPress:function(){a(),n(t)}},"\uc608"),r.a.createElement(x,{onPress:a},"\ub2eb\uae30"))}))})),Object(D.a)(z,"stop_modal",(function(e){e.id;var t=e.project,n=e.stopF;return r.a.createElement(I,null,(function(e){var a=e.closeModal;return r.a.createElement("div",null,r.a.createElement("div",null,r.a.createElement(_,null,t,"\uc744 \uc815\ub9d0\ub85c \uc911\uc9c0 \ud558\uc2dc\uaca0\uc2b5\ub2c8\uae4c?")),r.a.createElement(x,{primary:!0,onPress:function(){a(),n(t)}},"\uc608"),r.a.createElement(x,{onPress:a},"\ub2eb\uae30"))}))})),Object(D.a)(z,"update_modal",(function(e){e.id;var t=e.project,n=e.updateF;return r.a.createElement(I,null,(function(e){var a=e.closeModal;return r.a.createElement("div",null,r.a.createElement("div",null,r.a.createElement(_,null,t,"\uc744 \uc815\ub9d0\ub85c \uc5c5\ub370\uc774\ud2b8 \ud558\uc2dc\uaca0\uc2b5\ub2c8\uae4c?")),r.a.createElement(x,{primary:!0,onPress:function(){a(),n(t)}},"\uc608"),r.a.createElement(x,{onPress:a},"\ub2eb\uae30"))}))})),z)),U=S(r.a.createElement(x,{disabled:!0},"pending"))(x);function A(e){var t=e.project,n=e.status,a=e.pid,s=void 0===a?"-":a,o=e.uptime,c=void 0===o?"-":o,i=e.startF,l=e.stopF,u=e.updateF,d=e.isPending;return r.a.createElement("div",{className:"status"},r.a.createElement("span",{className:"item"},t),r.a.createElement("span",{className:"item"},n),r.a.createElement("span",{className:"item"},"Down"===n?"-":s),r.a.createElement("span",{className:"item"},"Down"===n?"-":c),r.a.createElement("span",{className:"item"},r.a.createElement(T,null,r.a.createElement(I,null,(function(e){var a=e.openModal;return r.a.createElement(U,{isPending:d,onPress:"Down"===n?function(){return a("start_modal",{project:t,startF:i})}:function(){return a("stop_modal",{project:t,stopF:l})}},"Down"===n?"start":"stop")}))),r.a.createElement(T,null,r.a.createElement(I,null,(function(e){var n=e.openModal;return r.a.createElement(U,{isPending:d,onPress:function(){return n("update_modal",{project:t,updateF:u})}},"update")})))))}A.defaultProps={isPending:!1};var B=A,L=n(2),J=n.n(L),W=(n(180),S(r.a.createElement(x,{disabled:!0},"pending"))(x));function q(e){var t=e.refresh,n=e.nowDate,a=e.isPending;return r.a.createElement("span",{className:"info"},r.a.createElement("span",{className:"info-item"},n),r.a.createElement(W,{isPending:a,onPress:function(){return t()}},"refresh"))}q.propType={refresh:J.a.func.isRequired,nowDate:J.a.string};var R=q;var Y=function(e){var t=e.histories.map((function(e){var t=e.occuredDate,n=e.project,a=e.command,r=e.result;return t.toLocaleString()+" "+n+" "+a+" "+(r?"\uc131\uacf5":"\uc2e4\ud328")})).join("\n");return r.a.createElement("textarea",Object.assign({},Object(O.css)({width:"100%"},{height:"200px"},{overflowY:"scroll"}),{readOnly:!0,value:t}))},$=(n(181),n(21));var G=function(){return r.a.createElement("div",{className:"item-title"},r.a.createElement("span",{className:"item"},"project"),r.a.createElement("span",{className:"item"},"status"),r.a.createElement("span",{className:"item"},"pid"),r.a.createElement("span",{className:"item"},"uptime"),r.a.createElement("span",{className:"item"},"control"))};var K=function(e){var t=e.dep;return r.a.createElement("span",null,"\uc2e4\ud589\uc21c\uc11c : ",t)},Q=function(e){Object(f.a)(n,e);var t=Object(h.a)(n);function n(e){var a;return Object(d.a)(this,n),(a=t.call(this,e)).maxHistory=50,a.getDependency=Object(u.a)(l.a.mark((function e(){var t,n;return l.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,g.a.get(a.basicUrl+"/dependency");case 2:t=e.sent,n=t.data.dependency,console.log(n),a.setState({dep:n});case 6:case"end":return e.stop()}}),e)}))),a.getSummary=Object(u.a)(l.a.mark((function e(){var t,n;return l.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return a.doPending(),e.next=3,g.a.get(a.basicUrl+"/summary");case 3:t=e.sent,n=t.data.projects,a.setState({projects:n,nowDate:new Date}),a.doIdle();case 7:case"end":return e.stop()}}),e)}))),a.refresh=Object(u.a)(l.a.mark((function e(){return l.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:a.getSummary().finally((function(){return a.getDependency()}));case 1:case"end":return e.stop()}}),e)}))),a.start=a.controlEndpoint("start"),a.stop=a.controlEndpoint("stop"),a.update=a.controlEndpoint("update"),a.state={isPending:!1,projects:[],nowDate:new Date,histories:[],dep:""},a.getSummary=a.getSummary.bind(Object(p.a)(a)),a.start=a.start.bind(Object(p.a)(a)),a.stop=a.stop.bind(Object(p.a)(a)),a.update=a.update.bind(Object(p.a)(a)),a.basicUrl=window.location.protocol+"//"+window.location.hostname+":3001/dropship/api/v1",a}return Object(m.a)(n,[{key:"pushHistory",value:function(e){var t,n;return this.state.histories.length===this.maxHistory?(t=[]).concat.apply(t,[e].concat(Object(c.a)(this.state.histories.slice(0,this.maxHistory-1)))):(n=[]).concat.apply(n,[e].concat(Object(c.a)(this.state.histories)))}},{key:"componentDidMount",value:function(){this.getDependency(),this.getSummary()}},{key:"doPending",value:function(){this.setState({isPending:!0})}},{key:"doIdle",value:function(){this.setState({isPending:!1})}},{key:"controlEndpoint",value:function(e){var t=this;return function(){var n=Object(u.a)(l.a.mark((function n(a){return l.a.wrap((function(n){for(;;)switch(n.prev=n.next){case 0:return t.doPending(),n.next=3,Object($.trackPromise)(g.a.post(t.basicUrl+"/"+a+"/"+e).then((function(){var n=t.pushHistory({occuredDate:new Date,project:a,command:e,result:!0});t.setState({histories:n}),t.getSummary()})).catch((function(n){var r=t.pushHistory({occuredDate:new Date,project:a,command:e,result:!1});t.setState({histories:r})})).finally((function(){t.getDependency().finally((function(){return t.doIdle()}))})));case 3:case"end":return n.stop()}}),n)})));return function(e){return n.apply(this,arguments)}}()}},{key:"render",value:function(){var e=this,t=this.state,n=t.projects,a=t.nowDate,s=t.isPending;return r.a.createElement("section",null,r.a.createElement("div",{className:"container"},r.a.createElement("h1",{className:"manager"},"SmartSee Server Manager"),r.a.createElement("div",{className:"refresh-bar"},r.a.createElement(K,{dep:this.state.dep}),r.a.createElement(R,{refresh:this.refresh,nowDate:a.toLocaleString(),isPending:s})),r.a.createElement(r.a.Fragment,null,r.a.createElement("div",{className:"status-box"},r.a.createElement(G,null),r.a.createElement("div",{className:"items"},n.sort((function(e,t){return e.project.localeCompare(t.project)})).map((function(t){return r.a.createElement(B,{className:"item",key:t.project,project:t.project,status:t.status,uptime:t.uptime,pid:t.pid,startF:e.start,stopF:e.stop,updateF:e.update,isPending:s})}))),r.a.createElement(Y,{histories:this.state.histories})))))}}]),n}(a.Component),V=n(81),X=n.n(V),Z=function(e){return Object($.usePromiseTracker)().promiseInProgress&&r.a.createElement("div",{style:{width:"100%",height:"100",position:"absolute",left:0,top:20}},r.a.createElement(X.a,{type:"ThreeDots",color:"#2BAD60",height:"100",width:"100"}))};var ee=function(){return r.a.createElement("div",{className:"App"},r.a.createElement("nav",null,r.a.createElement(Q,null),r.a.createElement(Z,null)))};Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));o.a.render(r.a.createElement(ee,null),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then((function(e){e.unregister()})).catch((function(e){console.error(e.message)}))},83:function(e,t,n){e.exports=n(204)},88:function(e,t,n){},89:function(e,t,n){}},[[83,1,2]]]);
//# sourceMappingURL=main.a008f42d.chunk.js.map