function create_user(){
    var Sid=$("#Sid").val();
    var Telephone=$("#Telephone").val();
    var Password=$("#password").val();
    console.log(Sid,Telephone,Password)
    $.ajax({
        method:"post",
        url: "http://localhost:8080/dapp/creatUser",
        data: { Sid:Sid, Telephone:Telephone, Password: Password },
        success: function (data) {
            // swal("注册成功!", "您的账户是: " + data.data, "success");
            swal({
                title: "注册成功",
                text: '您的账户是：<span style="color:red">'+ data.data +'</span><br/><span style="color:red">请务必记住，一旦丢失无法找回<span><br/><a style="color:#3b3bf4" href="/sign_in">点击进行登录</a><br/>10秒后自动关闭。',
                imageUrl: "../static/image/check.png",
                html: true,
                timer: 10000,
                showConfirmButton: false
            });
            console.log("success insert user",data);
        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
}