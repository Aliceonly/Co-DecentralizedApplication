var apply_button = document.querySelector("#apply_button");
var time_out = document.querySelector("#time_out");
let time = window.sessionStorage.getItem("times")
if (time==0){
    let times=3
    time_out.style.display = "none";
    document.getElementById("apply_time").innerHTML = times;
}if (time == 1){
    let times=2
    time_out.style.display = "none";
    document.getElementById("apply_time").innerHTML = times;
}if (time == 2){
    let times=1
    time_out.style.display = "none";
    document.getElementById("apply_time").innerHTML = times;
}if (time == 3){
    let times=0
    apply_button.style.display = "none";
    document.getElementById("apply_time").innerHTML = times;
}

function apply(){
    $.ajax({
        method: "get",
        url: "/Apply_5",
        success: function (data) {
            console.log(data.data)
        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
}
 
 
 
 
 