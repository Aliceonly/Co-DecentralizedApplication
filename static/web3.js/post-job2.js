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
  var account = window.sessionStorage.getItem("Global_Account");
  let balance = window.sessionStorage.getItem("balance");
  let residue = balance - taskmoney;
  if (balance < taskmoney) {
    swal({
      title: "余额不足",
      text: '<span style="color:red">当前余额：</span>' + balance + '<br/><span style="color:red">发布订单后的余额：</span>' + residue + ' <br/>',
      imageUrl: "../static/image/info1.png",
      html: true,
      showCancelButton: true,
      showconfirmButton: true,
    })
  } else {
    $.ajax({
      method: "post",
      url: "http://localhost:8080/dapp/creatTask",
      data: { taskname: taskname, tasktime: tasktime + ":" + tasktime2, account: account, taskmoney: taskmoney, taskplace1: taskplace1 + taskplace2, taskplace3: taskplace3, taskcontent: taskcontent },
      beforeSend: function () {
        swal({
          title: "正在发布中，请稍等几秒.....",
          text: '<span style="color:red">请不要离开此页面、直至下一个弹框出现</br>否则交易可能失败!</sapn>',
          html: true,
          imageUrl: "../static/image/wait.png",
          showconfirmButton: true,
        })
      },
      success: function (data) {
        console.log("success data", data);
        console.log("成功");
        console.log(data.data);
        //截取逗号前的字符串
        let result = data.data
        console.log(result.split(',')[0]);
        swal({
          title: "发布成功",
          text: '该订单签名是：<span style="color:red"><br/>' + result + '<br/>（签名将应用于后期确认订单操作，请务必复制牢记）</span>',
          imageUrl: "../static/image/check.png",
          html: true,
          showCancelButton: true,
          showconfirmButton: true,
        });
      },
      error: function (data) {
        console.log("error====>", error)
        console.log("error data===>", data)
        swal("OMG!", "发布失败请重试！", "error");
      }
    })
  }
}