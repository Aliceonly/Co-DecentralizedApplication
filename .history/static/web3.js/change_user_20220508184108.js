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
    document.getElementById("Account").value=Account ;
    document.getElementById("Sid").value=sid
    document.getElementById("Sname").value=sname
    document.getElementById("Sage").value=sage
    document.getElementById("Tele").value=telephone
    document.getElementById("Major").value=Major
    document.getElementById("Grade").value=Grade
    console.log
    $.ajax({
        method: "post",
        url: "http://localhost:8080/dapp/change_user_info",
        data: { Account: Account, Sid: Sid, Sname: Sname, Sage: Sage, Telephone: Telephone, Major: Major, Grade: Grade},
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

 