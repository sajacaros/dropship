(this["webpackJsonpdropship-ui"]=this["webpackJsonpdropship-ui"]||[]).push([[0],{16:function(e,t,a){},22:function(e,t,a){e.exports=a(49)},27:function(e,t,a){},28:function(e,t,a){},49:function(e,t,a){"use strict";a.r(t);var n=a(0),r=a.n(n),s=a(17),c=a.n(s),o=(a(27),a(28),a(3)),i=a.n(o),u=a(5),m=a(18),p=a(19),l=a(4),d=a(20),f=a(21),h=a(6),v=a.n(h),b=a(2),E=a.n(b);a(16);function g(e){var t=e.project,a=e.status,n=e.pid,s=e.uptime,c=e.startF,o=e.stopF,i=e.updateF;return r.a.createElement("div",{className:"status"},r.a.createElement("span",{className:"item"},t),r.a.createElement("span",{className:"item"},a),"Down"===a?r.a.createElement(r.a.Fragment,null,r.a.createElement("span",{className:"item"},"-"),r.a.createElement("span",{className:"item"},"-"),r.a.createElement("button",{className:"item",onClick:function(){return c(t)}},"start")):r.a.createElement(r.a.Fragment,null,r.a.createElement("span",{className:"item"},n),r.a.createElement("span",{className:"item"},s),r.a.createElement("button",{className:"item",onClick:function(){return o(t)}},"stop")),r.a.createElement("button",{className:"item",onClick:function(){return i(t)}},"update"))}g.prototype={project:E.a.string.isRequired,status:E.a.string.isRequired,uptime:E.a.string,pid:E.a.number,startF:E.a.func.isRequired,stopF:E.a.func.isRequired,updateF:E.a.func.isRequired};var j=g;function w(e){var t=e.refresh;return r.a.createElement("div",null,r.a.createElement("button",{onClick:function(){return t()}},"refresh"))}w.prototype={refresh:E.a.func.isRequired};var y=w,k=function(e){Object(f.a)(a,e);var t=Object(d.a)(a);function a(e){var n;return Object(m.a)(this,a),(n=t.call(this,e)).getSummary=Object(u.a)(i.a.mark((function e(){var t,a;return i.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,v.a.get(n.basicUrl+"/summary");case 2:t=e.sent,a=t.data.projects,n.setState({projects:a,isLoading:!1});case 5:case"end":return e.stop()}}),e)}))),n.start=function(){var e=Object(u.a)(i.a.mark((function e(t){return i.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,v.a.post(n.basicUrl+t+"/start").then((function(){return n.getSummary()}));case 2:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),n.stop=function(){var e=Object(u.a)(i.a.mark((function e(t){return i.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,v.a.post(n.basicUrl+t+"/stop").then((function(){return n.getSummary()}));case 2:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),n.update=function(){var e=Object(u.a)(i.a.mark((function e(t){return i.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,v.a.post(n.basicUrl+t+"/update").then((function(){return n.getSummary()}));case 2:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),n.state={isLoading:!0,projects:[]},n.getSummary=n.getSummary.bind(Object(l.a)(n)),n.start=n.start.bind(Object(l.a)(n)),n.stop=n.stop.bind(Object(l.a)(n)),n.update=n.update.bind(Object(l.a)(n)),console.log("not dev"),n.basicUrl=window.location.protocol+"//"+window.location.hostname+":3001/dropship/api/v1",n}return Object(p.a)(a,[{key:"componentDidMount",value:function(){this.getSummary()}},{key:"render",value:function(){var e=this,t=this.state,a=t.projects,n=t.isLoading;return r.a.createElement("section",{className:"container"},r.a.createElement(y,{refresh:this.getSummary}),n?r.a.createElement("div",{className:"loader"},r.a.createElement("span",{className:"loader_text"},"'Loading...'")):r.a.createElement("div",{className:"items"},a.sort((function(e,t){return e.project.localeCompare(t.project)})).map((function(t){return r.a.createElement(j,{key:t.project,project:t.project,status:t.status,uptime:t.uptime,pid:t.pid,startF:e.start,stopF:e.stop,updateF:e.update})}))))}}]),a}(n.Component);var N=function(){return r.a.createElement("div",{className:"App"},r.a.createElement("nav",null,r.a.createElement(k,null)))};Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));c.a.render(r.a.createElement(N,null),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then((function(e){e.unregister()})).catch((function(e){console.error(e.message)}))}},[[22,1,2]]]);
//# sourceMappingURL=main.36810a04.chunk.js.map