/**
 * Created by liujinxing on 2015/8/25.
 */
// 成功执行函数
function success(msg, fun)
{
    art.dialog({
        width:300,
        height:100,
        title:'温馨提醒',
        content:msg,
        icon: 'succeed',
        lock:true,
        okVal:'确定',
        ok:function(){},
        close:fun,
    });
}

// 询问弹框
function quest(title, content, val, success)
{
    return art.dialog({
        width:250,
        height:100,
        title:title,
        content:'<span style="font-weight:bold">'+content+'</span>',
        icon:'question',
        okVal:val,
        ok:success,
        lock:true,
        cancelVal: '再想想',
        cancel: function(){}
    });
}

// 失败执行函数
function warning(msg, time)
{
    if (msg == undefined) msg  = '服务器繁忙,请稍候再试...';
    if (time== undefined) time = 1;
    art.dialog({
        width:250,
        height:100,
        title:'温馨提醒',
        content:'<span style="font-weight:bold">'+msg+'</span>',
        icon:'warning',
        okVal:'确定',
        time:1,
        ok:function(){},
    });
}

/**
 * showDiv() 显示窗口
 */
function showDiv(obj, title, func, true_str, false_str, width, func_false)
{
    if ( func_false == undefined ) func_false = function(){$(this).dialog('close');};
    if ( false_str == undefined ) false_str = '返回';
    if ( width == undefined ) width = 'auto';

    // 默认赋值
    var arrBut = new Array({
        text: false_str,
        'class':'btn btn-xs',
        click:func_false
    });

    // 添加确定按钮
    if ( true_str == undefined ) true_str = '确定';
    if ( true_str != '')
    {
        arrBut.push({
            text: true_str,
            "class" : "btn btn-primary btn-xs",
            click: func
        });
    }

    var dialog = obj.dialog({
        modal: true,
        title: title,
        title_html: true,
        width: 'width',
        buttons:arrBut
    });

    return dialog;
}

/**
 * showDataTable()  显示表单数据信息
 * @param obj       作用的表格对象
 * @param aoColnmns 表头信息设置
 * @param length    分页长度
 * @param url       分页提交的页面
 * @returns object  返回一个对象
 */
function showDataTable(obj, aoColnmns, length, url , params)
{
    var objData = {
        "bAutoWidth": true,
        "bPaginate": true,
        "bLengthChange": true, // 是否可以调整分页
        "aoColumns": aoColnmns,
        "iDisplayStart": 0,
        "iDisplayLength": length,
        "bServerSide": true,
        "sAjaxSource": url,
        "bRetrieve":true,
        "bDestroy":true,
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
        },
    };

    if (params != undefined && params != '')
    {
        for( var i in params )
        {
            objData[i] = params[i];
        }
    }

    return obj.DataTable(objData);
}

// 判断为空
function empty(val)
{
    return ( val == undefined || val =='' || val == 0);
}

// 表格数据的显示处理
function handleTh(arrData)
{
    var html = '';
    for (var i in arrData)
    {
        html += '<td>'+arrData[i]+'</td>';
    }

    return html;
}

/**
 * initEdit() 编辑文本的初始化处理
 * @param string strContent 文本默认值
 */
function initEdit(strContent)
{
    if (strContent == undefined) strContent = '';
    // 表单初始化
    if ($('#addForm')[0]) $('#addForm')[0].reset();

    // 文本编辑器初始化
    $('.cleditorToolbar').css({
        'background':'rgba(0, 0, 0, 0) url("/Public/Admin/img/toolbar.gif") repeat scroll 0 0',
        'height':'53px',
    });

    // 初始化赋值
    $('.cleditor').val(strContent);
    $('.cleditor').cleditor()[0].updateFrame();
}

/**
 * initForm() 表单初始化赋值
 * @param objForm 需要赋值的表单对象
 * @param arrData 赋值的数据
 */
function initForm(objForm, arrData, isRest)
{
    if (isRest == undefined || isRest == true) objForm.reset();
    for (var i in arrData)
    {
        if (objForm[i] != undefined)
        {
            // 单选按钮
            if (objForm[i].type == undefined)
            {
                $(objForm).find('input[name='+i+']').removeAttr('checked').parent().removeClass('checked');
                $(objForm).find('input[name='+i+'][value=' + arrData[i] + ']').attr('checked', true).parent().addClass('checked');
            }
            else
                objForm[i].value = arrData[i];
        }
    }
}

/**
 * fluidChange() 流动布局的转换(选择器)
 * @param select 那个布局需要转换
 * @param type   是隐藏还是显示(默认显示)
 */
