/**
 * Created by liujinxing on 2015/8/25.
 */
// 判断为空
function empty(val)
{
    return val == undefined || val == '';
}

// 判断一个数组中是否存在某个值 , 返回布尔值
function in_array(val , arr) {
    for(var i in arr) if (arr[i] === val ) return true ;
    return false ;
}

// 处理其他参数
function handleParams(params)
{
    var other = '';
    if (params != undefined && typeof params == "object") {
        for (var i in params) other += ' ' + i + '="' + params[i] + '" '
    }
    
    return other;
}

// 生成Label
function Label(content, params)
{
    return '<label ' + handleParams(params) + '> ' + content + ' </label>';
}

// 生成input标签
function createInput(type, params)
{
    return '<input type="' + type + '" ' + handleParams(params) +  ' />';
}

// 生成radio
function createRadio(data, checked, params)
{
    var html = '', params = handleParams(params);
    if (data != undefined && typeof data == "object")
    {
        for (var i in data) 
        {
            var check = checked == i ? ' checked="checked" ' : '';
            html += '<label> <input type="radio" ' + params + check + ' value="' + i + '" /> ' + data[i] + ' </label> ';
        }
    }

    return html;
}

// 生成select
function createSelect(data, selected, params)
{
    var html = '', params = handleParams(params);
    if (data != undefined && typeof data == "object")
    {

        html += '<select ' + params + '>';
        for (var i in data) 
        {
            var select = i == selected ? ' selected="selected" ' : '';
            html += '<option value="' + i + '" ' + select + ' >' + data[i] + '</option>';
        }

        html += '</select>';
    }

    return html;
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
function verifyUpload(uploadObj, size, allowType, fileurl)
{
    // 初始化定义
    var obj     = uploadObj.files[0],                               // 上传对象信息
        arr     = [false, '对不起！上传文件超过指定值...'],         // 定义错误信息
        num     = obj.name.indexOf('.'),		                    // 获取文件的后缀名
        fileext = obj.name.substr(num + 1).toLocaleLowerCase();     // 获取文件后缀名

    // 默认赋值(支持上传文件类型)
    if (allowType == undefined) allowType = ['jpeg', 'jpg', 'gif', 'png'];

    // 判断大小
    if (obj.size < size)
    {
        arr[1] = '对不起！上传文件类型错误...';
        if (in_array(fileext, allowType))
        {
            // 追加参数
            if ( fileurl != undefined ) 
            {			
                var link = uploadObj.url.indexOf('?') >= 0 ? '&' : '?' ;
                uploadObj.url += link + 'fileurl=' + fileurl ;
            }

            arr = [true, '文件上传成功！'];
        }
    }

    // 返回数据
    return arr;
}