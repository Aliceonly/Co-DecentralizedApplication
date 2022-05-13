var sign_in_up = document.querySelector("#sign_in_up");
$.ajax({
    method: "get",
    url: "/succ_login",
    success: function (data) {
        if(data.data == ""){
            console.log("未登录",data.data)
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
