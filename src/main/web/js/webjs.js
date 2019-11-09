function submbtn(){

}
function validity(){
  var flag = true;
  var val = document.forms["userform"]["uname"].value;
  if(val === ""){
    alert('Username column empty!');
    flag = false;
  }
  return flag;
}
