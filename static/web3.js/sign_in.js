function Login(){
    var Sid=$("#Sid").val();
    var Password=$("#password").val();
    if(Sid == ""){
        alert("请输入账号")
    }
    if(Password == ""){
        alert("请输入密码")
    }
    // console.log(Sid,Password)
    $.ajax({
        method:"post",
        url: "http://localhost:8080/dapp/login",
        data: { Sid:Sid, Password: Password },
        success: function (data) {
            console.log(data)
            if (data.data==1){
                console.log("success login");
            }else{
                console.log("faile login")
            }
        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
}