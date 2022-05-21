let data =  JSON.parse(sessionStorage.getItem('Shared_order_detail'));
console.log(data);

document.getElementById("LaunchTime").innerHTML = data.LaunchTime;
document.getElementById("Taskname").innerHTML = data.Taskname;
document.getElementById("Amount").innerHTML =  data.Amount;
document.getElementById("Category").innerHTML = data.Category;
document.getElementById("State").innerHTML = data.State;
document.getElementById("Add").innerHTML =  data.Add 
document.getElementById("Block").innerHTML = data.Block;
document.getElementById("Timestamp").innerHTML = data.Timestamp;

console.log(parseInt(data.Timestamp))
var a=parseInt(data.Timestamp)


//失败接单
state = data.State;
console.log("状态", state)

//情况一
function fail_confirm() {
    var
        closeInSeconds = 5,
        displayText = ' #1 秒后将自动跳转首页',
        timer;

    swal({
        title: "抱歉，此订单已被接收！",
        text: displayText.replace(/#1/, closeInSeconds),
        imageUrl: "../static/image/info1.png",
        timer: closeInSeconds * 1000,
        showconfirmButton: true,
    }, function () {
        window.location.href = "/"
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

//情况二
function demo7() {
    var
        closeInSeconds = 5,
        displayText = ' #1 秒后将自动跳转首页',
        timer;
    swal({
        title: "抱歉，自己无法接单!",
        text: displayText.replace(/#1/, closeInSeconds),
        imageUrl: "../static/image/info1.png",
        timer: closeInSeconds * 1000,
        showconfirmButton: true,
    }, function () {
        window.location.href = "/"
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


//成功接单
function succ_confirm() {
    var
        closeInSeconds = 5,
        displayText = ' #1 秒后将自动跳转确认页面',
        timer;

    swal({
        title: "恭喜您，接单成功!",
        text: displayText.replace(/#1/, closeInSeconds),
        imageUrl: "../static/image/check.png",
        timer: closeInSeconds * 1000,
        showconfirmButton: true,
    }, function () {
        window.location.href = "/"
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

let Account = window.sessionStorage.getItem("Global_Account");
function Confirmtask() {
    if (data.Add == Account) {
        console.log("Add===>", Add)
        console.log("Account=====>", Account)
        demo7()
    } else {
        if (state == "已接单") {
            fail_confirm()
        } else {
            $.ajax({
                method: "post",
                url: "http://localhost:8080/dapp/confirmtask",
                data: { account: Account, timestap: a },
                success: function (data) {
                    console.log(data.data)
                    if (data.data == "confirmtrue") {
                        console.log("接单成功")
                        succ_confirm()
                    }
                },
                error: function (data) {
                    console.log("error====>", error)
                    console.log("error data===>", data)
                }
            })
        }

    }
}

