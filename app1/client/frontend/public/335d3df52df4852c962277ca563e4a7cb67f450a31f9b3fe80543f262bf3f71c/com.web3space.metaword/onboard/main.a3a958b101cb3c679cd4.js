"use strict";(self.webpackChunkspinner=self.webpackChunkspinner||[]).push([[179],{2861:function(e,t,n){var r,i,o=n(2322),a=n(2784),s=n(7029),c=Object.freeze({INFO:"info",EXIT:"exit",PERMISSION:"permission"}),l=(r=function(e,t){return r=Object.setPrototypeOf||{__proto__:[]}instanceof Array&&function(e,t){e.__proto__=t}||function(e,t){for(var n in t)Object.prototype.hasOwnProperty.call(t,n)&&(e[n]=t[n])},r(e,t)},function(e,t){if("function"!=typeof t&&null!==t)throw new TypeError("Class extends value "+String(t)+" is not a constructor or null");function n(){this.constructor=e}r(e,t),e.prototype=null===t?Object.create(t):(n.prototype=t.prototype,new n)}),u=function(e){function t(){var t=e.call(this,"System is not ready")||this;return t.name="SYSTEM_NOT_READY",t}return l(t,e),t}(Error),f=(i=Error,l((function(){var e=i.call(this,"Message type is not allowed")||this;return e.name="NOT_ALLOWED_TYPE",e}),i),function(e){function t(){var t=e.call(this,"Access is denied")||this;return t.name="REQUEST_PERMISSION_FAILED",t}return l(t,e),t}(Error)),p=function(){function e(){this._listeners={},this.on=this.on.bind(this),this.removeEventListener=this.removeEventListener.bind(this),this.emit=this.emit.bind(this)}return e.prototype.on=function(e,t){return this._listeners[e]||(this._listeners[e]=[]),this._listeners[e].push(t),this},e.prototype.removeEventListener=function(e,t){this._listeners[e]&&(this._listeners[e]=this._listeners[e].filter((function(e){return e!==t})))},e.prototype.removeAllEventListeners=function(e){this._listeners[e]&&delete this._listeners[e]},e.prototype.emit=function(e){for(var t=[],n=1;n<arguments.length;n++)t[n-1]=arguments[n];var r=this._listeners[e];r&&r.forEach((function(e){return e.apply(void 0,t)}))},e}(),d=p,h=function(){var e=function(t,n){return e=Object.setPrototypeOf||{__proto__:[]}instanceof Array&&function(e,t){e.__proto__=t}||function(e,t){for(var n in t)Object.prototype.hasOwnProperty.call(t,n)&&(e[n]=t[n])},e(t,n)};return function(t,n){if("function"!=typeof n&&null!==n)throw new TypeError("Class extends value "+String(n)+" is not a constructor or null");function r(){this.constructor=t}e(t,n),t.prototype=null===n?Object.create(n):(r.prototype=n.prototype,new r)}}(),m=function(e,t,n,r){return new(n||(n=Promise))((function(i,o){function a(e){try{c(r.next(e))}catch(e){o(e)}}function s(e){try{c(r.throw(e))}catch(e){o(e)}}function c(e){var t;e.done?i(e.value):(t=e.value,t instanceof n?t:new n((function(e){e(t)}))).then(a,s)}c((r=r.apply(e,t||[])).next())}))},g=function(e,t){var n,r,i,o,a={label:0,sent:function(){if(1&i[0])throw i[1];return i[1]},trys:[],ops:[]};return o={next:s(0),throw:s(1),return:s(2)},"function"==typeof Symbol&&(o[Symbol.iterator]=function(){return this}),o;function s(s){return function(c){return function(s){if(n)throw new TypeError("Generator is already executing.");for(;o&&(o=0,s[0]&&(a=0)),a;)try{if(n=1,r&&(i=2&s[0]?r.return:s[0]?r.throw||((i=r.return)&&i.call(r),0):r.next)&&!(i=i.call(r,s[1])).done)return i;switch(r=0,i&&(s=[2&s[0],i.value]),s[0]){case 0:case 1:i=s;break;case 4:return a.label++,{value:s[1],done:!1};case 5:a.label++,r=s[1],s=[0];continue;case 7:s=a.ops.pop(),a.trys.pop();continue;default:if(!((i=(i=a.trys).length>0&&i[i.length-1])||6!==s[0]&&2!==s[0])){a=0;continue}if(3===s[0]&&(!i||s[1]>i[0]&&s[1]<i[3])){a.label=s[1];break}if(6===s[0]&&a.label<i[1]){a.label=i[1],i=s;break}if(i&&a.label<i[2]){a.label=i[2],a.ops.push(s);break}i[2]&&a.ops.pop(),a.trys.pop();continue}s=t.call(e,a)}catch(e){s=[6,e],r=0}finally{n=i=0}if(5&s[0])throw s[1];return{value:s[0]?s[1]:void 0,done:!0}}([s,c])}}},b={ready:-1,"get-person-info":-1,"get-wallet":-1,"get-view-on-boarding":-1,"check-performance":-1,"get-raw-seed":-1,"get-raw-seed-confirm":-1,"create-wallet":-1,"get-node":-1,"connect-node":-1,"scan-qr":-1,"get-file-zip":-1,"get-balance-wallet":-1,"unzip-process":-1,"unzip-file":-1,"read-abi-file":-1,"get-all-smart-contract":-1,"watch-approve":-1,"send-transaction":-1,"take-picture":-1,"get-transaction":-1,"read-abi-string":-1,"get-white-list-pagination":-1,"get-my-setting":-1,"set-status-notification":-1,"backup-data":-1,"on-zip-process":-1,"on-zip-result":-1,"sync-data-to-watch":-1,"insert-white-list":-1,"check-white-list-is-exist":-1,"copy-clipboard":-1,"capture-screen":-1,"get-qr-from-image":-1,"get-qr-from-camera":-1,"get-from-clipboard":-1,"select-image":-1,"check-amount":-1,"set-pin-code":-1,"check-pin-code":-1,"request-face-touch-id":-1,"on-off-confirm-watch":-1,tel:-1,"register-bottom":-1,"register-bottom-off":-1,"on-click-submit-button-bottom":-1,"select-date":-1,"spinner-register-data":-1,"event-onError":-1,"spinner-close":-1,"open-spinner":-1,"on-spinner-selected":-1,"on-listen-info":-1,"spinner-selected":-1,"searching-on-spinner":-1,"update-d-app":-1,"on-delete-d-app":-1,"delete-d-app":-1,"on-unzip-process":-1,"on-download-process":-1,"on-start-download-app":-1,"on-start-verify-app":-1,"checking-sign-app":-1,"has-device-touch":-1,"get-app-info-from-url":-1,"on-zip-start":-1,"on-receive-message":-1,"get-backup-files":-1,"share-item":-1,"restore-by-file":-1,"unzip-fail":-1,"deploy-d-app":-1,"deploy-sc":-1,"open-dapp":-1,"get-all-d-app":-1,"start-server":-1,"get-status-connected":-1,"get-wallet-info":-1,"open-server-socket":-1,"connect-to-server-socket":-1,"send-file":-1,"check-is-online":-1,"get-data":-1,"new-tab":-1,"navigate-browser-tabs":-1,search:-1,next:-1,back:-1,"show-modal-categories":-1,reload:-1,share:-1,"add-bookmark":-1,"amount-tab":-1,"focus-search":-1,"navigate-bookmark":-1,"navigate-subscription":-1,"navigate-history":-1,"navigate-extension":-1,"navigate-setting":-1,"exit-browser":-1,"add-shortcut":-1,"create-bookmark-folder":-1,"close-add-bookmark":-1,"save-bookmark":-1,"get-bookmark-folders":-1,"chosen-bookmark":-1,"get-no-folder-bookmarks":-1,"navigate-bookmark-detail":-1,"get-bookmarks":-1,"delete-bookmark":-1,"delete-bookmark-folder":-1,"choose-bookmark":-1,"get-tabs":-1,"close-tab":-1,"close-browser-tabs":-1,"close-all":-1,"get-histories":-1,"choose-history":-1,"delete-history":-1,"update-navigate":-1},y=window;function v(e){return m(this,void 0,void 0,(function(){return g(this,(function(t){switch(t.label){case 0:return[4,new Promise((function(t){var n=setInterval((function(){"number"!=typeof b[e]&&(t(),clearInterval(n))}),1e3)}))];case 1:return[2,t.sent()]}}))}))}function w(e){return m(this,void 0,void 0,(function(){return g(this,(function(t){switch(t.label){case 0:return void 0!==b[e.command]&&(b[e.command]=0),y.webkit&&y.webkit.messageHandlers&&y.webkit.messageHandlers.callbackHandler?y.webkit.messageHandlers.callbackHandler.postMessage(JSON.stringify(e)):console.log("not found globalWindow.webkit"),void 0===b[e.command]?[2]:[4,v(e.command)];case 1:return t.sent(),[2,b[e.command]]}}))}))}var k,_,x,E=new(function(e){function t(){var t=e.call(this)||this;return t._isReady=!!(y.webkit&&y.webkit.messageHandlers&&y.webkit.messageHandlers.callbackHandler)||!!window.opener,t._hasNotch=!1,t._subscribe(),t}return h(t,e),Object.defineProperty(t.prototype,"statusNotch",{get:function(){return this._hasNotch},enumerable:!1,configurable:!0}),t.prototype.setStatusNotch=function(e){void 0===e&&(e=!1),this._hasNotch=e},Object.defineProperty(t.prototype,"isReady",{get:function(){return this._isReady},enumerable:!1,configurable:!0}),t.prototype.send=function(e,t){return m(this,void 0,void 0,(function(){var t,n,r;return g(this,(function(i){switch(i.label){case 0:if(!this.isReady)throw new u;return y.webkit&&y.webkit.messageHandlers?[4,w(e)]:[3,2];case 1:if(null==(t=i.sent()))return[2];if("string"==typeof t)return[2,this._handleJsonStringMessage(t)];if(n=t.command,r=t.data,b[n]=-1,console.log("response =>",r),void 0===r)throw"Command not response - ".concat(JSON.stringify(t));if(!0!==r.success)throw r.message;return[2,t.data];case 2:return[4,this._postMessageToWindow(e)];case 3:return[2,i.sent()]}}))}))},t.prototype.exit=function(){return m(this,void 0,void 0,(function(){return g(this,(function(e){switch(e.label){case 0:return[4,this.send(c.EXIT)];case 1:return e.sent(),[2]}}))}))},t.prototype.requestPermission=function(e,t){var n;return void 0===t&&(t=void 0),m(this,void 0,void 0,(function(){var r,i;return g(this,(function(o){switch(o.label){case 0:return[4,this.send(c.PERMISSION,{permission:e,option:t})];case 1:if(r=o.sent(),!(null===(n=null==r?void 0:r.payload)||void 0===n?void 0:n.success)){if(i=new f,!(null==r?void 0:r.payload))throw i.success=!1,i.permission=e,i;throw Object.assign(i,r.payload),i}return r.payload.permission=e,[2,r.payload]}}))}))},t.prototype._postMessageToWindow=function(e){var t=this,n=this;return new Promise((function(r){var i,o=function(e){r(e),this.removeEventListener(o)};t.on(o.bind(t,n)),null===(i=window.opener)||void 0===i||i.postMessage(e,"*")}))},t.prototype._subscribe=function(){var e=this;window.addEventListener("flutterInAppWebViewPlatformReady",(function(){e._isReady=!0,e.emit("ready")})),window.addEventListener("message",(function(t){var n=t.data;return"string"==typeof n?e._handleJsonStringMessage(n,!0):"ios"==n.platform?(console.log("=== React - 33333 ====",n,n.command),!0!==n.isSocket?(console.log("=== React - 44444 ====",n.data),b[n.command]=n.data):void e.emit(n.command,n.data)):(e.emit(n.command,n.data),n)}))},t.prototype._handleJsonStringMessage=function(e,t){if(e)try{var n=JSON.parse(e),r=n.command,i=n.data;if(!0!==i.success)throw t&&this.emit(r,n),i.message;return t&&this.emit(r,n),n.data}catch(e){throw e.toString()}},t}(d)),O=n(5221),S=function(){return a.createElement(I,null,"• Allows users to connect to Web3 and access resources like decentralized applications.",a.createElement("br",null),"• Manage and access all DApps.",a.createElement("br",null),"• Private and unlimited connections with Web3")},I=O.ZP.p(k||(_=["\n  font-size: 15px;\n  line-height: 25px;\n  color: #667386;\n  padding: 0 20px;\n  display: block;\n"],x||(x=_.slice(0)),k=Object.freeze(Object.defineProperties(_,{raw:{value:Object.freeze(x)}})))),j=n(3132),P=n(5206),A=n.p+"_/_/packages/assets/images/onboarding/pose.png",z=n.p+"_/_/packages/assets/images/onboarding/which1.png",N=n.p+"_/_/packages/assets/images/onboarding/which2.png",R=n.p+"_/_/packages/assets/images/onboarding/which3.png",T=n.p+"_/_/packages/assets/images/onboarding/which4.png",M=n.p+"_/_/packages/assets/images/onboarding/which5.png",D=n(253);function L(e,t){(null==t||t>e.length)&&(t=e.length);for(var n=0,r=new Array(t);n<t;n++)r[n]=e[n];return r}var C,H=function(e){var t,n,r=e.setDistance,i=e.children,o=(t=(0,a.useState)(!1),n=2,function(e){if(Array.isArray(e))return e}(t)||function(e,t){var n=null==e?null:"undefined"!=typeof Symbol&&e[Symbol.iterator]||e["@@iterator"];if(null!=n){var r,i,o,a,s=[],c=!0,l=!1;try{if(o=(n=n.call(e)).next,0===t){if(Object(n)!==n)return;c=!1}else for(;!(c=(r=o.call(n)).done)&&(s.push(r.value),s.length!==t);c=!0);}catch(e){l=!0,i=e}finally{try{if(!c&&null!=n.return&&(a=n.return(),Object(a)!==a))return}finally{if(l)throw i}}return s}}(t,n)||function(e,t){if(e){if("string"==typeof e)return L(e,t);var n=Object.prototype.toString.call(e).slice(8,-1);return"Object"===n&&e.constructor&&(n=e.constructor.name),"Map"===n||"Set"===n?Array.from(e):"Arguments"===n||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n)?L(e,t):void 0}}(t,n)||function(){throw new TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}()),s=o[0],c=o[1];return(0,a.useEffect)((function(){var e;return s||(e=setTimeout((function(){c(!0)}),100)),function(){return e&&clearTimeout(e)}}),[s]),a.createElement(P.E.div,{onPan:function(e,t){var n=t.offset.x;s&&c(!1),r(n)},style:{touchAction:"none",height:"100%"}},i)},F=(0,a.memo)(H);function W(e,t){return function(e){if(Array.isArray(e))return e}(e)||function(e,t){var n=null==e?null:"undefined"!=typeof Symbol&&e[Symbol.iterator]||e["@@iterator"];if(null!=n){var r,i,o,a,s=[],c=!0,l=!1;try{if(o=(n=n.call(e)).next,0===t){if(Object(n)!==n)return;c=!1}else for(;!(c=(r=o.call(n)).done)&&(s.push(r.value),s.length!==t);c=!0);}catch(e){l=!0,i=e}finally{try{if(!c&&null!=n.return&&(a=n.return(),Object(a)!==a))return}finally{if(l)throw i}}return s}}(e,t)||function(e,t){if(e){if("string"==typeof e)return J(e,t);var n=Object.prototype.toString.call(e).slice(8,-1);return"Object"===n&&e.constructor&&(n=e.constructor.name),"Map"===n||"Set"===n?Array.from(e):"Arguments"===n||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n)?J(e,t):void 0}}(e,t)||function(){throw new TypeError("Invalid attempt to destructure non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}()}function J(e,t){(null==t||t>e.length)&&(t=e.length);for(var n=0,r=new Array(t);n<t;n++)r[n]=e[n];return r}var q,B,Z,U=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:5,t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:0,n=arguments.length>2&&void 0!==arguments[2]?arguments[2]:0;return{duration:e,repeat:1/0,delay:t,repeatDelay:n}},Y=function(e){var t=e.isBounceRight,n=(0,j._)(),r=(0,j._)(),i=(0,j._)(),o=(0,j._)(),s=(0,j._)(),c=W((0,a.useState)(0),2),l=c[0],u=c[1],f=W((0,a.useState)((function(){return t?"bounceInRight":"bounceInLeft"})),1)[0],p=W((0,a.useState)((function(){return t?[0,700,1e3,1300,1600,1800]:[1800,1600,1300,1e3,700,0]})),1)[0];return(0,a.useEffect)((function(){var e=setTimeout((function(){u(0)}),300);return function(){return clearTimeout(e)}}),[l]),(0,a.useEffect)((function(){n.start({top:[-30,-50,-30],transition:U()}),r.start({scale:[1,.5,1],transition:U(6)}),o.start({top:[200,150,200],transition:U(4)}),i.start({x:[-50,0,-50],transition:U(4)}),s.start({x:[-50,0,-50],transition:U(4)})}),[]),a.createElement(a.Fragment,null,a.createElement(F,{setDistance:u},a.createElement(X,{className:"wrapContent"},a.createElement("div",{className:"wrapPerson"},a.createElement(D.f,{className:"person",animationIn:f,animationInDelay:p[1]},a.createElement("img",{src:A,alt:"Person",style:{right:0-l/2}})),a.createElement(D.f,{className:"wrapAnimate",style:{zIndex:15},animationIn:f,animationInDelay:p[0]},a.createElement(P.E.img,{animate:n,className:"which1",src:z,style:{right:-180-l}})),a.createElement(D.f,{className:"wrapAnimate",animationIn:f,animationInDelay:p[2]},a.createElement(P.E.img,{animate:r,className:"which2",src:N,style:{right:-150-l}})),a.createElement(D.f,{className:"wrapAnimate",style:{zIndex:15},animationIn:f,animationInDelay:p[3]},a.createElement(P.E.img,{animate:i,className:"which3",src:R,style:{left:-30+l}})),a.createElement(D.f,{className:"wrapAnimate",animationIn:f,animationInDelay:p[4]},a.createElement(P.E.img,{animate:o,className:"which4",src:T,style:{left:-200+l}})),a.createElement(D.f,{className:"wrapAnimate",animationIn:f,animationInDelay:p[5]},a.createElement(P.E.img,{animate:s,className:"which5",src:M,style:{left:-180+l}}))))))},X=O.ZP.div(C||(C=function(e,t){return t||(t=e.slice(0)),Object.freeze(Object.defineProperties(e,{raw:{value:Object.freeze(t)}}))}(["\n  position: relative;\n  height: 100%;\n  display: flex;\n  justify-content: center;\n  align-items: center;\n  top: 10px;\n  height: 100%;\n  .wrapPerson {\n    width: 310px;\n    height: 394px;\n    position: relative;\n    display: flex;\n    justify-content: center;\n  }\n  img, .wrapAnimate {\n    position: absolute;\n  }\n  .person {\n    z-index: 10;\n   img {\n    position: relative;\n    transition: all 0.5s linear;\n   }\n  }\n  .which1 {\n    z-index: 8;\n    transition: all 0.5s linear;\n  }\n  .which2 {\n    z-index: 15;\n    transition: all 0.5s linear;\n    top: 160px;\n  }\n  .which3 {\n    transition: all 0.5s linear;\n    top: 220px;\n  }\n  .which4 {\n    z-index: 8;\n    transition: all 0.5s linear;\n  }\n  .which5 {\n    z-index: 8;\n    transition: all 0.5s linear;\n  }\n"]))),$=n.p+"_/_/packages/assets/images/bg.png",G=function(e,t){return Object.defineProperty?Object.defineProperty(e,"raw",{value:t}):e.raw=t,e},Q=function(){return Q=Object.assign||function(e){for(var t,n=1,r=arguments.length;n<r;n++)for(var i in t=arguments[n])Object.prototype.hasOwnProperty.call(t,i)&&(e[i]=t[i]);return e},Q.apply(this,arguments)},V=O.ZP.div(q||(q=G(["\n  display: flex;\n  width: 100%;\n  height: calc(100vh - 20px);\n  flex-direction: column;\n  box-sizing: border-box;\n  background: #fdfdfd;\n  position: relative;\n  overflow: hidden;\n"],["\n  display: flex;\n  width: 100%;\n  height: calc(100vh - 20px);\n  flex-direction: column;\n  box-sizing: border-box;\n  background: #fdfdfd;\n  position: relative;\n  overflow: hidden;\n"]))),K=O.ZP.div(B||(B=G(["\n  display: flex;\n  flex-direction: column;\n  width: 100%;\n  height: 85%;\n  background-repeat: no-repeat, repeat;\n  background-size: 100% 100%;\n  .wrapContent {\n    @media (max-height: 750px) {\n      transform: scale(0.9);\n    }\n    @media (max-height: 710px) {\n      transform: scale(0.8);\n    }\n  }\n"],["\n  display: flex;\n  flex-direction: column;\n  width: 100%;\n  height: 85%;\n  background-repeat: no-repeat, repeat;\n  background-size: 100% 100%;\n  .wrapContent {\n    @media (max-height: 750px) {\n      transform: scale(0.9);\n    }\n    @media (max-height: 710px) {\n      transform: scale(0.8);\n    }\n  }\n"])));s.createRoot(document.getElementById("root")).render((0,o.jsx)((function(){var e=(0,a.useState)(E.isReady),t=e[0],n=e[1],r=function(){n(E.isReady)};return(0,a.useEffect)((function(){if(E.isReady)return function(){E.removeEventListener("ready",r)};E.on("ready",r)}),[]),(0,a.useEffect)((function(){console.log(t,"core is ready ????")}),[t]),(0,o.jsx)(V,{children:(0,o.jsxs)(K,Q({style:{backgroundImage:"url(".concat($,")"),backgroundSize:"100% 100%"}},{children:[(0,o.jsx)(Y,{isBounceRight:!1}),(0,o.jsx)(S,{})]}))})}),{})),Z&&Z instanceof Function&&n.e(216).then(n.bind(n,9543)).then((function(e){var t=e.getCLS,n=e.getFID,r=e.getFCP,i=e.getLCP,o=e.getTTFB;t(Z),n(Z),r(Z),i(Z),o(Z)}))}},function(e){e.O(0,[216],(function(){return 2861,e(e.s=2861)})),e.O()}]);
//# sourceMappingURL=main.a3a958b101cb3c679cd4.js.map