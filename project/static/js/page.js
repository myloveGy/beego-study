var MePage = (function($){
	// 初始化函数
	function MePage(options)
	{
		this.options = {
			ajaxUrl  : '?',				 // 请求的地址
			ajaxType : 'GET',			 // 请求的方式
			html 	 : '<ul>{html}<ul>', // 分页外围HTML
			pageNum  : 5,				 // 数字分页条数
			pageSize : 10,				 // 分页长度
			pageHtml : {				 // 其他配置
				first : '首页',
				prev  : '上一页',
				next  : '下一页',
				last  : '尾页'
			},
		};

		this.options = $.extend(this.options, options);
		this.Params = {
			pageCurr : 0, 				 // 默认当前页
			intStats : 0,				 // 开始位置
			intEnd 	 : 10,				 // 结束位置
		};
	}

	// 请求页面获取数据
	MePage.prototype.ajaxUrl = function(){
		$.ajax({
			url  : this.options.ajaxUrl,
			type : this.options.ajaxType,
			data : this.Params,
			success:function(data) {

			},
			error: function(){

			}
		});
	};

	MePage.prototype.CreatePage = function() {
		var html = '';

	};

	MePage.prototype.init = function(){
		this.ajaxUrl();
	}

	return MePage;
})($); 