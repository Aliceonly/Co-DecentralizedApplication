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
    Account= document.getElementById("Account").innerHTML ;
    document.getElementById("Sid").value=Sid
    document.getElementById("Sname").value=Sname
    document.getElementById("Sage").value=Sage
    document.getElementById("Tele").value=Telephone
    document.getElementById("Major").value=Major
    document.getElementById("Grade").value=Grade
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

 