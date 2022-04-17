if (typeof window.ethereum !== 'undefined') {
    console.log('MetaMask is installed!')
      getAccount();
      ethereum.on('accountsChanged', function (accounts) {
          // Time to reload your interface with accounts[0]!
          console.log("现在的账户是:", accounts[0]);
          account=accounts[0];
          $("#address").html(accounts[0]);
      });
  } else {
    alert('plase install the MetaMask')
  }
  var account;
  async function getAccount() {
      yghaccounts = await ethereum.request({ method: 'eth_requestAccounts' });
      account = yghaccounts[0];
     $("#address").html(account);
     console.log(account)
  }
//   console.log(1)
function CreatTask(){
 var taskname=$("#ppname").val();
  var pid1=$("#ppid").val();
  var pctx1=$("#ppctx").val();
  console.log(pname1);
  console.log(pid1);
  console.log(pctx1);
  $.ajax({
 method:"post",
 url:"http://localhost:3000/form_validation/gethash",
 data:{pname1:pname1,pid1:pid1,pctx1:pctx1,account:account},
 success:function(data){
     console.log(data);
     document.getElementById("hash").innerHTML=data;
     console.log("完成");

  //    window.location.href="http://localhost:3000/basic_table";
 },
 error: function(data){
     console.log(data)
 }
}) 
}