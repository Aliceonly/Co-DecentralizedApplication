var sign_in_up = document.querySelector("#sign_in_up");
var self_info = document.querySelector("#self_info");
$.ajax({
    method: "get",
    url: "/succ_login",
    success: function (data) {
        // let Account_status = window.sessionStorage.setItem("Global_Account",data.data);
        if(data.data == ""){
            console.log("未登录",data.data)
            self_info.style.display = "none";
            window.sessionStorage.setItem("Account_status",data.data); //设置登录状态,私钥内容
        }else{
            console.log("已登录",data.data)
            sign_in_up.style.display = "none";
        }
    },
    error: function (data) {
        console.log("error====>", error)
        console.log("error data===>", data)
    }
})
