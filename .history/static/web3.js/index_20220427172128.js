function detaildata(e) {
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
            window.location.href = "/Detail?Taskname=" + result.Taskname + '&LaunchTime=' + result.LaunchTime+ '&Amount=' + result.Amount+ '&Category=' + result.Category+ '&State=' + result.State+ '&Add=' + result.Add; //window.location.href跳转新页面

        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
}