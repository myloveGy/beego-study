/**
 * file: base.js
 * desc: 主要函数库
 * user: liujx
 * date: 2016-6-15
 * 注 ： create开头函数用来配合生成HTML 第一个参数接收配置信息, 第二个接收值, 第三个接收默认值
 */
// 验证数据是否为空
function empty(val){return val==undefined||val==""}
// 判断值是否存在数组或者对象中
function in_array(val,arr){for(var i in arr){if(arr[i]===val){return true}}return false}
// 首字母大写
function ucfirst(str){return str.substr(0, 1).toUpperCase() + str.substr(1)}
// 连接参数为字符串
function handleParams(params){var other="";if(params!=undefined&&typeof params=="object"){for(var i in params){other+=" "+i+'="'+params[i]+'" '}}return other}
// 生成label
function Label(content,params){return"<label "+handleParams(params)+"> "+content+" </label>"}
// 生成Input
function createInput(params, type){return'<input type="'+type+'" '+handleParams(params)+" />"}
// 生成text
function createText(params) {console.info(params); return createInput(params, 'text')}
// 生成textarea
function createTextarea(params){if(empty(params)){params={"class":" form-control","rows":5}}else{params["class"]+=" form-control";params["rows"]=5}return"<textarea "+handleParams(params)+"></textarea>"}
// 生成radio
function createRadio(params, data, checked){
    var html="",params=handleParams(params);
    if(data!=undefined&&typeof data=="object"){
        for(var i in data){
            var check = checked == i ? ' checked="checked" ':"";
            html += '<label class="line-height-1 blue"> <input type="radio" '+params+check+' value="'+i+'"  /> <span class="lbl"> '+data[i]+" </span> </label>　 "
        }
    }
    return html
}

// 生成select
function createSelect(params, data, selected){
    var html = "", params = handleParams(params);
    if(data != undefined && typeof data == "object") {
        html += "<select "+params+">";
        for(var i in data){
            var select = i == selected ? ' selected="selected" ':"";
            html += '<option value="'+i+'" '+select+" >"+data[i]+"</option>"
        }

        html += "</select>"}
    return html
}

// 生成上传文件类型 file
function createFile(params){
    if(params == undefined) params = {}
    var html = createInput(params, "hidden");
    params["name"]  = "UploadForm[" + params["name"] + "]";
    params["class"] = "input-file uniform_on fileUpload";
    html += createInput(params, "file");
    html += "<p class='bg-success p-5 m-3 isHide' onclick='$(this).hide()'></p>";
    return html
}

// 添加时间天
function createDate(params) {
    return '<div class="input-group bootstrap-timepicker"> \
        <input class="form-control date-picker" id="id-date-picker-1" type="text" data-date-format="yyyy-mm-dd" /> \
        <span class="input-group-addon"><i class="fa fa-calendar bigger-110"></i></span> \
        </div>';
}

// 添加时间分秒
function createTime(params) {
    return '<div class="input-group bootstrap-timepicker"> \
        <input id="timepicker1" type="text" class="form-control" /> \
        <span class="input-group-addon"><i class="fa fa-clock-o bigger-110"></i></span> \
        </div>';
}

// 添加时间
function createDatetime(params) {
    return '<div class="input-group"> \
        <input id="date-timepicker1" type="text" class="form-control" /> \
        <span class="input-group-addon"><i class="fa fa-clock-o bigger-110"></i></span> \
        </div>';
}

// 添加时间段
function createDaterange(params) {
    return '<div class="input-group"> \
        <span class="input-group-addon"><i class="fa fa-calendar bigger-110"></i></span> \
        <input class="form-control" type="text" name="date-range-picker" id="id-date-range-picker-1" /> \
        </div>';
}

// 生成多语言配置信息
function createLang(obj)
{
    var html = '<div class="col-sm-12 m-pl-0"><div class="tabbable">', mli = '', mDiv = '';
    if (language)
    {
        obj['class'] = 'form-control';
        var n = obj.name, m = 1;
        for (var i in language)
        {
            obj.name = n;
            obj.name += '[' + i + ']';
            var params = handleParams(obj), mid = n + m, input = obj.type == undefined ? "<textarea " + params + "></textarea>" : "<input type=\"text\" "+ params +"/>";
            mli += '<li class="' + (m == 1 ? 'active' : '') + '">\
                    <a href="#' + mid + '" data-toggle="tab">' + language[i] + '</a>\
                </li>';
            mDiv += '<div class="tab-pane ' + (m == 1 ? 'active' : '') + '" id="'+ mid + '">\
                    ' + input + '\
                </div>';
            m ++;
        }
    }

    html += '<ul class="nav nav-tabs padding-12 tab-color-blue background-blue">\
            ' + mli + '\
            </ul>\
            <div class="tab-content">\
            ' + mDiv + '\
            </div>\
        </div></div>';

    return html;
}

