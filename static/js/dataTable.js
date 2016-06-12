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
function setOperate(td, data, rowArr, row, col)
{
	$(td).html(createButtons([
		{"data":row, "title":"查看", "className":"btn-success", "cClass":"me-table-view",  "icon":"fa-search-plus",  "sClass":"blue"},
		{"data":row, "title":"编辑", "className":"btn-info", "cClass":"me-table-edit", "icon":"fa-pencil-square-o",  "sClass":"green"},
		{"data":row, "title":"删除", "className":"btn-danger", "cClass":"me-table-del", "icon":"fa-trash-o",  "sClass":"red"}
	]));
}

// 多选按钮信息
var oCheckBox = {
		"data": 	 null, 
		"bSortable": false, 
		"class": 	 "center", 
		"title": 	 '<label class="position-relative"><input type="checkbox" class="ace" /><span class="lbl"></span></label>',
		"render": 	 function(data){
			return '<label class="position-relative"><input type="checkbox" value="' + data["id"] + '" class="ace" /><span class="lbl"></span></label>';
        }
    };

// 默认操作选项
var oOperate = {"data": null, "title":"操作", "bSortable":false, "createdCell":setOperate};
var oOperateDetails = {"data":null, "title":"操作", "bSortable":false, "createdCell":function(td, data, rowArr, row, col){
	$(td).html(createButtons([
		{"data":row, "title":"查看", "className":"btn-success", "cClass":"me-table-view-detail",  "icon":"fa-search-plus",  "sClass":"blue"},
		{"data":row, "title":"编辑", "className":"btn-info", "cClass":"me-table-edit-detail", "icon":"fa-pencil-square-o",  "sClass":"green"},
		{"data":row, "title":"删除", "className":"btn-danger", "cClass":"me-table-del-detail", "icon":"fa-trash-o",  "sClass":"red"}
	]));
}};

var oTableLanguage = {
	// 显示
	"sLengthMenu": 	 "每页 _MENU_ 条记录",
	"sZeroRecords":  "没有找到记录",
	"sInfo": 		 "显示 _START_ 到 _END_ 共有 _TOTAL_ 条数据",
	"sInfoEmpty": 	 "无记录",
	"sInfoFiltered": "(从 _MAX_ 条记录过滤)",
	"sSearch": 		"搜索：",
	// 分页
	"oPaginate": {
		"sFirst": 	 "首页",
		"sPrevious": "上一页",
		"sNext": 	 "下一页",
		"sLast": 	 "尾页"
	}
};

