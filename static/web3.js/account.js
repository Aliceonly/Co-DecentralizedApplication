let Account = window.sessionStorage.getItem("Global_Account")
document.getElementById("account").innerHTML = Account;
function resume() {
    $.ajax({
        method: "post",
        url: "http://localhost:8080/dapp/user_info",
        data: { Account: Account},
        success: function (data) {
            sessionStorage.setItem('QueryUser',JSON.stringify(data.data))
            console.log(data.data)
            window.location.href="/resume"
        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
}