function fluidChange(select, type)
{
    $('.box' + select).parent().show();
    if (type == undefined || type == 'show')      // 默认第一个隐藏 选择的显示
    {
        $('.box:first').slideUp(1000);
        $('.box' + select).slideDown(1000);
    }
    else                                          // 显示
    {
        $('.box:first').slideDown(1000);
        $('.box' + select).slideUp(1000);
    }
}

/**
 * handleEditData() 数据添加的处理
 * @param   obj        添加数据的表单对象
 * @param   strUrl     请求的URL
 * @param   funName    成功执行的函数
 * @param   isVali     是否验证
 * @returns {boolean}  返回false
 */
function handleEditData(obj, strUrl, funName, isVali, isChange, isEdit)
{
    isVali = false;    // 全部验证
    if (isChange == undefined)  isChange = false;    // 是否装换内容
    if (isEdit   == undefined)    isEdit = false;    // 是否是修改数据
    if (isVali || obj.validate().form())
    {
        var strCls = '.' + obj.attr('id') + '_error';
        var loging = art.dialog({
            title:'<span style="color:green">数据正在提交,请耐心等待...</span>',
            width: 'auto',
            height: 'auto',
            lock:true,
        }).show();

        // 数据提交
        $.ajax({
            url:strUrl,
            type:'post',
            dataType:'json',
            data:obj.serialize(),
            success:function(json)
            {
                loging.close();
                if (json.status == 1)
                {
                    success(json.msg, function(){
                        $(strCls).html('');             // 错误清空
                        if (obj[0]) obj[0].reset();     // 表单重置

                        // 添加数据
                        if (isEdit)
                            isEdit.parents('tr').html(handleTh(json.data));
                        else
                            $('#showTable tbody').prepend('<tr>' + handleTh(json.data) + '</tr>');

                        // 是否装换内容
                        if (isChange)
                            fluidChange(':last', 'hide');
                        else
                            obj.parent().dialog('close');

                        funName();                      // 执行其他操作
                    });

                    return false;
                }

                $(strCls).html(json.msg).show();
                return false;
            },
            error:function()
            {
                loging.close();
                $(strCls).html('服务器繁忙,请稍候再试...').show();
            }
        })
    }

    return false;
}

// 初始化话表格
function initTable(arrData)
{
    $('.table-detail details').html('');
    for (var i in arrData)
    {
        $('.detailTable_' + i).html(arrData[i]);
    }
}

/**
 * 判断一个数组中是否存在某个值 , 返回布尔值
 * @method 		in_array()
 * @for 		所属类名
 * @param		mixed	val		需要查找的值
 * @param		mixed	arr		查找的数组
 * @return		boolean 返回一个布尔值( 存在true )
 */
function in_array( val , arr ) {
    for( var i in arr ) {
        if ( arr[i] === val ) return true ;
    }
    return false ;
}

/**
 * JsFileUpload Ajax 上传文件使用验证函数( 验证上传文件大小和类型 )
 * 在使用JsFileUpload 上传时才能使用 需要接收一个上传对象
 * @method 	verifyUpload()
 * @for 	所属类名
 * @param	object	 uploadObj	fileupload上传对象
 * @param	int		 size		允许上传文件的大小
 * @param	array  	 allowType 	允许上传的文件类型
 * @param 	string 	 fileurl    多次上传时 , 上一次上传文件的地址
 * @return 	array   返回一个保存上传信息的数组
 *			0 => success OR error
 *			1 => 错误信息
 */
function verifyUpload( uploadObj , size , allowType , fileurl )
{
    // 初始化定义
    var Obj     = uploadObj.files[0],                               // 上传对象信息
        arr     = [false, '对不起！上传文件超过指定值...'],             // 定义错误信息
        num     = Obj.name.indexOf('.') ; 		                    // 获取文件的后缀名
        fileext = Obj.name.substr( num + 1 ).toLocaleLowerCase();   // 获取文件后缀名

    // 默认赋值(支持上传文件类型)
    if (allowType == undefined) allowType = ['jpeg', 'jpg', 'gif', 'png'];

    // 判断大小
    if (Obj.size < size)
    {
        arr[1] = '对不起！上传文件类型错误...';
        if (in_array(fileext, allowType))
        {
            if ( fileurl != undefined ) {			// 追加参数
                var link = uploadObj.url.indexOf('?') >= 0 ? '&' : '?' ;
                uploadObj.url += link + 'fileurl=' + fileurl ;
            }

            arr = [true, '文件上传成功！'];
        }
    }

    // 返回数据
    return arr;
}