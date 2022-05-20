function Deletedata(e) {
    Timestamp = e.parentElement.parentElement.parentElement.children[1].children[0].children[1].children[0].innerText
    var Timestamp1 = Timestamp.replace(/[^0-9]/ig, "");
    console.log(Timestamp1);
    swal({
        title: "您确定要删除吗？",
        text: "您确定要删除这条数据？",
        type: "warning",
        showCancelButton: true,
        closeOnConfirm: false,
        confirmButtonText: "是的，我要删除",
        confirmButtonColor: "#ec6c62"
    }, function () {
        $.ajax({
            method: "post",
            url: "http://localhost:8080/dapp/deleteData",
            data: { timestamp: Timestamp1 },
            success: function () {
                // swal("操作成功!", "已成功删除数据！", "success");
                swal({
                    title: "操作成功",
                    text: '<span style="color:red">已成功删除数据！</span><a style="color:#3b3bf4" href="/all_order">点击返回查看</a>。<br/>5秒后自动关闭。',
                    imageUrl: "../static/image/check.png",
                    html: true,
                    timer: 5000,
                    showConfirmButton: false
                  });

            },
            error: function (data) {
                swal("OMG", "删除操作失败了!", "error");
                console.log("error====>", error)
                console.log("error data===>", data)
            }
        })
    })
}
