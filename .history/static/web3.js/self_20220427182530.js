function Deletedata(e){
    Timestamp = e.parentElement.parentElement.parentElement.children[1].children[0].children[1].children[0].innerText
    var Timestamp1 = Timestamp.replace(/[^0-9]/ig, "");
    console.log(Timestamp1);
    $.ajax({
        method: "post",
        url: "http://localhost:8080/dapp/deleteData",
        data: { timestamp: Timestamp1 },
        success: function (data) {
            // alert("删除成功") 
            confirm("确定要删除订单嘛？删除了可就没了哦！")
            window.location.href ="/delete_succ"
            // console.log("result=====>", result);
           
        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
}
