(window.webpackJsonp=window.webpackJsonp||[]).push([[16],{236:function(e,t,n){"use strict";n.d(t,"a",function(){return l});var r=n(372),a=n(15),i=n.n(a),c=n(1),o=n.n(c),u=n(237),l=o.a.memo(function(e){return o.a.createElement(r.a,Object.assign({},e,{className:i()(u.ph,e.className)}))})},237:function(e,t,n){e.exports={ph:"page-header_ph__1jMB9"}},299:function(e,t,n){"use strict";n.d(t,"c",function(){return s}),n.d(t,"d",function(){return m}),n.d(t,"a",function(){return f}),n.d(t,"f",function(){return h}),n.d(t,"b",function(){return j}),n.d(t,"e",function(){return v});var r=n(241),a=n(242),i=n.n(a);function c(){var e=Object(r.a)(['\n    mutation runScriptMutation(\n        $projectId: String!\n        $id: String!\n        $input: ExecutionInput!\n    ) {\n        output(projectId: $projectId, id: $id, input: $input)\n            @rest(method: "POST", path: "/projects/{args.projectId}/exec") {\n            jobId\n        }\n    }\n']);return c=function(){return e},e}function o(){var e=Object(r.a)(['\n    mutation deleteScriptMutation($projectId: String!, $id: String!) {\n        output(projectId: $projectId, id: $id)\n            @rest(\n                method: "DELETE"\n                path: "/projects/{args.projectId}/scripts/{args.id}"\n            ) {\n            id\n            rev\n        }\n    }\n']);return o=function(){return e},e}function u(){var e=Object(r.a)(['\n    mutation updateScriptMutation(\n        $projectId: String!\n        $id: String!\n        $input: Script!\n    ) {\n        output(projectId: $projectId, id: $id, input: $input)\n            @rest(\n                method: "PUT"\n                path: "projects/{args.projectId}/scripts/{args.id}"\n            ) {\n            id\n            rev\n        }\n    }\n']);return u=function(){return e},e}function l(){var e=Object(r.a)(['\n    mutation createScriptMutation($projectId: String!, $input: Script!) {\n        output(projectId: $projectId, input: $input)\n            @rest(method: "POST", path: "projects/{args.projectId}/scripts") {\n            id\n            rev\n        }\n    }\n']);return l=function(){return e},e}function d(){var e=Object(r.a)(['\n    query findScriptsQuery($projectId: String!, $id: String!) {\n        output(projectId: $projectId, id: $id)\n            @rest(\n                type: "Script"\n                path: "/projects/{args.projectId}/scripts/{args.id}"\n            ) {\n            id\n            rev\n            createdAt\n            name\n            description\n            execution\n            persistence\n        }\n    }\n']);return d=function(){return e},e}function p(){var e=Object(r.a)(['\n    query findScripts($projectId: String!, $query: Query) {\n        output(projectId: $projectId, query: $query)\n            @rest(\n                type: "[Script]"\n                path: "/projects/{args.projectId}/scripts?{args.query}"\n            ) {\n            paging\n            data\n        }\n    }\n']);return p=function(){return e},e}var s=i()(p()),m=i()(d()),f=i()(l()),h=i()(u()),j=i()(o()),v=i()(c())},323:function(e,t,n){"use strict";n.d(t,"a",function(){return l});var r=n(549),a=n(15),i=n.n(a),c=n(1),o=n.n(c),u=n(324),l=o.a.memo(function(e){var t=e.children,n=e.className;return o.a.createElement(r.a,{className:i()(u.card,n)},t)})},324:function(e,t,n){e.exports={card:"card_card__3zxNn"}},536:function(e,t,n){e.exports={button:"panel_button__1fWh9"}},539:function(e,t,n){"use strict";n.r(t);var r=n(13),a=n(16),i=n(22),c=n(21),o=n(23),u=n(550),l=n(427),d=n(428),p=n(547),s=n(328),m=n(49),f=n.n(m),h=n(1),j=n.n(h),v=n(299),b=u.a.TextArea;function _(e){var t=e.form,n=e.value,r=void 0===n?"":n;return j.a.createElement(p.a.Item,{label:"Description"},t.getFieldDecorator("description",{initialValue:r,rules:[{min:10}]})(j.a.createElement(b,{rows:3})))}function g(e){var t=e.form,n=e.value,r=void 0===n?"":n,a=e.type,i=void 0===a?"":a;return j.a.createElement(p.a.Item,{label:"Name"},t.getFieldDecorator("name",{initialValue:r,rules:[{required:!0,message:"Please input ".concat(i," name!")},{min:3}]})(j.a.createElement(u.a,null)))}var E=n(36),y=n(31),I=n(137),O=n(540),k=n(74),$=n.n(k),S=n(69),C=n(504),F=n.n(C),q=n(362),w=n.n(q);var x=n(67);function N(e,t,n,r,a,i){if(null!=r.id){var c=a.readQuery({query:t,variables:n})[e],o=i.data?i.data.output:{},u=Object(E.a)({},c,o,r);a.writeQuery({query:t,data:Object(x.a)({},e,u)})}}var B=n(323),D=n(256),P=n(236),M=n(536),T=function(e){function t(){return Object(r.a)(this,t),Object(i.a)(this,Object(c.a)(t).apply(this,arguments))}return Object(o.a)(t,e),Object(a.a)(t,[{key:"render",value:function(){var e=this.props,t=e.title,n=e.onBack;return j.a.createElement(P.a,{title:t,onBack:n,extra:this.__renderButtons()})}},{key:"__renderButtons",value:function(){var e=this.props,t=e.onCancel,n=e.onSave,r=e.onDelete,a=e.touched,i=e.valid,c=[];return"function"===typeof t&&c.push(j.a.createElement(D.a,{key:"cancel",className:M.button,icon:"close",disabled:!a,onClick:t},"Cancel")),"function"===typeof n&&c.push(j.a.createElement(D.a,{key:"save",type:"primary",className:M.button,icon:"save",disabled:!a||!i,onClick:n},"Save")),"function"===typeof r&&c.push(j.a.createElement(D.a,{key:"delete",type:"danger",className:M.button,icon:"delete",onClick:r},"Delete")),c}}]),t}(j.a.PureComponent),V=function(e){function t(e){var n;return Object(r.a)(this,t),(n=Object(i.a)(this,Object(c.a)(t).call(this,e))).__renderForm=n.__renderForm.bind(Object(y.a)(Object(y.a)(n))),n.__handleCompletedMutation=n.__handleCompletedMutation.bind(Object(y.a)(Object(y.a)(n))),n.__handleError=n.__handleError.bind(Object(y.a)(Object(y.a)(n))),n.__updateCache=n.__updateCache.bind(Object(y.a)(Object(y.a)(n))),n}return Object(o.a)(t,e),Object(a.a)(t,[{key:"render",value:function(){var e=this,t=this.props,n=t.id,r=t.projectId,a=t.fetch,i=t.create,c=t.update,o=$()(n)||"new"===n?i:c,u={id:n,projectId:r};return j.a.createElement(S.c,{query:a,variables:u,skip:this.__isNew(),onError:this.__handleError},function(t){return j.a.createElement(S.b,{mutation:o,variables:u,update:e.__updateCache,onCompleted:e.__handleCompletedMutation,onError:e.__handleError},function(n,r){return e.__renderForm(t,n,r)})})}},{key:"__renderForm",value:function(e,t,n){var r,a=this,i=this.props,c=i.id,o=i.projectId,u=i.form,l=i.children,d=i.title,s=i.onBack,m=i.onDelete,f=i.delete,v=e.loading||n.loading,b=e.data?e.data.output:void 0,_=e.error||n.error,g=u.isFieldsTouched(),y=function e(t){return null!=t&&F()(t,function(t){return null==t||!!w()(t)&&e(t)})}(u.getFieldsError()),O="string"===typeof d?d:d(b,v,_),k=function(){a.props.form.validateFieldsAndScroll(function(e,n){e||t({variables:{id:c,projectId:o,input:Object(E.a)({},n,{id:c})}})})},$=function(){a.props.form.resetFields()};this.__isNew()||null==f||(r=function(){n.client.mutate({mutation:f,variables:{id:c,projectId:o}}).then(function(e){"function"===typeof m&&m()}).catch(a.__handleError)});var S=l({form:u,touched:g,valid:y,loading:v,value:b}),C=S;return null!=S&&null!=S.type&&(C=S.type===h.Fragment?S.props.children:l),j.a.createElement(I.a,{spinning:v},j.a.createElement(p.a,null,j.a.Children.map(C,function(e,t){var n=t.toFixed();return 0===t?j.a.createElement(B.a,{key:n},j.a.createElement(T,{title:O,touched:g,valid:y,onBack:s,onSave:k,onCancel:$,onDelete:r}),e):j.a.createElement(B.a,{key:n},e)})))}},{key:"__isNew",value:function(){var e=this.props.id;return $()(e)||"new"===e}},{key:"__handleCompletedMutation",value:function(e){var t=this.props,n=t.form,r=t.onCreate;n.resetFields(),this.__isNew()&&"function"===typeof r&&r(e.output.id)}},{key:"__handleError",value:function(e){O.a.error({message:e.message})}},{key:"__updateCache",value:function(e,t){var n=this.props,r=n.projectId;N("output",n.fetch,{id:t.data?t.data.output.id:this.props.id,projectId:r},this.__getFormValues(),e,t)}},{key:"__getFormValues",value:function(){var e=this.props.id,t=this.props.form.getFieldsValue();return Object(E.a)({},t,{id:e})}}]),t}(j.a.Component),Q=p.a.create({name:"form_container "})(V),A=n(63);n.d(t,"default",function(){return L});var J=u.a.TextArea,L=function(e){function t(){return Object(r.a)(this,t),Object(i.a)(this,Object(c.a)(t).apply(this,arguments))}return Object(o.a)(t,e),Object(a.a)(t,[{key:"render",value:function(){var e=this,t=this.props.projectId,n=this.getParam("id"),r=function(){e.navigate(e.slicePathBack(1))};return j.a.createElement(Q,{id:n,projectId:t,title:function(e){return e?e.name:""},fetch:v.d,create:v.a,update:v.f,delete:v.b,onCreate:function(t){var n=e.slicePathBack(1);e.navigate("".concat(n,"/").concat(t))},onDelete:r,onBack:r},function(e){var t=e.form,n=e.value,r=void 0===n?{}:n,a=t.getFieldDecorator;return j.a.createElement(h.Fragment,null,j.a.createElement(l.a,{gutter:16},j.a.createElement(d.a,{lg:12},j.a.createElement("fieldset",null,j.a.createElement("legend",null,"General"),j.a.createElement(g,{form:t,value:r.name}),j.a.createElement(_,{form:t,value:r.description}))),j.a.createElement(d.a,{lg:12},j.a.createElement("fieldset",null,j.a.createElement("legend",null,"Persistence"),j.a.createElement(p.a.Item,null,a("persistence.enabled",{valuePropName:"checked",initialValue:f()(r,"persistence.enabled",!1)})(j.a.createElement(s.a,null,"Enable")))))),j.a.createElement(l.a,null,j.a.createElement(d.a,{lg:24},j.a.createElement("fieldset",null,j.a.createElement("legend",null,"Execution"),j.a.createElement(p.a.Item,{label:"Query"},a("execution.query",{initialValue:f()(r,"execution.query",""),rules:[{required:!0,message:"Input script query"},{min:8}]})(j.a.createElement(J,{rows:10,minLength:8,required:!0})))))))})}}]),t}(A.a)}}]);
//# sourceMappingURL=16.751b2b14.chunk.js.map