// 状态信息
function statusToString(td, data, rowdatas, row, col) {
	var str = '<span class="label label-' + (data == 1 ? 'success">启用' : 'warning">禁用') + '</span>';
	$(td).html(str);
}

// 时间戳列，值转换
function dateTimeString(td, cellData, rowData, row, col)
{
	$(td).html(timeFormat(cellData, 'yyyy-MM-dd hh:mm:ss'));
}

// 推荐信息
function recommendToString(td, data, rowdata, row, col)
{
    var str = '<span class="label label-' + (data == 1 ? 'success">推荐' : 'warning">不推荐') + '</span>';
    $(td).html(str);
}

// 图片显示
function stringToImage(td, data, rowdatas, row, col)
{
    if (!empty(data)) {
        var alt = empty(rowdatas) ? '图片详情信息' : rowdatas.Title;
        $(td).html('<img width="100px" layer-src="' + data + '" src="' + data + '" alt="' + alt + '" onclick="myTable.seeImage(' + row + ');" />')
    }
}

// 设置表单信息
function setOperate(td, data, rowdata, row, col)
{
	var str = '<a class="btn btn-success btn-xs" href="javascript:;"  onclick="myTable.views('+row+');"><i class="glyphicon glyphicon-zoom-in"></i></a> ';
	str += '<a class="btn btn-info btn-xs" href="javascript:;" onclick="myTable.update('+row+');"><i class="glyphicon glyphicon-edit "></i></a> ';
	str += '<a class="btn btn-danger btn-xs" href="javascript:;" onclick="myTable.delete('+row+');"><i class="glyphicon glyphicon-trash "></i></a>';
	$(td).html(str);
}

var loadIndex = 0;
var isSearch  = false;
function fnServerData(sSource, aoData, fnCallback) {
    var intLayer = layer.load(),
    // 请求前5个参数(1, 2)有用 1记录有多少个列 每个列有5个字段 后5个字段有用()
        attributes = aoData[2].value.split(","), obj = [],len = attributes.length + 1, mSkey = len * 5, mSort = mSkey + 2;
    if ( ! isSearch)
    {
        for (var i in attributes)
        {
            var key = 6 + i * 5,tmpData = aoData[key];
            if (tmpData.value != undefined && tmpData.value != "" && tmpData.value != "All")
            {
                var mKey = attributes[i]
                aoData.push({"name":'params['+mKey+']', "value": tmpData.value});
            }
        }

        // 添加快速查询
        if (aoData[mSkey].value != undefined && aoData[mSkey].value != "")
        {
            aoData.push({"name":'params[search]', "value": aoData[mSkey].value});
        }

        // 添加排序字段信息
        if (aoData[mSort].value != undefined && aoData[mSort].value != "")
        {
            var tmpkey = parseInt(aoData[mSort].value);
            aoData.push({"name":'params[orderBy]', "value": attributes[tmpkey]});
        }
    } else {
        $('.msearch').each(function(){
            var v = $(this).val();
            if ( ! empty(v) && v != 'All') aoData.push({"name":'params[' + $(this).attr('name') + ']', "value": $(this).val()});
        });
    }

    // ajax请求
    $.ajax({
        url: sSource,
        data: aoData,
        type: 'post',
        dataType: 'json',
        cache: false,
        success: function(data) {
            layer.close(intLayer)
            // 判断返回数据
            if (data.status != 1)
            {
                layer.msg('出现错误:' + data.msg, {time:2000, icon:5});
                return false;
            }

            $.fn.dataTable.defaults['bFilter'] = true;
            fnCallback(data.data);
        },
        error: function() {
            layer.close(intLayer);
            layer.msg("服务器繁忙,请稍候再试...", {time:2000});
        }
    });
};

