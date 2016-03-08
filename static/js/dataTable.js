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
	str += '<a class="btn btn-info" href="javascript:myTable.info('+row+');"><i class="icon-edit "></i></a> ';
	str += '<a class="btn btn-danger" href="javascript:myTable.delete('+row+');"><i class="icon-trash "></i></a>';
	$(td).html(str);
}

var MeTable = (function($) {
	var fnServerData = function(sSource, aoData, fnCallback) {
		// ajax请求
		$.ajax({
			url: sSource,
			data: aoData,
			type: 'post',
			dataType: 'json',
			success: function(data) {
				// 判断返回数据
				if (data.Status != 1) {
					layer.msg('出现错误:' + data.Message, {time:3000, icon:6});
					return false;
				}

				$.fn.dataTable.defaults['bFilter'] = true;
				fnCallback(data.Data);
			},
			error: function(msg) {
				console.log(msg);
				alert(msg);
			}
		});
	};

	// 构造函数初始化配置
	function MeTable(options, tableOptions) {
		// 表格信息配置
		this.tableOptions = {
			'bStateSave': true,
			"fnServerData": fnServerData,						// 获取数据的处理函数
			"sAjaxSource": "AjaxIndex",							// 获取数据地址
			"bLengthChange": true, 								// 是否可以调整分页
			"bAutoWidth": true,
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
			dialogId: "editDataDialog",
			tableId: "TableDataList",
		};

		this.tableOptions = $.extend(this.tableOptions, tableOptions);
		this.options = $.extend(this.options, options);
		this.actionType = "";
	}

	// 生成表格对象
	MeTable.prototype.Request = function() {
		this.table = $("#" + this.options.tableId).DataTable(this.tableOptions);
	};

	// 修改数据信息
	MeTable.prototype.edit = function(row) {
		var data = this.table.data()[row];
		var self = this;
		self.actionType = "update";
		this.tableOpt.aoColumns.forEach(function(v, k) {
			if (v.data == null || v.isUpdate === false) {
				return;
			}
			var obj = $('#' + self.opt.dialogId + ' *[name="' + v.data + '"].edit');
			self.setValue(obj, data[v.data]);
		});

		if (typeof self.opt.editBefore === "function") {
			if (self.opt.editBefore(data) === false) {
				return;
			}
		}

		$('#' + self.opt.dialogId).modal({
			backdrop: "static"
		});
	};

	// 关闭弹框
	MeTable.prototype.closeDialog = function() {
		$('#' + this.opt.dialogId).modal('hide');
	};

	// 表格搜索
	MeTable.prototype.search = function() {
		this.table.draw(true);
	};

	// 表格数据的添加
	MeTable.prototype.insert = function() {
		this.clearData();
		this.actionType = "add";
		if (typeof this.opt.addBefore === "function") {
			if (this.opt.addBefore() === false) {
				return;
			}
		}
		$('#' + this.opt.dialogId).modal({
			backdrop: "static"
		});
	};

	// 删除数据
	MeTable.prototype.delete = function(row) {
		var data = this.table.data()[row], self = this;
		data.operateType = "delete";
		// 询问框
		layer.confirm('您确定需要删除这条数据吗?', {
			title:'确认操作',
			btn:['确定','取消'],
			shift:4
			// 确认删除
		}, function(){
			$.ajax({
				url: self.options.baseUrl,
				type: "post",
				data:data,
				success: function(data) {
					if (data.success === true) {
						self.table.draw(false);
						alert(data.data);
					} else {
						alert(data.err);
					}
				}
			});

			// 取消删除
		}, function(){
			layer.msg('您取消了删除操作！', {time:1000});
		});
	};

	// 数据新增和修改的执行
	MeTable.prototype.saveData = function() {
		var data = {};
		var self = this;
		if (self.actionType == "") return;
		var url = self.options.baseUrl;
		$.ajax({
			url: url,
			data: JSON.stringify(data),
			type: "post",
			dataType: 'json',
			success: function(data) {
				if (data.success === true) {
					if (typeof self.options.saveAfter === "function") {
						if (self.options.saveAfter(data) === false) return;
					}
					self.closeDialog();
					if (self.table) self.table.draw(false);
					alert(data.data);
				} else {
					alert(data.err);
				}
			}
		});

		self.actionType = "";
	};

	return MeTable;
})($);
