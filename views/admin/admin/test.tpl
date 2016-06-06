<!DOCTYPE html>
<html lang="en">
	<head>
		<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
		<meta charset="utf-8" />
		<title>User Profile Page - Ace Admin</title>

		<meta name="description" content="3 styles with inline editable feature" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0" />
		<!-- bootstrap & fontawesome -->
		<link rel="stylesheet" href="/static/assets/css/bootstrap.min.css" />
		<link rel="stylesheet" href="/static/assets/css/font-awesome.min.css" />
		<!-- page specific plugin styles -->
		<link rel="stylesheet" href="/static/assets/css/jquery-ui.custom.min.css" />
		<link rel="stylesheet" href="/static/assets/css/jquery.gritter.css" />
		<link rel="stylesheet" href="/static/assets/css/select2.css" />
		<link rel="stylesheet" href="/static/assets/css/datepicker.css" />
		<link rel="stylesheet" href="/static/assets/css/bootstrap-editable.css" />
		<!-- text fonts -->
		<link rel="stylesheet" href="/static/assets/css/ace-fonts.css" />
		<!-- ace styles -->
		<link rel="stylesheet" href="/static/assets/css/ace.min.css" id="main-ace-style" />
		<!--[if lte IE 9]>
			<link rel="stylesheet" href="/static/assets/css/ace-part2.min.css" />
		<![endif]-->
		<link rel="stylesheet" href="/static/assets/css/ace-skins.min.css" />
		<link rel="stylesheet" href="/static/assets/css/ace-rtl.min.css" />
		<!--[if lte IE 9]>
		  <link rel="stylesheet" href="/static/assets/css/ace-ie.min.css" />
		<![endif]-->

		<!-- inline styles related to this page -->

		<!-- ace settings handler -->
		<script src="/static/assets/js/ace-extra.min.js"></script>
		<!-- HTML5shiv and Respond.js for IE8 to support HTML5 elements and media queries -->
		<!--[if lte IE 8]>
		<script src="/static/assets/js/html5shiv.min.js"></script>
		<script src="/static/assets/js/respond.min.js"></script>
		<![endif]-->

		<!-- basic scripts -->
		<!--[if !IE]> -->
		<script type="text/javascript">
			window.jQuery || document.write("<script src='/static/assets/js/jquery.min.js'>"+"<"+"/script>");
		</script>

		<!-- <![endif]-->

		<!--[if IE]>
		<script type="text/javascript">
		 window.jQuery || document.write("<script src='/static/assets/js/jquery1x.min.js'>"+"<"+"/script>");
		</script>
		<![endif]-->
		<script type="text/javascript">
			if('ontouchstart' in document.documentElement) document.write("<script src='/static/assets/js/jquery.mobile.custom.min.js'>"+"<"+"/script>");
		</script>
		<script src="/static/assets/js/bootstrap.min.js"></script>

		<!-- page specific plugin scripts -->

		<!--[if lte IE 8]>
		  <script src="/static/assets/js/excanvas.min.js"></script>
		<![endif]-->
	</head>

	<body class="no-skin">
		<!-- #section:basics/navbar.layout -->
		<div id="navbar" class="navbar navbar-default">
		    <script type="text/javascript">
		        try { ace.settings.check('navbar' , 'fixed')}catch(e){}
		    </script>

		    <div class="navbar-container" id="navbar-container">
		        <!-- #section:basics/sidebar.mobile.toggle -->
		        <button type="button" class="navbar-toggle menu-toggler pull-left" id="menu-toggler">
		            <span class="sr-only">Toggle sidebar</span>
		            <span class="icon-bar"></span>
		            <span class="icon-bar"></span>
		            <span class="icon-bar"></span>
		        </button>

		        <div class="navbar-header pull-left">
		            <a href="/admin/site" class="navbar-brand">
		                <small>GO WEB 后台管理</small>
		            </a>
		        </div>

		        <div class="navbar-buttons navbar-header pull-right" role="navigation">
		            <ul class="nav ace-nav">
		                <!-- 用户信息显示 -->
		                <li class="light-blue">
		                    <a data-toggle="dropdown" href="#" class="dropdown-toggle">
		                        <img class="nav-user-photo" src="/static/avatars/avatar.jpg" alt="Jason's Photo" />
		                            <span class="user-info">
		                                <small>欢迎登录</small>{{$.admin.Username}}
		                            </span>
		                        <i class="ace-icon fa fa-caret-down"></i>
		                    </a>

		                    <ul class="user-menu dropdown-menu-right dropdown-menu dropdown-yellow dropdown-caret dropdown-close">
		                        <li>
		                            <a href="#">
		                                <i class="ace-icon fa fa-cog"></i>
		                                设置
		                            </a>
		                        </li>
		                        <li>
		                            <a href="profile.html">
		                                <i class="ace-icon fa fa-user"></i>
		                                个人信息
		                            </a>
		                        </li>
		                        <li class="divider"></li>
		                        <li>
		                            <a href="/admin/logout">
		                                <i class="ace-icon fa fa-power-off"></i>
		                                退出
		                            </a>
		                        </li>
		                    </ul>
		                </li>

		            </ul>
		        </div>
		    </div>
		</div>

		<!-- /section:basics/navbar.layout -->
		<div class="main-container" id="main-container">
			<script type="text/javascript">
				try{ace.settings.check('main-container' , 'fixed')}catch(e){}
			</script>

			<!-- #section:basics/sidebar -->
			<div id="sidebar" class="sidebar responsive">
		        <script type="text/javascript">
		            try{ace.settings.check('sidebar' , 'fixed')}catch(e){}
		        </script>

		        <div class="sidebar-shortcuts" id="sidebar-shortcuts">
		            <div class="sidebar-shortcuts-large" id="sidebar-shortcuts-large">
		                <button class="btn btn-success">
		                    <i class="ace-icon fa fa-signal"></i>
		                </button>

		                <button class="btn btn-info">
		                    <i class="ace-icon fa fa-pencil"></i>
		                </button>

		                <button class="btn btn-warning">
		                    <i class="ace-icon glyphicon glyphicon-user"></i>
		                </button>

		                <button class="btn btn-danger">
		                    <i class="ace-icon fa fa-cogs"></i>
		                </button>
		            </div>
		            <div class="sidebar-shortcuts-mini" id="sidebar-shortcuts-mini">
		                <span class="btn btn-success"></span>
		                <span class="btn btn-info"></span>
		                <span class="btn btn-warning"></span>
		                <span class="btn btn-danger"></span>
		            </div>
		        </div>

		        <!--左侧导航栏信息-->
		        <ul class="nav nav-list">
		            <li class="admin-site">
		                <a  href="/admin/admin/index.html" >
		                    <i class="menu-icon glyphicon glyphicon-user"></i>
		                    <span class="menu-text"> 管理员信息 </span>
		                </a>
		            </li>
		            <li class="categoryindex">
		                <a  href="/admin/category/index.html" >
		                    <i class="menu-icon glyphicon glyphicon-list"></i>
		                    <span class="menu-text"> 文章分类信息 </span>
		                </a>
		            </li>
		            <li class="menuindex">
		                <a href="/admin/menu/index.html" >
		                    <i class="menu-icon glyphicon glyphicon-th"></i>
		                    <span class="menu-text"> 导航栏信息 </span>
		                </a>
		            </li>
		            <li class="">
		                <a  href="#" class="dropdown-toggle">
		                    <i class="menu-icon fa fa-desktop"></i>
		                    <span class="menu-text"> UI界面&amp;元素 </span>
		                    <b class="arrow fa fa-angle-down"></b>
		                </a>
		                <b class="arrow"></b>
		                <!--第二级别-->
		                <ul class="submenu">
		                    <li class=" ">
		                        <a  href="#" class="dropdown-toggle" >
		                            <i class="menu-icon fa fa-caret-right"></i>
		                            布局<b class="arrow fa fa-angle-down"></b>
		                        </a>
		                        <b class="arrow"></b>
		                        <ul class="submenu">
		                            <li class="othertop">
		                                <a href="/admin/other/top.html">
		                                <i class="menu-icon fa fa-caret-right"></i>
		                                头部导航</a>
		                                <b class="arrow"></b>
		                            </li>
		                        </ul>
		                    </li>
		                </ul>
		            </li>
		            <li class="">
		                <a  href="#" class="dropdown-toggle">
		                    <i class="menu-icon fa fa-file-o"></i>
		                    <span class="menu-text"> 其他页面 </span>
		                    <b class="arrow fa fa-angle-down"></b>
		                </a>
		                <b class="arrow"></b>
		                <!--第二级别-->
		                <ul class="submenu">
		                    <li class="othererror404 ">
		                        <a  href="/admin/other/error404.html"  >
		                        <i class="menu-icon fa fa-caret-right"></i>
		                        Error 404</a>
		                    </li>
		                    <li class="othererror500 ">
		                        <a  href="/admin/other/error500.html"  >
		                        <i class="menu-icon fa fa-caret-right"></i>
		                        Error 500</a>
		                    </li>
		                    <li class="otherblankpage ">
		                        <a  href="/admin/other/blankpage.html"  >
		                        <i class="menu-icon fa fa-caret-right"></i>
		                        空白页</a>
		                    </li>
		                </ul>
		            </li>
		        </ul>

		        <div class="sidebar-toggle sidebar-collapse" id="sidebar-collapse">
		            <i class="ace-icon fa fa-angle-double-left" data-icon1="ace-icon fa fa-angle-double-left" data-icon2="ace-icon fa fa-angle-double-right"></i>
		        </div>

		        <script type="text/javascript">
		            try{ace.settings.check('sidebar' , 'collapsed')}catch(e){}
		        </script>
		    </div>

			<!-- /section:basics/sidebar -->
			<div class="main-content">
				<!--头部可固定导航信息-->
		        <div class="breadcrumbs" id="breadcrumbs">
		            <script type="text/javascript">
		                try{ace.settings.check('breadcrumbs' , 'fixed')}catch(e){}
		            </script>

		            <!--面包屑信息-->
		            <ul class="breadcrumb">
		                <li>
		                    <i class="ace-icon fa fa-home home-icon"></i>
		                    <a href="/admin/site">首页</a>
		                </li>
		                <li class="active"></li>
		            </ul>

		            <!--搜索-->
		            <div class="nav-search" id="nav-search">
		                <form class="form-search">
		                    <span class="input-icon">
		                        <input type="text" placeholder="搜索/static." class="nav-search-input" id="nav-search-input" autocomplete="off" />
		                        <i class="ace-icon fa fa-search nav-search-icon"></i>
		                    </span>
		                </form>
		            </div>
		        </div>

				<!-- /section:basics/content.breadcrumbs -->
				<div class="page-content">
					<!--样式设置信息-->
		            <div class="ace-settings-container" id="ace-settings-container">
		                <div class="btn btn-app btn-xs btn-warning ace-settings-btn" id="ace-settings-btn">
		                    <i class="ace-icon fa fa-cog bigger-150"></i>
		                </div>


		                <div class="ace-settings-box clearfix" id="ace-settings-box">
		                    <div class="pull-left width-50">
		                        <div class="ace-settings-item">
		                            <div class="pull-left">
		                                <select id="skin-colorpicker" class="hide">
		                                    <option data-skin="no-skin" value="#438EB9">#438EB9</option>
		                                    <option data-skin="skin-1" value="#222A2D">#222A2D</option>
		                                    <option data-skin="skin-2" value="#C6487E">#C6487E</option>
		                                    <option data-skin="skin-3" value="#D0D0D0">#D0D0D0</option>
		                                </select>
		                            </div>
		                            <span>&nbsp; 选择皮肤 </span>
		                        </div>

		                        <div class="ace-settings-item">
		                            <input type="checkbox" class="ace ace-checkbox-2" id="ace-settings-navbar" />
		                            <label class="lbl" for="ace-settings-navbar"> 固定导航栏 </label>
		                        </div>

		                        <div class="ace-settings-item">
		                            <input type="checkbox" class="ace ace-checkbox-2" id="ace-settings-sidebar" />
		                            <label class="lbl" for="ace-settings-sidebar"> 固定侧边栏 </label>
		                        </div>

		                        <div class="ace-settings-item">
		                            <input type="checkbox" class="ace ace-checkbox-2" id="ace-settings-breadcrumbs" />
		                            <label class="lbl" for="ace-settings-breadcrumbs"> 固定的面包屑导航</label>
		                        </div>

		                        <div class="ace-settings-item">
		                            <input type="checkbox" class="ace ace-checkbox-2" id="ace-settings-rtl" />
		                            <label class="lbl" for="ace-settings-rtl"> 从右到左（替换）</label>
		                        </div>

		                        <div class="ace-settings-item">
		                            <input type="checkbox" class="ace ace-checkbox-2" id="ace-settings-add-container" />
		                            <label class="lbl" for="ace-settings-add-container">
		                                缩小显示
		                            </label>
		                        </div>
		                    </div>

		                    <div class="pull-left width-50">
		                        <div class="ace-settings-item">
		                            <input type="checkbox" class="ace ace-checkbox-2" id="ace-settings-hover" />
		                            <label class="lbl" for="ace-settings-hover"> 菜单收缩</label>
		                        </div>

		                        <div class="ace-settings-item">
		                            <input type="checkbox" class="ace ace-checkbox-2" id="ace-settings-compact" />
		                            <label class="lbl" for="ace-settings-compact"> 简单菜单</label>
		                        </div>

		                        <div class="ace-settings-item">
		                            <input type="checkbox" class="ace ace-checkbox-2" id="ace-settings-highlight" />
		                            <label class="lbl" for="ace-settings-highlight"> 当前菜单标记变换</label>
		                        </div>
		                    </div>
		                </div>
		            </div>

					<div class="page-content-area">
						<div class="page-header">
		                    <h1> 我的信息
		                    <small>
		                        <i class="ace-icon fa fa-angle-double-right"></i>
		                        编辑我的信息
		                    </small>
		                    </h1>
		                </div>

						<div class="row">
							<div class="col-xs-12">
								<div class="clearfix">
									<div class="pull-left alert alert-success no-margin">
										<button data-dismiss="alert" class="close" type="button">
											<i class="ace-icon fa fa-times"></i>
										</button>
										<i class="ace-icon fa fa-umbrella bigger-120 blue"></i>
										点击下面的图片或配置文件领域来编辑 ...
									</div>

									<div class="pull-right">
										<span class="green middle bolder">选择配置文件: &nbsp;</span>
										<div class="btn-toolbar inline middle no-margin">
											<div class="btn-group no-margin" data-toggle="buttons">
												<label class="btn btn-sm btn-yellow active">
													<span class="bigger-110">1</span>
													<input type="radio" value="1">
												</label>
												<label class="btn btn-sm btn-yellow">
													<span class="bigger-110">2</span>
													<input type="radio" value="2">
												</label>
												<label class="btn btn-sm btn-yellow">
													<span class="bigger-110">3</span>
													<input type="radio" value="3">
												</label>
											</div>
										</div>
									</div>
								</div>
								<div class="hr dotted"></div>
								{{template "/admin/admin/view.html" .}}
								{{template "/admin/admin/info.html" .}}				
								{{template "/admin/admin/edit.html" .}}
							</div>
						</div>
					{{.LayoutContent}}
					</div>
				</div>
			</div>

			<div class="footer">
				<div class="footer-inner">
					<!-- #section:basics/footer -->
					<div class="footer-content">
						<span class="bigger-120">
							<span class="blue bolder">Ace</span>
							Application &copy; 2013-2014
						</span>

						&nbsp; &nbsp;
						<span class="action-buttons">
							<a href="#">
								<i class="ace-icon fa fa-twitter-square light-blue bigger-150"></i>
							</a>

							<a href="#">
								<i class="ace-icon fa fa-facebook-square text-primary bigger-150"></i>
							</a>

							<a href="#">
								<i class="ace-icon fa fa-rss-square orange bigger-150"></i>
							</a>
						</span>
					</div>

					<!-- /section:basics/footer -->
				</div>
			</div>

			<a href="#" id="btn-scroll-up" class="btn-scroll-up btn btn-sm btn-inverse">
				<i class="ace-icon fa fa-angle-double-up icon-only bigger-110"></i>
			</a>
		</div><!-- /.main-container -->
		<script src="/static/assets/js/jquery-ui.custom.min.js"></script>
		<script src="/static/assets/js/jquery.ui.touch-punch.min.js"></script>
		<script src="/static/assets/js/jquery.gritter.min.js"></script>
		<script src="/static/assets/js/bootbox.min.js"></script>
		<script src="/static/assets/js/jquery.easypiechart.min.js"></script>
		<script src="/static/assets/js/date-time/bootstrap-datepicker.min.js"></script>
		<script src="/static/assets/js/jquery.hotkeys.min.js"></script>
		<script src="/static/assets/js/bootstrap-wysiwyg.min.js"></script>
		<script src="/static/assets/js/select2.min.js"></script>
		<script src="/static/assets/js/fuelux/fuelux.spinner.min.js"></script>
		<script src="/static/assets/js/x-editable/bootstrap-editable.min.js"></script>
		<script src="/static/assets/js/x-editable/ace-editable.min.js"></script>
		<script src="/static/assets/js/jquery.maskedinput.min.js"></script>

		<!-- ace scripts -->
		<script src="/static/assets/js/ace-elements.min.js"></script>
		<script src="/static/assets/js/ace.min.js"></script>

		<!-- inline scripts related to this page -->
		<script type="text/javascript">
			// jQuery(function($) {
			
			// 	//editables on first profile page
			// 	$.fn.editable.defaults.mode = 'inline';
			// 	$.fn.editableform.loading = "<div class='editableform-loading'><i class='ace-icon fa fa-spinner fa-spin fa-2x light-blue'></i></div>";
			//     $.fn.editableform.buttons = '<button type="submit" class="btn btn-info editable-submit"><i class="ace-icon fa fa-check"></i></button>'+
			//                                 '<button type="button" class="btn editable-cancel"><i class="ace-icon fa fa-times"></i></button>';    
				
			// 	//editables 
				
			// 	//text editable
			//     $('#username')
			// 	.editable({
			// 		type: 'text',
			// 		name: 'username'
			//     });
			
			
				
			// 	//select2 editable
			// 	var countries = [];
			//     $.each({ "CA": "Canada", "IN": "India", "NL": "Netherlands", "TR": "Turkey", "US": "United States"}, function(k, v) {
			//         countries.push({id: k, text: v});
			//     });
			
			// 	var cities = [];
			// 	cities["CA"] = [];
			// 	$.each(["Toronto", "Ottawa", "Calgary", "Vancouver"] , function(k, v){
			// 		cities["CA"].push({id: v, text: v});
			// 	});
			// 	cities["IN"] = [];
			// 	$.each(["Delhi", "Mumbai", "Bangalore"] , function(k, v){
			// 		cities["IN"].push({id: v, text: v});
			// 	});
			// 	cities["NL"] = [];
			// 	$.each(["Amsterdam", "Rotterdam", "The Hague"] , function(k, v){
			// 		cities["NL"].push({id: v, text: v});
			// 	});
			// 	cities["TR"] = [];
			// 	$.each(["Ankara", "Istanbul", "Izmir"] , function(k, v){
			// 		cities["TR"].push({id: v, text: v});
			// 	});
			// 	cities["US"] = [];
			// 	$.each(["New York", "Miami", "Los Angeles", "Chicago", "Wysconsin"] , function(k, v){
			// 		cities["US"].push({id: v, text: v});
			// 	});
				
			// 	var currentValue = "NL";
			//     $('#country').editable({
			// 		type: 'select2',
			// 		value : 'NL',
			// 		//onblur:'ignore',
			//         source: countries,
			// 		select2: {
			// 			'width': 140
			// 		},		
			// 		success: function(response, newValue) {
			// 			if(currentValue == newValue) return;
			// 			currentValue = newValue;
						
			// 			var new_source = (!newValue || newValue == "") ? [] : cities[newValue];
						
			// 			//the destroy method is causing errors in x-editable v1.4.6+
			// 			//it worked fine in v1.4.5
			// 			/**			
			// 			$('#city').editable('destroy').editable({
			// 				type: 'select2',
			// 				source: new_source
			// 			}).editable('setValue', null);
			// 			*/
						
			// 			//so we remove it altogether and create a new element
			// 			var city = $('#city').removeAttr('id').get(0);
			// 			$(city).clone().attr('id', 'city').text('Select City').editable({
			// 				type: 'select2',
			// 				value : null,
			// 				//onblur:'ignore',
			// 				source: new_source,
			// 				select2: {
			// 					'width': 140
			// 				}
			// 			}).insertAfter(city);//insert it after previous instance
			// 			$(city).remove();//remove previous instance
						
			// 		}
			//     });
			
			// 	$('#city').editable({
			// 		type: 'select2',
			// 		value : 'Amsterdam',
			// 		//onblur:'ignore',
			//         source: cities[currentValue],
			// 		select2: {
			// 			'width': 140
			// 		}
			//     });
			
			
				
			// 	//custom date editable
			// 	$('#signup').editable({
			// 		type: 'adate',
			// 		date: {
			// 			//datepicker plugin options
			// 			    format: 'yyyy/mm/dd',
			// 			viewformat: 'yyyy/mm/dd',
			// 			 weekStart: 1
						 
			// 			//,nativeUI: true//if true and browser support input[type=date], native browser control will be used
			// 			//,format: 'yyyy-mm-dd',
			// 			//viewformat: 'yyyy-mm-dd'
			// 		}
			// 	})
			
			//     $('#age').editable({
			//         type: 'spinner',
			// 		name : 'age',
			// 		spinner : {
			// 			min : 16,
			// 			max : 99,
			// 			step: 1,
			// 			on_sides: true
			// 			//,nativeUI: true//if true and browser support input[type=number], native browser control will be used
			// 		}
			// 	});
				
			
			//     $('#login').editable({
			//         type: 'slider',
			// 		name : 'login',
					
			// 		slider : {
			// 			 min : 1,
			// 			  max: 50,
			// 			width: 100
			// 			//,nativeUI: true//if true and browser support input[type=range], native browser control will be used
			// 		},
			// 		success: function(response, newValue) {
			// 			if(parseInt(newValue) == 1)
			// 				$(this).html(newValue + " hour ago");
			// 			else $(this).html(newValue + " hours ago");
			// 		}
			// 	});
			
			// 	$('#about').editable({
			// 		mode: 'inline',
			//         type: 'wysiwyg',
			// 		name : 'about',
			
			// 		wysiwyg : {
			// 			//css : {'max-width':'300px'}
			// 		},
			// 		success: function(response, newValue) {
			// 		}
			// 	});
				
				
				
			// 	// *** editable avatar *** //
			// 	try {//ie8 throws some harmless exceptions, so let's catch'em
			
			// 		//first let's add a fake appendChild method for Image element for browsers that have a problem with this
			// 		//because editable plugin calls appendChild, and it causes errors on IE at unpredicted points
			// 		try {
			// 			document.createElement('IMG').appendChild(document.createElement('B'));
			// 		} catch(e) {
			// 			Image.prototype.appendChild = function(el){}
			// 		}
			
			// 		var last_gritter
			// 		$('#avatar').editable({
			// 			type: 'image',
			// 			name: 'avatar',
			// 			value: null,
			// 			image: {
			// 				//specify ace file input plugin's options here
			// 				btn_choose: 'Change Avatar',
			// 				droppable: true,
			// 				maxSize: 110000,//~100Kb
			
			// 				//and a few extra ones here
			// 				name: 'avatar',//put the field name here as well, will be used inside the custom plugin
			// 				on_error : function(error_type) {//on_error function will be called when the selected file has a problem
			// 					if(last_gritter) $.gritter.remove(last_gritter);
			// 					if(error_type == 1) {//file format error
			// 						last_gritter = $.gritter.add({
			// 							title: 'File is not an image!',
			// 							text: 'Please choose a jpg|gif|png image!',
			// 							class_name: 'gritter-error gritter-center'
			// 						});
			// 					} else if(error_type == 2) {//file size rror
			// 						last_gritter = $.gritter.add({
			// 							title: 'File too big!',
			// 							text: 'Image size should not exceed 100Kb!',
			// 							class_name: 'gritter-error gritter-center'
			// 						});
			// 					}
			// 					else {//other error
			// 					}
			// 				},
			// 				on_success : function() {
			// 					$.gritter.removeAll();
			// 				}
			// 			},
			// 		    url: function(params) {
			// 				// ***UPDATE AVATAR HERE*** //
			// 				//for a working upload example you can replace the contents of this function with 
			// 				//examples/profile-avatar-update.js
			
			// 				var deferred = new $.Deferred
			
			// 				var value = $('#avatar').next().find('input[type=hidden]:eq(0)').val();
			// 				if(!value || value.length == 0) {
			// 					deferred.resolve();
			// 					return deferred.promise();
			// 				}
			
			
			// 				//dummy upload
			// 				setTimeout(function(){
			// 					if("FileReader" in window) {
			// 						//for browsers that have a thumbnail of selected image
			// 						var thumb = $('#avatar').next().find('img').data('thumb');
			// 						if(thumb) $('#avatar').get(0).src = thumb;
			// 					}
								
			// 					deferred.resolve({'status':'OK'});
			
			// 					if(last_gritter) $.gritter.remove(last_gritter);
			// 					last_gritter = $.gritter.add({
			// 						title: 'Avatar Updated!',
			// 						text: 'Uploading to server can be easily implemented. A working example is included with the template.',
			// 						class_name: 'gritter-info gritter-center'
			// 					});
								
			// 				 } , parseInt(Math.random() * 800 + 800))
			
			// 				return deferred.promise();
							
			// 				// ***END OF UPDATE AVATAR HERE*** //
			// 			},
						
			// 			success: function(response, newValue) {
			// 			}
			// 		})
			// 	}catch(e) {}
				
				
			
			// 	//another option is using modals
			// 	$('#avatar2').on('click', function(){
			// 		var modal = 
			// 		'<div class="modal fade">\
			// 		  <div class="modal-dialog">\
			// 		   <div class="modal-content">\
			// 			<div class="modal-header">\
			// 				<button type="button" class="close" data-dismiss="modal">&times;</button>\
			// 				<h4 class="blue">Change Avatar</h4>\
			// 			</div>\
			// 			\
			// 			<form class="no-margin">\
			// 			 <div class="modal-body">\
			// 				<div class="space-4"></div>\
			// 				<div style="width:75%;margin-left:12%;"><input type="file" name="file-input" /></div>\
			// 			 </div>\
			// 			\
			// 			 <div class="modal-footer center">\
			// 				<button type="submit" class="btn btn-sm btn-success"><i class="ace-icon fa fa-check"></i> Submit</button>\
			// 				<button type="button" class="btn btn-sm" data-dismiss="modal"><i class="ace-icon fa fa-times"></i> Cancel</button>\
			// 			 </div>\
			// 			</form>\
			// 		  </div>\
			// 		 </div>\
			// 		</div>';
					
					
			// 		var modal = $(modal);
			// 		modal.modal("show").on("hidden", function(){
			// 			modal.remove();
			// 		});
			
			// 		var working = false;
			
			// 		var form = modal.find('form:eq(0)');
			// 		var file = form.find('input[type=file]').eq(0);
			// 		file.ace_file_input({
			// 			style:'well',
			// 			btn_choose:'Click to choose new avatar',
			// 			btn_change:null,
			// 			no_icon:'ace-icon fa fa-picture-o',
			// 			thumbnail:'small',
			// 			before_remove: function() {
			// 				//don't remove/reset files while being uploaded
			// 				return !working;
			// 			},
			// 			allowExt: ['jpg', 'jpeg', 'png', 'gif'],
			// 			allowMime: ['image/jpg', 'image/jpeg', 'image/png', 'image/gif']
			// 		});
			
			// 		form.on('submit', function(){
			// 			if(!file.data('ace_input_files')) return false;
						
			// 			file.ace_file_input('disable');
			// 			form.find('button').attr('disabled', 'disabled');
			// 			form.find('.modal-body').append("<div class='center'><i class='ace-icon fa fa-spinner fa-spin bigger-150 orange'></i></div>");
						
			// 			var deferred = new $.Deferred;
			// 			working = true;
			// 			deferred.done(function() {
			// 				form.find('button').removeAttr('disabled');
			// 				form.find('input[type=file]').ace_file_input('enable');
			// 				form.find('.modal-body > :last-child').remove();
							
			// 				modal.modal("hide");
			
			// 				var thumb = file.next().find('img').data('thumb');
			// 				if(thumb) $('#avatar2').get(0).src = thumb;
			
			// 				working = false;
			// 			});
						
						
			// 			setTimeout(function(){
			// 				deferred.resolve();
			// 			} , parseInt(Math.random() * 800 + 800));
			
			// 			return false;
			// 		});
							
			// 	});
			
				
			
			// 	//////////////////////////////
			// 	$('#profile-feed-1').ace_scroll({
			// 		height: '250px',
			// 		mouseWheelLock: true,
			// 		alwaysVisible : true
			// 	});
			
			// 	$('a[ data-original-title]').tooltip();
			
			// 	$('.easy-pie-chart.percentage').each(function(){
			// 	var barColor = $(this).data('color') || '#555';
			// 	var trackColor = '#E2E2E2';
			// 	var size = parseInt($(this).data('size')) || 72;
			// 	$(this).easyPieChart({
			// 		barColor: barColor,
			// 		trackColor: trackColor,
			// 		scaleColor: false,
			// 		lineCap: 'butt',
			// 		lineWidth: parseInt(size/10),
			// 		animate:false,
			// 		size: size
			// 	}).css('color', barColor);
			// 	});
			  
			// 	///////////////////////////////////////////
			
			// 	//right & left position
			// 	//show the user info on right or left depending on its position
			// 	$('#user-profile-2 .memberdiv').on('mouseenter touchstart', function(){
			// 		var $this = $(this);
			// 		var $parent = $this.closest('.tab-pane');
			
			// 		var off1 = $parent.offset();
			// 		var w1 = $parent.width();
			
			// 		var off2 = $this.offset();
			// 		var w2 = $this.width();
			
			// 		var place = 'left';
			// 		if( parseInt(off2.left) < parseInt(off1.left) + parseInt(w1 / 2) ) place = 'right';
					
			// 		$this.find('.popover').removeClass('right left').addClass(place);
			// 	}).on('click', function(e) {
			// 		e.preventDefault();
			// 	});
			
		
			// 	$('#user-profile-3')
			// 	.find('input[type=file]').ace_file_input({
			// 		style:'well',
			// 		btn_choose:'Change avatar',
			// 		btn_change:null,
			// 		no_icon:'ace-icon fa fa-picture-o',
			// 		thumbnail:'large',
			// 		droppable:true,
					
			// 		allowExt: ['jpg', 'jpeg', 'png', 'gif'],
			// 		allowMime: ['image/jpg', 'image/jpeg', 'image/png', 'image/gif']
			// 	})
			// 	.end().find('button[type=reset]').on(ace.click_event, function(){
			// 		$('#user-profile-3 input[type=file]').ace_file_input('reset_input');
			// 	})
			// 	.end().find('.date-picker').datepicker().next().on(ace.click_event, function(){
			// 		$(this).prev().focus();
			// 	})
			// 	$('.input-mask-phone').mask('(999) 999-9999');
			
			
			
			// 	////////////////////
			// 	//change profile
			// 	$('[data-toggle="buttons"] .btn').on('click', function(e){
			// 		var target = $(this).find('input[type=radio]');
			// 		var which = parseInt(target.val());
			// 		$('.user-profile').parent().addClass('hide');
			// 		$('#user-profile-'+which).parent().removeClass('hide');
			// 	});
			// });
		</script>
	</body>
