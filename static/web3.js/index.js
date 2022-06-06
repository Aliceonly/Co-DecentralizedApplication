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

function detaildata(e) {
    if (Account_status == "") {
        Login_status()
    } else {
        result = []
        Timestamp = e.parentElement.parentElement.parentElement.children[1].children[0].children[1].children[0].innerText
        var Timestamp1 = Timestamp.replace(/[^0-9]/ig, "");
        console.log(Timestamp1);
        $.ajax({
            method: "post",
            url: "http://localhost:8080/dapp/detailData",
            data: { timestamp: Timestamp1 },
            success: function (data) {
                console.log("success data", data);
                result = data.data
                // alert(data.data.Add)
                console.log("result=====>", result);
                window.location.href = "/Detail?Taskname=" + result.Taskname + '&LaunchTime=' + result.LaunchTime + '&Amount=' + result.Amount + '&Category=' + result.Category + '&State=' + result.State + '&Add=' + result.Add + '&Timestamp=' + result.Timestamp + '&Block=' + result.Block; //window.location.href跳转新页面

            },
            error: function (data) {
                console.log("error====>", error)
                console.log("error data===>", data)
            }
        })
    }
}

let a = document.getElementById("state").innerHTML
console.log("aaaaa",a);


// function Campus_order(){
    // let entry="校园跑腿类型";
    // $.ajax({
        // method: "post",
        // url: "http://localhost:8080/dapp/Campus_order_detail",
        // data: { entry: entry },
        // success: function (data) {
            // console.log("success data", data);
            // result = data.data
            // console.log("result=====>", result);
            // sessionStorage.setItem('Query_campus_order_detail', JSON.stringify(result))
            // window.location.href="/detail_qukuilySend"
        // },
        // error: function (data) {
            // console.log("error====>", error)
            // console.log("error data===>", data)
        // }
    // })
// }
// 
// function CWork_order(){
    // let entry="校园兼职类型";
    // $.ajax({
        // method: "post",
        // url: "http://localhost:8080/dapp/CWork_order_detail",
        // data: { entry: entry },
        // success: function (data) {
            // console.log("success data", data);
            // result = data.data
            // console.log("result=====>", result);
            // sessionStorage.setItem('Query_CWork_order_detail', JSON.stringify(result))
            // window.location.href="/detail_campus" 
        // },
        // error: function (data) {
            // console.log("error====>", error)
            // console.log("error data===>", data)
        // }
    // })
// }
// 
// 
// function Shared_order(){
    // let entry="共享服务类型";
    // $.ajax({
        // method: "post",
        // url: "http://localhost:8080/dapp/Shared_order_detail",
        // data: { entry: entry },
        // success: function (data) {
            // console.log("success data", data);
            // result = data.data
            // console.log("result=====>", result);
            // sessionStorage.setItem('Query_Shared_order_detail', JSON.stringify(result))
            // window.location.href="/detail_shared"
        // },
        // error: function (data) {
            // console.log("error====>", error)
            // console.log("error data===>", data)
        // }
    // })
// }


let Account = window.sessionStorage.getItem("Global_Account");
function collectOrder(e) {
    if (Account_status == "") {
        Login_status()
    } else {
        Timestamp = e.parentElement.parentElement.parentElement.children[1].children[0].children[1].children[0].innerText
        var Timestamp1 = Timestamp.replace(/[^0-9]/ig, "");
        console.log(Timestamp1);
        $.ajax({
            method: "post",
            url: "http://localhost:8080/dapp/collectOrder",
            data: { timestamp: Timestamp1 ,account:Account},
            success: function (data) {
                console.log("success data", data);
                result = data.data
                console.log("result=====>", result);
                swal("Good!", "收藏成功", "success");
            },
            error: function (data) {
                swal("OMG!", "收藏失败，请勿重复收藏", "error");
                console.log("error====>", error)
                console.log("error data===>", data)
            }
        })
        
    }
}




function CancleUser() {
    alert(45)
    console.log("CancleUser")
    $.ajax({
        method: "post",
        url: "http://localhost:8080/dapp/CancelUser",
        data: { account: Account },
        success: function (data) {
            console.log(data.data)
            Cancle_succ()
            $.ajax({
                method: "get",
                url: "/exit_account",
                success: function (data) {
                    console.log("目前账户情况：", data)
                },
                error: function (data) {
                    console.log("error====>", error)
                    console.log("error data===>", data)
                }
            })

        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
}


function exit_login() {
    alert("iui")
    console.log("exit");
    $.ajax({
        method: "get",
        url: "/exit_account",
        success: function (data) {
            console.log("目前账户情况：", data)
            demo7()
        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })

}


function GetBalcnce() {
    $.ajax({
        method: "get",
        url: "/GetBlance",
        success: function (data) {
            console.log("目前账户情况：", data.data)
            window.location.href="GetBalance_show";
            sessionStorage.setItem("balance",data.data);
        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
}




