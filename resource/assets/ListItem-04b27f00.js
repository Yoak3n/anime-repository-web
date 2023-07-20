import{f as D,g as m,h as p,i as s,j as v,k as t,l as M,m as j,d as h,n as H,p as B,q as f,r as E,t as I,s as N,v as O,x as n,y as S,z as T,A as V}from"./index-31937f98.js";const K=o=>{const{textColor2:r,cardColor:e,modalColor:l,popoverColor:d,dividerColor:a,borderRadius:b,fontSize:i,hoverColor:c}=o;return{textColor:r,color:e,colorHover:c,colorModal:l,colorHoverModal:m(l,c),colorPopover:d,colorHoverPopover:m(d,c),borderColor:a,borderColorModal:m(l,a),borderColorPopover:m(d,a),borderRadius:b,fontSize:i}},q={name:"List",common:D,self:K},A=q,F=p([s("list",`
 --n-merged-border-color: var(--n-border-color);
 --n-merged-color: var(--n-color);
 --n-merged-color-hover: var(--n-color-hover);
 margin: 0;
 font-size: var(--n-font-size);
 transition:
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 padding: 0;
 list-style-type: none;
 color: var(--n-text-color);
 background-color: var(--n-merged-color);
 `,[v("show-divider",[s("list-item",[p("&:not(:last-child)",[t("divider",`
 background-color: var(--n-merged-border-color);
 `)])])]),v("clickable",[s("list-item",`
 cursor: pointer;
 `)]),v("bordered",`
 border: 1px solid var(--n-merged-border-color);
 border-radius: var(--n-border-radius);
 `),v("hoverable",[s("list-item",`
 border-radius: var(--n-border-radius);
 `,[p("&:hover",`
 background-color: var(--n-merged-color-hover);
 `,[t("divider",`
 background-color: transparent;
 `)])])]),v("bordered, hoverable",[s("list-item",`
 padding: 12px 20px;
 `),t("header, footer",`
 padding: 12px 20px;
 `)]),t("header, footer",`
 padding: 12px 0;
 box-sizing: border-box;
 transition: border-color .3s var(--n-bezier);
 `,[p("&:not(:last-child)",`
 border-bottom: 1px solid var(--n-merged-border-color);
 `)]),s("list-item",`
 position: relative;
 padding: 12px 0; 
 box-sizing: border-box;
 display: flex;
 flex-wrap: nowrap;
 align-items: center;
 transition:
 background-color .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 `,[t("prefix",`
 margin-right: 20px;
 flex: 0;
 `),t("suffix",`
 margin-left: 20px;
 flex: 0;
 `),t("main",`
 flex: 1;
 `),t("divider",`
 height: 1px;
 position: absolute;
 bottom: 0;
 left: 0;
 right: 0;
 background-color: transparent;
 transition: background-color .3s var(--n-bezier);
 pointer-events: none;
 `)])]),M(s("list",`
 --n-merged-color-hover: var(--n-color-hover-modal);
 --n-merged-color: var(--n-color-modal);
 --n-merged-border-color: var(--n-border-color-modal);
 `)),j(s("list",`
 --n-merged-color-hover: var(--n-color-hover-popover);
 --n-merged-color: var(--n-color-popover);
 --n-merged-border-color: var(--n-border-color-popover);
 `))]),G=Object.assign(Object.assign({},f.props),{size:{type:String,default:"medium"},bordered:Boolean,clickable:Boolean,hoverable:Boolean,showDivider:{type:Boolean,default:!0}}),u=S("n-list"),Q=h({name:"List",props:G,setup(o){const{mergedClsPrefixRef:r,inlineThemeDisabled:e,mergedRtlRef:l}=H(o),d=B("List",l,r),a=f("List","-list",F,A,o,r);E(u,{showDividerRef:I(o,"showDivider"),mergedClsPrefixRef:r});const b=N(()=>{const{common:{cubicBezierEaseInOut:c},self:{fontSize:g,textColor:x,color:C,colorModal:z,colorPopover:R,borderColor:$,borderColorModal:P,borderColorPopover:k,borderRadius:_,colorHover:w,colorHoverModal:y,colorHoverPopover:L}}=a.value;return{"--n-font-size":g,"--n-bezier":c,"--n-text-color":x,"--n-color":C,"--n-border-radius":_,"--n-border-color":$,"--n-border-color-modal":P,"--n-border-color-popover":k,"--n-color-modal":z,"--n-color-popover":R,"--n-color-hover":w,"--n-color-hover-modal":y,"--n-color-hover-popover":L}}),i=e?O("list",void 0,b,o):void 0;return{mergedClsPrefix:r,rtlEnabled:d,cssVars:e?void 0:b,themeClass:i==null?void 0:i.themeClass,onRender:i==null?void 0:i.onRender}},render(){var o;const{$slots:r,mergedClsPrefix:e,onRender:l}=this;return l==null||l(),n("ul",{class:[`${e}-list`,this.rtlEnabled&&`${e}-list--rtl`,this.bordered&&`${e}-list--bordered`,this.showDivider&&`${e}-list--show-divider`,this.hoverable&&`${e}-list--hoverable`,this.clickable&&`${e}-list--clickable`,this.themeClass],style:this.cssVars},r.header?n("div",{class:`${e}-list__header`},r.header()):null,(o=r.default)===null||o===void 0?void 0:o.call(r),r.footer?n("div",{class:`${e}-list__footer`},r.footer()):null)}}),U=h({name:"ListItem",setup(){const o=T(u,null);return o||V("list-item","`n-list-item` must be placed in `n-list`."),{showDivider:o.showDividerRef,mergedClsPrefix:o.mergedClsPrefixRef}},render(){const{$slots:o,mergedClsPrefix:r}=this;return n("li",{class:`${r}-list-item`},o.prefix?n("div",{class:`${r}-list-item__prefix`},o.prefix()):null,o.default?n("div",{class:`${r}-list-item__main`},o):null,o.suffix?n("div",{class:`${r}-list-item__suffix`},o.suffix()):null,this.showDivider&&n("div",{class:`${r}-list-item__divider`}))}});export{U as N,Q as a};