var MeTable = (function($) {
	// 构造函数初始化配置
	function MeTable(options, tableOptions) {
        isSearch = options.searchPosition == 'top';
		// 表格信息配置
		this.tableOptions = {
			'bStateSave': true,
			"fnServerData": fnServerData,						// 获取数据的处理函数
			"sAjaxSource": "search",							// 获取数据地址
			"bLengthChange": true, 								// 是否可以调整分页
			"bAutoWidth": false,
            "bPaginate": true,
            "iDisplayStart": 0,
            "iDisplayLength": 10,
            "bServerSide": true,
            "bRetrieve": true,
            "bDestroy": true,
            "processing": true,
            "serverSide": true,
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
            },
            "order":[[0, "desc"]],
		};

		// 自定义信息配置
		this.options = {
			dialogId : "myModal",           // 编辑 Modal ID
			tableId  : "showTable",         // 显示表格 ID
			formId   : "updateForm",        // 修改表单 ID
			baseUrl  : "update",            // 修改提交地址
			isHaveError: false,             // 是否存在错误
            searchPosition: "middle",       // 搜索框出现位置
		};

		this.tableOptions = $.extend(this.tableOptions, tableOptions);
		this.options 	  = $.extend(this.options, options);
		this.formOptions  = $.extend({
			"method": 	"post", 
			"id": 		this.options.formId, 
			"class": 	"form-horizontal",
			"name": 	this.options.formId,
			"action":   this.options.baseUrl, 
		}, this.options.formOptions);
		this.actionType   = "";
	}

	// 处理表单信息
	MeTable.prototype.CreateForm = function() {
		var attributes = this.tableOptions.aoColumns, self = this, form = "", search = "", views = "", formParams = handleParams(this.formOptions);
		form += '<form ' + formParams + '></fieldset>';
		views += '<table class="table table-bordered table-striped table-detail">';
		// 处理生成表单
		attributes.forEach(function(k, v) {
			// 初始化详情信息
			views += '<tr>';
			views += '<td width="25%">' + k.title + '</td><td class="views-info data-info-' + k.data + '"></td>';
			views += '</tr>';
			
			// 处理搜索
			if (k.search != undefined)
			{
				var tmpOptions = {"name":k.sName, "vid":v, "class":"msearch"},html = '';
                if (k.search.options) $.extend(tmpOptions, k.search.options);
                if ( self.options.searchPosition == 'top') tmpOptions['placeholder'] = '请输入' + k.title;
				switch (k.search.type)
				{
					case "select":
						k.value["All"] = "全部"; 
						html += createSelect(k.value, "All", tmpOptions)
						delete k.value['All']
					break;
					default:
						html += createInput('text', tmpOptions);
				}

				search += Label(k.title + " : " + html) + ' ';
				
			}

			// 处理编辑
			if (k.edit != undefined) 
			{
				// 处理其他参数
				if (k.edit.options == undefined) k.edit.options = {};
				k.edit.options["name"]  = k.sName;
				k.edit.options["class"] = "form-control";
				if (k.edit.type == undefined) k.edit.type = "text"

				if ( k.edit.type == "hidden" ) 
				{
					form += createInput('hidden', k.edit.options)
				}
				else 
				{

					form += '<div class="form-group">' + Label(k.title, {"class":"col-sm-3 control-label"}) + '<div class="col-sm-9">';

					// 判断类型
					switch (k.edit.type)
					{
                        // 单选
						case "radio":
							k.edit.options['class'] = 'ace valid';
							form += createRadio(k.value, k.edit.default, k.edit.options);
							break;
                        // 多选
                        case "checkbox":
                            k.edit.options['class'] = 'ace m-checkbox';
                            k.edit.options['name']  = k.sName + '[]';
                            form += createCheckbox(k.value, k.edit.default, k.edit.options);
                            break;
                        // 下拉
						case "select":
							form += createSelect(k.value, k.edit.default, k.edit.options);
							break;
                        // 文件上传
						case "file":
							form += createFile(k.edit.options);
							break;
                        // 文本
                        case "textarea":
                            form += createTextarea(k.edit.options); //'<textarea class="form-control" rows="3" ></textarea>';
                            break;
						// 多语言
						case 'lang':
							form += createLangInput(k.edit.options);
							break;
                        // 时间
                        case "time":
                            if (!empty(k.value)) k.edit.options["value"] = k.value
                            k.edit.options["class"] += " time";
                            form += '<div class="col-sm-6 m-pl-0">' + createInput('text', k.edit.options) + '</div>';
                            break;
                        // 输入框
						default:
							if (!empty(k.value)) k.edit.options["value"] = k.value	
							form += createInput(k.edit.type, k.edit.options);
					}
				}

				form += '</div></div>';
			}
		});

		form += '</fieldset></form>';
		views += '</table>';
		this.SearchHtml = search;

		$('#' + this.options.dialogId).find('div.modal-body').html(form);
		$("#data-info").html(views);
	};

	// 生成表格对象
	MeTable.prototype.init = function() {
		this.CreateForm();
		this.options.title = $('h1').text();
		this.table = $("#" + this.options.tableId).DataTable(this.tableOptions);
		var self = this;

        // 判断初始化处理(搜索添加位置)
        if (this.options.searchPosition == 'middle')
        {
            $('#showTable_filter').append(self.SearchHtml);
            // 表格添加搜索事件
            $('.msearch').on('keyup change', function () {
                self.table.column(parseInt($(this).attr('vid'))).search($(this).val()).draw();
            });

            // 处理搜索信息
            $('#showTable_wrapper div.row div.col-xs-6:first').removeClass('col-xs-6').addClass('col-xs-2');
            $('#showTable_wrapper div.row div.col-xs-6:first').removeClass('col-xs-6').addClass('col-xs-10');
        }
        else
        {
            // 在表单中添加搜索信息
            $('#' + this.options.searchFormId).prepend(self.SearchHtml);
        }
	};

	// 表格搜索
	MeTable.prototype.search = function() {
		this.table.draw();
	};

	// 初始化表单对象
	MeTable.prototype.initForm = function(data)
	{
		layer.closeAll();
		// 弹出标题显示
		var title = this.options.title + (this.actionType == "insert" ? "新增" : "编辑");
		$("#" + this.options.dialogId).find('h4').html(title);

		// 表单处理
		objForm = document.updateForm
		if (objForm != undefined)
		{
			$(objForm).find('input[type=hidden]').val('');                                  // 隐藏按钮充值
            $(objForm).find('input[type=checkbox]').each(function(){
                $(this).attr('checked', false);
                if ($(this).get(0)) $(this).get(0).checked = false;
            });                                                                             // 多选菜单
			objForm.reset();                                                                // 表单重置
			if (data != undefined)
			{
				for (var i in data)
				{
                    // 多语言处理 以及多选配置
                    if (typeof data[i]  ==  'object')
                    {
                        for (var x in data[i])
                        {
                            var key = i + '[' + x + ']';
                            // 对语言
                            if (objForm[key] != undefined)
								objForm[key].value = data[i][x];
                            else {
                                // 多选按钮
								if (parseInt(data[i][x]) > 0) {
									$('input[type=checkbox][value=' + data[i][x] + ']').attr('checked', true).each(function(){this.checked=true});
								}
                            }
                        }
                    }

                    // 其他除密码的以外的数据
					if (objForm[i] != undefined && objForm[i].type != "password")
					{
                        var obj = $(objForm[i]), tmp = data[i];
                        // 时间处理
                        if (obj.hasClass('time')) tmp = timeFormat(parseInt(tmp), 'yyyy/MM/dd hh:mm:ss');
						objForm[i].value = tmp;
					}
				}
			}
		}

		// 弹出表单信息
		$('#' + this.options.dialogId).modal({backdrop: "static"});
	}

	// 查询详情
	MeTable.prototype.views = function(row) {
        if (loadIndex != 0) return false;
		var self = this, data = this.table.data()[row];
		// 处理的数据
		if (data != undefined)
		{
			// 循环处理显示信息
			this.tableOptions.aoColumns.forEach(function(k, v) {
				var tmpKey   = k.data,tmpValue = data[tmpKey],dataInfo = $('.data-info-' + tmpKey)
				if (k.edit != undefined && k.edit.type == 'password') tmpValue = "******";
				// 赋值
				if (k.createdCell != undefined && typeof k.createdCell == "function")
					k.createdCell(dataInfo, tmpValue, data, row, undefined)
				else 
					dataInfo.html(tmpValue)
			});

			// 弹出显示
			loadIndex = layer.open({
			    type: 1,
			    shade: 0.3,
                shadeClose:true,
			    title: self.options.title + '详情',
			    content: $('#data-info'), 		// 捕获的元素
			    area:['50%', 'auto'],
			    cancel: function(index){layer.close(index);},
                end:function(){
                    $('.views-info').html('')
                    loadIndex = 0;
                }
			});
		}

	}

    // 新增之前的处理
    MeTable.prototype.insertShow = function(){
        return true;
    };

	// 表格数据的添加
	MeTable.prototype.insert = function() {
		this.actionType = "insert";
        this.insertShow();
		this.initForm();
	};

    // 数据修改之前的处理
    MeTable.prototype.updateShow = function(obj) {
        return true;
    };

	// 修改数据信息
	MeTable.prototype.update = function(row) {
		this.actionType = "update";
        this.updateShow(this.table.data()[row]);
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

	// 修改数据之后的处理
    MeTable.prototype.beforeSave = function(){
        return true;
    }

	// 数据新增和修改的执行
	MeTable.prototype.saveData = function(data) {
		layer.closeAll();
		var self = this;
		// 判断类型
		if (this.actionType == "") return;

		// 新增和修改验证数据
		if (this.actionType !== "delete")
		{
			var isHave = false;
			if ( ! $("#" + this.options.formId).validate({
				errorPlacement:function(error, errorPlacement) {
					if (isHave != false) return false;
					isHave = layer.tips($(error).html(), errorPlacement, {tips: [3], time:1000, end:function(){
						isHave = false;
					}})
				},
			}).form()) return false;
			// 提交数据
			data = $('#' + this.options.formId).serialize();
			data += "&actionType=" + this.actionType;
		}
		else
		{
			data.actionType = this.actionType;
		}

		var intLoad = layer.load();
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
				var intIcon = json.status == 1 ? 6 : 5;
				layer.msg(json.msg, {time:1000, icon:intIcon})

				// 判断操作成功
				if (json.status == 1)
				{
                    // 修改之后的处理
                    self.beforeSave(json.data);
					self.table.draw(false);
					if (self.actionType !== "delete") $("#" + self.options.dialogId).modal('hide');
					self.actionType = "";
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