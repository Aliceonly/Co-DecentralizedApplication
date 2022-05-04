function Login(){
    var Sid=$("#Sid").val();
    var Password=$("#password").val();
    console.log(Sid,Password)
    $.ajax({
        method:"post",
        url: "http://localhost:8080/dapp/login",
        data: { Sid:Sid, Password: Password },
        success: function (data) {
            console.log("success login",dat);
        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
}