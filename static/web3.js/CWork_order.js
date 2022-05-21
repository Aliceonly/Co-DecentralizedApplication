data = JSON.parse(sessionStorage.getItem('Query_CWork_order_detail'))
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
    document.getElementById("CWork_order_detail").innerHTML = result;
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
        <li>
        <i class='bx bx-briefcase'></i>
        ${e.Timestamp}
        </li>
        <li>
        <i class='bx bx-briefcase'></i>
        ${e.Amount}
        </li>
        <li>
        <i class='bx bx-location-plus'></i>
        ${e.State}
        </li>
        <li>
        <i class='bx bx-stopwatch'></i>
        ${e.LaunchTime}
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
    document.getElementById("CWork_order_detail").innerHTML = result;
}

 
