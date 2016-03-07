/**
 * stringTo 值的转换
 * @param  string type  值的类型
 * @param  mixed  value 值
 * @return mixed  返回处理后的值
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
			value = parseInt(value);
		case "bool":
			value = value === "true" || value === true || value === 1 || value === "1";
		case "float32":
		case "float64":
	}

  	return value;
}

/**
 * Format() 将时间转换为字符串
 * @param string fmt 转换的格式
 */
Date.prototype.Format = function(fmt) {
	var o = {
		"M+": this.getMonth() + 1, 						// 月份
		"d+": this.getDate(), 							// 日
		"h+": this.getHours(), 							// 小时
		"m+": this.getMinutes(), 						// 分
		"s+": this.getSeconds(), 						// 秒
		"q+": Math.floor((this.getMonth() + 3) / 3), 	// 季度
		"S": this.getMilliseconds() 					// 毫秒
	};

	// 判断年份
  	if (/(y+)/.test(fmt))
  	{
  		fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
  	} 

  	// 其他字符串替换
  	for (var k in o)
  	{
    	if (new RegExp("(" + k + ")").test(fmt))
    	{
    		fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
    	} 
  	}
  	
  	return fmt;
}

/**
 * timeToString() 时间戳转换为字符串
 * @param  int time 时间戳
 * @return string   返回字符串
 */
function timeToString(time) {
	var date = new Date(time*1000)
	return date.Format("yyyy-MM-dd hh:mm:ss")
}

// 我的datatable处理对象
var MeTable = (function(){

	// 处理回调函数
	var fnServerData = function(sSource, aoData, fnCallback) {
		$.ajax({
			url: sSource,
			data: args,
			type: 'post',
			dataType: 'json',
			success: function(data) {
				if (data.Result != 1) 
				{
					layer.closeAll();
					layer.msg('异常:' + data.ResultString);
				}

				$.fn.dataTable.defaults['bFilter'] = true;
				fnCallback(data.ResultObject);
			},
			error: function(msg) {
				console.log(msg);
				alert(msg);
			}
    	});
	}

	/**
	 * MeTable() 构造函数,初始化操作
	 * @param obj options      
	 * @param obj tableOptions 
	 */
	function MeTable(options, tableOptions) {
		// table表格配置信息
		this.tableOptions = {
			"processing": true,
			"serverSide": true,
			'bStateSave': true,
			"sAjaxSource": "AjaxIndex", 	//	这个是请求的地址
			"fnServerData": fnServerData 	// 	获取数据的处理函数
		};

		// 自己的配置信息
		this.options = {
			dialogId:"editDataDialog",
			tableId:"showDataTable"
		};

		// 继承处理
		this.tableOptions = $.extend(this.tableOptions, tableOptions);
		this.options = $.extend(this.options, options);

		// 定义操作类型
		this.actionType = "";
	}

	// 初始化请求数据
	MeTable.prototype.Request = function() {
		this.table = $("#" + this.options.tableId).DataTable(this.tableOptions)
	}

	// 关闭弹窗

	// 修改数据
	MeTable.prototype.edit = function(row) {
		var data = this.table.data()[row];
		this.actionType = "update";

		// 根据字段进行表单的处理
		this.tableOptions.aoColumns.forEach(function(v, k){
			if (v.data == null || v.isUpdate === false) 
			{
				return;
			}

			// 处理表单赋值
			var obj = $("#" + this.options.dialogId + ' *[name="' + v.data + '"]')
			this.saveValue(obj, data[v.data]);
		});

		// 判断是否有修改之前的处理
		if (typeof this.options.editBefore === "function") 
		{
			if (this.options.editBefore(data) == false) return 
		}

		// 弹出修改表单信息
		$('#' + this.options.dialogId).modal({
	      	backdrop: "static"
	    });
	};

	// 添加数据
	MeTable.prototype.add = function() {
		this.clearData();
		this.actionType = "insert"

		// 判断之前是否有操作
		if (typeof this.options.addBefore === "function") {
			if (this.options.addBefore() === false) {
				return;
			}
	    }

	    // 显示表单
	    $('#' + this.options.dialogId).modal({
	      backdrop: "static"
	    });
	};

	// 删除数据
	MeTable.prototype.delete = function(row) {
	    var data = this.table.data()[row];
	    if ( ! confirm("是否要删除Id=" + data[this.options.primaryKey])) return;
		$.ajax({
			url: this.options.baseUrl + "Del?id=" + data[this.options.primaryKey],
			type: "post",
			success: function(data) {
				if (data.success === true) {
					this.oTable.draw(false);
					alert(data.data);
				} else {
					alert(data.err);
				}
			}
		});
	    console.log("删除", data[self.options.primaryKey]);
  	};

  	// 修改和删除数据
  	MeTable.prototype.saveData = function() {
  		// 为空直接返回
  		if (this.actionType == "") return;
  		var self = this;
    
		$.ajax({
			url: url,
			data: JSON.stringify(data),
			type: "post",
			dataType: 'json',
			success: function(data) {
				self.closeDialog();
				if (self.table)
				{
					self.table.draw(false);
					alert(data.data);
				}
				else 
				{
					alert(data.err);
				}
			}
		});

		// 重新设置类型
		self.actionType = "";
	};

})($);