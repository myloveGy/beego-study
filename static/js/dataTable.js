/**
 * stringTo() 值的转换
 * @param  string   type  需要转换成的值类型
 * @param  mixed    value 值
 * @return mixed    返回处理值
 */
function stringTo(type, value) {
	switch (type) 
	{
		case "int":
		case "int8":
		case "int16":
		case "int32":
		case "int64":
		case "uint":
		case "uint8":
		case "uint16":
		case "uint32":
		case "uint64":
			return parseInt(value);
		case "bool":
			return value === "true" || value === true || value === 1 || value == "1";
		case "float32":
		case "float64":
	}

	return value;
};

/**
 * Date.Format() 时间处理为字符串
 * @param string fmt 处理的格式
 */
Date.prototype.Format = function(fmt) {
	// 定义处理格式
	var o = {
		"M+": this.getMonth() + 1, 					 // 月份
		"d+": this.getDate(), 						 // 日
		"h+": this.getHours(), 						 // 小时
		"m+": this.getMinutes(), 					 // 分
		"s+": this.getSeconds(), 					 // 秒
		"q+": Math.floor((this.getMonth() + 3) / 3), // 季度
		"S": this.getMilliseconds() 				 // 毫秒
	};

	// 处理年份
	if (/(y+)/.test(fmt))
	{
		fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
	}

	// 处理其他信息
	for (var k in o)
	{
		if (new RegExp("(" + k + ")").test(fmt))
		{
			fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
		}
	}
    
  	return fmt;
};

/**
 * timeToString 时间戳转换为字符串
 * @param  int 		time 	时间格式
 * @return string   返回格式化字符串时间信息
 */
function timeToString(time) {
	var date = new Date(time * 1000);
	return date.Format("yyyy-MM-dd hh:mm:ss");
};

// 状态信息
function statusToString(td, data, rowdatas, row, col) {
	var str = '<span class="label label-' + (data == 1 ? 'success">启用' : 'important">停用') + '</span>';
	$(td).html(str);	
}

// 时间戳列，值转换
function dateTimeString(td, cellData, rowData, row, col)
{
	$(td).html(timeToString(cellData));
}

// 设置表单信息
function setOperate(td, data, rowdata, row, col)
{
	var str = '<a class="btn btn-success" href="javascript:myTable.info('+row+');"><i class="icon-zoom-in "></i></a> ';
	str += '<a class="btn btn-info" href="javascript:myTable.update('+row+');"><i class="icon-edit "></i></a> ';
	str += '<a class="btn btn-danger" href="javascript:myTable.delete('+row+');"><i class="icon-trash "></i></a>';
	$(td).html(str);
}

