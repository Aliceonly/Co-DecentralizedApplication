function delete(e){
    Timestamp = e.parentElement.parentElement.parentElement.children[1].children[0].children[1].children[0].innerText
    var Timestamp1 = Timestamp.replace(/[^0-9]/ig, "");
    console.log(Timestamp1);
}