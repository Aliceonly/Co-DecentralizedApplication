function Campus_order(){
    let entry="校园跑腿类型";
    $.ajax({
        method: "post",
        url: "http://localhost:8080/dapp/Campus_order_detail",
        data: { entry: entry },
        success: function (data) {
            console.log("success data", data);
            result = data.data
            console.log("result=====>", result);
            sessionStorage.setItem('Query_campus_order_detail', JSON.stringify(result))
            window.location.href="/detail_qukuilySend"
        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
}

function CWork_order(){
    let entry="校园兼职类型";
    $.ajax({
        method: "post",
        url: "http://localhost:8080/dapp/CWork_order_detail",
        data: { entry: entry },
        success: function (data) {
            console.log("success data", data);
            result = data.data
            console.log("result=====>", result);
            sessionStorage.setItem('Query_CWork_order_detail', JSON.stringify(result))
            window.location.href="/detail_campus" 
        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
}


function Shared_order(){
    let entry="共享服务类型";
    $.ajax({
        method: "post",
        url: "http://localhost:8080/dapp/Shared_order_detail",
        data: { entry: entry },
        success: function (data) {
            console.log("success data", data);
            result = data.data
            console.log("result=====>", result);
            sessionStorage.setItem('Query_Shared_order_detail', JSON.stringify(result))
            window.location.href="/detail_shared"
        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
}