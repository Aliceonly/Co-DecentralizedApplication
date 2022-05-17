let Account = window.sessionStorage.getItem("Global_Account")
let Account_status = window.sessionStorage.getItem("Account_status");
console.log("Account_status===》", Account_status);
console.log("Account", Account);
if (Account_status != "") {
    document.getElementById("Self_account").innerHTML = Account;
    function release_order() {
        $.ajax({
            method: "post",
            url: "http://localhost:8080/dapp/Self_Order_show",
            data: { Account: Account },
            success: function (data) {
                sessionStorage.setItem('Query_release_order', JSON.stringify(data.data))
                console.log(data.data)
                window.location.href = "/release_order"
            },
            error: function (data) {
                console.log("error====>", error)
                console.log("error data===>", data)
            }
        })
    }
}

function Login_status() {
    var
        closeInSeconds = 5,
        displayText = ' #1 秒后将自动跳转登录页面',
        timer;
    swal({
        title: "请先进行登录操作!",
        text: displayText.replace(/#1/, closeInSeconds),
        imageUrl: "../static/image/info1.png",
        timer: closeInSeconds * 1000,
        showconfirmButton: true,
    }, function () {
        window.location.href = "/sign_in"
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


if (Account_status == "") {
    Login_status()
}
