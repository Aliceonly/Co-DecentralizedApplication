if (typeof window.ethereum !== 'undefined') {
  console.log('MetaMask is installed!')
  getAccount();
  ethereum.on('accountsChanged', function (accounts) {
    // Time to reload your interface with accounts[0]!
    console.log("现在的账户是:", accounts[0]);
    account = accounts[0];
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
function CreatTask() {
  var taskname = $("#exampleInput1").val();
  var tasktime = $("#exampleInput2").val();
  var tasktime2 = $("#exampleInput4").val();
  var taskmoney = $("#exampleInput3").val();
  var taskplace1 = $("#place option:selected").text();//地区
  var taskplace2 = $("#place2 option:selected").text();//主要地区
  var taskplace3 = $("#place3 option:selected").text();//工作类型
  var taskcontent = $('#exampleFormControlTextarea1').val();
  console.log(taskname, tasktime, taskmoney, tasktime2);
  console.log(taskplace1, taskplace2, taskplace3, taskcontent);

  $.ajax({
    method: "post",
    url: "http://localhost:8080/dapp/creatTask",
    data: { taskname: taskname, tasktime: tasktime + ":" + tasktime2, account: account, taskmoney: taskmoney, taskplace1: taskplace1 + taskplace2, taskplace3: taskplace3, taskcontent: taskcontent },
    success: function (data) {
      console.log("success data", data);
      console.log("成功");
      //  window.location.href ="http://localhost:8080/create_succ"
      swal({
        title: "发布成功",
        text:'您的订单时间戳是：<span style="color:red">'+ data.data +'（时间戳可用于订单查询）</span><br/><a style="color:#3b3bf4" href="/"> 点击返回首页查看</a><br/>10秒后自动关闭。',
        imageUrl: "../static/image/check.png",
        html: true,
        timer: 5000,
        showConfirmButton: false
      });
    },
    error: function (data) {
      console.log("error====>", error)
      console.log("error data===>", data)
      swal("OMG!", "发布失败请重试！", "error");
    }
  })
}