function demo_1() {
    swal("这是一个信息提示框!");
};
function demo_2() {
    swal("Good!", "登陆成功", "success");
};
function demo_3() {
    swal("OMG!", "登录失败，密码错误请重试！", "error");
};
function demo_4() {
    swal({
        title: "您确定要删除吗？",
        text: "您确定要删除这条数据？",
        type: "warning",
        showCancelButton: true,
        closeOnConfirm: false,
        confirmButtonText: "是的，我要删除",
        confirmButtonColor: "#ec6c62"
    }, function () {
        $.ajax({
            url: "do.php",
            type: "DELETE"
        }).done(function (data) {
            swal("操作成功!", "已成功删除数据！", "success");
        }).error(function (data) {
            swal("OMG", "删除操作失败了!", "error");
        });
    });
};
function demo_5() {
    swal({
        title: "发布成功",
        text: '<span style="color:red">您已发布成功，</span><a style="color:#3b3bf4" href="inde.html">点击返回首页查看</a>。<br/>5秒后自动关闭。',
        imageUrl: "check.png",
        html: true,
        timer: 5000,
        showConfirmButton: false
    });
};
function demo_6() {
    swal({
        title: "输入框来了",
        text: "这里可以输入并确认:",
        type: "input",
        showCancelButton: true,
        closeOnConfirm: false,
        animation: "slide-from-top",
        inputPlaceholder: "填点东西到这里面吧"
    }, function (inputValue) {
        if (inputValue === false) return false;
        if (inputValue === "") {
            swal.showInputError("请输入!");
            return false
        }
        swal("棒极了!", "您填写的是: " + inputValue, "success");
    });
};

function info_login() {
    var
        closeInSeconds = 5,
        displayText = ' #1 秒后将自动跳转登录页面',
        timer;

    swal({
        title: "请进行登录!",
        text: displayText.replace(/#1/, closeInSeconds),
        imageUrl: "../static/image/info1.png",
        timer: closeInSeconds * 1000,
        showCancelButton: true,
    }, function () {
            window.location.href = "/account"
        }
    );

    timer = setInterval(function () {
        closeInSeconds--;
        if (closeInSeconds < 0) {
            clearInterval(timer);
        }

        $('.sweet-alert > p').text(displayText.replace(/#1/, closeInSeconds));

    }, 1000);
}


// swal({
    // title: "您确定要删除吗？",
    // text: "您确定要删除这条数据？",
    // type: "warning",
    // showCancelButton: true,
    // closeOnConfirm: false,
    // confirmButtonText: "是的，我要删除",
    // confirmButtonColor: "#ec6c62"
// }, function () {
    // $.ajax({
        // url: "do.php",
        // type: "DELETE"
    // }).done(function (data) {
        // swal("操作成功!", "已成功删除数据！", "success");
    // }).error(function (data) {
        // swal("OMG", "删除操作失败了!", "error");
    // });
// });