// 多选按钮 checkbox
function createCheckbox(params, data, checked, arr, isHave)
{
    if (arr == undefined) arr = 'col-xs-6';
    var html = '', params = handleParams(params);
    if (data != undefined && typeof data=="object")
    {
        if (isHave == undefined) html += '<div class="checkbox col-xs-12"><label><input type="checkbox" class="ace checkbox-all" /><span class="lbl"> 选择全部 </span></label></div>';
        for (var i in data)
        {
            html += '<div class="checkbox ' + arr + '"><label><input type="checkbox" ' + params + ' value="' + i + '" /><span class="lbl"> ' + data[i] + ' </span></label></div>';
        }
    }

    return html;
}

// 生成按钮
function createButtons(data) {
    var div1   = '<div class="hidden-sm hidden-xs btn-group">',
        div2   = '<div class="hidden-md hidden-lg"><div class="inline position-relative"><button data-position="auto" data-toggle="dropdown" class="btn btn-minier btn-primary dropdown-toggle"><i class="ace-icon fa fa-cog icon-only bigger-110"></i></button><ul class="dropdown-menu dropdown-only-icon dropdown-yellow dropdown-menu-right dropdown-caret dropdown-close">';
    // 添加按钮信息
    if(data != undefined && typeof data == "object")
    {
        for(var i in data)
        {
            div1 += ' <button class="btn ' + data[i]['className'] + ' '+  data[i]['cClass'] + ' btn-xs" table-data="' + data[i]['data'] + '"><i class="ace-icon fa ' + data[i]["icon"] + ' bigger-120"></i></button> ';
            div2 += '<li><a title="' + data[i]['title'] + '" data-rel="tooltip" class="tooltip-info ' + data[i]['cClass'] + '" href="javascript:;" data-original-title="' + data[i]['title'] + '" table-data="' + data[i]['data'] + '"><span class="' + data[i]['sClass'] + '"><i class="ace-icon fa ' + data[i]['icon'] + ' bigger-120"></i></span></a></li>'; 
        }
    }

    return div1 + '</div>' + div2 + '</ul></div></div>';
}

// 生成表单对象
function createForm(k)
{
    var form = '';
    // 处理其他参数
    if (k.edit.options == undefined) k.edit.options = {}; // 容错处理
    if (!k.edit.type) k.edit.type = "text";
    k.edit.options["name"]  = k.sName;
    k.edit.options["class"] = "form-control";
    if (k.edit.type == undefined) k.edit.type = "text"

    if ( k.edit.type == "hidden" )
        form += createInput(k.edit.options, 'hidden');
    else
    {
        // 判断类型
        form += '<div class="form-group">' + Label(k.title, {"class":"col-sm-3 control-label"}) + '<div class="col-sm-9">';

        // 单选选按钮添加样式
        if (k.edit.type == "radio") k.edit.options['class'] = 'ace valid';
        // 多选按钮处理
        if (k.edit.type == "checkbox") {
            k.edit.options['class'] = 'ace m-checkbox';
            k.edit.options['name']  = k.sName + '[]';
        }

        // 时间处理
        if (k.edit.type == "time") {
            if (!empty(k.value)) {k.edit.options["value"] = k.value}
            k.edit.options["class"] += " time";
        }

        // 默认输入框处理
        if (k.edit.type == "text") if (!empty(k.value)) k.edit.options["value"] = k.value

        // 使用函数
        var func = 'create' + k.edit.type.substr(0, 1).toUpperCase() + k.edit.type.substr(1);
        form += window[func](k.edit.options, k.value, k.edit.default) + '</div></div>';
    }

    return form;
}

// 生成查看详情的Table
function createViewTr(title, data) {
    return '<tr><td width="25%">' + title + '</td><td class="views-info data-info-' + data + '"></td></tr>'
}

// 生成查表单信息
function createSearchForm(k, v) {
    k.search.options = $.extend({"name":k.sName, "vid":v, "class":"msearch"}, k.search.options);
    if (k.search.type == "select") {k.value["All"] = "全部";}
    var html = window['create' + ucfirst(k.search.type)](k.search.options, k.value, "All");
    if (k.search.type == "select") delete k.value["All"]
    return Label(k.title + " : " + html) + ' ';
}

