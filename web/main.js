var ids = document.getElementsByClassName('id');
var firsts = document.getElementsByClassName('first');
var lasts = document.getElementsByClassName('last');
var removes = document.getElementsByClassName('remove');
var req = document.getElementById("req");

var xhr = new XMLHttpRequest();
var url = "http://localhost:3000/people"

xhr.responseType = ""; // ne dönemsi gerektiğini belirtiyor. (ARAŞTIR.)

xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
        console.log(xhr);
        var response = JSON.parse(xhr.responseText);
        show(response);
    }
};
  xhr.open("GET",url, true);
  xhr.send();

function show(arr) {
  for(var i = 0; i < arr.length; i++) {
    var j = i+1;
    ids[i].innerHTML = arr[i].ID;
    firsts[i].innerHTML = arr[i].FirstName;
    lasts[i].innerHTML = arr[i].LastName;
    removes[i].innerHTML = "<button type='button' name='button' style='margin-left: 10px;' onclick='deleate(" + j + ")'>Sil</button>" +'<br>';
  }
}

function deleate(num) {
  var new_url = url + "/" + num;
  xhr.open("DELETE", new_url, true);
  // console.log(new_url);
  xhr.send();
}

function postA(){
  var data = {}
  data.FirstName = "Kamil";
  data.LastName = "KAPLAN";
  var json = JSON.stringify(data);
  var i = ids.length;
  xhr.open("POST", url + "/" + (i+1) , true);
  xhr.send(json);
}

function puthA(){
  var data = {};
  data.firstname = "John2";
  data.lastname  = "Snow2";
  var json = JSON.stringify(data)

  var xhr = new XMLHttpRequest();
  xhr.open("PUT", url+'/12', true);
  xhr.onload = function () {
  	var users = JSON.parse(xhr.responseText);
  	if (xhr.readyState == 4 && xhr.status == "200") {
  		console.table(users);
  	} else {
  		console.error(users);
  	}
  }
  xhr.send(json);

}
