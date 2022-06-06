let balance = window.sessionStorage.getItem("balance");
document.getElementById("user_balcne").innerHTML = balance;
function Applyfor(){
    $.ajax({
        method: "get",
        url: "/Frequency",
        success: function (data) {
            console.log(data.data)
            window.sessionStorage.setItem("times",data.data); 
            window.location.href="/Apply"
        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
    
}
