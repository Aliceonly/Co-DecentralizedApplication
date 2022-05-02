function Deletedata(e){
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
           
        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
}
}