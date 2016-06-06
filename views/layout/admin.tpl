<!DOCTYPE html>
<html lang="zh">
<head>
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
    <meta charset="utf-8" />
    <title>我的GO后台</title>

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


    <!-- 公共的JS文件 -->
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

    <script src="/static/assets/js/ace-elements.min.js"></script>
    <script src="/static/assets/js/ace.min.js"></script>
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

    <!--主要内容信息-->
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

            <!--主要内容信息-->
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
                        {{.LayoutContent}}
                    </div>
                </div>
            </div>

        </div>
    </div>

    <!--尾部信息-->
    <div class="footer">
        <div class="footer-inner">
            <div class="footer-content">
                <span class="bigger-120">
                    <span class="blue bolder"> Liujinxing </span>
                    个人 GO WEB 项目 &copy; 2016-2018
                </span>
            </div>
        </div>
    </div>
    <a href="#" id="btn-scroll-up" class="btn-scroll-up btn btn-sm btn-inverse">
        <i class="ace-icon fa fa-angle-double-up icon-only bigger-110"></i>
    </a>
</div>
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
<script src="/static/js/layer/layer.js"></script>
</body>
</html>
<script type="text/javascript">
    $(function(){
        var select = '.admin-site';
        // 导航栏样式装换
        $(select).addClass('active').parentsUntil('ul.nav-list').addClass('active open');
        // 隐藏和显示
        $('a[data-action=close]:first').click(function(){
            $(select).children('a').append('<span class="badge badge-primary tooltip-error" title="显示">显示</span>').bind('click', function (e) {
                e.preventDefault();
                $('div.widget-box:first').fadeIn();
                $(this).unbind('click').find('span:last').remove();
                return false;
            });;
        })
    })
</script>