// 服务器数据处理
function fnServerData(sSource, aoData, fnCallback) {
    var intLayer   = layer.load(),
        attributes = aoData[2].value.split(","),
        mSort 	   = (attributes.length + 1) * 5 + 2;

    // 添加查询条件
    $('.msearch').each(function(){
        var v = $(this).val();
        if ( ! empty(v) && v != 'All') aoData.push({"name":'params[' + $(this).attr('name') + ']', "value": $(this).val()});
    });

    // 添加排序字段信息
    if (aoData[mSort].value != undefined && aoData[mSort].value != "")
    {
        var tmpkey = parseInt(aoData[mSort].value);
        aoData.push({"name":'params[orderBy]', "value": attributes[tmpkey]});
    }

    // ajax请求
    $.ajax({
        url: sSource,
        data: aoData,
        type: 'post',
        dataType: 'json',
        success: function(data) {
            layer.close(intLayer);
            // 判断返回数据
            if (data.status != 1) return layer.msg('出现错误:' + data.msg, {time:2000, icon:5});

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
	function MeTable(options, tableOptions, detailOptions) {
		// 表格信息配置
		this.tableOptions = {
			"fnServerData": fnServerData,		// 获取数据的处理函数
			"sAjaxSource": "search",			// 获取数据地址
			"bLengthChange": true, 				// 是否可以调整分页
			"bAutoWidth": false,           	 	// 是否自动计算列宽
            "bPaginate": true,					// 是否使用分页
            "iDisplayStart": 0,
            "iDisplayLength": 10,
            "bServerSide": true,		 		// 是否开启从服务器端获取数据
            "bRetrieve": true,
            "bDestroy": true,
            // "processing": true,				// 是否使用加载进度条
            "sPaginationType": "full_numbers",  // 分页样式
            "oLanguage": oTableLanguage,		// 语言配置
            "order":[[1, "desc"]],
		};

		// 自定义信息配置
		this.options = {
			sModal: 	  "#myModal", 		// 编辑Modal选择器
			sTitle: 	  "",				// 表格的标题
			sTable: 	  "#showTable", 	// 显示表格选择器
			sFormId:  	  "#editForm",		// 编辑表单选择器
			sBaseUrl:     "update",			// 编辑数据提交URL
			sSearchHtml:  "",				// 搜索信息
			sSearchType:  "middle",			// 搜索表单位置
			sSearchForm:  "#searchForm",	// 搜索表单选择器
			bRenderH1: 	  true,				// 是否渲染H1内容
			iViewLoading: 0, 				// 详情加载Loading
			bViewFull: 	  false,			// 详情打开的方式 1 2 打开全屏
		};

		// 配置信息修改和继承
		this.tableOptions = $.extend(this.tableOptions, tableOptions);
		this.options 	  = $.extend(this.options, options);
		this.formOptions  = $.extend({
			"method": "post",
			"id": 	  "editForm",
			"class":  "form-horizontal",
			"name":   "editForm",
			"action": this.options.sBaseUrl,
		}, this.options.formOptions);

		// 操作类型
		this.actionType     = "";	  // 默认没有类型
		this.bHandleDetails = false;  // 默认没有开启详情处理
		this.oDetails 		= null;   // 详情配置为空

		// 详情配置的处理
		if (detailOptions != undefined && typeof detailOptions == "object")
		{
			this.bHandleDetails = true;
			this.oDetailParams  = null;
			this.oDetailObject  = null;
			var self = this;
			this.oDetails 		= {
				sTable:   "#detailTable",
				sModal:   "#myDetailModal",
				sBaseUrl: "edit",
				oTableOptions: {
					"bPaginate": 	 false,             // 不使用分页
					"bLengthChange": false,             // 是否可以调整分页
					"bServerSide": 	 true,		 		// 是否开启从服务器端获取数据
					"bAutoWidth": 	 false,
					"sAjaxSource":	"view",
					"fnServerData": function(sSource, aoData, fnCallback) {
						if (self.oDetailParams)
						{
							var intLayer = layer.load()
							for (var i in self.oDetailParams) aoData.push({name:i, value:self.oDetailParams[i]})
							// ajax请求
							$.ajax({
								url: sSource,
								data: aoData,
								type: 'post',
								dataType: 'json',
								success: function(data) {
									layer.close(intLayer);
									// 判断返回数据
									if (data.status != 1) return layer.msg('出现错误:' + data.msg, {time:2000, icon:5});
									$.fn.dataTable.defaults['bFilter'] = true;
									fnCallback(data.data);
									if (self.oDetailObject) self.oDetailObject.child(function(){return $('#detailTable').parent().html();}).show()
								},
								error: function() {
									layer.close(intLayer);
									layer.msg("服务器繁忙,请稍候再试...", {time:2000});
								}
							});
						}

					},		// 获取数据的处理函数
					"searching": 	 false,
					"ordering":  	 false,
					"oLanguage": 	 oTableLanguage,		// 语言配置
				}
			};

			detailOptions.oTableOptions = $.extend(this.oDetails.oTableOptions, detailOptions.oTableOptions)
			this.oDetails 		= $.extend(this.oDetails, detailOptions)
		}
	}

	// 处理表单信息
	MeTable.prototype.CreateForm = function() {
		var self       = this,
			formParams = handleParams(this.formOptions),
			form 	   = '<form ' + formParams + '><fieldset>',
			views 	   = '<table class="table table-bordered table-striped table-detail">';

		// 处理生成表单
		this.tableOptions.aoColumns.forEach(function(k, v) {
			views += createViewTr(k.title, k.data);
			if (k.edit != undefined) form += createForm(k);

			// 处理搜索
			if (k.search != undefined)
			{
				var tmpOptions = {"name":k.sName, "vid":v, "class":"msearch"},html = '';
                if (k.search.options) $.extend(tmpOptions, k.search.options);
                if ( self.options.sSearchPosition == 'top') tmpOptions['placeholder'] = '请输入' + k.title;
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

				self.options.sSearchHtml += Label(k.title + " : " + html) + ' ';
			}
		});

		// 生成HTML
		var Modal = '<div class="isHide" id="data-info"> ' + views +  ' </table></div> \
				    <div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myImageLabel"> \
				        <div class="modal-dialog" role="document"> \
				            <div class="modal-content"> \
				                <div class="modal-header"> \
				                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button> \
				                    <h4 class="modal-title">编辑详情</h4> \
				                </div> \
				                <div class="modal-body">' + form + '</fieldset></form></div> \
				                <div class="modal-footer"> \
				                    <button type="button" class="btn btn-default" data-dismiss="modal">取消</button> \
				                    <button type="button" class="btn btn-primary btn-image me-table-save">确定</button> \
				                </div> \
				            </div> \
				        </div> \
				    </div>';

		// 处理详情编辑信息
		if (this.bHandleDetails) {
			form  = '<form id="meDetailForm" class="form-horizontal" action="' + this.oDetails.sBaseUrl + '" name="meDetailForm" method="post" enctype="multipart/form-data"><fieldset>';
			views = '<table class="table table-bordered table-striped table-detail">';
			// 处理生成表单
			this.oDetails.oTableOptions.aoColumns.forEach(function(k) {
				views += createViewTr(k.title, k.data + '-detail'); // 查看详情信息
				if (k.edit != undefined) form += createForm(k);		// 编辑表单信息
			});

			// 添加详情输入框
			Modal += '<div class="isHide" id="data-info-detail"> ' + views +  ' </table></div> \
					<div class="modal fade" id="myDetailModal" tabindex="-1" role="dialog"> \
				        <div class="modal-dialog" role="document"> \
				            <div class="modal-content"> \
				                <div class="modal-header"> \
				                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button> \
				                    <h4 class="modal-title">编 辑</h4> \
				                </div> \
				                <div class="modal-body">' + form + '</fieldset></form></div> \
				                <div class="modal-footer"> \
				                    <button type="button" class="btn btn-default" data-dismiss="modal">取消</button> \
				                    <button type="button" class="btn btn-primary btn-image me-table-save-detail">确定</button> \
				                </div> \
				            </div> \
				        </div> \
				    </div>';
		}

		// 向页面添加HTML
    	$("body").append(Modal);
	};

	// 生成表格对象
	MeTable.prototype.init = function() {
		var self = this
		this.CreateForm();

		// 初始化主要表格
		this.table   = $(this.options.sTable).DataTable(this.tableOptions);

		// 初始化详情表格
		if (this.bHandleDetails) this.details = $(this.oDetails.sTable).DataTable(this.oDetails.oTableOptions)

		// 判断是否渲染H1
        if (this.options.bRenderH1) $('h1').html(this.options.sTitle);

        // 判断初始化处理(搜索添加位置)
        if (this.options.sSearchType == 'middle')
        {
            $('#showTable_filter').html('<form action="post" id="searchForm">' + self.options.sSearchHtml + '</form>');

            // 表格添加搜索事件
            $('.msearch').on('keyup change', function () {
                self.table.column(parseInt($(this).attr('vid'))).search($(this).val()).draw();
            });

            // 处理搜索信息
            $('#showTable_wrapper div.row div.col-xs-6:first').removeClass('col-xs-6').addClass('col-xs-2').next().removeClass('col-xs-6').addClass('col-xs-10');
        } else {
        	// 添加搜索表单信息
        	$(this.options.sSearchForm).append(self.options.sSearchHtml);
        }

		// 新增
		$('.me-table-insert').click(function(evt){evt.preventDefault();self.insert();});

		// 修改
		$(document).on('click', '.me-table-edit', function(evt){evt.preventDefault();self.update($(this).attr('table-data'))});

		// 删除
		$(document).on('click', '.me-table-del', function(evt){evt.preventDefault();self.delete($(this).attr('table-data'))});

		// 修改
		$(document).on('click', '.me-table-edit-detail', function(evt){evt.preventDefault();self.updateDetail($(this).attr('table-data'))});

		// 查看
		$(document).on('click', '.me-table-view', function(evt){evt.preventDefault();self.view($(this).attr('table-data'))});

		// 刷新
		$('.me-table-reload').click(function(evt){evt.preventDefault();self.search();}); 

		// 删除全部	
		$('.me-table-delete').click(function(evt){evt.preventDefault();self.deleteAll();}); 

		// 保存
		$('.me-table-save').click(function(evt){evt.preventDefault();self.save();});

		// 行选择
        $(document).on('click', self.options.sTable + ' th input:checkbox' , function(){
            var that = this;
            $(this).closest('table').find('tr > td:first-child input:checkbox')
            .each(function(){
                this.checked = that.checked;
                $(this).closest('tr').toggleClass('selected');
            });
        });
	};

	// 表格搜索
	MeTable.prototype.search = function() {this.table.draw();};

	// 初始化表单对象
	MeTable.prototype.initForm = function(data)
	{
		layer.closeAll();
		// 弹出标题显示
		var title = this.options.sTitle + (this.actionType == "insert" ? "新增" : "编辑");
		$(this.options.sModal).find('h4').html(title);

		// 表单处理
		objForm = $(this.options.sFormId).get(0)
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
		$(this.options.sModal).modal({backdrop: "static"});
	}

	// 查看详情
	MeTable.prototype.view = function(row) {
        if (this.options.iViewLoading != 0) return false;
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
			this.options.iViewLoading = layer.open({
			    type: 1,
			    shade: 0.3,
                shadeClose:true,
			    title: self.options.sTitle + '详情',
			    content: $('#data-info'), 		// 捕获的元素
			    area:['50%', 'auto'],
			    cancel: function(index){layer.close(index);},
                end:function(){
                    $('.views-info').html('')
                    self.options.iViewLoading = 0;
                },
				maxmin: true
			});

			// 展开全屏(解决内容过多问题)
			if (this.options.bViewFull) layer.full(this.options.iViewLoading)
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
	MeTable.prototype.updateDetail = function(row) {
		this.actionType = "updateDetail";
		$(this.oDetails.sModal).modal({backdrop: "static"});
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
			self.save(data);
			// 取消删除
		}, function(){
			layer.msg('您取消了删除操作！', {time:800});
		});
	};

	// 删除全部数据
	MeTable.prototype.deleteAll = function() {
        var data = [], self = this;
        this.actionType = "deleteAll";

        // 数据添加
        $(this.options.sTable + " tbody input:checkbox:checked").each(function(){data.push($(this).val());});

        // 数据为空
        if (data.length < 1)  return bootbox.alert({
            title:"温馨提醒",
            message:"您没有选择需要删除的数据 ! ",
        });

        // 确认操作
        bootbox.dialog({
            title:"温馨提醒",
            size:"small",
            message:'<p style="padding-left:15px; color:red">确定需要删除这' + data.length + '条数据吗?</p>',
            buttons:{
                succee:{
                    label:'<span class="ui-button-text"><i class="ace-icon fa fa-trash-o bigger-110"></i> 确定删除 </span>',
                    className:"btn btn-danger",
                    callback:function() {
                        self.save({"ids":data.join(',')});
                    }
                },
                cell:{
                    label:"取消",
                    className:"btn-default",
                    callback:function(){
                        layer.msg("您取消了删除操作！");
                    }
                }
            },
        });
	};

	// 修改数据之后的处理
    MeTable.prototype.beforeSave = function(){return true;};

	// 数据新增和修改的执行
	MeTable.prototype.save = function(data) {
		layer.closeAll();
		var self = this;
		// 判断类型
		if (this.actionType == "") return false;

		// 新增和修改验证数据
		if (this.actionType !== "delete" && this.actionType !== "deleteAll")
		{
			// 数据验证
			if ( ! $(this.options.sFormId).validate(validatorError).form()) return false;

			// 提交数据
			data = $(this.options.sFormId).serialize();
			data += "&actionType=" + this.actionType;
		}
		else
			data.actionType = this.actionType;

		var intLoad = layer.load();
		// ajax提交数据
		$.ajax({
			url:self.options.sBaseUrl,
			type:'POST',
			data:data,
			dataType:'json',
			success:function(json)
			{
				layer.close(intLoad);

				// 提示信息
				var intIcon = json.status == 1 ? 6 : 5;
				layer.msg(json.msg, {icon:intIcon})

				// 判断操作成功
				if (json.status == 1)
				{
                    // 修改之后的处理
                    self.beforeSave(json.data);
					self.table.draw(false);
					if (self.actionType !== "delete") $(self.options.sModal).modal('hide');
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