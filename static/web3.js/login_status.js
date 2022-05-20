let Account_status = window.sessionStorage.getItem("Account_status");
console.log("Account_status===》",Account_status);

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

 
 
if(Account_status ==""){
    Login_status()
}