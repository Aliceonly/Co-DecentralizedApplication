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
 var taskname=$("#exampleInput1").val();
  var tasktime=$("#exampleInput2").val();
  var tasktime2=$("#exampleInput4").val();
  var taskmoney=$("#exampleInput3").val();
  var taskplace1=$("#place option:selected").text();
  var taskplace2=$("#place2 option:selected").text();
  var taskplace3=$("#place3 option:selected").text();
  var taskcontent
  // console.log(taskname,tasktime,taskmoney,taskplace,tasktime2);
//   $.ajax({
//  method:"post",
//  url:"http://localhost:3000/form_validation/gethash",
//  data:{pname1:pname1,pid1:pid1,pctx1:pctx1,account:account},
//  success:function(data){
//      console.log(data);
//      document.getElementById("hash").innerHTML=data;
//      console.log("完成");

//   //    window.location.href="http://localhost:3000/basic_table";
//  },
//  error: function(data){
//      console.log(data)
//  }
// }) 
}