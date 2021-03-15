var isHave = false, Obj = {errorPlacement: validate}, lLayer, l, loginUrl = "/guest/login";

// 请求错误回应
function requestError() {
    layer.close(l);
    layer.msg('服务器繁忙,请稍候再试...');
}

// 验证函数
function validate(error, errorPlacement) {
    if (isHave !== false) {
        return false;
    }

    isHave = layer.tips($(error).html(), errorPlacement, {
        tips: [3], time: 1000, end: function () {
            isHave = false;
        }
    });
}

// 验证登录
function isLogin(isLogin) {
    try {
        let data = JSON.parse(window.localStorage.getItem("user-info"));
        if (!data || !data.username) {
            !isLogin && sLogin();
            return false
        }

        return true
    } catch (e) {
        !isLogin && sLogin();
        return false
    }
}

// 执行登录
function userLogin(obj) {
    if ($(obj).validate(Obj).form()) {
        l = layer.load();
        $.ajax({
            url: loginUrl,
            data: $(obj).serialize(),
            type: 'post',
            dataType: 'json',
            error: requestError,
            success: function (json) {
                layer.close(l);
                var i = json.code === 10000 ? 6 : 5;
                if (json.code === 10000) layer.close(lLayer);
                layer.msg(json.msg, {
                    icon: i, time: 1000, end: function () {
                        if (json.code === 10000) changUser(json.data);
                    }
                })
            },
        })
    }

    return false;
}

// 用户登录和退出切换
function changUser(params) {
    var sx = params !== null ? '.no-login' : '.is-login',
        sh = params !== null ? '.is-login' : '.no-login';
    window.localStorage.setItem("user-info", JSON.stringify(params))
    $('.user').html(params && params.username ? params.username : '');
    $(sx).fadeOut(1000, function () {
        $(sh).fadeIn(1000)
    });
}

// 显示登录页面
function sLogin() {
    var w = $(window).width() > 355 ? '355px' : '100%', s = Math.floor(Math.random() * 7);
    lLayer = layer.open({
        title: '用户登录',
        type: 1,
        shift: s,
        shadeClose: true,
        area: [w, 'auto'],
        content: $('#login').html()
    });
}

// 添加文章数据
function article(params) {
    var html = '', len = params.length;
    for (var i = 0; i < len; i++) {
        var data = params[i];
        if (empty(data.img)) data.img = $('div.blogs figure img').attr('src');
        html += '<div class="blogs"><figure><img src="' + data.img + '" alt=""></figure><ul><h3><a href="/article/view/' + data.id + '">' + data.title + '</a></h3><p>引导语：' + data.content + '</p>';
        html += '<p class="autor"><span class="lm f_l"><a href="/article/view/' + data.id + '">个人博客</a></span><span class="dtime f_l">' + timeFormat(data.create_time) + '</span><span class="viewnum f_r">浏览（<a href="/article/view/' + data.id + '">0</a>）</span><span class="pingl f_r">评论（<a href="/article/view/' + data.id + '">0</a>）</span></p>';
        html += '</ul></div>';
    }
    $('.topnews h2').after(html);
}

$(function () {
    if (isLogin) {
        // 需要发送请求是否已经登录
        $.ajax({
            url: "/guest/detail",
            dataType: "json"
        }).done(function (data) {
            if (data.code === 10000 && !data.data) {
                window.localStorage.removeItem("user-info")
            }
        })
    }

    // 用户退出
    $('.logout').click(function () {
        if (isLogin()) {
            layer.confirm('您确定要退出登录吗?', {title: '温馨提醒', btn: ['确定退出', '取消'], icon: 3, shift: 4}, function () {
                l = layer.load();
                $.get('/guest/logout', function (json) {
                    layer.msg(json.msg, {
                        time: 1000,
                        end: function () {
                            if (json.code === 10000) {
                                changUser(null)
                            }
                        }
                    });
                }, 'json').always(function () {
                    layer.close(l);
                });
            }, function () {
                layer.msg('您取消了退出登录！', {time: 1000});
            });
        }
    });

    // 登录可以操作
    $('.is-user').click(function (e) {
        e.preventDefault();
        return isLogin();
    });

    // 回顶部自动判断
    $(window).scroll(function () {
        if ($(window).scrollTop() > 250) {
            $('#tbox').fadeIn();
        } else {
            $('#tbox').fadeOut();
        }
    });

    // 回到顶部
    $('#gotop').click(function () {
        $('body,html').animate({scrollTop: 0}, 1000);
    });

    // 弹出model
    $('.publish-article').click(function () {
        $('#myModal').modal();
    });
    // 关闭modal
    $('#myModal').on('hide.bs.modal', function (e) {
        document.article.reset();
    });
    // 发布文章
    $('.btn-article').click(function () {
        if (isLogin()) {
            if ($('.article').validate(Obj).form()) {
                l = layer.load();
                $.ajax({
                    url: '/user/article/create',
                    data: $('.article').serialize(),
                    type: 'post',
                    dataType: 'json',
                    success: function (json) {
                        layer.close(l);
                        var s = json.code === 10000 ? 6 : 5;
                        if (json.code === 10000) {
                            $('#myModal').modal('hide');
                        }
                        layer.msg(json.msg, {
                            time: 2000, icon: s, end: function () {
                                article([json.data]);
                            }
                        })
                    },
                    error: requestError,
                })
            }
        }
        return false;
    })

    // 弹出model
    $('.file-upload').click(function () {
        $('#myImage').modal();
    });

    // 关闭modal
    $('#myImage').on('hide.bs.modal', function (e) {
        document.image.reset();
    });

    // 上传图片
    $('.btn-image').click(function () {
        if (isLogin()) {
            if ($('.image').validate(Obj).form()) {
                l = layer.load();
                $.ajax({
                    url: '/user/image/create',
                    data: $('.image').serialize(),
                    type: 'post',
                    dataType: 'json',
                    success: function (json) {
                        layer.close(l);
                        var s = json.code === 10000 ? 6 : 5;
                        if (json.code === 10000) {
                            $('#myImage').modal('hide');
                        }
                        layer.msg(json.msg, {time: 2000, icon: s})
                    },
                    error: requestError,
                })
            }
        }
        return false;
    })
    // 图片显示
    layer.photos({photos: "#layer-photos-demo"});
})

function FileUpload(url, selector, type, size) {
    $(selector).fileupload({
        url: url,
        dataType: 'json',
    }).on('fileuploaddone', function (e, data) {
        if (data.result.code === 10000) {
            layer.msg("图片上传成功")
            $(selector).parent().find("input[type=hidden]").val(data.result.data.path)
        } else {
            layer.msg("图片上传失败:" + data.result.msg)
        }
    });
}

FileUpload("/user/image/upload", '.fileUpload', undefined, 200000000);