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
            console.log("success insert user",data);
        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
}