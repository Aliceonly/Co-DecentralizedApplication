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
document.getElementById("Taskname").innerHTML = decodeURI(Taskname) ;

const Amount = list.map(item => item.Amount);
document.getElementById("Amount").innerHTML = Amount;

const Category = list.map(item => item.Category);
document.getElementById("Category").innerHTML = decodeURI(Category);

const State = list.map(item => item.State);
document.getElementById("State").innerHTML = decodeURI(State);

const Add_1 = list.map(item => item.Add);
const Add = decodeURI(Add_1)
document.getElementById("Add").innerHTML = Add.slice(0,35);

const Block = list.map(item => item.Block);
document.getElementById("Block").innerHTML = decodeURI(Block);

const Timestamp = list.map(item => item.Timestamp);
document.getElementById("Timestamp").innerHTML = decodeURI(Timestamp);

