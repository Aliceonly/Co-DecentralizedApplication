function sorry_info() {
    var
        closeInSeconds = 5,
        displayText = ' #1 秒后将自动跳转发布订单页面',
        timer;
    swal({
        title: "抱歉,未找到您的订单!",
        text: displayText.replace(/#1/, closeInSeconds),
        imageUrl: "../static/image/info1.png",
        timer: closeInSeconds * 1000,
        showCancelButton: true,
    }, function () {
        window.location.href = "/post_job"
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
data = JSON.parse(sessionStorage.getItem('Query_Unaccept_order'))
console.log(data);
if (data == undefined) {
    var result = "";
    result += `<div class="row">
    <div class="col-lg-8">
    <div class="row">
    <div class="col-md-6 col-sm-6">
    <div class="blog-card">
    <div class="blog-text">
    <ul>
    <li>
    <i class='bx bxs-user'></i>
    <span>not found</span>
    </li>
    <li>
    <i class='bx bx-calendar'></i>
    <span>not found</span>
    </li>
    </ul>
    <h3>
    <a href="blog-details.html">not found</a>
    </h3>
    <p>金额：not found</p>
    <p>类型：not found</p>
    <p>时间戳：not found</p>
    </div>
    </div>
    </div>
    </div>
    </div>
     </div>`
    document.getElementById("release_order").innerHTML = result;
    sorry_info()

} else {
    var result = "";
    data.forEach(e => {
        // ${e.title}
        result += `<div class="row">
    <div class="col-lg-8">
    <div class="row">
    <div class="col-md-6 col-sm-6">
    <div class="blog-card">
    <div class="blog-text">
    <ul>
    <li>
    <i class='bx bxs-user'></i>
    <span>${e.Add}</span>
    </li>
    <li>
    <i class='bx bx-calendar'></i>
    <span>${e.LaunchTime}</span>
    </li>
    </ul>
    <h3>
    <a>${e.Taskname}</a>
    </h3>
    <p>金额：${e.Amount}</p>
    <p>类型：${e.Category}</p>
    <p>接单状态：${e.State} </p>
    <p>时间戳：${e.Timestamp}</p>
    <a onclick="Cancle_order(this)" class="blog-btn">
    <button>取消</button>
    <i class='bx bx-plus bx-spin'></i>
    </a>
    </div>
    </div>
    </div>
    </div>
    </div>
     </div>`
    })

    document.getElementById("release_order").innerHTML = result;
}



function demo7() {
    var
        closeInSeconds = 5,
        displayText = ' #1 秒后将自动返回上一级',
        timer;

    swal({
        title: "取消成功!",
        text: displayText.replace(/#1/, closeInSeconds),
        imageUrl: "../static/image/check.png",
        timer: closeInSeconds * 1000,
        showCancelButton: true, //有这个就有取消按钮
        showconfirmButton: true,
    }, function () {
            window.location.href = "/self_resume_order"
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



 
 

function Cancle_order(e) {
    var Timestamp_new = e.parentElement.children[5].innerText
    Timestamp = Timestamp_new.slice(4,14)
    console.log(Timestamp);
    var account = e.parentElement.children[0].innerText
    console.log(account);
    console.log(account.slice(3,43));

    console.log(parseInt(Timestamp))
    var a=parseInt(Timestamp)
    // var b = a.slice(2,)

    $.ajax({
        method: "post",
        url: "http://localhost:8080/dapp/canceltask",
        data: { account: account, timestap: a},
        beforeSend: function () {
            swal({
              title: "正在取消中，请稍等几秒......",
              imageUrl: "../static/image/wait.png",
              showconfirmButton: true,
            })
        },
        success: function (data){
            console.log("success data", data);
            result = data.data
            console.log("result=====>", result);
            demo7()
        },
        error: function (data) {
            swal("OMG", "删除操作失败了!", "error");
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
}






























