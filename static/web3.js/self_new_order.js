Account = window.sessionStorage.getItem("Global_Account");
function collect_order(){
    $.ajax({
        method: "post",
        url: "http://localhost:8080/dapp/queryCollectOrder",
        data: { Account: Account},
        success: function (data) {
            console.log(data);
            result =data.data;
            sessionStorage.setItem('collect_signal_Order',JSON.stringify(result))
            console.log(data.data)
            window.location.href="/collect_order"
        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
}