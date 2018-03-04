package doc

// The CSS comes from: https://github.com/sindresorhus/github-markdown-css
const (
	cssGeneral =`.markdown-body hr:after,.markdown-body hr:before{content: "";display: table;}.markdown-body .anchor:focus,.markdown-body a:active,.markdown-body a:hover{outline: 0;}@font-face{font-family: octicons-anchor;src: url(data:font/woff;charset=utf-8;base64,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) format('woff');}.readme{border: 1px solid #ddd;border-radius: 3px;display: table;margin: 20px auto;padding: 30px;width: 790px;}.markdown-body ol,.markdown-body td,.markdown-body th,.markdown-body ul{padding: 0;}.markdown-body{color: #333;font-family: "Helvetica Neue",Helvetica,"Segoe UI",Arial,freesans,sans-serif;font-size: 16px;line-height: 1.6;ms-text-size-adjust: 100%%;overflow: hidden;webkit-text-size-adjust: 100%%;word-wrap: break-word;}.markdown-body strong{font-weight: 700;}.markdown-body h1{margin: .67em 0;}.markdown-body img{border: 0;}.markdown-body hr{box-sizing: content-box;moz-box-sizing: content-box;}.markdown-body input{color: inherit;font: 13px/1.4 Helvetica,arial,freesans,clean,sans-serif,"Segoe UI Emoji","Segoe UI Symbol";line-height: normal;margin: 0;}.markdown-body code,.markdown-body pre{font: 12px Consolas,"Liberation Mono",Menlo,Courier,monospace;}.markdown-body html input[disabled]{cursor: default;}.markdown-body input[type=checkbox]{box-sizing: border-box;moz-box-sizing: border-box;padding: 0;}.markdown-body *{box-sizing: border-box;moz-box-sizing: border-box;}.markdown-body a{background: 0 0;color: #4183c4;text-decoration: none;}.markdown-body a:active,.markdown-body a:focus,.markdown-body a:hover{text-decoration: underline;}.markdown-body hr:after{clear: both;}.markdown-body blockquote{margin: 0;}.markdown-body h1,.markdown-body h2{border-bottom: 1px solid #eee;padding-bottom: .3em;}.markdown-body ol ol,.markdown-body ul ol{list-style-type: lower-roman;}.markdown-body ol ol ol,.markdown-body ol ul ol,.markdown-body ul ol ol,.markdown-body ul ul ol{list-style-type: lower-alpha;}.markdown-body dd{margin-left: 0;}.markdown-body pre{word-wrap: normal;}.markdown-body .octicon{display: inline-block;font: normal normal 16px octicons-anchor;line-height: 1;moz-osx-font-smoothing: grayscale;moz-user-select: none;ms-user-select: none;text-decoration: none;user-select: none;webkit-font-smoothing: antialiased;webkit-user-select: none;}.markdown-body .octicon-link:before{content: '\f05c';}.markdown-body>:first-child{margin-top: 0!important;}.markdown-body>:last-child{margin-bottom: 0!important;}.markdown-body .anchor{bottom: 0;display: block;left: 0;margin-left: -30px;padding-left: 30px;padding-right: 6px;position: absolute;top: 0;}.markdown-body h1,.markdown-body h2,.markdown-body h3,.markdown-body h4,.markdown-body h5,.markdown-body h6{font-weight: 700;line-height: 1.4;margin-bottom: 16px;margin-top: 1em;position: relative;}.markdown-body h1 .octicon-link,.markdown-body h2 .octicon-link,.markdown-body h3 .octicon-link,.markdown-body h4 .octicon-link,.markdown-body h5 .octicon-link,.markdown-body h6 .octicon-link{color: #000;display: none;vertical-align: middle;}.markdown-body h1:hover .anchor,.markdown-body h2:hover .anchor,.markdown-body h3:hover .anchor,.markdown-body h4:hover .anchor,.markdown-body h5:hover .anchor,.markdown-body h6:hover .anchor{line-height: 1;margin-left: -30px;padding-left: 8px;text-decoration: none;}.markdown-body h1:hover .anchor .octicon-link,.markdown-body h2:hover .anchor .octicon-link,.markdown-body h3:hover .anchor .octicon-link,.markdown-body h4:hover .anchor .octicon-link,.markdown-body h5:hover .anchor .octicon-link,.markdown-body h6:hover .anchor .octicon-link{display: inline-block;}.markdown-body h1{font-size: 2.25em;line-height: 1.2;}.markdown-body h2{font-size: 1.75em;line-height: 1.225;}.markdown-body h3{font-size: 1.5em;line-height: 1.43;}.markdown-body h4{font-size: 1.25em;}.markdown-body h5{font-size: 1em;}.markdown-body h6{color: #777;font-size: 1em;}.markdown-body blockquote,.markdown-body dl,.markdown-body ol,.markdown-body p,.markdown-body pre,.markdown-body table,.markdown-body ul{margin-bottom: 16px;margin-top: 0;}.markdown-body hr{background: #e7e7e7;border: 0;height: 4px;margin: 16px 0;overflow: hidden;padding: 0;}.markdown-body ol,.markdown-body ul{padding-left: 2em;}.markdown-body ol ol,.markdown-body ol ul,.markdown-body ul ol,.markdown-body ul ul{margin-bottom: 0;margin-top: 0;}.markdown-body li>p{margin-top: 16px;}.markdown-body dl{padding: 0;}.markdown-body dl dt{font-size: 1em;font-style: italic;font-weight: 700;margin-top: 16px;padding: 0;}.markdown-body dl dd{margin-bottom: 16px;padding: 0 16px;}.markdown-body blockquote{border-left: 4px solid #ddd;color: #777;padding: 0 15px;}.markdown-body blockquote>:first-child{margin-top: 0;}.markdown-body blockquote>:last-child{margin-bottom: 0;}.markdown-body table{border-collapse: collapse;border-spacing: 0;display: block;overflow: auto;width: 100%%;word-break: keep-all;word-break: normal;}.markdown-body table th{font-weight: 700;}.markdown-body table td,.markdown-body table th{border: 1px solid #ddd;padding: 6px 13px;}.markdown-body table tr{background-color: #fff;border-top: 1px solid #ccc;}.markdown-body table tr:nth-child(2n){background-color: #f8f8f8;}.markdown-body img{box-sizing: border-box;max-width: 100%%;moz-box-sizing: border-box;}.markdown-body code{background-color: rgba(0,0,0,.04);border-radius: 3px;font-size: 85%%;margin: 0;padding: .2em 0;}.markdown-body code:after,.markdown-body code:before{content: "\00a0";letter-spacing: -.2em;}.markdown-body pre>code{background: 0 0;border: 0;font-size: 100%%;margin: 0;padding: 0;white-space: pre;word-break: normal;}.markdown-body .highlight{color: %s;margin-bottom: 16px;}.markdown-body .highlight pre,.markdown-body pre{background-color: %s;border-radius: 3px;font-size: 85%%;line-height: 1.45;overflow: auto;padding: 16px;}.markdown-body .highlight pre{margin-bottom: 0;word-break: normal;}.markdown-body pre code{background-color: transparent;border: 0;display: inline;line-height: inherit;margin: 0;max-width: initial;overflow: initial;padding: 0;word-wrap: normal;}.markdown-body pre code:after,.markdown-body pre code:before{content: normal;}.markdown-body .pl-c{color: %s;}.markdown-body .pl-c1,.markdown-body .pl-mdh,.markdown-body .pl-mm,.markdown-body .pl-mp,.markdown-body .pl-mr,.markdown-body .pl-s1 .pl-v,.markdown-body .pl-s3,.markdown-body .pl-sc,.markdown-body .pl-sv{color: #0086b3;}.markdown-body .pl-e,.markdown-body .pl-en{color: %s;}.markdown-body .pl-s1 .pl-s2,.markdown-body .pl-smi,.markdown-body .pl-smp,.markdown-body .pl-stj,.markdown-body .pl-vo,.markdown-body .pl-vpf{color: %s;}.markdown-body .pl-ent{color: #63a35c;}.markdown-body .pl-k,.markdown-body .pl-s,.markdown-body .pl-st{color: %s;}.markdown-body .pl-pds,.markdown-body .pl-s1,.markdown-body .pl-s1 .pl-pse .pl-s2,.markdown-body .pl-sr,.markdown-body .pl-sr .pl-cce,.markdown-body .pl-sr .pl-sra,.markdown-body .pl-sr .pl-sre,.markdown-body .pl-src,.markdown-body .pl-v{color: #df5000;}.markdown-body .pl-id{color: #b52a1d;}.markdown-body .pl-ii{background-color: #b52a1d;color: #f8f8f8;}.markdown-body .pl-sr .pl-cce{color: #63a35c;font-weight: 700;}.markdown-body .pl-ml{color: #693a17;}.markdown-body .pl-mh,.markdown-body .pl-mh .pl-en,.markdown-body .pl-ms{color: #1d3e81;font-weight: 700;}.markdown-body .pl-mq{color: teal;}.markdown-body .pl-mi{color: %s;font-style: italic;}.markdown-body .pl-mb{color: %s;font-weight: 700;}.markdown-body .pl-md,.markdown-body .pl-mdhf{background-color: #ffecec;color: #bd2c00;}.markdown-body .pl-mdht,.markdown-body .pl-mi1{background-color: #eaffea;color: #55a532;}.markdown-body .pl-mdr{color: #795da3;font-weight: 700;}.markdown-body .pl-mo{color: #1d3e81;}.markdown-body kbd{background-color: #e7e7e7;background-image: linear-gradient(#fefefe,#e7e7e7);background-image: -webkit-linear-gradient(#fefefe,#e7e7e7);background-repeat: repeat-x;border: 1px solid #cfcfcf;border-radius: 2px;color: #000;display: inline-block;font: 11px Consolas,"Liberation Mono",Menlo,Courier,monospace;line-height: 10px;padding: 3px 5px;}a.source-button,a.toc-button{border-radius: 3px;cursor: pointer;font-size: large;font-weight: 700;}.markdown-body .task-list-item{list-style-type: none;}.markdown-body .task-list-item+.task-list-item{margin-top: 3px;}.markdown-body .task-list-item input{float: left;margin: .3em 0 .25em -1.6em;vertical-align: middle;}.markdown-body :checked+.radio-label{border-color: #4183c4;position: relative;z-index: 1;}a.toc-button:hover{background: #E8E4D9;color: #A88C3F;}a.toc-button{color: #465DA6;padding: .3em .5em .5em;}a.source-button:hover{background: #E8E4D9;color: #A88C3F;}a.source-button{color: #465DA6;}.markdown-body #user-content-note {background-color:#6ab0de;color:white;}.markdown-body #user-content-note p {background-color:#e7f2fa;color:black;padding: 10px;}.markdown-body #user-content-warning {background-color:#f0b37e;color:white;}.markdown-body #user-content-warning p {background-color:#ffedcc;color:black;padding: 10px;}`
)
var (
	//Predefined css style color
	BuiltinCssStyle = map[int][]string{
		0: []string{"black",   "#f7f7f7", "#969896", "#795da3", "#333",    "#a71d5d", "#333",    "#333"},      //cssGitHub
		1: []string{"#dcdcdc", "#3f3f3f", "#7f9f7f", "#efef8f", "#efef8f", "#e3ceab", "#efef8f", "#efef8f"},   //cssZenburn
		2: []string{"#695d69", "#f7f3f7", "#379a37", "#776977", "#776977", "#7b59c0", "#776977", "#776977"},   //cssLake
		3: []string{"#5e6e5e", "#f0fff0", "#379a37", "#40a070", "#776977", "#ad2bee", "#776977", "#776977"},   //cssSeaSide
		4: []string{"#5e6e5e", "#fbebd4", "#379a37", "#40a070", "#a57a4c", "#ad2bee", "#a57a4c", "#a57a4c"},   //cssKimbieLight
		5: []string{"#0ff",    "navy",    "#888",    "#ff0",    "#fff",    "#fff",    "#a57a4c", "#a57a4c"},   //cssLightBlue
		6: []string{"#bbbb9d", "#282c34", "#969896", "#51b6c2", "#7cc379", "#c678dd", "#333",    "#333"},      //cssAtomDark
		7: []string{"#51616b", "#f2fbff", "grey",    "green",   "#008bb7", "#e68538", "#333",    "#333"},      //cssForgottenLight
	}

)