var MeTable = (function($) {
	var fnServerData = function(sSource, aoData, fnCallback) {

		var intLayer = layer.load();

		// 请求前5个参数(1, 2)有用 1记录有多少个列 每个列有5个字段 后5个字段有用()
		var attributes = aoData[2].value.split(","), obj = [],len = attributes.length + 1, mSkey = len * 5, mSort = mSkey + 2;
		for (var i in attributes)
		{
			var key = 6 + i * 5,tmpData = aoData[key];
			if (tmpData.value != undefined && tmpData.value != "" && tmpData.value != "全部")
			{
				var mKey = attributes[i]
				obj.push({"name":mKey, "value": tmpData.value});
			}
		}

		// 添加快速查询
		if (aoData[mSkey].value != undefined && aoData[mSkey].value != "") 
		{
			obj.push({"name":"search", "value": aoData[mSkey].value});
		}

		// 添加排序字段信息
		if (aoData[mSort].value != undefined && aoData[mSort].value != "") 
		{
			var tmpkey = parseInt(aoData[mSort].value)
			obj.push({"name":"orderBy", "value": attributes[tmpkey]});
		}

		// 查询数据使用json格式传输
		aoData.push({"name":"msearch", "value":JSON.stringify(obj)})
		
		// ajax请求
		$.ajax({
			url: sSource,
			data: aoData,
			type: 'post',
			dataType: 'json',
			cache:false,
			success: function(data) {
				layer.close(intLayer)
				// 判断返回数据
				if (data.Status != 1) 
				{
					layer.msg('出现错误:' + data.Message, {time:1000, icon:6});
					return false;
				}

				$.fn.dataTable.defaults['bFilter'] = true;
				fnCallback(data.Data);
			},
			error: function(msg) {
				layer.close(intLayer);
				layer.msg("服务器繁忙,请稍候再试...", {time:1000});
			}
		});
	};

	// 构造函数初始化配置
	function MeTable(options, tableOptions) {
		// 表格信息配置
		this.tableOptions = {
			'bStateSave': true,
			"fnServerData": fnServerData,						// 获取数据的处理函数
			"sAjaxSource": "ajaxindex",							// 获取数据地址
			"bLengthChange": true, 								// 是否可以调整分页
			"bAutoWidth": false,
	        "bPaginate": true,
	        "iDisplayStart": 0,
	        "iDisplayLength": 10,
	        "bServerSide": true,
	        "bRetrieve": true,
	        "bDestroy": true,
	        "sPaginationType":"full_numbers",
			"oLanguage": {
			    // 显示
			    "sLengthMenu": "每页 _MENU_ 条记录",
			    "sZeroRecords": "没有找到记录",
			    "sInfo": "显示 _START_ 到 _END_ 共有 _TOTAL_ 条数据",
			    "sInfoEmpty": "无记录",
			    "sInfoFiltered": "(从 _MAX_ 条记录过滤)",
			    "sSearch": "搜索：",

			    // 分页
			    "oPaginate": {
			        "sFirst": "首页",
			        "sPrevious": "上一页",
			        "sNext": "下一页",
			        "sLast": "尾页"
				}	
			}
		};

		// 自定义信息配置
		this.options = {
			dialogId: "myModal",
			tableId: "TableDataList",
			formId: "updateForm",
		};

		this.tableOptions = $.extend(this.tableOptions, tableOptions);
		this.options = $.extend(this.options, options);
		this.actionType = "";
	}

	// 处理表单信息
	MeTable.prototype.CreateForm = function() {
		var attributes = this.tableOptions.aoColumns, self = this, form = "", search = "";
		form += '<form class="form-horizontal" id="'+this.options.formId+'" name="'+this.options.formId+'" action="'+this.options.baseUrl +'" method="post"></fieldset>';
		// 处理生成表单
		attributes.forEach(function(k, v) {
			// 处理搜索
			if (k.search != undefined)
			{
				switch (k.search.type)
				{
					case "select":
						search += '<label> ' + k.title + ' : <select name="' + k.data +'" class="msearch" vid="' + v + '">';
						search += '<option> 全部 </option>'
						if (k.edit.value != undefined)
						{
							for (var x in k.edit.value)
							{
								search += '<option value="' + x + '">' + k.edit.value[x] + '</option>';
							}
						}

						search += '</select></label> ';
					break;
					default:
						search += '<label> '+ k.title +': <input type="text" name="' + k.data + '" vid="' + v + '" class="msearch" value="" /></label> ';
				}
				
			}

			// 处理编辑
			if (k.edit != undefined) 
			{
				// 处理其他参数
				var other = ' name="'+ k.data +'" ';
				if (k.edit.options != undefined)
				{
					for (var i in k.edit.options) other += i + '="' + k.edit.options[i] + '" '
				}

				if ( k.edit.type == "hidden" ) 
				{
					form += '<input type="hidden" ' + other + '/>';
				}
				else 
				{
					form += '<div class="control-group">';
					form += '	<label class="control-label">' + k.title + '</label>';
					form += '	<div class="controls">';

					// 判断类型
					switch (k.edit.type)
					{
						case "text":
							form += '<input class="input-xlarge focused" type="text" ' + other +' />';
							break;
						case "radio":
							if (k.edit.value != undefined)
							{
								for (var x in k.edit.value)
								{
									other += x == k.edit.default ? ' checked="checked" ' : '';
									form += "<label>";
									form += '<input class="input-xlarge focused" type="radio" '+ other +' value="' + x + '" />';
									form += '<span> ' + k.edit.value[x] + ' </span>';
									form += "<label>";
								}
							}
							break;
						case "select":
							form += '<select ' + other +'>';
							form += '<option> -- 请选择 -- </option>'
							if (k.edit.value != undefined)
							{
								for (var x in k.edit.value)
								{
									var selected = x == k.edit.default ? ' selected="selected" ' : '';
									form += '<option value="' + x + '" ' + selected + '>' + k.edit.value[x] + '</option>';
								}
							}

							form += '</select>';
							break;	
						default:	
							form += '<input class="input-xlarge focused" type="text" '+ other +' />';
					}
				}

				
				
				form += '	</div>';
				form += '</div>';
			}
		});

		form += '</fieldset></form>';
		this.SearchHtml = search;
		$('#' + this.options.dialogId).find('div.modal-body').html(form);
	};

	// 生成表格对象
	MeTable.prototype.init = function() {
		this.CreateForm();
		this.options.title = $('h2').text();
		this.table = $("#" + this.options.tableId).DataTable(this.tableOptions);
		var self = this;

		// 分页样式
        $('#showTable_paginate').addClass('pagination pagination-centered');
        // 去掉搜索placeholder属性(不去掉白色看不见)
        $('input[type=search]').removeAttr('placeholder');
        // 添加搜索信息
        $('#showTable_filter').append(self.SearchHtml);
        // 表格添加搜索事件
        $('input.msearch').live('keyup', function () {
            self.table.column(parseInt($(this).attr('vid'))).search($(this).val()).draw();
        });
        $('select.msearch').live('change', function () {
            self.table.column(parseInt($(this).attr('vid'))).search($(this).val()).draw();
        });
	};

	// 表格搜索
	MeTable.prototype.search = function() {
		this.table.draw(true);
	};

	// 初始化表单对象
	MeTable.prototype.initForm = function(data)
	{
		// 弹出标题显示
		var title = this.options.title + (this.actionType == "insert" ? "新增" : "编辑");
		$("#" + this.options.dialogId).find('h3').html(title);

		// 表单处理
		objForm = document.updateForm
		if (objForm != undefined)
		{
			objForm.reset();
			if (data != undefined)
			{
				for (var i in data)
				{
					if (objForm[i] != undefined) objForm[i].value = data[i];
				}
			}
		}

		// 弹出表单信息
		$('#' + this.options.dialogId).modal({
			backdrop: "static"
		});
	}

	// 初始化

	// 表格数据的添加
	MeTable.prototype.insert = function() {
		this.actionType = "insert";
		this.initForm();
	};

	// 修改数据信息
	MeTable.prototype.update = function(row) {
		this.actionType = "update";

		// 初始化表单
		this.initForm(this.table.data()[row])
	};

	// 删除数据
	MeTable.prototype.delete = function(row) {
		var data = this.table.data()[row], self = this;
		this.actionType = "delete";
		// 询问框
		layer.confirm('您确定需要删除这条数据吗?', {
			title: '确认操作',
			btn: ['确定','取消'],
			shift: 4,
			icon: 0
			// 确认删除
		}, function(){
			self.saveData(data);
			// 取消删除
		}, function(){
			layer.msg('您取消了删除操作！', {time:800});
		});
	};

	// 数据新增和修改的执行
	MeTable.prototype.saveData = function(data) {
		layer.closeAll();
		// 判断类型
		if (this.actionType == "") return;

		// 新增和修改验证数据
		if (this.actionType !== "delete")
		{
			if ( ! $("#" + this.options.formId).validate().form()) return false;
			// 提交数据
			data = $('#' + this.options.formId).serialize();
			data += "&actionType=" + this.actionType;
		}
		else
		{
			data.actionType = this.actionType;
		}

		var self = this, intLoad = layer.load();
		// ajax提交数据
		$.ajax({
			url:self.options.baseUrl,
			type:'POST',
			data:data,
			dataType:'json',
			success:function(json)
			{
				layer.close(intLoad);

				// 提示信息
				var intIcon = json.Status == 1 ? 6 : 5;
				layer.msg(json.Message, {time:1000, icon:intIcon})

				// 判断操作成功
				if (json.Status == 1)
				{
					self.table.draw(false);
					if (self.table.actionType !== "delete") $("#" + self.options.dialogId).modal('hide');
					self.table.actionType = "";
					return false;
				}
			},
			error:function(){
				layer.close(intLoad);
				layer.msg("服务器繁忙,请稍候再试...", {time:1000, icon:2})
			}
		});

		return false;
	};

	return MeTable;
})($);