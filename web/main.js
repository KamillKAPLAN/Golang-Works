var xhr = new XMLHttpRequest();
var url = "http://localhost:3000/people"

xhr.responseType = ""; // ne dönemsi gerektiğini belirtiyor. (ARAŞTIR.)

xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
        var response = JSON.parse(xhr.responseText);
        myFunction(response);
        alert(response.ip);
    }
};
xhr.open("GET",url, true);
xhr.send();
function myFunction(arr) {
    var out = "";
    var i;
    for(i = 0; i < arr.length; i++) {
        out += 'Firstname: ' + arr[i].FirstName + ' ' + arr[i].LastName + '<br>';
    }
    document.getElementById("demo").innerHTML = out;
}
