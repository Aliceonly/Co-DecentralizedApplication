data = JSON.parse(sessionStorage.getItem('Query_campus_order_detail'))
console.log(data);
if (data == undefined) {
    var result = "";
    result += `<div class="col-lg-12">
    <div class="job-card-two">
    <div class="row align-items-center">
    <div class="col-md-1">
    </div>
    <div class="col-md-8">
    <div class="job-info">
    <h3>
    <a>Not found </a>
    </h3>
    <ul>
    <li>
    <i class='bx bx-briefcase'></i>
    Not found 
    </li>
    <li>
    <i class='bx bx-briefcase'></i>
    Not found 
    </li>
    <li>
    <i class='bx bx-location-plus'></i>
    Not found 
    </li>
    <li>
    <i class='bx bx-stopwatch'></i>
    Not found 
    </li>
    </ul>
    <span>Not found </span>
    </div>
    </div>
    <div class="col-md-3">
     <div class="theme-btn text-end">
    <a class="default-btn">
    Not found 
    </a>
    </div>
    </div>
    </div>
    </div>
    </div>
    `
    document.getElementById("Campus_order_detail").innerHTML = result;
}else {
    var result = "";
    data.forEach(e => {
        // ${e.title}
        result += `<div class="col-lg-12">
        <div class="job-card-two">
        <div class="row align-items-center">
        <div class="col-md-1">
        </div>
        <div class="col-md-8">
        <div class="job-info">
        <h3>
        <a>${e.Taskname} </a>
        </h3>
        <ul>
        <li>时间戳：${e.Timestamp}</li>
        <li>金额：${e.Amount}</li>
        <li>接单状态：${e.State}
        </li>
        <li>订单时间：${e.LaunchTime}
        </li>
        </ul>
        <span>${e.Category}</span>
        </div>
        </div>
        <div class="col-md-3">
         <div class="theme-btn text-end">
        <a class="default-btn">
        查看更多
        </a>
        </div>
        </div>
        </div>
        </div>
        </div>
        `
    })
    document.getElementById("Campus_order_detail").innerHTML = result;
}

 