// 生成编辑和查看详细modal
function createModal(oModal, oViews) {
    return '<div class="isHide" '+ handleParams(oViews['params']) +'> ' + oViews['html'] +  ' </table></div> \
            <div class="modal fade" '+ handleParams(oModal['params']) +' tabindex="-1" role="dialog" > \
                <div class="modal-dialog" role="document"> \
                    <div class="modal-content"> \
                        <div class="modal-header"> \
                            <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button> \
                            <h4 class="modal-title"> 编 辑 </h4> \
                        </div> \
                        <div class="modal-body">' + oModal['html'] + '</fieldset></form></div> \
                        <div class="modal-footer"> \
                            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button> \
                            <button type="button" class="btn btn-primary btn-image ' + oModal['bClass'] + '">确定</button> \
                        </div> \
                    </div> \
                </div> \
            </div>';
}


// 验证上传文件
function verifyUpload(uploadObj,size,allowType,fileurl){var obj=uploadObj.files[0],arr=[false,"对不起！上传文件超过指定值..."],num=obj.name.indexOf("."),fileext=obj.name.substr(num+1).toLocaleLowerCase();if(allowType==undefined){allowType=["jpeg","jpg","gif","png"]}if(obj.size<size){arr[1]="对不起！上传文件类型错误...";if(in_array(fileext,allowType)){if(fileurl!=undefined){var link=uploadObj.url.indexOf("?")>=0?"&":"?";uploadObj.url+=link+"fileurl="+fileurl}arr=[true,"文件上传成功！"]}}return arr}

// 文件上传
function FileUpload(url, select, allowType, size){
    $(select).fileupload({
        dataType:"json",
        url:url,
        beforeSend:function(e,data){
            var arr=verifyUpload(data, size, allowType, $(select).parent().find("input[type=hidden]").val());
            if(!arr[0])
            {
                layer.msg(arr[1],{time:1500,icon:5});
                return false
            }
        },
        success:function(json){
            if(json.status==1){
                layer.msg(json.msg, {time:1500,icon:6,end:function(){
                    var str = json.data.image + " " + json.msg;
                    $(select).next('p').html(str).show().parent().find("input[type=hidden]").val(json.data.fileDir)
                }});
                return false
            }

            layer.msg(json.msg,{time:1500,icon:5})
        },
        error:function(){
            layer.msg('服务器繁忙,请稍候再试...', {time:1500});
        }
    })
}

// 时间格式化
Date.prototype.Format=function(fmt){var o={"M+":this.getMonth()+1,"d+":this.getDate(),"h+":this.getHours(),"m+":this.getMinutes(),"s+":this.getSeconds(),"q+":Math.floor((this.getMonth()+3)/3),"S":this.getMilliseconds()};if(/(y+)/.test(fmt)){fmt=fmt.replace(RegExp.$1,(this.getFullYear()+"").substr(4-RegExp.$1.length))}for(var k in o){if(new RegExp("("+k+")").test(fmt)){fmt=fmt.replace(RegExp.$1,(RegExp.$1.length==1)?(o[k]):(("00"+o[k]).substr((""+o[k]).length)))}}return fmt};
// 根据时间戳返回时间字符串
function timeFormat(time,str){if(empty(str)){str="yyyy-MM-dd"}var date=new Date(time*1000);return date.Format(str)}
// 值的转换
function stringTo(type,value){switch(type){case"int":case"int8":case"int16":case"int32":case"int64":case"uint":case"uint8":case"uint16":case"uint32":case"uint64":return parseInt(value);case"bool":return value==="true"||value===true||value===1||value=="1";case"float32":case"float64":}return value};
// 初始化表单信息
function InitForm(select, data) {
    objForm = $(select).get(0); // 获取表单对象
    if (objForm != undefined)
    {
        $(objForm).find('input[type=hidden]').val('');                                  // 隐藏按钮充值
        $(objForm).find('input[type=checkbox]').each(function(){$(this).attr('checked', false);if ($(this).get(0)) $(this).get(0).checked = false;});                                                                             // 多选菜单
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
}

// 详情表单赋值
function viewTable(object, data, tClass, row)
{
    // 循环处理显示信息
    object.forEach(function(k) {
        var tmpKey = k.data,tmpValue = data[tmpKey],dataInfo = $(tClass + tmpKey);
        if (k.edit != undefined && k.edit.type == 'password') tmpValue = "******";
        (k.createdCell != undefined && typeof k.createdCell == "function") ? k.createdCell(dataInfo, tmpValue, data, row, undefined) : dataInfo.html(tmpValue);
    });
}