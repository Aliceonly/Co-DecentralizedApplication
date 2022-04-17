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
  var taskcontent=$('#exampleFormControlTextarea1').val();
  console.log(taskname,tasktime,taskmoney,tasktime2);
  console.log(taskplace1,taskplace2,taskplace3,taskcontent);
  
  $.ajax({
 method:"post",
 url:"http://localhost:8080/post_job/creatTask",
 data:{taskname:taskname,tasktime:tasktime+tasktime2,account:account,taskmoney:taskmoney,taskplace1:taskplace1+taskplace2,taskplace3:taskplace3,taskcontent:taskcontent},
 success:function(data){
     console.log(data);
    //  document.getElementById("hash").innerHTML=data;
     console.log("完成");
 },
 error: function(data){
     console.log(data)
 }
}) 
}