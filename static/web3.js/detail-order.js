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

url = window.location.href
console.log("url========>", url)
// 使用正则表达式取出参数的值。
const queryURLParameter = (url) => {
    let regx = /([^&?=]+)=([^&?=]+)/g;
    let obj = {};
    url.replace(regx, (...args) => {
        if (obj[args[1]]) {
            obj[args[1]] = Array.isArray(obj[args[1]])
                ? obj[args[1]]
                : [obj[args[1]]];
            obj[args[1]].push(args[2]);
        } else {
            obj[args[1]] = args[2];
        }
    });
    return obj;
};
console.log("detail-order_Js_query=======>>>", queryURLParameter(url));
data = queryURLParameter(url)
// 
const list = [data];
const LaunchTime = list.map(item => item.LaunchTime);
document.getElementById("LaunchTime").innerHTML = LaunchTime;

const Taskname = list.map(item => item.Taskname);
document.getElementById("Taskname").innerHTML = decodeURI(Taskname);

const Amount = list.map(item => item.Amount);
document.getElementById("Amount").innerHTML = Amount;

const Category = list.map(item => item.Category);
document.getElementById("Category").innerHTML = decodeURI(Category);

const State = list.map(item => item.State);
document.getElementById("State").innerHTML = decodeURI(State);

const Add_1 = list.map(item => item.Add);
const Add = decodeURI(Add_1)
document.getElementById("Add").innerHTML = Add.slice(0, 35);

const Block = list.map(item => item.Block);
document.getElementById("Block").innerHTML = decodeURI(Block);

const Timestamp = list.map(item => item.Timestamp);
let Timestamp_new = decodeURI(Timestamp);
document.getElementById("Timestamp").innerHTML = Timestamp_new;
console.log(parseInt(Timestamp_new))
var a=parseInt(Timestamp_new)

 

let Account = window.sessionStorage.getItem("Global_Account");
function Confirmtask() {
    if (Add == Account) {
        console.log("Add===>", Add)
        console.log("Account=====>", Account)
        demo7()
    }else{
        $.ajax({
            method: "post",
            url: "http://localhost:8080/dapp/confirmtask",
            data: { account: Account, timestap:a},
            success: function (data) {
                console.log(data.data)
            },
            error: function (data) {
                console.log("error====>", error)
                console.log("error data===>", data)
            }
        })
         
    }
}