</html>
<script type="text/javascript">
	$(function() {
		// 详情中图片上传
		$('#avatar2').on('click', function(){
			var modal = 
			'<div class="modal fade">\
			  <div class="modal-dialog">\
			   <div class="modal-content">\
				<div class="modal-header">\
					<button type="button" class="close" data-dismiss="modal">&times;</button>\
					<h4 class="blue">跟换头像</h4>\
				</div>\
				\
				<form class="no-margin m-image" action="/admin/upload" method="post">\
				 <div class="modal-body">\
					<div class="space-4"></div>\
					<div style="width:75%;margin-left:12%;"><input type="file" name="avatar" /></div>\
				 </div>\
				\
				 <div class="modal-footer center">\
					<button type="submit" class="btn btn-sm btn-success"><i class="ace-icon fa fa-check"></i> 确定 </button>\
					<button type="button" class="btn btn-sm" data-dismiss="modal"><i class="ace-icon fa fa-times"></i> 取消 </button>\
				 </div>\
				</form>\
			  </div>\
			 </div>\
			</div>';
			var modal = $(modal);
			// 取消
			modal.modal("show").on("hidden", function(){ modal.remove();});
	
			var working = false,
				form 	= modal.find('form:eq(0)');
				file 	= form.find('input[type=file]').eq(0);

			// 图片上传
			file.ace_file_input({
				style:'well',
				btn_choose:'点击选择新的头像',
				btn_change:null,
				no_icon:'ace-icon fa fa-picture-o',
				thumbnail:'small',
				before_remove: function() {
					return !working;
				},

				// 允许上传的头像
				allowExt: ['jpg', 'jpeg', 'png', 'gif'],
				allowMime: ['image/jpg', 'image/jpeg', 'image/png', 'image/gif']
			});
	
			// 表单提交
			var ie_timeout = null;
			form.on('submit', function(){
				if( ! file.data('ace_input_files')) return false;	
				// return false;
				var $form = $(form);
				var file_input = file;
				var upload_in_progress = false;					
				var deferred ;
				if( "FormData" in window ) {
					formData_object = new FormData();
					$.each($form.serializeArray(), function(i, item) {
						formData_object.append(item.name, item.value);							
					});

					$form.find('input[type=file]').each(function(){
						var field_name = $(this).attr('name');
						var files = $(this).data('ace_input_files');
						if(files && files.length > 0) {
							for(var f = 0; f < files.length; f++) {
								formData_object.append(field_name, files[f]);
							}
						}
					});

					upload_in_progress = true;
					file_input.ace_file_input('loading', true);
					
					deferred = $.ajax({
						        url: $form.attr('action'),
						       type: $form.attr('method'),
						processData: false,//important
						contentType: false,//important
						   dataType: 'json',
						       data: formData_object
					})

				}
				else 
				{
					deferred = new $.Deferred 
					var temporary_iframe_id = 'temporary-iframe-'+(new Date()).getTime()+'-'+(parseInt(Math.random()*1000));
					var temp_iframe = 
							$('<iframe id="'+temporary_iframe_id+'" name="'+temporary_iframe_id+'" \
							frameborder="0" width="0" height="0" src="about:blank"\
							style="position:absolute; z-index:-1; visibility: hidden;"></iframe>')
							.insertAfter($form)

					$form.append('<input type="hidden" name="temporary-iframe-id" value="'+temporary_iframe_id+'" />');
					temp_iframe.data('deferrer' , deferred);
					$form.attr({
								  method: 'POST',
								 enctype: 'multipart/form-data',
								  target: temporary_iframe_id //important
								});

					upload_in_progress = true;
					file_input.ace_file_input('loading', true);
					$form.get(0).submit();
					ie_timeout = setTimeout(function(){
						ie_timeout = null;
						temp_iframe.attr('src', 'about:blank').remove();
						deferred.reject({'status':'fail', 'message':'Timeout!'});
					} , 30000);
				}

				deferred
				.done(function(result) {// 成功
					if (result.status == 1)
					{
						form.find('button').removeAttr('disabled');
						form.find('input[type=file]').ace_file_input('enable');
						form.find('.modal-body > :last-child').remove();
						modal.modal("hide");
						$('#avatar2').get(0).src = result.data.fileDir;
						working = false;
					} else {
						layer.msg(result.msg, {icon:5})
					}
					
				})
				.fail(function(result) {//failure
					layer.msg("页面没有响应");
				})
				.always(function() {//called on both success and failure
					if(ie_timeout) clearTimeout(ie_timeout)
					ie_timeout = null;
					upload_in_progress = false;
					file_input.ace_file_input('loading');
				});

				deferred.promise();
				return false;
			});
					
		});
	
		$('#user-profile-2 .memberdiv').on('mouseenter touchstart', function(){
			var $this = $(this);
			var $parent = $this.closest('.tab-pane');
			var off1 = $parent.offset();
			var w1 = $parent.width();
			var off2 = $this.offset();
			var w2 = $this.width();
			var place = 'left';
			if( parseInt(off2.left) < parseInt(off1.left) + parseInt(w1 / 2) ) place = 'right';
			$this.find('.popover').removeClass('right left').addClass(place);
		}).on('click', function(e) {
			e.preventDefault();
		});
	
		// 图片上传
		$('#user-profile-3')
		.find('input[type=file]').ace_file_input({
			style:'well',
			btn_choose:'Change avatar',
			btn_change:null,
			no_icon:'ace-icon fa fa-picture-o',
			thumbnail:'large',
			droppable:true,
			
			allowExt: ['jpg', 'jpeg', 'png', 'gif'],
			allowMime: ['image/jpg', 'image/jpeg', 'image/png', 'image/gif']
		})
		.end().find('button[type=reset]').on(ace.click_event, function(){
			$('#user-profile-3 input[type=file]').ace_file_input('reset_input');
		})
		.end().find('.date-picker').datepicker().next().on(ace.click_event, function(){
			$(this).prev().focus();
		})
		$('.input-mask-phone').mask('(999) 999-9999');
	
	
	
		////////////////////
		//change profile
		$('[data-toggle="buttons"] .btn').on('click', function(e){
			var target = $(this).find('input[type=radio]');
			var which = parseInt(target.val());
			$('.user-profile').parent().addClass('hide');
			$('#user-profile-'+which).parent().removeClass('hide');
		});
	});
</script>
