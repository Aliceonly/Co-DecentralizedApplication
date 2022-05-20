data = JSON.parse(sessionStorage.getItem('QueryUser'));
console.log(data);
Account = data.Account
Sid = data.Sid
Sname= data.Sname
Sage = data.Sage
Telephone = data.Telephone
Major = data.Major
Grade = data.Grade
document.getElementById("Account").innerHTML = Account;
document.getElementById("Sid").value=Sid
document.getElementById("Sname").value=Sname
document.getElementById("Sage").value=Sage
document.getElementById("Tele").value=Telephone
document.getElementById("Major").value=Major
document.getElementById("Grade").value=Grade

function change_user_info() {
    // account= $("Account").val()
    // if (account == undefined){
        account=document.getElementById("Account").innerHTML
        telephone = document.getElementById("Tele").value
    // }
    sid= $("Sid").val()
    if (sid == undefined){
        sid=document.getElementById("Sid").value
    }
    sname=$("#Sname").val()
    if (sname == undefined){
        sname=document.getElementById("Sname").value
    }
    sage=$("Sage").val()
    if (sage == undefined){
        sage=document.getElementById("Sage").value
    }
    telephone=$("Tele").val()
    if (telephone == undefined){
        telephone=document.getElementById("Tele").value
    }

    major=$("Major").val()
    if (major == undefined){
        major=document.getElementById("Major").value
    }
    grade=$("Grade").val()
    if (grade == undefined){
        grade=document.getElementById("Grade").value
    }
    console.log(sid,sname)
    $.ajax({
        method: "post",
        url: "http://localhost:8080/dapp/change_user_info",
        data: { account:account, sid: sid,sname:sname, sage: sage, selephone: telephone, major: major, grade: grade},
        success: function (data) {
            if(data.data==1){
                // swal("Good!", "修改成功", "success");
                swal({
                    title: "修改成功",
                    text: '<span style="color:red">点击</span><a style="color:#3b3bf4" href="/account">我的</a><span style="color:red">去查看个人信息吧！</span><br/>5秒后自动关闭。',
                    imageUrl: "../static/image/check.png",
                    html: true,
                    timer: 10000,
                    showConfirmButton: false
                });
            }
        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
}

 