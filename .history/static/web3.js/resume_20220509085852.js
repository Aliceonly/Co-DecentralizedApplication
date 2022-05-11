data = JSON.parse(sessionStorage.getItem('QueryUser'));
console.log(data.data)
console.log(data.Grade);
console.log(data.Account);
document.getElementById("Account").innerHTML = data.Account;
document.getElementById("sid").innerHTML = data.Sid;
document.getElementById("Sname").innerHTML = data.Sname;
document.getElementById("Major").innerHTML = data.Major;
document.getElementById("Tele").innerHTML = data.Telephone;
document.getElementById("Sage").innerHTML = data.Sage;
document.getElementById("Grade").innerHTML = data.Grade;