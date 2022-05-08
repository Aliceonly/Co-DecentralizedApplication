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
    account= document.getElementById("Account").value ;
    sid document.getElementById("Sid").value
    document.getElementById("Sname").value=sname
    document.getElementById("Sage").value=sage
    document.getElementById("Tele").value=telephone
    document.getElementById("Major").value=major
    document.getElementById("Grade").value=grade
    console.log(sname)
    $.ajax({
        method: "post",
        url: "http://localhost:8080/dapp/change_user_info",
        data: { account:account, sid: sid,sname:sname, sage: sage, selephone: telephone, major: major, grade: grade},
        success: function (data) {
            if(data.data==1){
                swal("Good!", "修改成功", "success");
            }
        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
}

 