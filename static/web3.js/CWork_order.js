let Account_status = window.sessionStorage.getItem("Account_status");
console.log("Account_status===》", Account_status);

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
        showCancelButton: true, 
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

data = JSON.parse(sessionStorage.getItem('Query_CWork_order_detail'))
console.log(data);
if (data == undefined) {
    swal("OMG!", "一个订单也没有", "error");
} else {
    var result = "";
    data.forEach(e => {
        result += `<div class="col-lg-12">
        <div class="job-card-two">
        <div class="row align-items-center">
        <div class="col-md-1">
        </div>
        <div class="col-md-8">
        <div class="job-info">
        <h3>${e.Taskname}</h3>
        <ul>
        <li>
        <p>薪资：${e.Amount}</p>
        </li>
        <li>
        <p >状态：${e.State}</p>
        </li>
        <li >
        <p>时间戳：${e.Timestamp}</p>
        </li>
        <li>
        <p>时间：${e.LaunchTime}</p>
         </li>
        </ul>
        <span>${e.Category}</span>
        </div>
        </div>
        <div class="col-md-3">
         <div class="theme-btn text-end">
        <button>
        <a onclick="Read_more(this)" class="default-btn">查看更多 </a>
        </button>
        </div>
        </div>
        </div>
        </div>
        </div>
        `
    })
    document.getElementById("CWork_order_detail").innerHTML = result;
}

function Read_more(e) {
    if (Account_status == "") {
        Login_status()
    } else {
        Timestamp = (e.parentElement.parentElement.parentElement.parentElement.parentElement.children[0].children[1].children[0].children[1].children[2].innerText).slice(4,14)
        console.log(Timestamp);
        $.ajax({
            method: "post",
            url: "http://localhost:8080/dapp/CWork_order_Read_more",
            data: { timestamp: Timestamp },
            success: function (data) {
                console.log("success data", data);
                result = data.data
                sessionStorage.setItem('CWork_order_detail', JSON.stringify(data.data))
                console.log("result=====>", result);
                window.location.href="CWork_order_Read_more"
            },
            error: function (data) {
                console.log("error====>", error)
                console.log("error data===>", data)
            }
        })
    }
}



